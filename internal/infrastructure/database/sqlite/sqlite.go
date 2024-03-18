package sqlite

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Sqlite *gorm.DB

func SQLiteSetup() {
	file, _ := os.OpenFile("gorm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	dbFilePath := "development.db"

	// If the file doesn't exist, create it

	// You can use a database creation command here if you're using a specific database
	// For example, if you're using SQLite, you can use the following command:
	// cmd := exec.Command("sqlite3", dbFilePath)

	// In this example, we'll just create an empty file for demonstration purposes
	if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
		file, err := os.Create(dbFilePath)
		if err != nil {
			fmt.Println("Error creating database file:", err)
			return
		}
		defer file.Close()

		fmt.Println("Database file created successfully.")
	}

	sqlite, err := gorm.Open(
		sqlite.Open(dbFilePath),
		&gorm.Config{
			Logger: logger.New(
				log.New(file, "\r\n", log.LstdFlags), // set log file and format
				logger.Config{
					LogLevel: logger.Info, // set log level
				},
			),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	Sqlite = sqlite

	// Development only db seed to improve development agility
	actorMigration()
	experimentMigration()
	goalMigration()
	variationMigration()

	fmt.Println("Seed completed...")
}

func actorMigration() {
	seeds := "DROP TABLE IF EXISTS actors;" +

		"CREATE TABLE IF NOT EXISTS actors (" +
		"id text," +
		"variation_id text" +
		"identifier text" +
		");"

	fmt.Println("Running actor seed...")

	result := Sqlite.Exec(seeds)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func experimentMigration() {
	seeds := "DROP TABLE IF EXISTS experiments;" +

		"CREATE TABLE IF NOT EXISTS experiments (" +
		"id text," +
		"name text," +
		"key text," +
		"status bigint" +
		");"

	fmt.Println("Running order seed...")

	result := Sqlite.Exec(seeds)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func goalMigration() {
	seeds := "DROP TABLE IF EXISTS goals;" +

		"CREATE TABLE IF NOT EXISTS goals (" +
		"id text," +
		"actor_id text," +
		"key text" +
		");"

	fmt.Println("Running product seed...")

	result := Sqlite.Exec(seeds)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}

func variationMigration() {
	seeds := "DROP TABLE IF EXISTS variations;" +

		"CREATE TABLE IF NOT EXISTS variations (" +
		"id text," +
		"experiment_id text," +
		"key text," +
		"percentage decimal" +
		");"

	fmt.Println("Running product seed...")

	result := Sqlite.Exec(seeds)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
}
