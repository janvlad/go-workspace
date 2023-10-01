package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var check bool
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Informe a sequencia: ")
	seq, _ := reader.ReadString('\n')

Bigloop:
	for i := 2; i < len(seq); i += 2 {
		for j := 0; j < i; j += 2 {
			if seq[j] == seq[i] {
				check = true
				break Bigloop
			}
		}
	}
	fmt.Println(check)
}
