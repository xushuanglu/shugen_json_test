package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"common.dh.cn/utils"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) EchoJsonOk(msg ...interface{}) {
	if msg == nil {
		msg = []interface{}{"ok"}
	}
	c.Data["json"] = utils.P{"msg": msg[0]}
	c.ServeJSON()
}

//前台用户登录
func (c *LoginController) GetJson() {

	//cmd := "[{'children': [{'children': [],'id': 8,'name': '类目11-11'}],'id': 3,'name': '类目11'},{'children': [{'children': [{'children': [],'id': 9,'name': '类目2-2-1'}],'id': 5,'name': '类目2-2'},{'children': [{'children': [],'id': 7,'name': '类目2-4'},{'children': [{'children': [],'id': 12,'name': '类目'}],'id': 10,'name': '测试'}],'id': 6,'name': '类目2-3-1'}],'id': 4,'name': '类目2'}]"

	cmd := "{'admins':['381cf09d','3966eb66','381cf0f2'],'children':[{'admins':[],'children':[{'admins':[],'children':[],'company':'381cf09b','id':'5ec8ae584ed79b003966eb20','metadata':{'cellPhone':'','country':'','createSubDep':1,'departmentNo':'001_1','description':'','email':'','leader':'','nationCode':'','region':''},'name':'org1_1','parent':'5ec8ae2a4ed79b003966eb17','roles':['5ec8ae584ed79b003966eb21']}],'company':'381cf09b','id':'5ec8ae2a4ed79b003966eb17','metadata':{'cellPhone':'','country':'','createSubDep':1,'departmentNo':'001','description':'','email':'','leader':'','nationCode':'','region':''},'name':'org1','roles':['5ec8ae2a4ed79b003966eb18','5ecb74bc268d1400381cf0f1']},{'admins':[],'children':[],'company':'381cf09b','id':'5ec8ae444ed79b003966eb1d','metadata':{'cellPhone':'','country':'','createSubDep':1,'departmentNo':'003','description':'','email':'','leader':'','nationCode':'','region':''},'name':'org3','roles':['5ec8ae454ed79b003966eb1e']}],'deviceModelIds':['5151','5145','5156','5147','5148','5187',null],'id':'381cf09b','metadata':{'address':'','country':'','createSubDep':1,'description':'','email':'lijuan.du@rootcloud.com','leader':'sany_india_new'},'name':'三一印度_new_测试'}"

	str := strings.Replace(string(cmd), "'", "\"", -1)
	str = strings.Replace(str, "\n", "", -1)

	dat := make(map[string]interface{})

	// var dat []map[string]interface{}
	if err := json.Unmarshal([]byte(str), &dat); err == nil {
		fmt.Println(dat)
	} else {
		fmt.Println(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(dat)

	c.Data["json"] = utils.P{"data": dat}
	c.ServeJSON()

}
