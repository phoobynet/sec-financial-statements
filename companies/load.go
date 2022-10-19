package companies

import (
	tickers "github.com/phoobynet/sec-company-tickers"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func Load(db *gorm.DB) {
	tx := db.Begin()
	log.Printf("Loading companies...")

	companyTickers, companyTickersErr := tickers.Get(nil)

	if companyTickersErr != nil {
		log.Fatalln("Failed to retrieve tickers from the SEC")
	}

	for _, companyTicker := range companyTickers {
		cik := strconv.Itoa(companyTicker.CIK)

		company := &Company{
			CIK:      cik,
			Name:     companyTicker.Name,
			Symbol:   companyTicker.Symbol,
			Exchange: companyTicker.Exchange,
		}

		companyErr := tx.Create(company).Error

		if companyErr != nil {
			log.Fatalln("Failed to create company", companyErr)
		}
	}

	tx.Commit()

}
