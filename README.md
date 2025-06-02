# Courses Package

The SpurtCMS Courses Package allows you to create, manage, and deliver educational content with ease. It supports organizing content into courses, modules, and lessons, making it ideal for building e-learning platforms, internal training systems, or paid course sites. This package gives you the tools to structure and present educational material effectively.


## Features

- Create and manage multiple courses with categories and tags
- Organize content into modules and lessons
- Assign courses based on membership level or group
- Restrict courses to specific users, groups, or subscription plans
- Add quizzes or tests to evaluate learning


# Installation

``` bash
go get github.com/spurtcms/courses
```


# Usage Example


``` bash
import (
	"github.com/spurtcms/auth"
	"github.com/spurtcms/courses"
)

func main() {

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "SecretKey@123",
		DB: &gorm.DB{},
		RoleId: 1,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Courses", auth.CRUD)

	CoursesConfig = courses.CoursesSetup(courses.Config{
		DB:               DB,
		AuthEnable:       true,
		PermissionEnable: false,
		Auth:             NewAuth,
	})

	//Courses
	if permisison {

		//list Course
		courseslist, count, err := CoursesConfig.CoursesList(limt, offset, filter, TenantId)
		if err != nil {
			fmt.Println(err)
		}

		//create Course
		createerr := CoursesConfig.CreateCourse(Create)
		if createerr != nil {
			fmt.Println(createerr)
		}


		//update Course
	    err := CoursesConfig.UpdateCourses(courseidint, userid, UpdateCourse, TenantId)
	    if err != nil {
		    fmt.Println(err)
	    }

		// delete Course
		err := CoursesConfig.DeleteCourses(id, userid, TenantId)
		if err != nil {
			fmt.Println(err)
			c.Redirect(301, "/403-page")
			return
		}

        //create Lessons
        err := CoursesConfig.CreateLessons(create)
		if err != nil {
			fmt.Println(err)
		}

        //edit Lesson
        lesson, err := CoursesConfig.EditLessons(lessonid, courseid, TenantId)
		if err != nil {
			fmt.Println(err)
		}

        //update Lesson
        err := CoursesConfig.UpdateLessons(update)
		if err != nil {
			fmt.Println(err)
		}

        //delete Lesson
        err := CoursesConfig.DeleteLessons(lessonid, userid, courseid, TenantId)
		if err != nil {
			fmt.Println(err)
			c.Redirect(301, "/403-page")
			return
		}

        //create Section
        err := CoursesConfig.CreateSections(create)
		if err != nil {
			fmt.Println(err)
		}

        //edit Section
        section, err := CoursesConfig.EditSections(sectionid, courseid, TenantId)
		if err != nil {
			fmt.Println(err)
		}

        //update Section
        err := CoursesConfig.UpdateSections(Update)
		if err != nil {
			fmt.Println(err)
		}

        //delete Section
        err := CoursesConfig.DeleteSections(sectionid, userid, courseid, TenantId)
		if err != nil {
			fmt.Println(err)
			c.Redirect(301, "/403-page")
			return
		}
        
	}
}

```

# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/courses/issues]. 
