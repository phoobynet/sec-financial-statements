package companies

type Company struct {
	CIK      string `gorm:"index"`
	Name     string
	Symbol   string `gorm:"index"`
	Exchange string `gorm:"index"`
}
