package main

import (
	"flag"
	"fmt"
	"github.com/phoobynet/sec-financial-statements/companies"
	"github.com/phoobynet/sec-financial-statements/database"
	"github.com/phoobynet/sec-financial-statements/quarterly"
	"github.com/phoobynet/sec-financial-statements/sics"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	sourceZip := flag.String("z", "", "source zip file")
	outPath := flag.String("o", "", "The output directory")
	flag.Parse()

	zipFileStat, zipFileStatErr := os.Stat(*sourceZip)

	if os.IsNotExist(zipFileStatErr) {
		flag.PrintDefaults()
		log.Fatalln("Source zip file does not exist")
	}

	if *outPath == "" {
		flag.PrintDefaults()
		return
	}

	zipFileStat.Name()

	dbPath := filepath.Join(*outPath, strings.TrimSuffix(zipFileStat.Name(), ".zip")+".db")

	if _, err := os.Stat(dbPath); !os.IsNotExist(err) {
		log.Fatalln("Database already exists")
	}

	fmt.Printf("Database path: %s\n", dbPath)

	db := database.Init(dbPath)
	sics.Load(db)
	companies.Load(db)
	quarterly.Load(db, *sourceZip)
	database.CreateIndexes(db)
}
