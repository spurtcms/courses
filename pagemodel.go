package spaces

import (
	"gorm.io/gorm"
)

// delete page group
func (SpaceModel) DeletePageGroup(tblpage *TblPagesGroup, id int, DB *gorm.DB) error {

	if err := DB.Model(TblPagesGroup{}).Where("tbl_pages_groups.id=?", id).UpdateColumns(map[string]interface{}{"is_deleted": tblpage.IsDeleted, "deleted_on": tblpage.DeletedOn, "deleted_by": tblpage.DeletedBy}).Error; err != nil {

		return err
	}

	return nil
}

func (SpaceModel) GetPageDetailsBySpaceId(getpg *[]TblPage, id int, DB *gorm.DB) (*[]TblPage, error) {

	if err := DB.Table("tbl_pages").Where("tbl_pages.is_deleted = ? and tbl_pages.spaces_id = ?", 0, id).Find(&getpg).Error; err != nil {

		return &[]TblPage{}, err
	}

	return getpg, nil
}

/*Get page log*/
func (SpaceModel) GetPageLogDetails(spaceid, pageid int, pageids []int, DB *gorm.DB) (tblpagelog []Tblpagealiaseslog, err error) {

	query := DB.Table("tbl_page_aliases_logs").Select("tbl_page_aliases_logs.created_by,tbl_page_aliases_logs.created_on,tbl_page_aliases_logs.status,tbl_users.username,max(tbl_page_aliases_logs.modified_by) as modified_by,max(tbl_page_aliases_logs.modified_on) as modified_on").Joins("inner join tbl_pages on tbl_pages.id = tbl_page_aliases_logs.page_id").Joins("inner join tbl_users on tbl_users.id = tbl_page_aliases_logs.created_by").Group("tbl_page_aliases_logs.created_by,tbl_page_aliases_logs.created_on,tbl_page_aliases_logs.status,tbl_users.username").Order("tbl_page_aliases_logs.created_on desc").Find(&tblpagelog)

	if spaceid != 0 {

		query = query.Where("tbl_pages.spaces_id=?", spaceid)
	}

	if pageid != 0 {

		query = query.Where("tbl_pages.page_id=?", pageid)
	}

	if err := query.Error; err != nil {

		return []Tblpagealiaseslog{}, err
	}

	return tblpagelog, nil
}

func (SpaceModel) SelectPage(pagereq GetPageReq, DB *gorm.DB) (tblpage []TblPage, err error) {

	query := DB.Table("tbl_pages").Where("is_deleted =0 ")

	if len(pagereq.PageIds) != 0 {

		query = query.Where("id in (?)", pagereq.PageIds)

	}

	if pagereq.Spaceid != 0 {

		query = query.Where("spaces_id = ?", pagereq.Spaceid)
	}

	if pagereq.PageId != 0 {

		query = query.Where("id = ? ", pagereq.PageId)
	}

	query.Find(&tblpage)

	if err := query.Error; err != nil {

		return []TblPage{}, err

	}

	return tblpage, nil
}

func (SpaceModel) PageAliases(ids []int, id int, memberaccess bool, memberid int, DB *gorm.DB) (tblpagegroup Tblpagealiases, err error) {

	query := DB.Table("tbl_page_aliases").Select("tbl_page_aliases.*,tbl_pages.page_group_id,tbl_users.username")

	query.Joins("inner join tbl_pages on tbl_pages.id = tbl_page_aliases.page_id")

	query.Joins("inner join tbl_users on tbl_users.id = tbl_page_aliases.created_by").Where("tbl_pages.is_deleted=0 and tbl_page_aliases.is_deleted=0")

	if memberaccess {

		
	}

	if len(ids) > 0 {

		query = query.Where("page_id in (?)", ids)
	}

	if id != 0 {

		query = query.Where("page_id =?", id)
	}

	query.Find(&tblpagegroup)

	if err := query.Error; err != nil {

		return Tblpagealiases{}, err

	}

	return tblpagegroup, nil
}

func (SpaceModel) SelectGroup(tblgroup *[]TblPagesGroup, id int, grpid []int, DB *gorm.DB) error {

	query := DB.Table("tbl_pages_groups").Where("spaces_id = ? and is_deleted=0", id)

	if len(grpid) != 0 {

		query = query.Where("id in (?)", grpid)

	}

	query.Find(&tblgroup)

	if err := query.Error; err != nil {

		return err

	}

	return nil
}

func (SpaceModel) PageGroup(tblpagegroup *TblPagesGroupAliases, id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_pages_group_aliases").Where("is_deleted = 0 and page_group_id = ?", id).First(&tblpagegroup).Error; err != nil {

		return err

	}

	return nil
}
