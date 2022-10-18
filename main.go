package main

import (
	"archive/zip"
	"flag"
	"fmt"
	tickers "github.com/phoobynet/sec-company-tickers"
	"github.com/phoobynet/sec-financial-statements/companies"
	"github.com/phoobynet/sec-financial-statements/database"
	. "github.com/phoobynet/sec-financial-statements/nums"
	. "github.com/phoobynet/sec-financial-statements/pres"
	"github.com/phoobynet/sec-financial-statements/sics"
	. "github.com/phoobynet/sec-financial-statements/submissions"
	. "github.com/phoobynet/sec-financial-statements/tags"
	. "github.com/phoobynet/sec-financial-statements/utils"
	sicScraper "github.com/phoobynet/sec-sic-scraper"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

	log.Printf("Loading SIC codes...")
	s, err := sicScraper.Get(nil)

	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range s {
		sicErr := db.Create(&sics.SIC{
			Code:          strconv.Itoa(item.Code),
			IndustryTitle: item.IndustryTitle,
			Office:        item.Office,
		}).Error

		if sicErr != nil {
			log.Fatalln(sicErr)
		}
	}

	log.Printf("Loading SIC codes...done")

	log.Printf("Loading companies...")

	companyTickers, companyTickersErr := tickers.Get(nil)

	if companyTickersErr != nil {
		log.Fatalln("Failed to retrieve tickers from the SEC")
	}

	for _, companyTicker := range companyTickers {
		cik := strconv.Itoa(companyTicker.CIK)

		company := &companies.Company{
			CIK:      cik,
			Name:     companyTicker.Name,
			Symbol:   companyTicker.Symbol,
			Exchange: companyTicker.Exchange,
		}

		companyErr := db.Create(company).Error

		if companyErr != nil {
			log.Fatalln("Failed to create company", companyErr)
		}
	}

	finStatementZip, migrateErr := zip.OpenReader(*sourceZip)

	if migrateErr != nil {
		panic(migrateErr)
	}

	for _, file := range finStatementZip.File {
		switch file.Name {
		case "sub.txt":
			ProcessFile[Submission](db, file, func(line string) *Submission {
				return ProcessLine[Submission](line)
			})
		case "num.txt":
			ProcessFile[Num](db, file, func(line string) *Num {
				return ProcessLine[Num](line)
			})
		case "pre.txt":
			ProcessFile[Pre](db, file, func(line string) *Pre {
				return ProcessLine[Pre](line)
			})
		case "tag.txt":
			ProcessFile[Tag](db, file, func(line string) *Tag {
				return ProcessLine[Tag](line)
			})
		}
	}

	log.Printf("Creating indexes...")

	log.Printf("Creating indexes...submissions")
	// submissions indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_submissions_cik ON submissions (cik)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_submissions_adsh ON submissions (adsh)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_submissions_form ON submissions (form)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_submissions_fp ON submissions (fp)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_submissions_fy ON submissions (fy)")

	log.Printf("Creating indexes...nums")
	// num indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_nums_adsh ON nums (adsh)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_nums_tag_version ON nums (tag, version)")

	log.Printf("Creating indexes...tags")
	// tag indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_tags_tag_version ON tags (tag, version)")

	log.Printf("Creating indexes...pres")
	// pre indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_pres_adsh ON pres (adsh)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_pres_tag_version ON pres (tag, version)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_pres_adsh_tag_version ON pres (adsh,tag, version)")

	log.Printf("Creating indexes...done")
}
