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
	// Note: increased default CookieExpire to 12h (from upstream default of 168h) to
	// reduce session lifetime for better security in my personal deployment.
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
