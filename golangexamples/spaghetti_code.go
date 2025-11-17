package main

/*
ANTI-PATTERN: Spaghetti Code

Code with complex and tangled control flow that's difficult to understand and maintain.
The logic jumps around making it hard to follow the execution path.
*/

import (
	"fmt"
)

// Order represents an order in the system
type Order struct {
	ID             int
	Items          []Item
	UserType       string
	PaymentMethod  string
	DiscountCode   string
	ShippingMethod string
	Status         string
}

type Item struct {
	Product string
	Price   float64
}

// ProcessOrder - A nightmare of nested conditions and tangled logic
func ProcessOrder(orderID int, userType, paymentMethod, discountCode, shippingMethod string, items []Item) *Order {
	status := ""
	total := 0.0

	if userType == "premium" {
		if paymentMethod == "credit" {
			if len(items) > 5 {
				for _, item := range items {
					total += item.Price
				}
				if discountCode != "" {
					if discountCode == "SAVE20" {
						total = total * 0.8
						status = "processing"
					} else if discountCode == "SAVE10" {
						total = total * 0.9
						if shippingMethod == "express" {
							total += 20
							status = "processing"
						} else {
							total += 5
							status = "pending"
						}
					} else {
						status = "invalid_code"
						return nil
					}
				} else {
					if shippingMethod == "express" {
						total += 20
						status = "processing"
					} else {
						status = "processing"
					}
				}
			} else {
				if discountCode != "" {
					for _, item := range items {
						total += item.Price
					}
					if discountCode == "SAVE20" {
						total = total * 0.8
						status = "processing"
					} else {
						status = "invalid_code"
					}
				} else {
					for _, item := range items {
						total += item.Price
					}
					status = "processing"
				}
			}
		} else if paymentMethod == "paypal" {
			for _, item := range items {
				total += item.Price
			}
			if total > 100 {
				if discountCode == "SAVE20" {
					total = total * 0.8
					status = "processing"
				} else {
					status = "processing"
				}
			} else {
				status = "pending"
			}
		} else {
			status = "invalid_payment"
			return nil
		}
	} else if userType == "regular" {
		if paymentMethod == "credit" {
			for _, item := range items {
				total += item.Price
			}
			if discountCode != "" {
				if discountCode == "SAVE10" {
					total = total * 0.9
					if len(items) > 3 {
						status = "processing"
					} else {
						status = "pending"
					}
				} else {
					status = "invalid_code"
				}
			} else {
				if len(items) > 3 {
					status = "processing"
				} else {
					status = "pending"
				}
			}
		} else {
			status = "cash_only_for_regular"
		}
	} else {
		if paymentMethod == "credit" {
			for _, item := range items {
				total += item.Price
			}
			status = "guest_order"
		} else {
			status = "invalid"
		}
	}

	// More tangled logic for shipping
	if status == "processing" {
		if shippingMethod == "express" {
			if userType == "premium" {
				// Free express for premium
			} else {
				total += 20
			}
		} else if shippingMethod == "standard" {
			if total > 50 {
				// Free shipping
			} else {
				total += 5
			}
		}
	}

	// Even more tangled validation
	if status != "" {
		if status != "invalid" {
			if total > 0 {
				if userType == "premium" || userType == "regular" {
					return &Order{
						ID:             orderID,
						Items:          items,
						UserType:       userType,
						PaymentMethod:  paymentMethod,
						DiscountCode:   discountCode,
						ShippingMethod: shippingMethod,
						Status:         status,
					}
				}
			}
		}
	}

	return nil
}

// OrderProcessor - Another example: tangled state management
type OrderProcessor struct {
	state      string
	tempValue  interface{}
	flag1      bool
	flag2      bool
	counter    int
	retryCount int
}

func NewOrderProcessor() *OrderProcessor {
	return &OrderProcessor{
		state:   "init",
		counter: 0,
	}
}

func (op *OrderProcessor) Process(data interface{}) interface{} {
	if op.state == "init" {
		op.tempValue = data
		op.state = "loading"
		op.flag1 = true
		op.counter++
		if op.counter > 5 {
			op.flag2 = true
			op.state = "ready"
			return op.handleReady()
		}
		return op.handleLoading()
	} else if op.state == "loading" {
		if op.flag1 {
			op.tempValue = fmt.Sprintf("%v%v", op.tempValue, data)
			if op.flag2 {
				op.state = "processing"
				return op.handleProcessing()
			}
			op.counter++
			if op.counter > 10 {
				op.state = "error"
				return nil
			}
			return op.handleLoading()
		}
	} else if op.state == "processing" {
		if op.flag1 && op.flag2 {
			if op.retryCount < 3 {
				op.retryCount++
				op.state = "retrying"
				return op.handleRetrying()
			}
			op.state = "complete"
			return op.handleComplete()
		}
	} else if op.state == "retrying" {
		if op.retryCount >= 3 {
			op.state = "failed"
			return nil
		}
		op.state = "processing"
		return op.handleProcessing()
	}
	return nil
}

func (op *OrderProcessor) handleReady() interface{} {
	if op.flag1 && op.flag2 {
		op.state = "done"
		return op.tempValue
	} else if op.flag1 {
		op.state = "loading"
		return nil
	}
	return nil
}

func (op *OrderProcessor) handleLoading() interface{} {
	if op.counter > 7 {
		op.flag2 = true
		op.state = "ready"
		return op.handleReady()
	}
	return nil
}

func (op *OrderProcessor) handleProcessing() interface{} {
	// More spaghetti logic
	if op.tempValue != nil {
		if op.counter > 5 {
			if op.flag1 {
				return op.tempValue
			}
		}
	}
	return nil
}

func (op *OrderProcessor) handleRetrying() interface{} {
	// Yet more tangled logic
	if op.retryCount > 0 {
		if op.retryCount < 3 {
			return op.handleProcessing()
		}
	}
	return nil
}

func (op *OrderProcessor) handleComplete() interface{} {
	return op.tempValue
}

// ValidateUser - More spaghetti with deeply nested conditions
func ValidateUser(username, password, email, userType string, age int, country string) (bool, string) {
	if username != "" {
		if len(username) >= 3 {
			if len(username) <= 20 {
				if password != "" {
					if len(password) >= 8 {
						if email != "" {
							if userType == "admin" {
								if age >= 21 {
									if country == "US" || country == "UK" {
										return true, "valid"
									}
									return false, "admin must be in US or UK"
								}
								return false, "admin must be 21+"
							} else if userType == "regular" {
								if age >= 18 {
									if country != "" {
										return true, "valid"
									}
									return false, "country required"
								}
								return false, "must be 18+"
							} else {
								if age >= 13 {
									return true, "valid"
								}
								return false, "must be 13+"
							}
						}
						return false, "email required"
					}
					return false, "password too short"
				}
				return false, "password required"
			}
			return false, "username too long"
		}
		return false, "username too short"
	}
	return false, "username required"
}

func main() {
	// Example usage
	items := []Item{
		{Product: "Item1", Price: 10.0},
		{Product: "Item2", Price: 20.0},
	}

	order := ProcessOrder(1, "premium", "credit", "SAVE20", "express", items)
	if order != nil {
		fmt.Printf("Order processed: %+v\n", order)
	}

	// Trying to understand what this code does is like untangling spaghetti
}
