package spaces

import "time"

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


func (spaces *Spaces) DeletedPageGroupAliases(deletedBy int) error {

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
