(function (win, headVar) {

 var doc = win.document,
 domWaiters = [],
 queue      = [], 
 handlers   = {}, 
 assets     = {}, 
 isAsync    = "async" in doc.createElement("script") || "MozAppearance" in doc.documentElement.style || win.opera,
 isHeadReady,
 isDomReady,
 api     = win[headVar] = (win[headVar] || function () { api.ready.apply(null, arguments); }),
 PRELOADING = 1,
 PRELOADED  = 2,
 LOADING    = 3,
 LOADED     = 4;
 api.Info={};
 api.InfoJS="";
 api._js="{{.js}}";

if (isAsync) {
	api.load = function () {
		var args      = arguments,
		    callback = args[args.length - 1],
		    items    = {};

		if (!isFunction(callback)) {
			callback = null;
		}

		each(args, function (item, i) {
				if (item !== callback) {
				item             = getAsset(item);
				items[item.name] = item;

				load(item, callback && i === args.length - 2 ? function () {
					if (allLoaded(items)) {
					one(callback);
					}

					} : null);
				}
				});

		return api;
	};
} else {
	api.load = function () {
		var args = arguments,
		    rest = [].slice.call(args, 1),
		    next = rest[0];
		if (!isHeadReady) {
			queue.push(function () {
					api.load.apply(null, args);
					});

			return api;
		}            
		if (!!next) {
			each(rest, function (item) {
					if (!isFunction(item)) {
					preLoad(getAsset(item));
					}
					});
			load(getAsset(args[0]), isFunction(next) ? next : function () {
					api.load.apply(null, rest);
					});                
		}
		else {
			load(getAsset(args[0]));
		}

		return api;
	};
}
api.js = api.load;
api.ready = function (key, callback) {
	if (key === doc) {
		if (isDomReady) {
			one(callback);
		}
		else {
			domWaiters.push(callback);
		}

		return api;
	}
	if (isFunction(key)) {
		callback = key;
		key      = "ALL";
	}
	if (typeof key !== 'string' || !isFunction(callback)) {
		return api;
	}
	var asset = assets[key];
	if (asset && asset.state === LOADED || key === 'ALL' && allLoaded() && isDomReady) {
		one(callback);
		return api;
	}

	var arr = handlers[key];
	if (!arr) {
		arr = handlers[key] = [callback];
	}
	else {
		arr.push(callback);
	}

	return api;
};
api.ready(doc, function () {

		if (allLoaded()) {
		each(handlers.ALL, function (callback) {
			one(callback);
			});
		}

		if (api.feature) {
		api.feature("domloaded", true);
		}
		});


/* private functions
 *********************/
function noop() {
}

function each(arr, callback) {
	if (!arr) {
		return;
	}

	// arguments special type
	if (typeof arr === 'object') {
		arr = [].slice.call(arr);
	}

	// do the job
	for (var i = 0, l = arr.length; i < l; i++) {
		callback.call(arr, arr[i], i);
	}
}
function is(type, obj) {
	var clas = Object.prototype.toString.call(obj).slice(8, -1);
	return obj !== undefined && obj !== null && clas === type;
}

function isFunction(item) {
	return is("Function", item);
}

function isArray(item) {
	return is("Array", item);
}

function toLabel(url) {
	var items = url.split("/"),
	    name = items[items.length - 1],
	    i    = name.indexOf("?");

	return i !== -1 ? name.substring(0, i) : name;
}

function one(callback) {
	callback = callback || noop;

	if (callback._done) {
		return;
	}

	callback();
	callback._done = 1;
}

function getAsset(item) {
	var asset = {};

	if (typeof item === 'object') {
		for (var label in item) {
			if (!!item[label]) {
				asset = {
				name: label,
				url : item[label]
				};
			}
		}
	}
	else {
		asset = {
		name: toLabel(item),
      		url : item
		};
	}
	var existing = assets[asset.name];
	if (existing && existing.url === asset.url) {
		return existing;
	}

	assets[asset.name] = asset;
	return asset;
}

function allLoaded(items) {
	items = items || assets;

	for (var name in items) {
		if (items.hasOwnProperty(name) && items[name].state !== LOADED) {
			return false;
		}
	}

	return true;
}


function onPreload(asset) {
	asset.state = PRELOADED;

	each(asset.onpreload, function (afterPreload) {
			afterPreload.call();
			});
}

function preLoad(asset, callback) {
	if (asset.state === undefined) {

		asset.state     = PRELOADING;
		asset.onpreload = [];

		loadAsset({ url: asset.url, type: 'cache' }, function () {
				onPreload(asset);
				});
	}
}

function load(asset, callback) {
	callback = callback || noop;

	if (asset.state === LOADED) {
		callback();
		return;
	}

	// INFO: why would we trigger a ready event when its not really loaded yet ?
	if (asset.state === LOADING) {
		api.ready(asset.name, callback);
		return;
	}

	if (asset.state === PRELOADING) {
		asset.onpreload.push(function () {
				load(asset, callback);
				});
		return;
	}

	asset.state = LOADING;

	loadAsset(asset, function () {
			asset.state = LOADED;
			callback();
			each(handlers[asset.name], function (fn) {
				one(fn);
				});
			if (isDomReady && allLoaded()) {
			each(handlers.ALL, function (fn) {
				one(fn);
				});
			}
			});
}
function loadAsset(asset, callback) {
	callback = callback || noop;

	var ele;
	if (/\.css[^\.]*$/.test(asset.url)) {
		ele      = doc.createElement('link');
		ele.type = 'text/' + (asset.type || 'css');
		ele.rel  = 'stylesheet';
		ele.href = asset.url;
	}
	else {
		ele      = doc.createElement('script');
		ele.type = 'text/' + (asset.type || 'javascript');
		ele.src  = asset.url;
	}

	ele.onload  = ele.onreadystatechange = process;
	ele.onerror = error;

	ele.async = false;
	ele.defer = false;

	function error(event) {
		event = event || win.event;
		ele.onload = ele.onreadystatechange = ele.onerror = null;
		callback();
	}

	function process(event) {
		event = event || win.event;
		if (event.type === 'load' || (/loaded|complete/.test(ele.readyState) && (!doc.documentMode || doc.documentMode < 9))) {
			ele.onload = ele.onreadystatechange = ele.onerror = null;
			callback();
		}
	}

	api.writeO(ele);
}

function domReady() {
	if (!doc.body) {
		win.clearTimeout(api.readyTimeout);
		api.readyTimeout = win.setTimeout(domReady, 50);
		return;
	}

	if (!isDomReady) {
		isDomReady = true;
		each(domWaiters, function (fn) {
				one(fn);
				});
	}
}

function domContentLoaded() {
	// W3C
	if (doc.addEventListener) {
		doc.removeEventListener("DOMContentLoaded", domContentLoaded, false);
		domReady();
	}

	// IE
	else if (doc.readyState === "complete") {
		doc.detachEvent("onreadystatechange", domContentLoaded);
		domReady();
	}
};

if (doc.readyState === "complete") {
	domReady();
}

// W3C
else if (doc.addEventListener) {
	doc.addEventListener("DOMContentLoaded", domContentLoaded, false);
	win.addEventListener("load", domReady, false);
}

// IE
else {
	doc.attachEvent("onreadystatechange", domContentLoaded);
	win.attachEvent("onload", domReady);
	var top = false;

	try {
		top = win.frameElement == null && doc.documentElement;
	} catch (e) { }

	if (top && top.doScroll) {
		(function doScrollCheck() {
		 if (!isDomReady) {
		 try {
		 top.doScroll("left");
		 } catch (error) {
		 win.clearTimeout(api.readyTimeout);
		 api.readyTimeout = win.setTimeout(doScrollCheck, 50);
		 return;
		 }
		 domReady();
		 }
		 })();
	}
}

setTimeout(function () {
		isHeadReady = true;
		each(queue, function (fn) {
			fn();
			});

		}, 300);

if (!window.console) window.console = {log: function() {}};
api.log=function(s){
	if(parseInt(api._log_flag)){
		var p = '';
		if (typeof  s=== 'string') {
			p=s;
			if(typeof bcdata_sp !="undefined"){
				p+="&sp="+bcdata_sp;
			}
		}else if (typeof s === 'object'){
			if(typeof bcdata_sp !="undefined"){
				s.sp=bcdata_sp;
			}
			for(var i in s){
				p+=i+"="+escape(s[i])+"&";
			}
		}else{
			p=s;
		}
		var logUrl = 'http://'+api._log+'/stat.log.test' + "?" +p+"&rand="+Math.random() ;
		var img = new Image(1,1);
		img.src = logUrl ;
		img.onload = function(){return;}
	}
}
/*
var PARAMTER_VALUE = null;    
function getParamter(paramName) {    
    if(!PARAMTER_VALUE) {
        PARAMTER_VALUE = new Array();    
        var paramStr = location.search.substring(1);    
        var paramArr = paramStr.split("&");    
        var len = paramArr.length;    
        var tempArr;    
        for(var i = 0; i < len; i++) {    
            tempArr = paramArr[i].split("=");    
            PARAMTER_VALUE[tempArr[0]] = tempArr[1];    
        }    
    }    
    var paramValue = PARAMTER_VALUE[paramName];    
    if(paramValue) {    
        return paramValue;    
    }    
} 
*/

api._isMobile=function() {
    var a = navigator.userAgent||navigator.vendor||window.opera;
    return /(android|bb\d+|meego).+mobile|avantgo|bada\/|blackberry|blazer|compal|elaine|fennec|hiptop|iemobile|ip(hone|od)|iris|kindle|lge |maemo|midp|mmp|mobile.+firefox|netfront|opera m(ob|in)i|palm( os)?|phone|p(ixi|re)\/|plucker|pocket|psp|series(4|6)0|symbian|treo|up\.(browser|link)|vodafone|wap|windows (ce|phone)|xda|xiino/i.test(a)||/1207|6310|6590|3gso|4thp|50[1-6]i|770s|802s|a wa|abac|ac(er|oo|s\-)|ai(ko|rn)|al(av|ca|co)|amoi|an(ex|ny|yw)|aptu|ar(ch|go)|as(te|us)|attw|au(di|\-m|r |s )|avan|be(ck|ll|nq)|bi(lb|rd)|bl(ac|az)|br(e|v)w|bumb|bw\-(n|u)|c55\/|capi|ccwa|cdm\-|cell|chtm|cldc|cmd\-|co(mp|nd)|craw|da(it|ll|ng)|dbte|dc\-s|devi|dica|dmob|do(c|p)o|ds(12|\-d)|el(49|ai)|em(l2|ul)|er(ic|k0)|esl8|ez([4-7]0|os|wa|ze)|fetc|fly(\-|_)|g1 u|g560|gene|gf\-5|g\-mo|go(\.w|od)|gr(ad|un)|haie|hcit|hd\-(m|p|t)|hei\-|hi(pt|ta)|hp( i|ip)|hs\-c|ht(c(\-| |_|a|g|p|s|t)|tp)|hu(aw|tc)|i\-(20|go|ma)|i230|iac( |\-|\/)|ibro|idea|ig01|ikom|im1k|inno|ipaq|iris|ja(t|v)a|jbro|jemu|jigs|kddi|keji|kgt( |\/)|klon|kpt |kwc\-|kyo(c|k)|le(no|xi)|lg( g|\/(k|l|u)|50|54|\-[a-w])|libw|lynx|m1\-w|m3ga|m50\/|ma(te|ui|xo)|mc(01|21|ca)|m\-cr|me(rc|ri)|mi(o8|oa|ts)|mmef|mo(01|02|bi|de|do|t(\-| |o|v)|zz)|mt(50|p1|v )|mwbp|mywa|n10[0-2]|n20[2-3]|n30(0|2)|n50(0|2|5)|n7(0(0|1)|10)|ne((c|m)\-|on|tf|wf|wg|wt)|nok(6|i)|nzph|o2im|op(ti|wv)|oran|owg1|p800|pan(a|d|t)|pdxg|pg(13|\-([1-8]|c))|phil|pire|pl(ay|uc)|pn\-2|po(ck|rt|se)|prox|psio|pt\-g|qa\-a|qc(07|12|21|32|60|\-[2-7]|i\-)|qtek|r380|r600|raks|rim9|ro(ve|zo)|s55\/|sa(ge|ma|mm|ms|ny|va)|sc(01|h\-|oo|p\-)|sdk\/|se(c(\-|0|1)|47|mc|nd|ri)|sgh\-|shar|sie(\-|m)|sk\-0|sl(45|id)|sm(al|ar|b3|it|t5)|so(ft|ny)|sp(01|h\-|v\-|v )|sy(01|mb)|t2(18|50)|t6(00|10|18)|ta(gt|lk)|tcl\-|tdg\-|tel(i|m)|tim\-|t\-mo|to(pl|sh)|ts(70|m\-|m3|m5)|tx\-9|up(\.b|g1|si)|utst|v400|v750|veri|vi(rg|te)|vk(40|5[0-3]|\-v)|vm40|voda|vulc|vx(52|53|60|61|70|80|81|83|85|98)|w3c(\-| )|webc|whit|wi(g |nc|nw)|wmlb|wonu|x700|yas\-|your|zeto|zte\-/i.test(a.substr(0,4));
}
api.isMobile=function(){
	var u = window.navigator.userAgent.toLowerCase();
	var h = '';
	if((/AppleWebKit.*mobile/i.test(u)) || (api._isMobile())  || (/android/i.test(u)) || (/MIDP|SymbianOS|NOKIA|SAMSUNG|LG|NEC|TCL|Alcatel|BIRD|DBTEL|Dopod|PHILIPS|HAIER|LENOVO|MOT-|Nokia|SonyEricsson|SIE-|Amoi|ZTE/.test(u))) {
		return true;
	}else{
		return false;
	}
}

})(window,"xx");
