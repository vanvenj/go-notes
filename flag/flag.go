package main

import (
	"flag"
	"fmt"
)

// https://gobyexample.com/command-line-flags
func main() {

	wordPtr := flag.String("word", "foo", "a string") // 指针
	numbPtr := flag.Int("numb", 10, "an int")
	boolPtr := flag.Bool("bool", false, "a bool")

	var svar string                                      // 字符串
	flag.StringVar(&svar, "svar", "bar", "a string var") // 将svar的指针指向flag

	flag.Parse()

	fmt.Println("word:", *wordPtr) // 取指针的数据
	fmt.Println("word:", wordPtr)  // 打印word的地址
	fmt.Println("numb:", *numbPtr)
	fmt.Println("bool:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("svar:", &svar)       // 打印svar地址
	fmt.Println("tail:", flag.Args()) //打印没有flag的数据

}
