package main

import (
	"fmt"
	"time"
)

type Phone interface {
	call()
	play(s string)
}

type NokiaPhone struct {
	name string
	id   int16
}

type IPhone struct {
}

func (iPhone IPhone) play(s string) {
	fmt.Println("playing " + s)
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (nokiaPhone NokiaPhone) play(s string) {
	fmt.Println("playing " + s)
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	//for i:= 0; i < 10; i++ {
	//	go gotest(i)
	//}
	//
	//time.Sleep(10 * time.Millisecond)
	//testInterface()
	//testStruct()
	//go say("world")
	//say("hello")
	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
}

func gotest(i int) {
	j := "liuzhao"
	fmt.Println(j)
	fmt.Printf("go routine %d!\n", i)
}

func testInterface() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}

func testStruct() {
	book := Book{"算法", "王小波", "小说", 34}
	book.showInfo()
	//book.setTitle("黄金时代")
	//book.showInfo()

	doubanBook := DoubanBook{Book{"算法", "王小波", "小说", 34}, 8.6}
	fmt.Println(doubanBook.getScore())

	var behavior, be Behavior
	behavior = &Book{"算法", "王小波", "小说", 34}
	behavior.read("算法")

	be = &DoubanBook{Book{"算法", "王小波", "小说", 34}, 8.6}
	be.read("算法")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func getSequence() func() int {
	i := 0
	fmt.Println("i的值为:", i)
	return func() int {
		i++
		return i
	}
}
