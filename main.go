package main

import "fmt"

func getText() {
	var str string
	fmt.Print("Введите текст: ")
	fmt.Scan(&str)

	fmt.Println("Вы ввели данный текст:",str)
}

func main() {
	const UsdEur = 0.85
	const UsdRub = 76.79
	const EurRub = UsdRub / UsdEur

	fmt.Println(EurRub)

	getText()
}

func convertCurrency (number float64,currency1 string,currency2 string) float64{
	return 0
}