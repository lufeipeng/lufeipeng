package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	_ "github.com/lufeipeng/lufeipeng/DefenceInfoEditor/routers"
	controllers "github.com/lufeipeng/lufeipeng/DefenceInfoEditor/controllers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}


// TestMain is a sample to run an endpoint test
func TestMain(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestMain", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	        Convey("The Result Should Not Be Empty", func() {
	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
	        })
	})
}

func TestReflect(t *testing.T) {

	/*
	for _, piInfo := range piInfoArray {
		printObject("T is ", reflect.ValueOf(piInfo).FieldByName("T").String() );
		printObject("N is ", reflect.ValueOf(piInfo).FieldByName("L").String() );
	}*/
	
	/*
	var info1 PositionNumInfo
	printObject("pp value is ", reflect.ArrayOf(reflect.ValueOf(ppInfo).Len(), reflect.TypeOf(info1) ) );
	
	var info2 PlantsInfo
	printObject("pi value is ", reflect.ArrayOf(reflect.ValueOf(piInfo).Len(), reflect.TypeOf(info2) ) );
	*/
	/*
	ppInfoDetail := ppInfo.([]PositionNumInfo);
	reflect.ValueOf(piInfo).([]PlantsInfo);
	
	printObject("pp info is", reflect.ValueOf(ppInfo).Len());	
	printObject("pi info is", reflect.ValueOf(piInfo).Len());	
	*/
	
	//beego.Debug("type:", reflect.TypeOf(objectInfo) );
		
	/*
	var oneDefenceOnfo DefenceOneLevelUpInfo;
	err := mapstructure.Decode(data, &oneDefenceOnfo);	
	if err != nil {
	 	beego.Debug("getWorldInitString " + worldName + ", error:" + err.Error())
 	}
	for _, plantInfo := range oneDefenceOnfo.pi {
		beego.Debug("getWorldInitString " + plantInfo.T);
	}
	
	for _, positionInfo := range oneDefenceOnfo.pp {
		beego.Debug("getWorldInitString " + positionInfo.T);
	}
	*/
	
	/*
	fmt.Println("type:", reflect.TypeOf(l1Objects) );
	v := reflect.ValueOf(l1Objects);
	v.MapKeys();
	defenceNumInfo := v.MapIndex("pp");
	for i:= 0; i < defenceNumInfo.Len(); i++ {
		oneDefenceNumInfo := defenceNumInfo.Index(i);
		beego.Debug("getWorldInitString " + oneDefenceNumInfo.FieldByName("T").String() + ", value:" + oneDefenceNumInfo.FieldByName("n").String() )
	}*/
	/*
	for level, info := range defenceLevelUpInfo.LevelUpInfos {
		beego.Debug("GetWorldLevelupString " + level);
		
		for _, plantInfo := range info.pi {
			beego.Debug("GetWorldLevelupString " + plantInfo.T);
		}
		
		for _, positionNumInfo := range info.pp {
			beego.Debug("GetWorldLevelupString " + positionNumInfo.T);
		}
	}*/
	}


func TestView(t *testing.T) {
	mapName := controllers.GetWorldLevelupString("egypt_levelup");
	beego.Trace("testing", "TestMain", mapName)
}