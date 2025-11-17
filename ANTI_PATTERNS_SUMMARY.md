# Anti-Patterns Summary Guide

A comprehensive reference for common software anti-patterns with examples in Python and Go.

---

## 1. Boat Anchor

**Definition:** Code that serves no purpose but is kept around "just in case" it might be needed. This weighs down the codebase like a boat anchor.

### Python Examples (`pythonexamples/boat_anchor.py`)

**1. Legacy XML Exporter** - Never used but instantiated
- XMLExporter created on every UserManager initialization
- Complex config loading that runs but produces no value
- "Just in case we need XML export again someday"

**2. Deprecated Email Notification** - Replaced but kept
- Old SMTP email notification code commented out
- New webhook system in place for years
- Nobody has needed email since 2019

**3. Old Database Connections** - From previous architecture
- MySQL connection strings from old system
- Backup database handler for removed backup system
- Constants for deprecated API endpoints

**4. Unused Methods** - Never called anywhere
- `export_to_xml()` - planned feature never implemented
- `_send_email_notification()` - hasn't been called in years
- Helper functions for features that were removed

### Go Examples (`golangexamples/boat_anchor.go`)

**1. Deprecated Payment Provider** - Old integration kept
- `OldPaymentProvider` struct for service replaced in 2020
- Complete implementation that's never instantiated
- "Keeping for historical reference"

**2. Legacy Error Types** - Defined but unused
- `LegacyError` with constructors that are never called
- Multiple error types for systems that don't exist
- Dead code that compiles but serves no purpose

**3. Feature Flags for Removed Features** - Constants nobody checks
- `DeprecatedFeatureFlag` for feature removed in 2019
- `ExperimentalFeatureAlpha` from 2018 experiment
- Magic numbers nobody remembers the purpose of

**4. Commented Legacy Code** - "DO NOT DELETE"
- Entire backup handler system commented out
- Multiple generations of implementations in comments
- Fear-based preservation

### Why This Matters

**Impact:**
- Increases cognitive load when reading code
- Makes codebase harder to navigate
- Creates confusion about what's actually used
- Waste build and test time

**Detection:**
- Code coverage tools show 0% coverage
- Grep for usages returns nothing
- Git history shows no recent modifications
- Comments say "TODO: Remove after migration" from years ago

**Solution:**
- Use version control (git) as your safety net
- Delete with confidence - you can always restore
- Document removal decisions in commit messages
- Regular cleanup sprints to remove boat anchors

---

## 2. Spaghetti Code

**Definition:** Code with complex and tangled control flow that's difficult to understand and maintain. Logic jumps around making it hard to follow the execution path.

### Python Examples (`pythonexamples/spaghetti_code.py`)

**1. ProcessOrder Function** - Deeply nested conditionals
- 6+ levels of nested if/else statements
- Complex combinations of user type, payment method, discount codes
- Same calculations duplicated in multiple branches
- Impossible to trace all execution paths

**2. OrderProcessor State Machine** - Tangled state management
- Multiple boolean flags (`flag1`, `flag2`) with unclear meaning
- State transitions depend on complex flag combinations
- Methods call each other in circular patterns
- Counter increments scattered throughout

**3. ValidateUser Function** - Pyramid of doom
- 8 levels of nested conditionals
- Each level checks one more validation rule
- Error messages buried deep in nesting
- Cannot easily add new validation rules

### Go Examples (`golangexamples/spaghetti_code.go`)

**1. ProcessOrder Function** - 150+ lines of nested logic
- User type × payment method × discount code × shipping method combinations
- Total amount calculated in multiple scattered locations
- Shipping cost logic intertwined with discount logic
- No clear separation of concerns

**2. OrderProcessor Struct** - Complex state machine
- Five different states with unclear transitions
- Multiple flags and counters affecting behavior
- Methods calling other methods in unpredictable order
- State transition logic duplicated across methods

**3. ValidateUser Function** - 9 levels of nesting
- Progressive validation with each level checking more
- Age restrictions vary by user type and country
- Early returns deep in nesting make control flow unclear

