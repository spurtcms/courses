package courses

import (
	"fmt"
	"time"

	categories "github.com/spurtcms/categories"
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
	CategoryId  string    `gorm:"type:character varying"`
	Status      int       `gorm:"type:integer"`
	TenantId    string    `gorm:"type:character varying"`
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
	CategoryId       string    `gorm:"type:character varying"`
	Status           int       `gorm:"type:integer"`
	TenantId         string    `gorm:"type:character varying"`
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
	Offer            string    `gorm:"column:offer"`
}

//TblCourseSettings

type TblCourseSettings struct {
	Id           int       `gorm:"primaryKey;auto_increment;type:serial"`
	CourseId     int       `gorm:"type:integer"`
	Certificate  int       `gorm:"type:integer"`
	Comments     int       `gorm:"type:integer"`
	Offer        string    `gorm:"type:character varying"`
	Visibility   string    `gorm:"type:character varying"`
	StartDate    string    `gorm:"type:character varying"`
	SignUpLimits int       `gorm:"type:integer"`
	Duration     string    `gorm:"type:character varying"`
	CreatedOn    time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy    int       `gorm:"type:integer"`
	IsDeleted    int       `gorm:"type:integer"`
	DeletedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"DEFAULT:NULL"`
	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy   int       `gorm:"DEFAULT:NULL"`
	TenantId     string    `gorm:"type:character varying"`
}

//Settings Page List Table

type TblSettingsPages struct {
	Id           int    `gorm:"primaryKey;auto_increment;type:serial"`
	Title        string `gorm:"type:character varying"`
	Description  string `gorm:"type:character varying"`
	ImageName    string `gorm:"type:character varying"`
	ImagePath    string `gorm:"type:character varying"`
	CategoryId   string `gorm:"type:character varying"`
	Status       int    `gorm:"type:integer"`
	Certificate  int    `gorm:"type:integer"`
	Comments     int    `gorm:"type:integer"`
	Offer        string `gorm:"type:character varying"`
	Visibility   string `gorm:"type:character varying"`
	StartDate    string `gorm:"type:character varying"`
	SignUpLimits int    `gorm:"type:integer"`
	Duration     string `gorm:"type:character varying"`
}

//Create Section

type TblSection struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	Title      string    `gorm:"type:character varying"`
	Content    string    `gorm:"type:character varying"`
	CourseId   int       `gorm:"type:integer"`
	TenantId   string    `gorm:"type:character varying"`
	OrderIndex int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy  int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:integer;DEFAULT:0"`
}

//Create Lesson fields

type TblLesson struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	CourseId   int       `gorm:"type:integer"`
	SectionId  int       `gorm:"type:integer"`
	Title      string    `gorm:"type:character varying"`
	Content    string    `gorm:"type:character varying"`
	EmbedLink  string    `gorm:"type:character varying"`
	FileName   string    `gorm:"type:character varying"`
	FilePath   string    `gorm:"type:character varying"`
	LessonType string    `gorm:"type:character varying"`
	OrderIndex int       `gorm:"type:integer"`
	TenantId   string    `gorm:"type:character varying"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy  int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:integer;DEFAULT:0"`
}

func (Coursesmodels CoursesModel) ListCourses(limit, offset int, filter Filter, tenantid string, DB *gorm.DB) (courselist []TblCourses, count int64, err error) {

	query := DB.Table("tbl_courses").Select("tbl_courses.*,tbl_course_settings.offer,tbl_users.profile_image_path,tbl_users.first_name,tbl_users.last_name,tbl_users.username").Joins("inner join tbl_users on tbl_courses.created_by=tbl_users.id").Joins("inner join tbl_course_settings on tbl_courses.id=tbl_course_settings.course_id").Where("tbl_courses.is_deleted=0 and tbl_courses.tenant_id=?", tenantid)

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

	createsettings := TblCourseSettings{
		CourseId:  course.Id,
		Comments:  0,
		Offer:     "free",
		CreatedBy: course.CreatedBy,
		CreatedOn: course.CreatedOn,
		IsDeleted: 0,
		TenantId:  course.TenantId,
	}

	if err := DB.Table("tbl_course_settings").Create(&createsettings).Error; err != nil {

		return err
	}

	return nil

}

