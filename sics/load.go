package sics

import (
	sicScraper "github.com/phoobynet/sec-sic-scraper"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func Load(db *gorm.DB) {
	tx := db.Begin()
	log.Printf("Loading SIC codes...")
	s, err := sicScraper.Get(nil)

	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range s {
		sicErr := tx.Create(&SIC{
			Code:          strconv.Itoa(item.Code),
			IndustryTitle: item.IndustryTitle,
			Office:        item.Office,
		}).Error

		if sicErr != nil {
			log.Fatalln(sicErr)
		}
	}

	tx.Commit()

	log.Printf("Loading SIC codes...done")
}
