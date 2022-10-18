package queries

import (
	"github.com/phoobynet/sec-financial-statements/database"
	"google.golang.org/appengine/log"
)

const (
	searchSQL = `
		
	`
)

func GetAllCompanies() []CompanySearchResult {
	const sql = `
		select c.cik, c.symbol, c.name, c.exchange, si.office, si.industry_title industry
				from companies c
					   inner join subs s on c.cik = s.cik
					   left join sics si on s.sic = si.code
				group by c.cik
		order by c.name, c.symbol
	`

	db := database.Get()

	var searchResults []CompanySearchResult

	err := db.Raw(sql).Scan(&searchResults).Error

	if err != nil {
		log.Errorf(nil, "error searching for companies: %v", err)
		return nil
	}

	return searchResults
}
