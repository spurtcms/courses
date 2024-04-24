package spaces

import (
	"strings"

	"github.com/spurtcms/categories"
)

// spacesetup
func SpaceSetup(config *Config) *Spaces {

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

		// categorynames = append(categorynames, child_page)

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

					// categorynames = append(categorynames, child)

					goto CLOOP

				} else {

					// categorynames = append(categorynames, child)

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

	Spacemodel.CreateSpace(space, spaces.DB)

	var spacealiase Tblspacesaliases

	spacealiase.SpacesName = SPC.Name

	spacealiase.SpacesDescription = SPC.Description

	spacealiase.ImagePath = SPC.ImagePath

	spacealiase.LanguageId = SPC.LanguageId

	spacealiase.CreatedOn = CurrentTime

	spacealiase.CreatedBy = SPC.CreatedBy

	spacealiase.SpacesSlug = strings.ToLower(spacealiase.SpacesName)

	spacealiase.SpacesId = space.Id

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
func (spaces *Spaces) DeleteSpace(spaceid int, deletedBy int) error {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return autherr
	}

	var space TblSpaces

	space.DeletedOn = CurrentTime

	space.DeletedBy = deletedBy

	space.IsDeleted = 1

	err := Spacemodel.DeleteSpace(&space, spaceid, spaces.DB)

	if err != nil {

		return err

	}

	return nil

}

// clone space - func helps to create duplicate space, using given space id
func (spaces *Spaces) CloneSpace(spaceid int, createdBy int) (Tblspacesaliases, error) {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return Tblspacesaliases{}, autherr
	}

	space, err := spaces.SpaceDetail(SpaceDetail{SpaceId: spaceid})

	if err != nil {

		return Tblspacesaliases{}, err

	}

	var cspace tblspaces

	cspace.PageCategoryId = space.PageCategoryId

	cspace.CreatedBy = createdBy

	cspace.CreatedOn = CurrentTime

	latestspace, serr := Spacemodel.CreateSpace(cspace, spaces.DB)

	if serr != nil {

		return Tblspacesaliases{}, serr
	}

	var cspaceali Tblspacesaliases

	cspaceali.CategoryId = space.PageCategoryId

	cspaceali.SpacesId = latestspace.Id

	cspaceali.SpacesName = space.SpacesName

	cspaceali.SpacesSlug, _ = CreateSlug(space.SpacesName)

	cspaceali.SpacesDescription = space.SpacesDescription

	cspaceali.CreatedBy = createdBy

	cspaceali.CreatedOn = CurrentTime

	spacealise, err := Spacemodel.CreateSpaceAliase(cspaceali, spaces.DB)

	return spacealise, err
}
