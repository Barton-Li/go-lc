package main

import "fmt"

type Singleton struct {
	Name string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{Name: "Singleton"}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	fmt.Println(s1.Name)
	s2 := GetInstance()
	fmt.Println(s2.Name)
	if s1 == s2 {
		fmt.Println("s1 == s2")
	}else{
		fmt.Println("s1 != s2")
	}
}