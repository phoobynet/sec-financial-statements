package subs

// Submission data set; this includes one record for each XBRL submission. The set includes fields of information pertinent to the submission and the filing entity. Information is extracted from the SEC’s EDGAR system and the filings submitted to the SEC by registrants.

// Sub is a single submission to the SEC
type Sub struct {
	ADSH       string `gorm:"column:adsh"`
	CIK        string `gorm:"column:cik"`
	Name       string `gorm:"column:name"`
	SIC        string `gorm:"column:sic"`
	CountryBA  string `gorm:"column:countryba"`
	StPrBA     string `gorm:"column:stprba"`
	CityBA     string `gorm:"column:cityba"`
	ZipBA      string `gorm:"column:zipba"`
	BAS1       string `gorm:"column:bas1"`
	BAS2       string `gorm:"column:bas2"`
	BAPh       string `gorm:"column:baph"`
	CountryMA  string `gorm:"column:countryma"`
	StPrMA     string `gorm:"column:stprma"`
	CityMA     string `gorm:"column:cityma"`
	ZipMA      string `gorm:"column:zipma"`
	MAS1       string `gorm:"column:mas1"`
	MAS2       string `gorm:"column:mas2"`
	CountryInc string `gorm:"column:countryinc"`
	StateInc   string `gorm:"column:stprinc"`
	EIN        string `gorm:"column:ein"`
	Former     string `gorm:"column:former"`
	Changed    string `gorm:"column:changed"`
	AFS        string `gorm:"column:afs"`
	WKSI       string `gorm:"column:wksi"`
	FYE        string `gorm:"column:fye"`
	Form       string `gorm:"column:form"`
	Period     string `gorm:"column:period"`
	FY         string `gorm:"column:fy"`
	FP         string `gorm:"column:fp"`
	Filed      string `gorm:"column:filed"`
	Accepted   string `gorm:"column:accepted"`
	PrevRpt    string `gorm:"column:prevrpt"`
	Detail     string `gorm:"column:detail"`
	Instance   string `gorm:"column:instance"`
	NCIKS      string `gorm:"column:nciks"`
	ACIKS      string `gorm:"column:aciks"`
}
