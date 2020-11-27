package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("./driver/main")

	// Пайпим стандартные дескрипторы
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Пайпим дополнительные дескрипторы
	// Правда в доках сказано что в винде не поддерживается ExtraFiles
	// https://golang.org/pkg/os/exec/#Cmd
	cmd.ExtraFiles = []*os.File{os.Stdout}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- END ---")
}
