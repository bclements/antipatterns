"""
ANTI-PATTERN: Leaky Abstractions

"All non-trivial abstractions, to some degree, are leaky." - Joel Spolsky

An abstraction that claims to hide implementation details but forces you to understand
the underlying implementation to use it correctly or efficiently.
"""

# ==============================================================================
# Example 1: ORM - The Classic Leaky Abstraction
# ==============================================================================

class User:
    """Simulated ORM model"""
    def __init__(self, id, name, company_id):
        self.id = id
        self.name = name
        self.company_id = company_id
        self._profile = None

    @property
    def profile(self):
        if self._profile is None:
            # LEAK: This looks like attribute access but triggers a database query!
            print(f"[DB QUERY] SELECT * FROM profiles WHERE user_id = {self.id}")
            self._profile = Profile(self.id)
        return self._profile


class Profile:
    """Simulated profile model"""
    def __init__(self, user_id):
        self.user_id = user_id
        self.avatar_url = f"https://cdn.example.com/avatar_{user_id}.png"
        self.bio = "Sample bio"


class QuerySet:
    """Simulated ORM QuerySet"""
    def __init__(self, data):
        self.data = data
        self._prefetched = False

    def filter(self, **kwargs):
        print(f"[DB QUERY] SELECT * FROM users WHERE {kwargs}")
        return self

    def select_related(self, *fields):
        """
        LEAK: You need to understand SQL JOINs to use the ORM efficiently!
        The abstraction claims to hide SQL but you must know about JOINs.
        """
        print(f"[DB QUERY] SELECT * FROM users LEFT JOIN {fields}")
        self._prefetched = True
        return self

    def all(self):
        return self.data


# LEAKY ABSTRACTION: The ORM looks simple but hides performance traps
def get_users_with_avatars_BAD(company_id):
    """
    PROBLEM: This looks clean but causes N+1 query problem!

    You must understand:
    - SQL query execution
    - Lazy loading behavior
    - Database round-trip costs

    The abstraction leaks because simple Python code causes hidden database queries.
    """
    users = QuerySet([
        User(1, "Alice", company_id),
        User(2, "Bob", company_id),
        User(3, "Charlie", company_id)
    ]).filter(company_id=company_id).all()

    avatars = []
    for user in users:
        # LEAK: Each access triggers a hidden database query!
        # Looks like simple attribute access, but it's not.
        avatars.append(user.profile.avatar_url)

    return avatars


def get_users_with_avatars_GOOD(company_id):
    """
    SOLUTION: You must understand SQL to use the ORM correctly.
    The abstraction has leaked - you need implementation knowledge.
    """
    users = QuerySet([
        User(1, "Alice", company_id),
        User(2, "Bob", company_id),
        User(3, "Charlie", company_id)
    ]).filter(company_id=company_id).select_related('profile').all()

    avatars = []
    for user in users:
        avatars.append(user.profile.avatar_url)

    return avatars


# ==============================================================================
# Example 2: File Context Manager - Leaks Buffer Behavior
# ==============================================================================

class BufferedFileReader:
    """
    LEAKY ABSTRACTION: Claims to hide file I/O complexity but leaks buffering behavior
    """
    def __init__(self, filename, buffer_size=8192):
        self.filename = filename
        self.buffer_size = buffer_size
        self.file = None

    def __enter__(self):
        self.file = open(self.filename, 'r')
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        if self.file:
            self.file.close()

    def read_line(self):
        """
        LEAK: This looks like it reads one line, but the buffering behavior
        means it actually reads buffer_size bytes from disk!

        Users must understand:
        - OS-level buffering
        - Filesystem caching
        - Memory vs disk I/O tradeoffs
        """
        return self.file.readline()


def process_large_file_BAD():
    """
    PROBLEM: The abstraction leaks - you need to understand buffering
    to know why this might be slow or memory-intensive.
    """
    with BufferedFileReader('/tmp/large_file.txt') as reader:
        line = reader.read_line()
        # LEAK: Even though you only read one line, the entire buffer
        # was loaded. This is invisible to the user!
        return line


