package author

type manager struct {
	repo Repository
}

// Creates a new author manager
func NewManager(repo Repository) *manager {
	return &manager{repo}
}

//Create
func (m *manager) Create(author *Author) (*Author, error) {
	return m.repo.Create(author)
}

//Update
func (m *manager) Update(author *Author) (*Author, error) {
	return m.repo.Update(author)
}

//Delete
func (m *manager) Delete(id int64) error {
	return m.repo.Delete(id)
}

//Get
func (m *manager) Find(id int) (*Author, error) {
	return m.repo.Find(id)
}

//GetByName
func (m *manager) FindByName(name string) (*Author, error) {
	return m.repo.FindByName(name)
}

//List
func (m *manager) List() ([]*Author, error) {
	return m.repo.List()
}
