package models

import (
	"gorm.io/gorm"
	"time"
)

type OrganisationSubmission struct {
	gorm.Model
	AccessionNumber   string `gorm:"index:idx_submissions_accession_number"`
	CIK               int    `gorm:"index:idx_submissions_cik"`
	FiscalYearEnd     string // mmdd, e.g. 0331 for March 31
	Form              string `gorm:"index:idx_submissions_form"`
	Period            time.Time
	FiscalYearFocus   int
	FiscalPeriodFocus string //Fiscal Period Focus (as defined in EFM Ch. 6) within Fiscal Year. The 10-Q for the 1st, 2nd and 3rd quarters would have a fiscal period focus of Q1, Q2 (or H1), and Q3 (or M9) respectively, and a 10-K would have a fiscal period focus of FY.
	PreviousReport    bool
	Detail            string
	Instance          string
	NumberOfCIKs      int
	AdditionalCIKs    string
}

type Organisation struct {
	gorm.Model
	CIK               int    `gorm:"index:idx_organizations_cik"`
	Symbol            string `gorm:"index:idx_organizations_symbol"`
	Name              string
	FormerName        string
	FormerNameChanged time.Time
	SIC               int
}

type OrganisationAddress struct {
	gorm.Model
	CIK             int    `gorm:"index:idx_addresses_cik"`
	Type            string // mailing or business
	Street1         string
	Street2         string
	City            string
	StateOrProvince string
	Zip             string
	Country         string
}