### Why This Matters

**Impact:**
- Extremely difficult to understand and modify
- High bug risk when making changes
- Nearly impossible to test all code paths
- Team velocity slows dramatically
- Onboarding new developers takes weeks

**Detection:**
- Cyclomatic complexity > 15
- Nesting depth > 5 levels
- Functions longer than 100 lines
- Cannot explain function in one sentence

**Solution:**
- Extract methods/functions for each responsibility
- Use early returns to reduce nesting
- Replace complex conditions with well-named boolean variables
- Apply strategy pattern for branching logic
- Break into smaller, single-purpose functions

---

## 3. Golden Hammer

**Definition:** Using the same solution/tool for every problem regardless of whether it's appropriate. "If all you have is a hammer, everything looks like a nail."

### Python Examples (`pythonexamples/golden_hammer.py`)

**1. Regex for Everything** - When you love regular expressions too much
- Using regex to check if number is even: `r'^-?\d*[02468]$'`
- Regex for arithmetic operations
- Regex to reverse strings
- All problems become regex problems

**2. Classes for Everything** - OOP zealot
- `MathOperations` class with only static methods
- Creating classes when simple functions would suffice
- Inheritance used just to reuse one method (Car inherits from Animal!)
- Every concept must be a class

**3. Design Patterns Everywhere** - Pattern addiction
- Singleton + Factory for creating strings
- Observer pattern for simple callbacks
- Strategy pattern for two-line functions
- Complexity without benefit

**4. Database for Everything** - DB as only data structure
- Storing temporary variables in database
- Calculation results persisted unnecessarily
- Map operations become SQL queries

### Go Examples (`golangexamples/golden_hammer.go`)

**1. Regex Fanatic** - Same as Python
- Regex for even number checking
- Regex for simple arithmetic
- When you have `regexp` package, everything is a pattern

