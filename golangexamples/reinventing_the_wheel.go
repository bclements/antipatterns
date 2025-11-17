package main

/*
ANTI-PATTERN: Reinventing the Wheel

Reimplementing functionality that already exists in standard libraries or
well-established third-party packages. This wastes time and often results
in buggier, less efficient code.
*/

import (
	"fmt"
	"math/rand"
	"strings"
)

// REINVENTING THE WHEEL: Custom JSON parser
// Go has encoding/json package!
func parseJSON(jsonString string) map[string]interface{} {
	/*
		Custom JSON parser - DON'T DO THIS!
		Use: import "encoding/json"; json.Unmarshal([]byte(jsonString), &result)
	*/
	result := make(map[string]interface{})
	// 200 lines of buggy JSON parsing logic...
	// Missing edge cases, security issues, poor performance
	return result
}

// REINVENTING THE WHEEL: Custom HTTP client
// Use net/http package!
type HTTPClient struct {
	/*
		Custom HTTP client - DON'T DO THIS!
		Use: import "net/http"; http.Get(url) or http.Post(url, contentType, body)
	*/
}

func (hc *HTTPClient) Get(url string) (string, error) {
	// 100+ lines reimplementing what net/http does better
	return "", nil
}

func (hc *HTTPClient) Post(url string, data []byte) (string, error) {
	// Missing features: timeout, retries, TLS verification, etc.
	return "", nil
}

// REINVENTING THE WHEEL: Custom date/time formatting
// Go has time.Format()!
func formatDate(year, month, day int) string {
	/*
		Custom date formatter - DON'T DO THIS!
		Use: import "time"; time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
	*/
	months := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	return fmt.Sprintf("%s %d, %d", months[month-1], day, year)
}

// REINVENTING THE WHEEL: Custom CSV parser
// Go has encoding/csv package!
func parseCSV(csvString string) [][]string {
	/*
		Custom CSV parser - DON'T DO THIS!
		Use: import "encoding/csv"; csv.NewReader(strings.NewReader(csvString)).ReadAll()
	*/
	lines := strings.Split(csvString, "\n")
	result := make([][]string, 0)
	for _, line := range lines {
		// This breaks on quoted commas, embedded newlines, etc.
		result = append(result, strings.Split(line, ","))
	}
	return result
}

// REINVENTING THE WHEEL: Custom UUID generator
// Use github.com/google/uuid package!
func generateUUID() string {
	/*
		Custom UUID generator - DON'T DO THIS!
		Use: import "github.com/google/uuid"; uuid.New().String()
	*/
	chars := "0123456789abcdef"
	result := make([]byte, 32)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// REINVENTING THE WHEEL: Custom sorting algorithm
// Go has sort package!
func bubbleSort(arr []int) []int {
	/*
		Custom sorting - DON'T DO THIS!
		Use: import "sort"; sort.Ints(arr)
	*/
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// REINVENTING THE WHEEL: Custom email validator
// Use regex or github.com/badoux/checkmail!
func validateEmail(email string) bool {
	/*
		Custom email validator - DON'T DO THIS!
		Use: regexp.MatchString() with proper pattern or a validation library
	*/
	// Overly simplistic validation missing many edge cases
	if strings.Contains(email, "@") {
		parts := strings.Split(email, "@")
		if len(parts) == 2 && strings.Contains(parts[1], ".") {
			return true
		}
	}
	return false
}

// REINVENTING THE WHEEL: Custom logging system
// Go has log package!
type CustomLogger struct {
	/*
		Custom logger - DON'T DO THIS!
		Use: import "log"; log.Println() or logrus/zap for more features
	*/
	filename string
}

func (cl *CustomLogger) Log(message string) {
	// Missing: log levels, formatting, rotation, structured logging
	fmt.Println(message)
}

// REINVENTING THE WHEEL: Custom configuration parser
// Use encoding/json, gopkg.in/yaml.v2, or github.com/spf13/viper!
func parseConfig(configFile string) map[string]string {
	/*
		Custom config parser - DON'T DO THIS!
		Use: encoding/json or yaml parser or viper
	*/
	config := make(map[string]string)
	// Breaks on comments, nested structures, complex values, etc.
	// Just use standard parsers!
	return config
}

// REINVENTING THE WHEEL: Custom template engine
// Go has text/template and html/template!
func renderTemplate(template string, context map[string]string) string {
	/*
		Custom template engine - DON'T DO THIS!
		Use: import "text/template"; template.New("name").Parse(tmpl).Execute(writer, context)
	*/
	result := template
	for key, value := range context {
		placeholder := fmt.Sprintf("{{%s}}", key)
		result = strings.ReplaceAll(result, placeholder, value)
	}
	return result
}

// REINVENTING THE WHEEL: Custom argument parser
// Use flag package or github.com/spf13/cobra!
func parseArguments(args []string) map[string]string {
	/*
		Custom CLI argument parser - DON'T DO THIS!
		Use: import "flag"; flag.String("name", "default", "usage")
		Or use cobra for complex CLI applications
	*/
	result := make(map[string]string)
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			parts := strings.SplitN(arg[2:], "=", 2)
			if len(parts) == 2 {
				result[parts[0]] = parts[1]
			}
		}
	}
	return result
}

