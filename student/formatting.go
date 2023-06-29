package student

type StudentFormatter struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FormatStudent(category Student) StudentFormatter {
	studentFormatter := StudentFormatter{}
	studentFormatter.ID = category.ID
	studentFormatter.Name = category.Name

	return studentFormatter

}

func FormatStudents(category []Student) []StudentFormatter {
	studentsFormatter := []StudentFormatter{}

	for _, student := range category {
		studentFormatter := FormatStudent(student)
		studentsFormatter = append(studentsFormatter, studentFormatter)
	}

	return studentsFormatter

}
