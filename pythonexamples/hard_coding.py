"""
ANTI-PATTERN: Hard Coding

Embedding configuration values, credentials, file paths, and other values
directly in the source code instead of using configuration files or environment variables.
Makes the code inflexible and difficult to deploy to different environments.
"""

class DatabaseConnection:
    """HARD CODING: Database credentials embedded in code"""

    def __init__(self):
        # HARD CODING: Credentials should be in environment variables or config files
        self.host = "prod-db-server-01.company.com"
        self.port = 5432
        self.username = "admin"
        self.password = "SuperSecret123!"  # SECURITY ISSUE!
        self.database = "production_db"

    def connect(self):
        # Now changing environments requires modifying source code
        connection_string = f"postgresql://{self.username}:{self.password}@{self.host}:{self.port}/{self.database}"
        return connection_string


class EmailService:
    """HARD CODING: Email configuration embedded"""

    def send_email(self, to, subject, body):
        # HARD CODING: SMTP settings should be configurable
        smtp_server = "smtp.gmail.com"
        smtp_port = 587
        smtp_username = "noreply@company.com"
        smtp_password = "EmailPassword123"  # SECURITY ISSUE!

        # HARD CODING: Email templates in code
        if subject == "welcome":
            body = """
            Welcome to our service!
            Thanks for signing up.
            Visit us at https://www.company.com
            """

        # Simulate sending email
        print(f"Sending email via {smtp_server}:{smtp_port}")


class APIClient:
    """HARD CODING: API endpoints and keys embedded"""

    def __init__(self):
        # HARD CODING: API configuration should be externalized
        self.base_url = "https://api.production.company.com/v1"
        self.api_key = "sk_live_abc123def456ghi789"  # SECURITY ISSUE!
        self.timeout = 30
        self.max_retries = 3

    def make_request(self, endpoint):
        # HARD CODING: Full URLs constructed with hardcoded base
        url = f"{self.base_url}/{endpoint}"
        headers = {
            "Authorization": f"Bearer {self.api_key}",
            "Content-Type": "application/json"
        }
        return url, headers


class FileManager:
    """HARD CODING: File paths embedded in code"""

    def save_file(self, filename, content):
        # HARD CODING: Absolute paths that won't work on other systems
        base_path = "/home/john/projects/myapp/uploads"
        file_path = f"{base_path}/{filename}"

        # HARD CODING: Log file path
        log_path = "/var/log/myapp/file_operations.log"

        print(f"Saving to {file_path}")
        print(f"Logging to {log_path}")

    def get_config(self):
        # HARD CODING: Configuration file path
        config_file = "/etc/myapp/config.json"
        return config_file


class PaymentProcessor:
    """HARD CODING: Payment gateway credentials and settings"""

    def process_payment(self, amount, card):
        # HARD CODING: Stripe API keys in code
        stripe_secret_key = "sk_live_51Hxxxxxxxxxxxxx"  # SECURITY ISSUE!
        stripe_public_key = "pk_live_51Hxxxxxxxxxxxxx"  # SECURITY ISSUE!

        # HARD CODING: Payment thresholds and rules
        if amount > 1000:
            # HARD CODING: Fee structure in code
            fee = amount * 0.029 + 0.30
        else:
            fee = amount * 0.035 + 0.30

        # HARD CODING: Currency
        currency = "USD"

        return amount + fee


class ApplicationConfig:
    """HARD CODING: All configuration in code"""

    # HARD CODING: Feature flags
    ENABLE_NEW_UI = True
    ENABLE_BETA_FEATURES = False
    ENABLE_DEBUG_MODE = False

    # HARD CODING: Business rules
    MAX_LOGIN_ATTEMPTS = 3
    SESSION_TIMEOUT_MINUTES = 30
    PASSWORD_MIN_LENGTH = 8

    # HARD CODING: Third-party service keys
    GOOGLE_MAPS_API_KEY = "AIzaSyXXXXXXXXXXXXXXXXXX"  # SECURITY ISSUE!
    SENDGRID_API_KEY = "SG.XXXXXXXXXXXXXXXX"  # SECURITY ISSUE!
    AWS_ACCESS_KEY = "AKIAXXXXXXXXXXXXXXXX"  # SECURITY ISSUE!
    AWS_SECRET_KEY = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"  # SECURITY ISSUE!

    # HARD CODING: URLs
    HOMEPAGE_URL = "https://www.company.com"
    API_DOCS_URL = "https://docs.company.com/api"
    SUPPORT_EMAIL = "support@company.com"

    # HARD CODING: Database settings
    DB_POOL_SIZE = 10
    DB_POOL_TIMEOUT = 30
    DB_MAX_OVERFLOW = 5


