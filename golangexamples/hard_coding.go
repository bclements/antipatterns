package main

/*
ANTI-PATTERN: Hard Coding

Embedding configuration values, credentials, file paths, and other values
directly in the source code instead of using configuration files or environment variables.
Makes the code inflexible and difficult to deploy to different environments.
*/

import (
	"fmt"
)

// DatabaseConnection - HARD CODING: Database credentials embedded in code
type DatabaseConnection struct {
	// HARD CODING: Credentials should be in environment variables or config files
	Host     string
	Port     int
	Username string
	Password string // SECURITY ISSUE!
	Database string
}

func NewDatabaseConnection() *DatabaseConnection {
	return &DatabaseConnection{
		Host:     "prod-db-server-01.company.com",
		Port:     5432,
		Username: "admin",
		Password: "SuperSecret123!", // SECURITY ISSUE!
		Database: "production_db",
	}
}

func (dbc *DatabaseConnection) Connect() string {
	// Now changing environments requires modifying source code
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		dbc.Username, dbc.Password, dbc.Host, dbc.Port, dbc.Database)
	return connectionString
}

// EmailService - HARD CODING: Email configuration embedded
type EmailService struct{}

func (es *EmailService) SendEmail(to, subject, body string) {
	// HARD CODING: SMTP settings should be configurable
	smtpServer := "smtp.gmail.com"
	smtpPort := 587
	smtpUsername := "noreply@company.com"
	smtpPassword := "EmailPassword123" // SECURITY ISSUE!

	// HARD CODING: Email templates in code
	if subject == "welcome" {
		body = `
		Welcome to our service!
		Thanks for signing up.
		Visit us at https://www.company.com
		`
	}

	// Simulate sending email
	fmt.Printf("Sending email via %s:%d\n", smtpServer, smtpPort)
	_ = smtpUsername
	_ = smtpPassword
	_ = body
}

// APIClient - HARD CODING: API endpoints and keys embedded
type APIClient struct {
	// HARD CODING: API configuration should be externalized
	BaseURL    string
	APIKey     string // SECURITY ISSUE!
	Timeout    int
	MaxRetries int
}

func NewAPIClient() *APIClient {
	return &APIClient{
		BaseURL:    "https://api.production.company.com/v1",
		APIKey:     "sk_live_abc123def456ghi789", // SECURITY ISSUE!
		Timeout:    30,
		MaxRetries: 3,
	}
}

func (ac *APIClient) MakeRequest(endpoint string) (string, map[string]string) {
	// HARD CODING: Full URLs constructed with hardcoded base
	url := fmt.Sprintf("%s/%s", ac.BaseURL, endpoint)
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", ac.APIKey),
		"Content-Type":  "application/json",
	}
	return url, headers
}

// FileManager - HARD CODING: File paths embedded in code
type FileManager struct{}

func (fm *FileManager) SaveFile(filename string, content []byte) {
	// HARD CODING: Absolute paths that won't work on other systems
	basePath := "/home/john/projects/myapp/uploads"
	filePath := fmt.Sprintf("%s/%s", basePath, filename)

	// HARD CODING: Log file path
	logPath := "/var/log/myapp/file_operations.log"

	fmt.Printf("Saving to %s\n", filePath)
	fmt.Printf("Logging to %s\n", logPath)
}

func (fm *FileManager) GetConfig() string {
	// HARD CODING: Configuration file path
	configFile := "/etc/myapp/config.json"
	return configFile
}

// PaymentProcessor - HARD CODING: Payment gateway credentials and settings
type PaymentProcessor struct{}

func (pp *PaymentProcessor) ProcessPayment(amount float64, card string) float64 {
	// HARD CODING: Stripe API keys in code
	stripeSecretKey := "sk_live_51Hxxxxxxxxxxxxx" // SECURITY ISSUE!
	stripePublicKey := "pk_live_51Hxxxxxxxxxxxxx" // SECURITY ISSUE!

	_ = stripeSecretKey
	_ = stripePublicKey
	_ = card

	// HARD CODING: Payment thresholds and rules
	var fee float64
	if amount > 1000 {
		// HARD CODING: Fee structure in code
		fee = amount*0.029 + 0.30
	} else {
		fee = amount*0.035 + 0.30
	}

	// HARD CODING: Currency
	currency := "USD"
	_ = currency

	return amount + fee
}

// ApplicationConfig - HARD CODING: All configuration in code
type ApplicationConfig struct{}

// HARD CODING: Feature flags
const (
	EnableNewUI         = true
	EnableBetaFeatures  = false
	EnableDebugMode     = false
)

// HARD CODING: Business rules
const (
	MaxLoginAttempts      = 3
	SessionTimeoutMinutes = 30
	PasswordMinLength     = 8
)

// HARD CODING: Third-party service keys
const (
	GoogleMapsAPIKey = "AIzaSyXXXXXXXXXXXXXXXXXX"                     // SECURITY ISSUE!
	SendGridAPIKey   = "SG.XXXXXXXXXXXXXXXX"                          // SECURITY ISSUE!
	AWSAccessKey     = "AKIAXXXXXXXXXXXXXXXX"                         // SECURITY ISSUE!
	AWSSecretKey     = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"     // SECURITY ISSUE!
)

// HARD CODING: URLs
const (
	HomepageURL  = "https://www.company.com"
	APIDocsURL   = "https://docs.company.com/api"
	SupportEmail = "support@company.com"
)

// HARD CODING: Database settings
const (
	DBPoolSize     = 10
	DBPoolTimeout  = 30
	DBMaxOverflow  = 5
)

