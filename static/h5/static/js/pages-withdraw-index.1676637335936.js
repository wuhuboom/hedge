(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-withdraw-index"],{"1a00":function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return i}));var i={uNavbar:n("56f0").default,uToast:n("3487").default},a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"main"},[n("u-navbar",{attrs:{"is-back":!0,"border-bottom":!1,"z-index":"1200","title-bold":!0,"back-icon-color":"#000000E6","title-color":"#000000E6",title:t.navTitle,"custom-back":t.goBackBtnClick}}),n("v-uni-view",{staticClass:"page_main"},[n("v-uni-view",{staticClass:"type_container"},[n("v-uni-view",{staticClass:"title"},[t._v(t._s(t.$t("withdraw.index.page.choose.type.text")))]),n("v-uni-view",{staticClass:"content"},t._l(t.typeList,(function(e,i){return n("v-uni-view",{key:i,class:t.chooseTypeIndex===e.id?"seleBtn active":"seleBtn",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.chooseTypeClick(e)}}},[t._v(t._s(e.name))])})),1)],1),n("v-uni-view",{staticClass:"money_container"},[n("v-uni-view",{staticClass:"title"},[t._v(t._s(t.$t("withdraw.index.page.choose.money.text")))]),n("v-uni-view",{staticClass:"content"},t._l(t.moneyList,(function(e,i){return n("v-uni-view",{key:i,class:t.chooseMoneyIndex===e.id?"seleBtn active":"seleBtn",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.chooseMoneyClick(e)}}},[n("v-uni-image",{staticStyle:{height:"16px"},attrs:{src:t.chooseMoneyIndex===e.id?"../../static/images/qian2.png":"../../static/images/qian.png",mode:""}}),t._v(t._s(e.value))],1)})),1)],1),n("v-uni-view",{staticClass:"form_container"},[n("v-uni-view",{staticClass:"txt"},[t._v(t._s(t.$t("withdraw.index.page.input.money.text")))]),n("v-uni-view",{staticClass:"input_m"},[n("v-uni-input",{attrs:{type:"text",value:"",placeholder:t.inputPlace,"placeholder-class":"place"},model:{value:t.money_input,callback:function(e){t.money_input=e},expression:"money_input"}})],1)],1),n("v-uni-view",{staticClass:"u-line"}),n("v-uni-view",{staticClass:"submitBtn",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.submitMoneyClick.apply(void 0,arguments)}}},[t._v(t._s(t.$t("withdraw.index.page.now.withdraw.btn.text")))])],1),n("u-toast",{ref:"uToast"})],1)},r=[]},"1de5":function(t,e,n){"use strict";t.exports=function(t,e){return e||(e={}),t=t&&t.__esModule?t.default:t,"string"!==typeof t?t:(/^['"].*['"]$/.test(t)&&(t=t.slice(1,-1)),e.hash&&(t+=e.hash),/["'() \t\n]/.test(t)||e.needQuotes?'"'.concat(t.replace(/"/g,'\\"').replace(/\n/g,"\\n"),'"'):t)}},"29a3":function(t,e,n){var i=n("ee33");i.__esModule&&(i=i.default),"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("5f91437f",i,!0,{sourceMap:!1,shadowMode:!1})},"2d10":function(t,e,n){"use strict";n.r(e);var i=n("1a00"),a=n("328d");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("e3e4");var o,s=n("f0c5"),c=Object(s["a"])(a["default"],i["b"],i["c"],!1,null,"b4df9dfa",null,!1,i["a"],o);e["default"]=c.exports},"328d":function(t,e,n){"use strict";n.r(e);var i=n("7207"),a=n.n(i);for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);e["default"]=a.a},"32de":function(t,e,n){"use strict";n.r(e);var i=n("6c82"),a=n.n(i);for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);e["default"]=a.a},3487:function(t,e,n){"use strict";n.r(e);var i=n("aa59"),a=n("32de");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("cc14");var o,s=n("f0c5"),c=Object(s["a"])(a["default"],i["b"],i["c"],!1,null,"36307caf",null,!1,i["a"],o);e["default"]=c.exports},"4b92":function(t,e,n){var i=n("9752");i.__esModule&&(i=i.default),"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("64e2e004",i,!0,{sourceMap:!1,shadowMode:!1})},"56f0":function(t,e,n){"use strict";n.r(e);var i=n("9170"),a=n("de48");for(var r in a)"default"!==r&&function(t){n.d(e,t,(function(){return a[t]}))}(r);n("e697");var o,s=n("f0c5"),c=Object(s["a"])(a["default"],i["b"],i["c"],!1,null,"1d7f90d0",null,!1,i["a"],o);e["default"]=c.exports},"6c82":function(t,e,n){"use strict";n("c975"),n("a9e3"),n("b64b"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={name:"u-toast",props:{zIndex:{type:[Number,String],default:""}},data:function(){return{isShow:!1,timer:null,config:{params:{},title:"",type:"",duration:2e3,isTab:!1,url:"",icon:!0,position:"center",callback:null,back:!1},tmpConfig:{}}},computed:{iconName:function(){if(["error","warning","success","info"].indexOf(this.tmpConfig.type)>=0&&this.tmpConfig.icon){var t=this.$u.type2icon(this.tmpConfig.type);return t}},uZIndex:function(){return this.isShow?this.zIndex?this.zIndex:this.$u.zIndex.toast:"999999"}},methods:{show:function(t){var e=this;this.tmpConfig=this.$u.deepMerge(this.config,t),this.timer&&(clearTimeout(this.timer),this.timer=null),this.isShow=!0,this.timer=setTimeout((function(){e.isShow=!1,clearTimeout(e.timer),e.timer=null,"function"===typeof e.tmpConfig.callback&&e.tmpConfig.callback(),e.timeEnd()}),this.tmpConfig.duration)},hide:function(){this.isShow=!1,this.timer&&(clearTimeout(this.timer),this.timer=null)},timeEnd:function(){if(this.tmpConfig.url){if("/"!=this.tmpConfig.url[0]&&(this.tmpConfig.url="/"+this.tmpConfig.url),Object.keys(this.tmpConfig.params).length){var t="";/.*\/.*\?.*=.*/.test(this.tmpConfig.url)?(t=this.$u.queryParams(this.tmpConfig.params,!1),this.tmpConfig.url=this.tmpConfig.url+"&"+t):(t=this.$u.queryParams(this.tmpConfig.params),this.tmpConfig.url+=t)}this.tmpConfig.isTab?uni.switchTab({url:this.tmpConfig.url}):uni.navigateTo({url:this.tmpConfig.url})}else this.tmpConfig.back&&this.$u.route({type:"back"})}}};e.default=i},7207:function(t,e,n){"use strict";var i=n("4ea4");Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0,n("96cf");var a=i(n("1da1")),r=n("d8ab"),o=n("0a93"),s={data:function(){return{navTitle:"",backgroundOBj:{background:"transparent"},moneyList:[{id:1,value:100},{id:2,value:300},{id:3,value:500},{id:4,value:1e3},{id:5,value:2e3},{id:6,value:3e3},{id:7,value:5e3},{id:8,value:1e4},{id:9,value:2e4},{id:10,value:5e4},{id:11,value:1e5}],typeList:[{id:1,name:"UPI"},{id:2,name:"USDT"}],chooseMoneyIndex:1,chooseTypeIndex:1,inputPlace:this.$t("withdraw.index.page.inpuit.money.place.text"),money_input:"100",scrollTop:0,urlFromStr:null}},mixins:[o.myMixin],onShow:function(){var t=this;this.initLang(),this.initData();var e=this.getOpenerEventChannel();e.on("withdrawEventClick",function(){var e=(0,a.default)(regeneratorRuntime.mark((function e(n){return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:t.urlFromStr=n.from;case 1:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}())},methods:{initLang:function(){var t=uni.getStorageSync("run_lang");this.$i18n.locale=t},initData:function(){this.navTitle=this.$t("home.withdraw")},goBackBtnClick:function(){var t=uni.getStorageSync("run_queryWithdrawUrlFrom")||"other";this.run_withdrawState_Bool||(this.urlFromStr=t),"user"===this.urlFromStr?uni.switchTab({url:"/pages/user/user",animationType:"pop-in",animationDuration:200}):"home"===this.urlFromStr&&uni.switchTab({url:"/pages/home/index",animationType:"pop-in",animationDuration:200})},chooseTypeClick:function(t){this.chooseTypeIndex=t.id},chooseMoneyClick:function(t){this.chooseMoneyIndex=t.id,this.money_input=t.value},submitMoneyClick:function(){var t=this;return(0,a.default)(regeneratorRuntime.mark((function e(){var n,i;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return n={},n.kinds=t.chooseTypeIndex+1,n.amount=t.money_input,e.next=5,(0,r.withdrawalReq)(n,t.localLoginToken);case 5:i=e.sent,200===i.code&&uni.showToast({title:i.msg,icon:"none",duration:2e3});case 7:case"end":return e.stop()}}),e)})))()}}};e.default=s},"785c":function(t,e,n){"use strict";n("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i=uni.getSystemInfoSync(),a={},r={name:"u-navbar",props:{height:{type:[String,Number],default:""},backIconColor:{type:String,default:"#606266"},backIconName:{type:String,default:"nav-back"},backIconSize:{type:[String,Number],default:"44"},backText:{type:String,default:""},backTextStyle:{type:Object,default:function(){return{color:"#606266"}}},title:{type:String,default:""},titleWidth:{type:[String,Number],default:"250"},titleColor:{type:String,default:"#606266"},titleBold:{type:Boolean,default:!1},titleSize:{type:[String,Number],default:32},isBack:{type:[Boolean,String],default:!0},background:{type:Object,default:function(){return{background:"#ffffff"}}},isFixed:{type:Boolean,default:!0},immersive:{type:Boolean,default:!1},borderBottom:{type:Boolean,default:!0},zIndex:{type:[String,Number],default:""},customBack:{type:Function,default:null}},data:function(){return{menuButtonInfo:a,statusBarHeight:i.statusBarHeight}},computed:{navbarInnerStyle:function(){var t={};return t.height=this.navbarHeight+"px",t},navbarStyle:function(){var t={};return t.zIndex=this.zIndex?this.zIndex:this.$u.zIndex.navbar,Object.assign(t,this.background),t},titleStyle:function(){var t={};return t.left=(i.windowWidth-uni.upx2px(this.titleWidth))/2+"px",t.right=(i.windowWidth-uni.upx2px(this.titleWidth))/2+"px",t.width=uni.upx2px(this.titleWidth)+"px",t},navbarHeight:function(){return this.height?this.height:44}},created:function(){},methods:{goBack:function(){"function"===typeof this.customBack?this.customBack.bind(this.$u.$parent.call(this))():uni.navigateBack()}}};e.default=r},"7d54":function(t,e,n){t.exports=n.p+"static/img/btn.png"},9170:function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return i}));var i={uIcon:n("79ab").default},a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{},[n("v-uni-view",{staticClass:"u-navbar",class:{"u-navbar-fixed":t.isFixed,"u-border-bottom":t.borderBottom},style:[t.navbarStyle]},[n("v-uni-view",{staticClass:"u-status-bar",style:{height:t.statusBarHeight+"px"}}),n("v-uni-view",{staticClass:"u-navbar-inner",style:[t.navbarInnerStyle]},[t.isBack?n("v-uni-view",{staticClass:"u-back-wrap",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.goBack.apply(void 0,arguments)}}},[n("v-uni-view",{staticClass:"u-icon-wrap"},[n("u-icon",{attrs:{name:t.backIconName,color:t.backIconColor,size:t.backIconSize}})],1),t.backText?n("v-uni-view",{staticClass:"u-icon-wrap u-back-text u-line-1",style:[t.backTextStyle]},[t._v(t._s(t.backText))]):t._e()],1):t._e(),t.title?n("v-uni-view",{staticClass:"u-navbar-content-title",style:[t.titleStyle]},[n("v-uni-view",{staticClass:"u-title u-line-1",style:{color:t.titleColor,fontSize:t.titleSize+"rpx",fontWeight:t.titleBold?"bold":"normal"}},[t._v(t._s(t.title))])],1):t._e(),n("v-uni-view",{staticClass:"u-slot-content"},[t._t("default")],2),n("v-uni-view",{staticClass:"u-slot-right"},[t._t("right")],2)],1)],1),t.isFixed&&!t.immersive?n("v-uni-view",{staticClass:"u-navbar-placeholder",style:{width:"100%",height:Number(t.navbarHeight)+t.statusBarHeight+"px"}}):t._e()],1)},r=[]},9752:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */.u-navbar[data-v-1d7f90d0]{width:100%}.u-navbar-fixed[data-v-1d7f90d0]{position:fixed;left:0;right:0;top:0;z-index:991}.u-status-bar[data-v-1d7f90d0]{width:100%}.u-navbar-inner[data-v-1d7f90d0]{display:flex;flex-direction:row;justify-content:space-between;position:relative;align-items:center}.u-back-wrap[data-v-1d7f90d0]{display:flex;flex-direction:row;align-items:center;flex:1;flex-grow:0;padding:%?14?% %?14?% %?14?% %?24?%}.u-back-text[data-v-1d7f90d0]{padding-left:%?4?%;font-size:%?30?%}.u-navbar-content-title[data-v-1d7f90d0]{display:flex;flex-direction:row;align-items:center;justify-content:center;flex:1;position:absolute;left:0;right:0;height:%?60?%;text-align:center;flex-shrink:0}.u-navbar-centent-slot[data-v-1d7f90d0]{flex:1}.u-title[data-v-1d7f90d0]{line-height:%?60?%;font-size:%?32?%;flex:1}.u-navbar-right[data-v-1d7f90d0]{flex:1;display:flex;flex-direction:row;align-items:center;justify-content:flex-end}.u-slot-content[data-v-1d7f90d0]{flex:1;display:flex;flex-direction:row;align-items:center}',""]),t.exports=e},aa59:function(t,e,n){"use strict";n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return r})),n.d(e,"a",(function(){return i}));var i={uIcon:n("79ab").default},a=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("v-uni-view",{staticClass:"u-toast",class:[t.isShow?"u-show":"","u-type-"+t.tmpConfig.type,"u-position-"+t.tmpConfig.position],style:{zIndex:t.uZIndex}},[n("v-uni-view",{staticClass:"u-icon-wrap"},[t.tmpConfig.icon?n("u-icon",{staticClass:"u-icon",attrs:{name:t.iconName,size:30,color:t.tmpConfig.type}}):t._e()],1),n("v-uni-text",{staticClass:"u-text"},[t._v(t._s(t.tmpConfig.title))])],1)},r=[]},b3d0:function(t,e,n){var i=n("24fb"),a=n("1de5"),r=n("7d54");e=i(!1);var o=a(r);e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */uni-page-body[data-v-b4df9dfa]{height:100%;background-color:#fff}.main[data-v-b4df9dfa]{display:flex;flex-direction:column;height:100%;overflow:hidden}.main .page_main[data-v-b4df9dfa]{padding:0 25px;margin-top:17px}.main .page_main .type_container[data-v-b4df9dfa]{display:flex;flex-direction:column}.main .page_main .type_container .content[data-v-b4df9dfa]{margin-top:15px;flex:1;display:flex;flex-direction:row;align-items:center;flex-wrap:wrap}.main .page_main .type_container .content .seleBtn[data-v-b4df9dfa]{margin-right:15px;margin-bottom:12px;min-width:70px;box-sizing:border-box;padding-right:5px;height:27px;border-radius:5px;background-color:#f6f6f6;color:#666;font-size:14px;font-weight:400;display:flex;align-items:center;justify-content:center}.main .page_main .type_container .content .active[data-v-b4df9dfa]{background:#34d155;box-shadow:0 4px 4px 0 rgba(55,210,84,.4);color:#fff}.main .page_main .money_container[data-v-b4df9dfa]{display:flex;flex-direction:column;margin-top:20px}.main .page_main .money_container .content[data-v-b4df9dfa]{margin-top:15px;flex:1;display:flex;flex-direction:row;align-items:center;flex-wrap:wrap}.main .page_main .money_container .content .seleBtn[data-v-b4df9dfa]{margin-right:15px;margin-bottom:12px;min-width:70px;box-sizing:border-box;padding-right:5px;height:27px;border-radius:5px;background-color:#f6f6f6;color:#666;font-size:14px;font-weight:400;display:flex;align-items:center}.main .page_main .money_container .content .seleBtn uni-image[data-v-b4df9dfa]{margin-left:7px;margin-right:9px;width:16px;vertical-align:middle}.main .page_main .money_container .content .seleBtn[data-v-b4df9dfa]:nth-child(4n+4){margin-right:0}.main .page_main .money_container .content .active[data-v-b4df9dfa]{background:#34d155;box-shadow:0 4px 4px 0 rgba(55,210,84,.4);color:#fff;max-resolution:9px}.main .page_main .form_container[data-v-b4df9dfa]{display:flex;flex-direction:row;padding:10px 0;margin-top:20px}.main .page_main .form_container .txt[data-v-b4df9dfa]{font-size:14px;font-weight:400;color:#2d2d2d;display:inline-block;margin-right:19px}.main .page_main .form_container .input_m[data-v-b4df9dfa] .place{color:#c0c4cc;font-size:13px;font-weight:400}.main .page_main .u-line[data-v-b4df9dfa]{margin:0;border-bottom:1px solid #d6d7d9;width:100%;-webkit-transform:scaleY(.5);transform:scaleY(.5);border-top-color:#d6d7d9;border-right-color:#d6d7d9;border-left-color:#d6d7d9}.main .page_main .submitBtn[data-v-b4df9dfa]{width:280px;height:43px;line-height:43px;background:url('+o+") no-repeat 50%;background-size:contain;margin:55px auto 20px;font-size:16px;color:#fff;text-align:center}body.?%PAGE?%[data-v-b4df9dfa]{background-color:#fff}",""]),t.exports=e},ba0f:function(t,e,n){var i=n("b3d0");i.__esModule&&(i=i.default),"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=n("4f06").default;a("768921ae",i,!0,{sourceMap:!1,shadowMode:!1})},cc14:function(t,e,n){"use strict";var i=n("29a3"),a=n.n(i);a.a},de48:function(t,e,n){"use strict";n.r(e);var i=n("785c"),a=n.n(i);for(var r in i)"default"!==r&&function(t){n.d(e,t,(function(){return i[t]}))}(r);e["default"]=a.a},e3e4:function(t,e,n){"use strict";var i=n("ba0f"),a=n.n(i);a.a},e697:function(t,e,n){"use strict";var i=n("4b92"),a=n.n(i);a.a},ee33:function(t,e,n){var i=n("24fb");e=i(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */.u-toast[data-v-36307caf]{position:fixed;z-index:-1;transition:opacity .3s;text-align:center;color:#fff;border-radius:%?8?%;background:#585858;display:flex;flex-direction:row;align-items:center;justify-content:center;font-size:%?28?%;opacity:0;pointer-events:none;padding:%?18?% %?40?%}.u-toast.u-show[data-v-36307caf]{opacity:1}.u-icon[data-v-36307caf]{margin-right:%?10?%;display:flex;flex-direction:row;align-items:center;line-height:normal}.u-position-center[data-v-36307caf]{left:50%;top:50%;-webkit-transform:translate(-50%,-50%);transform:translate(-50%,-50%);max-width:70%}.u-position-top[data-v-36307caf]{left:50%;top:20%;-webkit-transform:translate(-50%,-50%);transform:translate(-50%,-50%)}.u-position-bottom[data-v-36307caf]{left:50%;bottom:20%;-webkit-transform:translate(-50%,-50%);transform:translate(-50%,-50%)}.u-type-primary[data-v-36307caf]{color:#2979ff;background-color:#ecf5ff;border:1px solid #d7eafe}.u-type-success[data-v-36307caf]{color:#19be6b;background-color:#dbf1e1;border:1px solid #bef5c8}.u-type-error[data-v-36307caf]{color:#fa3534;background-color:#fef0f0;border:1px solid #fde2e2}.u-type-warning[data-v-36307caf]{color:#f90;background-color:#fdf6ec;border:1px solid #faecd8}.u-type-info[data-v-36307caf]{color:#909399;background-color:#f4f4f5;border:1px solid #ebeef5}.u-type-default[data-v-36307caf]{color:#fff;background-color:#585858}',""]),t.exports=e}}]);