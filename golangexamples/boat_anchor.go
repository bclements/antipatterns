package main

/*
ANTI-PATTERN: Boat Anchor

Code that serves no purpose but is kept around "just in case" it might be needed.
This weighs down the codebase like a boat anchor.
*/

import (
	"fmt"
)

// UserManager manages users with lots of boat anchor code
type UserManager struct {
	users []User

	// BOAT ANCHOR: This was for the legacy XML export feature we removed 2 years ago
	// But keeping it just in case we need it again
	xmlExporter *XMLExporter // Never used anymore

	// BOAT ANCHOR: This was for the backup system that was replaced
	// Keeping it just in case
	// backupHandler *BackupHandler
}

type User struct {
	ID    int
	Name  string
	Email string
}

// NewUserManager creates a new user manager
func NewUserManager() *UserManager {
	return &UserManager{
		users:       make([]User, 0),
		xmlExporter: NewXMLExporter(), // Instantiated but never used
	}
}

// AddUser adds a user to the system
func (um *UserManager) AddUser(user User) {
	um.users = append(um.users, user)

	// BOAT ANCHOR: Old method - used to notify via email, now we use webhooks
	// But maybe we'll need email again someday?
	// um.sendEmailNotification(user)

	// NEW METHOD:
	um.sendWebhookNotification(user)
}

func (um *UserManager) sendWebhookNotification(user User) {
	fmt.Printf("Sending webhook for %s\n", user.Name)
}

// BOAT ANCHOR: This hasn't been called in years but "we might need it"
func (um *UserManager) sendEmailNotification(user User) {
	// Legacy email notification - UNUSED
	smtpServer := "old.smtp.server.com"
	// 50 lines of email code that's never executed...
	_ = smtpServer
}

// BOAT ANCHOR: Feature we planned but never implemented
func (um *UserManager) ExportToXML() error {
	// Planned XML export feature - NEVER IMPLEMENTED
	// TODO: Implement this when we have time (been here for 3 years)
	return nil
}

// XMLExporter - BOAT ANCHOR: Legacy struct that's instantiated but never used
type XMLExporter struct {
	config map[string]interface{}
}

func NewXMLExporter() *XMLExporter {
	return &XMLExporter{
		config: loadLegacyConfig(),
	}
}

func loadLegacyConfig() map[string]interface{} {
	// Complex initialization that happens every time but is never used
	return map[string]interface{}{
		"format":  "xml",
		"version": "1.0",
	}
}

func (xe *XMLExporter) Export(data interface{}) error {
	// This method is never called anywhere
	return nil
}

// BOAT ANCHOR: Keeping these around "just in case we need to rollback"
const (
	OldDatabaseConnectionString = "mysql://oldserver:3306/legacy_db"
	OldAPIEndpoint              = "http://deprecated.api.com/v1"
	OldEncryptionKey            = "old_key_that_we_dont_use_anymore"
)

// BOAT ANCHOR: Old processing function that's been replaced but kept around
func legacyDataProcessing(data []byte) ([]byte, error) {
	/*
		This was the old way we processed data before we switched to the new system.
		Keeping it here just in case the new system has issues and we need to rollback.
		Note: The new system has been stable for 2 years now.
	*/
	return data, nil
}

// BOAT ANCHOR: Deprecated interface that nothing implements anymore
type LegacyProcessor interface {
	ProcessLegacy(data []byte) error
	ValidateLegacy(data []byte) bool
	TransformLegacy(data []byte) []byte
}

// BOAT ANCHOR: Old constants that may or may not still be relevant
const (
	MaxRetriesOldSystem      = 5     // For old retry logic we don't use
	TimeoutOldSystem         = 30000 // Milliseconds - old system timeout
	LegacyAPIVersion         = "v1"  // We're on v3 now
	DeprecatedFeatureFlag    = true  // Feature was removed in 2019
	ExperimentalFeatureAlpha = false // Experiment concluded in 2018
)

// BOAT ANCHOR: Helper functions for features that no longer exist
func convertToLegacyFormat(data interface{}) interface{} {
	// Conversion logic for a format we no longer use
	// But someone said "don't delete it, we might need it"
	return data
}

func validateLegacySchema(schema string) bool {
	// Validates against a schema we deprecated 3 years ago
	return true
}

// BOAT ANCHOR: Struct for a third-party integration we no longer use
type OldPaymentProvider struct {
	APIKey   string
	Endpoint string
	Version  int
}

// We switched to a new payment provider in 2020, but this is still here
func NewOldPaymentProvider(apiKey string) *OldPaymentProvider {
	return &OldPaymentProvider{
		APIKey:   apiKey,
		Endpoint: "https://old-payment-provider.com/api",
		Version:  1,
	}
}

func (opp *OldPaymentProvider) ProcessPayment(amount float64) error {
	// Never called anymore, but "might be useful for historical reference"
	return nil
}

// BOAT ANCHOR: Error types we defined but never use
type LegacyError struct {
	Code    int
	Message string
	Details map[string]interface{}
}

func (e *LegacyError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// These error constructors are defined but never used
func NewLegacyAuthError(msg string) *LegacyError {
	return &LegacyError{Code: 1001, Message: msg}
}

func NewLegacyValidationError(msg string) *LegacyError {
	return &LegacyError{Code: 1002, Message: msg}
}

func main() {
	// Only the AddUser functionality is actually used
	um := NewUserManager()
	um.AddUser(User{ID: 1, Name: "John", Email: "john@example.com"})

	// Everything else in this file is boat anchor code that's never called
}
