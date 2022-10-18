package tags

type Tag struct {
	Tag      string `gorm:"column:tag"`
	Version  string `gorm:"column:version"`
	Custom   string `gorm:"column:custom"`
	Abstract string `gorm:"column:abstract"`
	DataType string `gorm:"column:datatype"`
	IOrd     string `gorm:"column:iord"`
	CRDR     string `gorm:"column:crdr"`
	TLabel   string `gorm:"column:tlabel"`
	Doc      string `gorm:"column:doc"`
}
