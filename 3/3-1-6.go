package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"Деревня", "Село", "Село2"}, // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Большой город"},   // города с населением 100-999 тыс. человек
		1000: []string{"Миллионик"},                // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"Село":      100,
		"Миллионик": 500,
		"Деревня":   10,
		"Село2":     50,
		"Город":     150,
	}

	for val := range groupCity {
		for val2 := range groupCity[val] {
			fmt.Printf("Цикл по справочнику: ключ - %d, названия города - %v\n", val, groupCity[val][val2])
			if val != 100 {
				fmt.Printf("По условию задачи удаляем все значения по ключу %d\n", val)
				delete(cityPopulation, groupCity[val][val2])
			}
		}
	}

	for value := range cityPopulation {
		fmt.Printf("Город %v с населением %d ", value, cityPopulation[value])
	}
}
