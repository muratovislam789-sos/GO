package main

import (
	"fmt"
)

type Course struct {
	Id    int
	Name  string
	Price int
}

type Student struct {
	Name    string
	Courses []Course
}

type School struct {
	catalog  []Course
	students []Student
}

func (s *School) AddCourse(c Course) {
	s.catalog = append(s.catalog, c)
}

func (s *School) AddStudent(name string) {
	s.students = append(s.students, Student{Name: name})
}

func (s *School) Enroll(name string, courseId int) {
	var student *Student
	for i := range s.students {
		if s.students[i].Name == name {
			student = &s.students[i]
			break
		}
	}
	if student == nil {
		return
	}

	for _, c := range s.catalog {
		if c.Id == courseId {
			student.Courses = append(student.Courses, c)
			return
		}
	}
}

func (s *School) TotalCost(name string) int {
	for _, stu := range s.students {
		if stu.Name == name {
			total := 0
			for _, c := range stu.Courses {
				total += c.Price
			}
			return total
		}
	}
	return 0
}

func (s *School) PrintStudent(name string) {
	for _, stu := range s.students {
		if stu.Name == name {
			fmt.Printf("Student: %s\n", stu.Name)
			fmt.Print("Courses: ")
			for i, c := range stu.Courses {
				if i > 0 {
					fmt.Print(", ")
				}
				fmt.Printf("%s - %d", c.Name, c.Price)
			}
			fmt.Printf("\nTotal: %d\n", s.TotalCost(name))
			return
		}
	}
}

func main() {
	school := School{}

	school.AddCourse(Course{Id: 1, Name: "Go Basics", Price: 20000})
	school.AddCourse(Course{Id: 2, Name: "SQL", Price: 15000})
	school.AddCourse(Course{Id: 3, Name: "HTML & CSS", Price: 12000})

	school.AddStudent("Ali")
	school.AddStudent("Mansur")

	school.Enroll("Ali", 1)
	school.Enroll("Ali", 2)
	school.Enroll("Mansur", 2)
	school.Enroll("Mansur", 3)

	school.PrintStudent("Ali")
	school.PrintStudent("Mansur")
}