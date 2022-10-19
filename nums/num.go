package nums

type Num struct {
	ADSH     string `gorm:"column:adsh"`
	Tag      string `gorm:"column:tag"`
	Version  string `gorm:"column:version"`
	COReg    string `gorm:"column:coreg"`
	DDate    string `gorm:"column:ddate"`
	QTRS     string `gorm:"column:qtrs"`
	UOM      string `gorm:"column:uom"`
	Value    string `gorm:"column:value"`
	Footnote string `gorm:"column:footnote"`
}
