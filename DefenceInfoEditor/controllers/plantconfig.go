package controllers

import(
	"strings"
	"strconv"
)

var AndroidPlantMap map[string]int
var IosPlantMap map[string]int


func init() {
	AndroidPlantMap = make(map[string]int);
	IosPlantMap = make(map[string]int);
	initAndroid();
	initIos();
}

func TransToPlatform(str string, platform string) string{
	if(platform == "android"){
		for key, value := range AndroidPlantMap {
			str = strings.Replace(str, key, strconv.Itoa(value), -1);
		}
	}else if(platform == "ios"){
		for key, value := range IosPlantMap {
			str = strings.Replace(str, key, strconv.Itoa(value), -1);
		}
	}
	return str;
}

func initAndroid(){
	     IosPlantMap["peashooter"] = 1001;
          IosPlantMap["sunflower"] = 1002;
          IosPlantMap["wallnut"] = 1003;
          IosPlantMap["potatomine"] = 1004;
          IosPlantMap["cabbagepult"] = 1005;
          IosPlantMap["iceburg"] = 1006;
          IosPlantMap["bloomerang"] = 1007;
        IosPlantMap["twinsunflower"] = 1008;
          IosPlantMap["bonkchoy"] = 1009;
          IosPlantMap["springbean"] = 1010;
          IosPlantMap["spikeweed"] = 1011;
          IosPlantMap["snapdragon"] = 1012;
        IosPlantMap["powerlily"] = 1013;
          IosPlantMap["squash"] = 1014;
          IosPlantMap["chilibean"] = 1015;
          IosPlantMap["splitpea"] = 1016;
          IosPlantMap["jalapeno"] = 1017;
          IosPlantMap["gravebuster"] = 1018;
        IosPlantMap["snowpea"] = 1019;
          IosPlantMap["torchwood"] = 1020;
          IosPlantMap["kernelpult"] = 1021;
          IosPlantMap["lightningreed"] = 1022;
        IosPlantMap["coconutcannon"] = 1023;
          IosPlantMap["melonpult"] = 1024;
          IosPlantMap["peapod"] = 1025;
          IosPlantMap["imitater"] = 1026;
          IosPlantMap["repeater"] = 1027;
          IosPlantMap["spikerock"] = 1028;
        IosPlantMap["tallnut"] = 1029;
        IosPlantMap["threepeater"] = 1030;
          IosPlantMap["wintermelon"] = 1031;
          IosPlantMap["cherry_bomb"] = 1032;
        IosPlantMap["peach"] = 1033;
        IosPlantMap["firegourd"] = 1034;
          IosPlantMap["turnip"] = 1035;
        IosPlantMap["bamboo"] = 1036;
        IosPlantMap["magnifyinggrass"] = 1037;
        IosPlantMap["marigold"] = 1038;
          IosPlantMap["laser_bean"] = 1039;
          IosPlantMap["starfruit"] = 1040;
        IosPlantMap["blover"] = 1041;
        IosPlantMap["empea"] = 1042;
        IosPlantMap["citron"] = 1043;
          IosPlantMap["holonut"] = 1044;
        IosPlantMap["powerplant"] = 1045;
        IosPlantMap["smallcherry"] = 1046;
        IosPlantMap["carrotlauncher"] = 1047;
        IosPlantMap["carrotmissile"] = 1048;
        //dark
        IosPlantMap["puffshroom"] = 1049;
        IosPlantMap["fumeshroom"] = 1050;
        IosPlantMap["hypnoshroom"] = 1051;
        IosPlantMap["sunshroom"] = 1052;
        IosPlantMap["sunbean"] = 1053;
        IosPlantMap["peanut"] = 1054;
        IosPlantMap["magnetshroom"] = 1055;
        IosPlantMap["streetlamp"] = 1056;
        IosPlantMap["coffeebean"] = 1057;
        IosPlantMap["iceshroom"] = 1058;
        IosPlantMap["fireshroom"] = 1059;
        IosPlantMap["oakshooter"] = 1060;
        IosPlantMap["dandelion"] = 1061;
        IosPlantMap["broccoli"] = 1062;
        IosPlantMap["pamegranate"] = 1063;
       
        //beach
        IosPlantMap["lilypad"] = 1064;
        IosPlantMap["bowlingbulb"] = 1065;
        IosPlantMap["tanglekelp"] = 1066;
        IosPlantMap["banana"] = 1067;
        IosPlantMap["guacodile"] = 1068;
        IosPlantMap["homingthistle"] = 1069;
        IosPlantMap["chomper"] = 1070;
        IosPlantMap["lemon"] = 1071;
        IosPlantMap["ghostpepper"] = 1072;
        IosPlantMap["sweetpotato"] = 1073;
        IosPlantMap["cracker"] = 1074;
        IosPlantMap["lotusshower"] = 1075;
        IosPlantMap["sapfling"] = 1076;
       
        //iceAge
        IosPlantMap["hurrikale"] = 1077;
        IosPlantMap["firepeashooter"] = 1078;
        IosPlantMap["hotpotato"] = 1079;
        IosPlantMap["pepperpult"] = 1080;
        IosPlantMap["chardguard"] = 1081;
        IosPlantMap["stunion"] = 1082;
        IosPlantMap["xshot"] = 1083;
        IosPlantMap["rafflesia"] = 1084;
        IosPlantMap["acorn"] = 1085;
       
        IosPlantMap["chestnut"] = 1086;
        IosPlantMap["smallChestnut"] = 1087;
        IosPlantMap["sugarcane"] = 1088;
       
        //skycity
        IosPlantMap["doublesamara"] = 1089;
        IosPlantMap["anthurium"] = 1090;
        IosPlantMap["asparagus"] = 1091;
        IosPlantMap["saucer"] = 1092;
        IosPlantMap["horsebean"] = 1093;
        IosPlantMap["groundcherry"] = 1094;
       
        IosPlantMap["pineapple"] = 1095;
        IosPlantMap["bashopult"] = 1096;
        IosPlantMap["magicshroom"] = 1097;
        IosPlantMap["roseswordman"] = 1098;
        IosPlantMap["electricblueberry"] = 1099;
       
        IosPlantMap["greenturnip"] = 111001;
        IosPlantMap["birthsunflower"] = 111002;
        IosPlantMap["endurian"] = 111003;
        IosPlantMap["pumpkinwitch"] = 111004;
       
        //lostcity
        IosPlantMap["sunpod"]   = 111005;
        IosPlantMap["goldleaf"] = 111006;
        IosPlantMap["sungun"]   = 111007;
        IosPlantMap["akee"]     = 111008;
        IosPlantMap["redstinger"] = 111009;
        IosPlantMap["stallia"]    = 111010;
}

