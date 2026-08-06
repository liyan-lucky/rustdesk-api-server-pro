package config

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"rustdesk-api-server-pro/util"

	"gopkg.in/yaml.v3"
)

var serverConfigWriteMu sync.Mutex

type ServerConfig struct {
	DebugMode       bool                `yaml:"debugMode"`
	Db              *DbConfig           `yaml:"db"`
	SignKey         string              `yaml:"signKey"`
	HttpConfig      *HttpConfig         `yaml:"httpConfig"`
	SmtpConfig      *SmtpConfig         `yaml:"smtpConfig"`
	JobsConfig      *JobsConfig         `yaml:"jobsConfig"`
	OIDC            *OIDCConfig         `yaml:"oidc"`
	OAuth           *OAuthConfig        `yaml:"oauth"`
	RustdeskServer  *RustdeskServerConfig `yaml:"rustdeskServer"`
}

type RustdeskServerConfig struct {
	IdServer    string `yaml:"idServer"`
	RelayServer string `yaml:"relayServer"`
	ApiServer   string `yaml:"apiServer"`
	Key         string `yaml:"key"`
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

type OAuthConfig struct {
	Providers []OAuthProviderConfig `yaml:"providers"`
}

type OAuthProviderConfig struct {
	Type                  string   `yaml:"type"`
	Name                  string   `yaml:"name"`
	DisplayName           string   `yaml:"displayName"`
	Enabled               bool     `yaml:"enabled"`
	Issuer                string   `yaml:"issuer"`
	AuthorizationEndpoint string   `yaml:"authorizationEndpoint"`
	TokenEndpoint         string   `yaml:"tokenEndpoint"`
	UserinfoEndpoint      string   `yaml:"userinfoEndpoint"`
	JWKSURI               string   `yaml:"jwksUri"`
	RedirectURL           string   `yaml:"redirectUrl"`
	ClientID              string   `yaml:"clientId"`
	ClientSecret          string   `yaml:"clientSecret"`
	TeamID                string   `yaml:"teamId"`
	KeyID                 string   `yaml:"keyId"`
	PrivateKey            string   `yaml:"privateKey"`
	Scopes                []string `yaml:"scopes"`
	BindByEmail           bool     `yaml:"bindByEmail"`
	AutoCreateAdmin       bool     `yaml:"autoCreateAdmin"`
	AccountRole           string   `yaml:"accountRole"`
	AutoCreateUser        bool     `yaml:"autoCreateUser"`
	StateTTLSeconds       int      `yaml:"stateTtlSeconds"`
	TicketTTLSeconds      int      `yaml:"ticketTtlSeconds"`
	SuccessRedirect       string   `yaml:"successRedirect"`
	FailureRedirect       string   `yaml:"failureRedirect"`
	SubjectClaim          string   `yaml:"subjectClaim"`
	EmailClaim            string   `yaml:"emailClaim"`
	NameClaim             string   `yaml:"nameClaim"`
	PictureClaim          string   `yaml:"pictureClaim"`
	Prompt                string   `yaml:"prompt"`
	AllowedEmailDomains   []string `yaml:"allowedEmailDomains"`
}

var (
	wd, _    = os.Getwd()
	yamlFile = filepath.Join(wd, "server.yaml")
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
			StaticDir: "/app/dist",
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
		OAuth: &OAuthConfig{
			Providers: []OAuthProviderConfig{},
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
		return cfg
	}

	if envPort := os.Getenv("PORT"); envPort != "" {
		if !strings.HasPrefix(envPort, ":") {
			envPort = ":" + envPort
		}
		cfg.HttpConfig.Port = envPort
	}

	return cfg
}

func WriteServerConfig(cfg *ServerConfig) {
	_ = SaveServerConfig(cfg)
}

// SaveServerConfig persists the complete configuration atomically. Callers that
// expose configuration editing should use this variant so write failures are
// visible to the administrator instead of being silently ignored.
func SaveServerConfig(cfg *ServerConfig) error {
	serverConfigWriteMu.Lock()
	defer serverConfigWriteMu.Unlock()

	bytes, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	dir := filepath.Dir(yamlFile)
	if err = os.MkdirAll(dir, 0750); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(dir, ".server-*.yaml")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)
	if err = tmp.Chmod(0600); err == nil {
		_, err = tmp.Write(bytes)
	}
	if closeErr := tmp.Close(); err == nil {
		err = closeErr
	}
	if err != nil {
		return err
	}
	return os.Rename(tmpName, yamlFile)
}

func IsUnsafeSignKey(signKey string) bool {
	key := strings.TrimSpace(signKey)
	if len(key) < 32 {
		return true
	}
	unsafeKeys := map[string]struct{}{
		"sercrethatmaycontainch@r$32chars":     {},
		"CHANGE_ME_TO_A_RANDOM_32_BYTE_SECRET": {},
		"please-change-this-sign-key":          {},
	}
	_, unsafe := unsafeKeys[key]
	return unsafe
}

func (cfg *ServerConfig) OAuthProviders() []OAuthProviderConfig {
	providers := []OAuthProviderConfig{}
	seen := map[string]struct{}{}

	if cfg != nil && cfg.OIDC != nil {
		legacy := OAuthProviderConfig{
			Type:                "oidc",
			Name:                cfg.OIDC.ProviderName,
			DisplayName:         "OIDC",
			Enabled:             cfg.OIDC.Enabled,
			Issuer:              cfg.OIDC.Issuer,
			RedirectURL:         cfg.OIDC.RedirectURL,
			ClientID:            cfg.OIDC.ClientID,
			ClientSecret:        cfg.OIDC.ClientSecret,
			Scopes:              cfg.OIDC.Scopes,
			BindByEmail:         cfg.OIDC.BindByEmail,
			AutoCreateAdmin:     cfg.OIDC.AutoCreateAdmin,
			AccountRole:         "admin",
			StateTTLSeconds:     cfg.OIDC.StateTTLSeconds,
			TicketTTLSeconds:    cfg.OIDC.TicketTTLSeconds,
			SuccessRedirect:     cfg.OIDC.SuccessRedirect,
			FailureRedirect:     cfg.OIDC.FailureRedirect,
			SubjectClaim:        cfg.OIDC.SubjectClaim,
			EmailClaim:          cfg.OIDC.EmailClaim,
			NameClaim:           cfg.OIDC.NameClaim,
			PictureClaim:        cfg.OIDC.PictureClaim,
			Prompt:              cfg.OIDC.Prompt,
			AllowedEmailDomains: cfg.OIDC.AllowedEmailDomains,
		}
		if legacy.Name == "" {
			legacy.Name = "oidc"
		}
		providers = append(providers, legacy)
		seen[legacy.Name] = struct{}{}
	}

	if cfg != nil && cfg.OAuth != nil {
		for _, provider := range cfg.OAuth.Providers {
			name := strings.TrimSpace(provider.Name)
			if name == "" {
				name = strings.TrimSpace(provider.Type)
			}
			if name == "" {
				name = "oidc"
			}
			if _, ok := seen[name]; ok {
				continue
			}
			provider.Name = name
			provider.AccountRole = strings.ToLower(strings.TrimSpace(provider.AccountRole))
			if provider.AccountRole != "user" {
				provider.AccountRole = "admin"
			}
			providers = append(providers, provider)
			seen[name] = struct{}{}
		}
	}
	return providers
}
