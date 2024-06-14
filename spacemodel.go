package spaces

import (
	"time"

	"gorm.io/gorm"
)

/*spaceList*/
func (SpaceModel) SpaceList(spacereq SpaceListReq, spaceid []int, DB *gorm.DB) (tblspace []Tblspacesaliases, spacecount int64, err error) {

	query := DB.Table("tbl_spaces_aliases").Select("tbl_spaces_aliases.*,tbl_spaces.page_category_id,tbl_categories.parent_id").
		Joins("inner join tbl_spaces on tbl_spaces_aliases.spaces_id = tbl_spaces.id").
		Joins("inner join tbl_languages on tbl_languages.id = tbl_spaces_aliases.language_id").
		Joins("inner join tbl_categories on tbl_categories.id = tbl_spaces.page_category_id").
		Where("tbl_spaces.is_deleted = 0 and tbl_spaces_aliases.is_deleted = 0").Order("tbl_spaces.id desc")

	if spacereq.LanguageEnable {

		query = query.Where("tbl_spaces_aliases.language_id = ?", spacereq.SetLanguageId)
	}

	if spacereq.MemberAccessControl {

		subquery := DB.Table("tbl_access_control_pages").Select("tbl_access_control_pages.entry_id").Joins("inner join tbl_access_control_user_group on tbl_access_control_user_group.id = tbl_access_control_pages.access_control_user_group_id").
			Joins("inner join tbl_member_groups on tbl_member_groups.id = tbl_access_control_user_group.member_group_id").Joins("inner join tbl_members on tbl_members.member_group_id = tbl_member_groups.id")

		innerSubQuery := DB.Table("tbl_channel_entries").Select("tbl_channel_entries.channel_id").Joins("inner join tbl_channels on tbl_channels.id = tbl_channel_entries.channel_id").Joins("inner join tbl_spaces on tbl_spaces.id = tbl_spaces_aliases.spaces_id").Where("tbl_spaces.is_deleted = 0 and tbl_spaces_aliases.is_deleted = 0")

		subquery = subquery.Where("tbl_access_control_pages.spaces_id in (?)", innerSubQuery)

		query = query.Where("tbl_spaces_aliases.spaces_id not in (?)", subquery)

	}

	if len(spaceid) != 0 {

		query = query.Where("tbl_spaces.id in (?)", spaceid)
	}

	if spacereq.Keyword != "" {

		query = query.Where("LOWER(TRIM(tbl_spaces_aliases.spaces_name)) LIKE LOWER(TRIM(?))", "%"+spacereq.Keyword+"%")
	}
	if spacereq.CategoryId > 0 && spacereq.CategoryId != 0 {

		query = query.Where("tbl_spaces.page_category_id IN (?)", spacereq.CategoryId)
	}

	if spacereq.Limit != 0 {

		query.Limit(spacereq.Limit).Offset(spacereq.Offset).Order("tbl_spaces.id desc").Find(&tblspace)

	} else {

		query.Find(&tblspace).Count(&spacecount)

		return tblspace, spacecount, nil
	}

	return tblspace, 0, nil
}

// get last update
func (SpaceModel) GetLastUpdatePageAliases(tblpageali *TblPageAliases, spaceid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_page_aliases").Select("max(tbl_page_aliases.modified_on) as modified_on").Joins("inner join tbl_pages on tbl_pages.Id = tbl_page_aliases.page_id").Where("tbl_pages.spaces_Id=?", spaceid).Group("tbl_page_aliases.id").First(tblpageali).Error; err != nil {
		return err
	}

	return nil
}

func (SpaceModel) GetSpacealiaseDetails(spaceid int, spaceslug string, DB *gorm.DB) (TblSpacesAliase Tblspacesaliases, err error) {

	query := DB.Table("tbl_spaces_aliases").First(&TblSpacesAliase)

	if spaceid > 0 {

		query = query.Where("spaces_id=?", spaceid)

	}

	if spaceslug != "" {

		query = query.Where("spaces_slug=?", spaceid)
	}

	if err := query.Error; err != nil {

		return Tblspacesaliases{}, err
	}

	return TblSpacesAliase, nil
}

func (SpaceModel) GetSpaceDetails(id int, DB *gorm.DB) (tblspace tblspaces, err error) {

	if err := DB.Table("tbl_spaces").Select("tbl_spaces.created_on,tbl_spaces.modified_on,tbl_users.username").Where("tbl_spaces.id=?", id).Joins("inner join tbl_users on tbl_users.id = tbl_spaces.created_by").First(&tblspace).Error; err != nil {

		return tblspaces{}, err
	}

	return tblspace, nil
}

