package model

type Employee struct {
	ID           int    `json:"id" bson:"id"`
	Name         string `json:"name" bson:"name"`
	Sex          string `json:"sex" bson:"sex"`
	Age          int    `json:"age" bson:"age"`
	Salary       int    `json:"salary" bson:"salary"`
	DepartmentID int    `json:"department_id" bson:"department_id"`
}
