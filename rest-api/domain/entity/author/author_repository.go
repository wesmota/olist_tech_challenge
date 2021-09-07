package author

import (
	"log"

	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

func InitAuthorRepository(db *gorm.DB) *authorRepository {
	return &authorRepository{
		db: db,
	}
}

//Creates a new author
func (a *authorRepository) Create(author *Author) (*Author, error) {
	err := a.db.Create(author).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return author, nil
}

//Deletes an author
func (a *authorRepository) Delete(id int64) error {
	err := a.db.Delete(Author{}, "id = ?", id).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//Updates an author
func (a *authorRepository) Update(author *Author) (*Author, error) {
	err := a.db.Save(author).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return author, nil
}

//Finds an author by id
func (a *authorRepository) Find(id int) (*Author, error) {
	author := &Author{}
	err := a.db.First(author, "id = ?", id).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return author, nil
}

//Finds an author by name
func (a *authorRepository) FindByName(name string) (*Author, error) {
	a.db.AutoMigrate(&Author{})
	author := &Author{}
	err := a.db.Find(author, "name = ?", name).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return author, nil
}

//List all authors
func (a *authorRepository) List() ([]*Author, error) {
	authors := []*Author{}
	err := a.db.Find(&authors).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return authors, nil
}
