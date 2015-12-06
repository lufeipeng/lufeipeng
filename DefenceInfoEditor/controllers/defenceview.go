package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	//"strings"
	//"github.com/jmoiron/jsonq"
)

type ViewDefenceController struct {
	beego.Controller
}

func (c *ViewDefenceController) Post() {
	textValue := c.Ctx.Input.Query("value");
	platform := c.Ctx.Input.Query("platform");
	
	resultMap := make(map[string]string);
	resultMap["defenceStr"] = textValue
	resultMap["detailInfo"] = TransToPlatform(textValue, platform);
	resultBytes, _ := json.Marshal(resultMap);
	c.Ctx.WriteString( string(resultBytes) );
	}