// SendNotification - HARD CODING: Notification service configuration
func SendNotification(userID int, message string) {
	// HARD CODING: Slack webhook
	slackWebhook := "https://hooks.slack.com/services/T00/B00/XXXXXXXXXXXXXXXX"

	// HARD CODING: Notification rules
	if len(message) > 100 {
		message = message[:100] + "..."
	}

	// HARD CODING: Channels
	channel := "#notifications"
	username := "Bot"

	fmt.Printf("Sending to %s\n", slackWebhook)
	_ = channel
	_ = username
	_ = userID
}

// GetUserAvatar - HARD CODING: CDN and storage URLs
func GetUserAvatar(userID int) string {
	// HARD CODING: CDN URL
	cdnBase := "https://cdn.company.com"

	// HARD CODING: Default avatar path
	defaultAvatar := fmt.Sprintf("%s/images/default-avatar.png", cdnBase)

	// HARD CODING: S3 bucket
	s3Bucket := "company-user-avatars"

	_ = s3Bucket
	_ = userID
	return defaultAvatar
}

// RateLimitCheck - HARD CODING: Rate limiting rules in code
func RateLimitCheck(userID int) bool {
	// HARD CODING: Rate limits
	maxRequestsPerMinute := 60
	maxRequestsPerHour := 1000
	maxRequestsPerDay := 10000

	// HARD CODING: IP whitelist
	whitelistedIPs := []string{
		"192.168.1.100",
		"10.0.0.50",
		"172.16.0.25",
	}

	_ = maxRequestsPerMinute
	_ = maxRequestsPerHour
	_ = maxRequestsPerDay
	_ = whitelistedIPs
	_ = userID

	return true
}

// CacheManager - HARD CODING: Cache configuration
type CacheManager struct {
	// HARD CODING: Redis connection
	RedisHost     string
	RedisPort     int
	RedisPassword string // SECURITY ISSUE!
	RedisDB       int

	// HARD CODING: Cache TTL values
	UserCacheTTL    int // 1 hour
	ProductCacheTTL int // 30 minutes
	SessionCacheTTL int // 2 hours
}

func NewCacheManager() *CacheManager {
	return &CacheManager{
		RedisHost:       "redis.production.company.com",
		RedisPort:       6379,
		RedisPassword:   "RedisPassword123", // SECURITY ISSUE!
		RedisDB:         0,
		UserCacheTTL:    3600,  // 1 hour
		ProductCacheTTL: 1800,  // 30 minutes
		SessionCacheTTL: 7200,  // 2 hours
	}
}

// GenerateReport - HARD CODING: Report generation settings
func GenerateReport() string {
	// HARD CODING: Report title and format
	reportTitle := "Monthly Sales Report - Company Inc."
	reportFormat := "PDF"

	// HARD CODING: Colors and styling
	primaryColor := "#0066CC"
	secondaryColor := "#FF9900"
	fontFamily := "Arial"
	fontSize := 12

	// HARD CODING: Company info
	companyName := "Company Inc."
	companyAddress := "123 Business St, City, State 12345"
	companyPhone := "(555) 123-4567"
	companyEmail := "info@company.com"

	_ = reportFormat
	_ = primaryColor
	_ = secondaryColor
	_ = fontFamily
	_ = fontSize
	_ = companyName
	_ = companyAddress
	_ = companyPhone
	_ = companyEmail

	return reportTitle
}

// HARD CODING: Environment-specific URLs
const (
	AuthURL   = "https://auth.company.com"
	APIURL    = "https://api.company.com"
	WebURL    = "https://www.company.com"
	SocketURL = "wss://socket.company.com"
)

// ValidateUsername - HARD CODING: Validation rules should be configurable
func ValidateUsername(username string) (bool, string) {
	// HARD CODING: Character limits and rules
	if len(username) < 3 {
		return false, "Username must be at least 3 characters"
	}
	if len(username) > 20 {
		return false, "Username must be at most 20 characters"
	}

	// HARD CODING: Blocked usernames
	blockedUsernames := []string{"admin", "root", "administrator", "system"}
	for _, blocked := range blockedUsernames {
		if username == blocked {
			return false, "Username is not allowed"
		}
	}

	return true, "Valid"
}

// HARD CODING: Logging configuration
type Logger struct {
	LogLevel  string
	LogFormat string
	LogFile   string
}

func NewLogger() *Logger {
	return &Logger{
		LogLevel:  "INFO",                           // Should be configurable
		LogFormat: "json",                           // Should be configurable
		LogFile:   "/var/log/myapp/application.log", // Should be configurable
	}
}

// HARD CODING: Feature toggles in code
var features = map[string]bool{
	"new_dashboard":    true,
	"beta_ui":          false,
	"experimental_api": false,
	"dark_mode":        true,
}

// HARD CODING: Third-party endpoints
const (
	StripeEndpoint    = "https://api.stripe.com/v1"
	PayPalEndpoint    = "https://api.paypal.com/v1"
	SendGridEndpoint  = "https://api.sendgrid.com/v3"
	TwilioEndpoint    = "https://api.twilio.com/2010-04-01"
)

func main() {
	/*
		All of these hardcoded values should be:
		1. Moved to environment variables (.env file)
		2. Moved to configuration files (config.yaml, config.json)
		3. Stored in a secrets management system (AWS Secrets Manager, HashiCorp Vault)
		4. Made configurable per environment (dev, staging, production)

		This makes the code more flexible, secure, and maintainable.
	*/

	db := NewDatabaseConnection()
	fmt.Println("Database connection:", db.Connect())

	api := NewAPIClient()
	url, _ := api.MakeRequest("users")
	fmt.Println("API URL:", url)

	// All the hardcoded values make this code:
	// - Insecure (credentials in code)
	// - Inflexible (can't change config without code changes)
	// - Hard to deploy to different environments
	// - Difficult to test
}
