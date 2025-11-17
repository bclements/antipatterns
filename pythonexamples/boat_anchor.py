"""
ANTI-PATTERN: Boat Anchor

Code that serves no purpose but is kept around "just in case" it might be needed.
This weighs down the codebase like a boat anchor.
"""

class UserManager:
    def __init__(self):
        self.users = []
        # OLD: This was for the legacy XML export feature we removed 2 years ago
        # But keeping it just in case we need it again
        self.xml_exporter = XMLExporter()  # Never used anymore

        # OLD: This was for the backup system that was replaced
        # Keeping it commented out just in case
        # self.backup_handler = BackupHandler()

    def add_user(self, user):
        self.users.append(user)

        # OLD METHOD: Used to notify via email, now we use webhooks
        # But maybe we'll need email again someday?
        # self._send_email_notification(user)

        # NEW METHOD:
        self._send_webhook_notification(user)

    def _send_webhook_notification(self, user):
        print(f"Sending webhook for {user}")

    # BOAT ANCHOR: This hasn't been called in years but "we might need it"
    def _send_email_notification(self, user):
        """Legacy email notification - UNUSED"""
        smtp_server = "old.smtp.server.com"
        # 50 lines of email code that's never executed...
        pass

    # BOAT ANCHOR: Feature we planned but never implemented
    def export_to_xml(self):
        """Planned XML export feature - NEVER IMPLEMENTED"""
        # TODO: Implement this when we have time (been here for 3 years)
        pass


class XMLExporter:
    """BOAT ANCHOR: Legacy class that's instantiated but never used"""
    def __init__(self):
        self.config = self._load_legacy_config()

    def _load_legacy_config(self):
        # Complex initialization that happens every time but is never used
        return {"format": "xml", "version": "1.0"}

    def export(self, data):
        # This method is never called anywhere
        pass


# BOAT ANCHOR: Keeping this around "just in case we need to rollback"
OLD_DATABASE_CONNECTION_STRING = "mysql://oldserver:3306/legacy_db"
OLD_API_ENDPOINT = "http://deprecated.api.com/v1"
OLD_ENCRYPTION_KEY = "old_key_that_we_dont_use_anymore"
