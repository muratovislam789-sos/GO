package main

import (
	"fmt"
)

// ---------- User ----------
type User struct {
	name string
	age  int
}

func NewUser(name string, age int) User {
	return User{name: name, age: age}
}

func (u User) GetName() string { return u.name }
func (u User) GetAge() int     { return u.age }

func (u *User) SetName(name string) {
	if name != "" {
		u.name = name
	}
}

func (u *User) SetAge(age int) {
	if age >= 0 {
		u.age = age
	}
}

func (u User) Introduce() {
	fmt.Printf("Привет! Я %s, мне %d лет.\n", u.name, u.age)
}

// ---------- Course ----------
type Course struct {
	Title         string
	MaxScore      int
	StudentsCount int
	MaxStudents   int
}

func (c *Course) GetInfo() {
	fmt.Printf("Курс: %s | Макс. балл: %d | Студентов: %d/%d\n",
		c.Title, c.MaxScore, c.StudentsCount, c.MaxStudents)
}

func (c *Course) AddStudent() bool {
	if c.StudentsCount < c.MaxStudents {
		c.StudentsCount++
		return true
	}
	fmt.Println("Нельзя добавить больше студентов!")
	return false
}

// ---------- Student ----------
type Student struct {
	User
	Score  int
	Course *Course
}

func (s *Student) Study() {
	if s.Score < s.Course.MaxScore {
		s.Score += 10
		if s.Score > s.Course.MaxScore {
			s.Score = s.Course.MaxScore
		}
	}
}

func (s *Student) GetInfo() {
	fmt.Printf("Студент %s | Балл: %d/%d\n", s.GetName(), s.Score, s.Course.MaxScore)
}

func (s *Student) Act() { s.Study() }

// ---------- Teacher ----------
type Teacher struct {
	User
	Subject string
}

func (t *Teacher) Teach() {
	fmt.Printf("Преподаватель %s ведет предмет: %s\n", t.GetName(), t.Subject)
}

func (t *Teacher) GetInfo() {
	fmt.Printf("Преподаватель %s | Предмет: %s\n", t.GetName(), t.Subject)
}

func (t *Teacher) Act() { t.Teach() }

func (t *Teacher) GradeStudent(s *Student, points int) {
	s.Score += points
	if s.Score > s.Course.MaxScore {
		s.Score = s.Course.MaxScore
	}
	fmt.Printf("%s получил %d баллов.\n", s.GetName(), s.Score)
}

// ---------- Интерфейс ----------
type Participant interface {
	GetInfo()
	Act()
}

// ---------- Функции ----------
func AverageScore(students []*Student) float64 {
	if len(students) == 0 {
		return 0
	}
	sum := 0
	for _, s := range students {
		sum += s.Score
	}
	return float64(sum) / float64(len(students))
}

func BestStudent(students []*Student) *Student {
	if len(students) == 0 {
		return nil
	}
	best := students[0]
	for _, s := range students {
		if s.Score > best.Score {
			best = s
		}
	}
	return best
}

// ---------- main ----------
func main() {
	course := Course{Title: "Go Basics", MaxScore: 100, MaxStudents: 3}

	student1 := &Student{User: NewUser("Али", 20), Course: &course}
	student2 := &Student{User: NewUser("Айжан", 22), Course: &course}
	student3 := &Student{User: NewUser("Марат", 19), Course: &course}
	teacher := &Teacher{User: NewUser("Иван", 35), Subject: "Go Programming"}

	// Добавляем студентов
	if course.AddStudent() {
		fmt.Println("Али добавлен")
	}
	if course.AddStudent() {
		fmt.Println("Айжан добавлена")
	}
	if course.AddStudent() {
		fmt.Println("Марат добавлен")
	}

	participants := []Participant{student1, student2, student3, teacher}

	fmt.Println("\n--- Информация ---")
	course.GetInfo()
	for _, p := range participants {
		p.GetInfo()
		p.Act()
	}

	fmt.Println("\n--- Оценки ---")
	teacher.GradeStudent(student1, 15)
	teacher.GradeStudent(student2, 20)
	teacher.GradeStudent(student3, 5)

	fmt.Printf("\nСредний балл студентов: %.2f\n", AverageScore([]*Student{student1, student2, student3}))

	best := BestStudent([]*Student{student1, student2, student3})
	if best != nil {
		fmt.Printf("Лучший студент: %s с %d баллами\n", best.GetName(), best.Score)
	}
}
