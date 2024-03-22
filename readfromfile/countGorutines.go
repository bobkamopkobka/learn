package readfromfile

// Функция для обработки строк в отдельной горутине

import "sync"

func ProcessLines(lines []string, wg *sync.WaitGroup, resultChan chan<- map[string]int) {
	// Отложенное уменьшение счетчика WaitGroup при завершении функции
	defer wg.Done()

	// Создаем карту для подсчета вхождений в пределах горутины
	localCounts := make(map[string]int)

	// Подсчитываем вхождения строк в пределах горутины
	for _, line := range lines {
		localCounts[line]++
	}

	// Отправляем результат в канал
	resultChan <- localCounts
}
