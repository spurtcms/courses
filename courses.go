package courses

import (
	"fmt"
	"time"

	"github.com/spurtcms/auth/migration"
)

func CoursesSetup(config Config) *Courses {

	migration.AutoMigration(config.DB, config.DataBaseType)

	fmt.Println("hello")

	return &Courses{
		AuthEnable:       config.AuthEnable,
		Permissions:      config.Permissions,
		PermissionEnable: config.PermissionEnable,
		Auth:             config.Auth,
		DB:               config.DB,
	}

}

func (courses *Courses) CoursesList(limit, offset int, filter Filter, tenantid string) (list []TblCourses, Count int64, err error) {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return []TblCourses{}, 0, Autherr

	}

	if filter.Status == "Draft" {

		filter.Status = "0"

	} else if filter.Status == "Published" {

		filter.Status = "1"

	} else if filter.Status == "Unpublished" {

		filter.Status = "2"

	}

	courseslist, _, _ := Coursemodels.ListCourses(limit, offset, filter, tenantid, courses.DB)

	_, count, err := Coursemodels.ListCourses(0, 0, filter, tenantid, courses.DB)
	if err != nil {

		return []TblCourses{}, 0, err
	}

	return courseslist, count, nil

}

func (courses *Courses) CreateCourse(create TblCourse) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Create := TblCourse{
		Title:       create.Title,
		Description: create.Description,
		CategoryId:  create.CategoryId,
		ImagePath:   create.ImagePath,
		ImageName:   create.ImageName,
		CreatedBy:   create.CreatedBy,
		CreatedOn:   createdon,
		Status:      create.Status,
		IsDeleted:   create.IsDeleted,
		TenantId:    create.TenantId,
	}

	err := Coursemodels.CreateCourse(Create, courses.DB)

	if err != nil {

		return err
	}

	return nil

}

func (courses *Courses) EditCourses(id int, tenantid string) (courselist TblCourse, err error) {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return TblCourse{}, Autherr
	}

	courselist, err = Coursemodels.EditCourse(id, tenantid, courses.DB)
	if err != nil {
		fmt.Println(err)
	}

	return courselist, nil

}

func (courses *Courses) DeleteCourses(id, userid int, tenantid string) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	deletedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	deletedby := userid

	err := Coursemodels.DeleteCourse(id, tenantid, deletedby, deletedon, courses.DB)

	if err != nil {

		return err
	}

	return nil

}

func (courses *Courses) MultiSelectDeleteCourse(courseids []int, modifiedby int, tenantid string) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	var Course TblCourse

	Course.DeletedBy = modifiedby

	Course.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Course.IsDeleted = 1

	Course.TenantId = tenantid

	err := Coursemodels.MultiSelectCourseDelete(&Course, courseids, courses.DB)
	if err != nil {

		return err

	}
	return nil
}
