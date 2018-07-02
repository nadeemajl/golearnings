package main

import "fmt"

func main() {
	fmt.Println("Hi there!")

	s1 := "hello"
	sb1 := []byte(s1)
	fmt.Println("hello as bytes ", sb1)

	s2 := "world"
	sb2 := []byte(s2)
	fmt.Println("world as bytes", sb2)

	s3 := append(sb1, sb2...)

	fmt.Println("Joined bytes array ", s3)

	fmt.Println("first 5 bytes are ", s3[:5])
	fmt.Println("6th byte to end of splice ", s3[5:])
}