// REINVENTING THE WHEEL: Custom hash function
// Go has crypto/* packages!
func hashString(s string) int {
	/*
		Custom hash function - DON'T DO THIS!
		Use: import "crypto/sha256"; sha256.Sum256([]byte(s))
	*/
	// Not cryptographically secure, collision-prone
	hashValue := 0
	for _, char := range s {
		hashValue = (hashValue*31 + int(char)) % (1 << 32)
	}
	return hashValue
}

// REINVENTING THE WHEEL: Custom random string generator
// Use crypto/rand for security!
func generateRandomString(length int) string {
	/*
		Custom random string generator - DON'T DO THIS!
		Use: import "crypto/rand"; use crypto/rand.Read() for secure random
	*/
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// REINVENTING THE WHEEL: Custom URL parser
// Go has net/url package!
func parseURL(urlStr string) map[string]string {
	/*
		Custom URL parser - DON'T DO THIS!
		Use: import "net/url"; url.Parse(urlStr)
	*/
	result := make(map[string]string)

	// Broken implementation missing many edge cases
	parts := strings.Split(urlStr, "://")
	if len(parts) > 1 {
		result["protocol"] = parts[0]
		rest := parts[1]

		if idx := strings.Index(rest, "/"); idx != -1 {
			result["domain"] = rest[:idx]
			result["path"] = rest[idx:]
		} else {
			result["domain"] = rest
			result["path"] = "/"
		}
	}

	return result
}

// REINVENTING THE WHEEL: Custom retry logic
// Use github.com/avast/retry-go or similar!
func retryFunction(fn func() error, maxAttempts int) error {
	/*
		Custom retry logic - DON'T DO THIS!
		Use: github.com/avast/retry-go or implement with exponential backoff
	*/
	// Missing: exponential backoff, specific exception handling
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if err := fn(); err == nil {
			return nil
		}
	}
	return fmt.Errorf("max attempts reached")
}

// REINVENTING THE WHEEL: Custom database ORM
// Use GORM, sqlx, or database/sql!
type CustomORM struct {
	/*
		Custom ORM - DON'T DO THIS!
		Use: import "gorm.io/gorm" or "github.com/jmoiron/sqlx"
	*/
	connection interface{}
}

func (orm *CustomORM) Save(obj interface{}) error {
	// Missing: relationships, migrations, query optimization, etc.
	return nil
}

func (orm *CustomORM) Find(id int) (interface{}, error) {
	// Incomplete implementation
	return nil, nil
}

// REINVENTING THE WHEEL: Custom serialization
// Go has encoding/gob, encoding/json, or use msgpack!
func serializeObject(obj interface{}) string {
	/*
		Custom serialization - DON'T DO THIS!
		Use: import "encoding/gob" or "encoding/json"
	*/
	// Broken for complex objects, missing types, etc.
	return fmt.Sprintf("%+v", obj)
}

