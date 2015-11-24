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
		             var obj = JSON.parse(data);
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
	   param['value'] = $('#parameters')
	   $.post('./view',
		        param,
		        function (data) {
				     try {
			             var obj = JSON.parse(data);
			             showDataInTextField(data);
			         } catch (e) {
			             console.log(e);
			         }
		        }
		);
}

function showDataInTextField(data) {
	  listHTML = data;
	  var listHTML = "<textarea name=\"MSG\" id="" cols=100 rows=50>";
	  listHTML += data;
	  listHTML += "</textarea>";
      $('#result').html(listHTML);
}