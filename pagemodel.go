package spaces

import (
	"strconv"
	"time"

	"github.com/spurtcms/member"
	"gorm.io/gorm"
)

func (SpaceModel) GetPageDetailsBySpaceId(getpg *[]TblPage, id int, DB *gorm.DB) (*[]TblPage, error) {

	if err := DB.Table("tbl_pages").Where("tbl_pages.is_deleted = ? and tbl_pages.spaces_id = ?", 0, id).Find(&getpg).Error; err != nil {

		return &[]TblPage{}, err
	}

	return getpg, nil
}

/*Get page log*/
func (SpaceModel) GetPageLogDetails(spaceid, pageid int, pageids []int, DB *gorm.DB) (tblpagelog []Tblpagealiaseslog, err error) {

	query := DB.Table("tbl_page_aliases_logs").Select("tbl_page_aliases_logs.created_by,tbl_page_aliases_logs.created_on,tbl_page_aliases_logs.status,tbl_users.username,max(tbl_page_aliases_logs.modified_by) as modified_by,max(tbl_page_aliases_logs.modified_on) as modified_on")

	query.Joins("inner join tbl_pages on tbl_pages.id = tbl_page_aliases_logs.page_id")

	query.Joins("inner join tbl_users on tbl_users.id = tbl_page_aliases_logs.created_by").Group("tbl_page_aliases_logs.created_by,tbl_page_aliases_logs.created_on,tbl_page_aliases_logs.status,tbl_users.username").Order("tbl_page_aliases_logs.created_on desc")

	if spaceid != 0 {

		query = query.Where("tbl_pages.spaces_id=?", spaceid)
	}

	if pageid != 0 {

		query = query.Where("tbl_pages.page_id=?", pageid)
	}

	query.Find(&tblpagelog)

	if err := query.Error; err != nil {

		return []Tblpagealiaseslog{}, err
	}

	return tblpagelog, nil
}

func (SpaceModel) SelectPage(pagereq GetPageReq, DB *gorm.DB) (tblpage []TblPage, singlepage TblPage, err error) {

	query := DB.Table("tbl_pages").Where("is_deleted =0 ")

	if len(pagereq.PageIds) != 0 {

		query = query.Where("id in (?)", pagereq.PageIds)

	}

	if pagereq.Spaceid != 0 {

		query = query.Where("spaces_id = ?", pagereq.Spaceid)
	}

	if pagereq.PageId != 0 {

		query = query.Where("id = ? ", pagereq.PageId)

		query.First(singlepage)

		if err := query.Error; err != nil {

			return []TblPage{}, TblPage{}, err

		}

		return []TblPage{}, singlepage, nil
	}

	query.Find(&tblpage)

	if err := query.Error; err != nil {

		return []TblPage{}, TblPage{}, err

	}

	return tblpage, TblPage{}, nil
}

