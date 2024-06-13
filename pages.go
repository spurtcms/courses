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

	var (
		group      []TblPagesGroup
		pagegroups []PageGroups
	)

	Spacemodel.SelectGroup(&group, spaceid, []int{}, spaces.DB)
	for _, group := range group {

		var (
			pagegroup  TblPagesGroupAliases
			page_group PageGroups
		)

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
		pagelog.PageId = val.PgId
		pagelog.LanguageId = 1
		pagelog.PageTitle = val.Name
		pagelog.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		pagelog.PageDescription = val.Content
		pagelog.PageId = val.PgId
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
func (spaces *Spaces) CreatePage(SpaceId int, NewPages []Pages, Status string) ([]Tblpagealiases, error) {

	autherr := AuthandPermission(spaces)
	if autherr != nil {

		return []Tblpagealiases{}, autherr
	}

	var pagealis []Tblpagealiases
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
		var pageali Tblpagealiases
		pageali.LanguageId = 1
		pageali.PageId = pageret.Id
		pageali.PageTitle = val.Name
		pageali.PageDescription = val.Content
		pageali.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
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
func (spaces *Spaces) CreateSubpage(SpaceId int, NewSubPage []SubPages) ([]Tblpagealiases, error) {

	var subpgs []Tblpagealiases
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
		var pageali Tblpagealiases
		pageali.LanguageId = 1
		pageali.PageId = pageret.Id
		pageali.PageTitle = val.Name
		pageali.PageDescription = val.Content
		pageali.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
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
		Spacemodel.UpdatePage(&uptpage, val.PgId, spaces.DB)

		var uptpageali TblPageAliases
		uptpageali.PageTitle = val.Name
		uptpageali.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		uptpageali.PageDescription = val.Content
		uptpageali.OrderIndex = val.OrderIndex
		uptpageali.Status = val.Status
		uptpageali.ModifiedBy = val.ModifiedBy
		uptpageali.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		uptpageali.LastRevisionDate, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		uptpageali.ReadTime = val.ReadTime
		updpg, err := Spacemodel.UpdatePageAliase(&uptpageali, val.PgId, spaces.DB)

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
		Spacemodel.UpdatePage(&uptpage, val.SpgId, spaces.DB)

		var uptpageali TblPageAliases
		uptpageali.PageTitle = val.Name
		uptpageali.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		uptpageali.PageDescription = val.Content
		uptpageali.PageSuborder = val.OrderIndex
		uptpageali.Status = val.Status
		uptpageali.ModifiedBy = val.CreatedBy
		uptpageali.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		uptpageali.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		uptpageali.ReadTime = val.ReadTime
		updsub, err := Spacemodel.UpdatePageAliase(&uptpageali, val.SpgId, spaces.DB)

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

	var (
		err  error
		page TblPage
	)

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

	var (
		err  error
		page TblPage
	)

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


/*Create page*/
func (spaces *Spaces) CreatePages(Pagec PageCreate, userid int) error {

	var status string
	if Pagec.Status == "publish" {
		status = "publish"
	} else if Pagec.Status == "save" {
		status = "draft"
	}

	type TempCheck struct {
		FrontId    int
		NewFrontId int
		DBid       int
	}

	spaceId := Pagec.SpaceId

	var (
		Temparr []TempCheck
		err error
	)

	/*Create Group*/
	for _, val := range Pagec.NewGroup {

		/*Group create tbl_page_group*/
		var groups TblPagesGroup
		groups.SpacesId = spaceId
		groups.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		groups.CreatedBy = userid
		grpreturn, _ := Spacemodel.CreatePageGroup(&groups, spaces.DB)

		/*group aliases tbl_page_group_aliases*/
		var groupali TblPagesGroupAliases
		groupali.PageGroupId = grpreturn.Id
		groupali.GroupName = strings.ToUpper(val.Name)
		groupali.GroupSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		groupali.LanguageId = 1
		groupali.OrderIndex = val.OrderIndex
		groupali.CreatedBy = userid
		groupali.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		_, err = Spacemodel.CreatePageGroupAliases(&groupali, spaces.DB)

	}

	/*Update Group*/
	for _, val := range Pagec.UpdateGroup {

		var uptgroup TblPagesGroupAliases
		uptgroup.GroupName = strings.ToUpper(val.Name)
		uptgroup.GroupSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		uptgroup.LanguageId = 1
		uptgroup.OrderIndex = val.OrderIndex
		uptgroup.ModifiedBy = userid
		uptgroup.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		_, err = Spacemodel.UpdatePageGroupAliases(&uptgroup, val.GroupId, spaces.DB)
	}

	/*Create Pages*/
	for _, val := range Pagec.NewPages {

		var newgrpid int
		newgrpid = val.Pgroupid
		if val.NewGrpId != 0 {

			for _, grp := range Pagec.NewGroup {

				if val.Pgroupid == grp.GroupId && val.NewGrpId == grp.NewGroupId {
					var getgid TblPagesGroupAliases
					Spacemodel.GetPageGroupByName(&getgid, spaceId, grp.Name, spaces.DB)
					newgrpid = getgid.PageGroupId

					break

				}
			}
		}

		/*page creation tbl_page*/
		var page TblPage
		page.PageGroupId = newgrpid
		page.SpacesId = spaceId
		page.ParentId = 0
		page.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		page.CreatedBy = userid
		pageret, _ := Spacemodel.CreatePage(&page, spaces.DB)

		for _, newval := range Pagec.NewSubPage {

			if newval.ParentId == 0 && newval.NewParentId == val.ParentId || newval.ParentId == 0 && newval.NewParentId == val.NewPgId || newval.ParentId == val.PgId && newval.NewParentId == 0 || newval.ParentId == val.NewPgId && newval.NewParentId == 0 {
				
				var Temarr TempCheck
				Temarr.FrontId = newval.SpgId
				Temarr.NewFrontId = newval.NewSpId
				Temarr.DBid = pageret.Id

				Temparr = append(Temparr, Temarr)

			}

		}

		/*page creation tbl_page_aliases*/
		var pageali Tblpagealiases
		pageali.LanguageId = 1
		pageali.PageId = pageret.Id
		pageali.PageTitle = val.Name
		pageali.PageDescription = val.Content
		pageali.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		pageali.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pageali.CreatedBy = userid
		pageali.OrderIndex = val.OrderIndex
		pageali.Status = status
		pageali.LastRevisionDate, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pageali.LastRevisionNo = 1
		pageali.Access = "public"
		pageali.ReadTime = val.ReadTime

		_, err = Spacemodel.CreatepageAliases(&pageali, spaces.DB)

		/*This is for log*/
		var pagelog TblPageAliasesLog
		pagelog.PageId = pageret.Id
		pagelog.LanguageId = 1
		pagelog.PageTitle = val.Name
		pagelog.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		pagelog.PageDescription = val.Content
		pagelog.PageId = pageret.Id
		pagelog.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pagelog.CreatedBy = userid
		pagelog.Status = status
		pagelog.Access = "public"
		pagelog.ReadTime = val.ReadTime

		Spacemodel.PageAliasesLog(&pagelog, spaces.DB)

	}

	/*Update Group*/
	for _, val := range Pagec.UpdatePages {
		
		newgrpid := val.Pgroupid
		
		if val.NewGrpId != 0 {

			for _, grp := range Pagec.NewGroup {

				if val.NewGrpId == grp.NewGroupId {
					var getgid TblPagesGroupAliases
					Spacemodel.GetPageGroupByName(&getgid, spaceId, grp.Name, spaces.DB)
					newgrpid = getgid.PageGroupId
					break

				}
			}

		}

		var uptpage TblPage
		uptpage.PageGroupId = newgrpid
		uptpage.ParentId = 0
		Spacemodel.UpdatePage(&uptpage, val.PgId, spaces.DB)

		var uptpageali TblPageAliases
		uptpageali.PageTitle = val.Name
		uptpageali.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		uptpageali.PageDescription = val.Content
		uptpageali.OrderIndex = val.OrderIndex
		uptpageali.Status = status
		uptpageali.ModifiedBy = userid
		uptpageali.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		uptpageali.LastRevisionDate, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		uptpageali.ReadTime = val.ReadTime

		_, err = Spacemodel.UpdatePageAliase(&uptpageali, val.PgId, spaces.DB)

		/*This is for log*/
		var pagelog TblPageAliasesLog
		pagelog.LanguageId = 1
		pagelog.PageId = val.PgId
		pagelog.PageTitle = val.Name
		pagelog.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		pagelog.PageDescription = val.Content
		pagelog.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pagelog.CreatedBy = userid
		pagelog.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pagelog.ModifiedBy = userid
		pagelog.Status = status
		pagelog.Access = "public"
		pagelog.ReadTime = val.ReadTime

		Spacemodel.PageAliasesLog(&pagelog, spaces.DB)

	}

	/*createsub*/
	for _, val := range Pagec.NewSubPage {

		var (
			newgrpid int
			pgid int
		)

		newgrpid = val.PgroupId
		pgid = val.ParentId
		newpgid := val.NewParentId

		if val.NewParentId != 0 {
			for _, pg := range Pagec.NewPages {
				if pg.PgId == pgid && pg.NewPgId == newpgid {
					var getpage TblPageAliases
					Spacemodel.GetPageDataByName(&getpage, spaceId, pg.Name, spaces.DB)
					pgid = getpage.PageId
					break
				}

			}
		}

		if val.NewPgroupId != 0 {
			for _, grp := range Pagec.NewGroup {
				if val.NewPgroupId == grp.NewGroupId {
					var getgid TblPagesGroupAliases
					Spacemodel.GetPageGroupByName(&getgid, spaceId, grp.Name, spaces.DB)
					newgrpid = getgid.PageGroupId
					break

				}
			}

		}

		/*page creation tbl_page*/
		var page TblPage
		page.PageGroupId = newgrpid
		page.SpacesId = spaceId
		page.ParentId = pgid
		page.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		page.CreatedBy = userid
		pageret, _ := Spacemodel.CreatePage(&page, spaces.DB)

		/*page creation tbl_page_aliases*/
		var pageali Tblpagealiases
		pageali.LanguageId = 1
		pageali.PageId = pageret.Id
		pageali.PageTitle = val.Name
		pageali.PageDescription = val.Content
		pageali.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		pageali.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pageali.CreatedBy = userid
		pageali.PageSuborder = val.OrderIndex
		pageali.Status = status
		pageali.Access = "public"
		pageali.LastRevisionDate, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pageali.LastRevisionNo = 1
		pageali.ReadTime = val.ReadTime

		_, err = Spacemodel.CreatepageAliases(&pageali, spaces.DB)

		/*This is for log*/
		var pagelog TblPageAliasesLog
		pagelog.LanguageId = 1
		pagelog.PageId = pageret.Id
		pagelog.PageTitle = val.Name
		pagelog.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		pagelog.PageDescription = val.Content
		pagelog.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pagelog.CreatedBy = userid
		pagelog.Status = status
		pagelog.Access = "public"
		pagelog.ReadTime = val.ReadTime

		Spacemodel.PageAliasesLog(&pagelog, spaces.DB)

	}

	/*update subpages*/
	for _, val := range Pagec.UpdateSubPage {

		newgrpid := val.PgroupId

		if val.NewPgroupId != 0 {

			for _, grp := range Pagec.NewGroup {

				if val.NewPgroupId == grp.NewGroupId {

					var getgid TblPagesGroupAliases

					Spacemodel.GetPageGroupByName(&getgid, spaceId, grp.Name, spaces.DB)

					newgrpid = getgid.PageGroupId

					break

				}
			}

		}

		var uptpage TblPage
		uptpage.PageGroupId = newgrpid
		uptpage.ParentId = val.ParentId
		Spacemodel.UpdatePage(&uptpage, val.SpgId, spaces.DB)

		var uptpageali TblPageAliases
		uptpageali.PageTitle = val.Name
		uptpageali.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		uptpageali.PageDescription = val.Content
		uptpageali.PageSuborder = val.OrderIndex
		uptpageali.Status = status
		uptpageali.ModifiedBy = userid
		uptpageali.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		uptpageali.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		uptpageali.ReadTime = val.ReadTime

		_, err = Spacemodel.UpdatePageAliase(&uptpageali, val.SpgId, spaces.DB)

		/*This is for log*/
		var pagelog TblPageAliasesLog
		pagelog.PageTitle = val.Name
		pagelog.LanguageId = 1
		pagelog.PageId = val.SpgId
		pagelog.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		pagelog.PageDescription = val.Content
		pagelog.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pagelog.CreatedBy = userid
		pagelog.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pagelog.ModifiedBy = userid
		pagelog.Status = status
		pagelog.Access = "public"
		pagelog.ReadTime = val.ReadTime

		Spacemodel.PageAliasesLog(&pagelog, spaces.DB)

	}

	/*DeleteFunc*/

	// var deleteGroup deleteGRP

	// json.Unmarshal([]byte(deletegroup), &deleteGroup)

	for _, val := range Pagec.DeleteGroup {

		var deletegroup TblPagesGroup
		deletegroup.DeletedBy = userid
		deletegroup.IsDeleted = 1
		deletegroup.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		Spacemodel.DeletePageGroup(&deletegroup, DeletePageGroupreq{GroupId: val.GroupId}, spaces.DB)

		var deletegroupali TblPagesGroupAliases
		deletegroupali.DeletedBy = userid
		deletegroupali.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		deletegroupali.IsDeleted = 1

		err = Spacemodel.DeletePageGroupAliases(&deletegroupali, DeletePageGroupreq{GroupId: val.GroupId}, spaces.DB)

	}

	for _, val := range Pagec.DeletePages {

		var deletegroup TblPage
		deletegroup.DeletedBy = userid
		deletegroup.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		Spacemodel.DeletePage(&deletegroup, DeletePagereq{Id: val.PgId}, spaces.DB)

		var deletegroupali TblPageAliases
		deletegroupali.DeletedBy = userid
		deletegroupali.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		deletegroupali.IsDeleted = 1

		err = Spacemodel.DeletePageAliase(&deletegroupali, DeletePagereq{Id: val.PgId}, spaces.DB)

		/*This is for log*/
		var pagelog TblPageAliasesLog
		pagelog.PageTitle = val.Name
		pagelog.LanguageId = 1
		pagelog.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		pagelog.PageDescription = val.Content
		pagelog.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pagelog.CreatedBy = userid
		pagelog.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pagelog.DeletedBy = userid
		pagelog.Status = status
		pagelog.Access = "public"

		Spacemodel.PageAliasesLog(&pagelog, spaces.DB)

	}

	// var deleteSubPage deleteSUB

	// json.Unmarshal([]byte(deletesub), &deleteSubPage)

	for _, val := range Pagec.DeleteSubPage {

		var deletegroup TblPage
		deletegroup.DeletedBy = userid
		deletegroup.IsDeleted = 1
		deletegroup.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		Spacemodel.DeletePage(&deletegroup, DeletePagereq{Id: val.SpgId}, spaces.DB)

		var deletegroupali TblPageAliases
		deletegroupali.DeletedBy = userid
		deletegroupali.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		deletegroupali.IsDeleted = 1
		err = Spacemodel.DeletePageAliase(&deletegroupali, DeletePagereq{Id: val.SpgId}, spaces.DB)

		/*This is for log*/
		var pagelog TblPageAliasesLog
		pagelog.PageTitle = val.Name
		pagelog.LanguageId = 1
		pagelog.PageSlug = strings.ToLower(strings.ReplaceAll(val.Name, " ", "_"))
		pagelog.PageDescription = val.Content
		pagelog.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pagelog.CreatedBy = userid
		pagelog.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
		pagelog.DeletedBy = userid
		pagelog.Status = status
		pagelog.Access = "public"

		Spacemodel.PageAliasesLog(&pagelog, spaces.DB)

	}

	if status == "publish" && len(Pagec.NewPages) == 0 && len(Pagec.NewSubPage) == 0 && len(Pagec.UpdatePages) == 0 && len(Pagec.UpdateSubPage) == 0 {

		var page []TblPage

		Spacemodel.SelectedPage(&page, spaceId, []int{}, spaces.DB)

		var id []int

		for _, val := range page {

			id = append(id, val.Id)

			var page Tblpagealiases
			Spacemodel.getPageAliases(&page, val.Id, spaces.DB)

			/*This is for log*/
			var pagelog TblPageAliasesLog
			pagelog.PageId = val.Id
			pagelog.PageTitle = page.PageTitle
			pagelog.LanguageId = 1
			pagelog.PageSlug = strings.ToLower(strings.ReplaceAll(page.PageTitle, " ", "_"))
			pagelog.PageDescription = page.PageDescription
			pagelog.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
			pagelog.CreatedBy = userid
			pagelog.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
			pagelog.ModifiedBy = userid
			pagelog.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
			pagelog.DeletedBy = userid
			pagelog.Status = "publish"
			pagelog.Access = "public"
			Spacemodel.PageAliasesLog(&pagelog, spaces.DB)

		}

		Spacemodel.UpdatePageAliasePublishStatus(id, userid, spaces.DB)

	}

	Temparr = nil

	if err != nil {

		return err
	}

	return nil
}
