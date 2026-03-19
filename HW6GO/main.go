package main

import "fmt"

//////////////////// ЗАДАЧА 1 ////////////////////

type Book struct {
	Title       string
	Author      string
	pages       int
	isAvailable bool
}

// Конструктор
func newBook(title, author string, pages int) *Book {
	return &Book{
		Title:       title,
		Author:      author,
		pages:       pages,
		isAvailable: true,
	}
}

func (b *Book) Info() {
	fmt.Println("Название:", b.Title)
	fmt.Println("Автор:", b.Author)
	fmt.Println("Страниц:", b.pages)
	fmt.Println("Доступна:", b.isAvailable)
}

func (b *Book) Borrow() {
	if !b.isAvailable {
		fmt.Println("Книга уже выдана")
		return
	}
	b.isAvailable = false
	fmt.Println("Книга выдана")
}

func (b *Book) ReturnBook() {
	b.isAvailable = true
	fmt.Println("Книга возвращена")
}

func (b *Book) GetPages() int {
	return b.pages
}

func (b *Book) SetPages(p int) {
	if p <= 0 {
		fmt.Println("Ошибка: некорректное количество страниц")
		return
	}
	b.pages = p
}

//////////////////// ЗАДАЧА 2 ////////////////////

type Worker interface {
	Work() string
	GetName() string
}

type Programmer struct {
	Name     string
	Language string
}

func (p Programmer) Work() string {
	return "Программист " + p.Name + " пишет код на " + p.Language
}

func (p Programmer) GetName() string {
	return p.Name
}

type Designer struct {
	Name string
	Tool string
}

func (d Designer) Work() string {
	return "Дизайнер " + d.Name + " делает макет в " + d.Tool
}

func (d Designer) GetName() string {
	return d.Name
}

func ShowWork(w Worker) {
	fmt.Println("Имя:", w.GetName())
	fmt.Println(w.Work())
}

//////////////////// ЗАДАЧА 3 ////////////////////

type Product struct {
	Name     string
	price    float64
	Quantity int
}

// Конструктор
func newProduct(name string, price float64, quantity int) *Product {
	if price < 0 {
		price = 0
	}
	if quantity < 0 {
		quantity = 0
	}
	return &Product{
		Name:     name,
		price:    price,
		Quantity: quantity,
	}
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) SetPrice(newPrice float64) {
	if newPrice < 0 {
		fmt.Println("Ошибка: цена не может быть отрицательной")
		return
	}
	p.price = newPrice
}

func (p *Product) Buy(amount int) {
	if amount <= 0 {
		fmt.Println("Ошибка: некорректное количество")
		return
	}
	if amount > p.Quantity {
		fmt.Println("Недостаточно товара")
		return
	}
	p.Quantity -= amount
	fmt.Println("Покупка:", amount)
}

func (p *Product) Restock(amount int) {
	if amount <= 0 {
		fmt.Println("Ошибка пополнения")
		return
	}
	p.Quantity += amount
}

func (p *Product) Info() {
	fmt.Println("Товар:", p.Name)
	fmt.Println("Цена:", p.price)
	fmt.Println("Количество:", p.Quantity)
}

//////////////////// MAIN ////////////////////

func main() {

	fmt.Println("=== ЗАДАЧА 1: BOOK ===")
	book := newBook("Гарри Поттер", "Роулинг", 500)

	book.Info()
	book.Borrow()
	book.Borrow()
	book.ReturnBook()
	book.SetPages(600)
	book.Info()

	fmt.Println("\n=== ЗАДАЧА 2: WORKERS ===")

	p1 := Programmer{Name: "Али", Language: "Go"}
	p2 := Programmer{Name: "Дима", Language: "Java"}

	d1 := Designer{Name: "Алина", Tool: "Figma"}
	d2 := Designer{Name: "Оля", Tool: "Photoshop"}

	ShowWork(p1)
	ShowWork(p2)
	ShowWork(d1)
	ShowWork(d2)

	fmt.Println("\n=== ЗАДАЧА 3: PRODUCT ===")

	product := newProduct("Телефон", 1000, 5)

	product.Info()
	product.Buy(2)
	product.Buy(10)
	product.SetPrice(1200)
	product.Restock(5)
	product.Info()
}
