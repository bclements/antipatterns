"""
ANTI-PATTERN: Reinventing the Wheel

Reimplementing functionality that already exists in standard libraries or
well-established third-party packages. This wastes time and often results
in buggier, less efficient code.
"""

# REINVENTING THE WHEEL: Custom JSON parser
# Python has built-in json module!
def parse_json(json_string):
    """
    Custom JSON parser - DON'T DO THIS!
    Use: import json; json.loads(json_string)
    """
    result = {}
    # 200 lines of buggy JSON parsing logic...
    # Missing edge cases, security issues, poor performance
    return result


# REINVENTING THE WHEEL: Custom HTTP client
# Use requests library or httpx!
class HTTPClient:
    """
    Custom HTTP client - DON'T DO THIS!
    Use: import requests; requests.get(url)
    """

    def get(self, url):
        """Incomplete HTTP GET implementation"""
        # 100+ lines reimplementing what requests does better
        pass

    def post(self, url, data):
        """Incomplete HTTP POST implementation"""
        # Missing features: timeout, retries, SSL verification, etc.
        pass


# REINVENTING THE WHEEL: Custom date/time formatting
# Python has datetime.strftime()!
def format_date(year, month, day):
    """
    Custom date formatter - DON'T DO THIS!
    Use: from datetime import datetime; datetime(year, month, day).strftime('%Y-%m-%d')
    """
    # Buggy implementation that doesn't handle leap years, etc.
    months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun",
              "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]
    return f"{months[month-1]} {day}, {year}"


# REINVENTING THE WHEEL: Custom CSV parser
# Python has csv module!
def parse_csv(csv_string):
    """
    Custom CSV parser - DON'T DO THIS!
    Use: import csv; csv.reader(file)
    """
    lines = csv_string.split('\n')
    result = []
    for line in lines:
        # This breaks on quoted commas, embedded newlines, etc.
        result.append(line.split(','))
    return result


# REINVENTING THE WHEEL: Custom UUID generator
# Python has uuid module!
import random

def generate_uuid():
    """
    Custom UUID generator - DON'T DO THIS!
    Use: import uuid; uuid.uuid4()
    """
    # Not actually following UUID spec, not guaranteed unique
    chars = "0123456789abcdef"
    return ''.join(random.choice(chars) for _ in range(32))


