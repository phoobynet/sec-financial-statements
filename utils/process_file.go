package utils

import (
	"archive/zip"
	"bufio"
	"fmt"
	"gorm.io/gorm"
	"io"
	"log"
	"reflect"
)

func ProcessFile[T any](db *gorm.DB, file *zip.File, create func(line string) *T) {
	t := reflect.TypeOf(new(T))
	reader, readerErr := file.Open()

	if readerErr != nil {
		panic(readerErr)
	}

	defer func(subFileReader io.ReadCloser) {
		_ = subFileReader.Close()
	}(reader)

	scanner := bufio.NewScanner(reader)

	header := true

	fmt.Printf("Processing %s\n", file.Name)

	counter := 0

	for scanner.Scan() {
		if header {
			header = false
			continue
		}

		row := create(scanner.Text())

		rowErr := db.Create(row).Error

		if rowErr != nil {
			log.Fatalf("%s error inserting buffer: %v", t.Name(), rowErr)
		}
		counter++

		if counter%1000 == 0 {
			fmt.Printf("Processed %d rows\n", counter)
		}
	}
}
