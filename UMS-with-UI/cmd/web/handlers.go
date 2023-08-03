package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"strconv"

	"github.com/Vitalis-Maina/internal/data"
)

type Student struct {
	Name         string `json:"student_name"`
	CourseID     int64  `json:"course_id"`
	DepartmentID int64  `json:"department_id"`
}

type Department struct {
	ID   int64  `json:"department_id"`
	Name string `json:"department_name"`
}

type Courses struct {
	ID           int64  `json:"course_id"`
	Name         string `json:"course_name"`
	DepartmentID int64  `json:"department_id"`
}
type Lecturers struct {
	ID           int64  `json:"lecturer_id"`
	Name         string `json:"lecturer_name"`
	DepartmentID int64  `json:"department_id"`
	CourseID     int64  `json:"course_id"`
}

type Units struct {
	ID         int64  `json:"unit_id"`
	Name       string `json:"unit_name"`
	CourseID   int64  `json:"course_id"`
	LecturerID int64  `json:"lecturer_id"`
}
type StudentUnits struct {
	StudentID  int    `json:"student_id"`
	UnitID     int    `json:"unit_id"`
	Name       string `json:"name"`
	UnitName   string `json:"unit_name"`
	CourseID   int    `json:"course_id"`
	CourseName string `json:"course_name"`
	Category   string `json:"category"`
}

func (app *application) showStudent(w http.ResponseWriter, r *http.Request) {

	s, err := app.models.GetStudents()

	if err != nil {
		http.Error(w, "failed to fetch data", http.StatusInternalServerError)
	}

	data := UMSData{Students: s}

	tmpl, err := template.ParseFiles("./ui/static/students.html")
	if err != nil {
		http.Error(w, "Failed to parse files", http.StatusNotFound)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Fprintln(w, "failed to execute template", err)
	}
}
func (app *application) showDepartments(w http.ResponseWriter, r *http.Request) {
	d, err := app.models.GetDepartments()

	if err != nil {
		http.Error(w, "error while fetching departments", http.StatusInternalServerError)
	}
	data := UMSData{Departments: d}

	tmpl, err := template.ParseFiles("./ui/static/departments.html")
	if err != nil {
		http.Error(w, "failed to parse files", http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		fmt.Fprintln(w, "failed to execute template", err)
	}
}

func (app *application) showCourses(w http.ResponseWriter, r *http.Request) {
	c, err := app.models.GetCourses()
	if err != nil {
		http.Error(w, "failed to fetch courses", http.StatusInternalServerError)
	}
	data := UMSData{Courses: c}

	tmpl, err := template.ParseFiles("./ui/static/courses.html")
	if err != nil {
		http.Error(w, "failed to parse files", http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "failed to exectute template", http.StatusInternalServerError)
	}
}
func (app *application) showLecturer(w http.ResponseWriter, r *http.Request) {
	lecturers, err := app.models.GetLecturers()
	if err != nil {
		http.Error(w, "Failed to fetch lecturers", http.StatusInternalServerError)
	}

	data := UMSData{Lecturers: lecturers}

	tmpl, err := template.ParseFiles("./ui/static/lecturers.html")
	if err != nil {
		http.Error(w, "failed to parse files", http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "failed to execute template", http.StatusInternalServerError)
	}
}

func (app *application) showUnits(w http.ResponseWriter, r *http.Request) {
	units, err := app.models.GetUnits()

	if err != nil {
		http.Error(w, "failed to fetch units", http.StatusInternalServerError)
	}
	data := UMSData{Units: units}

	tmpl, err := template.ParseFiles("./ui/static/units.html")
	if err != nil {
		http.Error(w, "failed to parse files", http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "failed to execute template", http.StatusInternalServerError)
	}
}

func (app *application) showStudelntUnits(w http.ResponseWriter, r *http.Request) {
	studentunits, err := app.models.GetStudentUnits()
	if err != nil {
		http.Error(w, "failed to fetch student units", http.StatusInternalServerError)
	}
	data := UMSData{StudentUnits: studentunits}

	tmpl, err := template.ParseFiles("./ui/static/studentUnits.html")
	if err != nil {
		http.Error(w, "failed to parse files", http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "failed to execute template", http.StatusInternalServerError)
	}
}
func (app *application) showStudentUnits(w http.ResponseWriter, r *http.Request) {
	studentunits, err := app.models.GetStudentUnits()
	if err != nil {
		http.Error(w, "failed to fetch student units", http.StatusInternalServerError)
	}
	data := UMSData{StudentUnits: studentunits}

	tmpl, err := template.ParseFiles("./ui/static/studentUnits.html")
	if err != nil {
		http.Error(w, "failed to parse files", http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "failed to execute template", http.StatusInternalServerError)
	}
}

func (app *application) deleteStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/deleteStudent.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		student_name := r.Form.Get("student_name")
		if student_name == "" {
			http.Error(w, "cant submit empty fields", http.StatusBadRequest)
			return
		}
		app.models.DeleteStudent(student_name)
		fmt.Fprintf(w, "student %s has been successfully deleted", student_name)

	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}
