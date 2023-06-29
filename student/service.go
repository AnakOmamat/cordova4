package student

type Service interface {
	RegisterStudent(input RegisterStudentInput) (Student, error)
	GetStudentByID(input GetStudentDetailInput) (Student, error)
	GetAllStudent() ([]Student, error)
	UpdateStudent(input GetStudentDetailInput, inputData RegisterStudentInput) (Student, error)
	DeleteStudent(input GetStudentDetailInput) (Student, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterStudent(input RegisterStudentInput) (Student, error) {
	student := Student{}
	student.Name = input.Name
	newStudent, err := s.repository.Save(student)
	if err != nil {
		return newStudent, err
	}
	return newStudent, nil
}

func (s *service) GetStudentByID(input GetStudentDetailInput) (Student, error) {
	student, err := s.repository.FindByID(input.ID)

	if err != nil {
		return student, err
	}
	return student, nil
}

func (s *service) GetAllStudent() ([]Student, error) {
	student, err := s.repository.FindAll()
	if err != nil {
		return student, err
	}
	return student, nil

}

func (s *service) UpdateStudent(input GetStudentDetailInput, inputData RegisterStudentInput) (Student, error) {
	student, err := s.repository.FindByID(input.ID)
	if err != nil {
		return student, err
	}

	student.Name = inputData.Name

	updatedStudent, err := s.repository.Update(student)
	if err != nil {
		return updatedStudent, err
	}

	return updatedStudent, nil
}

func (s *service) DeleteStudent(input GetStudentDetailInput) (Student, error) {
	student, err := s.repository.FindByID(input.ID)
	if err != nil {
		return student, err
	}

	deleteStudent, err := s.repository.Delete(student)
	if err != nil {
		return deleteStudent, err
	}

	return deleteStudent, nil
}
