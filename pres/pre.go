package pres

type Pre struct {
	ADSH    string `gorm:"column:adsh"`
	Report  string `gorm:"column:report"`
	Line    string `gorm:"column:line"`
	Stmt    string `gorm:"column:stmt"`
	Inpth   string `gorm:"column:inpth"`
	RFile   string `gorm:"column:rfile"`
	Tag     string `gorm:"column:tag"`
	Version string `gorm:"column:version"`
	PLabel  string `gorm:"column:plabel"`
}
