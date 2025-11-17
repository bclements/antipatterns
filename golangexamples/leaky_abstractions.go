package main

/*
ANTI-PATTERN: Leaky Abstractions

"All non-trivial abstractions, to some degree, are leaky." - Joel Spolsky

An abstraction that claims to hide implementation details but forces you to understand
the underlying implementation to use it correctly or efficiently.
*/

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

// ==============================================================================
// Example 1: Interface Leaking Through Error Types
// ==============================================================================

// Cache interface claims to abstract storage, but errors leak the implementation
type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
}

// RedisCache implementation
type RedisCache struct {
	connected bool
}

// LEAK: Redis-specific errors
var (
	ErrRedisPoolExhausted = errors.New("redis: connection pool exhausted")
	ErrRedisTimeout       = errors.New("redis: operation timeout")
	ErrRedisConnectionLost = errors.New("redis: connection lost")
)

func (r *RedisCache) Get(key string) (string, error) {
	if !r.connected {
		return "", ErrRedisConnectionLost // LEAK: Implementation-specific error
	}
	return "value", nil
}

func (r *RedisCache) Set(key string, value string) error {
	if !r.connected {
		return ErrRedisPoolExhausted // LEAK: Implementation-specific error
	}
	return nil
}

func (r *RedisCache) Delete(key string) error {
	return nil
}

// LEAKY ABSTRACTION: Code that uses Cache is coupled to Redis implementation
func ProcessDataBAD(cache Cache) error {
	/*
		PROBLEM: The interface claims to hide the implementation, but we're
		forced to check for Redis-specific errors!

		You must understand:
		- Redis connection pooling
		- Redis-specific error conditions
		- Implementation details that should be hidden

		The abstraction leaks through error types.
	*/
	err := cache.Set("key", "value")

	// LEAK: Checking for implementation-specific error!
	if err == ErrRedisPoolExhausted {
		// Now we're coupled to Redis - the abstraction has leaked
		fmt.Println("Redis pool exhausted - retry logic")
		time.Sleep(100 * time.Millisecond)
		return cache.Set("key", "value")
	}

	// LEAK: Another Redis-specific error check
	if err == ErrRedisTimeout {
		fmt.Println("Redis timeout - adjust timeout settings")
	}

	return err
}

// Better approach: Use error wrapping to hide implementation
type CacheError struct {
	Op  string
	Err error
}

func (e *CacheError) Error() string {
	return fmt.Sprintf("cache %s: %v", e.Op, e.Err)
}

func (e *CacheError) Unwrap() error {
	return e.Err
}

// ==============================================================================
// Example 2: io.Reader/Writer Leaking Buffer Management
// ==============================================================================

// DataReader claims to abstract data reading, but leaks buffer behavior
type DataReader struct {
	source io.Reader
	buffer []byte
}

func NewDataReader(source io.Reader, bufferSize int) *DataReader {
	return &DataReader{
		source: source,
		buffer: make([]byte, bufferSize),
	}
}

func (dr *DataReader) ReadNext() (string, error) {
	/*
		LEAK: The abstraction looks simple, but you need to understand:
		- Buffer size affects performance dramatically
		- EOF handling semantics from io.Reader
		- When buffers are allocated vs reused
		- Memory implications of buffer size

		The abstraction leaks io.Reader semantics and buffer management.
	*/
	n, err := dr.source.Read(dr.buffer)
	if err != nil && err != io.EOF {
		return "", err
	}
	return string(dr.buffer[:n]), nil
}

func ProcessFileBAD() {
	/*
		PROBLEM: This looks simple, but performance completely depends on
		understanding buffer management which should be hidden!
	*/
	reader := NewDataReader(strings.NewReader("hello world"), 1)

	// LEAK: Tiny buffer size causes many system calls
	// You need to understand this to use the abstraction efficiently!
	data, _ := reader.ReadNext()
	fmt.Println(data)
}

func ProcessFileGOOD() {
	/*
		SOLUTION: Use appropriate buffer size, but this means understanding
		the underlying I/O model. The abstraction leaked!
	*/
	reader := NewDataReader(strings.NewReader("hello world"), 4096)

	data, _ := reader.ReadNext()
	fmt.Println(data)
}

// ==============================================================================
// Example 3: sync.Pool Leaking Memory Semantics
// ==============================================================================

// ObjectPool claims to manage object lifecycle, but leaks GC behavior
type ObjectPool struct {
	pool *sync.Pool
}

func NewObjectPool() *ObjectPool {
	return &ObjectPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &ExpensiveObject{}
			},
		},
	}
}

type ExpensiveObject struct {
	Data [1024]byte
}

