package queries

import (
	"github.com/phoobynet/sec-financial-statements/database"
	"google.golang.org/appengine/log"
)

func SearchForCompanies(query string) []CompanySearchResult {
	const sql = `
		select c.cik, c.symbol, c.name, c.exchange, si.office, si.industry_title industry
		from companies c
			   inner join subs s on c.cik = s.cik
			   left join sics si on s.sic = si.code
		where c.symbol = ?
		   or c.name like ?
		group by c.cik
		collate NOCASE
	`
	db := database.Get()

	var searchResults []CompanySearchResult

	queryErr := db.Raw(searchSQL, query, "%"+query+"%").Scan(&searchResults).Error

	if queryErr != nil {
		log.Errorf(nil, "error searching for companies: %v", queryErr)
		return nil
	}

	return searchResults
}