func (app *application) deleteDepartment(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/deleteDepartment.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		r.ParseForm()

		department_id := r.Form.Get("department_id")
		if department_id == "" {
			http.Error(w, "cant submit empty fields", http.StatusBadRequest)
			return
		}
		departmentID, err := strconv.Atoi(department_id)
		if err != nil {
			http.Error(w, "invalid department id", http.StatusBadRequest)
		}

		err = app.models.DeleteDepartment(departmentID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "department with id: %d has been successfully deleted", departmentID)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (app *application) deleteCourse(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/deleteCourse.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		r.ParseForm()

		course_id := r.Form.Get("course_id")
		if course_id == "" {
			http.Error(w, "cant submit empty fields", http.StatusBadRequest)
			return
		}
		courseID, err := strconv.Atoi(course_id)
		if err != nil {
			http.Error(w, "invalid course if", http.StatusBadRequest)
		}
		err = app.models.DeleteCourse(courseID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "course with id: %d has been successffully deleted", courseID)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}
func (app *application) deleteLecturer(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/deleteLecturer.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		r.ParseForm()

		lecturer_id := r.Form.Get("lecturer_id")

		if lecturer_id == "" {
			http.Error(w, "cannot submit empty field", http.StatusBadRequest)
			return
		}
		lecturerID, err := strconv.Atoi(lecturer_id)
		if err != nil {
			http.Error(w, "invalid lecturer id", http.StatusBadRequest)
		}
		err = app.models.DeleteLecturer(lecturerID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "lecturer with id:%d has been successfully deleted", lecturerID)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}
func (app *application) deleteUnit(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/deleteUnit.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		r.ParseForm()

		unit_id := r.Form.Get("unit_id")

		if unit_id == "" {
			http.Error(w, "you cannot submit empty field", http.StatusBadRequest)
		}
		unitID, err := strconv.Atoi(unit_id)
		if err != nil {
			http.Error(w, "invalid unit id", http.StatusBadRequest)
		}
		err = app.models.DeleteUnit(unitID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "unit with id:%d has been successfully deleted ", unitID)
	}

}
func (app *application) deleteStudentUnit(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/deleteStudentUnit.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}

	} else if r.Method == "POST" {
		r.ParseForm()

		student_id := r.Form.Get("student_id")
		unit_id := r.Form.Get("unit_id")

		if student_id == "" || unit_id == "" {
			http.Error(w, "you cannot submit empty fields", http.StatusBadRequest)
			return
		}

		studentID, err := strconv.Atoi(student_id)
		if err != nil {
			http.Error(w, "invalid student id", http.StatusBadRequest)
		}

		unitID, err := strconv.Atoi(unit_id)
		if err != nil {
			http.Error(w, "invalid unit id", http.StatusBadRequest)
		}

		err = app.models.DeleteStudentUnit(studentID, unitID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Student unit %d - %d has been successfully deleted", studentID, unitID)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func (app *application) insertStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/createStudent.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		studentName := r.Form.Get("name")
		courseID := r.Form.Get("course_id")
		departmentID := r.Form.Get("department_id")

		if studentName == "" || courseID == "" || departmentID == "" {
			http.Error(w, "cant submit empty fields", http.StatusBadRequest)
			return
		}

		courseIDInt, err := strconv.ParseInt(courseID, 10, 64)
		if err != nil {
			http.Error(w, "Invalid course ID", http.StatusBadRequest)
			return
		}
		departmentIDInt, err := strconv.ParseInt(departmentID, 10, 64)
		if err != nil {
			http.Error(w, "Invalid department ID", http.StatusBadRequest)
			return
		}

		// Create a Student struct with the form data
		student := Student{
			Name:         studentName,
			CourseID:     courseIDInt,
			DepartmentID: departmentIDInt,
		}

		// Perform the necessary operations, such as inserting the student into the database
		err = app.models.InsertStudent(data.Student(student))
		if err != nil {
			http.Error(w, "Failed to insert student", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "Student inserted successfully")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (app *application) insertDepartment(w http.ResponseWriter, r *http.Request) {
	var department Department
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/createDepartment.html")
		if err != nil {
			http.Error(w, "failed to parse files", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {

		r.ParseForm()

		department_id := r.Form.Get("department_id")
		department_name := r.Form.Get("department_name")

		if department_id == "" || department_name == "" {
			http.Error(w, "cant submit empty fields", http.StatusBadRequest)
			return
		}

		departmentID, err := strconv.ParseInt(department_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid department id", http.StatusBadRequest)
		}

		department = Department{
			ID:   departmentID,
			Name: department_name,
		}
		err = app.models.InsertDepartment(data.Department(department))
		if err != nil {
			http.Error(w, "Failed to insert department", http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, "department inserted successfully")

	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func (app *application) insertCourse(w http.ResponseWriter, r *http.Request) {
	var course Courses
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/createCourse.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		r.ParseForm()

		course_id := r.Form.Get("course_id")
		course_name := r.Form.Get("course_name")
		department_id := r.Form.Get("department_id")

		if course_id == "" || course_name == "" || department_id == "" {
			http.Error(w, "cant submit empty fields id ", http.StatusBadRequest)
			return
		}

		courseID, err := strconv.ParseInt(course_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid course id", http.StatusBadRequest)
		}
		departmentID, err := strconv.ParseInt(department_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid department id", http.StatusBadRequest)
		}
		course = Courses{
			ID:           courseID,
			Name:         course_name,
			DepartmentID: departmentID,
		}
		err = app.models.InsertCourse(data.Courses(course))
		if err != nil {
			http.Error(w, "Failed to insert course", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "Course inserted successfully")
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func (app *application) insertLecturer(w http.ResponseWriter, r *http.Request) {
	var lecturer Lecturers
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/createLecturer.html")
		if err != nil {
			http.Error(w, " failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		r.ParseForm()

		lecturer_id := r.Form.Get("lecturer_id")
		lecturer_name := r.Form.Get("lecturer_name")
		department_id := r.Form.Get("department_id")
		course_id := r.Form.Get("course_id")

		//validation checks
		if lecturer_id == "" || lecturer_name == "" || department_id == "" || course_id == "" {
			http.Error(w, "cant be empty", http.StatusBadRequest)
			return
		}
		lecturerID, err := strconv.ParseInt(lecturer_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
		}
		courseId, err := strconv.ParseInt(course_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
		}
		departmentID, err := strconv.ParseInt(department_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
		}

		lecturer = Lecturers{
			ID:           lecturerID,
			Name:         lecturer_name,
			CourseID:     courseId,
			DepartmentID: departmentID,
		}
		err = app.models.InsertLecturer(data.Lecturers(lecturer))
		if err != nil {
			http.Error(w, "Failed to insert lecturer", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "Lecturer inserted successfully")
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func (app *application) insertUnit(w http.ResponseWriter, r *http.Request) {
	var unit Units
	if r.Method == "GET" {

		tmpl, err := template.ParseFiles("./ui/static/createUnit.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}

	} else if r.Method == "POST" {
		r.ParseForm()

		unit_id := r.Form.Get("unit_id")
		unit_name := r.Form.Get("unit_name")
		course_id := r.Form.Get("course_id")
		lecturer_id := r.Form.Get("lecturer_id")

		if unit_id == "" || unit_name == "" || course_id == "" || lecturer_id == "" {
			http.Error(w, "cant be empty", http.StatusBadRequest)
			return
		}
		unitID, err := strconv.ParseInt(unit_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid unit id", http.StatusBadRequest)
		}
		courseID, err := strconv.ParseInt(course_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid course id", http.StatusBadRequest)
		}
		lecturerId, err := strconv.ParseInt(lecturer_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid lecturer id", http.StatusBadRequest)
		}

		unit = Units{
			ID:         unitID,
			Name:       unit_name,
			CourseID:   courseID,
			LecturerID: lecturerId,
		}

		err = app.models.InsertUnit(data.Units(unit))
		if err != nil {
			http.Error(w, "Failed to insert unit", http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "Unit inserted successfully")
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

/*
	func (app *application) insertStudentUnit(w http.ResponseWriter, r *http.Request) {
		var studentUnit StudentUnits
		if r.Method == "GET" {
			tmpl, err := template.ParseFiles("./ui/static/createStudentUnit.html")
			if err != nil {
				http.Error(w, "failed to load template", http.StatusInternalServerError)
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				http.Error(w, "failed to render template", http.StatusInternalServerError)
				return
			}
		} else if r.Method == "POST" {

			r.ParseForm()

			student_id := r.Form.Get("student_id")
			unit_id := r.Form.Get("unit_id")

			if student_id == "" || unit_id == "" {
				http.Error(w, "cannot be empty", http.StatusBadRequest)
				return
			}
			studentID, err := strconv.ParseInt(student_id, 10, 64)
			if err != nil {
				http.Error(w, "invalid student id", http.StatusBadRequest)
			}
			unitID, err := strconv.ParseInt(unit_id, 10, 64)
			if err != nil {
				http.Error(w, "invalid unit id", http.StatusBadRequest)
			}
			studentUnit = StudentUnits{
				StudentID: studentID,
				UnitID:    unitID,
			}
			err = app.models.InsertStudentUnit(data.StudentUnits(studentUnit))
			if err != nil {
				http.Error(w, "Failed to insert student unit", http.StatusInternalServerError)
				return
			}

			fmt.Fprintln(w, "StudentUnit inserted successfully")
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

}

func (app *application) updateStudentUnit(w http.ResponseWriter, r *http.Request) {

	var studentUnit StudentUnits

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/updateStudentUnits.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		r.ParseForm()

		student_id := r.Form.Get("student_id")
		new_student_id := r.Form.Get("new_student_id")
		unit_id := r.Form.Get("unit_name")
		new_unit_id := r.Form.Get("new_unit_id")

		if student_id == "" || unit_id == "" || new_student_id == "" || new_unit_id == "" {
			http.Error(w, "cant submit empty fields", http.StatusBadRequest)
			return
		}
		studentID, err := strconv.Atoi(student_id)
		if err != nil {
			http.Error(w, "invalid student id", http.StatusBadRequest)
		}
		newStudentID, err := strconv.ParseInt(new_student_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid new student id", http.StatusBadRequest)
		}
		unitID, err := strconv.Atoi(unit_id)
		if err != nil {
			http.Error(w, "invalid unit id", http.StatusBadRequest)
		}
		newUnitID, err := strconv.ParseInt(new_unit_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid new unit id", http.StatusBadRequest)
		}

		studentUnit = StudentUnits{
			StudentID: newStudentID,
			UnitID:    newUnitID,
		}
		err = app.models.UpdateStudentUnits(studentID, unitID, data.StudentUnits(studentUnit))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "studentUnit  %d - %d updated successfully", studentID, unitID)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}
*/
func (app *application) updateStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/updateStudent.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		r.ParseForm()

		student_name := r.Form.Get("student_name")
		newName := r.Form.Get("new_name")
		courseID := r.Form.Get("course_id")
		departmentID := r.Form.Get("department_id")

		if student_name == "" || newName == "" || courseID == "" || departmentID == "" {
			http.Error(w, "cannot be empty", http.StatusBadRequest)
			return
		}
		course_id, err := strconv.ParseInt(courseID, 10, 64)
		if err != nil {
			http.Error(w, "invalid course id", http.StatusBadRequest)
		}
		department_id, err := strconv.ParseInt(departmentID, 10, 64)
		if err != nil {
			http.Error(w, "invalid department id", http.StatusBadRequest)
		}
		student := Student{
			Name:         newName,
			CourseID:     course_id,
			DepartmentID: department_id,
		}
		err = app.models.UpdateStudent(student_name, data.Student(student))
		if err != nil {
			http.Error(w, "failed to update unit", http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Unit with id %s updated successfully", student_name)

	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}
func (app *application) updateUnit(w http.ResponseWriter, r *http.Request) {

	var unit Units
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/updateUnit.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}

	} else if r.Method == "POST" {
		r.ParseForm()

		unit_id := r.Form.Get("unit_id")
		new_id := r.Form.Get("new_id")
		unit_name := r.Form.Get("unit_name")
		course_id := r.Form.Get("course_id")
		lecturer_id := r.Form.Get("lecturer_id")

		if unit_id == "" || new_id == "" || unit_name == "" || course_id == "" || lecturer_id == "" {
			http.Error(w, "cant submit empty fields", http.StatusBadRequest)
			return
		}

		unitID, err := strconv.Atoi(unit_id)
		if err != nil {
			http.Error(w, "invalid unit id", http.StatusBadRequest)
		}
		newID, err := strconv.ParseInt(new_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid new unit id", http.StatusBadRequest)
		}
		NewCourseID, err := strconv.ParseInt(course_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid course id", http.StatusBadRequest)
		}
		NewLecturerID, err := strconv.ParseInt(lecturer_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid lecturer id", http.StatusBadRequest)

		}

		unit = Units{
			ID:         newID,
			Name:       unit_name,
			CourseID:   NewCourseID,
			LecturerID: NewLecturerID,
		}
		err = app.models.UpdateUnit(unitID, data.Units(unit))
		if err != nil {
			http.Error(w, "failed to update unit", http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Unit with id %d updated successfully", unitID)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func (app *application) updateLecturer(w http.ResponseWriter, r *http.Request) {

	var lecturer Lecturers

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/updateLecturer.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		r.ParseForm()

		lecturer_id := r.Form.Get("lecturer_id")
		new_id := r.Form.Get("new_id")
		lecturer_name := r.Form.Get("lecturer_name")
		department_id := r.Form.Get("department_id")
		course_id := r.Form.Get("course_id")

		if lecturer_id == "" || new_id == "" || lecturer_name == "" || department_id == "" || course_id == "" {
			http.Error(w, "cant submit empty fields", http.StatusBadRequest)
			return
		}
		lecturerID, err := strconv.Atoi(lecturer_id)
		if err != nil {
			http.Error(w, "invalid lecturer id", http.StatusBadRequest)
		}

		newID, err := strconv.ParseInt(new_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid new lecturer id", http.StatusBadRequest)
		}
		departmentID, err := strconv.ParseInt(department_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid department id", http.StatusBadRequest)
		}
		courseID, err := strconv.ParseInt(course_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid course id", http.StatusBadRequest)
		}

		lecturer = Lecturers{
			ID:           newID,
			Name:         lecturer_name,
			DepartmentID: departmentID,
			CourseID:     courseID,
		}
		err = app.models.UpdateLecturer(lecturerID, data.Lecturers(lecturer))
		if err != nil {
			http.Error(w, "failed to update lecturer", http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "lecturer with id %d successfully updated", lecturerID)

	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func (app *application) updateDepartment(w http.ResponseWriter, r *http.Request) {

	var department Department

	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/updateDepartment.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		r.ParseForm()

		department_id := r.Form.Get("department_id")
		new_id := r.Form.Get("new_id")
		department_name := r.Form.Get("department_name")

		if department_id == "" || new_id == "" || department_name == "" {
			http.Error(w, "cant submit empty fields", http.StatusBadRequest)
			return
		}

		departmentID, err := strconv.Atoi(department_id)
		if err != nil {
			http.Error(w, "invalid department id", http.StatusBadRequest)
		}
		newID, err := strconv.ParseInt(new_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid department id", http.StatusBadRequest)
		}

		department = Department{
			ID:   newID,
			Name: department_name,
		}
		err = app.models.UpdateDepartment(departmentID, data.Department(department))
		if err != nil {
			http.Error(w, "failed to update department", http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Department with id %d updated successfully ", departmentID)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}

func (app *application) updateCourse(w http.ResponseWriter, r *http.Request) {

	var course Courses
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("./ui/static/updateCourse.html")
		if err != nil {
			http.Error(w, "failed to load template", http.StatusInternalServerError)
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "failed to render template", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		course_id := r.Form.Get("course_id")
		new_id := r.Form.Get("new_id")
		course_name := r.Form.Get("course_name")
		department_id := r.Form.Get("department_id")

		if course_id == "" || new_id == "" || course_name == "" || department_id == "" {
			http.Error(w, "you cannot submit empty fields", http.StatusBadRequest)
			return
		}
		courseID, err := strconv.Atoi(course_id)
		if err != nil {
			http.Error(w, "invalid course id", http.StatusBadRequest)
		}
		newID, err := strconv.ParseInt(new_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid new id", http.StatusBadRequest)
		}
		departmentID, err := strconv.ParseInt(department_id, 10, 64)
		if err != nil {
			http.Error(w, "invalid department id", http.StatusBadRequest)
		}

		course = Courses{
			ID:           newID,
			Name:         course_name,
			DepartmentID: departmentID,
		}
		err = app.models.UpdateCourse(courseID, data.Courses(course))
		if err != nil {
			http.Error(w, "failed to update course ", http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "course with id %d updated successfully", courseID)
	} else {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

}
