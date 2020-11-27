package main

import (
	"fmt"
	"log"
	"os"
)

// go build -o ./driver/main ./driver/main.go

// ./driver/main 3>out.txt

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Stdin.Fd(), os.Stdin.Name())
	fmt.Println(os.Stdout.Fd(), os.Stdout.Name())
	fmt.Println(os.Stderr.Fd(), os.Stderr.Name())

	// Создаем новый файловый дескриптор под номером 3
	f := os.NewFile(3, "fd3")

	fmt.Println(f.Fd(), f.Name())

	defer func () {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if _, err := f.Write([]byte("FOOBAR\n")); err != nil {
		log.Fatal(err)
	}
}
