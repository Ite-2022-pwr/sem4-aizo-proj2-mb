package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func SaveMapToCSV(inputMap map[int]float64, filename string) {
	// Open file
	filename += ".csv"
	file, err := os.Create(filename)
	Check(err)
	fmt.Println("Saving to file: ", filename)
	log.Println("Saving to file: ", filename)
	defer file.Close()

	// Write to file
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for key, value := range inputMap {
		err := writer.Write([]string{strconv.Itoa(key), strconv.FormatFloat(value, 'f', -1, 64)})
		Check(err)
	}
}
