package courses

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Filter struct {
	Keyword     string
	Category    string
	Status      string
	Sorting     string
	CourseTitle string
}

type CoursesModel struct {
	Userid     int
	DataAccess int
}

var Coursemodels CoursesModel

// Create function
type TblCourse struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title       string    `gorm:"type:character varying"`
	Description string    `gorm:"type:character varying"`
	ImageName   string    `gorm:"type:character varying"`
	ImagePath   string    `gorm:"type:character varying"`
	CategoryId  int       `gorm:"type:integer"`
	Status      int       `gorm:"type:integer"`
	TenantId    int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:integer"`
	ModifiedOn  time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:integer;DEFAULT:0"`
}

// List function
type TblCourses struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title            string    `gorm:"type:character varying"`
	Description      string    `gorm:"type:character varying"`
	ImageName        string    `gorm:"type:character varying"`
	ImagePath        string    `gorm:"type:character varying"`
	CategoryId       int       `gorm:"type:integer"`
	Status           int       `gorm:"type:integer"`
	TenantId         int       `gorm:"type:integer"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn        time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer;DEFAULT:0"`
	DateString       string    `gorm:"-"`
	CreatedTime      string    `gorm:"-"`
	ProfileImagePath string    `gorm:"column:profile_image_path"`
	FirstName        string    `gorm:"column:first_name"`
	LastName         string    `gorm:"column:last_name"`
	UserName         string    `gorm:"column:username"`
	NameString       string    `gorm:"-"`
}

func (Coursesmodels CoursesModel) ListCourses(limit, offset int, filter Filter, tenantid int, DB *gorm.DB) (courselist []TblCourses, count int64, err error) {

	query := DB.Table("tbl_courses").Select("tbl_courses.*,tbl_users.profile_image_path,tbl_users.first_name,tbl_users.last_name,tbl_users.username").Joins("inner join tbl_users on tbl_courses.created_by=tbl_users.id").Where("tbl_courses.is_deleted=0 and tbl_courses.tenant_id=?", tenantid)

	if filter.Sorting == "lastUpdated" {

		query = query.Order("modified_on desc")

	} else if filter.Sorting == "createdDate" {

		query = query.Order("created_on asc")

	} else if filter.Sorting == "asc" {

		query = query.Order("title asc")

	} else if filter.Sorting == "desc" {

		query = query.Order("title desc")

	} else {

		query = query.Order("tbl_courses.id desc")

	}

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(tbl_courses.title)) like LOWER(TRIM(?))", "%"+filter.Keyword+"%")

	}

	if filter.CourseTitle != "" {

		query = query.Where("LOWER(TRIM(tbl_courses.title)) like LOWER(TRIM(?))", "%"+filter.CourseTitle+"%")

	}

	if filter.Status != "" {

		query = query.Where("tbl_courses.status=?", filter.Status)

	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Find(&courselist)

		return courselist, count, nil

	}

	query.Find(&courselist).Count(&count)
	if query.Error != nil {

		return []TblCourses{}, 0, query.Error
	}

	return courselist, count, nil
}

func (Coursemodels CoursesModel) CreateCourse(course TblCourse, DB *gorm.DB) error {

	if err := DB.Table("tbl_courses").Create(&course).Error; err != nil {

		return err
	}
	fmt.Println("hello world courses created")

	return nil

}

func (Coursemodels CoursesModel) EditCourse(id, tenantid int, DB *gorm.DB) (courselist TblCourse, err error) {

	if err := DB.Table("tbl_courses").Where("id=? and tenant_id=? and is_deleted=0", id, tenantid).First(&courselist).Error; err != nil {

		return TblCourse{}, err
	}

	return courselist, nil

}

func (Coursemodels CoursesModel) DeleteCourse(id, tenantid, deletedby int, deletedon time.Time, DB *gorm.DB) error {

	if err := DB.Table("tbl_courses").Where("id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": deletedby, "deleted_on": deletedon}).Error; err != nil {

		return err
	}

	return nil
}

func (Coursemodels CoursesModel) MultiSelectCourseDelete(course *TblCourse, id []int, DB *gorm.DB) error {

	if err := DB.Table("tbl_courses").Where("id in (?) and tenant_id=?", id, course.TenantId).UpdateColumns(map[string]interface{}{"is_deleted": course.IsDeleted, "deleted_on": course.DeletedOn, "deleted_by": course.DeletedBy}).Error; err != nil {

		return err
	}

	return nil

}
