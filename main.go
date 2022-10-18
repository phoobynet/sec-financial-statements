package main

import (
	"flag"
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

	if *outPath == "" {
		flag.PrintDefaults()
		return
	}

	zipFileStat, zipFileStatErr := os.Stat(*sourceZip)

	if os.IsNotExist(zipFileStatErr) {
		flag.PrintDefaults()
		log.Fatalln("Source zip file does not exist")
	}

	zipFileStat.Name()

	dbPath := filepath.Join(*outPath, strings.TrimSuffix(zipFileStat.Name(), ".zip")+".db")

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Printf("Database path: %s\n", dbPath)

		log.Printf("Loading database (this may take a while)...")

		db := database.Init(dbPath)
		sics.Load(db)
		companies.Load(db)
		quarterly.Load(db, *sourceZip)
		database.CreateIndexes(db)
	} else {
		log.Printf("Database already exists at %s...starting server", dbPath)
	}
}
