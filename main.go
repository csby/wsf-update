package main

import "fmt"

func main() {
	err := host.Run()
	if err != nil {
		fmt.Println(err)
	}
}
