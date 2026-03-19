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
func (u User) Introduce() {
	fmt.Printf("Привет! Я %s, мне %d лет.\n", u.name, u.age)
}

// ---------- Course ----------
type Course struct {
	Title         string
	MaxScore      int
	StudentsCount int
}

func (c *Course) GetInfo() {
	fmt.Printf("Курс: %s | Макс. балл: %d | Студентов: %d\n", c.Title, c.MaxScore, c.StudentsCount)
}

func (c *Course) AddStudent() {
	c.StudentsCount++
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

func (t *Teacher) GradeStudent(s *Student) {
	fmt.Printf("%s получил %d баллов.\n", s.GetName(), s.Score)
}

// ---------- Интерфейс ----------
type Participant interface {
	GetInfo()
	Act()
}

// ---------- main ----------
func main() {
	course := Course{Title: "Go Basics", MaxScore: 100}

	student1 := &Student{User: NewUser("Али", 20), Course: &course}
	student2 := &Student{User: NewUser("Айжан", 22), Course: &course}
	teacher := &Teacher{User: NewUser("Иван", 35), Subject: "Go Programming"}

	course.AddStudent()
	course.AddStudent()

	participants := []Participant{student1, student2, teacher}

	for _, p := range participants {
		p.GetInfo()
		p.Act()
	}

	fmt.Println("\n--- Итог ---")
	student1.GetInfo()
	student2.GetInfo()
	teacher.GradeStudent(student1)
	teacher.GradeStudent(student2)

	avg := (student1.Score + student2.Score) / 2
	fmt.Printf("Средний балл студентов: %d\n", avg)
}
