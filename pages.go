package spaces

import (
	"log"
	"strings"
	"time"
)

// get page list
func (spaces *Spaces) GetPages(pagereq GetPageReq) ([]Pages, Pages, error) {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return []Pages{}, Pages{}, autherr
	}

	if pagereq.Spaceid != 0 {

		pages, err := GetPageBySpaceIdORPageIds(pagereq, spaces.DB)

		return pages, Pages{}, err

	} else if len(pagereq.PageIds) > 0 {

		pages, err := GetPageBySpaceIdORPageIds(pagereq, spaces.DB)

		return pages, Pages{}, err

	} else if pagereq.PageId != 0 {

		page, err := GetPageByPageId(pagereq, spaces.DB)

		return []Pages{}, page, err

	}

	return []Pages{}, Pages{}, nil

}

// get groups
func (spaces *Spaces) GetGroup(spaceid int) ([]PageGroups, error) {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return []PageGroups{}, autherr
	}

	var group []TblPagesGroup

	var pagegroups []PageGroups

	Spacemodel.SelectGroup(&group, spaceid, []int{}, spaces.DB)

	for _, group := range group {

		var pagegroup TblPagesGroupAliases

		var page_group PageGroups

		Spacemodel.PageGroup(&pagegroup, group.Id, spaces.DB)

		page_group.GroupId = pagegroup.PageGroupId

		page_group.Name = pagegroup.GroupName

		page_group.OrderIndex = pagegroup.OrderIndex

		pagegroups = append(pagegroups, page_group)

	}

	return pagegroups, nil
}

// get subpages
func (spaces *Spaces) GetSubPages(pagereq GetPageReq) ([]SubPages, SubPages, error) {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return []SubPages{}, SubPages{}, autherr
	}

	if pagereq.Spaceid != 0 || len(pagereq.PageIds) > 0 {

		pages, err := GetSubPageBySpaceIdORPageIds(pagereq, spaces.DB)

		return pages, SubPages{}, err

	} else if pagereq.PageId != 0 {

		page, err := GetSubPageByPageId(pagereq, spaces.DB)

		return []SubPages{}, page, err

	}

	return []SubPages{}, SubPages{}, nil

}

// get pages logs
func (spaces *Spaces) PageAliasesLog(spaceid int) ([]PageLog, error) {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return []PageLog{}, autherr
	}

	pagelog, err := Spacemodel.GetPageLogDetails(spaceid, 0, []int{}, spaces.DB)

	if err != nil {

		return []PageLog{}, err
	}

	var finallog []PageLog

	for _, val := range pagelog {

		var log PageLog

		log.Username = val.Username

		if val.ModifiedOn.IsZero() {

			log.Status = "draft"

		} else {

			log.Status = "Updated"
		}

		if val.Status == "publish" {

			log.Status = val.Status
		}

		log.Date = val.CreatedOn

		finallog = append(finallog, log)

	}

	return finallog, nil
}

// create pagelog
func (spaces *Spaces) CreatePagelog(spaceid int, NewPages []Pages, Status string) ([]TblPageAliasesLog, error) {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return []TblPageAliasesLog{}, autherr
	}

	var pglogs []TblPageAliasesLog

	/*Create Pages*/
	for _, val := range NewPages {

		/*This is for log*/
		var pagelog TblPageAliasesLog

		pagelog.PageId = val.PageId

		pagelog.LanguageId = 1

		pagelog.PageTitle = val.Title

		pagelog.PageSlug = strings.ToLower(strings.ReplaceAll(val.Title, " ", "_"))

		pagelog.PageDescription = val.Content

		pagelog.PageId = val.PageId

		pagelog.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		pagelog.CreatedBy = val.CreatedBy

		pagelog.Status = Status

		pagelog.Access = "public"

		pagelog.ReadTime = val.ReadTime

		pgl, err := Spacemodel.PageAliasesLog(&pagelog, spaces.DB)

		if err != nil {

			log.Println(err)
		}

		pglogs = append(pglogs, pgl)

	}

	return pglogs, nil
}

