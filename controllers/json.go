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

	cmd := "[{'children': [{'children': [],'id': 8,'name': '类目11-11'}],'id': 3,'name': '类目11'},{'children': [{'children': [{'children': [],'id': 9,'name': '类目2-2-1'}],'id': 5,'name': '类目2-2'},{'children': [{'children': [],'id': 7,'name': '类目2-4'},{'children': [{'children': [],'id': 12,'name': '类目'}],'id': 10,'name': '测试'}],'id': 6,'name': '类目2-3-1'}],'id': 4,'name': '类目2'}]"

	str := strings.Replace(string(cmd), "'", "\"", -1)
	str = strings.Replace(str, "\n", "", -1)

	var dat []map[string]interface{}
	if err := json.Unmarshal([]byte(str), &dat); err == nil {
		// fmt.Println(dat)
	} else {
		fmt.Println(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.Encode(dat)

	c.Data["json"] = dat
	c.ServeJSON()

}
