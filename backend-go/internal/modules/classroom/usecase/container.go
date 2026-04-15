package usecase

import (
	"go-enterprise-blueprint/internal/modules/classroom/usecase/addstudent"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/assignmentor"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/createclass"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/getclassdetail"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/listclasses"
	"go-enterprise-blueprint/internal/modules/classroom/usecase/updateaccess"
)

type Container struct {
	createClass    createclass.UseCase
	listClasses    listclasses.UseCase
	getClassDetail getclassdetail.UseCase
	addStudent     addstudent.UseCase
	updateAccess   updateaccess.UseCase
	assignMentor   assignmentor.UseCase
}

func NewContainer(createClass createclass.UseCase, listClasses listclasses.UseCase, getClassDetail getclassdetail.UseCase, addStudent addstudent.UseCase, updateAccess updateaccess.UseCase, assignMentor assignmentor.UseCase) *Container {
	return &Container{createClass: createClass, listClasses: listClasses, getClassDetail: getClassDetail, addStudent: addStudent, updateAccess: updateAccess, assignMentor: assignMentor}
}

func (c *Container) CreateClass() createclass.UseCase       { return c.createClass }
func (c *Container) ListClasses() listclasses.UseCase       { return c.listClasses }
func (c *Container) GetClassDetail() getclassdetail.UseCase { return c.getClassDetail }
func (c *Container) AddStudent() addstudent.UseCase         { return c.addStudent }
func (c *Container) UpdateAccess() updateaccess.UseCase     { return c.updateAccess }
func (c *Container) AssignMentor() assignmentor.UseCase     { return c.assignMentor }
