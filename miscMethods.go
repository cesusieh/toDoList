package main

import (
	"bufio"
	"fmt"
	"os"
)

func clear() {
	fmt.Printf("\x1bc")
}

func scan() string{
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	x := scanner.Text()
	return x
}
