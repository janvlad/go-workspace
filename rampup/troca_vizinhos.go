package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Insira a sequência de inteiros: ")
	seq, _ := reader.ReadString('\n')

	seqSlice := strings.Fields(seq)
	seqArray := make([]int, len(seqSlice))

	for i := 0; i < len(seqSlice); i++ {
		seqArray[i], _ = strconv.Atoi(seqSlice[i])
	}
	for i := 0; i < len(seqArray)-1; i += 2 {
		seqArray[i], seqArray[i+1] = seqArray[i+1], seqArray[i]
	}

	fmt.Println("Vizinhos trocados: ", seqArray)
}