// create page
func (spaces *Spaces) CreatePage(SpaceId int, NewPages []Pages, Status string) ([]TblPageAliases, error) {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return []TblPageAliases{}, autherr
	}

	var pagealis []TblPageAliases

	/*Create Pages*/
	for _, val := range NewPages {

		/*page creation tbl_page*/
		var page TblPage

		page.PageGroupId = val.Pgroupid
	
		page.SpacesId = SpaceId

		page.ParentId = 0

		page.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		page.CreatedBy = val.CreatedBy

		pageret, _ := Spacemodel.CreatePage(&page, spaces.DB)

		/*page creation tbl_page_aliases*/
		var pageali TblPageAliases

		pageali.LanguageId = 1

		pageali.PageId = pageret.Id

		pageali.PageTitle = val.Title

		pageali.PageDescription = val.Content

		pageali.PageSlug = strings.ToLower(strings.ReplaceAll(val.Title, " ", "_"))

		pageali.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		pageali.CreatedBy = val.CreatedBy

		pageali.OrderIndex = val.OrderIndex

		pageali.Status = val.Status

		pageali.LastRevisionDate, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		pageali.LastRevisionNo = 1

		pageali.Access = "public"

		pageali.ReadTime = val.ReadTime

		pagealise, err := Spacemodel.CreatepageAliases(&pageali, spaces.DB)

		if err != nil {

			log.Println(err)
		}

		pagealis = append(pagealis, pagealise)

	}

	return pagealis, nil

}

// create group
func (spaces *Spaces) CreateGroup(SpaceId int, NewGroup []PageGroups, CreatedBy int) ([]TblPagesGroupAliases, error) {

	var grpss []TblPagesGroupAliases

	for _, val := range NewGroup {

		/*Group create tbl_page_group*/
		var groups TblPagesGroup

		groups.SpacesId = SpaceId

		groups.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		groups.CreatedBy = CreatedBy

		grpreturn, _ := Spacemodel.CreatePageGroup(&groups, spaces.DB)

		/*group aliases tbl_page_group_aliases*/
		var groupali TblPagesGroupAliases

		groupali.PageGroupId = grpreturn.Id

		groupali.GroupName = strings.ToUpper(val.Name)

		groupali.GroupSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))

		groupali.LanguageId = 1

		groupali.OrderIndex = val.OrderIndex

		groupali.CreatedBy = CreatedBy

		groupali.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		grps, err := Spacemodel.CreatePageGroupAliases(&groupali, spaces.DB)

		if err != nil {

			return []TblPagesGroupAliases{}, err
		}

		grpss = append(grpss, grps)
	}

	return grpss, nil
}

// create subpage
func (spaces *Spaces) CreateSubpage(SpaceId int, NewSubPage []SubPages) ([]TblPageAliases, error) {

	var subpgs []TblPageAliases

	/*createsub*/
	for _, val := range NewSubPage {

		/*page creation tbl_page*/
		var page TblPage

		page.PageGroupId = val.PgroupId

		page.SpacesId = SpaceId

		page.ParentId = val.ParentId

		page.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		page.CreatedBy = val.CreatedBy

		pageret, _ := Spacemodel.CreatePage(&page, spaces.DB)

		/*page creation tbl_page_aliases*/
		var pageali TblPageAliases

		pageali.LanguageId = 1

		pageali.PageId = pageret.Id

		pageali.PageTitle = val.Title

		pageali.PageDescription = val.Content

		pageali.PageSlug = strings.ToLower(strings.ReplaceAll(val.Title, " ", "_"))

		pageali.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		pageali.CreatedBy = val.CreatedBy

		pageali.PageSuborder = val.OrderIndex

		pageali.Status = val.Status

		pageali.Access = "public"

		pageali.LastRevisionDate, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		pageali.LastRevisionNo = 1

		pageali.ReadTime = val.ReadTime

		subpg, err := Spacemodel.CreatepageAliases(&pageali, spaces.DB)

		if err != nil {

			log.Println(err)
		}

		subpgs = append(subpgs, subpg)

	}

	return subpgs, nil
}

// update page
func (spaces *Spaces) UpdatePage(SpaceId int, UpdatePages []Pages) ([]TblPageAliases, error) {

	var updpagss []TblPageAliases

	/*Update Group*/
	for _, val := range UpdatePages {

		var uptpage TblPage

		uptpage.PageGroupId = val.Pgroupid

		uptpage.ParentId = 0

		Spacemodel.UpdatePage(&uptpage, val.PageId, spaces.DB)

		var uptpageali TblPageAliases

		uptpageali.PageTitle = val.Title

		uptpageali.PageSlug = strings.ToLower(strings.ReplaceAll(val.Title, " ", "_"))

		uptpageali.PageDescription = val.Content

		uptpageali.OrderIndex = val.OrderIndex

		uptpageali.Status = val.Status

		uptpageali.ModifiedBy = val.ModifiedBy

		uptpageali.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		uptpageali.LastRevisionDate, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		uptpageali.ReadTime = val.ReadTime

		updpg,err := Spacemodel.UpdatePageAliase(&uptpageali, val.PageId, spaces.DB)

		if err != nil {

			log.Println(err)
		}

		updpagss = append(updpagss, updpg)

	}

	return updpagss, nil
}