func (SpaceModel) CreateSpace(tblspac tblspaces, DB *gorm.DB) (tblspace tblspaces, err error) {

	if err := DB.Table("tbl_spaces").Create(&tblspac).Error; err != nil {

		return tblspaces{}, err
	}

	return tblspac, nil
}

func (SpaceModel) CreateSpaceAliase(tblspac Tblspacesaliases, DB *gorm.DB) (tblspc Tblspacesaliases, err error) {

	if err := DB.Table("tbl_spaces_aliases").Create(&tblspac).Error; err != nil {

		return Tblspacesaliases{}, err
	}

	return tblspac, nil
}

/*Update Space*/
func (SpaceModel) UpdateSpaceAliases(tblspace *Tblspacesaliases, id int, DB *gorm.DB) error {

	DB.Table("tbl_spaces_aliases").Where("spaces_id = ?", tblspace.Id).UpdateColumns(map[string]interface{}{"spaces_name": tblspace.SpacesName, "spaces_description": tblspace.SpacesDescription, "spaces_slug": tblspace.SpacesSlug, "image_path": tblspace.ImagePath, "modified_by": tblspace.ModifiedBy, "modified_on": tblspace.ModifiedOn})

	return nil
}

/*Update Space*/
func (SpaceModel) UpdateSpace(tblspace *tblspaces, id int, DB *gorm.DB) error {

	if tblspace.PageCategoryId != 0 {

		DB.Table("tbl_spaces").Where("id = ?", tblspace.Id).UpdateColumns(map[string]interface{}{"page_category_id": tblspace.PageCategoryId, "modified_by": tblspace.ModifiedBy, "modified_on": tblspace.ModifiedOn})

	} else {

		DB.Table("tbl_spaces").Where("id = ?", tblspace.Id).UpdateColumns(map[string]interface{}{"modified_by": tblspace.ModifiedBy, "modified_on": tblspace.ModifiedOn})

	}
	return nil
}

/*Deleted space*/
func (SpaceModel) DeleteSpaceAliases(tblspace *TblSpacesAliases, id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_spaces_aliases").Where("spaces_id = ?", id).UpdateColumns(map[string]interface{}{"deleted_by": tblspace.DeletedBy, "deleted_on": tblspace.DeletedOn, "is_deleted": tblspace.IsDeleted}).Error; err != nil {

		return err
	}

	return nil
}

/*Deleted space*/
func (SpaceModel) DeleteSpace(tblspace *TblSpaces, id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_spaces").Where("id = ?", id).UpdateColumns(map[string]interface{}{"deleted_by": tblspace.DeletedBy, "deleted_on": tblspace.DeletedOn, "is_deleted": tblspace.IsDeleted}).Error; err != nil {

		return err
	}

	return nil
}

func (SpaceModel) PageCount(DB *gorm.DB) (count int64, err error) {
	if err := DB.Table("tbl_page_aliases").Where("is_deleted = 0").Count(&count).Error; err != nil {

		return 0, err
	}

	return count, nil
}

func (SpaceModel) NewpageCount(DB *gorm.DB) (count int64, err error) {

	if err := DB.Table("tbl_page_aliases").Where("is_deleted = 0 AND created_on >=?", time.Now().AddDate(0, 0, -10)).Count(&count).Error; err != nil {

		return 0, err
	}

	return count, nil
}

// Mostlyviewed List//

func (SpaceModel) MostlyViewList(Space *[]Tblspacesaliases, limit int, DB *gorm.DB) (err error) {

	query := DB.Table("tbl_spaces_aliases").Select("tbl_spaces_aliases.*,tbl_spaces.page_category_id,tbl_categories.parent_id").
		Joins("inner join tbl_spaces on tbl_spaces_aliases.spaces_id = tbl_spaces.id").
		Joins("inner join tbl_languages on tbl_languages.id = tbl_spaces_aliases.language_id").
		Joins("inner join tbl_categories on tbl_categories.id = tbl_spaces.page_category_id").
		Where("tbl_spaces.is_deleted = 0 and tbl_spaces_aliases.is_deleted = 0 and tbl_spaces_aliases.language_id = 1 and tbl_spaces_aliases.view_count!=0").Order("tbl_spaces_aliases.view_count desc").Limit(limit)

	query.Find(&Space)

	return nil

}

