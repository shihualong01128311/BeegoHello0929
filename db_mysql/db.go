package db_mysql

import (
	"BeegoHello0929/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)
var Db *sql.DB

func Connect()  {

	config:=beego.AppConfig
	dbDriver:=config.String("db_driver")
	dbUser:=config.String("db_user")
	dbPassword:=config.String("db_password")
	dbIp:=config.String("db_ip")
	dbName:=config.String("db_name")
	fmt.Println(dbDriver,dbUser,dbPassword)
	connUrl:=dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf-8"
	db,err:=sql.Open(dbDriver,connUrl)
	if err!=nil{
		fmt.Println(err.Error())
		panic("数据库连接错误，请检查配置")
	}
	Db=db
	fmt.Println(db)
	
}
func AddUser(u models.User) (int64,error) {
	md5Hash:=md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes:=md5Hash.Sum(nil)
	u.Password=hex.EncodeToString(passwordBytes)
	result,err:=Db.Exec("inster into user(name,birthday,address,password)"+
		"values(?,?,?,?)",u.Name,u.Birthday,u.Address,u.Password)
	if err!=nil{
		return -1, err
	}
	row,err:=result.RowsAffected()
	if err!=nil {
		return -1, err
	}
	return row,nil
}

