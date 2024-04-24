package spaces

import (
	"fmt"
	"time"
)

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

// get page list
func (spaces *Spaces) GetPage(pagereq GetPageReq) ([]Pages, error) {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return []Pages{}, autherr
	}

	var pages []Pages

	page, err := Spacemodel.SelectPage(pagereq, spaces.DB)

	if err != nil {

		return []Pages{}, err
	}

	if pagereq.Spaceid != 0 {

		var ids []int

		for _, page := range page {

			ids = append(ids, page.Id)

		}

		pagelog, err := Spacemodel.GetPageLogDetails(pagereq.Spaceid, 0, ids, spaces.DB)

		if err != nil {

			fmt.Println(err)
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

		var one_page Pages

		page_aliases, _ := Spacemodel.PageAliases(ids, 0, pagereq.Memberaccess, pagereq.MemberId, spaces.DB)

		one_page.PgId = page_aliases.PageId

		one_page.Name = page_aliases.PageTitle

		one_page.Content = page_aliases.PageDescription

		one_page.OrderIndex = page_aliases.OrderIndex

		one_page.Pgroupid = page_aliases.PageGroupId

		one_page.ParentId = page_aliases.ParentId

		one_page.CreatedDate = page_aliases.CreatedOn

		one_page.LastUpdate = page_aliases.ModifiedOn

		one_page.Username = page_aliases.Username

		one_page.ReadTime = page_aliases.ReadTime

		one_page.Log = finallog

		pages = append(pages, one_page)

	} else if len(pagereq.PageIds) > 0 {

	} else if pagereq.PageId != 0 {

	}

	return pages, nil

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
func (spaces *Spaces) GetSubPages(pagereq GetPageReq) ([]SubPages, error) {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return []SubPages{}, autherr
	}

	var subpages []SubPages

	page, err := Spacemodel.SelectPage(pagereq, spaces.DB)

	if err != nil {

		return []SubPages{}, err
	}

	for _, page := range page {

		pagelog, err := Spacemodel.GetPageLogDetails(pagereq.Spaceid, page.Id, []int{}, spaces.DB)

		if err != nil {

			fmt.Println(err)
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

		if page.ParentId != 0 {

			sid := page.Id

			var subpage SubPages

			page_aliases, _ := Spacemodel.PageAliases([]int{}, sid, pagereq.Memberaccess, pagereq.MemberId, spaces.DB)

			subpage.SpgId = page_aliases.PageId

			subpage.Name = page_aliases.PageTitle

			subpage.Content = page_aliases.PageDescription

			subpage.ParentId = page.ParentId

			subpage.OrderIndex = page_aliases.PageSuborder

			subpage.CreatedDate = page_aliases.CreatedOn

			subpage.LastUpdate = page_aliases.ModifiedOn

			subpage.Username = page_aliases.Username

			subpage.ReadTime = page_aliases.ReadTime

			subpage.Log = finallog

			subpages = append(subpages, subpage)

		}

	}

	return subpages, nil

}

// delete page aliases
func (spaces *Spaces) DeletePageAliases(delpage DeletePagereq) error {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return autherr
	}

	var pageali Tblpagealiases

	pageali.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	pageali.DeletedBy = delpage.DeletedBy

	pageali.IsDeleted = 1

	return nil

}

// delete pagegroup
func (spaces *Spaces) DeletedPageGroupAliases(delpage DeletePagereq) error {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return autherr
	}

	return nil
}

func (spaces *Spaces) DeletePage(delpage DeletePagereq) error {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return autherr
	}

	if len(delpage.SpaceIds) != 0 {

	} else if delpage.SpaceId != 0 {

	} else if len(delpage.GroupIds) != 0 {

	} else if delpage.GroupId != 0 {

	} else if len(delpage.Ids) != 0 {

	} else if delpage.Id != 0 {

	}

	return nil
}

func (spaces *Spaces) DeletedPageGroup(delpgg DeletePageGroupreq) error {

	autherr := AuthandPermission(spaces)

	if autherr != nil {

		return autherr
	}

	if len(delpgg.SpaceIds) != 0 {

	} else if delpgg.SpaceId != 0 {

	} else if len(delpgg.GroupIds) != 0 {

	} else if delpgg.GroupId != 0 {

	}

	return nil

}
