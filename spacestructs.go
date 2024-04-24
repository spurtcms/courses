package spaces

import (
	"time"

	"github.com/spurtcms/categories"
)

type SpaceListReq struct {
	Offset              int
	Limit               int
	Keyword             string //(filter)
	CategoryId          int    //(filter)
	LanguageEnable      bool   //(optional)
	SetLanguageId       string //(optional)
	PublishedSpace      bool   //(optional)-want publishedspace only then enable true
	MemberAccessControl bool   //(optional)-want restricted space content hide then enable true
	MemberId            int
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

type SpaceDetail struct {
	SpaceId      int
	SpaceSlug    string
	ImagePathUrl string
}

type SpaceModel struct{}

var Spacemodel SpaceModel
