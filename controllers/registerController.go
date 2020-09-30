package controllers

import (
	"BeegoHello0929/db_mysql"
	"BeegoHello0929/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController)Post() {
	dataBytes, err := ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	var user models.User
	err = json.Unmarshal(dataBytes, &user)
	if err != nil {
		r.Ctx.WriteString("数据解析错误，请重试")
		//result:=models.Result{
		//Code:0,
		//Message:"数据解析错误，请重试",
		//Data:nil,
	//}
		//r.Data["json"]=&result
		//r.ServeJSON()
		return
	}
	row, err := db_mysql.AddUser(user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("注册用户信息失败，请重试")
		//result:=models.Result{
			//Code:0,
			//Message:"注册用户信息失败，请重试",
			//Data:nil,
		//}
		//r.Data["json"]=&result
		//r.ServeJSON()
		return
	}
	fmt.Println(row)
	//md5Hash:=md5.New()
	//md5Hash.Write([]byte(user.Password))
	//user.Password=hex.EncodeToString(md5Hash.Sum(nil))
	//result:=models.Result{
	//code:1,
	//Message:"恭喜，注册用户信息成功",
	//Data:user,
//}
	//r.Data["json"]=&result()
	//r.ServeJSON()
	r.Ctx.WriteString("恭喜，注册用户信息成功")
}