// REINVENTING THE WHEEL: Custom markdown parser
// Use github.com/russross/blackfriday or github.com/gomarkdown/markdown!
func parseMarkdown(text string) string {
	/*
		Custom markdown parser - DON'T DO THIS!
		Use: github.com/russross/blackfriday/v2 or similar
	*/
	// Handles only basic cases, missing most markdown features
	text = strings.Replace(text, "**", "<strong>", 1)
	text = strings.Replace(text, "**", "</strong>", 1)
	text = strings.Replace(text, "*", "<em>", 1)
	text = strings.Replace(text, "*", "</em>", 1)
	return text
}

// REINVENTING THE WHEEL: Custom min/max functions
// Use standard comparison or math package helpers!
func min(a, b int) int {
	/*
		Custom min - DON'T DO THIS!
		Go 1.21+ has: import "cmp"; min := cmp.Or(a, b)
		Or just use: if a < b { return a }; return b
	*/
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	// Same issue as min
	if a > b {
		return a
	}
	return b
}

// REINVENTING THE WHEEL: Custom string contains
// Use strings.Contains()!
func stringContains(s, substr string) bool {
	/*
		Custom string contains - DON'T DO THIS!
		Use: import "strings"; strings.Contains(s, substr)
	*/
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// REINVENTING THE WHEEL: Custom HTTP router
// Use net/http.ServeMux, gorilla/mux, or gin!
type CustomRouter struct {
	/*
		Custom HTTP router - DON'T DO THIS!
		Use: net/http.ServeMux or github.com/gorilla/mux or github.com/gin-gonic/gin
	*/
	routes map[string]func(string) string
}

func (r *CustomRouter) AddRoute(path string, handler func(string) string) {
	// Missing: HTTP methods, path parameters, middleware, etc.
	if r.routes == nil {
		r.routes = make(map[string]func(string) string)
	}
	r.routes[path] = handler
}

// REINVENTING THE WHEEL: Custom error wrapping
// Go 1.13+ has built-in error wrapping with fmt.Errorf("%w", err)!
type CustomError struct {
	/*
		Custom error wrapping - DON'T DO THIS!
		Use: fmt.Errorf("context: %w", originalError)
	*/
	message       string
	originalError error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %v", e.message, e.originalError)
}

// REINVENTING THE WHEEL: Custom base64 encoding
// Go has encoding/base64 package!
func encodeBase64(data []byte) string {
	/*
		Custom base64 encoding - DON'T DO THIS!
		Use: import "encoding/base64"; base64.StdEncoding.EncodeToString(data)
	*/
	// Custom implementation would be buggy and slow
	const base64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	// ... complex implementation that's already done correctly in stdlib
	_ = base64Chars
	return ""
}

func main() {
	/*
		Why reinventing the wheel is bad:
		1. Wastes development time
		2. Results in buggy, incomplete implementations
		3. Missing edge cases and security considerations
		4. No community support or updates
		5. Harder to maintain
		6. Other developers have to learn your custom implementation
		7. Often worse performance than battle-tested libraries

		Always check if functionality exists in:
		- Go standard library (https://pkg.go.dev/std)
		- Well-established third-party packages
		- Framework-specific utilities

		Only create custom implementations when:
		- You have very specific requirements not met by existing solutions
		- Performance is critical and profiling shows existing solutions are bottlenecks
		- You need to avoid dependencies for valid reasons (deployment size, security, etc.)

		Common Go packages to know:
		- encoding/json, encoding/csv, encoding/xml - data formats
		- net/http - HTTP client and server
		- database/sql - database access
		- text/template, html/template - templating
		- crypto/* - cryptographic functions
		- time - date and time operations
		- flag - command-line parsing
		- log - logging
		- strings, bytes - string manipulation
		- sort - sorting
		- regexp - regular expressions
	*/

	fmt.Println("Don't reinvent the wheel - use existing libraries!")
}
