package courses

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spurtcms/auth/migration"
	"github.com/spurtcms/categories"
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

func (courses *Courses) EditCourses(id int, tenantid string) (courselist TblCourse, category []categories.Arrangecategories, err error) {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return TblCourse{}, []categories.Arrangecategories{}, Autherr
	}

	courselist, err = Coursemodels.EditCourse(id, tenantid, courses.DB)
	if err != nil {
		fmt.Println(err)
	}

	var FinalSelectedCategories []categories.Arrangecategories

	var idc []int

	ids := strings.Split(courselist.CategoryId, ",")

	for _, cid := range ids {

		convid, _ := strconv.Atoi(cid)

		idc = append(idc, convid)
	}

	GetSelectedCategory, _ := Coursemodels.GetCategoriseById(idc, courses.DB, tenantid)

	var addcat categories.Arrangecategories

	var individualid []categories.CatgoriesOrd

	for _, CategoriesArrange := range GetSelectedCategory {

		var individual categories.CatgoriesOrd

		individual.Id = CategoriesArrange.Id

		individual.Category = CategoriesArrange.CategoryName

		individualid = append(individualid, individual)

	}

	addcat.Categories = individualid

	FinalSelectedCategories = append(FinalSelectedCategories, addcat)

	return courselist, FinalSelectedCategories, nil

}

func (courses *Courses) UpdateCourses(id, userid int, update TblCourse, tenantid string) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	modified_on, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	updates := TblCourse{
		Title:       update.Title,
		Description: update.Description,
		ImageName:   update.ImageName,
		ImagePath:   update.ImagePath,
		CategoryId:  update.CategoryId,
		Status:      update.Status,
		ModifiedOn:  modified_on,
		ModifiedBy:  userid,
	}

	err := Coursemodels.UpdateCourse(id, tenantid, updates, courses.DB)

	if err != nil {

		return err
	}

	return nil

}

func (courses *Courses) EditCourseSettings(id int, tenantid string) (coursesettingslist TblCourseSettings, err error) {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return TblCourseSettings{}, Autherr
	}

	coursesettingslist, err = Coursemodels.EditCourseSettings(id, tenantid, courses.DB)
	if err != nil {
		fmt.Println(err)
	}

	return coursesettingslist, nil

}

func (courses *Courses) UpdateCourseSettings(id int, update TblCourseSettings, tenantid string) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	modified_on, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	updates := TblCourseSettings{
		CourseId:     update.CourseId,
		Certificate:  update.Certificate,
		Comments:     update.Comments,
		Offer:        update.Offer,
		Visibility:   update.Visibility,
		StartDate:    update.StartDate,
		SignUpLimits: update.SignUpLimits,
		Duration:     update.Duration,
		ModifiedOn:   modified_on,
		ModifiedBy:   update.ModifiedBy,
	}

	err := Coursemodels.UpdateCourseSettings(tenantid, updates, courses.DB)

	if err != nil {

		return err
	}

	return nil

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

//Create Section

func (courses *Courses) CreateSections(create TblSection) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Create := TblSection{
		Title:      create.Title,
		Content:    create.Content,
		CourseId:   create.CourseId,
		OrderIndex: create.OrderIndex,
		TenantId:   create.TenantId,
		CreatedOn:  createdon,
		CreatedBy:  create.CreatedBy,
		IsDeleted:  create.IsDeleted,
	}

	err := Coursemodels.CreateSection(Create, courses.DB)

	if err != nil {

		return err
	}

	return nil

}

//List Section

func (courses *Courses) ListSections(id int, tenantid string) (section []TblSection, err error) {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return []TblSection{}, Autherr
	}

	sectionlist, err := Coursemodels.SectionList(id, tenantid, courses.DB)

	if err != nil {

		return []TblSection{}, err
	}

	return sectionlist, nil

}

//Edit Section

func (courses *Courses) EditSections(sectionid int, coursesid int, tenantid string) (section TblSection, err error) {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return TblSection{}, Autherr
	}

	section, err = Coursemodels.EditSection(sectionid, coursesid, tenantid, courses.DB)
	if err != nil {
		fmt.Println(err)
	}

	return section, nil

}

//Update Section

func (courses *Courses) UpdateSections(update TblSection) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	modifiedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Update := TblSection{
		Id:         update.Id,
		Title:      update.Title,
		Content:    update.Content,
		CourseId:   update.CourseId,
		TenantId:   update.TenantId,
		ModifiedOn: modifiedon,
		ModifiedBy: update.ModifiedBy,
	}

	err := Coursemodels.UpdateSection(Update, courses.DB)

	if err != nil {

		return err
	}

	return nil

}

//Delete Section

func (courses *Courses) DeleteSections(sectionid, userid int, courseid int, tenantid string) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	deletedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	deletedby := userid

	err := Coursemodels.DeleteSection(sectionid, courseid, tenantid, deletedby, deletedon, courses.DB)

	if err != nil {

		return err
	}

	return nil

}

//Create Text

