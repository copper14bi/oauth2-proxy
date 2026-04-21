package options

import (
	"fmt"
	"net/url"
	"time"
)

// Options holds all configuration for the oauth2-proxy
type Options struct {
	// HTTP server options
	HTTPAddress  string `flag:"http-address" cfg:"http_address" env:"OAUTH2_PROXY_HTTP_ADDRESS"`
	HTTPSAddress string `flag:"https-address" cfg:"https_address" env:"OAUTH2_PROXY_HTTPS_ADDRESS"`

	// TLS options
	TLSCertFile string `flag:"tls-cert-file" cfg:"tls_cert_file" env:"OAUTH2_PROXY_TLS_CERT_FILE"`
	TLSKeyFile  string `flag:"tls-key-file" cfg:"tls_key_file" env:"OAUTH2_PROXY_TLS_KEY_FILE"`

	// Upstream options
	UpstreamURLs []string `flag:"upstream" cfg:"upstreams" env:"OAUTH2_PROXY_UPSTREAMS"`

	// OAuth2 provider options
	Provider        string `flag:"provider" cfg:"provider" env:"OAUTH2_PROXY_PROVIDER"`
	ProviderName    string `flag:"provider-display-name" cfg:"provider_display_name" env:"OAUTH2_PROXY_PROVIDER_DISPLAY_NAME"`
	ClientID        string `flag:"client-id" cfg:"client_id" env:"OAUTH2_PROXY_CLIENT_ID"`
	ClientSecret    string `flag:"client-secret" cfg:"client_secret" env:"OAUTH2_PROXY_CLIENT_SECRET"`
	ClientSecretFile string `flag:"client-secret-file" cfg:"client_secret_file" env:"OAUTH2_PROXY_CLIENT_SECRET_FILE"`

	// OAuth2 URLs
	LoginURL    string `flag:"login-url" cfg:"login_url" env:"OAUTH2_PROXY_LOGIN_URL"`
	RedeemURL   string `flag:"redeem-url" cfg:"redeem_url" env:"OAUTH2_PROXY_REDEEM_URL"`
	ProfileURL  string `flag:"profile-url" cfg:"profile_url" env:"OAUTH2_PROXY_PROFILE_URL"`
	ValidateURL string `flag:"validate-url" cfg:"validate_url" env:"OAUTH2_PROXY_VALIDATE_URL"`

	// Cookie options
	CookieName     string        `flag:"cookie-name" cfg:"cookie_name" env:"OAUTH2_PROXY_COOKIE_NAME"`
	CookieSecret   string        `flag:"cookie-secret" cfg:"cookie_secret" env:"OAUTH2_PROXY_COOKIE_SECRET"`
	CookieDomains  []string      `flag:"cookie-domain" cfg:"cookie_domains" env:"OAUTH2_PROXY_COOKIE_DOMAINS"`
	CookiePath     string        `flag:"cookie-path" cfg:"cookie_path" env:"OAUTH2_PROXY_COOKIE_PATH"`
	CookieExpire   time.Duration `flag:"cookie-expire" cfg:"cookie_expire" env:"OAUTH2_PROXY_COOKIE_EXPIRE"`
	CookieRefresh  time.Duration `flag:"cookie-refresh" cfg:"cookie_refresh" env:"OAUTH2_PROXY_COOKIE_REFRESH"`
	CookieSecure   bool          `flag:"cookie-secure" cfg:"cookie_secure" env:"OAUTH2_PROXY_COOKIE_SECURE"`
	CookieHTTPOnly bool          `flag:"cookie-httponly" cfg:"cookie_httponly" env:"OAUTH2_PROXY_COOKIE_HTTPONLY"`
	CookieSameSite string        `flag:"cookie-samesite" cfg:"cookie_samesite" env:"OAUTH2_PROXY_COOKIE_SAMESITE"`

	// Session options
	SessionStoreType string `flag:"session-store-type" cfg:"session_store_type" env:"OAUTH2_PROXY_SESSION_STORE_TYPE"`

	// Email / access control
	EmailDomains      []string `flag:"email-domain" cfg:"email_domains" env:"OAUTH2_PROXY_EMAIL_DOMAINS"`
	AllowedGroups     []string `flag:"allowed-group" cfg:"allowed_groups" env:"OAUTH2_PROXY_ALLOWED_GROUPS"`
	HtpasswdFile      string   `flag:"htpasswd-file" cfg:"htpasswd_file" env:"OAUTH2_PROXY_HTPASSWD_FILE"`
	HtpasswdUserGroup []string `flag:"htpasswd-user-group" cfg:"htpasswd_user_groups" env:"OAUTH2_PROXY_HTPASSWD_USER_GROUPS"`

	// Proxy behavior
	ReverseProxy      bool     `flag:"reverse-proxy" cfg:"reverse_proxy" env:"OAUTH2_PROXY_REVERSE_PROXY"`
	RealClientIPHeader string   `flag:"real-client-ip-header" cfg:"real_client_ip_header" env:"OAUTH2_PROXY_REAL_CLIENT_IP_HEADER"`
	WhitelistDomains  []string `flag:"whitelist-domain" cfg:"whitelist_domains" env:"OAUTH2_PROXY_WHITELIST_DOMAINS"`

	// Redirect URL
	RedirectURL string `flag:"redirect-url" cfg:"redirect_url" env:"OAUTH2_PROXY_REDIRECT_URL"`

	// Logging
	LoggingFilename       string `flag:"logging-filename" cfg:"logging_filename" env:"OAUTH2_PROXY_LOGGING_FILENAME"`
	StandardLogging       bool   `flag:"standard-logging" cfg:"standard_logging" env:"OAUTH2_PROXY_STANDARD_LOGGING"`
	RequestLogging        bool   `flag:"request-logging" cfg:"request_logging" env:"OAUTH2_PROXY_REQUEST_LOGGING"`
	AuthLogging           bool   `flag:"auth-logging" cfg:"auth_logging" env:"OAUTH2_PROXY_AUTH_LOGGING"`
}

// NewOptions returns a new Options struct with sensible defaults
func NewOptions() *Options {
	return &Options{
		HTTPAddress:        "127.0.0.1:4180",
		HTTPSAddress:       ":443",
		Provider:           "google",
		CookieName:         "_oauth2_proxy",
		CookiePath:         "/",
		CookieExpire:       168 * time.Hour,
		CookieRefresh:      0,
		CookieSecure:       true,
		CookieHTTPOnly:     true,
		CookieSameSite:     "",
		SessionStoreType:   "cookie",
		ReverseProxy:       false,
		RealClientIPHeader: "X-Real-IP",
		StandardLogging:    true,
		RequestLogging:     true,
		AuthLogging:        true,
	}
}

// Validate checks that required options are set and valid
func (o *Options) Validate() error {
	if o.ClientID == "" {
		return fmt.Errorf("missing setting: client-id")
	}
	if o.ClientSecret == "" && o.ClientSecretFile == "" {
		return fmt.Errorf("missing setting: client-secret or client-secret-file")
	}
	if o.CookieSecret == "" {
		return fmt.Errorf("missing setting: cookie-secret")
	}
	if len(o.UpstreamURLs) == 0 {
		return fmt.Errorf("missing setting: upstream")
	}
	for _, upstream := range o.UpstreamURLs {
		if _, err := url.Parse(upstream); err != nil {
			return fmt.Errorf("invalid upstream URL %q: %w", upstream, err)
		}
	}
	return nil
}
