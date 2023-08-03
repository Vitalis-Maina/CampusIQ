package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()

	//student routes

	router.HandleFunc("/v1/students/", app.showStudent)
	router.HandleFunc("/v1/students/create/", app.insertStudent)
	router.HandleFunc("/v1/students/update/", app.updateStudent)
	router.HandleFunc("/v1/students/delete/", app.deleteStudent)

	//Department routes

	router.HandleFunc("/v1/departments/", app.showDepartments)
	router.HandleFunc("/v1/departments/create/", app.showDepartments)
	router.HandleFunc("/v1/departments/update/", app.showDepartments)
	router.HandleFunc("/v1/departments/delete/", app.showDepartments)

	//course routes

	router.HandleFunc("/v1/courses/", app.showCourses)
	router.HandleFunc("/v1/courses/create/", app.showCourses)
	router.HandleFunc("/v1/courses/update/", app.showCourses)
	router.HandleFunc("/v1/courses/delete/", app.showCourses)

	//lecturer routes
	router.HandleFunc("/v1/lecturers/", app.showLecturer)
	router.HandleFunc("/v1/lecturers/create/", app.showLecturer)
	router.HandleFunc("/v1/lecturers/update/", app.showLecturer)
	router.HandleFunc("/v1/lecturers/delete/", app.showLecturer)

	//unit routes
	router.HandleFunc("/v1/units/", app.showUnits)
	router.HandleFunc("/v1/units/create/", app.showUnits)
	router.HandleFunc("/v1/units/update/", app.showUnits)
	router.HandleFunc("/v1/units/delete/", app.showUnits)

	//studentUnit routes

	router.HandleFunc("/v1/studentunits/", app.showStudentUnits)
	router.HandleFunc("/v1/studentunits/create/", app.insertStudentUnit)
	// router.HandleFunc("/v1/studentunits/update", app.showStudelntUnits)
	// router.HandleFunc("/v1/studentunits/delete/", app.showStudelntUnits)

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	router.Handle("/", fileserver)

	return router
}