func (SpaceModel) RecentlyViewList(Space *[]Tblspacesaliases, limit int, DB *gorm.DB) (err error) {

	query := DB.Table("tbl_spaces_aliases").Select("tbl_spaces_aliases.*,tbl_spaces.page_category_id,tbl_categories.parent_id").
		Joins("inner join tbl_spaces on tbl_spaces_aliases.spaces_id = tbl_spaces.id").
		Joins("inner join tbl_languages on tbl_languages.id = tbl_spaces_aliases.language_id").
		Joins("inner join tbl_categories on tbl_categories.id = tbl_spaces.page_category_id").
		Where("tbl_spaces.is_deleted = 0 and tbl_spaces_aliases.is_deleted = 0 and tbl_spaces_aliases.language_id = 1 and tbl_spaces_aliases.view_count!=0").Order("tbl_spaces_aliases.recent_time desc").Limit(limit)

	query.Find(&Space)

	return nil

}

// Name already exists
func (SpaceModel) CheckSpaceName(space *Tblspacesaliases, userid int, name string, DB *gorm.DB) error {

	if userid == 0 {

		if err := DB.Table("tbl_spaces_aliases").Where("LOWER(TRIM(spaces_name))=LOWER(TRIM(?)) and is_deleted=0", name).First(&space).Error; err != nil {

			return err
		}
	} else {

		if err := DB.Table("tbl_spaces_aliases").Where("LOWER(TRIM(spaces_name))=LOWER(TRIM(?)) and spaces_id not in (?) and is_deleted=0", name, userid).First(&space).Error; err != nil {

			return err
		}
	}

	return nil
}
func (SpaceModel) GetPageGroupData(group *[]TblPagesGroupAliases, id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_pages_group_aliases").Joins("inner join tbl_pages_groups on tbl_pages_group_aliases.page_group_id = tbl_pages_groups.id").Where("tbl_pages_groups.is_deleted = ? and tbl_pages_group_aliases.is_deleted = ? and tbl_pages_groups.spaces_id = ?", 0, 0, id).Find(&group).Error; err != nil {

		return err

	}

	return nil

}
func (SpaceModel) GetPageInPage(pageid *[]Tblpagealiases, spaceid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_page_aliases").Select("tbl_page_aliases.*").Joins("inner join tbl_pages on tbl_page_aliases.page_id = tbl_pages.id").Where("tbl_pages.is_deleted = ? and tbl_page_aliases.is_deleted = ? and tbl_pages.page_group_id = ? and  parent_id = ? and  tbl_pages.spaces_id = ?", 0, 0, 0, 0, spaceid).Find(&pageid).Error; err != nil {

		return err

	}

	return nil
}
func (SpaceModel) GetIdInPage(pageid *TblPagesGroupAliases, spaceid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_pages_group_aliases").Joins("inner join tbl_pages on tbl_pages_group_aliases.page_group_id = tbl_pages.page_group_id").Where("tbl_pages.is_deleted = ? and tbl_pages_group_aliases.is_deleted = ? and tbl_pages.page_group_id != ? and  tbl_pages.parent_id = ? and  tbl_pages.spaces_id = ?", 0, 0, 0, 0, spaceid).First(&pageid).Error; err != nil {
		return err

	}

	return nil
}
func (SpaceModel) GetPageAliasesInPage(data *[]Tblpagealiases, spacid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_page_aliases").Select("tbl_page_aliases.*,tbl_pages.parent_id,tbl_pages.page_group_id,tbl_pages.spaces_id").Joins("inner join tbl_pages on tbl_page_aliases.page_id = tbl_pages.id").Where("tbl_page_aliases.is_deleted = ? and tbl_pages.is_deleted = ? and tbl_pages.page_group_id != ? and  tbl_pages.parent_id = ? and  tbl_pages.spaces_id = ?", 0, 0, 0, 0, spacid).Find(&data).Error; err != nil {

		return err

	}

	return nil
}
func (SpaceModel) GetDetailsInPageAli(pagedetails *TblPagesGroupAliases, groupname string, spaceid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_pages_group_aliases").Joins("inner join tbl_pages_groups on tbl_pages_group_aliases.page_group_id = tbl_pages_groups.id").Where("tbl_pages_groups.is_deleted = ? and  tbl_pages_group_aliases.is_deleted = ? and  tbl_pages_group_aliases.group_name = ? and tbl_pages_groups.spaces_id = ? ", 0, 0, groupname, spaceid).Find(&pagedetails).Error; err != nil {

		return err

	}

	return nil
}
func (SpaceModel) GetPageAliasesInPageData(result *[]Tblpagealiases, spacid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_page_aliases").Select("tbl_page_aliases.*,tbl_pages.page_group_id,tbl_pages.parent_id").Joins("inner join tbl_pages on tbl_page_aliases.page_id = tbl_pages.id").Where("tbl_pages.parent_id != ? and  tbl_pages.spaces_id = ?", 0, spacid).Find(&result).Error; err != nil {

		return err

	}

	return nil
}
func (SpaceModel) GetDetailsInPageAlia(pageid *TblPagesGroupAliases, pagegroupid int, spaceid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_pages_group_aliases").Joins("inner join tbl_pages_groups on tbl_pages_group_aliases.page_group_id = tbl_pages_groups.id").Where("tbl_pages_group_aliases.page_group_id = ? and  tbl_pages_groups.spaces_id = ?", pagegroupid, spaceid).First(&pageid).Error; err != nil {

		return err

	}

	return nil
}
func (SpaceModel) AliasesInParentId(data *TblPageAliases, parentid int, spacid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_page_aliases").Select("tbl_page_aliases.*,tbl_pages.id,tbl_pages.parent_id").Joins("inner join tbl_pages on tbl_page_aliases.page_id = tbl_pages.id").Where("tbl_page_aliases.page_id = ? and  tbl_pages.spaces_id = ?", parentid, spacid).First(&data).Error; err != nil {

		return err

	}

	return nil
}
func (SpaceModel) LastLoopAliasesInPage(data *TblPageAliases, pagetitle string, spacid int, DB *gorm.DB) error {

	if err := DB.Table("tbl_page_aliases").Select("tbl_page_aliases.*").Joins("inner join tbl_pages on tbl_page_aliases.page_id = tbl_pages.id").Where("tbl_page_aliases.page_title = ? and  tbl_pages.spaces_id = ?", pagetitle, spacid).First(&data).Error; err != nil {

		return err

	}

	return nil
}

