package spaces

import (
	"strings"
	"time"

	"github.com/spurtcms/courses/migration"

	"github.com/spurtcms/categories"
)

// spacesetup
func SpaceSetup(config Config) *Spaces {

	migration.AutoMigration(config.DB, config.DataBaseType)

	return &Spaces{
		DB:               config.DB,
		AuthEnable:       config.AuthEnable,
		PermissionEnable: config.PermissionEnable,
		Auth:             config.Auth,
	}
}

/*spacelist*/
func (spaces *Spaces) SpaceList(spacelistreq SpaceListReq) (tblspace []Tblspacesaliases, totalcount int64, err error) {

	autherr := AuthandPermission(spaces)
	if autherr != nil {

		return []Tblspacesaliases{}, 0, autherr
	}

	spacess, _, spaceerr := Spacemodel.SpaceList(spacelistreq, []int{}, spaces.DB)
	if spaceerr != nil {

		return []Tblspacesaliases{}, 0, spaceerr
	}

	var SpaceDetails []Tblspacesaliases

	for _, space := range spacess {

		child_page, _ := categories.Categorymodel.GetCategoryById(space.PageCategoryId, spaces.DB)
		var categorynames []categories.TblCategories

		var flg int

		categorynames = append(categorynames, child_page)
		flg = child_page.ParentId

		var count int

		if flg != 0 {

		CLOOP:

			for {

				count++
				if count >= 50 { // for safe
					break //for safe
				}

				child, _ := categories.Categorymodel.GetCategoryById(flg, spaces.DB)
				flg = child.ParentId

				if flg != 0 {

					categorynames = append(categorynames, child)
					goto CLOOP
				} else {

					categorynames = append(categorynames, child)
					break
				}

			}

		}

		var reverseCategoryOrder []categories.TblCategories
		for i := len(categorynames) - 1; i >= 0; i-- {
			reverseCategoryOrder = append(reverseCategoryOrder, categorynames[i])

		}

		var pageupd TblPageAliases
		Spacemodel.GetLastUpdatePageAliases(&pageupd, space.Id, spaces.DB)
		space.SpaceFullDescription = space.SpacesDescription
		Spiltdescription := TruncateDescription(space.SpacesDescription, 85)
		space.SpacesDescription = Spiltdescription
		space.CategoryNames = reverseCategoryOrder

		SpaceDetails = append(SpaceDetails, space)

	}

	_, count, _ := Spacemodel.SpaceList(SpaceListReq{Keyword: spacelistreq.Keyword, CategoryId: spacelistreq.CategoryId}, []int{}, spaces.DB)

	return SpaceDetails, count, nil

}

/*SpaceDetail*/
func (spaces *Spaces) SpaceDetail(spd SpaceDetail) (space Tblspacesaliases, err error) {

	autherr := AuthandPermission(spaces)
	if autherr != nil {

		return Tblspacesaliases{}, autherr
	}

	spacename, err1 := Spacemodel.GetSpacealiaseDetails(spd.SpaceId, spd.SpaceSlug, spaces.DB)

	return spacename, err1

}

// create space
func (spaces *Spaces) SpaceCreation(SPC SpaceCreation) (tblspac Tblspacesaliases, err error) {

	autherr := AuthandPermission(spaces)
	if autherr != nil {

		return Tblspacesaliases{}, autherr
	}

	var space tblspaces
	space.PageCategoryId = SPC.CategoryId
	space.CreatedOn = CurrentTime
	space.CreatedBy = SPC.CreatedBy
	createspace, err := Spacemodel.CreateSpace(space, spaces.DB)
	if err != nil {
		return Tblspacesaliases{}, err
	}

	var spacealiase Tblspacesaliases
	spacealiase.SpacesName = SPC.Name
	spacealiase.SpacesDescription = SPC.Description
	spacealiase.ImagePath = SPC.ImagePath
	spacealiase.LanguageId = SPC.LanguageId
	spacealiase.CreatedOn = CurrentTime
	spacealiase.CreatedBy = SPC.CreatedBy
	spacealiase.SpacesSlug = strings.ToLower(spacealiase.SpacesName)
	spacealiase.SpacesId = createspace.Id
	Spacemodel.CreateSpaceAliase(spacealiase, spaces.DB)

	return spacealiase, nil

}

