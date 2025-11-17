package main

/*
ANTI-PATTERN: Copy and Paste Programming

Duplicating code instead of creating reusable functions or types.
This leads to maintenance nightmares when bugs need to be fixed in multiple places.
*/

import (
	"fmt"
	"time"
)

// ReportGenerator - Code duplication everywhere instead of creating reusable methods
type ReportGenerator struct{}

func (rg *ReportGenerator) GenerateSalesReport(salesData []map[string]interface{}) {
	// COPY-PASTE: Almost identical to expense report!
	fmt.Println("==================================================")
	fmt.Println("SALES REPORT")
	fmt.Println("==================================================")
	fmt.Printf("Generated: %s\n", time.Now().Format("2006-01-02"))
	fmt.Printf("Total Records: %d\n", len(salesData))
	fmt.Println("--------------------------------------------------")

	total := 0.0
	for _, item := range salesData {
		product := item["product"].(string)
		amount := item["amount"].(float64)
		fmt.Printf("%s: $%.2f\n", product, amount)
		total += amount
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("Total: $%.2f\n", total)
	fmt.Println("==================================================")
}

func (rg *ReportGenerator) GenerateExpenseReport(expenseData []map[string]interface{}) {
	// COPY-PASTE: Almost identical to sales report!
	fmt.Println("==================================================")
	fmt.Println("EXPENSE REPORT")
	fmt.Println("==================================================")
	fmt.Printf("Generated: %s\n", time.Now().Format("2006-01-02"))
	fmt.Printf("Total Records: %d\n", len(expenseData))
	fmt.Println("--------------------------------------------------")

	total := 0.0
	for _, item := range expenseData {
		product := item["product"].(string)
		amount := item["amount"].(float64)
		fmt.Printf("%s: $%.2f\n", product, amount)
		total += amount
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("Total: $%.2f\n", total)
	fmt.Println("==================================================")
}

func (rg *ReportGenerator) GenerateInventoryReport(inventoryData []map[string]interface{}) {
	// COPY-PASTE: Again, almost the same code!
	fmt.Println("==================================================")
	fmt.Println("INVENTORY REPORT")
	fmt.Println("==================================================")
	fmt.Printf("Generated: %s\n", time.Now().Format("2006-01-02"))
	fmt.Printf("Total Records: %d\n", len(inventoryData))
	fmt.Println("--------------------------------------------------")

	total := 0.0
	for _, item := range inventoryData {
		product := item["product"].(string)
		amount := item["amount"].(float64)
		fmt.Printf("%s: $%.2f\n", product, amount)
		total += amount
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("Total: $%.2f\n", total)
	fmt.Println("==================================================")
}

// UserValidator - More copy-paste nightmares
type UserValidator struct{}

func (uv *UserValidator) ValidateAdminUser(username, password, email string) (bool, string) {
	// COPY-PASTE: Validation logic duplicated
	if username == "" {
		return false, "Username is required"
	}
	if len(username) < 3 {
		return false, "Username too short"
	}
	if password == "" {
		return false, "Password is required"
	}
	if len(password) < 8 {
		return false, "Password too short"
	}
	if email == "" {
		return false, "Email is required"
	}
	if !containsChar(email, '@') {
		return false, "Invalid email"
	}
	// Admin-specific check
	if !hasPrefix(username, "admin_") {
		return false, "Admin username must start with admin_"
	}
	return true, "Valid"
}

func (uv *UserValidator) ValidateRegularUser(username, password, email string) (bool, string) {
	// COPY-PASTE: Same validation code with tiny difference
	if username == "" {
		return false, "Username is required"
	}
	if len(username) < 3 {
		return false, "Username too short"
	}
	if password == "" {
		return false, "Password is required"
	}
	if len(password) < 8 {
		return false, "Password too short"
	}
	if email == "" {
		return false, "Email is required"
	}
	if !containsChar(email, '@') {
		return false, "Invalid email"
	}
	// Regular user specific check
	if hasPrefix(username, "admin_") {
		return false, "Regular users cannot have admin_ prefix"
	}
	return true, "Valid"
}

func (uv *UserValidator) ValidateGuestUser(username, password, email string) (bool, string) {
	// COPY-PASTE: Yet again the same validation!
	if username == "" {
		return false, "Username is required"
	}
	if len(username) < 3 {
		return false, "Username too short"
	}
	if password == "" {
		return false, "Password is required"
	}
	if len(password) < 8 {
		return false, "Password too short"
	}
	if email == "" {
		return false, "Email is required"
	}
	if !containsChar(email, '@') {
		return false, "Invalid email"
	}
	// Guest specific check
	if !hasPrefix(username, "guest_") {
		return false, "Guest username must start with guest_"
	}
	return true, "Valid"
}

// Helper functions
func containsChar(s string, c rune) bool {
	for _, ch := range s {
		if ch == c {
			return true
		}
	}
	return false
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// COPY-PASTE: Database query duplication
type DB struct{}

func (db *DB) GetUserByID(userID int) (map[string]interface{}, error) {
	// COPY-PASTE: Same pattern repeated
	// connection := createConnection()
	// defer connection.Close()
	// query := "SELECT * FROM users WHERE id = ?"
	// result := executeQuery(connection, query, userID)
	// return result, nil
	return map[string]interface{}{"id": userID, "name": "User"}, nil
}

func (db *DB) GetProductByID(productID int) (map[string]interface{}, error) {
	// COPY-PASTE: Exact same pattern as above
	// connection := createConnection()
	// defer connection.Close()
	// query := "SELECT * FROM products WHERE id = ?"
	// result := executeQuery(connection, query, productID)
	// return result, nil
	return map[string]interface{}{"id": productID, "name": "Product"}, nil
}

func (db *DB) GetOrderByID(orderID int) (map[string]interface{}, error) {
	// COPY-PASTE: Again the same pattern
	// connection := createConnection()
	// defer connection.Close()
	// query := "SELECT * FROM orders WHERE id = ?"
	// result := executeQuery(connection, query, orderID)
	// return result, nil
	return map[string]interface{}{"id": orderID, "status": "pending"}, nil
}

func (db *DB) GetCustomerByID(customerID int) (map[string]interface{}, error) {
	// COPY-PASTE: One more time...
	// connection := createConnection()
	// defer connection.Close()
	// query := "SELECT * FROM customers WHERE id = ?"
	// result := executeQuery(connection, query, customerID)
	// return result, nil
	return map[string]interface{}{"id": customerID, "name": "Customer"}, nil
}

// APIHandler - COPY-PASTE: API endpoint duplication
type APIHandler struct{}

type Request struct {
	Data map[string]interface{}
}

func (r *Request) GetData() map[string]interface{} {
	return r.Data
}

type Response struct {
	Body   map[string]interface{}
	Status int
}

func (ah *APIHandler) HandleGetUser(request *Request) *Response {
	// COPY-PASTE: Error handling duplicated everywhere
	data := request.GetData()
	if data == nil {
		return &Response{
			Body:   map[string]interface{}{"error": "No data provided"},
			Status: 400,
		}
	}
	// Process user...
	return &Response{
		Body:   map[string]interface{}{"success": true},
		Status: 200,
	}
}

func (ah *APIHandler) HandleGetProduct(request *Request) *Response {
	// COPY-PASTE: Same error handling pattern
	data := request.GetData()
	if data == nil {
		return &Response{
			Body:   map[string]interface{}{"error": "No data provided"},
			Status: 400,
		}
	}
	// Process product...
	return &Response{
		Body:   map[string]interface{}{"success": true},
		Status: 200,
	}
}

func (ah *APIHandler) HandleGetOrder(request *Request) *Response {
	// COPY-PASTE: And again...
	data := request.GetData()
	if data == nil {
		return &Response{
			Body:   map[string]interface{}{"error": "No data provided"},
			Status: 400,
		}
	}
	// Process order...
	return &Response{
		Body:   map[string]interface{}{"success": true},
		Status: 200,
	}
}

func (ah *APIHandler) HandleGetCustomer(request *Request) *Response {
	// COPY-PASTE: Yet again the same error handling
	data := request.GetData()
	if data == nil {
		return &Response{
			Body:   map[string]interface{}{"error": "No data provided"},
			Status: 400,
		}
	}
	// Process customer...
	return &Response{
		Body:   map[string]interface{}{"success": true},
		Status: 200,
	}
}

// COPY-PASTE: Duplicate logging functions
func LogInfo(message string) {
	fmt.Printf("[INFO] %s - %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func LogWarning(message string) {
	// COPY-PASTE: Same as LogInfo with different prefix
	fmt.Printf("[WARNING] %s - %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func LogError(message string) {
	// COPY-PASTE: Same as LogInfo with different prefix
	fmt.Printf("[ERROR] %s - %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func LogDebug(message string) {
	// COPY-PASTE: Same as LogInfo with different prefix
	fmt.Printf("[DEBUG] %s - %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}

func main() {
	// All this code duplication means:
	// - Bugs need to be fixed in multiple places
	// - Changes require updating multiple locations
	// - Code is harder to maintain
	// - More opportunities for inconsistencies

	rg := &ReportGenerator{}
	salesData := []map[string]interface{}{
		{"product": "Widget", "amount": 100.0},
	}
	rg.GenerateSalesReport(salesData)
}
