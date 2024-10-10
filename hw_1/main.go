package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	fmt.Fprintf(out, "Ьэхося Чшхщ, ьсэозц охрыю ю чмъмшм - ьышъсцехц эмфжтн\n")
}
