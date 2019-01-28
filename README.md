# model_auto_example

### 1、exec ./db/create_db.sql in db

### 2、Switch directory to github.com/koalaone/model_auto
	
### 3、exec command 
```bash
./model_auto -import=github.com/model_auto_example/model -dbName=db_name -user=user_name -password=user_password -host=localhost -port=3306
```

### 4、copy "model" dir to "model_auto_example" dir

### 5、exec command
```bash
make vgo

make vendor

```

### 6、edit main.go
```go

package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koalaone/model_auto/mysql"
	"github.com/koalaone/model_auto_example/model"
)

func main() {
	dbConn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True"
	dbType := "mysql"
	err := mysql.InitConnDB(dbType, dbConn)
	if err != nil {
		log.Println("InitConnDB error:" + err.Error())
		return
	}

	wheres := make(map[string]interface{})
	wheres["uid"] = "u_1234"
	locUser, err := model.GetUserByWheres(wheres)
	if err != nil {
		log.Println("GetUserByWheres error:" + err.Error())
		return
	}

	if locUser == nil {
		newUser := model.User{
			Uid:   "u_1234",
			Uname: "koalaone",
			Age:   22,
		}
		err = model.CreateUser(&newUser)
		if err != nil {
			log.Println("CreateUser error:" + err.Error())
			return
		}
	} else {
		updates := make(map[string]interface{})
		updates["uname"] = "koala"
		err = model.UpdateUser(wheres, updates)
		if err != nil {
			log.Println("UpdateUser error:" + err.Error())
			return
		}
	}

	wheres = make(map[string]interface{})
	wheres["uid"] = "u_1234"
	locUser, err = model.GetUserByWheres(wheres)
	if err != nil {
		log.Println("GetUserByWheres error:" + err.Error())
		return
	}

	log.Printf("user info:%#v", locUser)
}


```