package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Kiln struct {
	BaseUrl string `yaml:"baseURL"`
	Token   string `yaml:"token"`
}

type Monitoring struct {
	Interval time.Duration `yaml:"intervam"`
	AccountIDS []string `yaml:"account_ids"`
}

type Alerts struct {
	Kind string `yaml:"kind"`
	DiscordWebhook string `yaml:"discord_webhook"`
}

type Rules struct {
	MinUptime float64 `yaml:"min_uptime"`
	MaxInactivePeriods int `yaml:"max_inactive_periods"`
	RewardDropThresholdPct float64 `yaml:"reward_drop_threshold_pct"`
}

type Config struct {
	Kiln Kiln `yaml:"kiln"`
	Monitoring Monitoring `yaml:"monitoring"`
	Alerts Alerts `yaml:"alerts"`
	Rules Rules `yaml:"rules"`
}

func Load(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	// Overrides via env (12-factor friendly)
	if v := os.Getenv("KILN_API_TOKEN"); v != "" {
		cfg.Kiln.Token = v
	}
	if v:= os.Getenv("DISCORD_WEBHOOK_URL"); v!= ""{
		cfg.Alerts.DiscordWebhook = v
	}
	if v:= os.Getenv("ALERT_KIND"); v!= "" {
		cfg.Alerts.Kind = v
	}

	return &cfg, nil
}