func (courses *Courses) CreateLessons(lesson TblLesson) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	createdon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Create := TblLesson{
		CourseId:   lesson.CourseId,
		SectionId:  lesson.SectionId,
		Title:      lesson.Title,
		Content:    lesson.Content,
		EmbedLink:  lesson.EmbedLink,
		OrderIndex: lesson.OrderIndex,
		TenantId:   lesson.TenantId,
		LessonType: lesson.LessonType,
		CreatedOn:  createdon,
		CreatedBy:  lesson.CreatedBy,
		IsDeleted:  lesson.IsDeleted,
	}

	err := Coursemodels.CreateLesson(Create, courses.DB)

	if err != nil {

		return err
	}

	return nil

}

//List Lesson

func (courses *Courses) ListLessons(id int, tenantid string) (lesson []TblLesson, err error) {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return []TblLesson{}, Autherr
	}

	lessonlist, err := Coursemodels.LessonList(id, tenantid, courses.DB)

	if err != nil {

		return []TblLesson{}, err
	}

	return lessonlist, nil

}

//Edit Lesson

func (courses *Courses) EditLessons(lessonid int, coursesid int, tenantid string) (lesson TblLesson, err error) {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return TblLesson{}, Autherr
	}

	lesson, err = Coursemodels.EditLesson(lessonid, coursesid, tenantid, courses.DB)
	if err != nil {
		fmt.Println(err)
	}

	return lesson, nil

}

// Update Lesson
func (courses *Courses) UpdateLessons(update TblLesson) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	modifiedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	Update := TblLesson{
		Id:         update.Id,
		CourseId:   update.CourseId,
		SectionId:  update.SectionId,
		Title:      update.Title,
		Content:    update.Content,
		EmbedLink:  update.EmbedLink,
		TenantId:   update.TenantId,
		ModifiedOn: modifiedon,
		ModifiedBy: update.ModifiedBy,
	}

	err := Coursemodels.UpdateLesson(Update, courses.DB)

	if err != nil {

		return err
	}

	return nil

}

//Delete Lesson

func (courses *Courses) DeleteLessons(lessonid, userid int, courseid int, tenantid string) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	deletedon, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	deletedby := userid

	err := Coursemodels.DeleteLesson(lessonid, courseid, tenantid, deletedby, deletedon, courses.DB)

	if err != nil {

		return err
	}

	return nil

}

//Update Order index for Lesson

func (courses *Courses) UpdateLessonOrderIndexes(Orderindex int, lessonid, courseid int, userid int, tenantid string) (bool, error) {

	autherr := AuthandPermission(courses)

	if autherr != nil {

		return false, autherr
	}

	var Lesson TblLesson

	Lesson.OrderIndex = Orderindex

	Lesson.TenantId = tenantid

	err := Coursemodels.UpdateLessonOrderIndex(&Lesson, lessonid, courseid, courses.DB, tenantid)

	if err != nil {

		return false, err
	}

	return true, nil
}

//Lesson OrderIndex Reorder

func (courses *Courses) UpdateLessonOrders(lessonids []int, tenantid string, userid, courseid int, sectionID int) error {
	autherr := AuthandPermission(courses)
	if autherr != nil {
		return autherr
	}

	var lesson TblLesson

	lesson.ModifiedBy = userid
	lesson.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	lesson.TenantId = tenantid

	for index, id := range lessonids {

		if id != 0 {

			lesson.Id = id

			lesson.OrderIndex = index + 1

			err := Coursemodels.UpdateLessonOrder(&lesson, courseid, sectionID, courses.DB)

			if err != nil {

				return err
			}
		}

	}

	return nil
}

//Publish Course

func (courses *Courses) StatusChanges(courseid int, status int, tenantid string) error {

	if Autherr := AuthandPermission(courses); Autherr != nil {

		return Autherr
	}

	err := Coursemodels.StatusChange(courseid, status, tenantid, courses.DB)

	if err != nil {

		return err
	}

	return nil

}

//Update Order index for section

func (courses *Courses) UpdateSectionOrderIndexes(Orderindex int, sectionid int, userid int, tenantid string) (bool, error) {

	autherr := AuthandPermission(courses)

	if autherr != nil {

		return false, autherr
	}

	var Section TblSection

	Section.OrderIndex = Orderindex

	err := Coursemodels.UpdateSectionOrderIndex(&Section, sectionid, courses.DB, tenantid)

	if err != nil {

		return false, err
	}

	return true, nil
}

//Section OrderIndex Reorder

func (courses *Courses) UpdateSectionOrders(sectionids []int, tenantid string, userid, courseid int) error {
	autherr := AuthandPermission(courses)
	if autherr != nil {
		return autherr
	}

	var section TblSection

	section.ModifiedBy = userid
	section.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	section.TenantId = tenantid

	for index, id := range sectionids {

		if id != 0 {

			section.Id = id

			section.OrderIndex = index + 1

			err := Coursemodels.UpdateSectionOrder(&section, courseid, courses.DB)

			if err != nil {

				return err
			}
		}

	}

	return nil
}
