
# Learning Package

With 'learning' package, experience an organized content creation process that allows effortless organization of content presentation on your website. Customize specific areas like the header, footer, or dedicated sections for different content types.  This package will empower you to create a structured website in Golang, with layout tailored to diverse content presentations, ensuring seamless content management with distinct layouts and specialized sections for an enhanced user experience

## Features

- SpaceList: Administrators can access a comprehensive list of all spaces available within the CMS, providing a comprehensive overview of the system's organizational structure and facilitating streamlined navigation.
- SpaceDetail: This function facilitates the retrieval of detailed information regarding specific spaces within the CMS, which helps comprehensive understanding of individual spaces' configurations and attributes.
- SpaceCreation: This function allows administrators to create new spaces within the CMS, facilitating the expansion and customization of the system's content architecture to suit evolving new needs.
- SpaceUpdate: Enables administrators to dynamically modify existing space configurations for its Title, description, cover image etc, ensuring flexibility and adaptability in aligning spaces with changing objectives or requirements.
- DeleteSpace: Facilitates the removal of obsolete or redundant spaces from the CMS, maintaining system cleanliness and optimizing resource utilization.
- CloneSpace: Provides administrators with the ability to duplicate existing spaces along with Page groups and pages mapped to it, enabling efficient content replication across different sections of the CMS.
- PageList: Retrieves a comprehensive list of pages available within the CMS, empowering administrators with a detailed overview of the system's content inventory and organizational structure.
- PageCategoryList: This function retrieves a structured list of page categories.

# Installation

``` bash
go get github.com/spurtcms/Learning
```


# Usage Example

``` bash
import (
	"github.com/spurtcms/auth"
	"github.com/spurtcms/Learning"
)

func main() {

	Auth := auth.AuthSetup(auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		SecretKey:  "SecretKey@123",
		DB: &gorm.DB{},
		Roleid:1,
	})

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Space", auth.CRUD)

	spaceauth := space.SpaceSetup(&space.Config{
		DB:               &gorm.DB{},
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	//spaces
	if permisison {

		//list space
		spacelist, count, err := spaceauth.SpaceList(space.SpaceListReq{Limit: 10, Offset: 0})
		fmt.Println(spacelist, count, err)

		//create space
		_, cerr := spaceauth.SpaceCreation(space.SpaceCreation{
			Name:        "Default_Space",
			Description: "default space",
			CategoryId:  1,
			CreatedBy:   1,
		})

		if cerr != nil {

			fmt.Println(cerr)
		}

		//update space
		uerr := spaceauth.SpaceUpdate(space.SpaceCreation{
			Name:        "Default Space",
			Description: "default space",
			CategoryId:  1,
			ModifiedBy:  1,
		}, 1)

		if uerr != nil {

			fmt.Println(uerr)

		}

		// delete space
		derr := spaceauth.DeleteSpace(1, 1)

		if derr != nil {

			fmt.Println(derr)

		}
		// page list
		pages, pagecount, perr := spaceauth.GetPages(space.GetPageReq{Spaceid: 2, PublishedPageonly: true})
		fmt.Println(pages, pagecount, perr)

		// create page
		pagedata := space.Pages{PageId: 1, Title: "Introduction", ParentId: 1}

		pagelog, cperr := spaceauth.CreatePagelog(2, []space.Pages{pagedata}, "public")
		fmt.Println(pagelog)

		if cperr != nil {

			fmt.Println(cperr)

		}

		// create group
		group := space.PageGroups{Name: "Default_Group", OrderIndex: 1}

		pagegroup, gcerr := spaceauth.CreateGroup(2, []space.PageGroups{group}, 1)
		fmt.Println(pagegroup)

		if gcerr != nil {

			fmt.Println(gcerr)

		}

	}

}

```

# Getting help
If you encounter a problem with the package,please refer [Please refer [(https://www.spurtcms.com/documentation/cms-admin)] or you can create a new Issue in this repo[https://github.com/spurtcms/learning/issues]. 
