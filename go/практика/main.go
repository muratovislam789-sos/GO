// ПРИАКТИКА 1
// package main

// import "fmt"

// // Задание 1
// type Product struct {
// 	Name  string
// 	Price float64
// 	Stock int
// }

// // Задание 3
// func IncreasePrice(product *Product) {
// 	product.Price = product.Price * 1.10
// }

// func main() {

// 	products := []Product{
// 		{Name: "Laptop", Price: 500000, Stock: 5},
// 		{Name: "Phone", Price: 300000, Stock: 10},
// 		{Name: "Headphones", Price: 50000, Stock: 20},
// 	}

// 	fmt.Println("Список товаров:")
// 	for _, p := range products {
// 		fmt.Println(p.Name, p.Price, p.Stock)
// 	}

// 	// Задание 2
// 	expensive := products[0]
// 	for _, p := range products {
// 		if p.Price > expensive.Price {
// 			expensive = p
// 		}
// 	}

// 	fmt.Println("Самый дорогой товар:", expensive.Name)

// 	IncreasePrice(&products[0])
// 	fmt.Println("Новая цена после увеличения:", products[0].Price)
// }

// // ПРИАКТИКА 2
// package main

// import "fmt"

// // ================= INTERFACE =================

// type Payment interface {
// 	Pay(amount float64) bool
// }

// // ================= TRANSACTION =================

// type Transaction struct {
// 	Method string
// 	Amount float64
// 	Fee    float64
// 	Total  float64
// 	Ok     bool
// }

// // ================= CARD PAYMENT =================

// type CardPayment struct {
// 	CardNumber  string
// 	Balance     float64
// 	FeePercent  float64
// 	Limit       float64
// 	IsBlocked   bool
// 	FailedCount int
// 	History     []Transaction
// }

// func (c *CardPayment) Pay(amount float64) bool {
// 	fee := amount * c.FeePercent / 100
// 	total := amount + fee
// 	ok := false

// 	if c.IsBlocked {
// 		fmt.Println("Карта заблокирована")
// 	} else if amount <= 0 {
// 		fmt.Println("Некорректная сумма")
// 	} else if amount > c.Limit {
// 		fmt.Println("Превышен лимит")
// 	} else if total > c.Balance {
// 		fmt.Println("Недостаточно средств на карте")
// 	} else {
// 		c.Balance -= total
// 		fmt.Println("Оплата картой:", amount)
// 		ok = true
// 	}

// 	c.History = append(c.History, Transaction{
// 		Method: "card",
// 		Amount: amount,
// 		Fee:    fee,
// 		Total:  total,
// 		Ok:     ok,
// 	})

// 	return ok
// }

// // ================= CASH PAYMENT =================

// type CashPayment struct {
// 	Money      float64
// 	FeePercent float64
// 	Limit      float64
// 	History    []Transaction
// }

// func (c *CashPayment) Pay(amount float64) bool {
// 	fee := amount * c.FeePercent / 100
// 	total := amount + fee
// 	ok := false

// 	if amount <= 0 {
// 		fmt.Println("Некорректная сумма")
// 	} else if amount > c.Limit {
// 		fmt.Println("Превышен лимит")
// 	} else if total > c.Money {
// 		fmt.Println("Недостаточно наличных")
// 	} else {
// 		c.Money -= total
// 		fmt.Println("Оплата наличными:", amount)
// 		ok = true
// 	}

// 	c.History = append(c.History, Transaction{
// 		Method: "cash",
// 		Amount: amount,
// 		Fee:    fee,
// 		Total:  total,
// 		Ok:     ok,
// 	})

// 	return ok
// }

// // ================= CRYPTO PAYMENT =================

// type CryptoPayment struct {
// 	Wallet     string
// 	Coins      float64
// 	FeePercent float64
// 	Limit      float64
// 	History    []Transaction
// }

// func (c *CryptoPayment) Pay(amount float64) bool {
// 	fee := amount * c.FeePercent / 100
// 	total := amount + fee
// 	ok := false

// 	if amount <= 0 {
// 		fmt.Println("Некорректная сумма")
// 	} else if amount > c.Limit {
// 		fmt.Println("Превышен лимит")
// 	} else if total > c.Coins {
// 		fmt.Println("Недостаточно крипты")
// 	} else {
// 		c.Coins -= total
// 		fmt.Println("Оплата криптой:", amount)
// 		ok = true
// 	}

// 	c.History = append(c.History, Transaction{
// 		Method: "crypto",
// 		Amount: amount,
// 		Fee:    fee,
// 		Total:  total,
// 		Ok:     ok,
// 	})

