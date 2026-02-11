package main

import "fmt"

func main() {
	const UsdEur = 0.85
	const UsdRub = 76.79
	const EurRub = UsdRub / UsdEur

	fmt.Println(EurRub)

	var str string
	fmt.Print("Введите текст: ")
	fmt.Scan(&str)

	fmt.Println("Вы ввели данный текст:",str)

	

	}

	func convertCurrency (number float64,currency1 string,currency2 string) float64 {
		
	}