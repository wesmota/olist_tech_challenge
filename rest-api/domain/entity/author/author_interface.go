package author

// CRUD Inteface for author
type Reader interface {
	Find(id int) (*Author, error)
	FindByName(name string) (*Author, error)
	List() ([]*Author, error)
}

type Writer interface {
	Create(author *Author) (*Author, error)
	Update(author *Author) (*Author, error)
	Delete(id int64) error
}

type Repository interface {
	Reader
	Writer
}
