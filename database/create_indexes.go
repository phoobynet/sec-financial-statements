package database

import (
	"gorm.io/gorm"
	"log"
)

func CreateIndexes(db *gorm.DB) {
	log.Printf("Creating indexes...")
	log.Printf("Creating indexes...submissions")
	// submissions indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_submissions_cik ON submissions (cik)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_submissions_adsh ON submissions (adsh)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_submissions_form ON submissions (form)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_submissions_fp ON submissions (fp)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_submissions_fy ON submissions (fy)")

	log.Printf("Creating indexes...nums")
	// num indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_nums_adsh ON nums (adsh)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_nums_tag_version ON nums (tag, version)")

	log.Printf("Creating indexes...tags")
	// tag indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_tags_tag_version ON tags (tag, version)")

	log.Printf("Creating indexes...pres")
	// pre indexes
	db.Exec("CREATE INDEX IF NOT EXISTS idx_pres_adsh ON pres (adsh)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_pres_tag_version ON pres (tag, version)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_pres_adsh_tag_version ON pres (adsh,tag, version)")

	log.Printf("Creating indexes...done")
}
