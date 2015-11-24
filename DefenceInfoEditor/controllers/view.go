package controllers

import (
	"github.com/astaxie/beego"
	//"time"
	"os"
	"path/filepath"
	"io/ioutil"
)

type ViewController struct {
	beego.Controller
}

func (c *ViewController) Get() {
	worldName := c.Ctx.Input.Query("worldName");
	beego.Debug("Post value is  " + worldName)
	
	var err error
	currentPath, err := os.Getwd()
	if err != nil {
		beego.Debug("Get Current Path failed")
	}
	confPath := filepath.Join(currentPath, "jsonfiles", worldName + ".json")
	
	beego.Debug("Current Path is " + confPath)
	
	beego.Debug("channel.conf path is " + confPath)

	buff, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic("open file failed!")
	}
	c.Ctx.WriteString( string(buff) );
}

func (c *ViewController) Post() {
	worldName := c.Ctx.Input.Query("worldName");
	configValue := c.Ctx.Input.Query("values");
	
	beego.Debug("Post value is  " + worldName)
	
	var err error
	currentPath, err := os.Getwd()
	if err != nil {
		beego.Debug("Get Current Path failed")
	}
	confPath := filepath.Join(currentPath, "jsonfiles", worldName + ".json")
	
	beego.Debug("Current Path is " + confPath)
	
	beego.Debug("channel.conf path is " + confPath)

	err2 := ioutil.WriteFile(confPath, ([] byte)(configValue), os.ModeAppend);
	if err2 != nil {
		panic("open file failed!")
	}
	c.Ctx.WriteString( configValue );
	}
