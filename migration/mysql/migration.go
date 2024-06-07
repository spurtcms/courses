package mysql

import (
	"gorm.io/gorm"
	"time"

	"gorm.io/datatypes"
)

type TblSpaces struct {
	Id             int       `gorm:"primaryKey;auto_increment"`
	PageCategoryId int       `gorm:"type:int"`
	CreatedOn      time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	CreatedBy      int       `gorm:"type:int"`
	ModifiedOn     time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy     int       `gorm:"DEFAULT:NULL"`
	DeletedOn      time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy      int       `gorm:"DEFAULT:NULL"`
	IsDeleted      int       `gorm:"DEFAULT:0"`
}

type TblSpacesAliases struct {
	Id                int
	SpacesId          int       `gorm:"type:int"`
	LanguageId        int       `gorm:"type:int"`
	SpacesName        string    `gorm:"type:varchar(255)"`
	SpacesSlug        string    `gorm:"type:varchar(255)"`
	SpacesDescription string    `gorm:"type:varchar(255)"`
	ImagePath         string    `gorm:"type:varchar(255)"`
	CreatedOn         time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	CreatedBy         int       `gorm:"type:int"`
	ModifiedOn        time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy        int       `gorm:"type:int;DEFAULT:NULL"`
	DeletedOn         time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy         int       `gorm:"type:int;DEFAULT:NULL"`
	IsDeleted         int       `gorm:"type:int;DEFAULT:0"`
	ViewCount         int       `gorm:"type:int"`
	RecentTime        time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
}

type TblPagesCategoriesAliases struct {
	Id                  int       `gorm:"primaryKey;auto_increment"`
	PageCategoryId      int       `gorm:"type:int"`
	LanguageId          int       `gorm:"type:int"`
	CategoryName        string    `gorm:"type:varchar(255)"`
	CategorySlug        string    `gorm:"type:varchar(255)"`
	CategoryDescription string    `gorm:"type:varchar(255)"`
	CreatedOn           time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	CreatedBy           int       `gorm:"type:int"`
	ModifiedOn          time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy          int       `gorm:"type:int;DEFAULT:NULL"`
	DeletedOn           time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy           int       `gorm:"type:int;DEFAULT:NULL"`
	IsDeleted           int       `gorm:"type:int;DEFAULT:0"`
	ParentId            int       `gorm:"type:int"`
}

type TblPagesCategories struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	CreatedOn  time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	CreatedBy  int       `gorm:"type:int"`
	ModifiedOn time.Time `gorm:"DEFAULT:NULL"`
	ModifiedBy int       `gorm:"type:int;DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"DEFAULT:NULL"`
	DeletedBy  int       `gorm:"type:int;DEFAULT:NULL"`
	IsDeleted  int       `gorm:"type:int;DEFAULT:0"`
}

