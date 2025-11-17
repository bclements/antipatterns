package main

/*
ANTI-PATTERN: Dead Code

Code that is never executed or called. It clutters the codebase and
makes it harder to maintain. Unlike Lava Flow (old code nobody dares remove),
this is code that provably never runs.
*/

import (
	"fmt"
)

// DEAD CODE: Function that's never called anywhere
func calculateLegacyTax(amount float64) float64 {
	// This function is never called in the entire codebase
	return amount * 0.15
}

// DEAD CODE: Struct that's never instantiated
type UnusedHelper struct {
	// Nobody creates instances of this struct
	value int
}

func (uh *UnusedHelper) Process() int {
	return uh.value * 2
}

// Order represents an order
type Order struct {
	ID     int
	Status string
	Items  []string
}

// ProcessOrder processes an order with lots of dead code inside
func ProcessOrder(order *Order) *Order {
	// DEAD CODE: Variable assigned but never used
	totalWeight := 0.0
	_ = totalWeight // Added to avoid compiler error in this example

	// DEAD CODE: Unreachable code after return
	if order.Status == "cancelled" {
		return nil
		fmt.Println("This will never print")   // DEAD CODE
		order.Status = "processed"             // DEAD CODE
	}

	// DEAD CODE: Condition that's never true
	if order.Status == "quantum" { // 'quantum' status doesn't exist in the system
		quantumProcess(order) // DEAD CODE
	}

	// DEAD CODE: Code after unconditional return
	if true {
		return order
		// Everything below is DEAD CODE
		order.Status = "completed"
		validateOrder(order)
	}

	// More DEAD CODE (unreachable)
	return processBackup(order)
}

// DEAD CODE: Helper function never called
func quantumProcess(order *Order) *Order {
	// This is called from dead code, so it's also dead code
	return order
}

// DEAD CODE: Helper function never called
func validateOrder(order *Order) bool {
	// Never called anywhere
	return true
}

// DEAD CODE: Helper function never called
func processBackup(order *Order) *Order {
	// Never called anywhere
	return order
}

// DataProcessor processes data with dead code inside
type DataProcessor struct {
	// DEAD CODE: Field that's never read
	unusedCounter int
	legacyFlag    bool
}

func (dp *DataProcessor) Process(data []int) []int {
	// DEAD CODE: Assignment that's overwritten before use
	result := make([]int, 0)
	result = make([]int, len(data))
	for i, v := range data {
		result[i] = v * 2
	}

	// DEAD CODE: Condition that's always false
	if len(data) < 0 { // Length can never be negative
		return make([]int, 0) // DEAD CODE
	}

	// DEAD CODE: Variable assigned but never used
	tempValue := result[0] * 3
	_ = tempValue // Added to avoid compiler error

	// DEAD CODE: Unreachable due to logic
	if len(result) > 0 {
		return result
	} else {
		// DEAD CODE: This else block is unreachable when result is created from data
		return dp.fallbackProcess(data)
	}
}

// DEAD CODE: Method that's never called
func (dp *DataProcessor) fallbackProcess(data []int) []int {
	// Never called anywhere
	return data
}

// DEAD CODE: Method that's never called
func (dp *DataProcessor) Reset() {
	// Never called anywhere
	dp.unusedCounter = 0
}

// CalculateDiscount calculates discount with dead branches
func CalculateDiscount(amount float64, userType string) float64 {
	var discount float64

	// DEAD CODE: These conditions are mutually exclusive
	if userType == "premium" {
		discount = 0.2
	} else if userType == "regular" {
		discount = 0.1
	} else if userType == "premium" { // DEAD CODE: Already handled above
		discount = 0.25 // DEAD CODE
	}

	// DEAD CODE: Unreachable return
	if amount > 0 {
		return amount * (1 - discount)
	} else {
		return 0
	}

	// DEAD CODE: After all paths return
	fmt.Println("Processing complete")
	return 0
}

// DEAD CODE: Commented out code that should be deleted
// func oldImplementation(data []int) []int {
//     // This is the old way we did things
//     result := make([]int, 0)
//     for _, item := range data {
//         result = append(result, item * 2)
//     }
//     return result
// }

