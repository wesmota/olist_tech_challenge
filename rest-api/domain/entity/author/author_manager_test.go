package author

import (
	"fmt"
	"testing"

	"github.com/wesmota/olist_tech_challenge/rest-api/database"
)

const DBPath = "../../../database/olist.db"

func TestCRUD(t *testing.T) {
	dbConn := database.ConnectDB(DBPath)
	sqlDB, err := dbConn.DB()
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	}
	defer sqlDB.Close()
	authorRepo := InitAuthorRepository(dbConn)
	authorManager := NewManager(authorRepo)
	author := Author{
		Name: "Test Author",
	}
	author2 := Author{
		Name: "Test Author",
	}

	t.Run("Create a new author", func(t *testing.T) {

		authorManager.Create(&author)
		if author.ID == 0 {
			t.Fail()
		}
		fmt.Println(author)

	})
	t.Run("Shouldn't create duplicate author", func(t *testing.T) {

		_, err := authorManager.Create(&author2)
		if err == nil {
			t.Fail()
		}
		if author2.ID != 0 {
			t.Fail()
		}

	})

	t.Run("List authors", func(t *testing.T) {
		authors, err := authorManager.List()
		if len(authors) == 0 {
			t.Fail()
		}
		if err != nil {
			t.Fail()
		}

	})

	t.Run("Find author by name", func(t *testing.T) {
		authorBD, err := authorManager.FindByName("Test Author")
		if authorBD.ID == 0 {
			t.Log("Author not found")
			t.Fail()
		}
		if err != nil {
			t.Log("Error: ", err.Error())
			t.Fail()
		}
		if authorBD.Name != "Test Author" {
			t.Log("Name not found.")
			t.Fail()
		}

	})

	t.Run("Update author", func(t *testing.T) {
		authorUp, err := authorManager.FindByName("Test Author")
		if err != nil {
			t.Fail()
		}
		fmt.Println("authorUp")
		fmt.Println(authorUp)
		authorUp.Name = "Test Author Updated"
		_, err = authorManager.Update(authorUp)
		if err != nil {
			t.Fail()
		}
		authorCheck, err := authorManager.FindByName(authorUp.Name)
		if authorCheck.ID != authorUp.ID {
			t.Fail()
		}
		if authorCheck.Name != authorUp.Name {
			t.Fail()
		}

	})

	t.Run("Delete author", func(t *testing.T) {
		err := authorManager.Delete(author.ID)
		if err != nil {
			t.Fail()
		}

	})

}
