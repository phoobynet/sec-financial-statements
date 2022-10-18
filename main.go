package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/phoobynet/sec-financial-statements/companies"
	"github.com/phoobynet/sec-financial-statements/database"
	"github.com/phoobynet/sec-financial-statements/quarterly"
	"github.com/phoobynet/sec-financial-statements/queries"
	"github.com/phoobynet/sec-financial-statements/sics"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	sourceZip := flag.String("z", "", "source zip file")
	outPath := flag.String("o", "", "The output directory")
	port := flag.Int("p", 3000, "The port to listen on")
	flag.Parse()

	if *outPath == "" {
		flag.PrintDefaults()
		return
	}

	zipFileStat, zipFileStatErr := os.Stat(*sourceZip)

	if os.IsNotExist(zipFileStatErr) {
		flag.PrintDefaults()
		log.Fatalln("Source zip file does not exist")
	}

	zipFileStat.Name()

	dbPath := filepath.Join(*outPath, strings.TrimSuffix(zipFileStat.Name(), ".zip")+".db")

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Printf("Database path: %s\n", dbPath)

		log.Printf("Loading database (this may take a while)...")

		db := database.Init(dbPath)
		sics.Load(db)
		companies.Load(db)
		quarterly.Load(db, *sourceZip)
		database.CreateIndexes(db)
	} else {
		database.Init(dbPath)
		log.Printf("Database already exists at %s...starting server", dbPath)
	}

	// region web server
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/api/search", func(c *fiber.Ctx) error {
		query := c.Query("q")
		if query == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		searchResults := queries.SearchForCompanies(query)

		return c.JSON(searchResults)
	})

	app.Get("/api/companies", func(c *fiber.Ctx) error {
		return c.JSON(queries.GetAllCompanies())
	})

	app.Static("/", "./public")

	log.Fatalln(app.Listen(fmt.Sprintf(":%d", *port)))
	// endregion
}