func (SpaceModel) PageAliases(page GetPageReq, DB *gorm.DB) (tblpage []Tblpagealiases, singepage Tblpagealiases, err error) {

	query := DB.Table("tbl_page_aliases").Select("tbl_page_aliases.*,tbl_pages.parent_id,tbl_pages.page_group_id,tbl_users.username")

	query.Joins("inner join tbl_pages on tbl_pages.id = tbl_page_aliases.page_id")

	query.Joins("inner join tbl_users on tbl_users.id = tbl_page_aliases.created_by").Where("tbl_pages.is_deleted=0 and tbl_page_aliases.is_deleted=0")

	if page.Memberaccess {

		var mem member.TblMember

		DB.Model(member.TblMember{}).Where("is_deleted=0 and id=?", page.MemberId).First(&mem)

		if !page.ContentHideonly {

			subquery := `select tbl_access_control_pages.page_id from tbl_access_control_pages inner join tbl_access_control_pages on tbl_access_control_pages.access_control_user_group_id =tbl_access_control_user_groups.id Where member_group_id=` + strconv.Itoa(mem.MemberGroupId) + ` and tbl_access_control_user_groups.is_deleted=0`

			query = query.Where("tbl_page_aliases.id not in (?)", subquery)

		}

	}

	if page.PublishedPageonly {

		query = query.Where("tbl_page_aliases.status='publish'")
	}

	if page.Spaceid != 0 {

		query = query.Where("tbl_pages.spaces_id = ?", page.Spaceid)

	} else if len(page.PageIds) > 0 {

		query = query.Where("page_id in (?)", page.PageIds)

	} else if page.PageId != 0 {

		query = query.Where("page_id =?", page.PageId)

		query.First(&singepage)

		if err := query.Error; err != nil {

			return []Tblpagealiases{}, singepage, err

		}

		return tblpage, singepage, nil
	}

	query.Find(&tblpage)

	if err := query.Error; err != nil {

		return []Tblpagealiases{}, Tblpagealiases{}, err

	}

	return tblpage, singepage, nil
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

/*Delete PageAliases*/
func (SpaceModel) DeletePageAliase(tblpage *TblPageAliases, delpage DeletePagereq, DB *gorm.DB) error {

	query := DB.Model(TblPageAliases{})

	if delpage.SpaceId != 0 {

		subquery := `select id from tbl_pages where space_id = ` + strconv.Itoa(delpage.SpaceId) + ``

		query = query.Where("page_id=(?)", subquery)

	} else if len(delpage.GroupIds) != 0 {

		str := convertarrayintToString(delpage.GroupIds, ",")

		subquery := `select id from tbl_pages where page_group_id in (` + str + `)`

		query = query.Where("page_id=(?)", subquery)

	} else if delpage.GroupId != 0 {

		subquery := `select id from tbl_pages where page_group_id = ` + strconv.Itoa(delpage.GroupId) + ``

		query = query.Where("page_id = (?)", subquery)

	} else if len(delpage.Ids) != 0 {

		query.Where("page_id in (?)", delpage.Ids)

	} else if delpage.Id != 0 {

		query.Where("page_id=?", delpage.Id)

	}

	query.UpdateColumns(map[string]interface{}{"deleted_on": tblpage.DeletedOn, "deleted_by": tblpage.DeletedBy, "is_deleted": 1})

	if err := query.Error; err != nil {

		return err
	}

	return nil

}

/*Delete PageAliases*/
func (SpaceModel) DeletePage(tblpage *TblPage, delpage DeletePagereq, DB *gorm.DB) error {

	query := DB.Model(TblPage{})

	if delpage.SpaceId != 0 {

	} else if len(delpage.GroupIds) != 0 {

		query = query.Where("space_id=?", delpage.SpaceId)

	} else if delpage.GroupId != 0 {

		query = query.Where("page_group_id=?", delpage.GroupId)

	} else if len(delpage.Ids) != 0 {

		query.Where("id in (?)", delpage.Ids)

	} else if delpage.Id != 0 {

		query.Where("id=?", delpage.Id)

	}

	query.UpdateColumns(map[string]interface{}{"deleted_on": tblpage.DeletedOn, "deleted_by": tblpage.DeletedBy, "is_deleted": 1})

	if err := query.Error; err != nil {

		return err
	}

	return nil

}

/* Delete group */
func (SpaceModel) DeletePageGroup(tblpagegroup *TblPagesGroup, delpgg DeletePageGroupreq, DB *gorm.DB) error {

	query := DB.Model(TblPagesGroup{})

	if delpgg.SpaceId != 0 {

		query = query.Where("space_id=?", delpgg.SpaceId)

	} else if len(delpgg.GroupIds) != 0 {

		str := convertarrayintToString(delpgg.GroupIds, ",")

		subquery := `select id from tbl_page_groups where id in (` + str + `)`

		query = query.Where("id in (?)", subquery)

	} else if delpgg.GroupId != 0 {

		query = query.Where("id=?", delpgg.GroupId)

	}

	query.UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": tblpagegroup.DeletedOn, "deleted_by": tblpagegroup.DeletedBy})

	if err := query.Error; err != nil {

		return err

	}

	return nil
}

/* Delete Groupaliases */
func (SpaceModel) DeletePageGroupAliases(tblpagegroup *TblPagesGroupAliases, delpgg DeletePageGroupreq, DB *gorm.DB) error {

	query := DB.Model(TblPagesGroupAliases{})

	if delpgg.SpaceId != 0 {

		query = query.Where("space_id=?", delpgg.SpaceId)

	} else if len(delpgg.GroupIds) != 0 {

		str := convertarrayintToString(delpgg.GroupIds, ",")

		subquery := `select id from tbl_page_group_aliases where id in (` + str + `)`

		query = query.Where("page_group_id in (?)", subquery)

	} else if delpgg.GroupId != 0 {

		query = query.Where("page_group_id=?", delpgg.GroupId)

	}

	query.UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": tblpagegroup.DeletedOn, "deleted_by": tblpagegroup.DeletedBy})

	if err := query.Error; err != nil {

		return err

	}

	return nil
}

func (SpaceModel) CreatePageGroup(tblpagegroup *TblPagesGroup, DB *gorm.DB) (*TblPagesGroup, error) {

	if err := DB.Table("tbl_pages_groups").Create(&tblpagegroup).Error; err != nil {

		return &TblPagesGroup{}, err
	}

	return tblpagegroup, nil

}

/*Create PagegroupAliases */
func (SpaceModel) CreatePageGroupAliases(tblpagegroup *TblPagesGroupAliases, DB *gorm.DB) (TblPagesGroupAliases, error) {

	if err := DB.Table("tbl_pages_group_aliases").Create(&tblpagegroup).Error; err != nil {

		return *tblpagegroup, err
	}

	return *tblpagegroup, nil
}

/*pdate pagegroupAliases */
func (SpaceModel) UpdatePageGroupAliases(tblpagegroup *TblPagesGroupAliases, id int, DB *gorm.DB) (TblPagesGroupAliases, error) {

	query := DB.Table("tbl_pages_group_aliases").Where("page_group_id = ?", id).UpdateColumns(map[string]interface{}{"group_name": tblpagegroup.GroupName, "group_slug": tblpagegroup.GroupSlug, "group_description": tblpagegroup.GroupDescription, "language_id": tblpagegroup.LanguageId, "modified_on": tblpagegroup.ModifiedOn, "modified_by": tblpagegroup.ModifiedBy})

	if err := query.Error; err != nil {

		return TblPagesGroupAliases{}, err
	}

	return *tblpagegroup, nil
}

/*Create page log*/
func (SpaceModel) PageAliasesLog(tblpagelog *TblPageAliasesLog, DB *gorm.DB) (TblPageAliasesLog, error) {

	if err := DB.Table("tbl_page_aliases_logs").Create(&tblpagelog).Error; err != nil {

		return TblPageAliasesLog{}, err
	}

	return *tblpagelog, nil
}

/*CreatePage*/
func (SpaceModel) CreatePage(tblpage *TblPage, DB *gorm.DB) (*TblPage, error) {

	if err := DB.Table("tbl_pages").Create(&tblpage).Error; err != nil {

		return &TblPage{}, err
	}
	return tblpage, nil

}

// create PageAliases
func (SpaceModel) CreatepageAliases(tblpageAliases *Tblpagealiases, DB *gorm.DB) (Tblpagealiases, error) {

	if err := DB.Table("tbl_page_aliases").Create(&tblpageAliases).Error; err != nil {

		return *tblpageAliases, err
	}

	return *tblpageAliases, nil

}

/*update page*/
func (SpaceModel) UpdatePage(tblpage *TblPage, pageid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_pages").Where("id=?", pageid).UpdateColumns(map[string]interface{}{"page_group_id": tblpage.PageGroupId, "parent_id": tblpage.ParentId}).Error; err != nil {

		return err
	}

	return nil
}