// DEAD CODE: Import that's never used (would need to be in real file)
// import "crypto/md5" // md5 is never used

// DEAD CODE: Constants that are never referenced
const (
	UnusedConstant  = 42
	MaxRetries      = 3  // Never used
	DefaultTimeout  = 30 // Never used
)

// UserManager manages users with dead code
type UserManager struct {
	users []string
	// DEAD CODE: Field never accessed
	adminCount int
}

func (um *UserManager) AddUser(user string) string {
	um.users = append(um.users, user)

	// DEAD CODE: Variable assigned but never used
	timestamp := "2024-01-01"
	_ = timestamp

	// DEAD CODE: Computation that's never used
	userCount := len(um.users)
	_ = userCount

	return user
}

// DEAD CODE: Method never called
func (um *UserManager) GetAdminCount() int {
	// This method is never called
	return um.adminCount
}

// DEAD CODE: Method never called
func (um *UserManager) ClearUsers() {
	// This method is never called
	um.users = make([]string, 0)
}

// ComplexFunction has multiple dead code issues
func ComplexFunction(x, y int) int {
	// DEAD CODE: Early return makes everything below unreachable
	return x + y

	// All of this is DEAD CODE:
	var result int
	if x > 100 {
		result = x * 2
	} else {
		result = y * 2
	}

	temp := result / 2
	return temp
}

// SafeDivide divides numbers with dead exception handler
func SafeDivide(a, b float64) float64 {
	// DEAD CODE: We handle zero before it can panic
	if b == 0 {
		return 0
	}
	return a / b

	// DEAD CODE: Unreachable due to return above
	fmt.Println("Division complete")
	return 0
}

// ProcessItems processes items with dead loop
func ProcessItems(items []int) []int {
	if len(items) == 0 {
		return make([]int, 0)
	}

	result := make([]int, 0)

	// DEAD CODE: Loop that never executes (empty slice)
	for _, item := range []int{} {
		result = append(result, item) // DEAD CODE
	}

	// DEAD CODE: Condition that's always false after above check
	if len(items) == 0 {
		return nil // DEAD CODE
	}

	for _, item := range items {
		result = append(result, item*2)
	}

	return result
}

// DEAD CODE: Decorator that's never used
func unusedDecorator(fn func(int) int) func(int) int {
	// This decorator is never applied to any function
	return func(x int) int {
		fmt.Println("Before")
		result := fn(x)
		fmt.Println("After")
		return result
	}
}

// DEAD CODE: Global variable that's never read
var globalConfig = map[string]string{
	"setting1": "value1",
	"setting2": "value2",
} // Never accessed anywhere

// DEAD CODE: Interface that nothing implements
type UnusedInterface interface {
	DoSomething() error
	DoSomethingElse(int) string
}

// DEAD CODE: Struct that implements unused interface
type UnusedImplementation struct{}

func (ui *UnusedImplementation) DoSomething() error {
	return nil
}

func (ui *UnusedImplementation) DoSomethingElse(x int) string {
	return fmt.Sprintf("%d", x)
}

// DEAD CODE: Type alias that's never used
type UnusedAlias = map[string]interface{}

// DEAD CODE: Function with all dead branches
func DeadBranches(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}

	// DEAD CODE: All paths return above
	fmt.Println("This is unreachable")
	return 999
}

// DEAD CODE: Init function that does nothing useful
func init() {
	// This init function doesn't do anything that's actually used
	_ = globalConfig
	_ = UnusedConstant
}

// DEAD CODE: Method on pointer receiver that's never called
func (um *UserManager) unusedMethod() {
	// Never called anywhere
	fmt.Println("This is never called")
}

// DEAD CODE: Variadic function that's never called
func unusedVariadic(args ...interface{}) {
	// Never called anywhere
	for _, arg := range args {
		fmt.Println(arg)
	}
}

// DEAD CODE: Function that returns error but is never called
func unusedErrorReturner() error {
	// Never called anywhere
	return fmt.Errorf("this error is never seen")
}

func main() {
	// Only a tiny portion of the code above is actually used
	order := &Order{ID: 1, Status: "pending"}
	result := ProcessOrder(order)
	fmt.Printf("Order: %+v\n", result)

	// Everything else in this file is dead code that's never executed
}
