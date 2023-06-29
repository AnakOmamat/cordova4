package student

import "gorm.io/gorm"

type Repository interface {
	Save(student Student) (Student, error)
	FindByID(ID int) (Student, error)
	Update(student Student) (Student, error)
	FindAll() ([]Student, error)
	Delete(student Student) (Student, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(student Student) (Student, error) {
	err := r.db.Create(&student).Error
	if err != nil {
		return student, err
	}

	return student, nil

}

func (r *repository) FindByID(ID int) (Student, error) {

	var student Student

	err := r.db.Where("id = ?", ID).Find(&student).Error
	if err != nil {
		return student, err
	}
	return student, nil

}

func (r *repository) Update(student Student) (Student, error) {
	err := r.db.Save(&student).Error
	if err != nil {
		return student, err
	}
	return student, nil

}

func (r *repository) FindAll() ([]Student, error) {
	var student []Student

	err := r.db.Find(&student).Error
	if err != nil {
		return student, err
	}
	return student, nil
}

func (r *repository) Delete(student Student) (Student, error) {

	err := r.db.Delete(&student).Error
	if err != nil {
		return student, err
	}
	return student, nil
}
