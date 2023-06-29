package handler

import (
	"cordova4/auth"
	"cordova4/helper"
	"cordova4/student"
	"cordova4/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type studentHandler struct {
	studentService student.Service
	authService    auth.Service
}

func NewStudentHandler(studentService student.Service, authService auth.Service) *studentHandler {
	return &studentHandler{studentService, authService}

}

func (h *studentHandler) CreateStudent(c *gin.Context) {
	var input student.RegisterStudentInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create student", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newStudent, err := h.studentService.RegisterStudent(input)
	if err != nil {
		response := helper.APIResponse("Failed to create student", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Success to create student", http.StatusOK, "success", student.StudentFormatter(newStudent))
	c.JSON(http.StatusOK, response)
}

func (h *studentHandler) GetStudents(c *gin.Context) {

	students, err := h.studentService.GetAllStudent()
	if err != nil {
		response := helper.APIResponse("Error to get student", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List of student", http.StatusOK, "success", student.FormatStudents(students))
	c.JSON(http.StatusOK, response)
}

func (h *studentHandler) GetStudent(c *gin.Context) {
	var input student.GetStudentDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of student", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	studentDetail, err := h.studentService.GetStudentByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get detail of student", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("student detail", http.StatusOK, "success", student.StudentFormatter(studentDetail))
	c.JSON(http.StatusOK, response)
}

func (h *studentHandler) DeleteStudent(c *gin.Context) {
	var inputID student.GetStudentDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete student", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	deleteStudent, err := h.studentService.DeleteStudent(inputID)
	if err != nil {
		response := helper.APIResponse("Failed to delete student", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to delete student", http.StatusOK, "success", student.StudentFormatter(deleteStudent))
	c.JSON(http.StatusOK, response)
}

func (h *studentHandler) UpdateStudent(c *gin.Context) {
	var inputID student.GetStudentDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update student", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData student.RegisterStudentInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update student", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedStudent, err := h.studentService.UpdateStudent(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update student", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to update student", http.StatusOK, "success", student.StudentFormatter(updatedStudent))
	c.JSON(http.StatusOK, response)
}
