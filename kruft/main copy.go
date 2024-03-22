//package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)


func GetLines(r io.Reader, process func(number float64) error) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() { // Читаем строку, преобразуем её в float64 и сохраняем в слайс
		t := scanner.Text()
		if t == "" {
			fmt.Println("[WARN] blank line skipped")
			continue
		}

		number, err := strconv.ParseFloat(t, 64)
		if err != nil {
			return fmt.Errorf("cannot parse '%s': %w", t, err) // Возвращаем слайс с прочитанными числами и ошибку, если произошла ошибка преобразования
		}

		if err := process(number); err != nil {
			return fmt.Errorf("cannot process '%f': %w", number, err)
		}
	}

	if scanner.Err() != nil {
		return scanner.Err() // Возвращаем слайс и ошибку, если в сканере произошла ошибка
	}

	return nil
}

func main() {
	file, err := os.Open(defaultFileName) // Открываем файл на чтение
	if err != nil {
		panic(fmt.Errorf("cannot open file %s for reading: %w", defaultFileName, err)) // Возвращаем пустую строку и ошибку, если не удалось открыть файл
	}
	defer file.Close() // Закрываем файл по завершении функции

	results := make([]float64, 0)

	err = GetLines(file, func(number float64) error {
		results = append(results, number)
		return nil
	})
	if err != nil {
		panic(fmt.Errorf("cannot parse input: %w", err))
	}

	fmt.Println("total number of scanned floats: ", len(results))
}
