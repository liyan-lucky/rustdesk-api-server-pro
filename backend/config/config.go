package config

import (
	"os"
	"path"
	"rustdesk-api-server-pro/util"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	DebugMode  bool        `yaml:"debugMode"`
	Db         *DbConfig   `yaml:"db"`
	SignKey    string      `yaml:"signKey"`
	HttpConfig *HttpConfig `yaml:"httpConfig"`
	SmtpConfig *SmtpConfig `yaml:"smtpConfig"`
	JobsConfig *JobsConfig `yaml:"jobsConfig"`
	OIDC       *OIDCConfig `yaml:"oidc"`
}

type DbConfig struct {
	Driver   string `yaml:"driver"`
	Dsn      string `yaml:"dsn"`
	TimeZone string `yaml:"timeZone"`
	ShowSql  bool   `yaml:"showSql"`
}

type HttpConfig struct {
	PrintRequestLog bool   `yaml:"printRequestLog"`
	Port            string `yaml:"port"`
	StaticDir       string `yaml:"staticdir"`
}

type SmtpConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Encryption string `yaml:"encryption"` // none ssl/tls starttls
	From       string `yaml:"from"`
}

type DeviceCheckJob struct {
	Duration int `yaml:"duration"`
}

type JobsConfig struct {
	DeviceCheckJob *DeviceCheckJob `yaml:"deviceCheckJob"`
}

type OIDCConfig struct {
	Enabled             bool     `yaml:"enabled"`
	ProviderName        string   `yaml:"providerName"`
	Issuer              string   `yaml:"issuer"`
	ClientID            string   `yaml:"clientId"`
	ClientSecret        string   `yaml:"clientSecret"`
	RedirectURL         string   `yaml:"redirectUrl"`
	Scopes              []string `yaml:"scopes"`
	BindByEmail         bool     `yaml:"bindByEmail"`
	AutoCreateAdmin     bool     `yaml:"autoCreateAdmin"`
	StateTTLSeconds     int      `yaml:"stateTtlSeconds"`
	TicketTTLSeconds    int      `yaml:"ticketTtlSeconds"`
	SuccessRedirect     string   `yaml:"successRedirect"`
	FailureRedirect     string   `yaml:"failureRedirect"`
	SubjectClaim        string   `yaml:"subjectClaim"`
	EmailClaim          string   `yaml:"emailClaim"`
	NameClaim           string   `yaml:"nameClaim"`
	PictureClaim        string   `yaml:"pictureClaim"`
	Prompt              string   `yaml:"prompt"`
	AllowedEmailDomains []string `yaml:"allowedEmailDomains"`
}

var (
	wd, _    = os.Getwd()
	yamlFile = path.Join(wd, "server.yaml")
)

func GetDefaultServerConfig() *ServerConfig {
	return &ServerConfig{
		DebugMode: false,
		Db: &DbConfig{
			Driver:   "sqlite",
			Dsn:      "./server.db",
			ShowSql:  true,
			TimeZone: "Asia/Shanghai",
		},
		HttpConfig: &HttpConfig{
			Port:      ":8080",
			StaticDir: "dist",
		},
		SignKey: util.RandomString(32),
		JobsConfig: &JobsConfig{
			DeviceCheckJob: &DeviceCheckJob{
				Duration: 30,
			},
		},
		OIDC: &OIDCConfig{
			Enabled:          false,
			ProviderName:     "oidc",
			Scopes:           []string{"openid", "profile", "email"},
			BindByEmail:      true,
			AutoCreateAdmin:  false,
			StateTTLSeconds:  180,
			TicketTTLSeconds: 180,
			SuccessRedirect:  "/login",
			FailureRedirect:  "/login",
			SubjectClaim:     "sub",
			EmailClaim:       "email",
			NameClaim:        "name",
			PictureClaim:     "picture",
		},
	}
}

func GetServerConfig() *ServerConfig {
	cfg := GetDefaultServerConfig()
	bytes, err := os.ReadFile(yamlFile)
	if err != nil {
		WriteServerConfig(cfg)
		return cfg
	}

	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		WriteServerConfig(cfg)
		return cfg
	}
	return cfg
}

func WriteServerConfig(cfg *ServerConfig) {
	bytes, _ := yaml.Marshal(cfg)
	_ = os.WriteFile(yamlFile, bytes, 0755)
}
