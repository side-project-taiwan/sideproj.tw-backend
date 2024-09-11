package main

import (
	"fmt"
	"log"
	"os"
	"spt/internal/utility"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	rootDir, envPath, err := utility.GetProjectRootDirAndEnvPath()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Project root directory:", rootDir)
	fmt.Println("Environment file path:", envPath)

	g := gen.NewGenerator(gen.Config{
		OutPath: fmt.Sprintf("%s/internal/gorm_gen/models", rootDir),
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbDatabase)
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	g.UseDB(gormdb)

	// Get all table names from the database
	var tables []string
	err = gormdb.Raw("SHOW TABLES").Scan(&tables).Error
	if err != nil {
		panic("failed to get table names")
	}

	// Generate models for all tables in the database
	for _, table := range tables {
		g.ApplyBasic(g.GenerateModel(table))
	}

	// Old code for generating specific tables (commented out)
	// Generate model for the "table"
	// g.ApplyBasic(g.GenerateModel("project"))
	// g.ApplyBasic(g.GenerateModel("event"))
	g.Execute()
}
