(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["about"],{"0a06":function(e,t,n){"use strict";var r=n("c532"),a=n("30b5"),o=n("f6b49"),s=n("5270"),i=n("4a7b");function c(e){this.defaults=e,this.interceptors={request:new o,response:new o}}c.prototype.request=function(e){"string"===typeof e?(e=arguments[1]||{},e.url=arguments[0]):e=e||{},e=i(this.defaults,e),e.method?e.method=e.method.toLowerCase():this.defaults.method?e.method=this.defaults.method.toLowerCase():e.method="get";var t=[s,void 0],n=Promise.resolve(e);this.interceptors.request.forEach((function(e){t.unshift(e.fulfilled,e.rejected)})),this.interceptors.response.forEach((function(e){t.push(e.fulfilled,e.rejected)}));while(t.length)n=n.then(t.shift(),t.shift());return n},c.prototype.getUri=function(e){return e=i(this.defaults,e),a(e.url,e.params,e.paramsSerializer).replace(/^\?/,"")},r.forEach(["delete","get","head","options"],(function(e){c.prototype[e]=function(t,n){return this.request(r.merge(n||{},{method:e,url:t}))}})),r.forEach(["post","put","patch"],(function(e){c.prototype[e]=function(t,n,a){return this.request(r.merge(a||{},{method:e,url:t,data:n}))}})),e.exports=c},"0df6":function(e,t,n){"use strict";e.exports=function(e){return function(t){return e.apply(null,t)}}},1315:function(e,t,n){"use strict";var r={URL:"http://192.168.1.15:8001/"};t["a"]=r},"1d2b":function(e,t,n){"use strict";e.exports=function(e,t){return function(){for(var n=new Array(arguments.length),r=0;r<n.length;r++)n[r]=arguments[r];return e.apply(t,n)}}},2444:function(e,t,n){"use strict";(function(t){var r=n("c532"),a=n("c8af"),o={"Content-Type":"application/x-www-form-urlencoded"};function s(e,t){!r.isUndefined(e)&&r.isUndefined(e["Content-Type"])&&(e["Content-Type"]=t)}function i(){var e;return("undefined"!==typeof XMLHttpRequest||"undefined"!==typeof t&&"[object process]"===Object.prototype.toString.call(t))&&(e=n("b50d")),e}var c={adapter:i(),transformRequest:[function(e,t){return a(t,"Accept"),a(t,"Content-Type"),r.isFormData(e)||r.isArrayBuffer(e)||r.isBuffer(e)||r.isStream(e)||r.isFile(e)||r.isBlob(e)?e:r.isArrayBufferView(e)?e.buffer:r.isURLSearchParams(e)?(s(t,"application/x-www-form-urlencoded;charset=utf-8"),e.toString()):r.isObject(e)?(s(t,"application/json;charset=utf-8"),JSON.stringify(e)):e}],transformResponse:[function(e){if("string"===typeof e)try{e=JSON.parse(e)}catch(t){}return e}],timeout:0,xsrfCookieName:"XSRF-TOKEN",xsrfHeaderName:"X-XSRF-TOKEN",maxContentLength:-1,validateStatus:function(e){return e>=200&&e<300},headers:{common:{Accept:"application/json, text/plain, */*"}}};r.forEach(["delete","get","head"],(function(e){c.headers[e]={}})),r.forEach(["post","put","patch"],(function(e){c.headers[e]=r.merge(o)})),e.exports=c}).call(this,n("4362"))},"2d83":function(e,t,n){"use strict";var r=n("387f");e.exports=function(e,t,n,a,o){var s=new Error(e);return r(s,t,n,a,o)}},"2e67":function(e,t,n){"use strict";e.exports=function(e){return!(!e||!e.__CANCEL__)}},3086:function(e,t,n){"use strict";n.r(t);var r=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"components-layout-demo-basic"}},[n("a-layout",[n("a-layout-header",[e._v("Welcome To Explain SQL")]),n("a-layout",[n("a-layout-content",[n("div",{attrs:{id:"components-form-demo-advanced-search"}},[n("a-form",{staticClass:"ant-advanced-search-form",attrs:{form:e.form},on:{submit:e.handleSearch}},[n("a-row",{attrs:{gutter:24}},[n("a-form-item",[n("a-textarea",{attrs:{placeholder:"Please enter sql",rows:4},model:{value:e.sql,callback:function(t){e.sql=t},expression:"sql"}})],1)],1),n("a-row",[n("a-col",{style:{textAlign:"right"},attrs:{span:24}},[n("a-button",{attrs:{type:"primary","html-type":"submit"}},[e._v("Search")]),n("a-button",{style:{marginLeft:"8px"},on:{click:e.handleReset}},[e._v("Clear")])],1)],1)],1),n("div",[n("a-table",{attrs:{columns:e.columns,"data-source":e.slowRes,scroll:{x:1500,y:300}}},[n("a",{attrs:{slot:"action"},slot:"action"},[e._v("action")])])],1)],1)])],1),n("a-layout-footer",[e.showerr?n("div",[n("a-alert",{attrs:{message:"Error Tips",description:e.result,type:"error",closable:""},on:{close:e.onClose}})],1):e._e()])],1)],1)},a=[],o=n("bc3a"),s=n.n(o),i=n("1315"),c=[{title:"table",width:200,dataIndex:"table",key:"table",fixed:"left"},{title:"id",width:100,dataIndex:"id",key:"id",fixed:"left"},{title:"filtered",dataIndex:"filtered",key:"1",width:150},{title:"key",dataIndex:"key",key:"key",width:150},{title:"key_len",dataIndex:"key_len",key:"key_len",width:150},{title:"partitions",dataIndex:"partitions",key:"partitions",width:150},{title:"possible_keys",dataIndex:"possible_keys",key:"possible_keys",width:150},{title:"ref",dataIndex:"ref",key:"ref",width:150},{title:"rows",dataIndex:"rows",key:"7",width:150},{title:"select_type",dataIndex:"select_type",key:"select_type",width:150},{title:"type",dataIndex:"type",key:"type",width:150},{title:"Extra",dataIndex:"Extra",key:"Extra",width:150},{title:"Action",key:"operation",fixed:"right",width:100,scopedSlots:{customRender:"action"}}],u={data:function(){return{showerr:!1,slowRes:[],columns:c,size:"default",expand:!1,result:"",sql:"",form:this.$form.createForm(this,{name:"advanced_search"})}},computed:{count:function(){return this.expand?11:7}},updated:function(){},methods:{onClose:function(e){e.preventDefault(),this.showerr=!1},onChange:function(e){console.log("size checked",e.target.value),this.size=e.target.value},explainSql:function(){var e=this,t=new FormData;t.append("sql",this.sql),s.a.post(i["a"].URL+"action/explainSql",t).then((function(t){200==t.data.code?(e.slowRes=t.data.data,e.showerr=!1):(e.result=t.data.data.Message,e.showerr=!0)}))},handleSearch:function(e){e.preventDefault(),this.explainSql()},handleReset:function(){this.sql=""},toggle:function(){this.expand=!this.expand}}},f=u,l=(n("6124"),n("2877")),d=Object(l["a"])(f,r,a,!1,null,null,null);t["default"]=d.exports},"30b5":function(e,t,n){"use strict";var r=n("c532");function a(e){return encodeURIComponent(e).replace(/%40/gi,"@").replace(/%3A/gi,":").replace(/%24/g,"$").replace(/%2C/gi,",").replace(/%20/g,"+").replace(/%5B/gi,"[").replace(/%5D/gi,"]")}e.exports=function(e,t,n){if(!t)return e;var o;if(n)o=n(t);else if(r.isURLSearchParams(t))o=t.toString();else{var s=[];r.forEach(t,(function(e,t){null!==e&&"undefined"!==typeof e&&(r.isArray(e)?t+="[]":e=[e],r.forEach(e,(function(e){r.isDate(e)?e=e.toISOString():r.isObject(e)&&(e=JSON.stringify(e)),s.push(a(t)+"="+a(e))})))})),o=s.join("&")}if(o){var i=e.indexOf("#");-1!==i&&(e=e.slice(0,i)),e+=(-1===e.indexOf("?")?"?":"&")+o}return e}},"387f":function(e,t,n){"use strict";e.exports=function(e,t,n,r,a){return e.config=t,n&&(e.code=n),e.request=r,e.response=a,e.isAxiosError=!0,e.toJSON=function(){return{message:this.message,name:this.name,description:this.description,number:this.number,fileName:this.fileName,lineNumber:this.lineNumber,columnNumber:this.columnNumber,stack:this.stack,config:this.config,code:this.code}},e}},3934:function(e,t,n){"use strict";var r=n("c532");e.exports=r.isStandardBrowserEnv()?function(){var e,t=/(msie|trident)/i.test(navigator.userAgent),n=document.createElement("a");function a(e){var r=e;return t&&(n.setAttribute("href",r),r=n.href),n.setAttribute("href",r),{href:n.href,protocol:n.protocol?n.protocol.replace(/:$/,""):"",host:n.host,search:n.search?n.search.replace(/^\?/,""):"",hash:n.hash?n.hash.replace(/^#/,""):"",hostname:n.hostname,port:n.port,pathname:"/"===n.pathname.charAt(0)?n.pathname:"/"+n.pathname}}return e=a(window.location.href),function(t){var n=r.isString(t)?a(t):t;return n.protocol===e.protocol&&n.host===e.host}}():function(){return function(){return!0}}()},"467f":function(e,t,n){"use strict";var r=n("2d83");e.exports=function(e,t,n){var a=n.config.validateStatus;!a||a(n.status)?e(n):t(r("Request failed with status code "+n.status,n.config,null,n.request,n))}},"4a7b":function(e,t,n){"use strict";var r=n("c532");e.exports=function(e,t){t=t||{};var n={},a=["url","method","params","data"],o=["headers","auth","proxy"],s=["baseURL","url","transformRequest","transformResponse","paramsSerializer","timeout","withCredentials","adapter","responseType","xsrfCookieName","xsrfHeaderName","onUploadProgress","onDownloadProgress","maxContentLength","validateStatus","maxRedirects","httpAgent","httpsAgent","cancelToken","socketPath"];r.forEach(a,(function(e){"undefined"!==typeof t[e]&&(n[e]=t[e])})),r.forEach(o,(function(a){r.isObject(t[a])?n[a]=r.deepMerge(e[a],t[a]):"undefined"!==typeof t[a]?n[a]=t[a]:r.isObject(e[a])?n[a]=r.deepMerge(e[a]):"undefined"!==typeof e[a]&&(n[a]=e[a])})),r.forEach(s,(function(r){"undefined"!==typeof t[r]?n[r]=t[r]:"undefined"!==typeof e[r]&&(n[r]=e[r])}));var i=a.concat(o).concat(s),c=Object.keys(t).filter((function(e){return-1===i.indexOf(e)}));return r.forEach(c,(function(r){"undefined"!==typeof t[r]?n[r]=t[r]:"undefined"!==typeof e[r]&&(n[r]=e[r])})),n}},5270:function(e,t,n){"use strict";var r=n("c532"),a=n("c401"),o=n("2e67"),s=n("2444");function i(e){e.cancelToken&&e.cancelToken.throwIfRequested()}e.exports=function(e){i(e),e.headers=e.headers||{},e.data=a(e.data,e.headers,e.transformRequest),e.headers=r.merge(e.headers.common||{},e.headers[e.method]||{},e.headers),r.forEach(["delete","get","head","post","put","patch","common"],(function(t){delete e.headers[t]}));var t=e.adapter||s.adapter;return t(e).then((function(t){return i(e),t.data=a(t.data,t.headers,e.transformResponse),t}),(function(t){return o(t)||(i(e),t&&t.response&&(t.response.data=a(t.response.data,t.response.headers,e.transformResponse))),Promise.reject(t)}))}},"5a44":function(e,t,n){},6124:function(e,t,n){"use strict";var r=n("f2b8"),a=n.n(r);a.a},"7a77":function(e,t,n){"use strict";function r(e){this.message=e}r.prototype.toString=function(){return"Cancel"+(this.message?": "+this.message:"")},r.prototype.__CANCEL__=!0,e.exports=r},"7aac":function(e,t,n){"use strict";var r=n("c532");e.exports=r.isStandardBrowserEnv()?function(){return{write:function(e,t,n,a,o,s){var i=[];i.push(e+"="+encodeURIComponent(t)),r.isNumber(n)&&i.push("expires="+new Date(n).toGMTString()),r.isString(a)&&i.push("path="+a),r.isString(o)&&i.push("domain="+o),!0===s&&i.push("secure"),document.cookie=i.join("; ")},read:function(e){var t=document.cookie.match(new RegExp("(^|;\\s*)("+e+")=([^;]*)"));return t?decodeURIComponent(t[3]):null},remove:function(e){this.write(e,"",Date.now()-864e5)}}}():function(){return{write:function(){},read:function(){return null},remove:function(){}}}()},"7dbd":function(e,t,n){"use strict";var r=n("5a44"),a=n.n(r);a.a},"83b9":function(e,t,n){"use strict";var r=n("d925"),a=n("e683");e.exports=function(e,t){return e&&!r(t)?a(e,t):t}},"8df4b":function(e,t,n){"use strict";var r=n("7a77");function a(e){if("function"!==typeof e)throw new TypeError("executor must be a function.");var t;this.promise=new Promise((function(e){t=e}));var n=this;e((function(e){n.reason||(n.reason=new r(e),t(n.reason))}))}a.prototype.throwIfRequested=function(){if(this.reason)throw this.reason},a.source=function(){var e,t=new a((function(t){e=t}));return{token:t,cancel:e}},e.exports=a},b50d:function(e,t,n){"use strict";var r=n("c532"),a=n("467f"),o=n("30b5"),s=n("83b9"),i=n("c345"),c=n("3934"),u=n("2d83");e.exports=function(e){return new Promise((function(t,f){var l=e.data,d=e.headers;r.isFormData(l)&&delete d["Content-Type"];var p=new XMLHttpRequest;if(e.auth){var h=e.auth.username||"",m=e.auth.password||"";d.Authorization="Basic "+btoa(h+":"+m)}var y=s(e.baseURL,e.url);if(p.open(e.method.toUpperCase(),o(y,e.params,e.paramsSerializer),!0),p.timeout=e.timeout,p.onreadystatechange=function(){if(p&&4===p.readyState&&(0!==p.status||p.responseURL&&0===p.responseURL.indexOf("file:"))){var n="getAllResponseHeaders"in p?i(p.getAllResponseHeaders()):null,r=e.responseType&&"text"!==e.responseType?p.response:p.responseText,o={data:r,status:p.status,statusText:p.statusText,headers:n,config:e,request:p};a(t,f,o),p=null}},p.onabort=function(){p&&(f(u("Request aborted",e,"ECONNABORTED",p)),p=null)},p.onerror=function(){f(u("Network Error",e,null,p)),p=null},p.ontimeout=function(){var t="timeout of "+e.timeout+"ms exceeded";e.timeoutErrorMessage&&(t=e.timeoutErrorMessage),f(u(t,e,"ECONNABORTED",p)),p=null},r.isStandardBrowserEnv()){var g=n("7aac"),v=(e.withCredentials||c(y))&&e.xsrfCookieName?g.read(e.xsrfCookieName):void 0;v&&(d[e.xsrfHeaderName]=v)}if("setRequestHeader"in p&&r.forEach(d,(function(e,t){"undefined"===typeof l&&"content-type"===t.toLowerCase()?delete d[t]:p.setRequestHeader(t,e)})),r.isUndefined(e.withCredentials)||(p.withCredentials=!!e.withCredentials),e.responseType)try{p.responseType=e.responseType}catch(x){if("json"!==e.responseType)throw x}"function"===typeof e.onDownloadProgress&&p.addEventListener("progress",e.onDownloadProgress),"function"===typeof e.onUploadProgress&&p.upload&&p.upload.addEventListener("progress",e.onUploadProgress),e.cancelToken&&e.cancelToken.promise.then((function(e){p&&(p.abort(),f(e),p=null)})),void 0===l&&(l=null),p.send(l)}))}},bc3a:function(e,t,n){e.exports=n("cee4")},c345:function(e,t,n){"use strict";var r=n("c532"),a=["age","authorization","content-length","content-type","etag","expires","from","host","if-modified-since","if-unmodified-since","last-modified","location","max-forwards","proxy-authorization","referer","retry-after","user-agent"];e.exports=function(e){var t,n,o,s={};return e?(r.forEach(e.split("\n"),(function(e){if(o=e.indexOf(":"),t=r.trim(e.substr(0,o)).toLowerCase(),n=r.trim(e.substr(o+1)),t){if(s[t]&&a.indexOf(t)>=0)return;s[t]="set-cookie"===t?(s[t]?s[t]:[]).concat([n]):s[t]?s[t]+", "+n:n}})),s):s}},c401:function(e,t,n){"use strict";var r=n("c532");e.exports=function(e,t,n){return r.forEach(n,(function(n){e=n(e,t)})),e}},c532:function(e,t,n){"use strict";var r=n("1d2b"),a=Object.prototype.toString;function o(e){return"[object Array]"===a.call(e)}function s(e){return"undefined"===typeof e}function i(e){return null!==e&&!s(e)&&null!==e.constructor&&!s(e.constructor)&&"function"===typeof e.constructor.isBuffer&&e.constructor.isBuffer(e)}function c(e){return"[object ArrayBuffer]"===a.call(e)}function u(e){return"undefined"!==typeof FormData&&e instanceof FormData}function f(e){var t;return t="undefined"!==typeof ArrayBuffer&&ArrayBuffer.isView?ArrayBuffer.isView(e):e&&e.buffer&&e.buffer instanceof ArrayBuffer,t}function l(e){return"string"===typeof e}function d(e){return"number"===typeof e}function p(e){return null!==e&&"object"===typeof e}function h(e){return"[object Date]"===a.call(e)}function m(e){return"[object File]"===a.call(e)}function y(e){return"[object Blob]"===a.call(e)}function g(e){return"[object Function]"===a.call(e)}function v(e){return p(e)&&g(e.pipe)}function x(e){return"undefined"!==typeof URLSearchParams&&e instanceof URLSearchParams}function b(e){return e.replace(/^\s*/,"").replace(/\s*$/,"")}function w(){return("undefined"===typeof navigator||"ReactNative"!==navigator.product&&"NativeScript"!==navigator.product&&"NS"!==navigator.product)&&("undefined"!==typeof window&&"undefined"!==typeof document)}function S(e,t){if(null!==e&&"undefined"!==typeof e)if("object"!==typeof e&&(e=[e]),o(e))for(var n=0,r=e.length;n<r;n++)t.call(null,e[n],n,e);else for(var a in e)Object.prototype.hasOwnProperty.call(e,a)&&t.call(null,e[a],a,e)}function k(){var e={};function t(t,n){"object"===typeof e[n]&&"object"===typeof t?e[n]=k(e[n],t):e[n]=t}for(var n=0,r=arguments.length;n<r;n++)S(arguments[n],t);return e}function _(){var e={};function t(t,n){"object"===typeof e[n]&&"object"===typeof t?e[n]=_(e[n],t):e[n]="object"===typeof t?_({},t):t}for(var n=0,r=arguments.length;n<r;n++)S(arguments[n],t);return e}function I(e,t,n){return S(t,(function(t,a){e[a]=n&&"function"===typeof t?r(t,n):t})),e}e.exports={isArray:o,isArrayBuffer:c,isBuffer:i,isFormData:u,isArrayBufferView:f,isString:l,isNumber:d,isObject:p,isUndefined:s,isDate:h,isFile:m,isBlob:y,isFunction:g,isStream:v,isURLSearchParams:x,isStandardBrowserEnv:w,forEach:S,merge:k,deepMerge:_,extend:I,trim:b}},c8af:function(e,t,n){"use strict";var r=n("c532");e.exports=function(e,t){r.forEach(e,(function(n,r){r!==t&&r.toUpperCase()===t.toUpperCase()&&(e[t]=n,delete e[r])}))}},cee4:function(e,t,n){"use strict";var r=n("c532"),a=n("1d2b"),o=n("0a06"),s=n("4a7b"),i=n("2444");function c(e){var t=new o(e),n=a(o.prototype.request,t);return r.extend(n,o.prototype,t),r.extend(n,t),n}var u=c(i);u.Axios=o,u.create=function(e){return c(s(u.defaults,e))},u.Cancel=n("7a77"),u.CancelToken=n("8df4b"),u.isCancel=n("2e67"),u.all=function(e){return Promise.all(e)},u.spread=n("0df6"),e.exports=u,e.exports.default=u},d925:function(e,t,n){"use strict";e.exports=function(e){return/^([a-z][a-z\d\+\-\.]*:)?\/\//i.test(e)}},db27:function(e,t,n){"use strict";n.r(t);var r=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"components-layout-demo-basic"}},[n("a-layout",[n("a-layout",[n("a-layout-header",[e._v("log")]),n("a-layout-content",[n("div",[e.showerr?n("a-alert",{attrs:{message:"Warning",description:"当前数据库未打开慢SQL日志收集功能，请打开后重试.",type:"warning","show-icon":""}}):e._e()],1),n("a-descriptions",[n("a-descriptions-item",{attrs:{label:"Config Info"}},[n("p",{domProps:{innerHTML:e._s(e.slowRes.FileLog)}})])],1)],1)],1)],1)],1)},a=[],o=n("bc3a"),s=n.n(o),i=n("1315"),c=[{title:"table",width:200,dataIndex:"table",key:"table",fixed:"left"},{title:"id",width:100,dataIndex:"id",key:"id",fixed:"left"},{title:"filtered",dataIndex:"filtered",key:"1",width:150},{title:"key",dataIndex:"key",key:"key",width:150},{title:"key_len",dataIndex:"key_len",key:"key_len",width:150},{title:"partitions",dataIndex:"partitions",key:"partitions",width:150},{title:"possible_keys",dataIndex:"possible_keys",key:"possible_keys",width:150},{title:"ref",dataIndex:"ref",key:"ref",width:150},{title:"rows",dataIndex:"rows",key:"7",width:150},{title:"select_type",dataIndex:"select_type",key:"select_type",width:150},{title:"type",dataIndex:"type",key:"type",width:150},{title:"Extra",dataIndex:"Extra",key:"Extra",width:150},{title:"Action",key:"operation",fixed:"right",width:100,scopedSlots:{customRender:"action"}}],u={data:function(){return{showerr:!1,slowRes:{},columns:c,size:"default",expand:!1,result:"",sql:"",form:this.$form.createForm(this,{name:"advanced_search"})}},mounted:function(){this.getSlowStatus()},computed:{count:function(){return this.expand?11:7}},updated:function(){},methods:{getSlowStatus:function(){var e=this;s.a.post(i["a"].URL+"action/slowStatus",null).then((function(t){1!=t.data.status?e.showerr=!0:(e.getSlowInfo(),e.showerr=!1)}))},getSlowInfo:function(){var e=this;s.a.post(i["a"].URL+"action/slowInfo",null).then((function(t){200==t.data.code&&(e.slowRes=t.data.data)}))},onClose:function(e){console.log(e,"I was closed.")},onChange:function(e){console.log("size checked",e.target.value),this.size=e.target.value},handleSearch:function(e){e.preventDefault(),console.log(this.sql),this.getDBInfo()},handleReset:function(){this.sql=""},toggle:function(){this.expand=!this.expand}}},f=u,l=(n("7dbd"),n("2877")),d=Object(l["a"])(f,r,a,!1,null,null,null);t["default"]=d.exports},e683:function(e,t,n){"use strict";e.exports=function(e,t){return t?e.replace(/\/+$/,"")+"/"+t.replace(/^\/+/,""):e}},f2b8:function(e,t,n){},f6b49:function(e,t,n){"use strict";var r=n("c532");function a(){this.handlers=[]}a.prototype.use=function(e,t){return this.handlers.push({fulfilled:e,rejected:t}),this.handlers.length-1},a.prototype.eject=function(e){this.handlers[e]&&(this.handlers[e]=null)},a.prototype.forEach=function(e){r.forEach(this.handlers,(function(t){null!==t&&e(t)}))},e.exports=a},f820:function(e,t,n){"use strict";n.r(t);var r=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("a-divider",{attrs:{orientation:"left"}},[e._v("DataBase Status")]),n("a-row",{attrs:{type:"flex"}},[n("a-col",{attrs:{flex:2}},[n("a-descriptions",{attrs:{bordered:""}},[n("a-descriptions-item",{attrs:{label:"DataBase Name"}},[e._v(e._s(e.DbInfo.DataBaseName))]),n("a-descriptions-item",{attrs:{label:"Engine"}},[e._v(e._s(this.DbInfo.Engine))]),n("a-descriptions-item",{attrs:{label:"Port"}},[e._v(e._s(this.DbInfo.Port))]),n("a-descriptions-item",{attrs:{label:"DataSize"}},[e._v(e._s(this.DbInfo.DataSize))]),n("a-descriptions-item",{attrs:{label:"Uptime",span:2}},[e._v(e._s(this.DbInfo.Uptime))]),n("a-descriptions-item",{attrs:{label:"Processlist"}},[e._v(e._s(this.ProcesslistSize))]),n("a-descriptions-item",{attrs:{label:"QPS"}},[e._v(e._s(this.QPS)+"m/s")]),n("a-descriptions-item",{attrs:{label:"TPS"}},[e._v(e._s(this.TPS)+"m/s")]),n("a-descriptions-item",{attrs:{label:"SlowQueryLogFile Path"}},[e._v(e._s(this.DbInfo.SlowQueryLogFile))]),n("a-descriptions-item",{attrs:{label:"LongQuery Time"}},[e._v(e._s(this.DbInfo.LongQueryTime))]),n("a-descriptions-item",{attrs:{label:"SlowQueryLog Status"}},[e._v(e._s(this.DbInfo.SlowQueryLog))]),n("a-descriptions-item",{attrs:{label:"Status",span:3}},[n("a-badge",{attrs:{status:"processing"},domProps:{textContent:e._s(this.DbInfo.Status)}})],1)],1)],1),n("a-col",{attrs:{flex:3}},[n("div",{staticStyle:{"margin-left":"10%"}},[n("a-row",{attrs:{gutter:14}},[n("a-col",{attrs:{span:8}},[n("a-card",[n("a-statistic",{staticStyle:{"margin-right":"10px"},attrs:{title:"KeyBufferReadHits",value:e.keyBufferRead,precision:2,suffix:"%","value-style":{color:"#cf1322"}},scopedSlots:e._u([{key:"prefix",fn:function(){},proxy:!0}])})],1)],1),n("a-col",{attrs:{span:8}},[n("a-card",[n("a-statistic",{staticClass:"demo-class",attrs:{title:"KeyBufferWriteHits",value:e.keyBufferWrite,precision:2,suffix:"%","value-style":{color:"#cf1322"}},scopedSlots:e._u([{key:"prefix",fn:function(){},proxy:!0}])})],1)],1)],1)],1),n("div",{staticStyle:{"margin-left":"10%"}},[n("a-row",{attrs:{gutter:14}},[n("a-col",{attrs:{span:8}},[n("a-card",[n("a-statistic",{staticStyle:{"margin-right":"10px"},attrs:{title:"QueryCacheHits",value:e.queryCache,precision:2,suffix:"%","value-style":{color:"#cf1322"}},scopedSlots:e._u([{key:"prefix",fn:function(){},proxy:!0}])})],1)],1),n("a-col",{attrs:{span:8}},[n("a-card",[n("a-statistic",{staticClass:"demo-class",attrs:{title:"ThreadCacheHits",value:e.threadCache,precision:2,suffix:"%","value-style":{color:"#cf1322"}},scopedSlots:e._u([{key:"prefix",fn:function(){},proxy:!0}])})],1)],1)],1)],1)])],1),n("a-divider",{attrs:{orientation:"left"}},[e._v("Processlist")]),n("a-row",{attrs:{type:"flex"}},[n("a-col",{attrs:{flex:"100%"}},[n("a-table",{attrs:{columns:e.columns,"data-source":e.Processlist,rowKey:"Id"},scopedSlots:e._u([{key:"name",fn:function(t){return n("a",{},[e._v(e._s(t))])}},{key:"tags",fn:function(t){return n("span",{},e._l(t,(function(t){return n("a-tag",{key:t,attrs:{color:"loser"===t?"volcano":t.length>5?"geekblue":"green"}},[e._v(e._s(t.toUpperCase()))])})),1)}}])},[n("span",{attrs:{slot:"customTitle"},slot:"customTitle"},[n("a-icon",{attrs:{type:"smile-o"}}),e._v("Name ")],1)])],1)],1)],1)},a=[],o=n("bc3a"),s=n.n(o),i=n("1315"),c=[{title:"Id",dataIndex:"Id",key:"Id",slots:{title:"Id"},scopedSlots:{customRender:"Id"}},{title:"BaseName",dataIndex:"BaseName",key:"BaseName"},{title:"Command",dataIndex:"Command",key:"Command"},{title:"Host",key:"Host",dataIndex:"Host",scopedSlots:{customRender:"Host"}},{title:"User",key:"User",dataIndex:"User",scopedSlots:{customRender:"User"}},{title:"State",key:"State",dataIndex:"State",scopedSlots:{customRender:"State"}},{title:"Info",key:"Info",dataIndex:"Info",scopedSlots:{customRender:"Info"}},{title:"Time",key:"Time",dataIndex:"Time",scopedSlots:{customRender:"Time"}}],u={data:function(){return{DbInfo:{},Processlist:[],ProcesslistSize:0,TPS:0,QPS:0,keyBufferRead:"",keyBufferWrite:"",queryCache:"",threadCache:"",columns:c}},mounted:function(){this.getDBInfo(),this.getProcesslist(),this.getQPS(),this.getTPS(),this.getKeyBuffer(),this.getQueryCache(),this.getThreadCache()},methods:{formatDuring:function(e){var t=parseInt(e),n=0,r=0;t>60&&(n=parseInt(t/60),t=parseInt(t%60),n>60&&(r=parseInt(n/60),n=parseInt(n%60)));var a=parseInt(t)+"秒";return n>0&&(a=parseInt(n)+"分"+a),r>0&&(a=parseInt(r)+"小时"+a),a},getDBInfo:function(){var e=this;s.a.get(i["a"].URL+"base/dataBaseStatus",{params:{}}).then((function(t){e.DbInfo=t.data.data,e.DbInfo.Uptime=e.formatDuring(e.DbInfo.Uptime)}))},getProcesslist:function(){var e=this;s.a.get(i["a"].URL+"base/processlist",{params:{}}).then((function(t){e.Processlist=t.data.data,e.ProcesslistSize=e.Processlist.length}))},getQPS:function(){var e=this;s.a.get(i["a"].URL+"base/qps",{params:{}}).then((function(t){var n=String(t.data.data[0].Value);e.QPS=n.substring(0,5)}))},getTPS:function(){var e=this;s.a.get(i["a"].URL+"base/tps",{params:{}}).then((function(t){var n=String(t.data.data[0].Value);e.TPS=n.substring(0,5)}))},getKeyBuffer:function(){var e=this;s.a.get(i["a"].URL+"base/keyBuffer",{params:{}}).then((function(t){e.keyBufferRead=t.data.data[0].Value,e.keyBufferWrite=t.data.data[1].Value}))},getQueryCache:function(){var e=this;s.a.get(i["a"].URL+"base/queryCache",{params:{}}).then((function(t){e.queryCache=t.data.data.Value}))},getThreadCache:function(){var e=this;s.a.get(i["a"].URL+"base/threadCache",{params:{}}).then((function(t){e.threadCache=t.data.data.Value}))}}},f=u,l=n("2877"),d=Object(l["a"])(f,r,a,!1,null,null,null);t["default"]=d.exports}}]);
//# sourceMappingURL=about.2b9720af.js.map