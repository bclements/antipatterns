"""
ANTI-PATTERN: Copy and Paste Programming

Duplicating code instead of creating reusable functions or classes.
This leads to maintenance nightmares when bugs need to be fixed in multiple places.
"""

class ReportGenerator:
    """Code duplication everywhere instead of creating reusable methods"""

    def generate_sales_report(self, sales_data):
        print("=" * 50)
        print("SALES REPORT")
        print("=" * 50)
        print(f"Generated: 2024-01-15")
        print(f"Total Records: {len(sales_data)}")
        print("-" * 50)

        total = 0
        for item in sales_data:
            print(f"{item['product']}: ${item['amount']}")
            total += item['amount']

        print("-" * 50)
        print(f"Total: ${total}")
        print("=" * 50)

    def generate_expense_report(self, expense_data):
        # COPY-PASTE: Almost identical to sales_report!
        print("=" * 50)
        print("EXPENSE REPORT")
        print("=" * 50)
        print(f"Generated: 2024-01-15")
        print(f"Total Records: {len(expense_data)}")
        print("-" * 50)

        total = 0
        for item in expense_data:
            print(f"{item['product']}: ${item['amount']}")
            total += item['amount']

        print("-" * 50)
        print(f"Total: ${total}")
        print("=" * 50)

    def generate_inventory_report(self, inventory_data):
        # COPY-PASTE: Again, almost the same code!
        print("=" * 50)
        print("INVENTORY REPORT")
        print("=" * 50)
        print(f"Generated: 2024-01-15")
        print(f"Total Records: {len(inventory_data)}")
        print("-" * 50)

        total = 0
        for item in inventory_data:
            print(f"{item['product']}: ${item['amount']}")
            total += item['amount']

        print("-" * 50)
        print(f"Total: ${total}")
        print("=" * 50)


class UserValidator:
    """More copy-paste nightmares"""

    def validate_admin_user(self, username, password, email):
        # COPY-PASTE: Validation logic duplicated
        if not username:
            return False, "Username is required"
        if len(username) < 3:
            return False, "Username too short"
        if not password:
            return False, "Password is required"
        if len(password) < 8:
            return False, "Password too short"
        if not email:
            return False, "Email is required"
        if "@" not in email:
            return False, "Invalid email"
        # Admin-specific check
        if not username.startswith("admin_"):
            return False, "Admin username must start with admin_"
        return True, "Valid"

    def validate_regular_user(self, username, password, email):
        # COPY-PASTE: Same validation code with tiny difference
        if not username:
            return False, "Username is required"
        if len(username) < 3:
            return False, "Username too short"
        if not password:
            return False, "Password is required"
        if len(password) < 8:
            return False, "Password too short"
        if not email:
            return False, "Email is required"
        if "@" not in email:
            return False, "Invalid email"
        # Regular user specific check
        if username.startswith("admin_"):
            return False, "Regular users cannot have admin_ prefix"
        return True, "Valid"

    def validate_guest_user(self, username, password, email):
        # COPY-PASTE: Yet again the same validation!
        if not username:
            return False, "Username is required"
        if len(username) < 3:
            return False, "Username too short"
        if not password:
            return False, "Password is required"
        if len(password) < 8:
            return False, "Password too short"
        if not email:
            return False, "Email is required"
        if "@" not in email:
            return False, "Invalid email"
        # Guest specific check
        if not username.startswith("guest_"):
            return False, "Guest username must start with guest_"
        return True, "Valid"


# COPY-PASTE: Database query duplication
def get_user_by_id(user_id):
    connection = create_connection()
    cursor = connection.cursor()
    query = "SELECT * FROM users WHERE id = ?"
    cursor.execute(query, (user_id,))
    result = cursor.fetchone()
    cursor.close()
    connection.close()
    return result

def get_product_by_id(product_id):
    # COPY-PASTE: Exact same pattern as above
    connection = create_connection()
    cursor = connection.cursor()
    query = "SELECT * FROM products WHERE id = ?"
    cursor.execute(query, (product_id,))
    result = cursor.fetchone()
    cursor.close()
    connection.close()
    return result

def get_order_by_id(order_id):
    # COPY-PASTE: Again the same pattern
    connection = create_connection()
    cursor = connection.cursor()
    query = "SELECT * FROM orders WHERE id = ?"
    cursor.execute(query, (order_id,))
    result = cursor.fetchone()
    cursor.close()
    connection.close()
    return result

def create_connection():
    """Dummy function"""
    pass


# COPY-PASTE: API endpoint duplication
class APIHandler:
    def handle_get_user(self, request):
        # COPY-PASTE: Error handling duplicated everywhere
        try:
            data = request.get_data()
            if not data:
                return {"error": "No data provided"}, 400
            # Process user...
            return {"success": True}, 200
        except ValueError as e:
            return {"error": str(e)}, 400
        except Exception as e:
            return {"error": "Internal server error"}, 500

    def handle_get_product(self, request):
        # COPY-PASTE: Same error handling pattern
        try:
            data = request.get_data()
            if not data:
                return {"error": "No data provided"}, 400
            # Process product...
            return {"success": True}, 200
        except ValueError as e:
            return {"error": str(e)}, 400
        except Exception as e:
            return {"error": "Internal server error"}, 500

    def handle_get_order(self, request):
        # COPY-PASTE: And again...
        try:
            data = request.get_data()
            if not data:
                return {"error": "No data provided"}, 400
            # Process order...
            return {"success": True}, 200
        except ValueError as e:
            return {"error": str(e)}, 400
        except Exception as e:
            return {"error": "Internal server error"}, 500