func initIos(){
	AndroidPlantMap["sunflower"] = 1;
    AndroidPlantMap["peashooter"] = 2;
    AndroidPlantMap["wallnut"] = 3;
    AndroidPlantMap["tallnut"] = 4;
    AndroidPlantMap["bonkchoy"] = 5;
    AndroidPlantMap["cabbagepult"] = 6;
    AndroidPlantMap["melonpult"] = 7;
    AndroidPlantMap["cherry_bomb"] = 8;
    AndroidPlantMap["coconutcannon"] = 9;
    AndroidPlantMap["gravebuster"] = 10;
    AndroidPlantMap["iceburg"] = 11;
    AndroidPlantMap["laser_bean"] = 12;
    AndroidPlantMap["potatomine"] = 13;
    AndroidPlantMap["repeater"] = 14;
    AndroidPlantMap["snapdragon"] = 15;
    AndroidPlantMap["spikeweed"] = 16;
    AndroidPlantMap["threepeater"] = 17;
    AndroidPlantMap["torchwood"] = 18;
    AndroidPlantMap["kernelpult"] = 19;
    AndroidPlantMap["springbean"] = 20;
    AndroidPlantMap["snowpea"] = 21;
    AndroidPlantMap["chilibean"] = 22;
    AndroidPlantMap["splitpea"] = 23;
    AndroidPlantMap["lightningreed"] = 24;
    AndroidPlantMap["peapod"] = 25;
    AndroidPlantMap["magnifyinggrass"] = 26;
    AndroidPlantMap["bloomerang"] = 27;
    AndroidPlantMap["holonut"] = 28;
    AndroidPlantMap["empea"] = 29;
    AndroidPlantMap["blover"] = 30;
    AndroidPlantMap["starfruit"] = 31;
    AndroidPlantMap["imitater"] = 32;
    AndroidPlantMap["jalapeno"] = 33;
    AndroidPlantMap["wintermelon"] = 34;
    AndroidPlantMap["twinsunflower"] = 35;
    AndroidPlantMap["marigold"] = 36;
    AndroidPlantMap["spikerock"] = 37;
    AndroidPlantMap["powerlily"] = 38;
    AndroidPlantMap["squash"] = 39;
    AndroidPlantMap["citron"] = 40;
    AndroidPlantMap["powerplant"] = 41;
    AndroidPlantMap["turnip"] = 42;
    AndroidPlantMap["peach"] = 43;
    AndroidPlantMap["firegourd"] = 44;
    AndroidPlantMap["bamboo"] = 45;
    AndroidPlantMap["smallcherry"] = 46;
    AndroidPlantMap["carrotlauncher"] = 47;
    AndroidPlantMap["carrotmissile"] = 48;
    AndroidPlantMap["dandelion"] = 49;
    AndroidPlantMap["broccoli"] = 50;
    //dark
    AndroidPlantMap["puffshroom"] = 51;
    AndroidPlantMap["fumeshroom"] = 52;
    AndroidPlantMap["hypnoshroom"] = 53;
    AndroidPlantMap["sunshroom"] = 54;
    AndroidPlantMap["sunbean"] = 55;
    AndroidPlantMap["peanut"] = 56;
    AndroidPlantMap["magnetshroom"] = 57;
    AndroidPlantMap["streetlamp"] = 58;
    AndroidPlantMap["coffeebean"] = 59;
    AndroidPlantMap["iceshroom"] = 60;
    AndroidPlantMap["fireshroom"] = 61;
    AndroidPlantMap["oakshooter"] = 62;
    AndroidPlantMap["pamegranate"] = 63;
    
    AndroidPlantMap["chomper"] = 64;
    AndroidPlantMap["sweetpotato"] = 65;
    
    //beach
    AndroidPlantMap["tanglekelp"] = 66;
    AndroidPlantMap["banana"] = 67;
    AndroidPlantMap["guacodile"] = 68;
    AndroidPlantMap["homingthistle"] = 69;
    AndroidPlantMap["lilypad"] = 70;
    AndroidPlantMap["lemon"] = 71;
    AndroidPlantMap["ghostpepper"] = 72;
    AndroidPlantMap["bowlingbulb"] = 73;
    AndroidPlantMap["cracker"] = 74;
    AndroidPlantMap["lotusshower"] = 75;
    AndroidPlantMap["sapfling"] = 76;
    //iceAge
    AndroidPlantMap["hurrikale"] = 77;
    AndroidPlantMap["firepeashooter"] = 78;
    AndroidPlantMap["hotpotato"] = 79;
    AndroidPlantMap["pepperpult"] = 80;
    AndroidPlantMap["chardguard"] = 81;
    AndroidPlantMap["stunion"] = 82;
    AndroidPlantMap["rafflesia"] = 83;
    AndroidPlantMap["acorn"] = 84;
    
    //skycity
	AndroidPlantMap["doublesamara"] = 85;
	AndroidPlantMap["anthurium"] = 86;
    AndroidPlantMap["asparagus"] = 87;
    AndroidPlantMap["saucer"] = 88;
    AndroidPlantMap["horsebean"] = 89;
    AndroidPlantMap["groundcherry"] = 90;
    AndroidPlantMap["pineapple"] = 91;
    
    //lostcity
    AndroidPlantMap["sunpod"]   = 92;
    AndroidPlantMap["goldleaf"] = 93;
    AndroidPlantMap["sungun"]   = 94;
    AndroidPlantMap["akee"]     = 95;
    AndroidPlantMap["redstinger"] = 96;
    AndroidPlantMap["stallia"]    = 97;
    
    
    // daily reward plant, start from 300
    AndroidPlantMap["chestnut"] = 300;
    AndroidPlantMap["smallChestnut"] = 301;
    AndroidPlantMap["xshot"] = 302;
    AndroidPlantMap["sugarcane"] = 303;
    AndroidPlantMap["bashopult"] = 304;
    AndroidPlantMap["magicshroom"] = 305;
    AndroidPlantMap["roseswordman"] = 306;
    AndroidPlantMap["electricblueberry"] = 307;
    AndroidPlantMap["birthsunflower"] = 308;
    AndroidPlantMap["greenturnip"] = 309;
    AndroidPlantMap["endurian"] = 310;
    AndroidPlantMap["pumpkinwitch"] = 311;
    AndroidPlantMap["cottonyeti"] = 312;
}