func (Coursemodels CoursesModel) EditCourse(id int, tenantid string, DB *gorm.DB) (courselist TblCourse, err error) {

	if err := DB.Table("tbl_courses").Where("id=? and tenant_id=? and is_deleted=0", id, tenantid).First(&courselist).Error; err != nil {

		return TblCourse{}, err
	}

	return courselist, nil

}

func (Coursemodels CoursesModel) UpdateCourse(id int, tenantid string, update TblCourse, DB *gorm.DB) error {

	if update.ImagePath != "" {
		if err := DB.Table("tbl_courses").Where("id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"title": update.Title, "description": update.Description, "image_name": update.ImageName, "image_path": update.ImagePath, "category_id": update.CategoryId, "status": update.Status, "modified_on": update.ModifiedOn, "modified_by": update.ModifiedBy}).Error; err != nil {

			return err
		}
	} else {
		if err := DB.Table("tbl_courses").Where("id=? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"title": update.Title, "description": update.Description, "category_id": update.CategoryId, "status": update.Status, "modified_on": update.ModifiedOn, "modified_by": update.ModifiedBy}).Error; err != nil {

			return err
		}
	}

	return nil
}

func (Coursemodels CoursesModel) EditCourseSettings(id int, tenantid string, DB *gorm.DB) (coursesettinglist TblCourseSettings, err error) {

	if err := DB.Table("tbl_course_settings").Where("course_id=? and tenant_id=? and is_deleted=0", id, tenantid).First(&coursesettinglist).Error; err != nil {

		return TblCourseSettings{}, err
	}

	return coursesettinglist, nil

}

