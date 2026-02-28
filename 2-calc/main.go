package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var operation string
	var input string
	var numberSlice []float64

	for {
		fmt.Print("Введите вашу операцию (AVG/SUM/MED): ")
		fmt.Scan(&operation)

		if operation != "AVG" && operation != "SUM" && operation != "MED" {
			fmt.Println("Не коректное значение")
			continue
		}

		fmt.Print("Введите числа через запятую: ")
		fmt.Scan(&input)

		stringSlice := strings.Split(input, ",")

		numberSlice = []float64{}

		for _, s := range stringSlice {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}

			num, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Printf("Ошибка конвертации '%s': %v\n", s, err)
				continue
			}
			numberSlice = append(numberSlice, num)
		}

		if len(numberSlice) == 0 {
			fmt.Println("Не введено ни одного корректного числа")
			continue
		}

		stop := true

		for i := 0; i < len(numberSlice); i++ {
			if numberSlice[i] < 0 {
				fmt.Println("Число не может быть отрицательным")
				stop = false
				break
			}
		}

		if stop != true {
			continue
		}

		break
	}

	getOperation(operation, numberSlice)
}

func getOperation(operation string, num []float64) {
	switch operation {
	case "AVG":
		var sum float64
		for _, value := range num {
			sum += value
		}
		avg := sum / float64(len(num))
		fmt.Println("Среднее арифметическое из ваших чисел составляет:", avg)
	case "SUM":
		var sum float64
		for _, value := range num {
			sum += value
		}
		fmt.Println("Сумма ваших чисел составляет:", sum)
	case "MED":
		sortedNum := make([]float64, len(num))
		copy(sortedNum, num)

		sort.Float64s(sortedNum)

		var median float64
		n := len(sortedNum)

		if n%2 == 0 {
			median = (sortedNum[n/2-1] + sortedNum[n/2]) / 2.0
		} else {
			median = sortedNum[n/2]
		}

		fmt.Println("Медиана ваших чисел составляет:", median)
	}
}