func (op *ObjectPool) Get() *ExpensiveObject {
	/*
		LEAK: This looks like simple object pooling, but you must understand:
		- sync.Pool can evict objects at any GC
		- Objects are not guaranteed to be available
		- Memory semantics differ from manual pooling
		- No size limits - can grow unbounded

		The abstraction leaks Go's GC behavior and sync.Pool semantics.
	*/
	return op.pool.Get().(*ExpensiveObject)
}

func (op *ObjectPool) Put(obj *ExpensiveObject) {
	op.pool.Put(obj)
}

func UsePoolBAD() {
	/*
		PROBLEM: This looks like guaranteed object reuse, but sync.Pool
		doesn't guarantee anything! GC can clear it at any time.
	*/
	pool := NewObjectPool()

	obj := pool.Get()
	// Use object
	pool.Put(obj)

	// LEAK: You assume Get() returns the same object, but it might not!
	// You need to understand sync.Pool behavior to use it correctly.
	obj2 := pool.Get()
	if obj2 != obj {
		fmt.Println("Different object - sync.Pool was cleared!")
	}
}

// ==============================================================================
// Example 4: http.Client Leaking Connection Management
// ==============================================================================

// HTTPClient claims to abstract HTTP, but leaks connection pooling
type HTTPClient struct {
	// LEAK: These fields leak implementation details users need to understand
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 2,
		IdleConnTimeout:     90 * time.Second,
	}
}

func (hc *HTTPClient) Get(url string) (string, error) {
	/*
		LEAK: Simple Get() hides complex connection pooling behavior.

		You must understand:
		- Connection pooling and reuse
		- Keep-alive behavior
		- DNS caching
		- TLS session resumption
		- When connections are closed

		The abstraction leaks TCP connection management.
	*/
	fmt.Printf("[HTTP] GET %s (using connection pool)\n", url)
	return "response", nil
}

func MakeRequestsBAD() {
	/*
		PROBLEM: Looks simple, but connection behavior is completely opaque.
		Performance depends on understanding connection pooling!
	*/
	client := NewHTTPClient()

	// LEAK: Are these using the same connection? New connections?
	// You can't tell without understanding the implementation!
	for i := 0; i < 10; i++ {
		_, _ = client.Get(fmt.Sprintf("http://api.example.com/users/%d", i))
	}

	// LEAK: What happens to idle connections? When are they closed?
	// You need to understand MaxIdleConnsPerHost to tune performance!
}

// ==============================================================================
// Example 5: Channel Buffering Leaking Goroutine Behavior
// ==============================================================================

// WorkQueue claims to abstract async work, but leaks channel semantics
type WorkQueue struct {
	tasks chan func()
}

func NewWorkQueue(bufferSize int) *WorkQueue {
	/*
		LEAK: Buffer size dramatically affects behavior, but this is
		supposed to be an implementation detail!

		You must understand:
		- Buffered vs unbuffered channels
		- Blocking behavior
		- Goroutine scheduling
		- Memory usage of buffered channels
	*/
	return &WorkQueue{
		tasks: make(chan func(), bufferSize),
	}
}

func (wq *WorkQueue) Submit(task func()) error {
	/*
		LEAK: This looks like it always succeeds, but it blocks!
		You need to understand channel semantics to use it correctly.
	*/
	select {
	case wq.tasks <- task:
		return nil
	default:
		return fmt.Errorf("queue full")
	}
}

func (wq *WorkQueue) Start() {
	go func() {
		for task := range wq.tasks {
			task()
		}
	}()
}

func UseWorkQueueBAD() {
	/*
		PROBLEM: Behavior completely changes based on buffer size,
		but this is supposed to be hidden by the abstraction!
	*/
	queue := NewWorkQueue(1) // Tiny buffer
	queue.Start()

	// LEAK: This works fine
	_ = queue.Submit(func() { fmt.Println("Task 1") })

	// LEAK: This might fail because buffer is full!
	// You need to understand channel buffering to use this correctly.
	err := queue.Submit(func() { fmt.Println("Task 2") })
	if err != nil {
		fmt.Println("Queue full - needed to understand buffering!")
	}
}

func UseWorkQueueGOOD() {
	/*
		SOLUTION: Use larger buffer, but this means understanding
		the channel implementation. The abstraction leaked!
	*/
	queue := NewWorkQueue(1000)
	queue.Start()

	for i := 0; i < 100; i++ {
		_ = queue.Submit(func() { fmt.Println("Task") })
	}
}

// ==============================================================================
// Example 6: context.Context Leaking Cancellation Propagation
// ==============================================================================

// Service claims to abstract business logic, but leaks context semantics
type Service struct {
	name string
}