// 	return ok
// }

// // ================= PROCESS PAYMENT =================

// func ProcessPayment(p Payment, amount float64) {
// 	fmt.Println("Начинаем оплату")

// 	ok := p.Pay(amount)

// 	if card, okType := p.(*CardPayment); okType {
// 		if !ok {
// 			card.FailedCount++
// 			if card.FailedCount >= 3 {
// 				card.IsBlocked = true
// 				fmt.Println("Карта заблокирована после 3 неудачных попыток")
// 			}
// 		} else {
// 			card.FailedCount = 0
// 		}
// 	}

// 	if ok {
// 		fmt.Println("Оплата прошла успешно")
// 	} else {
// 		fmt.Println("Оплата не прошла")
// 	}

// 	fmt.Println()
// }

// // ================= MAIN TESTS =================

// func main() {
// 	fmt.Println("НОВАЯ ВЕРСИЯ КОДА")

// 	card := &CardPayment{
// 		CardNumber: "1234-5678",
// 		Balance:    1000,
// 		FeePercent: 2,
// 		Limit:      500,
// 	}

// 	cash := &CashPayment{
// 		Money:      500,
// 		FeePercent: 1,
// 		Limit:      300,
// 	}

// 	crypto := &CryptoPayment{
// 		Wallet:     "0xABC",
// 		Coins:      5,
// 		FeePercent: 3,
// 		Limit:      2,
// 	}

// 	// ❌ amount = -10
// 	ProcessPayment(card, -10)

// 	// ❌ больше баланса
// 	ProcessPayment(card, 2000)

// 	// ❌ больше лимита
// 	ProcessPayment(card, 600)

// 	// ✅ успешные оплаты
// 	ProcessPayment(card, 100)
// 	ProcessPayment(card, 100)
// 	ProcessPayment(card, 100)

// 	// ❌ 3 ошибки подряд → блокировка
// 	ProcessPayment(card, 2000)
// 	ProcessPayment(card, 2000)
// 	ProcessPayment(card, 2000)

// 	// Проверка других методов
// 	ProcessPayment(cash, 200)
// 	ProcessPayment(crypto, 1)
// }

// // ПРАКТИКА 3
// package main

// import "fmt"

// // ================= INTERFACES =================

// type Payment interface {
// 	Pay(amount float64) bool
// }

// type Refundable interface {
// 	Refund(amount float64) bool
// }

// // ================= CARD PAYMENT =================

// type CardPayment struct {
// 	CardNumber string
// 	Balance    float64
// }

// func (c *CardPayment) Pay(amount float64) bool {
// 	if amount <= 0 || amount > c.Balance {
// 		fmt.Println("Недостаточно средств или некорректная сумма")
// 		return false
// 	}
// 	c.Balance -= amount
// 	fmt.Println("Оплата картой:", amount)
// 	return true
// }

// func (c *CardPayment) Refund(amount float64) bool {
// 	if amount <= 0 {
// 		return false
// 	}
// 	c.Balance += amount
// 	fmt.Println("Возврат картой:", amount)
// 	return true
// }

// // ================= CASH PAYMENT =================

// type CashPayment struct {
// 	Money float64
// }

// func (c *CashPayment) Pay(amount float64) bool {
// 	if amount <= 0 || amount > c.Money {
// 		fmt.Println("Недостаточно наличных или некорректная сумма")
// 		return false
// 	}
// 	c.Money -= amount
// 	fmt.Println("Оплата наличными:", amount)
// 	return true
// }

// func (c *CashPayment) Refund(amount float64) bool {
// 	if amount <= 0 {
// 		return false
// 	}
// 	c.Money += amount
// 	fmt.Println("Возврат наличными:", amount)
// 	return true
// }

// // ================= CRYPTO PAYMENT =================

// type CryptoPayment struct {
// 	Coins float64
// }

// func (c *CryptoPayment) Pay(amount float64) bool {
// 	if amount <= 0 || amount > c.Coins {
// 		fmt.Println("Недостаточно крипты или некорректная сумма")
// 		return false
// 	}
// 	c.Coins -= amount
// 	fmt.Println("Оплата криптой:", amount)
// 	return true
// }

// func (c *CryptoPayment) Refund(amount float64) bool {
// 	if amount <= 0 {
// 		return false
// 	}
// 	c.Coins += amount
// 	fmt.Println("Возврат криптой:", amount)
// 	return true
// }

// // ================= CASHIER =================