**2. Channels for Everything** - Channel abuse
- Channel as a single value holder (`chan int` with buffer 1)
- Channel as a set (doesn't work well)
- Channel as a queue when slice would be better
- Goroutine + channel for synchronous operations

**3. Interfaces for Everything** - Interface overload
- Interface for simple addition function
- Interface with one implementation
- Interfaces that add complexity without flexibility

**4. Pointers for Everything** - Unnecessary pointer complexity
- Pointers to ints, bools, small structs
- Returning pointers to primitives
- Memory allocation overhead with no benefit

**5. Microservices for Everything** - Service per function
- `AddTwoNumbersService`
- `StringToUpperCaseService`
- Network overhead for trivial operations

### Why This Matters

**Impact:**
- Overcomplicated solutions to simple problems
- Performance overhead from inappropriate tools
- Maintenance burden from unnecessary abstraction
- Team must learn custom patterns instead of standard approaches

**Detection:**
- Same pattern/tool appears everywhere
- Simple operations have complex implementations
- New team members question design choices
- Performance problems from tool mismatch

**Solution:**
- Evaluate tools based on problem requirements
- Learn multiple approaches and when to use each
- Use simplest solution that works
- "Make it work, make it right, make it fast" - in that order

---

## 4. Copy and Paste Programming

**Definition:** Duplicating code instead of creating reusable functions or classes. This leads to maintenance nightmares when bugs need to be fixed in multiple places.

### Python Examples (`pythonexamples/copy_paste_programming.py`)

**1. Report Generator** - Three nearly identical report methods
- `generate_sales_report()`, `generate_expense_report()`, `generate_inventory_report()`
- Only difference: title and data source
- Same formatting logic duplicated 3 times
- Bug fixes must be applied to all three

**2. User Validator** - Validation logic duplicated
- `validate_admin_user()`, `validate_regular_user()`, `validate_guest_user()`
- Same username/password/email validation in each
- Only user-type-specific check differs
- Each has 15+ lines of identical code

**3. Database Queries** - Query pattern copied
- `get_user_by_id()`, `get_product_by_id()`, `get_order_by_id()`
- Same connection handling, query execution, cleanup
- Only table name differs
- Repeated connection management bugs

**4. API Handlers** - Error handling duplicated
- Same try/catch structure in every handler
- Identical error response formatting
- Same validation checks repeated
- Any error handling improvement requires 10+ file changes

### Go Examples (`golangexamples/copy_paste_programming.go`)

**1. Report Generator** - Identical report methods
- Three functions with 20+ lines of duplicate code each
- Same header, footer, totaling logic
- Only report title varies
- Cannot change format without touching all three

**2. User Validator** - Validation duplication
- Three validation functions sharing 90% of code
- Username, password, email checks repeated verbatim
- Prefix checking is the only difference
- Maintenance burden of keeping in sync

**3. Database Methods** - Generic query pattern duplicated
- `GetUserByID()`, `GetProductByID()`, `GetOrderByID()`, `GetCustomerByID()`
- Connection creation/cleanup duplicated
- Should be generic `GetByID[T](table string, id int) T`
- No use of Go generics or reflection

**4. Logging Functions** - Log level duplication
- `LogInfo()`, `LogWarning()`, `LogError()`, `LogDebug()`
- Same timestamp formatting in each
- Only log level prefix differs
- Should be `Log(level, message)`

### Why This Matters

**Impact:**
- Bug fixes require multiple changes (easy to miss one)
- Inconsistencies emerge over time
- Testing burden multiplied
- Code bloat and harder navigation
- Changes are risky (did we update all copies?)

**Detection:**
- Code similarity analysis tools
- Manual code review finds patterns
- Bugs appear in multiple places
- Similar function/method names

**Solution:**
- Extract common logic into reusable functions
- Use templates/generics for similar structures
- DRY principle: Don't Repeat Yourself
- Create utility functions and base classes
- Refactor before adding third copy

---

## 5. God Object

**Definition:** A class/struct that knows too much or does too much. It has too many responsibilities and becomes a central point that controls or knows about everything in the system. Violates the Single Responsibility Principle.

### Python Examples (`pythonexamples/god_object.py`)

**1. ApplicationManager** - Manages EVERYTHING
- User management (create, delete, update, authenticate)
- Database operations (connect, query, backup, migrate)
- Product management (add, remove, search, recommend)
- Order processing (create, cancel, track, ship)
- Payment processing (charge, refund, validate)
- Email service (send, queue, templates)
- Logging and audit trail
- Caching (set, get, invalidate)
- File storage (upload, delete, URL generation)
- Notifications (send, mark read, preferences)
- Analytics (track events, generate reports)
- Configuration management
- Feature flags
- Third-party integrations (Stripe, AWS, SendGrid)

**2. 50+ Methods** - All in one class
- Each method should be in a separate service
- Thousands of lines of code in one file
- Impossible to understand the full scope
- Every feature change touches this class

**3. Dozens of Fields** - State explosion
- 20+ instance variables
- Multiple data structures (lists, dicts, sets)
- Connection pools, caches, queues
- No clear organization

### Go Examples (`golangexamples/god_object.go`)

**1. ApplicationManager Struct** - 15+ responsibilities
- Same problems as Python version
- User authentication and authorization
- Database connection management
- Product catalog and inventory
- Order and shipping
- Payment processing
- Email and notifications
- Logging and analytics
- Caching and configuration
- File storage
- All third-party integrations

**2. 60+ Methods** - Massive interface
- CreateUser, DeleteUser, UpdateUserProfile
- ConnectToDatabase, ExecuteQuery, BackupDatabase
- ProcessPayment, RefundPayment, ValidateCreditCard
- SendEmail, SendWelcomeEmail, SendOrderConfirmation
- All business logic in one place

**3. Violates SOLID Principles**
- Single Responsibility: Has 15 responsibilities
- Open/Closed: Every feature requires modification
- Interface Segregation: Clients depend on unused methods
- Dependency Inversion: Tightly coupled to implementations

### Why This Matters

**Impact:**
- Impossible to understand completely
- Changes are high risk (ripple effects)
- Merge conflicts constantly
- Cannot test in isolation
- Team paralysis (everyone edits same file)
- Violates every SOLID principle
- Becomes the bottleneck for all work

**Detection:**
- Class > 500 lines
- More than 10 public methods
- More than 10 instance variables
- Multiple unrelated responsibilities
- Name ends in "Manager", "Util", "Helper", "Service"
- Every feature touches this class

**Solution:**
- Split into separate classes by responsibility
- Use dependency injection
- Create focused services (UserService, PaymentService, etc.)
- Apply Single Responsibility Principle
- Facade pattern to maintain compatibility during migration

---

## 6. Lava Flow

**Definition:** Dead code and forgotten functionality that remains in the codebase because nobody knows if it's safe to remove it. Like lava that has cooled and hardened, this code is now part of the landscape but serves no purpose.

### Python Examples (`pythonexamples/lava_flow.py`)

**1. Mysterious Legacy Function** - Unknown purpose
- `mysterious_legacy_function()` - "Don't touch! Bob said it's important"
- Bob left in 2017
- Nobody understands the logic
- Magic numbers (42, 1.337) with no explanation

**2. DataProcessor with Ancient History** - Layers of old code
- `legacy_mode` flag that's always False
- `use_old_algorithm` flag from 2018 migration
- `enable_experimental_feature` from 2016 experiment
- Methods for never-released features

**3. Legacy Configuration** - Unclear if used
- Dictionary of settings from various eras
- Some still used, some not - nobody knows which
- Dead API endpoints and feature flags
- "Safe to remove? Nobody knows"

**4. Commented Code Archaeology** - Multiple generations
- Four different implementations all commented out
- Dates from 2016, 2018, 2019, 2022
- "Old implementation - DO NOT DELETE"
- Should be in version control, not comments

**5. Emergency Fix from 2016** - Temporary became permanent
- `emergency_fix_for_bug_123()` with dead bug tracker link
- "TODO: Remove after proper fix" - still there in 2024
- "DO NOT REMOVE: This is critical! (Why? Unknown)"

### Go Examples (`golangexamples/lava_flow.go`)

**1. Legacy Processing Logic** - Ancient alternatives
- `legacyProcess()` method never executed
- System X decommissioned in 2017
- "Critical for backwards compatibility" (with what?)
- Complex logic nobody dares touch

**2. OldUserManager** - Deprecated but kept
- "DEPRECATED: Use NewUserManager instead"
- Last used: Unknown
- Safe to delete: Nobody knows
- Complete implementation still maintained

**3. Feature Flags from the Past** - Dead toggles
- `EnableOldAPI` - API removed in 2018
- `UseDeprecatedFormatter` - formatter replaced in 2019
- `LegacyModeEnabled` - never actually implemented
- All set to false, none checked

**4. Commented Imports and Globals** - Fossil code
- Import comments for removed dependencies
- Global variables from 2014 singleton pattern
- `initialized` flag that may not be checked
- Fear prevents cleanup

### Why This Matters

**Impact:**
- Increases codebase complexity
- Makes it harder to understand current behavior
- Creates fear of making changes
- New developers waste time understanding dead code
- Merge conflicts in code that doesn't matter
- Slows down searches and navigation

**Detection:**
- Code with comments like "DO NOT DELETE" or "KEEP FOR NOW"
- Flags/features always set to same value
- Dead bug tracker links
- Author left company years ago
- Multiple commented-out implementations

**Solution:**
- Use version control as safety net (git history)
- Remove commented code - it's in version control
- Feature flags should have sunset dates
- Documentation in commit messages, not code
- Regular archaeology sessions to remove lava flow
- Test coverage to gain confidence

---

## 7. Dead Code

**Definition:** Code that is never executed or called. It clutters the codebase and makes it harder to maintain. Unlike Lava Flow (old code nobody dares remove), this is code that provably never runs.

### Python Examples (`pythonexamples/dead_code.py`)

**1. Unreachable Code After Return**
- Code after `return` statement
- Code after unconditional `break` or `continue`
- Code in `else` branch of condition that's always true

**2. Unused Functions and Classes**
- `calculate_legacy_tax()` - never called anywhere
- `UnusedHelper` class - never instantiated
- `quantum_process()` - called from other dead code

**3. Impossible Conditions**
- `if len(data) < 0:` - length can never be negative
- `if True: return x; else: return y` - else never executes
- Duplicate conditions in if/elif chain

**4. Variables Assigned but Never Read**
- `total_weight = 0` then never used
- Result calculated but overwritten before use
- Parameters passed but never accessed

**5. Imports Never Used**
- `from typing import Dict` - Dict never referenced
- Entire modules imported but unused

### Go Examples (`golangexamples/dead_code.go`)

**1. Unreachable Code Patterns**
- Code after `return` in all branches
- Code after unconditional `panic()`
- Else blocks that can never execute

**2. Unused Functions and Methods**
- Functions defined but never called
- Methods on types that are never invoked
- Entire interfaces with no implementations

**3. Dead Constants and Globals**
- `UnusedConstant = 42` - never referenced
- Global maps initialized but never accessed
- Type aliases never used

**4. Impossible Branches**
- `if x > 0 { return 1 } else if x < 0 { return -1 } else { return 0 }; print("here")` - print never runs
- Duplicate conditions in switch cases

**5. Unused Struct Fields**
- Fields assigned in constructor but never read
- Parameters that are silently ignored
- Return values from functions that are never used

### Why This Matters

**Impact:**
- Clutters codebase and reduces readability
- Wastes CPU cycles (in some cases)
- Creates confusion about what actually runs
- False sense of functionality
- Increases maintenance burden
- Linters and static analysis catch most of these

**Detection:**
- Code coverage tools show 0% coverage
- Static analysis (unused variables, unreachable code)
- Linters flag dead code
- Go compiler errors on unused variables/imports
- Manual code inspection

**Solution:**
- Run static analysis tools regularly
- Enable linter rules for dead code
- Maintain code coverage reports
- Delete unused code aggressively
- Use "unused" directives only when necessary

---

## 8. Hard Coding

**Definition:** Embedding configuration values, credentials, file paths, and other values directly in the source code instead of using configuration files or environment variables. Makes the code inflexible and difficult to deploy to different environments.

### Python Examples (`pythonexamples/hard_coding.py`)

**1. Database Credentials** - Embedded secrets
```python
host = "prod-db-server-01.company.com"
password = "SuperSecret123!"  # SECURITY ISSUE!
```
- Credentials in source code
- Environment-specific values
- Security risk if code is shared

**2. API Keys and Secrets** - Hardcoded everywhere
```python
STRIPE_SECRET_KEY = "sk_live_abc123..."
AWS_ACCESS_KEY = "AKIAXXXXXXXX"
SENDGRID_API_KEY = "SG.XXXXXXXXX"
```
- Keys visible in version control
- Cannot rotate without code deploy
- Accidental exposure in logs/errors

**3. File Paths** - System-specific locations
```python
base_path = "/home/john/projects/myapp/uploads"
config_file = "/etc/myapp/config.json"
```
- Won't work on other systems
- Developer username in path
- Absolute paths instead of relative

**4. Business Rules** - Logic in constants
```python
MAX_LOGIN_ATTEMPTS = 3
SESSION_TIMEOUT_MINUTES = 30
PASSWORD_MIN_LENGTH = 8
```
- Cannot change without deployment
- Different per environment
- Should be configurable

**5. URLs and Endpoints** - Environment coupling
```python
API_URL = "https://api.production.company.com"
AUTH_URL = "https://auth.company.com"
```
- Hardcoded to production
- Cannot test against staging
- No environment flexibility

### Go Examples (`golangexamples/hard_coding.go`)

**1. Database Connection Hardcoded**
```go
Host:     "prod-db-server-01.company.com",
Password: "SuperSecret123!",
```
- Cannot switch environments
- Credentials in compiled binary
- Security nightmare

**2. Payment Provider Keys**
```go
stripeSecretKey := "sk_live_51Hxxxxxxxxxxxxx"
awsAccessKey := "AKIAXXXXXXXXXXXXXXXX"
```
- API keys in source
- Version control exposure
- Cannot rotate easily

**3. Feature Flags as Constants**
```go
const (
    EnableNewUI        = true
    EnableBetaFeatures = false
)
```
- Requires recompilation to change
- Cannot toggle per user/tenant
- No runtime control

**4. Rate Limits and Thresholds**
```go
maxRequestsPerMinute := 60
whitelistedIPs := []string{"192.168.1.100", "10.0.0.50"}
```
- Cannot adjust based on load
- IP whitelist in code
- No flexibility

### Why This Matters

**Impact:**
- **Security**: Credentials exposed in version control
- **Flexibility**: Cannot deploy to different environments
- **Testing**: Cannot easily test with test data
- **Rotation**: Cannot rotate secrets without redeployment
- **Scalability**: Cannot adjust limits dynamically
- **Compliance**: May violate security policies

**Detection:**
- Grep for passwords, API keys, connection strings
- Look for environment-specific values (prod URLs)
- Absolute file paths
- IP addresses and ports in code
- String literals that should be configurable

**Solution:**
- **Environment Variables**: Use `os.getenv()` / `os.Getenv()`
- **Config Files**: JSON, YAML, TOML configuration
- **Secrets Management**: AWS Secrets Manager, HashiCorp Vault
- **Feature Flags**: LaunchDarkly, Split.io, or custom service
- **Constants**: Only for truly constant values (π, regex patterns)

---

## 9. Reinventing the Wheel

**Definition:** Reimplementing functionality that already exists in standard libraries or well-established third-party packages. This wastes time and often results in buggier, less efficient code.

### Python Examples (`pythonexamples/reinventing_the_wheel.py`)

**1. Custom JSON Parser** - Don't parse JSON manually!
- Python has `import json; json.loads()`
- Custom parser will have bugs and edge cases
- Missing: unicode handling, escape sequences, nested objects

**2. Custom HTTP Client** - Use requests or httpx!
- `import requests; requests.get(url)` works perfectly
- Custom version missing: timeouts, retries, SSL, redirects
- Reinventing 10,000+ lines of battle-tested code

**3. Custom Sorting** - Python has `sorted()`!
- Implementing bubble sort (O(n²)) instead of using built-in Timsort (O(n log n))
- Built-in is faster and tested
- No reason to write custom sort

**4. Custom Date Formatting** - Use datetime.strftime()!
- `datetime.strftime('%Y-%m-%d')` handles all cases
- Custom version won't handle leap years, timezones
- Locale-aware formatting for free

**5. Custom Email Validation** - Use email-validator library!
- Email RFC is complex (3000+ pages)
- Simple `'@' in email` misses 99% of cases
- Well-tested libraries exist

**6. Custom Logging System** - Python has `logging` module!
- Missing: log levels, rotation, handlers, filters
- Built-in logging is flexible and powerful
- Integrates with all frameworks

### Go Examples (`golangexamples/reinventing_the_wheel.go`)

**1. Custom JSON Parser** - Use `encoding/json`!
```go
// Don't do this:
func parseJSON(s string) map[string]interface{}

// Use this:
json.Unmarshal([]byte(s), &result)
```

**2. Custom HTTP Client** - Use `net/http`!
```go
// Don't do this:
type HTTPClient struct { ... }

// Use this:
http.Get(url)
http.Post(url, contentType, body)
```

**3. Custom CSV Parser** - Use `encoding/csv`!
- `csv.NewReader(reader).ReadAll()`
- Handles quoted fields, embedded commas, newlines
- Custom parser breaks on edge cases

**4. Custom UUID Generator** - Use `github.com/google/uuid`!
- UUID spec is complex
- Need cryptographically random values
- Custom version won't follow RFC 4122

**5. Custom Sorting** - Use `sort` package!
```go
// Don't write bubble sort
// Use: sort.Ints(arr)
```

**6. Custom Base64 Encoding** - Use `encoding/base64`!
- Padding rules are tricky
- URL-safe variants exist
- Built-in is optimized

**7. Custom Template Engine** - Use `text/template`!
- `template.New("t").Parse(tmpl).Execute(wr, data)`
- Handles escaping, nested templates, functions
- Security built-in

### Why This Matters

**Impact:**
- **Time Waste**: Weeks/months on already-solved problems
- **Bugs**: Custom implementations have edge-case bugs
- **Maintenance**: You own the bugs and updates
- **Performance**: stdlib/libraries are optimized
- **Security**: Mature libraries have had security reviews
- **Team Knowledge**: Others must learn your custom version
- **No Updates**: Won't get improvements from community

**Detection:**
- Code review finds suspicious implementations
- "We have a custom X" for common functionality
- Large utility files/packages
- Avoiding standard library imports
- NIH (Not Invented Here) syndrome

**Solution:**
- **Check stdlib first**: Most languages have rich standard libraries
- **Research existing solutions**: PyPI, Go packages, npm
- **Benchmark before custom**: Prove stdlib is insufficient
- **Use battle-tested code**: Millions of users vs. just you
- **Focus on business logic**: Not infrastructure

**Valid Reasons to Reinvent:**
- Specific requirements stdlib can't meet (after investigation)
- Performance critical path (after profiling shows bottleneck)
- Dependency minimization for good reason (embedded systems, security)
- Educational purposes (clearly marked as such)

---

## 10. Leaky Abstractions

**Definition:** "All non-trivial abstractions, to some degree, are leaky." - Joel Spolsky

An abstraction that claims to hide implementation details but forces you to understand the underlying implementation to use it correctly or efficiently.

### Python Examples (`pythonexamples/leaky_abstractions.py`)

**1. ORM N+1 Query Problem** - The classic leak
```python
# Looks innocent:
for user in users:
    print(user.profile.avatar_url)  # N+1 queries!

# Must understand SQL JOINs:
users.select_related('profile')
```
- Abstraction promises to hide SQL
- Must understand JOINs, lazy loading, query optimization
- Performance traps hidden behind clean API

**2. ThreadPoolExecutor** - GIL leak
- Looks like parallel execution
- CPU-bound tasks don't actually run in parallel due to GIL
- Must understand threading vs multiprocessing
- Need implementation knowledge to use correctly

**3. Property Decorator** - Computation cost leak
```python
@property
def average(self):
    return sum(self._data) / len(self._data)
```
- Looks like attribute access
- Actually expensive O(n) computation
- No caching, recomputes every access
- Performance characteristics hidden

**4. HTTP Client** - Connection pooling leak
- Simple API hides connection reuse
- Must understand TCP connection lifecycle
- Keep-alive behavior affects performance
- Cannot reason about performance without understanding implementation

**5. List vs Generator** - Memory model leak
```python
[x*2 for x in range(1000000)]  # All in memory!
(x*2 for x in range(1000000))  # Lazy evaluation
```
- Similar syntax, vastly different behavior
- Must understand eager vs lazy evaluation
- Memory implications invisible from syntax

### Go Examples (`golangexamples/leaky_abstractions.go`)

**1. Interface Error Types** - Implementation leaks through errors
```go
type Cache interface { Get(); Set() }

// Leaks implementation:
if err == redis.ErrPoolExhausted { ... }
```
- Interface claims to hide implementation
- Error types reveal Redis underneath
- Code becomes coupled to implementation

**2. io.Reader Buffering** - Buffer management leak
```go
reader := bufio.NewReader(file)
reader.ReadString('\n')
```
- Must understand buffer size for performance
- EOF semantics leak from underlying io.Reader
- Memory vs performance tradeoffs exposed

**3. sync.Pool** - GC behavior leak
```go
obj := pool.Get()
pool.Put(obj)
obj2 := pool.Get()  // Might not be same object!
```
- Looks like guaranteed object reuse
- GC can clear pool at any time
- Must understand Go memory management

**4. Channel Buffering** - Goroutine semantics leak
```go
ch := make(chan int, 1)   // Buffered
ch := make(chan int)      // Unbuffered
```
- Buffer size dramatically changes blocking behavior
- Must understand channel semantics for correct use
- Implementation detail affects correctness

**5. context.Context** - Cancellation propagation leak
```go
func Process(ctx context.Context) {
    // When to check ctx.Done()? How often?
    // What's the difference between deadline and timeout?
}
```
- Simple API hides complex cancellation semantics
- Must understand propagation rules
- Error types leak implementation details

**6. http.Client** - Connection pool leak
```go
client := &http.Client{
    MaxIdleConnsPerHost: 2,  // Must tune this!
}
```
- Simple Get() hides connection pooling
- Must understand connection reuse for performance
- TLS session resumption, keep-alive invisible

### Why This Matters

**Impact:**
- **Surprising Performance**: Code looks fast but is slow
- **Requires Deep Knowledge**: Must understand implementation to use correctly
- **Defeats Abstraction Purpose**: If you need to know internals, why abstract?
- **Production Issues**: Works in dev, fails at scale
- **Debugging Complexity**: "Why is this slow?" requires understanding layers

**Detection:**
- Performance problems that require implementation knowledge
- Documentation says "for best performance, do X"
- Must read source code to use correctly
- Surprising behavior at scale
- Magic configuration parameters with no clear guidance

**Good vs Bad Leaky Abstractions:**
- **Good**: Leaks are documented, predictable, and graceful
- **Bad**: Surprises, undocumented behavior, sharp edges
- **Acceptable**: Minor leaks for significant simplification
- **Unacceptable**: Must understand implementation for basic use

**Solution:**
- **Document the leaks**: Explain what leaks and why
- **Provide escape hatches**: Allow control when needed
- **Make leaks explicit**: Parameters like buffer size
- **Performance guides**: Document common pitfalls
- **Accept some leakage**: Perfect abstraction is impossible

**When Leaks Are OK:**
- Tradeoff is clearly worth it
- Leaks are well-documented
- 80% of use cases don't need to know
- Escape hatches exist for advanced users

---

## Summary: Choosing Anti-Patterns for Your Presentation

### Best for Demonstrating Technical Depth:

1. **Leaky Abstractions** - Shows architectural thinking, understanding of multiple layers
2. **God Object** - Demonstrates SOLID principles and design patterns knowledge
3. **Spaghetti Code** - Shows ability to recognize and refactor complex logic

### Most Relatable in Production:

1. **Leaky Abstractions** - Everyone has debugged ORM N+1 queries
2. **Hard Coding** - Secret leaks happen to everyone
3. **Copy-Paste Programming** - Common in rapid development phases

### Great Discussion Starters:

1. **Leaky Abstractions** - "When is a leak acceptable?"
2. **Golden Hammer** - "When is it OK to use your favorite tool?"
3. **Reinventing the Wheel** - "When should you write custom code?"

### Language-Specific Showcases:

**Python:**
- Leaky Abstractions (GIL, ORM)
- Property decorator issues
- List vs generator tradeoffs

**Go:**
- Leaky Abstractions (channels, interfaces)
- Error handling patterns
- Concurrency pitfalls

### Recommendation for Presentation:

**Primary Example:** Leaky Abstractions (10-15 minutes)
- Most nuanced and demonstrates senior-level thinking
- Great Python (ORM) and Go (interface/channel) examples
- Leads to architectural discussions

**Supporting Examples:** (2-3 minutes each)
- God Object (shows SOLID principles)
- Copy-Paste Programming (shows refactoring skills)
- Hard Coding (shows security awareness)

This gives a 20-25 minute presentation with depth and breadth.
