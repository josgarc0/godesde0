package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func TabladeMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("ingrese un n+umero : ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", numero, i, i*numero)
	}

}
