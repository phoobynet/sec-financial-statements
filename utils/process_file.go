package utils

import (
	"archive/zip"
	"bufio"
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

	rowBuffer := make([]interface{}, 0, 100)

	counter := 0

	for scanner.Scan() {
		if header {
			header = false
			continue
		}

		row := create(scanner.Text())

		rowBuffer = append(rowBuffer, row)

		if len(rowBuffer) == 100 {
			db.Create(rowBuffer)
			rowBuffer = rowBuffer[:0]
			counter += 100
			log.Printf("%s processed: %10d\n", t.Name(), counter)
		}
	}

	if len(rowBuffer) > 0 {
		db.Create(rowBuffer)
		counter += len(rowBuffer)
		log.Printf("%s processed: %10d\n", t.Name(), counter)
	}
}
