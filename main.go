package main

import (
	"archive/zip"
	"flag"
	"github.com/phoobynet/sec-financial-statements/database"
	. "github.com/phoobynet/sec-financial-statements/submissions"
	. "github.com/phoobynet/sec-financial-statements/utils"
	"log"
	"os"
)

func main() {
	sourceZip := flag.String("s", "/Volumes/raid/Downloads/browser/2022q3.zip", "source zip file")
	databasePath := flag.String("database", "", "Path to the database file")
	flag.Parse()

	if _, exists := os.Stat(*sourceZip); os.IsNotExist(exists) {
		flag.PrintDefaults()
		log.Fatalln("Source zip file does not exist")
	}

	if *databasePath == "" {
		flag.PrintDefaults()
		return
	}

	db := database.Init(*databasePath)

	finStatementZip, migrateErr := zip.OpenReader(*sourceZip)

	if migrateErr != nil {
		panic(migrateErr)
	}

	for _, file := range finStatementZip.File {
		switch file.Name {
		case "sub.txt":
			ProcessFile[Submission](db, file, NewSubmission)
		}
	}
}
