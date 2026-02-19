package main

import "fmt"

func getText() {
	var str string
	fmt.Print("Введите текст: ")
	fmt.Scan(&str)
	fmt.Println("Вы ввели данный текст:", str)
}

// Функция для ввода исходной валюты
func getCurrency() string {
	var currency string
	for {
		fmt.Print("Введите исходную валюту RUB/USD/EUR: ")
		fmt.Scan(&currency)

		if currency != "RUB" && currency != "USD" && currency != "EUR" {
			fmt.Println("Вы ввели не корректную валюту")
			continue
		}
		return currency
	}
}

// Функция для ввода суммы
func getAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите сумму: ")
		fmt.Scan(&amount)

		if amount <= 0 {
			fmt.Println("Вы ввели не корректное число")
			continue
		}
		return amount
	}
}

// Функция для ввода целевой валюты
func getTargetCurrency(sourceCurrency string) string {
	var targetCurrency string
	for {
		var prompt string
		if sourceCurrency == "RUB" {
			prompt = "Введите целевую валюту USD/EUR: "
		} else if sourceCurrency == "USD" {
			prompt = "Введите целевую валюту RUB/EUR: "
		} else { // EUR
			prompt = "Введите целевую валюту RUB/USD: "
		}

		fmt.Print(prompt)
		fmt.Scan(&targetCurrency)

		// Проверяем, что целевая валюта допустима и не равна исходной
		if targetCurrency != "RUB" && targetCurrency != "USD" && targetCurrency != "EUR" {
			fmt.Println("Вы ввели не корректную валюту")
			continue
		}

		if targetCurrency == sourceCurrency {
			fmt.Println("Целевая валюта не может совпадать с исходной")
			continue
		}

		return targetCurrency
	}
}

// Функция конвертации валюты
func convertCurrency(amount float64, fromCurrency string, toCurrency string) float64 {
	const UsdEur = 0.85
	const UsdRub = 76.79
	const EurRub = UsdRub / UsdEur
	const EurUsd = 1 / UsdEur
	const RubUsd = 1 / UsdRub
	const RubEur = 1 / EurRub

	switch {
	case fromCurrency == "RUB" && toCurrency == "USD":
		return amount * RubUsd
	case fromCurrency == "RUB" && toCurrency == "EUR":
		return amount * RubEur
	case fromCurrency == "USD" && toCurrency == "RUB":
		return amount * UsdRub
	case fromCurrency == "USD" && toCurrency == "EUR":
		return amount * UsdEur
	case fromCurrency == "EUR" && toCurrency == "RUB":
		return amount * EurRub
	case fromCurrency == "EUR" && toCurrency == "USD":
		return amount * EurUsd
	default:
		return 0
	}
}

func main() {
	// Показываем курс EUR/RUB
	const UsdEur = 0.85
	const UsdRub = 76.79
	const EurRub = UsdRub / UsdEur
	fmt.Printf("Курс EUR/RUB: %.2f\n", EurRub)

	// Дополнительная функция из задания
	getText()

	// Основной цикл конвертации
	for {
		fmt.Println("\n--- Конвертер валют ---")

		// Шаг 1: Ввод исходной валюты
		sourceCurrency := getCurrency()

		// Шаг 2: Ввод суммы
		amount := getAmount()

		// Шаг 3: Ввод целевой валюты
		targetCurrency := getTargetCurrency(sourceCurrency)

		// Вычисление и вывод результата
		result := convertCurrency(amount, sourceCurrency, targetCurrency)
		fmt.Printf("Результат: %.2f %s = %.2f %s\n",
			amount, sourceCurrency, result, targetCurrency)

		// Спрашиваем, хочет ли пользователь продолжить
		var choice string
		fmt.Print("\nХотите выполнить еще одну конвертацию? (да/нет): ")
		fmt.Scan(&choice)
		if choice != "да" && choice != "Да" && choice != "yes" && choice != "Yes" {
			break
		}
	}

	fmt.Println("Программа завершена")
}
