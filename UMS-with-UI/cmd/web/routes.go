package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	// Home route
	// router.HandleFunc("/", app.home).Methods("GET")

	// Student routes
	studentRoutes := router.PathPrefix("/v1/students").Subrouter()
	studentRoutes.HandleFunc("/", app.showStudent).Methods("GET")

	studentRoutes.HandleFunc("/create", app.insertStudent)
	studentRoutes.HandleFunc("/update", app.updateStudent)
	studentRoutes.HandleFunc("/delete", app.deleteStudent)

	// Department routes
	departmentRoutes := router.PathPrefix("/v1/departments").Subrouter()

	departmentRoutes.HandleFunc("/", app.showDepartments).Methods("GET")
	departmentRoutes.HandleFunc("/create", app.insertDepartment)
	departmentRoutes.HandleFunc("/update", app.updateDepartment)
	departmentRoutes.HandleFunc("/delete", app.deleteDepartment)

	// Course routes
	courseRoutes := router.PathPrefix("/v1/courses").Subrouter()
	courseRoutes.HandleFunc("/", app.showCourses).Methods("GET")

	courseRoutes.HandleFunc("/create", app.insertCourse)
	courseRoutes.HandleFunc("/update", app.updateCourse)
	courseRoutes.HandleFunc("/delete", app.deleteCourse)

	// Lecturer routes
	lecturerRoutes := router.PathPrefix("/v1/lecturers").Subrouter()

	lecturerRoutes.HandleFunc("/", app.showLecturer).Methods("GET")
	lecturerRoutes.HandleFunc("/create", app.insertLecturer)
	lecturerRoutes.HandleFunc("/update", app.updateLecturer)
	lecturerRoutes.HandleFunc("/delete", app.deleteLecturer)

	// Unit routes
	unitRoutes := router.PathPrefix("/v1/units").Subrouter()

	unitRoutes.HandleFunc("/", app.showUnits).Methods("GET")
	unitRoutes.HandleFunc("/create", app.insertUnit)
	unitRoutes.HandleFunc("/update", app.updateUnit)
	unitRoutes.HandleFunc("/delete", app.deleteUnit)

	// StudentUnit routes
	studentUnitRoutes := router.PathPrefix("/v1/studentunits").Subrouter()
	studentUnitRoutes.HandleFunc("/", app.showStudentUnits).Methods("GET")

	// studentUnitRoutes.HandleFunc("/create", app.insertStudentUnit)
	// studentUnitRoutes.HandleFunc("/update", app.updateStudentUnit)
	studentUnitRoutes.HandleFunc("/delete", app.deleteStudentUnit)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/", http.StripPrefix("", fileServer))

	return router
}
