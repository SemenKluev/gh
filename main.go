package main

import "fmt"

// Функция для ввода текста
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

// Функция для создания map с курсами обмена между валютами
func getExchangeRates() map[[2]string]float64 {
	rates := make(map[[2]string]float64)

	// Курсы для конвертации из RUB
	rates[[2]string{"RUB", "USD"}] = 1.0 / 76.79          // RUB -> USD
	rates[[2]string{"RUB", "EUR"}] = 1.0 / (76.79 / 0.85) // RUB -> EUR

	// Курсы для конвертации из USD
	rates[[2]string{"USD", "RUB"}] = 76.79 // USD -> RUB
	rates[[2]string{"USD", "EUR"}] = 0.85  // USD -> EUR

	// Курсы для конвертации из EUR
	rates[[2]string{"EUR", "RUB"}] = 76.79 / 0.85 // EUR -> RUB
	rates[[2]string{"EUR", "USD"}] = 1.0 / 0.85   // EUR -> USD

	return rates
}

// Функция конвертации валюты с использованием map
func convertCurrency(amount float64, fromCurrency string, toCurrency string) float64 {
	rates := getExchangeRates()

	// Получаем курс из map по паре валют
	rate, exists := rates[[2]string{fromCurrency, toCurrency}]
	if !exists {
		return 0 // Если курс не найден
	}

	return amount * rate
}

func main() {
	// Получаем курсы валют
	rates := getExchangeRates()

	// Показываем курс EUR/RUB
	eurRubRate := rates[[2]string{"EUR", "RUB"}]
	fmt.Printf("Курс EUR/RUB: %.2f\n", eurRubRate)

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
