package spaces

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

// create page
func (spaces *Spaces) CreatePage() {

}

//create group
func (spaces *Spaces) CreateGroup(){

}

//create subpage
func (spaces *Spaces) CreateSubpage(){

}

//update page
func (spaces *Spaces) UpdatePage(){

}

//update Group
func (spaces *Spaces) UpdateGroup(){

}

//update Subpage
func (spaces *Spaces) UpdateSubpage(){
	
}

//deletepage
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

//delete pagegroup
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


//deletepage
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