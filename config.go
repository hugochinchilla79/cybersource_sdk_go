package cybersource

import (
	"fmt"
	"os"
)

type Environment string

const (
	EnvSandbox    Environment = "sandbox"
	EnvProduction Environment = "production"
)

type Config struct {
	MerchantID string
	ApiKey     string
	SecretKey  string

	BaseURL string

	Env Environment
}

func (c Config) Validate() error {
	if c.MerchantID == "" {
		return fmt.Errorf("MerchantID es requerido")
	}
	if c.ApiKey == "" {
		return fmt.Errorf("ApiKey es requerido")
	}
	if c.SecretKey == "" {
		return fmt.Errorf("SecretKey es requerido")
	}
	if c.BaseURL == "" {
		return fmt.Errorf("BaseURL es requerido")
	}
	return nil
}

func DefaultBaseURL(env Environment) string {
	switch env {
	case EnvProduction:
		return "https://api.cybersource.com"
	default:
		return "https://apitest.cybersource.com"
	}
}

func LoadConfigFromEnv() (Config, error) {
	env := Environment(os.Getenv("CYBS_ENV"))
	if env == "" {
		env = EnvSandbox
	}

	baseURL := os.Getenv("CYBS_BASE_URL")
	if baseURL == "" {
		baseURL = DefaultBaseURL(env)
	}

	cfg := Config{
		MerchantID: os.Getenv("CYBS_MERCHANT_ID"),
		ApiKey:     os.Getenv("CYBS_API_KEY"),
		SecretKey:  os.Getenv("CYBS_SECRET_KEY"),
		BaseURL:    baseURL,
		Env:        env,
	}

	return cfg, cfg.Validate()
}