# REINVENTING THE WHEEL: Custom sorting algorithm
# Python has sorted() and list.sort()!
def bubble_sort(arr):
    """
    Custom sorting - DON'T DO THIS!
    Use: sorted(arr) or arr.sort()
    """
    # Reimplementing a slow O(n²) algorithm when Python's sort is O(n log n)
    n = len(arr)
    for i in range(n):
        for j in range(0, n - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
    return arr


# REINVENTING THE WHEEL: Custom email validator
# Use email-validator library or regex!
def validate_email(email):
    """
    Custom email validator - DON'T DO THIS!
    Use: from email_validator import validate_email
    """
    # Overly simplistic validation missing many edge cases
    if '@' in email and '.' in email.split('@')[1]:
        return True
    return False


# REINVENTING THE WHEEL: Custom logging system
# Python has logging module!
class CustomLogger:
    """
    Custom logger - DON'T DO THIS!
    Use: import logging; logging.getLogger(__name__)
    """

    def __init__(self, filename):
        self.filename = filename

    def log(self, message):
        # Missing: log levels, formatting, rotation, handlers, filters
        with open(self.filename, 'a') as f:
            f.write(message + '\n')


# REINVENTING THE WHEEL: Custom configuration parser
# Use configparser, pyyaml, or toml!
def parse_config(config_file):
    """
    Custom config parser - DON'T DO THIS!
    Use: import configparser; configparser.ConfigParser()
    """
    config = {}
    with open(config_file) as f:
        for line in f:
            # Breaks on comments, sections, complex values, etc.
            if '=' in line:
                key, value = line.strip().split('=')
                config[key] = value
    return config


# REINVENTING THE WHEEL: Custom template engine
# Use Jinja2!
def render_template(template, context):
    """
    Custom template engine - DON'T DO THIS!
    Use: from jinja2 import Template; Template(template).render(context)
    """
    # Simplistic and buggy string replacement
    result = template
    for key, value in context.items():
        result = result.replace(f"{{{key}}}", str(value))
    return result


# REINVENTING THE WHEEL: Custom argument parser
# Use argparse!
def parse_arguments(args):
    """
    Custom CLI argument parser - DON'T DO THIS!
    Use: import argparse; argparse.ArgumentParser()
    """
    # Missing: help text, type validation, subcommands, defaults, etc.
    result = {}
    for arg in args:
        if arg.startswith('--'):
            parts = arg[2:].split('=')
            if len(parts) == 2:
                result[parts[0]] = parts[1]
    return result


# REINVENTING THE WHEEL: Custom hash function
# Use hashlib!
def hash_string(s):
    """
    Custom hash function - DON'T DO THIS!
    Use: import hashlib; hashlib.sha256(s.encode()).hexdigest()
    """
    # Not cryptographically secure, collision-prone
    hash_value = 0
    for char in s:
        hash_value = (hash_value * 31 + ord(char)) % (2**32)
    return hash_value


# REINVENTING THE WHEEL: Custom random string generator
# Use secrets module for security or random for non-security!
def generate_random_string(length):
    """
    Custom random string generator - DON'T DO THIS!
    Use: import secrets; secrets.token_urlsafe(length)
    """
    # Not cryptographically secure
    chars = "abcdefghijklmnopqrstuvwxyz0123456789"
    return ''.join(random.choice(chars) for _ in range(length))


# REINVENTING THE WHEEL: Custom URL parser
# Use urllib.parse!
def parse_url(url):
    """
    Custom URL parser - DON'T DO THIS!
    Use: from urllib.parse import urlparse
    """
    # Broken implementation missing many edge cases
    parts = url.split('://')
    protocol = parts[0] if len(parts) > 1 else ''
    rest = parts[1] if len(parts) > 1 else parts[0]

    domain = rest.split('/')[0]
    path = '/' + '/'.join(rest.split('/')[1:]) if '/' in rest else '/'

    return {'protocol': protocol, 'domain': domain, 'path': path}


# REINVENTING THE WHEEL: Custom retry logic
# Use tenacity library!
def retry_function(func, max_attempts=3):
    """
    Custom retry logic - DON'T DO THIS!
    Use: from tenacity import retry, stop_after_attempt
    """
    # Missing: exponential backoff, specific exception handling, logging
    for attempt in range(max_attempts):
        try:
            return func()
        except Exception:
            if attempt == max_attempts - 1:
                raise
            continue


# REINVENTING THE WHEEL: Custom caching decorator
# Use functools.lru_cache!
_cache = {}

def cache(func):
    """
    Custom cache decorator - DON'T DO THIS!
    Use: from functools import lru_cache; @lru_cache(maxsize=128)
    """
    # Missing: cache size limits, TTL, cache invalidation
    def wrapper(*args):
        if args not in _cache:
            _cache[args] = func(*args)
        return _cache[args]
    return wrapper


# REINVENTING THE WHEEL: Custom database ORM
# Use SQLAlchemy, Django ORM, or Peewee!
class CustomORM:
    """
    Custom ORM - DON'T DO THIS!
    Use: from sqlalchemy import create_engine, Column, Integer, String
    """

    def __init__(self, connection):
        self.connection = connection

    def save(self, obj):
        # Missing: relationships, migrations, query optimization, etc.
        pass

    def find(self, id):
        # Incomplete implementation
        pass


# REINVENTING THE WHEEL: Custom serialization
# Use pickle, json, or msgpack!
def serialize_object(obj):
    """
    Custom serialization - DON'T DO THIS!
    Use: import pickle; pickle.dumps(obj)
    """
    # Broken for complex objects, missing types, etc.
    return str(obj.__dict__)


# REINVENTING THE WHEEL: Custom markdown parser
# Use markdown library!
def parse_markdown(text):
    """
    Custom markdown parser - DON'T DO THIS!
    Use: import markdown; markdown.markdown(text)
    """
    # Handles only basic cases, missing most markdown features
    text = text.replace('**', '<strong>', 1).replace('**', '</strong>', 1)
    text = text.replace('*', '<em>', 1).replace('*', '</em>', 1)
    return text


"""
Why reinventing the wheel is bad:
1. Wastes development time
2. Results in buggy, incomplete implementations
3. Missing edge cases and security considerations
4. No community support or updates
5. Harder to maintain
6. Other developers have to learn your custom implementation
7. Often worse performance than battle-tested libraries

Always check if functionality exists in:
- Python standard library
- Well-established third-party packages (PyPI)
- Framework-specific utilities

Only create custom implementations when:
- You have very specific requirements not met by existing solutions
- Performance is critical and profiling shows existing solutions are bottlenecks
- You need to avoid dependencies for valid reasons (deployment size, security, etc.)
"""
