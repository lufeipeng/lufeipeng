apiConfig = {
    "egypt_initpos": {
    	 "apiDescription": "埃及世界",
    	   "parameters": {
               "worldName": "egypt_initpos",
           }
    },
    "pirate_initpos": {
    	 "apiDescription": "海盗关卡",
    	   "parameters": {
    		   "worldName": "pirate_initpos",
           }
    },
    "kongfu_initpos": {
    	"apiDescription": "功夫世界",
    	   "parameters": {
    		   "worldName": "kongfu_initpos",
           }
    },
    "egypt_levelup": {
   	 "apiDescription": "埃及世界升级",
   	   "parameters": {
              "worldName": "egypt_levelup",
          }
   },
   "pirate_levelup": {
   	 "apiDescription": "海盗关卡升级",
   	   "parameters": {
   		   "worldName": "pirate_levelup",
          }
   },
   "kongfu_levelup": {
   	"apiDescription": "功夫世界升级",
   	   "parameters": {
   		   "worldName": "kongfu_levelup",
          }
   }
};

function switchAPI(el) {
    apiName = $(el).attr('id');
    var config = apiConfig[apiName];

    var h = '', input;
    
    for (var i in config['parameters']) {
        h += '<div>';
        h += i;
        h += ' : ';
        h += '<input class="paramInput" id="param-' + i + "\" value='" + config['parameters'][i] + "' style=\"width: 280px;\">";
        h += '</div>';
    }

    $('#apiDescription').html(config['apiDescription']);
    $('#parameters').html(h);
    callView();
}

function callView() {
   param = {};
   $(".paramInput").each(function (i, n) {
       n = $(n);
       id = n.attr('id').substr(6);
       param[id] = n.val();
   });
   $.get('./view',
	        param,
	        function (data) {
			     try {
		             showDataInTextField(data);
		         } catch (e) {
		             console.log(e);
		         }
	        }
	);
}

function callSave() {
	param = {};
	   $(".paramInput").each(function (i, n) {
	       n = $(n);
	       id = n.attr('id').substr(6);
	       param[id] = n.val();
	   });
	   param['value'] = $('#defenceStr').val();
	   
	   try {
           var obj = JSON.parse(param['value']);
       } catch (e) {
    	   alert("小伙，保存的阵型Json数据有问题!!")
          return;
       }
	   
	   $.post('./view',
		        param,
		        function (data) {
				     try {
			             showDataInTextField(data);
			         } catch (e) {
			             console.log(e);
			         }
		        }
		);
}

function viewDefence(platform) {
	param = {};
	   $(".paramInput").each(function (i, n) {
	       n = $(n);
	       id = n.attr('id').substr(6);
	       param[id] = n.val();
	   });
	   param['value'] = $('#defenceStr').val();
	   param['platform'] = platform;
	   
	   try {
           var obj = JSON.parse(param['value']);
       } catch (e) {
    	   alert("小伙，保存的阵型Json数据有问题!!")
          return;
       }
	   
	   $.post('./defenceview',
		        param,
		        function (data) {
				     try {
			             showDataInTextField(data);
			         } catch (e) {
			             console.log(e);
			         }
		        }
		);
}

function showDataInTextField(data) {
      var obj = JSON.parse(data);
    
	  var listHTML = "<textarea name=\"MSG\" id=\"defenceStr\" style=\"width: 100%;height: 100%\">";
	  listHTML += obj.defenceStr;
	  listHTML += "</textarea>";
      $('#resultLeft').html(listHTML);
      
      
	  var listHTML = "<textarea name=\"MSG\" id=\"defenceInfo\" style=\"width: 100%;height: 100%\">";
	  listHTML += obj.detailInfo;
	  listHTML += "</textarea>";
      $('#resultRight').html(listHTML);

}