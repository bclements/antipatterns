"""
ANTI-PATTERN: Spaghetti Code

Code with complex and tangled control flow that's difficult to understand and maintain.
The logic jumps around making it hard to follow the execution path.
"""

def process_order(order_id, user_type, payment_method, items, discount_code, shipping_method):
    """A nightmare of nested conditions and tangled logic"""
    status = None
    total = 0

    if user_type == "premium":
        if payment_method == "credit":
            if len(items) > 5:
                total = sum([item['price'] for item in items])
                if discount_code:
                    if discount_code == "SAVE20":
                        total = total * 0.8
                        status = "processing"
                    elif discount_code == "SAVE10":
                        total = total * 0.9
                        if shipping_method == "express":
                            total += 20
                            status = "processing"
                        else:
                            total += 5
                            status = "pending"
                    else:
                        status = "invalid_code"
                        return None
                else:
                    if shipping_method == "express":
                        total += 20
                        status = "processing"
                    else:
                        status = "processing"
            else:
                if discount_code:
                    total = sum([item['price'] for item in items])
                    if discount_code == "SAVE20":
                        total = total * 0.8
                        status = "processing"
                    else:
                        status = "invalid_code"
                else:
                    total = sum([item['price'] for item in items])
                    status = "processing"
        elif payment_method == "paypal":
            total = sum([item['price'] for item in items])
            if total > 100:
                if discount_code == "SAVE20":
                    total = total * 0.8
                    status = "processing"
                else:
                    status = "processing"
            else:
                status = "pending"
        else:
            status = "invalid_payment"
            return None
    elif user_type == "regular":
        if payment_method == "credit":
            total = sum([item['price'] for item in items])
            if discount_code:
                if discount_code == "SAVE10":
                    total = total * 0.9
                    if len(items) > 3:
                        status = "processing"
                    else:
                        status = "pending"
                else:
                    status = "invalid_code"
            else:
                if len(items) > 3:
                    status = "processing"
                else:
                    status = "pending"
        else:
            status = "cash_only_for_regular"
    else:
        if payment_method == "credit":
            total = sum([item['price'] for item in items])
            status = "guest_order"
        else:
            status = "invalid"

    # More tangled logic for shipping
    if status == "processing":
        if shipping_method == "express":
            if user_type == "premium":
                # Free express for premium
                pass
            else:
                total += 20
        elif shipping_method == "standard":
            if total > 50:
                # Free shipping
                pass
            else:
                total += 5

    # Even more tangled validation
    if status:
        if status != "invalid":
            if total > 0:
                if user_type == "premium" or user_type == "regular":
                    return {"order_id": order_id, "total": total, "status": status}

    return None


# Another example: tangled state management
class OrderProcessor:
    def __init__(self):
        self.state = "init"
        self.temp_value = None
        self.flag1 = False
        self.flag2 = False
        self.counter = 0

    def process(self, data):
        if self.state == "init":
            self.temp_value = data
            self.state = "loading"
            self.flag1 = True
            self.counter += 1
            if self.counter > 5:
                self.flag2 = True
                self.state = "ready"
                return self._handle_ready()
            else:
                return self._handle_loading()
        elif self.state == "loading":
            if self.flag1:
                self.temp_value += data
                if self.flag2:
                    self.state = "processing"
                    return self._handle_processing()
                else:
                    self.counter += 1
                    if self.counter > 10:
                        self.state = "error"
                        return None
                    return self._handle_loading()
        # ... and so on with more tangled state transitions

    def _handle_ready(self):
        if self.flag1 and self.flag2:
            self.state = "done"
            return self.temp_value
        elif self.flag1:
            self.state = "loading"
            return None

    def _handle_loading(self):
        if self.counter > 7:
            self.flag2 = True
            self.state = "ready"
            return self._handle_ready()
        return None

    def _handle_processing(self):
        # More spaghetti...
        pass
