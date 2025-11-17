package main

/*
ANTI-PATTERN: Golden Hammer

Using the same solution/tool for every problem regardless of whether it's appropriate.
"If all you have is a hammer, everything looks like a nail."
*/

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// RegexFanatic - Someone who learned regex and now uses it for EVERYTHING
type RegexFanatic struct{}

// ValidateEmail - OK: Regex is appropriate here
func (rf *RegexFanatic) ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}

// IsEven - GOLDEN HAMMER: Using regex to check if a number is even!
// Should just use: number % 2 == 0
func (rf *RegexFanatic) IsEven(number int) bool {
	pattern := `^-?\d*[02468]$`
	matched, _ := regexp.MatchString(pattern, strconv.Itoa(number))
	return matched
}

// AddNumbers - GOLDEN HAMMER: Using regex for basic arithmetic!
// Should just use: a + b
func (rf *RegexFanatic) AddNumbers(a, b int) int {
	// This makes no sense and doesn't even work correctly
	combined := fmt.Sprintf("%d%d", a, b)
	re := regexp.MustCompile(`(\d+)`)
	matches := re.FindAllString(combined, -1)
	if len(matches) > 0 {
		result, _ := strconv.Atoi(matches[0])
		return result
	}
	return 0
}

// ReverseString - GOLDEN HAMMER: Using regex for string reversal
// Should just use a simple loop or strings.Builder
func (rf *RegexFanatic) ReverseString(text string) string {
	// This doesn't even work correctly!
	result := text
	for i := 0; i < len(text); i++ {
		pattern := fmt.Sprintf(`^(.{%d})(.)`, i)
		re := regexp.MustCompile(pattern)
		result = re.ReplaceAllString(result, "$2$1")
	}
	return result
}

// ChannelForEverything - Someone who uses channels for every data structure
type ChannelForEverything struct{}

// GOLDEN HAMMER: Using channel as a queue when a slice would be better
func (cfe *ChannelForEverything) CreateQueue() chan int {
	// A simple slice would be more appropriate for most queue use cases
	return make(chan int, 100)
}

// GOLDEN HAMMER: Using channel as a set (doesn't even work well)
func (cfe *ChannelForEverything) CreateSet() chan string {
	// A map[string]bool would be the right choice
	return make(chan string, 100)
}

// GOLDEN HAMMER: Using channel as a single value holder
func (cfe *ChannelForEverything) CreateVariable(value int) chan int {
	// Just use a variable!
	ch := make(chan int, 1)
	ch <- value
	return ch
}

// InterfaceForEverything - Someone who creates interfaces for everything
// even when concrete types would be simpler
type InterfaceForEverything struct{}

// GOLDEN HAMMER: Creating interfaces for simple operations
type Adder interface {
	Add(a, b int) int
}

type Subtractor interface {
	Subtract(a, b int) int
}

type Multiplier interface {
	Multiply(a, b int) int
}

// These should just be functions, not interfaces!
type Calculator struct{}

func (c Calculator) Add(a, b int) int      { return a + b }
func (c Calculator) Subtract(a, b int) int { return a - b }
func (c Calculator) Multiply(a, b int) int { return a * b }

// StructForEverything - Someone who uses structs even for simple data
type StructForEverything struct{}

// GOLDEN HAMMER: Creating a struct to return two values
type TwoInts struct {
	First  int
	Second int
}

// Should just return (int, int)
func (sfe *StructForEverything) GetTwoNumbers() TwoInts {
	return TwoInts{First: 1, Second: 2}
}

// GOLDEN HAMMER: Struct for a simple string pair
type StringPair struct {
	Key   string
	Value string
}

// A map or tuple would be better
func (sfe *StructForEverything) GetConfig() StringPair {
	return StringPair{Key: "timeout", Value: "30"}
}

// GoRoutineForEverything - Someone who uses goroutines for everything
type GoRoutineForEverything struct{}

// GOLDEN HAMMER: Using goroutine for a simple function call
func (gfe *GoRoutineForEverything) PrintMessage(msg string) {
	// No need for a goroutine here, adds unnecessary complexity
	go func() {
		fmt.Println(msg)
	}()
}

