package readfromfile

import (
	"bufio"
	"fmt"
	"io"
)

func GetLines(reaDd io.Reader) ([]string, error) {

	scanner := bufio.NewScanner(reaDd)
	result := make([]string, 0)
	for scanner.Scan() {
		teXtt := scanner.Text()
		if teXtt == "" {
			fmt.Println("Пустая строка, пропускаем.")
			continue
		}
		fmt.Println(teXtt)
		result = append(result, teXtt)
	}

	if scanner.Err() != nil {
			return nil, scanner.Err()
		}

	return result, scanner.Err()
}
