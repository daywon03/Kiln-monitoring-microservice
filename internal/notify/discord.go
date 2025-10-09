package notify

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    
    "github.com/daywon03/Kiln-monitoring-microservice/internal/rules"
)

// DiscordNotifier sends alerts to Discord via webhook
type DiscordNotifier struct {
    webhookURL string
    client     *http.Client
}

// DiscordMessage represents a Discord webhook message
type DiscordMessage struct {
    Content string `json:"content"`
}

// NewDiscordNotifier creates a new Discord notifier
func NewDiscordNotifier(webhookURL string) *DiscordNotifier {
    return &DiscordNotifier{
        webhookURL: webhookURL,
        client:     &http.Client{Timeout: 10 * time.Second},
    }
}

// Send sends the finding to Discord
func (n *DiscordNotifier) Send(finding rules.Finding) error {
    // Format message with emoji based on severity
    emoji := "ℹ️"
    switch finding.Severity {
    case rules.SeverityWarn:
        emoji = "⚠️"
    case rules.SeverityCrit:
        emoji = "🚨"
    }

    content := fmt.Sprintf("%s **%s** | %s | Account: `%s`",
        emoji,
        finding.Severity,
        finding.Message,
        finding.Context["account_id"],
    )

    message := DiscordMessage{Content: content}
    
    jsonData, err := json.Marshal(message)
    if err != nil {
        return fmt.Errorf("failed to marshal Discord message: %w", err)
    }

    req, err := http.NewRequest("POST", n.webhookURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return fmt.Errorf("failed to create Discord request: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := n.client.Do(req)
    if err != nil {
        return fmt.Errorf("failed to send Discord webhook: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return fmt.Errorf("Discord webhook returned status: %d", resp.StatusCode)
    }

    return nil
}
