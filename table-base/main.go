package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func reverse(tx string) (out string) {
	s := strings.Split(tx, "")
	for i := len(s) - 1; i != -1; i-- {
		out = out + s[i]
	}
	return
}

func sum() {
	fmt.Println("sum")
	return
}

func multiply() {
	fmt.Println("multiply")
	return
}

func divide() {
	fmt.Println("divide")
	return
}

func main() {
	fmt.Println(reverse("hola"))
}

var ErrNotEven = errors.New("not an even number")

func EvenOrSleep(n int) error {
	if n%2 == 0 {
		fmt.Println("****  even ")
		fmt.Println(n)

		time.Sleep(time.Duration(n) * time.Second)
		return nil
	}
	fmt.Println("not even")

	return ErrNotEven
}
