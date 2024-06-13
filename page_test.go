package spaces

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
)

// test getpages function
func TestGetPages(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Spaces", auth.CRUD)

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		pages, count, err := space.GetPages(GetPageReq{Spaceid: 2, PublishedPageonly: true})

		if err != nil {

			panic(err)
		}

		fmt.Println(pages, count)

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// test getgroup function
func TestGetGroup(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})
	pagegroup, err := space.GetGroup(2)

	if err != nil {

		panic(err)
	}
	fmt.Println(pagegroup)
}

// test getsubpages function
func TestGetSubPages(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})
	subpages, subpage, err := space.GetSubPages(GetPageReq{Spaceid: 2})

	if err != nil {

		panic(err)
	}
	fmt.Println(subpages, subpage)
}

// test pagealiaseslog function
func TestPageAliasesLog(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})
	pagelog, err := space.PageAliasesLog(2)

	if err != nil {

		panic(err)
	}
	fmt.Println(pagelog)
}

// test createpagelog function
func TestCreatePagelog(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	page := Pages{PgId: 1, Name: "Introduction", ParentId: 1}

	pagelog, err := space.CreatePagelog(2, []Pages{page}, "public")

	if err != nil {

		panic(err)
	}
	fmt.Println(pagelog)
}

// test createpage function
func TestCreatePage(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	page := Pages{Name: "Introduction", ParentId: 1}

	pagelog, err := space.CreatePage(2, []Pages{page}, "public")

	if err != nil {

		panic(err)
	}
	fmt.Println(pagelog)
}

// test creategroup function
func TestCreateGroup(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	group := PageGroups{Name: "Demo", OrderIndex: 1}

	pagegroup, err := space.CreateGroup(2, []PageGroups{group}, 1)

	if err != nil {

		panic(err)
	}
	fmt.Println(pagegroup)
}

// test createsubpage function
func TestCreateSubpage(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	spage := SubPages{Name: "Introduction", ParentId: 1}

	subpage, err := space.CreateSubpage(2, []SubPages{spage})

	if err != nil {

		panic(err)
	}
	fmt.Println(subpage)
}

// test updatepage function
func TestUpdatePage(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	page := Pages{PgId: 1, Name: "Introduction", OrderIndex: 1}

	pagelog, err := space.UpdatePage(2, []Pages{page})

	if err != nil {

		panic(err)
	}
	fmt.Println(pagelog)
}

// test updategroup function
func TestUpdateGroup(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	group := PageGroups{Name: "Demo1", OrderIndex: 1}

	pagegroup, err := space.UpdateGroup([]PageGroups{group}, 1)

	if err != nil {

		panic(err)
	}
	fmt.Println(pagegroup)
}

// test updatesubpage function
func TestUpdateSubpage(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	spage := SubPages{SpgId: 4, Name: "Introduction", ParentId: 1}

	subpage, err := space.UpdateSubpage(2, []SubPages{spage})

	if err != nil {

		panic(err)
	}
	fmt.Println(subpage)
}

// test deletepage function
func TestDeletePage(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	err := space.DeletePage(DeletePagereq{SpaceId: 1})

	if err != nil {

		panic(err)
	}
}

// test deletepagegroup function
func TestDeletedPageGroup(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	err := space.DeletedPageGroup(DeletePageGroupreq{GroupId: 1})

	if err != nil {

		panic(err)
	}
}

// test deletesubpage function
func TestDeleteSubPage(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	err := space.DeleteSubPage(DeletePagereq{Id: 4})

	if err != nil {

		panic(err)
	}
}
