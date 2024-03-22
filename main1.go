package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

// разбираем входящий слайс на горутины, обрабатываем их в канале
func calcGorutines(inputData []float64) ([]float64, error) {
    resultChan := make(chan float64)    // Канал для передачи результатов подсчета из горутин в основную горутину
    doneChan := make(chan struct{})     // Канал для сигнала об окончании работы горутин
    var wg sync.WaitGroup               // Используем WaitGroup для ожидания завершения горутин
    numOfWorkers := 4                   // Определяем количество горутин, которые будут вычислять корень квадратный
    dataForWorkers := len(inputData) / numOfWorkers
    remPerWorkers := len(inputData) % numOfWorkers

    for i := 0; i < numOfWorkers; i++ {
        start := i * dataForWorkers
        end := start + dataForWorkers
        if i == numOfWorkers-1 {
            end += remPerWorkers
        }
        wg.Add(1)
        go calculateSquares(inputData[start:end], &wg, resultChan)
    }

    // Запускаем горутину, которая закроет канал после завершения всех горутин
    go func() {
        wg.Wait()
        close(resultChan)
        close(doneChan) // Сигнализируем об окончании работы горутин
    }()

    // Обработка результатов из канала
    var result []float64
    for res := range resultChan {
        result = append(result, res)
    }

    <-doneChan // Ждем завершения всех горутин

    return result, nil
}

// Горутина для вычисления квадратного корня
func calculateSquares(dataL []float64, wg *sync.WaitGroup, resultChan chan<- float64) {
    defer wg.Done() // уменьшаем счетчик, когда горутина завершена
    for _, value := range dataL {
        squareRoot := math.Sqrt(value)
        resultChan <- squareRoot
    }
}

// чтение файла, разбор по строкам, выборка цифр, на выходе слайс float64
func readFiles(inputFileNames string) ([]float64, error) {
    fileXX, err := os.Open(inputFileNames)
    if err != nil {
        fmt.Println("Ошибка открытия файла:", err)
        return nil, err
    }
    defer fileXX.Close()
    resultXX := make([]float64, 0)
    scannerXX := bufio.NewScanner(fileXX)
    for scannerXX.Scan() {
        lineXX := scannerXX.Text()
        resXA, err := strconv.ParseFloat(lineXX, 64)
        if err != nil {
            //fmt.Println("Это не число.", err)
            continue
        } else {
            resultXX = append(resultXX, resXA)
        }
        if err := scannerXX.Err(); err != nil {
            fmt.Println("Ошибка чтения файла:", err)
            panic(err)
        }
    }
    return resultXX, err
}

func main() {
    fileToRead := "my.txt"
    readedData, err := readFiles(fileToRead)
    if err != nil {
        panic(err)
    } else {
        fmt.Println(readedData)
    }
    result, err := calcGorutines(readedData)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%.3f\n", result)
}