func (s *Service) ProcessRequest(ctx context.Context, request string) error {
	/*
		LEAK: This looks simple, but you must understand:
		- Context cancellation propagation
		- Deadline vs timeout vs cancellation
		- When to check context.Done()
		- Context value propagation

		The abstraction leaks context cancellation semantics.
	*/
	select {
	case <-ctx.Done():
		// LEAK: You need to check this, but when? How often?
		// The abstraction doesn't tell you!
		return ctx.Err()
	case <-time.After(100 * time.Millisecond):
		fmt.Printf("[%s] Processed: %s\n", s.name, request)
		return nil
	}
}

func CallServiceBAD() {
	/*
		PROBLEM: Context cancellation behavior is opaque.
		You need to understand context semantics to use this correctly!
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	service := &Service{name: "UserService"}

	// LEAK: This will timeout, but the error tells you nothing about
	// whether it was context timeout, cancellation, or deadline!
	err := service.ProcessRequest(ctx, "data")
	if err != nil {
		// LEAK: Is this context.DeadlineExceeded or context.Canceled?
		// You need to understand context error types!
		fmt.Printf("Error: %v\n", err)
	}
}

// ==============================================================================
// Example 7: Database Abstraction Leaking Transaction Semantics
// ==============================================================================

// DB claims to abstract database, but leaks transaction behavior
type DB struct {
	inTransaction bool
}

func (db *DB) Query(query string) ([]map[string]interface{}, error) {
	/*
		LEAK: This looks simple, but you must understand:
		- Transaction isolation levels
		- Lock behavior
		- Rollback semantics
		- Connection pooling

		The abstraction leaks SQL transaction semantics.
	*/
	fmt.Printf("[DB] Query: %s (in_transaction: %v)\n", query, db.inTransaction)
	return nil, nil
}

func (db *DB) Begin() error {
	/*
		LEAK: What isolation level? What happens to concurrent queries?
		You need to understand transaction semantics!
	*/
	db.inTransaction = true
	return nil
}

func (db *DB) Commit() error {
	db.inTransaction = false
	return nil
}

func (db *DB) Rollback() error {
	db.inTransaction = false
	return nil
}

func UseDBBAD() {
	/*
		PROBLEM: Transaction behavior is opaque. What happens if you
		forget to commit? What isolation level? The abstraction leaks!
	*/
	db := &DB{}

	_ = db.Begin()
	_, _ = db.Query("UPDATE users SET balance = balance - 100 WHERE id = 1")

	// LEAK: What happens if we crash here? Auto-rollback? Committed?
	// You need to understand transaction semantics!

	_ = db.Commit()
}

// ==============================================================================
// MAIN - Demonstrate Leaky Abstractions
// ==============================================================================

func main() {
	fmt.Println("=" + strings.Repeat("=", 79))
	fmt.Println("LEAKY ABSTRACTIONS ANTI-PATTERN")
	fmt.Println("=" + strings.Repeat("=", 79))
	fmt.Println()

	fmt.Println("Example 1: Interface Leaking Through Error Types")
	fmt.Println(strings.Repeat("-", 80))
	cache := &RedisCache{connected: false}
	_ = ProcessDataBAD(cache)
	fmt.Println()

	fmt.Println("Example 2: io.Reader Leaking Buffer Management")
	fmt.Println(strings.Repeat("-", 80))
	ProcessFileBAD()
	fmt.Println()

	fmt.Println("Example 3: sync.Pool Leaking Memory Semantics")
	fmt.Println(strings.Repeat("-", 80))
	UsePoolBAD()
	fmt.Println()

	fmt.Println("Example 5: Channel Buffering Leaking Goroutine Behavior")
	fmt.Println(strings.Repeat("-", 80))
	UseWorkQueueBAD()
	time.Sleep(200 * time.Millisecond) // Wait for goroutines
	fmt.Println()

	fmt.Println("Example 6: context.Context Leaking Cancellation")
	fmt.Println(strings.Repeat("-", 80))
	CallServiceBAD()
	fmt.Println()

	fmt.Println("Example 7: Database Leaking Transaction Semantics")
	fmt.Println(strings.Repeat("-", 80))
	UseDBBAD()
	fmt.Println()

	fmt.Println("=" + strings.Repeat("=", 79))
	fmt.Println("KEY INSIGHT: All abstractions leak to some degree.")
	fmt.Println("Good abstractions leak gracefully and document what leaks.")
	fmt.Println("Bad ones force you to understand implementation details to use")
	fmt.Println("them correctly, defeating the purpose of abstraction.")
	fmt.Println("=" + strings.Repeat("=", 79))
}
