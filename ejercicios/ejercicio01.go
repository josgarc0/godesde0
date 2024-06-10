package ejercicios

import (
	"strconv"
)

func RestoVariables(numChar string) (int, string) {
	numero, err := strconv.Atoi(numChar)
	var texto string
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}

	if numero > 100 {
		texto = "Es mayor a 100"
	} else {
		texto = "Es menor a 100"
	}
	return numero, texto
}