/*update pagealiases*/
func (SpaceModel) UpdatePageAliase(tblpageali *TblPageAliases, pageid int, DB *gorm.DB) (TblPageAliases, error) {

	query := DB.Table("tbl_page_aliases").Where("page_id=?", pageid).UpdateColumns(map[string]interface{}{
		"page_title": tblpageali.PageTitle, "page_slug": tblpageali.PageSlug, "modified_on": tblpageali.ModifiedOn,
		"modified_by": tblpageali.ModifiedBy, "page_description": tblpageali.PageDescription, "order_index": tblpageali.OrderIndex, "status": tblpageali.Status, "read_time": tblpageali.ReadTime})

	if err := query.Error; err != nil {

		return TblPageAliases{}, err

	}

	return *tblpageali, nil
}

/*PageGroup*/
func (SpaceModel) GetPageGroupByName(TblPagesGroupAliases *TblPagesGroupAliases, spaceid int, name string, DB *gorm.DB) error {

	if err := DB.Table("tbl_pages_group_aliases").Joins("inner join tbl_pages_groups on tbl_pages_groups.id=tbl_pages_group_aliases.page_group_id").Where("LOWER(TRIM(group_name))=LOWER(TRIM(?)) and tbl_pages_groups.spaces_id=? and tbl_pages_group_aliases.is_deleted=0", name, spaceid).Last(&TblPagesGroupAliases).Error; err != nil {

		return err
	}

	return nil
}

