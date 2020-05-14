package main

import (
	"fmt"
	"time"
)

func f1() {
	for {
		fmt.Println("in f1 method")
		time.Sleep(1 * time.Second)
		chan1 <- "chan1 in f1 method"
	}
}

func f2() {
	for {
		fmt.Println("in f2 method")
		time.Sleep(1 * time.Second)
	}
}

var chan1 = make(chan string)

func main() {
	go f1()
	go f2()
	select {} //可以正常运行

	//for{} //可以正常运行

	//for {
	//	select {}   //可以正常运行
	//}

	//for {
	//	select {
	//		case r := <-chan1:
	//			fmt.Println(r)
	//
	//	}   		//可以正常运行,需要配合Channel
	//}
}
