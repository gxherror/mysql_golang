package main

import (
	//. "fmt"
	//f  "fmt"
	//_ "fmt"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	//"time"
)

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func example(a, b int) (c, d int) {
	return c, d
}

//arg ->int slice
func tot_println(arg ...int) {
	for i, v := range arg {
		fmt.Println(i, v)
	}
}

//send a copy
func add1(a int) int {
	a += 1
	return a
}

//! channel , slice , map like point
//! can be send directly without *
func add2(a *int, f t_add1) int {
	fmt.Println(f(*a))
	*a += 1
	return *a
}

//run defer reverse order when return
func defer_test(n int) {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func throwpanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}() //!
	f()
	return b
}

//!
type t_add1 func(int) int

type skills []string //slice as dynamic array

type Human struct {
	name  string
	age   int
	phone int
}

type Student struct {
	Human
	skills
	score int
	phone int
}

//method use copy//! Capture
func (s Student) get_name() string {
	return s.name
}

//method use ptr
//auto detect ptr
func (s *Student) set_score(score int) {
	s.score = score
}

//
func (s Human) get_age() int {
	return s.age
}

func (s Student) get_age() int {
	return s.age
}

type Student_i interface {
	get_age() int
	get_name() string
}

func (s Student) String() string {
	return "<" + s.name + "-" + strconv.Itoa(s.age) + ">"
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)
	/*
		// 获取方法字段
		// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
		// 2. 再通过reflect.Type的Field获取其Field
		// 3. 最后通过Field的Interface()得到对应的value
		for i := 0; i < getType.NumField(); i++ {
			field := getType.Field(i)
			value := getValue.Field(i).Interface()
			fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
		}
	*/

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func fibonacci1(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci(c chan int, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	a := F()
	a[0]()
	a[1]()
	a[2]()
}
func F() []func() {
	b := make([]func(),3)
	for i := 0; i < 3; i++ {
		b[i] = func(j int ) func() {
			return func ()  {
				fmt.Println(&j, j)
			}
		}(i)
	}
	return b
}

/*

//"T"
//'T'
func main() {

	/*
	i := 10
	switch i {
	case 1:
		fmt.Println("1")
	case 10:
		fmt.Println("10")
		fallthrough
	default:
		fmt.Println("T")
	}
*/
/*
		c := make(chan int)
	    o := make(chan bool)
	    go func() {
	        for {
	            select {
	                case v := <- c:
	                    println(v)
	                case <- time.After(5 * time.Second):
	                    println("timeout")
	                    o <- true
	                    // 此处的break只是跳出了select循环，并未终止for循环，要用return才能终止这个子进程
	                	break
	            }
	        }
	    }()
	    <- o
*/
/*
		c := make(chan int)
	    quit := make(chan int)
	    go func() {
	        for i := 0; i < 10; i++ {
	            fmt.Println(<-c)
	        }
	        quit <- 0
	    }()
	    fibonacci(c, quit)
*/
/*
		c := make(chan int, 10)
	    go fibonacci(cap(c), c)
	    for i := range c {
	        fmt.Println(i)
	    }
*/
/*
		c := make(chan int, 2) // 修改 2 为 1 就报错，修改 2 为 3 可以正常运行
	    c <- 1
	    c <- 2
	    fmt.Println(<-c)
	    fmt.Println(<-c)
*/
/*
		//go say("world") // 开一个新的 Goroutines 执行
		//say("hello")    // 当前 Goroutines 执行
		a := []int{7, 2, 8, -9, 4, 0}

	    c := make(chan int)
	    go sum(a[:len(a)/2], c)
	    go sum(a[len(a)/2:], c)
	    x, y := <-c, <-c  // receive from c

	    fmt.Println(x, y, x + y)
*/
/*
	user := User{1, "Allen.Wu", 25}

	DoFiledAndMethod(user)
*/
/*
	var num float64 = 1.2345

	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("value: ", reflect.ValueOf(num))
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)
	fmt.Println(value.Type())
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)
*/
/*
	var x float64=3.4
	v:=reflect.ValueOf(x)
	fmt.Println(v.Type())
	student := Student{Human{"lihua", 18, 99999}, []string{"art", "swim"}, 99, 12345}
	//var i Student_i= student
	fmt.Println(student.String())
	println(student.get_name())
	student.set_score(88)
	fmt.Println(student.get_age()) //method inherit
	//defer_test(5)
	fmt.Println(add2(&w, add1))

	//var arr [5]byte
	/odr s:=`Hello
	/World`
	//s[0]='w' error
*/
/*
		a := [3]int{1, 2, 3}
		a[0] = 0
		w = 10
		b := [...]int{1, 2, 3}
		var bs []int = b[:]
		var dict1 map[string]int
		dict1 = make(map[string]int)
		ptr1 := new(int)
		var ptr2 *int
		ptr2 = &w
		fmt.Println(*ptr2)
		*ptr1 = 8
		fmt.Println(*ptr1)
		dict1["one"] = 1
		dict1["two"] = 2
		dict1["three"] = 3
		//bs=b[:]
		b[0] = 0
		c := []byte(s)
		c[0] = 'w'
		bs = append(bs, 4)
		v, ok := dict1["one"]
		fmt.Println(v, ok)
		var sum int


		for sum < 100 {
			sum *= 2
		}
		fmt.Println(sum)

		goto Label1
		for k, v := range dict1 {
			fmt.Println(k, v)
		}
	Label1:
		if x := w; x > 0 {
			fmt.Println("x greater than 0")
		} else {
			fmt.Println("x less than 0")
		}
*/
//}
