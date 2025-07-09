package main

import (
	"fmt"
)

type ErrChamada struct {
	Mensagem string
	Codigo   int
}

func (e ErrChamada) Error() string {
	return fmt.Sprintf("Err %d: %s", e.Codigo, e.Mensagem)
}

func dividirDoisNumeros(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, ErrChamada{Mensagem: "valor igual a  0", Codigo: 500}
	}

	return dividendo / divisor, nil
}

func dividirResultadoPanic(dividendo, divisor int) int {
	return dividendo / divisor
}

func main() {
	resultado, err := dividirDoisNumeros(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(resultado)

}