func (SpaceModel) DeletePageInSpace(page *TblPage, id []int, DB *gorm.DB) error {

	if err := DB.Table("tbl_pages").Where("tbl_pages.id IN ?", id).UpdateColumns(map[string]interface{}{"deleted_by": page.DeletedBy, "deleted_on": page.DeletedOn, "is_deleted": page.IsDeleted}).Error; err != nil {

		return err
	}

	return nil
}
func (SpaceModel) DeletePageAliInSpace(pageali *TblPageAliases, id []int, DB *gorm.DB) error {

	if err := DB.Table("tbl_page_aliases").Where("tbl_page_aliases.page_id IN ?", id).UpdateColumns(map[string]interface{}{"deleted_by": pageali.DeletedBy, "deleted_on": pageali.DeletedOn, "is_deleted": pageali.IsDeleted}).Error; err != nil {

		return err
	}

	return nil
}

// get pagegroup data by pass spaceid
func (SpaceModel) GetPageGroupDetailsBySpaceId(getpagegrp *[]TblPagesGroup, id int, DB *gorm.DB) (*[]TblPagesGroup, error) {

	if err := DB.Table("tbl_pages_group").Where("tbl_pages_group.is_deleted = ? and tbl_pages_group.spaces_id = ?", 0, id).Find(&getpagegrp).Error; err != nil {

		return &[]TblPagesGroup{}, err
	}

	return getpagegrp, nil
}

// delete page group
func (SpaceModel) SpaceDeletePageGroup(tblpage *TblPagesGroup, id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_pages_group").Where("tbl_pages_group.id=?", id).UpdateColumns(map[string]interface{}{"is_deleted": tblpage.IsDeleted, "deleted_on": tblpage.DeletedOn, "deleted_by": tblpage.DeletedBy}).Error; err != nil {

		return err
	}

	return nil
}

// delete page group aliases
func (SpaceModel) SpaceDeletePageGroupAliases(tblpageali *TblPagesGroupAliases, id int, DB *gorm.DB) error {

	if err := DB.Table("tbl_pages_group_aliases").Where("tbl_pages_group_aliases.page_group_id=?", id).UpdateColumns(map[string]interface{}{"is_deleted": tblpageali.IsDeleted, "deleted_on": tblpageali.DeletedOn, "deleted_by": tblpageali.DeletedBy}).Error; err != nil {

		return err
	}

	return nil
}

func (SpaceModel) ClonePages(pages *Tblpagealiases, DB *gorm.DB) error {

	if err := DB.Table("tbl_page_aliases").Omit("id").Create(&pages).Error; err != nil {

		return err

	}
	return nil

}
func (SpaceModel) ClonePagesGroup(pagegroup *TblPagesGroupAliases, DB *gorm.DB) error {

	if err := DB.Table("tbl_pages_group_aliases").Omit("id").Create(&pagegroup).Error; err != nil {

		return err

	}
	return nil
}
