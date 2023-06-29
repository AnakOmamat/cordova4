package student

type Student struct {
	ID   int    `gorm:"primary_key;not null"`
	Name string `gorm:"size:255;not null"`
}
