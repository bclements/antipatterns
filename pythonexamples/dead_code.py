"""
ANTI-PATTERN: Dead Code

Code that is never executed or called. It clutters the codebase and
makes it harder to maintain. Unlike Lava Flow (old code nobody dares remove),
this is code that provably never runs.
"""

# DEAD CODE: Function that's never called anywhere
def calculate_legacy_tax(amount):
    """This function is never called in the entire codebase"""
    return amount * 0.15


# DEAD CODE: Class that's never instantiated
class UnusedHelper:
    """Nobody creates instances of this class"""

    def __init__(self):
        self.value = 0

    def process(self):
        return self.value * 2


def process_order(order):
    """Main function with lots of dead code inside"""

    # DEAD CODE: Variable assigned but never used
    total_weight = 0

    # DEAD CODE: Unreachable code after return
    if order['status'] == 'cancelled':
        return None
        print("This will never print")  # DEAD CODE
        order['processed'] = True  # DEAD CODE

    # DEAD CODE: Condition that's never true
    if order.get('type') == 'quantum':  # 'quantum' type doesn't exist in the system
        quantum_process(order)  # DEAD CODE

    # DEAD CODE: Code after unconditional return
    if True:
        return order
        # Everything below is DEAD CODE
        order['timestamp'] = 'now'
        validate_order(order)

    # More DEAD CODE (unreachable)
    return process_backup(order)


# DEAD CODE: Helper function never called
def quantum_process(order):
    """This is called from dead code, so it's also dead code"""
    return order


# DEAD CODE: Helper function never called
def validate_order(order):
    """Never called anywhere"""
    return True


# DEAD CODE: Helper function never called
def process_backup(order):
    """Never called anywhere"""
    return order


class DataProcessor:
    """Class with dead code inside"""

    def __init__(self):
        # DEAD CODE: Attribute that's never read
        self.unused_counter = 0
        self.legacy_flag = False

    def process(self, data):
        # DEAD CODE: Assignment that's overwritten before use
        result = []
        result = data * 2

        # DEAD CODE: Condition that's always false
        if len(data) < 0:  # Length can never be negative
            return []  # DEAD CODE

        # DEAD CODE: Variable assigned but never used
        temp_value = result * 3

        # DEAD CODE: Unreachable due to previous return logic
        if result:
            return result
        else:
            # DEAD CODE: This else block is unreachable
            return self._fallback_process(data)

    # DEAD CODE: Method that's never called
    def _fallback_process(self, data):
        """Never called anywhere"""
        return data

    # DEAD CODE: Method that's never called
    def reset(self):
        """Never called anywhere"""
        self.unused_counter = 0


# DEAD CODE: Entire module-level code block
if __name__ == "__main__":
    # This looks like it might run, but this file is only ever imported
    # never executed as main, so all this code is dead
    print("Initializing...")
    processor = DataProcessor()
    processor.process([1, 2, 3])
    print("Done!")


def calculate_discount(amount, user_type):
    """Function with dead branches"""

    # DEAD CODE: These conditions are mutually exclusive with the code design
    if user_type == "premium":
        discount = 0.2
    elif user_type == "regular":
        discount = 0.1
    elif user_type == "premium":  # DEAD CODE: Already handled above
        discount = 0.25  # DEAD CODE

    # DEAD CODE: Unreachable return
    if amount > 0:
        return amount * (1 - discount)
    else:
        return 0

    # DEAD CODE: After all paths return
    print("Processing complete")


# DEAD CODE: Commented out code that should be deleted
# def old_implementation(data):
#     """This is the old way we did things"""
#     result = []
#     for item in data:
#         result.append(item * 2)
#     return result


# DEAD CODE: Import that's never used
from typing import Optional, Dict  # Dict is never used

# DEAD CODE: Constant that's never referenced
UNUSED_CONSTANT = 42
MAX_RETRIES = 3  # Never used
DEFAULT_TIMEOUT = 30  # Never used


class UserManager:
    """Class with dead code"""

    def __init__(self):
        self.users = []
        # DEAD CODE: Attribute never accessed
        self.admin_count = 0

    def add_user(self, user):
        self.users.append(user)

        # DEAD CODE: Variable assigned but never used
        timestamp = "2024-01-01"

        # DEAD CODE: Computation that's never used
        user_count = len(self.users)

        return user

    # DEAD CODE: Method never called
    def get_admin_count(self):
        """This method is never called"""
        return self.admin_count

    # DEAD CODE: Method never called
    def clear_users(self):
        """This method is never called"""
        self.users = []


def complex_function(x, y):
    """Function with multiple dead code issues"""

    # DEAD CODE: Early return makes everything below unreachable
    return x + y

    # All of this is DEAD CODE:
    if x > 100:
        result = x * 2
    else:
        result = y * 2

    temp = result / 2
    return temp


# DEAD CODE: Exception handler for exception that's never raised
def safe_divide(a, b):
    """Function with dead exception handler"""
    try:
        if b == 0:
            return 0  # We handle zero here
        return a / b
    except ZeroDivisionError:  # DEAD CODE: Never raised due to check above
        return 0  # DEAD CODE


# DEAD CODE: Loop that never executes
def process_items(items):
    """Function with dead loop"""

    if not items:
        return []

    result = []

    # DEAD CODE: We already returned if items is empty
    for item in []:  # Empty list loop
        result.append(item)  # DEAD CODE

    # DEAD CODE: Condition that's always False after above check
    if len(items) == 0:
        return None  # DEAD CODE

    return [x * 2 for x in items]


# DEAD CODE: Decorator that's never used
def unused_decorator(func):
    """This decorator is never applied to any function"""
    def wrapper(*args, **kwargs):
        print("Before")
        result = func(*args, **kwargs)
        print("After")
        return result
    return wrapper


# DEAD CODE: Global variable that's never read
GLOBAL_CONFIG = {
    'setting1': 'value1',
    'setting2': 'value2',
}  # Never accessed anywhere