// GOLDEN HAMMER: Using goroutine for sequential operations
func (gfe *GoRoutineForEverything) ProcessData(data []int) []int {
	// This makes it slower and more complex!
	resultChan := make(chan []int)
	go func() {
		result := make([]int, len(data))
		for i, v := range data {
			result[i] = v * 2
		}
		resultChan <- result
	}()
	return <-resultChan
}

// ReflectionForEverything - Someone who uses reflection when they shouldn't
type ReflectionForEverything struct{}

// GOLDEN HAMMER: Using reflection for simple type check
func (rfe *ReflectionForEverything) IsString(value interface{}) bool {
	// Should just use type assertion: _, ok := value.(string)
	return fmt.Sprintf("%T", value) == "string"
}

// PointerForEverything - Someone who uses pointers everywhere
type PointerForEverything struct{}

// GOLDEN HAMMER: Using pointers for small primitive types unnecessarily
func (pfe *PointerForEverything) AddNumbers(a *int, b *int) *int {
	// Pointers for ints add complexity without benefit
	result := *a + *b
	return &result
}

func (pfe *PointerForEverything) IsEven(n *int) *bool {
	// Returning pointer to bool is usually unnecessary
	result := *n%2 == 0
	return &result
}

// MapForEverything - Someone who uses maps for every data structure
type MapForEverything struct{}

// GOLDEN HAMMER: Using map as a list
func (mfe *MapForEverything) CreateList() map[int]string {
	// A slice []string would be more appropriate
	return make(map[int]string)
}

// GOLDEN HAMMER: Using map to store two values
func (mfe *MapForEverything) GetUserInfo() map[string]string {
	// A struct would be better
	return map[string]string{
		"name":  "John",
		"email": "john@example.com",
	}
}

// ContextForEverything - Someone who passes context.Context everywhere
// even when it's not needed

// GOLDEN HAMMER: Using context for simple value passing
func ProcessWithUnnecessaryContext(data string) string {
	// Context is overkill when you just need to pass a value
	// Should just pass the value as a parameter
	return strings.ToUpper(data)
}

// FactoryPatternForEverything - Using factories for simple object creation
type FactoryPatternForEverything struct{}

// GOLDEN HAMMER: Factory for a simple struct
type Point struct {
	X, Y int
}

type PointFactory struct{}

func NewPointFactory() *PointFactory {
	return &PointFactory{}
}

func (pf *PointFactory) CreatePoint(x, y int) *Point {
	// Just use: &Point{X: x, Y: y}
	// No factory needed!
	return &Point{X: x, Y: y}
}

// SingletonForEverything - Using singleton pattern unnecessarily
var (
	configInstance     *Config
	loggerInstance     *Logger
	validatorInstance  *Validator
	formatterInstance  *Formatter
	converterInstance  *Converter
	parserInstance     *Parser
	// Everything is a singleton!
)

type Config struct{ value string }
type Logger struct{ name string }
type Validator struct{}
type Formatter struct{}
type Converter struct{}
type Parser struct{}

// GOLDEN HAMMER: Making everything a singleton when it doesn't need to be

// MicroservicesForEverything - Someone who splits everything into microservices
/*
GOLDEN HAMMER: The developer learned about microservices and now wants to split
a simple application into 50 microservices:

- UserEmailValidationService
- UserFirstNameService
- UserLastNameService
- AddTwoNumbersService
- SubtractTwoNumbersService
- StringToUpperCaseService
- StringToLowerCaseService

Each service has its own repository, database, API, and deployment pipeline
for functionality that should just be simple functions!
*/

func main() {
	// Examples of golden hammer in action
	rf := &RegexFanatic{}
	fmt.Println("Is 4 even?", rf.IsEven(4)) // Using regex instead of modulo!

	pfe := &PointerForEverything{}
	a, b := 5, 10
	result := pfe.AddNumbers(&a, &b) // Unnecessary pointer complexity
	fmt.Println("Sum:", *result)

	// The golden hammer principle: when all you have is a hammer,
	// everything looks like a nail. Use the right tool for the job!
}
