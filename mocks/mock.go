package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Contagem(saida io.Writer) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(saida, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(saida, "Vai!")
}

func main() {
	Contagem(os.Stdout)
}