// type Op struct {
// 	Kind   string
// 	Method string
// 	Amount float64
// 	Ok     bool
// }

// type Cashier struct {
// 	DayTotal float64
// 	Ops      []Op
// }

// func (c *Cashier) Checkout(p Payment, method string, amount float64) {
// 	ok := p.Pay(amount)
// 	if ok {
// 		c.DayTotal += amount
// 	}
// 	c.Ops = append(c.Ops, Op{
// 		Kind:   "pay",
// 		Method: method,
// 		Amount: amount,
// 		Ok:     ok,
// 	})
// 	fmt.Printf("Чек: %s %.2f статус: %v\n", method, amount, ok)
// }

// func (c *Cashier) MakeRefund(r Refundable, method string, amount float64) {
// 	ok := r.Refund(amount)
// 	if ok {
// 		c.DayTotal -= amount
// 	}
// 	c.Ops = append(c.Ops, Op{
// 		Kind:   "refund",
// 		Method: method,
// 		Amount: amount,
// 		Ok:     ok,
// 	})
// 	fmt.Printf("Возврат: %s %.2f статус: %v\n", method, amount, ok)
// }

// // ================= LOGGED PAYMENT =================

// type LoggedPayment struct {
// 	Method string
// 	Inner  Payment
// 	Log    []string
// }

// func (l *LoggedPayment) Pay(amount float64) bool {
// 	l.Log = append(l.Log, fmt.Sprintf("До оплаты %s: %.2f", l.Method, amount))
// 	ok := l.Inner.Pay(amount)
// 	l.Log = append(l.Log, fmt.Sprintf("После оплаты %s: %v", l.Method, ok))
// 	return ok
// }

// // ================= MAIN =================

// func main() {
// 	fmt.Println("=== Мини-касса с refund ===")

// 	card := &CardPayment{CardNumber: "1234-5678", Balance: 1000}
// 	cash := &CashPayment{Money: 500}
// 	crypto := &CryptoPayment{Coins: 5}

// 	cashier := &Cashier{}

// 	// Checkout оплата
// 	cashier.Checkout(card, "card", 100)
// 	cashier.Checkout(cash, "cash", 200)
// 	cashier.Checkout(crypto, "crypto", 1)

// 	// Попытка возврата
// 	payments := []Payment{card, cash, crypto}
// 	amountRefund := 50.0
// 	for _, p := range payments {
// 		if r, ok := p.(Refundable); ok {
// 			cashier.MakeRefund(r, fmt.Sprintf("%T", p), amountRefund)
// 		} else {
// 			fmt.Println("Этот способ оплаты не поддерживает возврат")
// 		}
// 	}

// 	// LoggedPayment для карты
// 	loggedCard := &LoggedPayment{Method: "card", Inner: card}
// 	cashier.Checkout(loggedCard, "card", 150)

// 	fmt.Println("=== История операций ===")
// 	for _, op := range cashier.Ops {
// 		fmt.Printf("%s %s %.2f статус: %v\n", op.Kind, op.Method, op.Amount, op.Ok)
// 	}

// 	fmt.Println("=== Лог карты ===")
// 	for _, l := range loggedCard.Log {
// 		fmt.Println(l)
// 	}

// 	fmt.Printf("Итоговый DayTotal: %.2f\n", cashier.DayTotal)
// }

// ДЗ Move Interface Transport
package main

import "fmt"

// ================= INTERFACE =================

type Move interface {
	Move() string
}

// ================= TRANSPORT STRUCTS =================

type Car struct {
	Brand string
	Speed int
}

type Bike struct {
	Brand string
	Speed int
}

type Plane struct {
	Model string
	Speed int
}

// ================= MOVE METHODS =================

func (c Car) Move() string {
	return fmt.Sprintf("Машина %s едет со скоростью %d км/ч", c.Brand, c.Speed)
}

func (b Bike) Move() string {
	return fmt.Sprintf("Велосипед %s едет со скоростью %d км/ч", b.Brand, b.Speed)
}

func (p Plane) Move() string {
	return fmt.Sprintf("Самолет %s летит со скоростью %d км/ч", p.Model, p.Speed)
}

// ================= START JOURNEY =================

func StartJourney(m Move) {
	fmt.Println(m.Move())
}

// ================= MAIN =================

func main() {
	car := Car{"Toyota", 120}
	bike := Bike{"Giant", 25}
	plane := Plane{"Boeing 747", 900}

	transports := []Move{car, bike, plane}

	fmt.Println("=== Путешествие начинается ===")
	for _, t := range transports {
		StartJourney(t)
	}
}
