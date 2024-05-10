package spaces

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var SecretKey = "Secret123"

// Db connection
func DBSetup() (*gorm.DB, error) {

	dbConfig := map[string]string{
		"username": "postgres",
		"password": "****",
		"host":     "localhost",
		"port":     "5432",
		"dbname":   "spurtcms",
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "user=" + dbConfig["username"] + " password=" + dbConfig["password"] +
			" dbname=" + dbConfig["dbname"] + " host=" + dbConfig["host"] +
			" port=" + dbConfig["port"] + " sslmode=disable TimeZone=Asia/Kolkata",
	}), &gorm.Config{})

	if err != nil {

		log.Fatal("Failed to connect to database:", err)

	}
	if err != nil {

		return nil, err

	}

	return db, nil
}

// test spacelist function
func TestSpaceList(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, err:= Auth.IsGranted("Spaces", auth.CRUD)

	log.Println("permisision",permisison,err)

	space := SpaceSetup(&Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	if permisison {

		spacelist, count, err := space.SpaceList(SpaceListReq{Limit: 10, Offset: 0})

		if err != nil {

			panic(err)
		}

		fmt.Println(spacelist, count)

	} else {

		log.Println("permissions enabled not initialised")

	}

}

// test getspacedetail function
func TestGetSpaceDetail(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(&Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	spacedata, err := space.SpaceDetail(SpaceDetail{SpaceId: 1})

	if err != nil {

		panic(err)
	}

	fmt.Println(spacedata)

}

// test spacecreation function
func TestSpaceCreation(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(&Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	spacedata,err := space.SpaceCreation(SpaceCreation{Name: "Default Space",Description: "default space",CategoryId: 1,CreatedBy: 1})

	if err != nil {

		panic(err)
	}

	fmt.Println(spacedata)
}

//test spaceupdate function 
func TestSpaceUpdate(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(&Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	err := space.SpaceUpdate(SpaceCreation{Name: "Default Space",Description: "default space",CategoryId: 1,CreatedBy: 1},1)

	if err != nil {

		panic(err)
	}

}

func TestDeleteSpace(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(&Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	err := space.DeleteSpace(1,1)

	if err != nil {

		panic(err)
	}

}

// test clonespace function
func TestCloneSpace(t *testing.T) {

	db, _ := DBSetup()

	space := SpaceSetup(&Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
	})

	spacedata,err := space.CloneSpace(1,1)

	if err != nil {

		panic(err)
	}

	fmt.Println(spacedata)
}