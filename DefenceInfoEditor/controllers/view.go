package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"strconv"
	"strings"
	"github.com/jmoiron/jsonq"
	"os"
	"fmt"
	"sort"
	"reflect"
	"path/filepath"
	"io/ioutil"
	//"github.com/mitchellh/mapstructure"
)

type ViewController struct {
	beego.Controller
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
	PP []PositionInfo  `json:"pp"`
	PI []PlantsInfo  `json:"pi"`
	VI string  `json:"vi"`
	MN string  `json:"mn"`
	EMN string  `json:"emn"`
}

type DefenceInfo struct {
	DefenceInfoDetail DefenceInfoDetail `json:"d"`
}

type DefenceOneLevelUpInfo struct {
	PP map[string]PositionNumInfo
	PI map[string]PlantsInfo
}

func printObject(info string, v interface{}) {
	fmt.Printf(info + " %+v\n", v)
}

type PositionPrintNumInfo struct {
	plantId string
	num int
	level string
}


func (c *ViewController) Get() {
	worldName := c.Ctx.Input.Query("worldName");
	beego.Debug("Get value is  " + worldName)
	worldJsonStr := readJsonFileAsString(worldName)
	resultMap := make(map[string]string);
	resultMap["defenceStr"] = worldJsonStr
	if(strings.Contains(worldName, "init")) {
		resultMap["detailInfo"] = GetWorldInitStr(worldName)	
	}else{
		resultMap["detailInfo"] = GetWorldLevelUpStr(worldName)
	}
	resultBytes, _ := json.Marshal(resultMap);
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

func GetWorldLevelUpStr(levelUpWordName string) string {
	var initWorldName string;
	if(strings.Contains(levelUpWordName, "egypt")){
		initWorldName = "egypt_initpos"
	}else if(strings.Contains(levelUpWordName, "pirate")) {
		initWorldName = "pirate_initpos"
	}else if(strings.Contains(levelUpWordName, "kongfu")) {
		initWorldName = "kongfu_initpos"
	}
	beego.Debug("initWorldName", initWorldName);
	printNumInfoMap := GetWorldInitData(initWorldName);
	//fmt.Printf("printNumInfoMap, " + " %+v\n", printNumInfoMap)
	
	defenceLevelUpInfo := GetWorldLevelupInfo(levelUpWordName);
	
	sumPlantInfoMap := make(map[int](map[string]PositionPrintNumInfo) );
	
	var levels []int
    for k := range defenceLevelUpInfo {
        levels = append(levels, k)
    }
    sort.Ints(levels)
	var firstLevel = true;
	for _, level := range levels {
		oneLevelInfo, _ := defenceLevelUpInfo[level];		
		//FirstLevel
		sumPlantInfoMap[level] = make(map[string]PositionPrintNumInfo, 0);
		if(firstLevel) {
			sumPlantInfoMap[1] = make(map[string]PositionPrintNumInfo, 0);
			updateNumInfoMap, _:= sumPlantInfoMap[1];
			for key, value := range printNumInfoMap {
				updateNumInfoMap[key] = value;
			}
		}	
		updateNumInfoMap, _ := sumPlantInfoMap[level] 
		for key, value := range sumPlantInfoMap[level - 1] {
			updateNumInfoMap[key] = value;
		}		
		for plantId, v := range oneLevelInfo.PI {
			if info, ok := updateNumInfoMap[plantId]; ok {
				info.level = v.L;
			} else {
				var info PositionPrintNumInfo;
				info.level = v.L;
				info.num = 0;
				info.plantId = v.T;
				updateNumInfoMap[plantId] = info;
			}
		}
		for plantId, v := range oneLevelInfo.PP {
			if info, ok := updateNumInfoMap[plantId]; ok {
				vNum, _ := strconv.Atoi(v.N);
				info.num = info.num + vNum;
				updateNumInfoMap[plantId] = info;
			}else{
				fmt.Printf("sumPlantInfoMap Error, no pi info, %d \n ", plantId)
			}
		}
		//fmt.Printf("sumPlantInfoMap2 %+v\n", updateNumInfoMap)
	}
	//fmt.Printf("sumPlantInfoMap %+v\n", sumPlantInfoMap)
	return GetStrInfoFromLevelMap(sumPlantInfoMap);
}

func GetWorldInitData(worldName string) map[string]PositionPrintNumInfo {
	worldInitInfo := GetWorldInitInfo(worldName)
	printNumInfoMap := make(map[string]PositionPrintNumInfo, 0);
	for _,onePositionInfo := range worldInitInfo.DefenceInfoDetail.PP {
		for _,onePlantInfo := range worldInitInfo.DefenceInfoDetail.PI {
			//printObject("pp  is ", onePositionInfo);
			//printObject("pi  is ", onePlantInfo);
			
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
	return printNumInfoMap;
}

func GetWorldInitStr(worldName string) string {
	printNumInfoMap := GetWorldInitData(worldName);
	printObject("levelUpMap info  is ", printNumInfoMap);
	str := GetStrInfoFromInitMap(printNumInfoMap);
	printObject("GetStrInfoFromInfoMap info  is ", str);
	return str;
}

func GetStrInfoFromInitMap(mapInfo map[string]PositionPrintNumInfo) string {
	var str string;
	for _, plantInfo := range mapInfo {
		str += ("" + plantInfo.plantId + "      x" + strconv.Itoa(plantInfo.num) + "        lv." +  plantInfo.level + "\n" );
	}
	return str;
}

func GetStrInfoFromLevelMap(mapInfo map[int](map[string]PositionPrintNumInfo) ) string {
	var levels []int
    for k, _ := range mapInfo {
        levels = append(levels, k)
    }
    sort.Ints(levels)
    //fmt.Printf("GetStrInfoFromLevelMap, " + " %+v\n", levels)
    
	var str string;
	for _, level := range levels {
		info, _ := mapInfo[level];
		
		//fmt.Printf("GetStrInfoFromLevelMap, info" + " %+v\n", info)
		
		str += "等级" + strconv.Itoa(level);
		str += "\n";
		
		var plantIds []string
		for plantId, _ := range info {
			plantIds = append(plantIds,  plantId);
		}
		sort.Strings(plantIds);
		 
		for _, plantId := range plantIds {
			//fmt.Println("plantInfo,  id:",  strconv.Itoa(plantId)  )
			plantInfo, _:= info[ plantId ];
			//fmt.Printf("plantInfo, " + " %+v\n", plantInfo )
			str += ("" + plantInfo.plantId + "      x" + strconv.Itoa(plantInfo.num) + "        lv." +  plantInfo.level + "\n" );
		}
		str += "\n";
		str += "\n";
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
	
	for level := 2; level <= 30; level++ {
		var oneDefenceInfo DefenceOneLevelUpInfo;
		oneDefenceInfo.PP = make( map[string]PositionNumInfo, 0);
		oneDefenceInfo.PI = make( map[string]PlantsInfo, 0);
		
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
				oneDefenceInfo.PP[ positionNumInfo.T ]  = positionNumInfo;
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
				oneDefenceInfo.PI[ plantsInfo.T ] = plantsInfo
			}
		}
		levelUpMap[ level ] = oneDefenceInfo;
	}
	//printObject("levelUpMap info  is ", levelUpMap);

	return levelUpMap;
}


func GetWorldInitInfo(worldName string) DefenceInfo {
	defenceStr := readJsonFileAsString(worldName);
	var defenceInfo DefenceInfo	
	err := json.Unmarshal([]byte(defenceStr), &defenceInfo)
	if err != nil {
	 	beego.Debug("getWorldInitString " + worldName + ", value:" + defenceStr)
 	}
	plantSize := strconv.Itoa(len(defenceInfo.DefenceInfoDetail.PI));
	beego.Debug("getWorldInitString pi num is," + plantSize);
	plantInfoSize := strconv.Itoa(len(defenceInfo.DefenceInfoDetail.PP))
	beego.Debug("getWorldInitString pp num is," + plantInfoSize);
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
	
	resultMap := make(map[string]string);
	resultMap["defenceStr"] = configValue
	if(strings.Contains(worldName, "init")) {
		resultMap["detailInfo"] = GetWorldInitStr(worldName)	
	}else{
		resultMap["detailInfo"] = GetWorldLevelUpStr(worldName)
	}
	resultBytes, _ := json.Marshal(resultMap);
	c.Ctx.WriteString( string(resultBytes) );
	
	}
