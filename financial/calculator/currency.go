package currency

import (
	"fmt"
	"math/big"
)

func ConvertBRLToUSD(brlValue float64) float64 {
	// conexão API de confiança e segura
	// pra pegar em tempo real
	// o valor de cotação do dolar comercial

	var x float64

	return brlValue * x
}

func SumTwoCurrencysUsingFloat64(a float64, b float64) float64 {
	var newA, newB, temp big.Float

	newA.SetFloat64(a)

	newB.SetFloat64(b)

	temp.Add(&newA, &newB)

	sum, acc := temp.Float64()

	fmt.Println(acc)

	return sum
}
