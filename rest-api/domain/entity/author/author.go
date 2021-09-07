package author

type Author struct {
	ID   int64  `json:"ID" gorm:"column:id;type int;primary_key;not null;"`
	Name string `json:"Name" gorm:"column:name;type text;;not null;"`
}

func (Author) TableName() string {
	return "author"
}