def send_notification(user_id, message):
    """HARD CODING: Notification service configuration"""

    # HARD CODING: Slack webhook
    slack_webhook = "https://hooks.slack.com/services/T00/B00/XXXXXXXXXXXXXXXX"

    # HARD CODING: Notification rules
    if len(message) > 100:
        message = message[:100] + "..."

    # HARD CODING: Channels
    channel = "#notifications"
    username = "Bot"

    print(f"Sending to {slack_webhook}")


def get_user_avatar(user_id):
    """HARD CODING: CDN and storage URLs"""

    # HARD CODING: CDN URL
    cdn_base = "https://cdn.company.com"

    # HARD CODING: Default avatar path
    default_avatar = f"{cdn_base}/images/default-avatar.png"

    # HARD CODING: S3 bucket
    s3_bucket = "company-user-avatars"

    return default_avatar


def rate_limit_check(user_id):
    """HARD CODING: Rate limiting rules in code"""

    # HARD CODING: Rate limits
    max_requests_per_minute = 60
    max_requests_per_hour = 1000
    max_requests_per_day = 10000

    # HARD CODING: IP whitelist
    whitelisted_ips = [
        "192.168.1.100",
        "10.0.0.50",
        "172.16.0.25"
    ]

    return True


class CacheManager:
    """HARD CODING: Cache configuration"""

    def __init__(self):
        # HARD CODING: Redis connection
        self.redis_host = "redis.production.company.com"
        self.redis_port = 6379
        self.redis_password = "RedisPassword123"  # SECURITY ISSUE!
        self.redis_db = 0

        # HARD CODING: Cache TTL values
        self.user_cache_ttl = 3600  # 1 hour
        self.product_cache_ttl = 1800  # 30 minutes
        self.session_cache_ttl = 7200  # 2 hours


def generate_report():
    """HARD CODING: Report generation settings"""

    # HARD CODING: Report title and format
    report_title = "Monthly Sales Report - Company Inc."
    report_format = "PDF"

    # HARD CODING: Colors and styling
    primary_color = "#0066CC"
    secondary_color = "#FF9900"
    font_family = "Arial"
    font_size = 12

    # HARD CODING: Company info
    company_name = "Company Inc."
    company_address = "123 Business St, City, State 12345"
    company_phone = "(555) 123-4567"
    company_email = "info@company.com"

    return report_title


# HARD CODING: Environment-specific URLs
if True:  # Should check actual environment
    # Everything assumes production
    AUTH_URL = "https://auth.company.com"
    API_URL = "https://api.company.com"
    WEB_URL = "https://www.company.com"
    SOCKET_URL = "wss://socket.company.com"


# HARD CODING: Validation rules
def validate_username(username):
    """HARD CODING: Validation rules should be configurable"""

    # HARD CODING: Character limits and rules
    if len(username) < 3:
        return False, "Username must be at least 3 characters"
    if len(username) > 20:
        return False, "Username must be at most 20 characters"

    # HARD CODING: Blocked usernames
    blocked_usernames = ["admin", "root", "administrator", "system"]
    if username.lower() in blocked_usernames:
        return False, "Username is not allowed"

    return True, "Valid"


"""
All of these hardcoded values should be:
1. Moved to environment variables (.env file)
2. Moved to configuration files (config.yaml, config.json)
3. Stored in a secrets management system (AWS Secrets Manager, HashiCorp Vault)
4. Made configurable per environment (dev, staging, production)

This makes the code more flexible, secure, and maintainable.
"""
