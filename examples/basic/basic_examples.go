package main

import (
	"fmt"
	"strconv"
	"time"
)

const my_const string = "constant"

func main() {

	fmt.Println("Hello world")
	fmt.Println("go" + "lang")
	fmt.Println("1+1", 1+1)

	fmt.Println("7.0/3.0", 7.0/3.0)

	var i int
	g := "d"

	var a, b, c = 1, 2, 3
	d, e, f := 3, 4, 5
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
	fmt.Println(i)
	fmt.Println(f)
	fmt.Printf("%T\n", g)

	fmt.Printf("%T\n", my_const)

	const n = 7.2
	fmt.Print(n)
	fmt.Printf("%T\n", n)

	j := 1
	for j < 10 {
		println(j)
		j++
	}

	for j = 10; j < 20; j++ {
		println(j)
	}

	var s int = 1
	switch s {
	case 1:
		println(1)
	case 2:
		println(2)
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("Its weekend")
	default:
		println("Its not weekend :( ")
	}

	t := time.Now()

	switch {
	case t.Hour() > 5:
		println("Nothing else matters ")
	case s >= 1:
		println("S is the way to go")
	}

	whatAmIFunc := func(i interface{}) {
		switch t := i.(type) { // i.(type) type assertions
		case bool:
			println("It is boolean")
		case int:
			println("It is int")
		case float64:
			println("It is float64")
		default:
			fmt.Printf("Not known type %T\n", t)
		}
	}

	whatAmIFunc(true)
	whatAmIFunc(1)
	whatAmIFunc("string")

	fmt.Printf("Function type %T\n", whatAmIFunc)

	//arrays fixed length

	tab_a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("tab_a:", tab_a)

	var tab_two [2][3]int
	fmt.Println("tab_two", tab_two)

	// slices like lists

	slice1 := make([]string, 3)
	slice1[0] = "a"
	slice1[2] = "d"
	slice1[1] = "f"
	slice1 = append(slice1, "ww")

	fmt.Println(slice1)
	slice2 := make([]string, len(slice1))
	copy(slice2, slice1)

	println(slice1)
	println(slice2)

	fmt.Println(slice1[:5])
	fmt.Println(slice1[1:2])

	// maps
	map1 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3
	fmt.Println(map1)
	delete(map1, "d")
	fmt.Println(map1)
	map2 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(map2)

	// ranges (iterate over things with index)

	for k, v := range map1 {
		println(k, v)
	}
	for k := range map1 {
		println(k)
	}
	for index, _ := range slice1 {
		println(index)
	}

	// funcitons

	val1, val2 := vals()
	val3, _ := vals()
	println(val1, val2, val3)

	var nextInt = intSeq()
	println(nextInt())
	println(nextInt())
	println(nextInt())

	var nextInt1 = intSeq()
	println(nextInt1())

	// pointers
	myValueToChange := 5
	println(myValueToChange)
	zeroValueNormal(myValueToChange)
	println(myValueToChange)
	zeroValuePtr(&myValueToChange)
	println(myValueToChange)

	//structs
	fred := person{name: "Fred"}

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Fred"})

	// methods - receivers
	fmt.Println(fred.introduce())

	// interfaces
	// musi być &fred a nie fred z powodu (p *person) introduce()
	var introducable introduceYourself = &fred
	fmt.Println(introducable.introduce())

}

type introduceYourself interface {
	introduce() string
}

type person struct {
	name string
	age  int
}

func aaa(p *person) {
	println(p.age)
}

func (p *person) introduce() string {
	return "I am " + p.name + " and I have " + strconv.Itoa(p.age) + " years"
}

// interfaces
//var introducable introduceYourself = fred
// fmt.Println(introducable.introduce())

//func (p person) introduce() string {
//	return "I am " + p.name + " and I have " + strconv.Itoa(p.age) + " years"
//}

func zeroValueNormal(i int) {
	i = 0
}

func zeroValuePtr(i *int) {
	*i = 0
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func vals() (int, int) {
	return 3, 7
}
