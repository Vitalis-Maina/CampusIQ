package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	router := http.NewServeMux()

	//student routes

	router.HandleFunc("/v1/students/", app.showStudent)

	//Department routes

	router.HandleFunc("/v1/departments/", app.showDepartments)

	//course routes

	router.HandleFunc("/v1/courses/", app.showCourses)

	//lecturer routes
	router.HandleFunc("/v1/lecturers/", app.showLecturer)

	//unit routes
	router.HandleFunc("/v1/units/", app.showUnits)

	//studentUnit routes

	router.HandleFunc("/v1/studentunits/", app.showStudelntUnits)

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	router.Handle("/", fileserver)

	return router
}
