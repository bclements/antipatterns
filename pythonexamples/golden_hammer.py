"""
ANTI-PATTERN: Golden Hammer

Using the same solution/tool for every problem regardless of whether it's appropriate.
"If all you have is a hammer, everything looks like a nail."
"""

import re

class RegexFanatic:
    """Someone who learned regex and now uses it for EVERYTHING"""

    def validate_email(self, email):
        # OK: Regex is appropriate here
        pattern = r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
        return bool(re.match(pattern, email))

    def is_even(self, number):
        # GOLDEN HAMMER: Using regex to check if a number is even!
        # Should just use: number % 2 == 0
        return bool(re.match(r'^-?\d*[02468]$', str(number)))

    def add_numbers(self, a, b):
        # GOLDEN HAMMER: Using regex for basic arithmetic!
        # Just use: a + b
        result = str(a) + str(b)
        return int(re.sub(r'(\d+)(\d+)', r'\1\2', result))  # Makes no sense!

    def get_list_length(self, items):
        # GOLDEN HAMMER: Using regex to count list items
        # Should just use: len(items)
        string_rep = str(items)
        matches = re.findall(r',', string_rep)
        return len(matches) + 1 if matches or string_rep != '[]' else 0

    def reverse_string(self, text):
        # GOLDEN HAMMER: Using regex for string reversal
        # Should just use: text[::-1]
        result = text
        for i in range(len(text)):
            pattern = r'^(.{' + str(i) + r'})(.)'
            result = re.sub(pattern, r'\2\1', result)
        return result  # This doesn't even work correctly!


class InheritanceEverything:
    """Someone who uses inheritance for every code reuse scenario"""

    pass

# GOLDEN HAMMER: Using inheritance when composition would be better
class Animal:
    def breathe(self):
        return "breathing"

# A car is NOT an animal, but we're using inheritance just to reuse the breathe logic
class Car(Animal):
    """GOLDEN HAMMER: Car inheriting from Animal just to reuse code!"""
    def drive(self):
        # Now we have a car that can breathe...
        return f"driving while {self.breathe()}"

# GOLDEN HAMMER: Everything must be a class!
class MathOperations:
    """Using classes when simple functions would suffice"""

    @staticmethod
    def add(a, b):
        return a + b

    @staticmethod
    def subtract(a, b):
        return a - b

    @staticmethod
    def multiply(a, b):
        return a * b

    # Should just be simple functions, not a class!


# GOLDEN HAMMER: Using design patterns everywhere
class SingletonFactory:
    """Combining Singleton and Factory patterns when you just need a simple function"""
    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
        return cls._instance

    def create_string(self, text):
        """This is just creating a string... why is it a singleton factory?"""
        return str(text)

    def create_number(self, num):
        """Just returns a number..."""
        return int(num)


# GOLDEN HAMMER: Using a database for everything
class DatabaseForEverything:
    """Someone who puts everything in a database, even temporary data"""

    def __init__(self, db_connection):
        self.db = db_connection

    def calculate_sum(self, a, b):
        # GOLDEN HAMMER: Storing calculation in database instead of just returning it
        result = a + b
        self.db.execute("INSERT INTO calculations (operation, result) VALUES ('add', ?)", (result,))
        return result

    def temporary_variable(self, value):
        # GOLDEN HAMMER: Storing temporary variable in database!
        self.db.execute("INSERT INTO temp_vars (value) VALUES (?)", (value,))
        row = self.db.execute("SELECT value FROM temp_vars WHERE id = LAST_INSERT_ID()").fetchone()
        self.db.execute("DELETE FROM temp_vars WHERE id = LAST_INSERT_ID()")
        return row[0]


# GOLDEN HAMMER: Using microservices for everything
"""
The developer learned about microservices and now wants to split
a simple application into 50 microservices:

- UserEmailValidationService
- UserFirstNameService
- UserLastNameService
- AddTwoNumbersService
- SubtractTwoNumbersService
- StringToUpperCaseService
- StringToLowerCaseService

Each service has its own database, API, and deployment pipeline
for functionality that should just be simple functions!
"""
