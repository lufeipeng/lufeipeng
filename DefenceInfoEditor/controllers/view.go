package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"strconv"
	"strings"
	"github.com/jmoiron/jsonq"
	"os"
	"fmt"
	"reflect"
	"path/filepath"
	"io/ioutil"
	//"github.com/mitchellh/mapstructure"
)

type ViewController struct {
	beego.Controller
}

func (c *ViewController) Get() {
	worldName := c.Ctx.Input.Query("worldName");
	beego.Debug("Get value is  " + worldName)
	worldJsonStr := readJsonFileAsString(worldName)
	resultMap := make(map[string]string);
	resultMap["defenceStr"] = worldJsonStr
	resultMap["detailInfo"] = "current detail info is"
	
	resultBytes, _ := json.Marshal(resultMap);

	//mapName := GetWorldInitString(worldName);
	//beego.Debug("worldMapName is" + mapName)
	
	c.Ctx.WriteString( string(resultBytes) );
}


func readJsonFileAsString(worldName string) string {
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
	return string(buff);
}

type PositionInfo struct {
	T string  `json:"T"`
	R string  `json:"R"`
	C string  `json:"C"`
}

type PositionNumInfo struct {
	T string  `json:"T"`
	N string  `json:"N"`
}

type PlantsInfo struct {
	T string  `json:"T"`
	L string  `json:"L"`
	A string  `json:"A"`
	Y string  `json:"Y"`
}

type DefenceInfoDetail struct {
	pp []PositionInfo  `json:"pp"`
	pi []PlantsInfo  `json:"pi"`
	vi string  `json:"vi"`
	mn string  `json:"mn"`
	emn string  `json:"emn"`
}

type DefenceInfo struct {
	DefenceInfoDetail DefenceInfoDetail `json:"d"`
}

type DefenceOneLevelUpInfo struct {
	pp []PositionNumInfo  `json:"pp"`
	pi []PlantsInfo  `json:"pi"`
}

func printObject(info string, v interface{}) {
	fmt.Printf(info + " %+v\n", v)
}

type PositionPrintNumInfo struct {
	plantId string
	num int
	level string
}

func GetWorldInitStr(worldName string) string{
	worldInitInfo := GetWorldInitInfo(worldName)
	printNumInfoMap := make(map[string]PositionPrintNumInfo, 0);
	for _,onePositionInfo := range worldInitInfo.DefenceInfoDetail.pp {
		for _,onePlantInfo := range worldInitInfo.DefenceInfoDetail.pi {
			
			printObject("pp  is ", onePositionInfo);
			printObject("pi  is ", onePlantInfo);
			
			if(onePlantInfo.T == onePositionInfo.T) {
				if v, ok := printNumInfoMap[onePlantInfo.T]; ok {
					v.num++;
					printNumInfoMap[onePlantInfo.T] = v;
				} else {
					var positionPrintNumInfo PositionPrintNumInfo;
					positionPrintNumInfo.plantId = onePlantInfo.T;
					positionPrintNumInfo.num = 1;
					positionPrintNumInfo.level = onePlantInfo.L
					printNumInfoMap[onePlantInfo.T] = positionPrintNumInfo
				}
			}
		}
	}
	printObject("levelUpMap info  is ", printNumInfoMap);
	str := GetStrInfoFromInfoMap(printNumInfoMap);
	printObject("GetStrInfoFromInfoMap info  is ", str);
	return str;
}

func GetStrInfoFromInfoMap(mapInfo map[string]PositionPrintNumInfo) string{
	var str string;
	for _, info := range mapInfo {
		str += ("plantID is" + info.plantId + ",level:" + info.level + ",num" + strconv.Itoa(info.num) );
	}
	return str;
}


func GetWorldLevelupInfo(worldName string) map[int]DefenceOneLevelUpInfo {
	defenceStr := readJsonFileAsString(worldName);
			
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(defenceStr))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	
	levelUpMap := make(map[int]DefenceOneLevelUpInfo);
	
	for level := 1; level <= 30; level++ {
		var oneDefenceInfo DefenceOneLevelUpInfo;
		oneDefenceInfo.pp = make([]PositionNumInfo,0);
		oneDefenceInfo.pi = make([]PlantsInfo, 0);
		
		objectInfo, _ := jq.Object("l" + strconv.Itoa(level) );
		v := reflect.ValueOf(objectInfo);
		i := v.Interface()
		v1 := i.(map[string]interface{})

		ppInfoTypeArray := reflect.ValueOf(v1["pp"])
		piInfoTypeArray := reflect.ValueOf(v1["pi"])
		if(!ppInfoTypeArray.IsValid() && !piInfoTypeArray.IsValid() ) {
			continue;
		}
			
		if(ppInfoTypeArray.IsValid()){
			ppInfoArray := ppInfoTypeArray.Interface().([]interface{});
			for _, ppInfo := range ppInfoArray {
				var positionNumInfo PositionNumInfo
				positionNumInfo.T = reflect.ValueOf( reflect.ValueOf(ppInfo).Interface().(map[string]interface{})["T"] ).String();
				positionNumInfo.N = reflect.ValueOf( reflect.ValueOf(ppInfo).Interface().(map[string]interface{})["N"] ).String();
				oneDefenceInfo.pp = append(oneDefenceInfo.pp, positionNumInfo);
			}
		}
		
		if piInfoTypeArray.IsValid() {
			piInfoArray := piInfoTypeArray.Interface().([]interface{});
			for _, piInfo := range piInfoArray {
				var plantsInfo PlantsInfo
				plantsInfo.T = reflect.ValueOf( reflect.ValueOf(piInfo).Interface().(map[string]interface{})["T"] ).String();
				plantsInfo.L = reflect.ValueOf( reflect.ValueOf(piInfo).Interface().(map[string]interface{})["L"] ).String();
				plantsInfo.A = reflect.ValueOf( reflect.ValueOf(piInfo).Interface().(map[string]interface{})["A"] ).String();
				plantsInfo.Y = reflect.ValueOf( reflect.ValueOf(piInfo).Interface().(map[string]interface{})["Y"] ).String();
				oneDefenceInfo.pi = append(oneDefenceInfo.pi, plantsInfo);
			}
		}
		levelUpMap[ level ] = oneDefenceInfo;
	}
	printObject("levelUpMap info  is ", levelUpMap);

	return levelUpMap;
}


func GetWorldInitInfo(worldName string) DefenceInfo {
	defenceStr := readJsonFileAsString(worldName);
	var defenceInfo DefenceInfo;
	
	beego.Debug("GetWorldInitInfo " + worldName )
	
	err := json.Unmarshal([]byte(defenceStr), &defenceInfo)
	if err != nil {
	 	beego.Debug("getWorldInitString " + worldName + ", value:" + defenceStr)
 	}
	plantSize := strconv.Itoa(len(defenceInfo.DefenceInfoDetail.pi));
	beego.Debug("getWorldInitString pi num is" + plantSize);
	
	for _, plantInfo := range defenceInfo.DefenceInfoDetail.pi {
		beego.Debug("getWorldInitString " + plantInfo.T);
	}
	
	for _, positionInfo := range defenceInfo.DefenceInfoDetail.pp {
		beego.Debug("getWorldInitString " + positionInfo.T);
	} 
	printObject("defenceInfo init info  is ", defenceInfo);
	return defenceInfo;
}


func (c *ViewController) Post() {
	worldName := c.Ctx.Input.Query("worldName");
	configValue := c.Ctx.Input.Query("value");
	
	beego.Debug("Post value is  " + worldName + "value:" + configValue)
	
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
