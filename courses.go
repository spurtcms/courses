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

func (courses *Courses) CreateCourse(create TblCourses) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Create := TblCourses{
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
