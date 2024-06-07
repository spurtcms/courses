package spaces

import (
	"gorm.io/datatypes"
	"time"
)

type TblSpaces struct {
	Id             int
	PageCategoryId int
	CreatedOn      time.Time
	CreatedBy      int
	ModifiedOn     time.Time
	ModifiedBy     int
	DeletedOn      time.Time
	DeletedBy      int
	IsDeleted      int
}

type TblSpacesAliases struct {
	Id                int
	SpacesId          int
	LanguageId        int
	SpacesName        string
	SpacesSlug        string
	SpacesDescription string
	ImagePath         string
	CreatedOn         time.Time
	CreatedBy         int
	ModifiedOn        time.Time
	ModifiedBy        int
	DeletedOn         time.Time
	DeletedBy         int
	IsDeleted         int
	ViewCount         int
	RecentTime        time.Time
}

type TblPageAliases struct {
	Id               int
	PageId           int
	LanguageId       int
	PageTitle        string
	PageSlug         string
	PageDescription  string
	PublishedOn      time.Time
	Author           string
	Excerpt          string
	FeaturedImages   string
	Access           string
	MetaDetails      datatypes.JSONMap
	Status           string
	AllowComments    bool
	CreatedOn        time.Time
	CreatedBy        int
	ModifiedOn       time.Time
	ModifiedBy       int
	DeletedOn        time.Time
	DeletedBy        int
	IsDeleted        int
	OrderIndex       int
	PageSuborder     int
	LastRevisionDate time.Time
	LastRevisionNo   int
	ReadTime         int
}