func (Coursemodels CoursesModel) UpdateCourseSettings(tenantid string, update TblCourseSettings, DB *gorm.DB) error {

	if err := DB.Table("tbl_course_settings").Where("course_id=? and tenant_id=?", update.CourseId, tenantid).UpdateColumns(map[string]interface{}{"certificate": update.Certificate, "comments": update.Comments, "offer": update.Offer, "visibility": update.Visibility, "start_date": update.StartDate, "sign_up_limits": update.SignUpLimits, "duration": update.Duration, "modified_on": update.ModifiedOn, "modified_by": update.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil
}

func (Coursemodels CoursesModel) DeleteCourse(id int, tenantid string, deletedby int, deletedon time.Time, DB *gorm.DB) error {

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

func (Coursemodels CoursesModel) GetCategoriseById(id []int, DB *gorm.DB, tenantid string) (category []categories.TblCategories, err error) {

	if err := DB.Table("tbl_categories").Where("id in (?) and tenant_id=?", id, tenantid).Order("id asc").Find(&category).Error; err != nil {

		return category, err
	}

	return category, nil

}

//create sections

func (Coursemodels CoursesModel) CreateSection(section TblSection, DB *gorm.DB) error {

	if err := DB.Table("tbl_sections").Create(&section).Error; err != nil {

		return err
	}

	return nil

}

//ListSections

func (Coursemodels CoursesModel) SectionList(id int, tenantid string, DB *gorm.DB) (section []TblSection, err error) {

	if err := DB.Table("tbl_sections").Where("course_id=? and tenant_id=? and is_deleted=0", id, tenantid).Order("tbl_sections.order_index asc").Find(&section).Error; err != nil {

		return []TblSection{}, err
	}

	return section, nil
}

//Edit Section

func (Coursemodels CoursesModel) EditSection(sectionid int, coursesid int, tenantid string, DB *gorm.DB) (section TblSection, err error) {

	if err := DB.Debug().Table("tbl_sections").Where("id=? and course_id=? and tenant_id=?", sectionid, coursesid, tenantid).First(&section).Error; err != nil {

		return TblSection{}, err
	}

	return section, nil

}

//Update sections

func (Coursemodels CoursesModel) UpdateSection(update TblSection, DB *gorm.DB) error {

	if err := DB.Table("tbl_sections").Where("id=? and course_id=? and tenant_id=?", update.Id, update.CourseId, update.TenantId).UpdateColumns(map[string]interface{}{"title": update.Title, "content": update.Content, "modified_on": update.ModifiedOn, "modified_by": update.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil

}

//Section Delete

func (Coursemodels CoursesModel) DeleteSection(sectionid int, coursesid int, tenantid string, deletedby int, deletedon time.Time, DB *gorm.DB) error {

	if err := DB.Table("tbl_sections").Where("id=? and course_id=? and tenant_id=?", sectionid, coursesid, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": deletedby, "deleted_on": deletedon}).Error; err != nil {

		return err
	}

	return nil
}

//create Lesson

func (Coursemodels CoursesModel) CreateLesson(lesson TblLesson, DB *gorm.DB) error {

	if err := DB.Table("tbl_lessons").Create(&lesson).Error; err != nil {

		return err
	}

	return nil

}

//ListLesson

func (Coursemodels CoursesModel) LessonList(id int, tenantid string, DB *gorm.DB) (lesson []TblLesson, err error) {

	if err := DB.Table("tbl_lessons").Where("course_id=? and tenant_id=? and is_deleted=0", id, tenantid).Order("tbl_lessons.order_index asc").Find(&lesson).Error; err != nil {

		return []TblLesson{}, err
	}

	return lesson, nil
}

//Edit Lesson

func (Coursemodels CoursesModel) EditLesson(lessonid int, coursesid int, tenantid string, DB *gorm.DB) (lesson TblLesson, err error) {

	if err := DB.Debug().Table("tbl_lessons").Where("id=? and course_id=? and tenant_id=?", lessonid, coursesid, tenantid).First(&lesson).Error; err != nil {

		return TblLesson{}, err
	}

	return lesson, nil

}

//Update Lesson

func (Coursemodels CoursesModel) UpdateLesson(update TblLesson, DB *gorm.DB) error {

	if err := DB.Table("tbl_lessons").Where("id=? and course_id=? and tenant_id=?", update.Id, update.CourseId, update.TenantId).UpdateColumns(map[string]interface{}{"title": update.Title, "content": update.Content,
		"embed_link": update.EmbedLink, "modified_on": update.ModifiedOn, "modified_by": update.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil

}

//Delete Lesson

func (Coursemodels CoursesModel) DeleteLesson(lessonid int, coursesid int, tenantid string, deletedby int, deletedon time.Time, DB *gorm.DB) error {

	if err := DB.Table("tbl_lessons").Where("id=? and course_id=? and tenant_id=?", lessonid, coursesid, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": deletedby, "deleted_on": deletedon}).Error; err != nil {

		return err
	}

	return nil
}

//Update Order index for Lesson

func (Coursemodels CoursesModel) UpdateLessonOrderIndex(lesson *TblLesson, lessonid int, courseid int, DB *gorm.DB, tenantid string) error {

	if err := DB.Table("tbl_lessons").Where("id=? and course_id=? and tenant_id=?", lessonid, courseid, lesson.TenantId).UpdateColumns(map[string]interface{}{"order_index": lesson.OrderIndex}).Error; err != nil {

		return err
	}

	return nil
}

//Lesson Reorder

func (Coursemodels CoursesModel) UpdateLessonOrder(lesson *TblLesson, courseid int, sectionID int, DB *gorm.DB) error {

	if err := DB.Table("tbl_lessons").Where("id=? and course_id=? and tenant_id=?", lesson.Id, courseid, lesson.TenantId).UpdateColumns(map[string]interface{}{"order_index": lesson.OrderIndex, "section_id": sectionID, "modified_by": lesson.ModifiedBy, "modified_on": lesson.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

//Publish Course

func (Coursemodels CoursesModel) StatusChange(coursesid int, status int, tenantid string, DB *gorm.DB) error {

	if err := DB.Table("tbl_courses").Where("id=? and tenant_id=?", coursesid, tenantid).UpdateColumns(map[string]interface{}{"status": status}).Error; err != nil {

		return err
	}

	return nil
}

//Update Order index for section

func (Coursemodels CoursesModel) UpdateSectionOrderIndex(Section *TblSection, sectionid int, DB *gorm.DB, tenantid string) error {

	fmt.Println("Section:", Section)

	if err := DB.Table("tbl_sections").Where("id=? and tenant_id=?", sectionid, tenantid).UpdateColumns(map[string]interface{}{"order_index": Section.OrderIndex}).Error; err != nil {

		return err
	}

	return nil
}

//Section Reorder

func (Coursemodels CoursesModel) UpdateSectionOrder(Section *TblSection, courseid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_sections").Where("id=? and course_id=? and tenant_id=?", Section.Id, courseid, Section.TenantId).UpdateColumns(map[string]interface{}{"order_index": Section.OrderIndex, "modified_by": Section.ModifiedBy, "modified_on": Section.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}
