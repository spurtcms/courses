package postgres

import (
	"gorm.io/gorm"
	"time"

	"gorm.io/datatypes"
)

type TblSpaces struct {
	Id             int       `gorm:"primaryKey;auto_increment;type:serial"`
	PageCategoryId int       `gorm:"type:integer"`
	CreatedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy      int       `gorm:"type:integer"`
	ModifiedOn     time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy     int       `gorm:"DEFAULT:NULL"`
	DeletedOn      time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy      int       `gorm:"DEFAULT:NULL"`
	IsDeleted      int       `gorm:"DEFAULT:0"`
}

type TblSpacesAliases struct {
	Id                int       `gorm:"primaryKey;auto_increment;type:serial"`
	SpacesId          int       `gorm:"type:integer"`
	LanguageId        int       `gorm:"type:integer"`
	SpacesName        string    `gorm:"type:character varying"`
	SpacesSlug        string    `gorm:"type:character varying"`
	SpacesDescription string    `gorm:"type:character varying"`
	ImagePath         string    `gorm:"type:character varying"`
	CreatedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy         int       `gorm:"type:integer"`
	ModifiedOn        time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn         time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy         int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted         int       `gorm:"type:integer;DEFAULT:0"`
	ViewCount         int       `gorm:"type:integer"`
	RecentTime        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
}

type TblPagesCategoriesAliases struct {
	Id                  int       `gorm:"primaryKey;auto_increment;type:serial"`
	PageCategoryId      int       `gorm:"type:integer"`
	LanguageId          int       `gorm:"type:integer"`
	CategoryName        string    `gorm:"type:character varying"`
	CategorySlug        string    `gorm:"type:character varying"`
	CategoryDescription string    `gorm:"type:character varying"`
	CreatedOn           time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy           int       `gorm:"type:integer"`
	ModifiedOn          time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy          int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn           time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy           int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted           int       `gorm:"type:integer;DEFAULT:0"`
	ParentId            int       `gorm:"type:integer"`
}

type TblPagesCategories struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy  int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:integer;DEFAULT:0"`
}

type TblPagesGroup struct {
	Id         int       `gorm:"primaryKey;auto_increment;type:serial"`
	SpacesId   int       `gorm:"type:integer"`
	CreatedOn  time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy  int       `gorm:"type:integer"`
	ModifiedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"DEFAULT:NULL"`
	IsDeleted  int       `gorm:"DEFAULT:0"`
}
type TblPagesGroupAliases struct {
	Id               int       `gorm:"primaryKey;auto_increment;type:serial"`
	PageGroupId      int       `gorm:"type:integer"`
	LanguageId       int       `gorm:"type:integer"`
	GroupName        string    `gorm:"type:character varying"`
	GroupSlug        string    `gorm:"type:character varying"`
	GroupDescription string    `gorm:"type:character varying"`
	CreatedOn        time.Time `gorm:"type:timestamp without time zone"`
	CreatedBy        int       `gorm:"type:integer"`
	ModifiedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn        time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:integer;DEFAULT:0"`
	OrderIndex       int       `gorm:"type:integer"`
}
type TblPage struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	SpacesId    int       `gorm:"type:integer"`
	PageGroupId int       `gorm:"type:integer"`
	ParentId    int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:integer"`
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:integer;DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:integer;DEFAULT:0"`
}

type TblPageAliases struct {
	Id               int               `gorm:"primaryKey;auto_increment;type:serial"`
	PageId           int               `gorm:"type:integer"`
	LanguageId       int               `gorm:"type:integer"`
	PageTitle        string            `gorm:"type:character varying"`
	PageSlug         string            `gorm:"type:character varying"`
	PageDescription  string            `gorm:"type:character varying"`
	PublishedOn      time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	Author           string            `gorm:"type:character varying"`
	Excerpt          string            `gorm:"type:character varying"`
	FeaturedImages   string            `gorm:"type:character varying"`
	Access           string            `gorm:"type:character varying"`
	MetaDetails      datatypes.JSONMap `gorm:"type:jsonb"`
	Status           string            `gorm:"type:character varying"`
	AllowComments    bool              `gorm:"type:boolean"`
	CreatedOn        time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy        int               `gorm:"type:integer"`
	ModifiedOn       time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy       int               `gorm:"DEFAULT:NULL"`
	DeletedOn        time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy        int               `gorm:"DEFAULT:NULL"`
	IsDeleted        int               `gorm:"DEFAULT:0"`
	OrderIndex       int               `gorm:"type:integer"`
	PageSuborder     int               `gorm:"type:integer"`
	LastRevisionDate time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	LastRevisionNo   int               `gorm:"type:integer"`
	ReadTime         int               `gorm:"type:integer"`
}

type TblPageAliasesLog struct {
	Id              int               `gorm:"primaryKey;auto_increment;type:serial"`
	PageId          int               `gorm:"type:integer"`
	LanguageId      int               `gorm:"type:integer"`
	PageTitle       string            `gorm:"type:character varying"`
	PageSlug        string            `gorm:"type:character varying"`
	PageDescription string            `gorm:"type:character varying"`
	PublishedOn     time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	Author          string            `gorm:"type:character varying"`
	Excerpt         string            `gorm:"type:character varying"`
	FeaturedImages  string            `gorm:"type:character varying"`
	Access          string            `gorm:"type:character varying"`
	MetaDetails     datatypes.JSONMap `gorm:"type:jsonb"`
	Status          string            `gorm:"type:character varying"`
	AllowComments   bool              `gorm:"type:boolean"`
	CreatedOn       time.Time         `gorm:"type:timestamp without time zone"`
	CreatedBy       int               `gorm:"type:integer"`
	ModifiedOn      time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy      int               `gorm:"type:integer;DEFAULT:NULL"`
	DeletedOn       time.Time         `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy       int               `gorm:"type:integer;DEFAULT:NULL"`
	ReadTime        int               `gorm:"type:integer"`
}

func MigrationTables(DB *gorm.DB) {

	err := DB.AutoMigrate(
		&TblSpaces{},
		&TblSpacesAliases{},
		&TblPagesCategories{},
		&TblPagesCategoriesAliases{},
		&TblPagesGroup{},
		&TblPagesGroupAliases{},
		&TblPage{},
		&TblPageAliases{},
		&TblPageAliasesLog{},
	)

	if err != nil {

		panic(err)

	}

}