// update space
func (spaces *Spaces) SpaceUpdate(SPC SpaceCreation, spaceid int) error {

	autherr := AuthandPermission(spaces)
	if autherr != nil {

		return autherr
	}

	var spaceali Tblspacesaliases
	spaceali.Id = spaceid
	spaceali.SpacesName = SPC.Name
	spaceali.SpacesDescription = SPC.Description
	spaceali.SpacesSlug = strings.ToLower(SPC.Name)
	spaceali.ImagePath = SPC.ImagePath
	spaceali.ModifiedOn = CurrentTime
	spaceali.ModifiedBy = SPC.ModifiedBy
	err1 := Spacemodel.UpdateSpaceAliases(&spaceali, spaceid, spaces.DB)

	if err1 != nil {

		return err1
	}

	var space tblspaces
	space.Id = spaceid
	space.PageCategoryId = SPC.CategoryId
	space.ModifiedOn = CurrentTime
	space.ModifiedBy = SPC.ModifiedBy
	err2 := Spacemodel.UpdateSpace(&space, spaceid, spaces.DB)

	if err2 != nil {

		return err2
	}

	return nil

}

func (spaces *Spaces) DeleteSpaceAliase(spaceid int, deletedBy int) error {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return autherr
	}

	var spacealias TblSpacesAliases
	spacealias.DeletedOn = CurrentTime
	spacealias.DeletedBy = deletedBy
	spacealias.IsDeleted = 1
	err := Spacemodel.DeleteSpaceAliases(&spacealias, spaceid, spaces.DB)

	if err != nil {

		return err
	}

	return nil
}

/*Delete Space*/
func (spaces *Spaces) DeleteSpace(spaceid int, userid int) error {

	var (
		tblspaces TblSpacesAliases
		space     TblSpaces
		pageali   TblPageAliases
	)

	tblspaces.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	tblspaces.DeletedBy = userid
	tblspaces.IsDeleted = 1
	space.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	space.DeletedBy = userid
	space.IsDeleted = 1

	var deletedon, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	var deletedby = userid
	var isdeleted = 1

	pageali.DeletedOn = deletedon
	pageali.DeletedBy = deletedby
	pageali.IsDeleted = isdeleted
	err1 := Spacemodel.DeleteSpaceAliases(&tblspaces, spaceid, spaces.DB)
	if err1 != nil {
		return err1
	}

	err2 := Spacemodel.DeleteSpace(&space, spaceid, spaces.DB)

	if err2 != nil {
		return err2
	}

	var page []TblPage
	Spacemodel.GetPageDetailsBySpaceId(&page, spaceid, spaces.DB)

	var pid []int
	if len(page) != 0 {
		for _, v := range page {
			pid = append(pid, v.Id)
		}

		var pg TblPage
		pg.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pg.DeletedBy = userid
		pg.IsDeleted = 1
		Spacemodel.DeletePageInSpace(&pg, pid, spaces.DB)

		var pgali TblPageAliases
		pgali.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pgali.DeletedBy = userid
		pgali.IsDeleted = 1
		Spacemodel.DeletePageAliInSpace(&pgali, pid, spaces.DB)

		var pagegroup []TblPagesGroup
		Spacemodel.GetPageGroupDetailsBySpaceId(&pagegroup, spaceid, spaces.DB)

		var pagegroupid int
		for _, v := range page {
			v.Id = pagegroupid

		}

		var pggroupdel TblPagesGroup
		pggroupdel.DeletedBy = userid
		pggroupdel.IsDeleted = 1
		pggroupdel.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		Spacemodel.SpaceDeletePageGroup(&pggroupdel, pagegroupid, spaces.DB)

		var pggroupalidel TblPagesGroupAliases
		pggroupalidel.DeletedBy = userid
		pggroupalidel.IsDeleted = 1
		pggroupalidel.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		Spacemodel.SpaceDeletePageGroupAliases(&pggroupalidel, pagegroupid, spaces.DB)

	}

	return nil
}

