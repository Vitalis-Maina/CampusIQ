package main

import "github.com/Vitalis-Maina/internal/data"

type UMSData struct {
	Students     []data.Student
	Courses      []data.Courses
	Departments  []data.Department
	Units        []data.Units
	Lecturers    []data.Lecturers
	StudentUnits []data.StudentUnits
}