type TblPagesGroup struct {
	Id         int       `gorm:"primaryKey;auto_increment"`
	SpacesId   int       `gorm:"type:int"`
	CreatedOn  time.Time `gorm:"type:timestamp"`
	CreatedBy  int       `gorm:"type:int"`
	ModifiedOn time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	ModifiedBy int       `gorm:"DEFAULT:NULL"`
	DeletedOn  time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	DeletedBy  int       `gorm:"DEFAULT:NULL"`
	IsDeleted  int       `gorm:"DEFAULT:0"`
}
type TblPagesGroupAliases struct {
	Id               int       `gorm:"primaryKey;auto_increment"`
	PageGroupId      int       `gorm:"type:int"`
	LanguageId       int       `gorm:"type:int"`
	GroupName        string    `gorm:"type:varchar(255)"`
	GroupSlug        string    `gorm:"type:varchar(255)"`
	GroupDescription string    `gorm:"type:varchar(255)"`
	CreatedOn        time.Time `gorm:"type:timestamp"`
	CreatedBy        int       `gorm:"type:int"`
	ModifiedOn       time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	ModifiedBy       int       `gorm:"type:int;DEFAULT:NULL"`
	DeletedOn        time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	DeletedBy        int       `gorm:"type:int;DEFAULT:NULL"`
	IsDeleted        int       `gorm:"type:int;DEFAULT:0"`
	OrderIndex       int       `gorm:"type:int"`
}
type TblPage struct {
	Id          int       `gorm:"primaryKey;auto_increment"`
	SpacesId    int       `gorm:"type:int"`
	PageGroupId int       `gorm:"type:int"`
	ParentId    int       `gorm:"type:int"`
	CreatedOn   time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	CreatedBy   int       `gorm:"type:int"`
	ModifiedOn  time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"type:int;DEFAULT:NULL"`
	DeletedOn   time.Time `gorm:"type:timestamp;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:int;DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:int;DEFAULT:0"`
}

type TblPageAliases struct {
	Id               int               `gorm:"primaryKey;auto_increment"`
	PageId           int               `gorm:"type:int"`
	LanguageId       int               `gorm:"type:int"`
	PageTitle        string            `gorm:"type:varchar(255)"`
	PageSlug         string            `gorm:"type:varchar(255)"`
	PageDescription  string            `gorm:"type:varchar(255)"`
	PublishedOn      time.Time         `gorm:"type:timestamp;DEFAULT:NULL"`
	Author           string            `gorm:"type:varchar(255)"`
	Excerpt          string            `gorm:"type:varchar(255)"`
	FeaturedImages   string            `gorm:"type:varchar(255)"`
	Access           string            `gorm:"type:varchar(255)"`
	MetaDetails      datatypes.JSONMap `gorm:"type:jsonb"`
	Status           string            `gorm:"type:varchar(255)"`
	AllowComments    bool              `gorm:"type:bool"`
	CreatedOn        time.Time         `gorm:"type:timestamp;DEFAULT:NULL"`
	CreatedBy        int               `gorm:"type:int"`
	ModifiedOn       time.Time         `gorm:"type:timestamp;DEFAULT:NULL"`
	ModifiedBy       int               `gorm:"DEFAULT:NULL"`
	DeletedOn        time.Time         `gorm:"type:timestamp;DEFAULT:NULL"`
	DeletedBy        int               `gorm:"DEFAULT:NULL"`
	IsDeleted        int               `gorm:"DEFAULT:0"`
	OrderIndex       int               `gorm:"type:int"`
	PageSuborder     int               `gorm:"type:int"`
	LastRevisionDate time.Time         `gorm:"type:timestamp;DEFAULT:NULL"`
	LastRevisionNo   int               `gorm:"type:int"`
	ReadTime         int               `gorm:"type:int"`
}

type TblPageAliasesLog struct {
	Id              int               `gorm:"primaryKey;auto_increment"`
	PageId          int               `gorm:"type:int"`
	LanguageId      int               `gorm:"type:int"`
	PageTitle       string            `gorm:"type:varchar(255)"`
	PageSlug        string            `gorm:"type:varchar(255)"`
	PageDescription string            `gorm:"type:varchar(255)"`
	PublishedOn     time.Time         `gorm:"type:timestamp;DEFAULT:NULL"`
	Author          string            `gorm:"type:varchar(255)"`
	Excerpt         string            `gorm:"type:varchar(255)"`
	FeaturedImages  string            `gorm:"type:varchar(255)"`
	Access          string            `gorm:"type:varchar(255)"`
	MetaDetails     datatypes.JSONMap `gorm:"type:jsonb"`
	Status          string            `gorm:"type:varchar(255)"`
	AllowComments   bool              `gorm:"type:bool"`
	CreatedOn       time.Time         `gorm:"type:timestamp"`
	CreatedBy       int               `gorm:"type:int"`
	ModifiedOn      time.Time         `gorm:"type:timestamp;DEFAULT:NULL"`
	ModifiedBy      int               `gorm:"type:int;DEFAULT:NULL"`
	DeletedOn       time.Time         `gorm:"type:timestamp;DEFAULT:NULL"`
	DeletedBy       int               `gorm:"type:int;DEFAULT:NULL"`
	ReadTime        int               `gorm:"type:int"`
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
