package courses

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Filter struct {
	Keyword       string
	Category      string
	Status        string
	FromDate      string
	ToDate        string
	FirstName     string
	MemberProfile bool
	Level         string
	OrderId       int
	TransactionId string
	Gateway       string
}

type CoursesModel struct {
	Userid     int
	DataAccess int
}

var Coursemodels CoursesModel

type TblCourses struct {
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

func (Coursemodels CoursesModel) CreateCourse(course TblCourses, DB *gorm.DB) error {

	if err := DB.Table("tbl_courses").Create(&course).Error; err != nil {

		return err
	}
	fmt.Println("hello world courses created")

	return nil

}