// update Group
func (spaces *Spaces) UpdateGroup(UpdateGroup []PageGroups, createdBy int) ([]TblPagesGroupAliases, error) {

	var updgrps []TblPagesGroupAliases

	/*Update Group*/
	for _, val := range UpdateGroup {

		var uptgroup TblPagesGroupAliases

		uptgroup.GroupName = strings.ToUpper(val.Name)

		uptgroup.GroupSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))

		uptgroup.LanguageId = 1

		uptgroup.OrderIndex = val.OrderIndex

		uptgroup.ModifiedBy = createdBy

		uptgroup.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		updgr, err := Spacemodel.UpdatePageGroupAliases(&uptgroup, val.GroupId, spaces.DB)

		updgrps = append(updgrps, updgr)

		if err != nil {

			return []TblPagesGroupAliases{}, err
		}

	}

	return updgrps, nil
}

// update Subpage
func (spaces *Spaces) UpdateSubpage(SpaceId int, UpdateSubPage []SubPages) ([]TblPageAliases, error) {

	var tblpagealiase []TblPageAliases

	/*update subpages*/
	for _, val := range UpdateSubPage {

		var uptpage TblPage

		uptpage.PageGroupId = val.PgroupId

		uptpage.ParentId = val.ParentId

		Spacemodel.UpdatePage(&uptpage, val.SubPageId, spaces.DB)

		var uptpageali TblPageAliases

		uptpageali.PageTitle = val.Title

		uptpageali.PageSlug = strings.ToLower(strings.ReplaceAll(val.Title, " ", "_"))

		uptpageali.PageDescription = val.Content

		uptpageali.PageSuborder = val.OrderIndex

		uptpageali.Status = val.Status

		uptpageali.ModifiedBy = val.CreatedBy

		uptpageali.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		uptpageali.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		uptpageali.ReadTime = val.ReadTime

		updsub, err := Spacemodel.UpdatePageAliase(&uptpageali, val.SubPageId, spaces.DB)

		if err != nil {

			log.Println(err)
		}

		tblpagealiase = append(tblpagealiase, updsub)
	}

	return tblpagealiase, nil
}

// deletepage
func (spaces *Spaces) DeletePage(delpage DeletePagereq) error {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return autherr
	}

	var err error

	var page TblPage

	page.DeletedOn = CurrentTime

	page.DeletedBy = delpage.DeletedBy

	page.IsDeleted = 1

	err = Spacemodel.DeletePage(&page, delpage, spaces.DB)

	var pageali TblPageAliases

	pageali.DeletedOn = CurrentTime

	pageali.DeletedBy = delpage.DeletedBy

	pageali.IsDeleted = 1

	err = Spacemodel.DeletePageAliase(&pageali, delpage, spaces.DB)

	return err
}

// delete pagegroup
func (spaces *Spaces) DeletedPageGroup(delpgg DeletePageGroupreq) error {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return autherr
	}

	var page TblPagesGroup

	page.DeletedOn = CurrentTime

	page.DeletedBy = delpgg.DeletedBy

	page.IsDeleted = 1

	Spacemodel.DeletePageGroup(&page, delpgg, spaces.DB)

	var pageali TblPagesGroupAliases

	pageali.DeletedOn = CurrentTime

	pageali.DeletedBy = delpgg.DeletedBy

	pageali.IsDeleted = 1

	Spacemodel.DeletePageGroupAliases(&pageali, delpgg, spaces.DB)

	return nil

}

// deletepage
func (spaces *Spaces) DeleteSubPage(delpage DeletePagereq) error {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return autherr
	}

	var err error

	var page TblPage

	page.DeletedOn = CurrentTime

	page.DeletedBy = delpage.DeletedBy

	page.IsDeleted = 1

	err = Spacemodel.DeletePage(&page, delpage, spaces.DB)

	var pageali TblPageAliases

	pageali.DeletedOn = CurrentTime

	pageali.DeletedBy = delpage.DeletedBy

	pageali.IsDeleted = 1

	err = Spacemodel.DeletePageAliase(&pageali, delpage, spaces.DB)

	return err
}
