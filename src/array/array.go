package main

import (
	"errors"
	"fmt"
)

func main() {
	//var balance [] float32
	//balance = []float32{100.0, 1.0, 2.35}
	//for i:= 0; i < len(balance); i ++ {
	//	fmt.Println(balance[i])
	//}
	balance := []int32{5, 1, -9, 6, 45, 10, 5666, 6}

	//balance := []int32{}
	//avg := getAverage(balance)
	//fmt.Println(avg)
	//medium, err := getMedium(balance)
	//if err == nil {
	//	fmt.Println(medium)
	//} else {
	//	print(err.Error())
	//}
	c := make(chan int32)
	go sum(balance[:len(balance)/2], c)
	go sum(balance[len(balance)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}

func getAverage(array []int32) float32 {
	var sum, result float32
	len := len(array)
	for _, num := range array {
		sum += float32(num)
	}
	result = sum / float32(len)
	return result

}

func getMedium(array []int32) (medium float32, err error) {
	if len(array) == 0 {
		err = errors.New("数组为空")
		return
	}
	size := len(array)
	if size&1 == 0 {
		medium = float32(array[(size-1)>>1]+array[size>>1]) / 2
		return
	}
	medium = float32(array[size>>1])
	return
}

func sum(array []int32, c chan int32) {
	var sum int32 = 0
	for _, num := range array {
		sum += num
	}
	c <- sum
}
