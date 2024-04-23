package spaces

import (
	"time"

	"github.com/spurtcms/categories"
	"gorm.io/gorm"
)

type SpaceListReq struct {
	Offset         int
	Limit          int
	Keyword        string
	CategoryId     int
	LanguageEnable bool
	SetLanguageId  string
}

type tblspaces struct {
	Id             int
	PageCategoryId int
	CreatedOn      time.Time
	CreatedBy      int
	ModifiedOn     time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy     int       `gorm:"DEFAULT:NULL"`
	DeletedOn      time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy      int       `gorm:"DEFAULT:NULL"`
	IsDeleted      int       `gorm:"DEFAULT:0"`
	Username       string    `gorm:"-:migration;<-:false"`
	CreatedDate    string    `gorm:"-"`
	ModifiedDate   string    `gorm:"-"`
	SpaceName      string    `gorm:"-"`
}

type Tblspacesaliases struct {
	Id                   int
	SpacesId             int
	LanguageId           int
	SpacesName           string
	SpacesSlug           string
	SpacesDescription    string
	ImagePath            string
	CreatedOn            time.Time
	CreatedBy            int                        `gorm:"type:integer"`
	ModifiedOn           time.Time                  `gorm:"DEFAULT:NULL"`
	ModifiedBy           int                        `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn            time.Time                  `gorm:"DEFAULT:NULL"`
	DeletedBy            int                        `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted            int                        `gorm:"type:integer;DEFAULT:0"`
	PageCategoryId       int                        `gorm:"-:migration;column:page_category_id;<-:false"`
	ParentId             int                        `gorm:"-:migration;column:parent_id;<-:false"`
	CreatedDate          string                     `gorm:"-"`
	ModifiedDate         string                     `gorm:"-"`
	CategoryNames        []categories.TblCategories `gorm:"-"`
	CategoryId           int                        `gorm:"-:migration;column:category_id;<-:false"`
	FullSpaceAccess      bool                       `gorm:"-"`
	SpaceFullDescription string                     `gorm:"-"`
	ReadTime             string                     `gorm:"-"`
	ViewCount            int                        `gorm:"type:integer"`
	RecentTime           time.Time
}

type tblpagescategoriesaliases struct {
	Id                  int    `gorm:"primaryKey;auto_increment"`
	PageCategoryId      int    `gorm:"type:integer"`
	LanguageId          int    `gorm:"type:integer"`
	CategoryName        string `gorm:"type:character varying"`
	CategorySlug        string `gorm:"type:character varying"`
	CategoryDescription string `gorm:"type:character varying"`
	CreatedOn           time.Time
	CreatedBy           int       `gorm:"type:integer"`
	ModifiedOn          time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy          int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn           time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy           int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted           int       `gorm:"type:integer;DEFAULT:0"`
	ParentId            int       `gorm:"type:integer"`
}

type tblpagescategories struct {
	Id         int `gorm:"primaryKey;auto_increment"`
	CreatedOn  time.Time
	CreatedBy  int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:integer;DEFAULT:0"`
}

type SpaceCreation struct {
	Name        string
	Description string
	ImagePath   string
	CategoryId  int //child category id
	LanguageId  int //For specific language space
	CreatedBy   int
	ModifiedBy  int
}

type SpaceModel struct{}

var Spacemodel SpaceModel

/*spaceList*/
func (SpaceModel) SpaceList(spacereq SpaceListReq, spaceid []int, DB *gorm.DB) (tblspace []Tblspacesaliases, spacecount int64, err error) {

	query := DB.Model(TblSpacesAliases{}).Select("tbl_spaces_aliases.*,tbl_spaces.page_category_id,tbl_categories.parent_id").
		Joins("inner join tbl_spaces on tbl_spaces_aliases.spaces_id = tbl_spaces.id").
		Joins("inner join tbl_languages on tbl_languages.id = tbl_spaces_aliases.language_id").
		Joins("inner join tbl_categories on tbl_categories.id = tbl_spaces.page_category_id").
		Where("tbl_spaces.is_deleted = 0 and tbl_spaces_aliases.is_deleted = 0 and tbl_spaces_aliases.language_id = 1").Order("tbl_spaces.id desc")

	if len(spaceid) != 0 {

		query = query.Where("tbl_spaces.id in (?)", spaceid)
	}

	if spacereq.Keyword != "" {

		query = query.Where("LOWER(TRIM(tbl_spaces_aliases.spaces_name)) ILIKE LOWER(TRIM(?))", "%"+spacereq.Keyword+"%")
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

	if err := DB.Model(TblPageAliases{}).Select("max(tbl_page_aliases.modified_on) as modified_on").Joins("inner join tbl_pages on tbl_pages.Id = tbl_page_aliases.page_id").Where("tbl_pages.spaces_Id=?", spaceid).Group("tbl_page_aliases.id").First(tblpageali).Error; err != nil {
		return err
	}

	return nil
}

func (SpaceModel) GetSpacealiaseDetails(spaceid int, spaceslug string, DB *gorm.DB) (TblSpacesAliases Tblspacesaliases, err error) {

	query := DB.Table("tbl_spaces_aliases").First(&TblSpacesAliases)

	if spaceid > 0 {

		query = query.Where("spaces_id=?", spaceid)

	}

	if spaceslug != "" {

		query = query.Where("spaces_slug=?", spaceid)
	}

	if err := query.Error; err != nil {

		return Tblspacesaliases{}, err
	}

	return TblSpacesAliases, nil
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

	if err := DB.Model(TblSpacesAliases{}).Where("spaces_id = ?", id).UpdateColumns(map[string]interface{}{"deleted_by": tblspace.DeletedBy, "deleted_on": tblspace.DeletedOn, "is_deleted": tblspace.IsDeleted}).Error; err != nil {

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