/*GetPage*/
func (SpaceModel) GetPageDataByName(TblPageAliases *TblPageAliases, spaceid int, name string, DB *gorm.DB) error {

	if err := DB.Table("tbl_page_aliases").Select("tbl_page_aliases.*").Joins("inner join tbl_pages on tbl_pages.id=tbl_page_aliases.page_id").Where("page_title=? and tbl_pages.spaces_id=? and tbl_page_aliases.is_deleted=0", name, spaceid).Last(&TblPageAliases).Error; err != nil {

		return err
	}

	return nil
}

/*update pagealiases*/
func (SpaceModel) UpdatePageAliasePublishStatus(pageid []int, userid int, DB *gorm.DB) error {

	Formatdate, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	if err := DB.Table("tbl_page_aliases").Where("page_id in (?)", pageid).UpdateColumns(map[string]interface{}{"status": "publish", "modified_on": Formatdate, "modified_by": userid}).Error; err != nil {
		return err
	}

	return nil
}

/*Get page log*/
func (SpaceModel) GetPageLogDetailsByPageId(tblpagelog *[]Tblpagealiaseslog, spaceid int, pageid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_page_aliases_logs").Select("tbl_page_aliases_logs.created_by,tbl_page_aliases_logs.created_on,tbl_page_aliases_logs.status,tbl_users.username,max(tbl_page_aliases_logs.modified_by) as modified_by,max(tbl_page_aliases_logs.modified_on) as modified_on").Joins("inner join tbl_pages on tbl_pages.id = tbl_page_aliases_logs.page_id").Joins("inner join tbl_users on tbl_users.id = tbl_page_aliases_logs.created_by").Where("tbl_pages.spaces_id=? and page_id = ? ", spaceid, pageid).Group("tbl_page_aliases_logs.created_by,tbl_page_aliases_logs.created_on,tbl_page_aliases_logs.status,tbl_users.username").Order("tbl_page_aliases_logs.created_on desc").Find(&tblpagelog).Error; err != nil {

		return err
	}

	return nil
}
func (SpaceModel) getPageAliases(tblpagegroup *Tblpagealiases, id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_page_aliases").Select("tbl_page_aliases.*,tbl_pages.page_group_id,tbl_users.username").Joins("inner join tbl_pages on tbl_pages.id = tbl_page_aliases.page_id").Joins("inner join tbl_users on tbl_users.id = tbl_page_aliases.created_by").Where("page_id = ? and tbl_pages.is_deleted=0 and tbl_page_aliases.is_deleted=0", id).Find(&tblpagegroup).Error; err != nil {

		return err

	}

	return nil
}

func (SpaceModel) SelectedGroup(tblgroup *[]TblPagesGroup, id int, grpid []int, DB *gorm.DB) error {

	query := DB.Table("tbl_pages_groups").Where("spaces_id = ? and is_deleted=0", id)

	query.Find(&tblgroup)

	if err := query.Error; err != nil {

		return err

	}

	return nil
}

func (SpaceModel) SelectedPage(tblpage *[]TblPage, id int, pgid []int, DB *gorm.DB) error {

	query := DB.Table("tbl_pages").Where("spaces_id = ? and is_deleted =0 ", id)

	query.Find(&tblpage)

	if err := query.Error; err != nil {

		return err

	}

	return nil
}