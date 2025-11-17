"""
ANTI-PATTERN: Lava Flow

Dead code and forgotten functionality that remains in the codebase because
nobody knows if it's safe to remove it. Like lava that has cooled and hardened,
this code is now part of the landscape but serves no purpose.
"""

# LAVA FLOW: Nobody knows what this does or if it's still needed
# Found in codebase since 2015
def mysterious_legacy_function(data):
    """
    TODO: Figure out what this does
    NOTE: Don't touch this! Bob said it's important but Bob left in 2017
    """
    result = []
    for item in data:
        # Nobody understands this logic
        if item.get('legacy_flag'):
            processed = _process_legacy_data(item)
            if processed:
                result.append(processed)
    return result


def _process_legacy_data(item):
    """
    LAVA FLOW: Helper function that's never called anymore
    but we're afraid to delete it
    """
    # Magic numbers that nobody understands
    if item.get('value', 0) > 42:
        return item.get('value') * 1.337
    return None


class DataProcessor:
    """Modern data processor with lots of lava flow from previous implementations"""

    def __init__(self):
        # LAVA FLOW: Old configuration that might be important?
        self.legacy_mode = False  # Set to True for backwards compatibility
        self.use_old_algorithm = False  # TODO: Remove after migration (from 2018)
        self.enable_experimental_feature = True  # Was experimental in 2016

        # LAVA FLOW: Database connections that may or may not be used
        self.old_db_connection = None  # From when we used MySQL
        self.backup_db_connection = None  # For the backup system we removed

    def process_data(self, data):
        """Current implementation"""
        # LAVA FLOW: Check for legacy mode that's never actually True
        if self.legacy_mode:
            return self._legacy_process(data)  # Nobody knows what this does

        # LAVA FLOW: Feature flag that's always True
        if self.use_old_algorithm:
            return self._old_algorithm(data)  # Never executed

        # This is what actually runs
        return self._new_process(data)

    def _new_process(self, data):
        """The actual current implementation"""
        return [x * 2 for x in data]

    # LAVA FLOW: Legacy implementation from 2016
    def _legacy_process(self, data):
        """
        LEGACY PROCESSING LOGIC - DO NOT MODIFY
        Last modified: 2016-03-15 by John
        Note: This is critical for backwards compatibility with System X
        (System X was decommissioned in 2017)
        """
        result = []
        for item in data:
            # Complex logic that nobody understands
            temp = item * 1.5
            if temp > 100:
                temp = temp / 2
            result.append(temp)
        return result

    # LAVA FLOW: Old algorithm that's been replaced
    def _old_algorithm(self, data):
        """
        WARNING: This is the old algorithm
        Use _new_process instead
        Keeping this for reference
        (from 2018)
        """
        return [x * 1.5 for x in data]

    # LAVA FLOW: Method for a feature that was never released
    def export_to_xml(self, data):
        """
        Export data to XML format
        TODO: Complete implementation
        NOTE: Waiting for approval from stakeholders
        (stakeholder meeting was in 2015)
        """
        pass


# LAVA FLOW: Entire class that's no longer used
class OldUserManager:
    """
    DEPRECATED: Use NewUserManager instead
    Keeping this for backwards compatibility
    Last used: Unknown
    Safe to delete: Nobody knows
    """

    def __init__(self):
        self.users = {}

    def add_user(self, username, email):
        """Old way of adding users"""
        self.users[username] = {'email': email, 'created': '2015-01-01'}

    def get_user(self, username):
        """Old way of getting users"""
        return self.users.get(username)


# LAVA FLOW: Configuration from various eras
# Some are still used, some aren't, nobody knows which is which
LEGACY_CONFIG = {
    'old_api_endpoint': 'http://old-api.example.com',  # Dead endpoint
    'timeout': 30,  # Still used?
    'max_retries': 3,  # Maybe used?
    'use_cache': True,  # Probably used?
    'cache_ttl': 3600,  # Who knows?
    'enable_feature_x': False,  # Feature X removed in 2017
    'enable_feature_y': True,  # Feature Y never implemented
    'debug_mode': False,  # Safe to remove? Nobody knows
}

# LAVA FLOW: Constants that might be important
MAGIC_NUMBER = 42  # Don't change! (Why? Nobody remembers)
ANOTHER_MAGIC_NUMBER = 1337  # Used somewhere... maybe?
CRITICAL_THRESHOLD = 100  # This is critical! (For what?)


# LAVA FLOW: Commented code that's been there forever
def process_user_data(user):
    """Process user data"""

    # Old implementation - DO NOT DELETE
    # result = []
    # for key, value in user.items():
    #     if key != 'password':
    #         result.append({key: value})
    # return result

    # Newer implementation - USE THIS ONE
    # return {k: v for k, v in user.items() if k != 'password'}

    # Latest implementation (2019)
    # filtered = user.copy()
    # del filtered['password']
    # return filtered

    # Current implementation (2022)
    return {k: v for k, v in user.items() if k not in ['password', 'secret']}


# LAVA FLOW: Emergency fix from 2016 that became permanent
def emergency_fix_for_bug_123(data):
    """
    EMERGENCY FIX - 2016-07-15
    TODO: Remove this after proper fix is implemented
    BUG: https://bugtracker.old-company.com/123 (link dead)
    DO NOT REMOVE: This is critical! (Why? Unknown)
    """
    if data is None:
        return []
    return data


# LAVA FLOW: Import statements for libraries we no longer use
# import oldframework  # Removed from requirements.txt in 2018 but import still here
# from legacy_lib import LegacyClass  # Nobody knows if we still need this


# LAVA FLOW: Global variables from the dawn of time
_global_cache = {}  # Used? Unknown
_instance = None  # Singleton pattern from 2014
_initialized = False  # Is this checked anywhere?


class ModernClass:
    """Modern implementation with lava flow sprinkled throughout"""

    def __init__(self):
        # LAVA FLOW: Initialize things we might need?
        global _initialized
        if not _initialized:
            self._setup_legacy_compatibility()
            _initialized = True

    def _setup_legacy_compatibility(self):
        """
        LAVA FLOW: Setup legacy compatibility layer
        NOTE: This was for version 1.x clients
        (Version 1.x was sunset in 2019)
        """
        pass

    def process(self, data):
        """Modern process method"""
        # LAVA FLOW: Safety check that's never been triggered
        if hasattr(data, 'legacy_format'):
            data = self._convert_from_legacy(data)

        # Actual processing
        return data * 2

    def _convert_from_legacy(self, data):
        """
        LAVA FLOW: Convert from legacy format
        Last called: Never (according to logs from 2020)
        Safe to delete: Probably, but who wants to risk it?
        """
        return data
