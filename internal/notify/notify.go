package notify

import "github.com/daywon03/Kiln-monitoring-microservice/internal/rules"

// Notifier interface defines how alerts are sent
type Notifier interface {
    Send(finding rules.Finding) error
}