// clone space - func helps to create duplicate space, using given space id
func (spaces *Spaces) CloneSpace(createspace SpaceCreation, spaceid int, createdBy int) (Tblspacesaliases, error) {

	autherr := AuthandPermission(spaces)
	if autherr != nil {
		return Tblspacesaliases{}, autherr
	}

	space, err := spaces.SpaceDetail(SpaceDetail{SpaceId: spaceid})
	if err != nil {
		return Tblspacesaliases{}, err

	}

	var cspace tblspaces
	cspace.PageCategoryId = createspace.CategoryId
	cspace.CreatedBy = createdBy
	cspace.CreatedOn = CurrentTime
	tblspaces, serr := Spacemodel.CreateSpace(cspace, spaces.DB)
	if serr != nil {
		return Tblspacesaliases{}, serr
	}

	var cspaceali Tblspacesaliases
	cspaceali.CategoryId = createspace.CategoryId
	cspaceali.SpacesId = tblspaces.Id
	cspaceali.SpacesName = createspace.Name
	cspaceali.SpacesSlug, _ = CreateSlug(space.SpacesName)
	cspaceali.SpacesDescription = space.SpacesDescription
	cspaceali.CreatedBy = createdBy
	cspaceali.LanguageId = space.LanguageId
	cspaceali.CreatedOn = CurrentTime
	spacealise, err := Spacemodel.CreateSpaceAliase(cspaceali, spaces.DB)

	var pagegroupdata []TblPagesGroupAliases
	Spacemodel.GetPageGroupData(&pagegroupdata, spaceid, spaces.DB)

	for _, value := range pagegroupdata {

		var group TblPagesGroup
		group.SpacesId = tblspaces.Id
		group.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		groups, _ := Spacemodel.CreatePageGroup(&group, spaces.DB)
		pagegroup := value
		pagegroup.PageGroupId = groups.Id
		pagegroup.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		Spacemodel.ClonePagesGroup(&pagegroup, spaces.DB)
	}

	var pageId []Tblpagealiases
	Spacemodel.GetPageInPage(&pageId, spaceid, spaces.DB) //parentid 0 and groupid 0

	for _, val := range pageId {
		var page TblPage
		page.SpacesId = tblspaces.Id
		page.PageGroupId = 0
		page.ParentId = 0
		page.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pageid, _ := Spacemodel.CreatePage(&page, spaces.DB)

		// var pagesali TblPageAliases
		pagesali := val
		pagesali.PageId = pageid.Id
		pagesali.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		Spacemodel.ClonePages(&pagesali, spaces.DB)

	}

	var pagegroupaldata TblPagesGroupAliases
	Spacemodel.GetIdInPage(&pagegroupaldata, spaceid, spaces.DB) // parentid = 0 and groupid != 0

	var pagealiase []Tblpagealiases
	Spacemodel.GetPageAliasesInPage(&pagealiase, spaceid, spaces.DB) // parentid = 0 and groupid != 0

	for _, value := range pagealiase {
		var pageal TblPagesGroupAliases
		Spacemodel.GetDetailsInPageAli(&pageal, pagegroupaldata.GroupName, tblspaces.Id, spaces.DB)
		var page TblPage
		page.SpacesId = tblspaces.Id
		page.PageGroupId = pageal.PageGroupId
		page.ParentId = 0
		page.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pagess, _ := Spacemodel.CreatePage(&page, spaces.DB)

		// var pagesali TblPageAliases
		pagesali := value
		pagesali.PageId = pagess.Id
		pagesali.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		Spacemodel.ClonePages(&pagesali, spaces.DB)

	}

	var pagealiasedata []Tblpagealiases
	Spacemodel.GetPageAliasesInPageData(&pagealiasedata, spaceid, spaces.DB) // parentid != 0 and groupid = 0

	for _, result := range pagealiasedata {

		var newgroupid int
		if result.PageGroupId != 0 {
			var pagesgroupal TblPagesGroupAliases
			Spacemodel.GetDetailsInPageAlia(&pagesgroupal, result.PageGroupId, spaceid, spaces.DB) // parentid != 0 and groupid = 0
			var pageal TblPagesGroupAliases
			Spacemodel.GetDetailsInPageAli(&pageal, pagesgroupal.GroupName, tblspaces.Id, spaces.DB)
			newgroupid = pageal.PageGroupId

		}

		var pagealid TblPageAliases
		Spacemodel.AliasesInParentId(&pagealid, result.ParentId, spaceid, spaces.DB)

		var pageali TblPageAliases
		Spacemodel.LastLoopAliasesInPage(&pageali, pagealid.PageTitle, tblspaces.Id, spaces.DB)

		var page TblPage
		page.SpacesId = tblspaces.Id
		page.PageGroupId = newgroupid
		page.ParentId = pageali.PageId
		page.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pagealiid, _ := Spacemodel.CreatePage(&page, spaces.DB)

		// var pagesali TblPageAliases
		pagesali := result
		pagesali.PageId = pagealiid.Id
		Spacemodel.ClonePages(&pagesali, spaces.DB)

	}
	return spacealise, err
}

// Dashboard pagescount function
func (spaces *Spaces) DashboardPagesCount() (totalcount int, lasttendayscount int, err error) {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return 0, 0, autherr
	}

	allpagecount, err := Spacemodel.PageCount(spaces.DB)

	if err != nil {

		return 0, 0, err
	}

	Newpagecount, err := Spacemodel.NewpageCount(spaces.DB)

	if err != nil {

		return 0, 0, err
	}

	return int(allpagecount), int(Newpagecount), nil

}

// Check Name is already exits or not
func (spaces *Spaces) CheckSpaceName(id int, name string) (bool, error) {

	var space Tblspacesaliases

	err := Spacemodel.CheckSpaceName(&space, id, name, spaces.DB)

	if err != nil {

		return false, err
	}

	return true, nil
}