def process_large_file_GOOD():
    """
    SOLUTION: You must understand the buffering to use it efficiently.
    The abstraction leaked - you need implementation knowledge.
    """
    # You need to know about buffer_size to tune performance
    with BufferedFileReader('/tmp/large_file.txt', buffer_size=64) as reader:
        line = reader.read_line()
        return line


# ==============================================================================
# Example 3: Thread Pool Executor - Leaks Threading Model
# ==============================================================================

from concurrent.futures import ThreadPoolExecutor
import time

def expensive_io_task(task_id):
    """Simulates I/O-bound work"""
    time.sleep(0.1)
    return f"Task {task_id} complete"


def cpu_intensive_task(n):
    """
    LEAKY ABSTRACTION: ThreadPoolExecutor looks great for parallelism,
    but leaks the GIL (Global Interpreter Lock) limitation!

    LEAK: You must understand:
    - Python's GIL prevents true parallel CPU work
    - Thread vs Process differences
    - When to use ThreadPoolExecutor vs ProcessPoolExecutor
    """
    total = 0
    for i in range(n):
        total += i ** 2
    return total


def parallel_work_BAD():
    """
    PROBLEM: This looks like it will run in parallel, but the GIL means
    CPU-bound tasks won't actually benefit from multiple threads!

    The abstraction leaks - you need to understand the GIL.
    """
    with ThreadPoolExecutor(max_workers=4) as executor:
        # LEAK: These look like they'll run in parallel, but GIL prevents it!
        futures = [executor.submit(cpu_intensive_task, 1000000) for _ in range(4)]
        results = [f.result() for f in futures]
    return results


def parallel_work_GOOD():
    """
    SOLUTION: For I/O-bound work, threads work great.
    But you needed to understand the GIL to know the difference!
    """
    with ThreadPoolExecutor(max_workers=4) as executor:
        # These actually run in parallel because they're I/O-bound
        futures = [executor.submit(expensive_io_task, i) for i in range(4)]
        results = [f.result() for f in futures]
    return results


# ==============================================================================
# Example 4: Requests Library - Leaks Connection Pooling
# ==============================================================================

class HTTPClient:
    """
    LEAKY ABSTRACTION: Simple HTTP client that leaks connection management
    """
    def __init__(self):
        self._connections = {}

    def get(self, url):
        """
        LEAK: This looks like a simple GET request, but there's hidden
        connection pooling, timeout behavior, and retry logic.

        Users must understand:
        - TCP connection lifecycle
        - Connection pooling and reuse
        - Timeout vs connection timeout vs read timeout
        - DNS caching
        """
        print(f"[HTTP] GET {url}")
        # Simulated connection pooling (hidden from user)
        host = url.split('/')[2]
        if host not in self._connections:
            print(f"[TCP] Opening new connection to {host}")
            self._connections[host] = "connection_object"
        else:
            print(f"[TCP] Reusing existing connection to {host}")

        return {"status": 200, "data": "response"}


def make_requests_BAD():
    """
    PROBLEM: Looks simple, but performance depends on understanding
    connection pooling, which is hidden by the abstraction.
    """
    client = HTTPClient()

    # LEAK: Each call reuses connections, but you can't see it!
    # You need to understand this to reason about performance
    for i in range(10):
        response = client.get(f"https://api.example.com/users/{i}")

    # LEAK: What happens to connections when client goes out of scope?
    # Are they closed? Kept alive? You need to know!


def make_requests_GOOD():
    """
    SOLUTION: Use session explicitly, but this means understanding
    the underlying HTTP connection model. The abstraction leaked!
    """
    client = HTTPClient()

    # You understand connection reuse and manage it explicitly
    try:
        for i in range(10):
            response = client.get(f"https://api.example.com/users/{i}")
    finally:
        # Explicitly clean up (if the API provided this)
        pass


# ==============================================================================
# Example 5: List Comprehension vs Generator - Leaks Memory Model
# ==============================================================================

