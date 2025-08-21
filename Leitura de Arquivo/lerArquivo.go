package main

import (
	"fmt"
	"io"
	"os"
)
func main() {
	lerArquivoPorBytes("arquivo.txt")
}
func lerArquivoPorBytes(filepath string) {
	data, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}

	defer data.Close()
	pedacobyte := 100
	buffer := make([]byte, pedacobyte)

	for {
		n, err := data.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Erro ao ler o arquivo:", err)
			return
		}
		if err == io.EOF {
			fmt.Println("Fim do arquivo.")
			break
		}
		// fmt.Println(buffer)
		fmt.Println(buffer[:n])	
	}
}
