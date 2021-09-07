package reader

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/wesmota/olist_tech_challenge/rest-api/database"
	"github.com/wesmota/olist_tech_challenge/rest-api/domain/entity/author"
)

// CSVLoader is a struct that loads a CSV file into sqlite3 database
type csvLoader struct {
	filePath string
	bdPath   string
}

func NewCSVLoader(filePath string, sqlitePath string) *csvLoader {
	return &csvLoader{filePath: filePath,
		bdPath: sqlitePath}
}

// Load method loads a CSV file into sqlite3 database
func (loader *csvLoader) Load() error {
	//open csv file
	csvFile, err := os.Open(loader.filePath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	fmt.Println("Loading CSV file into database...")
	err = loader.importAuthors(csvFile)
	if err != nil {
		fmt.Println(err)
		//return err
	}

	fmt.Println("File path: ", loader.filePath)
	fmt.Println("Database path: ", loader.bdPath)
	return nil
}

func (loader *csvLoader) importAuthors(csvFile *os.File) error {
	// open author database
	dbConn := database.ConnectDB(loader.bdPath)
	sqlDB, err := dbConn.DB()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer sqlDB.Close()

	// read lines from csv file
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	var newAuthor author.Author
	// use manager to insert into database
	authorRepo := author.InitAuthorRepository(dbConn)
	authorManager := author.NewManager(authorRepo)

	// iterate over lines
	for _, line := range csvLines {
		fmt.Println(line)
		newAuthor = author.Author{
			Name: line[0],
		}
		authorDB, _ := authorManager.FindByName(newAuthor.Name)
		if authorDB.ID == 0 {
			_, err = authorManager.Create(&newAuthor)
			if err != nil {
				fmt.Println(err)
				return err

			}
		}

	}

	return nil
}
