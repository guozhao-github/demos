package main

import "fmt"

//1
type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

func (s *S) Set(age int) {
	s.Age = age
}

//3
func f(i I) { //interface作用是就体现在函数的参数处
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	s := S{}
	f(&s) //4

	s1 := S{90}

	fmt.Println(s1.Get())
}
