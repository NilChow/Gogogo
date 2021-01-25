package main

import (
	"fmt"
)

func main() {
	fmt.Println("Interface Test Start...")

	InterfaceInSlice()
}

type Person interface {
	Name() string
	Age() int
	Sex() bool
}

type Student struct {
	name	string
	age		int
	sex		bool
	grade	string
}

type Singer struct {
	name	string
	age		int
	sex		bool
	song	string
}

func (s *Student) Name() string {
	return s.name
}
func (s *Student) Age() int {
	return s.age
}
func (s *Student) Sex() bool {
	return s.sex
}

func (s *Singer) Name() string {
	return s.name
}
func (s *Singer) Age() int {
	return s.age
}
func (s *Singer) Sex() bool {
	return s.sex
}

// 使用非空接口作为切片元素
func InterfaceInSlice() {
	s := make(map[string][]Person)
	student_1 := &Student{"Iverson", 21, true, "1"}
	student_2 := &Student{"Jordan", 33, true, "12"}
	singer_1 := &Singer{"Ansr J", 28, true, "FREESTYLE"}
	singer_2 := &Singer{"PG-ONE", 23, true, "Fresh"}
	singer_3 := &Singer{"Nineone", 24, false, "PUMA"}

	var temp []Person
	temp = append(temp, student_1)
	temp = append(temp, student_2)
	temp = append(temp, singer_1)
	temp = append(temp, singer_2)
	s["1"] = temp
	temp = append(temp, singer_3)	// 不会被添加到切片中

	for _, v := range s["1"] {
		fmt.Printf("[Name: %s] [Age: %d] [Sex: %v]\n", v.Name(), v.Age(), v.Sex())
	}
}