def process_large_dataset_BAD(data_size=1000000):
    """
    LEAKY ABSTRACTION: List comprehension looks like generator but isn't!

    LEAK: You must understand:
    - Memory allocation for lists
    - Eager vs lazy evaluation
    - Iterator protocol

    The abstraction leaks because similar syntax has wildly different behavior.
    """
    # LEAK: This creates entire list in memory immediately!
    # Looks innocent but could cause OOM for large datasets
    numbers = [x * 2 for x in range(data_size)]

    # Now you have 1 million integers in memory
    return sum(numbers)


def process_large_dataset_GOOD(data_size=1000000):
    """
    SOLUTION: Use generator, but you needed to understand memory model
    to know when to use which! The abstraction leaked.
    """
    # Generator expression - lazy evaluation
    numbers = (x * 2 for x in range(data_size))

    # Memory efficient - but you needed to know this existed!
    return sum(numbers)


# ==============================================================================
# Example 6: Property Decorator - Leaks Computation Cost
# ==============================================================================

class DataAnalyzer:
    """
    LEAKY ABSTRACTION: @property makes methods look like attributes,
    hiding expensive computation.
    """
    def __init__(self, data):
        self._data = data

    @property
    def average(self):
        """
        LEAK: Looks like simple attribute access, but it's expensive computation!

        Users must understand:
        - Properties can run arbitrary code
        - Each access recomputes (no caching)
        - Performance characteristics are hidden
        """
        print("[EXPENSIVE] Calculating average...")
        return sum(self._data) / len(self._data)

    @property
    def sorted_data(self):
        """
        LEAK: Even worse - this sorts on every access!
        Looks like an attribute but has O(n log n) cost.
        """
        print("[EXPENSIVE] Sorting data...")
        return sorted(self._data)


def analyze_data_BAD():
    """
    PROBLEM: Looks like simple attribute access, but triggers
    expensive computation multiple times!
    """
    analyzer = DataAnalyzer(list(range(10000)))

    # LEAK: Each of these looks like attribute access but is expensive!
    print(analyzer.average)  # O(n) computation
    print(analyzer.average)  # O(n) computation AGAIN!
    print(analyzer.sorted_data)  # O(n log n) computation
    print(analyzer.sorted_data)  # O(n log n) computation AGAIN!

    # The abstraction leaks - you need to know it's not cached!


def analyze_data_GOOD():
    """
    SOLUTION: Cache the results, but you needed to understand
    the property was recomputing! Abstraction leaked.
    """
    analyzer = DataAnalyzer(list(range(10000)))

    # Explicitly cache because you know it's expensive
    avg = analyzer.average  # Compute once
    sorted_data = analyzer.sorted_data  # Compute once

    # Use cached values
    print(avg)
    print(sorted_data)


# ==============================================================================
# MAIN - Demonstrate the Leaky Abstractions
# ==============================================================================

def main():
    print("=" * 80)
    print("LEAKY ABSTRACTIONS ANTI-PATTERN")
    print("=" * 80)
    print()

    print("Example 1: ORM Leaky Abstraction (N+1 Query Problem)")
    print("-" * 80)
    print("BAD: Hidden N+1 queries")
    get_users_with_avatars_BAD(1)
    print()
    print("GOOD: Explicit join (but requires SQL knowledge!)")
    get_users_with_avatars_GOOD(1)
    print()

    print("=" * 80)
    print("Example 4: HTTP Client Connection Pooling Leak")
    print("-" * 80)
    make_requests_BAD()
    print()

    print("=" * 80)
    print("Example 6: Property Decorator Computation Cost Leak")
    print("-" * 80)
    analyze_data_BAD()
    print()

    print("=" * 80)
    print("KEY INSIGHT: All abstractions leak to some degree.")
    print("Good abstractions leak gracefully. Bad ones force you to")
    print("understand the implementation to use them correctly.")
    print("=" * 80)


if __name__ == "__main__":
    main()
