var ReactValuSearch=function(e){function t(t){for(var n,r,a=t[0],i=t[1],c=0,d=[];c<a.length;c++)r=a[c],Object.prototype.hasOwnProperty.call(o,r)&&o[r]&&d.push(o[r][0]),o[r]=0;for(n in i)Object.prototype.hasOwnProperty.call(i,n)&&(e[n]=i[n]);for(l&&l(t);d.length;)d.shift()()}var n={},o={0:0};function r(t){if(n[t])return n[t].exports;var o=n[t]={i:t,l:!1,exports:{}};return e[t].call(o.exports,o,o.exports,r),o.l=!0,o.exports}r.e=function(e){var t=[],n=o[e];if(0!==n)if(n)t.push(n[2]);else{var a=new Promise((function(t,r){n=o[e]=[t,r]}));t.push(n[2]=a);var i,c=document.createElement("script");c.charset="utf-8",c.timeout=120,r.nc&&c.setAttribute("nonce",r.nc),c.src=function(e){return r.p+""+({}[e]||e)+".js"}(e);var l=new Error;i=function(t){c.onerror=c.onload=null,clearTimeout(d);var n=o[e];if(0!==n){if(n){var r=t&&("load"===t.type?"missing":t.type),a=t&&t.target&&t.target.src;l.message="Loading chunk "+e+" failed.\n("+r+": "+a+")",l.name="ChunkLoadError",l.type=r,l.request=a,n[1](l)}o[e]=void 0}};var d=setTimeout((function(){i({type:"timeout",target:c})}),12e4);c.onerror=c.onload=i,document.head.appendChild(c)}return Promise.all(t)},r.m=e,r.c=n,r.d=function(e,t,n){r.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},r.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},r.t=function(e,t){if(1&t&&(e=r(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(r.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var o in e)r.d(n,o,function(t){return e[t]}.bind(null,o));return n},r.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return r.d(t,"a",t),t},r.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},r.p="https://cdn.search.valu.pro/xamk/",r.oe=function(e){throw console.error(e),e};var a=window.webpackJsonpReactValuSearch=window.webpackJsonpReactValuSearch||[],i=a.push.bind(a);a.push=t,a=a.slice();for(var c=0;c<a.length;c++)t(a[c]);var l=i;return r(r.s=110)}({110:function(e,t,n){"use strict";n.r(t);var o=n(23),r=new o.LazyValuSearch({load:function(){return Promise.all([n.e(3),n.e(4)]).then(n.bind(null,57))}}),a=window.location.hostname;r.init((function(){var e=Object(o.select)(".banner .search-field",HTMLInputElement),t=Object(o.select)(".search-submit",HTMLButtonElement),n=document.querySelector(".wrap .search-field");"www.xamk.fi"===a&&n instanceof HTMLInputElement&&(Object(o.select)(".search-icon",HTMLDivElement).addEventListener("click",(function(){r.activate()})),n.addEventListener("focus",(function(){r.load()})));return t.addEventListener("click",(function(){r.activate()})),e.addEventListener("focus",(function(){r.load()})),function(t,o){"www.xamk.fi"===a&&n instanceof HTMLInputElement&&t.bindInputAsOpener(n),t.bindInputAsOpener(e),t.initModal({uiStrings:o.getUiStrings()})}}))},111:function(module,exports){function _defineProperties(e,t){for(var n=0;n<t.length;n++){var o=t[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(e,o.key,o)}}function _createClass(e,t,n){return t&&_defineProperties(e.prototype,t),n&&_defineProperties(e,n),e}function isSearchActive(e){return e=e||"vs","undefined"!=typeof window&&-1!==window.location.search.indexOf(e+"_q=")}function createEmitter(){var e=[];return{once:function(t){e.push(t)},emit:function(){var t=arguments;e.forEach((function(e){e.apply(void 0,[].slice.call(t))})),e=[]}}}var asserted=!1;function assertMainEntryNotLoaded(){"undefined"!=typeof window&&(asserted||("rvs-loaded"in window&&console.error("[RVS] Bad lazy loading setup! The '@valu/react-valu-search' entry is already imported. Ensure 'import \"@valu/react-valu-search\";' is not executed before the .load() call"),asserted=!0))}var polyfillEmitter=createEmitter(),polyfillStatus="waiting";function loadPolyfill(e,t){if(t||(t={}),polyfillEmitter.once(e),"done"===polyfillStatus)return polyfillEmitter.emit();if("loading"!==polyfillStatus){if(void 0!==t.isModern?t.isModern:isModern())return polyfillEmitter.emit();polyfillStatus="loading",console.log("[RVS] Loading Valu Search polyfills v"+version),loadScript(t.url||POLYFILL_URL,(function(){polyfillStatus="done",polyfillEmitter.emit()}))}}function hasFnProps(e,t){if(!e)return!1;for(var n=0;n<t.length;n++){var o=t[n];if(o&&"function"!=typeof e[o])return!1}return!0}function isModern(){try{return eval("async function test() {}"),!0}catch(e){}return hasFnProps([],["find","findIndex","includes"])&&hasFnProps(Array,["from"])&&hasFnProps(Object,["assign"])&&hasFnProps("",["startsWith","endsWith"])&&hasFnProps(window,["Promise","URL","URLSearchParams","fetch"])}var version="18.2.1",POLYFILL_URL="https://cdn.search.valu.pro/react-valu-search/v"+version+"/polyfill.js";function loadScript(e,t){var n=window.document.createElement("script");n.src=e,n.type="text/javascript",t&&n.addEventListener("load",t),window.document.getElementsByTagName("script")[0].parentNode.appendChild(n)}function select(e,t){var n=document.querySelector(e);if(n instanceof t)return n;throw new Error('[RVS] Bad selector "'+e+'" for  "'+t.name+'" got: "'+String(n))}function onDomContentLoaded(e){if(/complete|interactive|loaded/.test(document.readyState))e();else{document.addEventListener("DOMContentLoaded",(function t(){e(),document.removeEventListener("DOMContentLoaded",t)}),!1)}}var LazyValuSearch=function(){function e(e){var t=this;if(this.module=void 0,this.options=void 0,this.onLoad=createEmitter(),this.replaced=void 0,this.instanceId=void 0,this.isActive=function(){return isSearchActive(t.options.instanceId)},this.isLoaded=function(){return!!t.module},this.replace=function(e){localStorage.valuSearchLoaderReplaced=e,window.location.reload()},this.load=function(){t.replaced||(assertMainEntryNotLoaded(),loadPolyfill((function(){if(t.module)return t.onLoad.emit(t.module.default,t.module);t.options.load().then((function(e){if(!e.default||"function"!=typeof e.default.initModal)throw new Error("[RVS] ValuSearchLoader: Invalid default export from the lazy loaded module. Expected ValuSearch instance.");if(t.instanceId!==e.default.instanceId)throw new Error('[RVS] The lazy loaded RVS instanceId "'+e.default.instanceId+'"  does not match with loader instanceId "'+t.instanceId+'"');t.module=e,t.onLoad.emit(t.module.default,t.module)}),(function(e){console.error("[RVS] ValuSearchLoader: Failed to load",e)}))}),t.options.polyfill))},this.activate=function(){t.onLoad.once((function(e){e.activate()})),t.load()},this.init=function(e){t.replaced||onDomContentLoaded((function(){var n=e();n&&t.onLoad.once(n),isSearchActive(t.options.instanceId)&&t.load()}))},this.instanceId=e.instanceId||"vs",this.options=e,void 0!==typeof window){var n,o=(null==(n=window.valuSearchLoader)?void 0:n.replaced)||!1;if(window.valuSearchLoader=this,!o){var r=localStorage.valuSearchLoaderReplaced;r&&(this.replaced=!0,loadScript(r))}}}return _createClass(e,[{key:"trapClassName",get:function(){return"valu-search-focus-trap-"+this.instanceId}}]),e}();exports.LazyValuSearch=LazyValuSearch,exports.onDomContentLoaded=onDomContentLoaded,exports.select=select},23:function(e,t,n){e.exports=n(111)}});
//# sourceMappingURL=loader.js.map