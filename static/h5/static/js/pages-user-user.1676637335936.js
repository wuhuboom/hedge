(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-user-user"],{"19a4":function(e,t,i){var a=i("e2ff");a.__esModule&&(a=a.default),"string"===typeof a&&(a=[[e.i,a,""]]),a.locals&&(e.exports=a.locals);var n=i("4f06").default;n("4375a11e",a,!0,{sourceMap:!1,shadowMode:!1})},"1de5":function(e,t,i){"use strict";e.exports=function(e,t){return t||(t={}),e=e&&e.__esModule?e.default:e,"string"!==typeof e?e:(/^['"].*['"]$/.test(e)&&(e=e.slice(1,-1)),t.hash&&(e+=t.hash),/["'() \t\n]/.test(e)||t.needQuotes?'"'.concat(e.replace(/"/g,'\\"').replace(/\n/g,"\\n"),'"'):e)}},"332f":function(e,t,i){"use strict";i.r(t);var a=i("4e4d"),n=i("4e76");for(var r in n)"default"!==r&&function(e){i.d(t,e,(function(){return n[e]}))}(r);i("fc55");var s,o=i("f0c5"),l=Object(o["a"])(n["default"],a["b"],a["c"],!1,null,"45b1310d",null,!1,a["a"],s);t["default"]=l.exports},"392e":function(e,t,i){var a=i("24fb");t=a(!1),t.push([e.i,".wrapListCell[data-v-3171b00c]{margin-top:10px;height:45px;background:#132736;border-radius:5px;display:flex;justify-content:space-between;align-items:center;padding:0 10px}.cell_arrowRight[data-v-3171b00c]{width:12px;height:17px}.cell_iconAndTitle[data-v-3171b00c]{display:flex;align-items:center}.cell_icon[data-v-3171b00c]{width:25px;height:25px}.cell_title[data-v-3171b00c]{margin-left:5px;font-size:14px;color:#fff}",""]),e.exports=t},"3d21":function(e,t,i){e.exports=i.p+"static/img/userimg.png"},"4b92":function(e,t,i){var a=i("9752");a.__esModule&&(a=a.default),"string"===typeof a&&(a=[[e.i,a,""]]),a.locals&&(e.exports=a.locals);var n=i("4f06").default;n("64e2e004",a,!0,{sourceMap:!1,shadowMode:!1})},"4bc6":function(e,t,i){e.exports=i.p+"static/img/back.png"},"4e0e":function(e,t,i){"use strict";var a;i.d(t,"b",(function(){return n})),i.d(t,"c",(function(){return r})),i.d(t,"a",(function(){return a}));var n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("v-uni-view",{staticClass:"wrapListCell",style:e.itemCellStyle},[a("v-uni-view",{staticClass:"cell_iconAndTitle"},[a("v-uni-image",{staticClass:"cell_icon",attrs:{src:e.itemCell.imageIcon,mode:""}}),a("v-uni-view",{staticClass:"cell_title"},[e._v(e._s(e.itemCell.name))])],1),e.showImgBool?a("v-uni-image",{staticClass:"cell_arrowRight",attrs:{src:i("4fe3"),mode:""}}):e._e()],1)},r=[]},"4e4d":function(e,t,i){"use strict";i.d(t,"b",(function(){return n})),i.d(t,"c",(function(){return r})),i.d(t,"a",(function(){return a}));var a={uNavbar:i("56f0").default,uIcon:i("79ab").default},n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("v-uni-view",{staticClass:"page-layout"},[a("v-uni-view",{staticClass:"back"},[a("v-uni-scroll-view",{staticClass:"scroll-Y",attrs:{"scroll-top":e.scrollTop,"scroll-y":"true"}},[a("v-uni-view",{staticClass:"main"},[a("u-navbar",{attrs:{"is-back":!1,background:e.backgroundOBj,"border-bottom":!1,"z-index":"1200","back-icon-color":"#fff","title-color":"#fff",title:e.navTitle}}),a("v-uni-view",{staticClass:"user_box"},[a("v-uni-view",{staticClass:"user_Image"},[a("v-uni-image",{staticStyle:{height:"70px"},attrs:{src:i("3d21"),mode:""}})],1),a("v-uni-view",{staticClass:"userMain"},[a("v-uni-view",{staticClass:"name"},[a("v-uni-text",[e._v(e._s(e.username))])],1),a("v-uni-view",{staticClass:"desc"},[e._v(e._s(e.$t("home.balance"))+":"+e._s(e.cashPledge))])],1)],1),a("v-uni-view",{staticClass:"user_detail"},[a("v-uni-view",{staticClass:"left"},[a("v-uni-view",{staticClass:"col1 u-line-1"},[e._v(e._s(e.$t("user.page.money.text"))),a("v-uni-image",{staticStyle:{height:"16px"},attrs:{src:i("e21f"),mode:""},on:{click:function(t){arguments[0]=t=e.$handleEvent(t),e.initPlayInfoData.apply(void 0,arguments)}}})],1),a("v-uni-view",{staticClass:"col2 u-line-1"},[e._v(e._s(e.balance))])],1),a("v-uni-view",{staticClass:"right"},[a("v-uni-view",{staticClass:"btn",on:{click:function(t){arguments[0]=t=e.$handleEvent(t),e.rechargeBtnClick.apply(void 0,arguments)}}},[a("v-uni-text",{staticStyle:{color:"rgb(255, 212, 10)"}},[e._v(e._s(e.$t("home.recharge")))])],1),a("v-uni-view",{staticClass:"btn",staticStyle:{"margin-left":"5px"},on:{click:function(t){arguments[0]=t=e.$handleEvent(t),e.withdrawBtnClick.apply(void 0,arguments)}}},[a("v-uni-text",[e._v(e._s(e.$t("home.withdraw")))])],1)],1)],1),a("v-uni-view",{staticClass:"lists"},e._l(e.dataUserList,(function(t,i){return a("v-uni-view",{key:t.id,staticClass:"list",on:{click:function(i){arguments[0]=i=e.$handleEvent(i),e.listCellClick(t)}}},[a("v-uni-image",{staticClass:"lf",style:"height:"+t.height+"px;",attrs:{src:t.imageIcon,mode:""}}),a("v-uni-view",{staticClass:"txt"},[e._v(e._s(t.name))]),a("u-icon",{attrs:{name:"arrow-right",color:"#fff"}})],1)})),1)],1)],1)],1)],1)},r=[]},"4e76":function(e,t,i){"use strict";i.r(t);var a=i("e9eb"),n=i.n(a);for(var r in a)"default"!==r&&function(e){i.d(t,e,(function(){return a[e]}))}(r);t["default"]=n.a},"4fe3":function(e,t,i){e.exports=i.p+"static/img/arrow.svg"},"56f0":function(e,t,i){"use strict";i.r(t);var a=i("9170"),n=i("de48");for(var r in n)"default"!==r&&function(e){i.d(t,e,(function(){return n[e]}))}(r);i("e697");var s,o=i("f0c5"),l=Object(o["a"])(n["default"],a["b"],a["c"],!1,null,"1d7f90d0",null,!1,a["a"],s);t["default"]=l.exports},"6e31":function(e,t,i){"use strict";Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var a={name:"wrapListCellView",props:{itemCellStyle:{type:String,default:"border-radius: 5px;"},itemCell:{type:Object,require:!0},showImgBool:{type:Boolean,default:!1}},mounted:function(){},data:function(){return{errCode:0,errMsg:""}}};t.default=a},"749d":function(e,t,i){e.exports=i.p+"static/img/myListBg.png"},"785c":function(e,t,i){"use strict";i("a9e3"),Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var a=uni.getSystemInfoSync(),n={},r={name:"u-navbar",props:{height:{type:[String,Number],default:""},backIconColor:{type:String,default:"#606266"},backIconName:{type:String,default:"nav-back"},backIconSize:{type:[String,Number],default:"44"},backText:{type:String,default:""},backTextStyle:{type:Object,default:function(){return{color:"#606266"}}},title:{type:String,default:""},titleWidth:{type:[String,Number],default:"250"},titleColor:{type:String,default:"#606266"},titleBold:{type:Boolean,default:!1},titleSize:{type:[String,Number],default:32},isBack:{type:[Boolean,String],default:!0},background:{type:Object,default:function(){return{background:"#ffffff"}}},isFixed:{type:Boolean,default:!0},immersive:{type:Boolean,default:!1},borderBottom:{type:Boolean,default:!0},zIndex:{type:[String,Number],default:""},customBack:{type:Function,default:null}},data:function(){return{menuButtonInfo:n,statusBarHeight:a.statusBarHeight}},computed:{navbarInnerStyle:function(){var e={};return e.height=this.navbarHeight+"px",e},navbarStyle:function(){var e={};return e.zIndex=this.zIndex?this.zIndex:this.$u.zIndex.navbar,Object.assign(e,this.background),e},titleStyle:function(){var e={};return e.left=(a.windowWidth-uni.upx2px(this.titleWidth))/2+"px",e.right=(a.windowWidth-uni.upx2px(this.titleWidth))/2+"px",e.width=uni.upx2px(this.titleWidth)+"px",e},navbarHeight:function(){return this.height?this.height:44}},created:function(){},methods:{goBack:function(){"function"===typeof this.customBack?this.customBack.bind(this.$u.$parent.call(this))():uni.navigateBack()}}};t.default=r},9170:function(e,t,i){"use strict";i.d(t,"b",(function(){return n})),i.d(t,"c",(function(){return r})),i.d(t,"a",(function(){return a}));var a={uIcon:i("79ab").default},n=function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("v-uni-view",{},[i("v-uni-view",{staticClass:"u-navbar",class:{"u-navbar-fixed":e.isFixed,"u-border-bottom":e.borderBottom},style:[e.navbarStyle]},[i("v-uni-view",{staticClass:"u-status-bar",style:{height:e.statusBarHeight+"px"}}),i("v-uni-view",{staticClass:"u-navbar-inner",style:[e.navbarInnerStyle]},[e.isBack?i("v-uni-view",{staticClass:"u-back-wrap",on:{click:function(t){arguments[0]=t=e.$handleEvent(t),e.goBack.apply(void 0,arguments)}}},[i("v-uni-view",{staticClass:"u-icon-wrap"},[i("u-icon",{attrs:{name:e.backIconName,color:e.backIconColor,size:e.backIconSize}})],1),e.backText?i("v-uni-view",{staticClass:"u-icon-wrap u-back-text u-line-1",style:[e.backTextStyle]},[e._v(e._s(e.backText))]):e._e()],1):e._e(),e.title?i("v-uni-view",{staticClass:"u-navbar-content-title",style:[e.titleStyle]},[i("v-uni-view",{staticClass:"u-title u-line-1",style:{color:e.titleColor,fontSize:e.titleSize+"rpx",fontWeight:e.titleBold?"bold":"normal"}},[e._v(e._s(e.title))])],1):e._e(),i("v-uni-view",{staticClass:"u-slot-content"},[e._t("default")],2),i("v-uni-view",{staticClass:"u-slot-right"},[e._t("right")],2)],1)],1),e.isFixed&&!e.immersive?i("v-uni-view",{staticClass:"u-navbar-placeholder",style:{width:"100%",height:Number(e.navbarHeight)+e.statusBarHeight+"px"}}):e._e()],1)},r=[]},9752:function(e,t,i){var a=i("24fb");t=a(!1),t.push([e.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */.u-navbar[data-v-1d7f90d0]{width:100%}.u-navbar-fixed[data-v-1d7f90d0]{position:fixed;left:0;right:0;top:0;z-index:991}.u-status-bar[data-v-1d7f90d0]{width:100%}.u-navbar-inner[data-v-1d7f90d0]{display:flex;flex-direction:row;justify-content:space-between;position:relative;align-items:center}.u-back-wrap[data-v-1d7f90d0]{display:flex;flex-direction:row;align-items:center;flex:1;flex-grow:0;padding:%?14?% %?14?% %?14?% %?24?%}.u-back-text[data-v-1d7f90d0]{padding-left:%?4?%;font-size:%?30?%}.u-navbar-content-title[data-v-1d7f90d0]{display:flex;flex-direction:row;align-items:center;justify-content:center;flex:1;position:absolute;left:0;right:0;height:%?60?%;text-align:center;flex-shrink:0}.u-navbar-centent-slot[data-v-1d7f90d0]{flex:1}.u-title[data-v-1d7f90d0]{line-height:%?60?%;font-size:%?32?%;flex:1}.u-navbar-right[data-v-1d7f90d0]{flex:1;display:flex;flex-direction:row;align-items:center;justify-content:flex-end}.u-slot-content[data-v-1d7f90d0]{flex:1;display:flex;flex-direction:row;align-items:center}',""]),e.exports=t},b05b:function(e,t,i){var a=i("392e");a.__esModule&&(a=a.default),"string"===typeof a&&(a=[[e.i,a,""]]),a.locals&&(e.exports=a.locals);var n=i("4f06").default;n("1fa629cb",a,!0,{sourceMap:!1,shadowMode:!1})},b15f:function(e,t,i){"use strict";var a=i("b05b"),n=i.n(a);n.a},b88f:function(e,t,i){"use strict";i.r(t);var a=i("4e0e"),n=i("c223");for(var r in n)"default"!==r&&function(e){i.d(t,e,(function(){return n[e]}))}(r);i("b15f");var s,o=i("f0c5"),l=Object(o["a"])(n["default"],a["b"],a["c"],!1,null,"3171b00c",null,!1,a["a"],s);t["default"]=l.exports},c223:function(e,t,i){"use strict";i.r(t);var a=i("6e31"),n=i.n(a);for(var r in a)"default"!==r&&function(e){i.d(t,e,(function(){return a[e]}))}(r);t["default"]=n.a},de48:function(e,t,i){"use strict";i.r(t);var a=i("785c"),n=i.n(a);for(var r in a)"default"!==r&&function(e){i.d(t,e,(function(){return a[e]}))}(r);t["default"]=n.a},e21f:function(e,t,i){e.exports=i.p+"static/img/reload.png"},e2ff:function(e,t,i){var a=i("24fb"),n=i("1de5"),r=i("4bc6"),s=i("ec86"),o=i("749d");t=a(!1);var l=n(r),c=n(s),u=n(o);t.push([e.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */uni-page-body[data-v-45b1310d]{display:block;height:100%;position:relative}.page-layout[data-v-45b1310d]{height:100%}.page-layout .back[data-v-45b1310d]{position:relative;height:100vh;background:url('+l+") no-repeat 0 0;background-size:100%!important}.page-layout .back .scroll-Y[data-v-45b1310d]{height:100%}.page-layout .back .scroll-Y .user_box[data-v-45b1310d]{margin:35px 14px 0;display:flex;align-items:center}.page-layout .back .scroll-Y .user_box .user_Image[data-v-45b1310d]{box-sizing:border-box;border:.5px solid #88a1d0;width:69px;height:69px;border-radius:50%;margin-right:12px}.page-layout .back .scroll-Y .user_box .user_Image uni-image[data-v-45b1310d]{width:69px;height:69px;vertical-align:middle;display:block;white-space:nowrap}.page-layout .back .scroll-Y .user_box .userMain[data-v-45b1310d]{flex:1;overflow:hidden}.page-layout .back .scroll-Y .user_box .userMain .name[data-v-45b1310d]{display:flex;align-items:center}.page-layout .back .scroll-Y .user_box .userMain .name uni-text[data-v-45b1310d]{display:inline-block;max-width:115px;font-size:13px;color:#fff;font-weight:700}.page-layout .back .scroll-Y .user_box .userMain .desc[data-v-45b1310d]{margin-top:9px;font-size:18px;color:#ffd40a;font-weight:400;display:flex;align-items:center}.page-layout .back .scroll-Y .user_detail[data-v-45b1310d]{height:81px;margin:11px 15px 7px;background:url("+c+") no-repeat;background-size:contain;display:flex;align-items:center;padding:0 15px}.page-layout .back .scroll-Y .user_detail .left[data-v-45b1310d]{flex:1;overflow:hidden}.page-layout .back .scroll-Y .user_detail .left .col1[data-v-45b1310d]{display:flex;align-items:center;font-size:12px;color:#e5eeff}.page-layout .back .scroll-Y .user_detail .left .col1 uni-image[data-v-45b1310d]{width:20px;margin-left:2px;vertical-align:middle;margin-left:10px}.page-layout .back .scroll-Y .user_detail .left .col2[data-v-45b1310d]{font-size:31px;color:#e5eeff}.page-layout .back .scroll-Y .user_detail .left .u-line-1[data-v-45b1310d]{display:-webkit-box!important;overflow:hidden;text-overflow:ellipsis;word-break:break-all;-webkit-line-clamp:1;-webkit-box-orient:vertical!important}.page-layout .back .scroll-Y .user_detail .right[data-v-45b1310d]{color:hsla(0,0%,100%,.0352941);display:flex;align-items:center}.page-layout .back .scroll-Y .user_detail .right .btn uni-text[data-v-45b1310d]{white-space:nowrap;font-size:17px;color:#fff;margin:0 9px}.page-layout .back .scroll-Y .lists[data-v-45b1310d]{padding-bottom:80px}.page-layout .back .scroll-Y .lists .list[data-v-45b1310d]{box-sizing:border-box;padding:0 24px 0 30px;display:flex;align-items:center;height:46px;background:url("+u+") no-repeat 50%;background-size:contain}.page-layout .back .scroll-Y .lists .list .lf[data-v-45b1310d]{width:26px;vertical-align:middle;margin-right:22px}.page-layout .back .scroll-Y .lists .list .txt[data-v-45b1310d]{flex:1;color:#fff;font-size:15px}",""]),e.exports=t},e697:function(e,t,i){"use strict";var a=i("4b92"),n=i.n(a);n.a},e9eb:function(e,t,i){"use strict";var a=i("4ea4");i("4160"),i("159b"),Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0,i("96cf");var n=a(i("1da1")),r=i("d8ab"),s=i("0a93"),o=a(i("b88f")),l={components:{wrapListCell:o.default},data:function(){return{navTitle:this.$t("tabar.mine"),backgroundOBj:{background:"transparent"},scrollTop:0,dataUserList:[],username:"",balance:0,cashPledge:0,vipNum:0,reateAmount:0,dataUserListCellStyle:"border-radius: 10px;",noticeNum:0,currDevOsNum:null,downUrl:"",realSymbolName:"",realSymbol:"",defaultServerAdd:"",leverTree:""}},filters:{saveTwoStr:function(e){return e.toFixed(2)}},mixins:[s.myMixin,s.withdrawMinxi,s.rechargeMinxi],onShow:function(){this.initLang(),this.initPlayInfoData(),this.getServerAddData()},methods:{initLang:function(){var e=uni.getStorageSync("run_lang");this.$i18n.locale=e,uni.setTabBarItem({index:0,text:this.$t("tabar.home")}),uni.setTabBarItem({index:1,text:this.$t("tabar.competition")}),uni.setTabBarItem({index:2,text:this.$t("tabar.sportsresults")}),uni.setTabBarItem({index:3,text:this.$t("tabar.mine")}),this.leverTree?this.dataUserList=[{id:1,height:20,imageIcon:"../../static/images/info_l.png",name:this.$t("user.list.cell.my.info.text")},{id:2,height:24,imageIcon:"../../static/images/account_l.png",name:this.$t("user.list.cell.my.account.text")},{id:3,height:30,imageIcon:"../../static/images/order_l.png",name:this.$t("user.list.cell.recharge.record.text")},{id:4,height:30,imageIcon:"../../static/images/order_l.png",name:this.$t("user.list.cell.order.record.text")},{id:5,height:26,imageIcon:"../../static/images/money_l.png",name:this.$t("user.list.cell.fund.record.text")},{id:6,height:26,imageIcon:"../../static/images/money_l.png",name:this.$t("user.list.cell.order.status.text")},{id:8,height:26,imageIcon:"../../static/images/agent_l.png",name:this.$t("user.list.cell.agent.top.text")},{id:9,height:30,imageIcon:"../../static/images/notice_l.png",name:this.$t("user.list.cell.sys.notice.text")},{id:10,height:26,imageIcon:"../../static/images/logout_l.png",name:this.$t("user.list.cell.account.logout.text")}]:this.dataUserList=[{id:1,height:20,imageIcon:"../../static/images/info_l.png",name:this.$t("user.list.cell.my.info.text")},{id:2,height:24,imageIcon:"../../static/images/account_l.png",name:this.$t("user.list.cell.my.account.text")},{id:3,height:30,imageIcon:"../../static/images/order_l.png",name:this.$t("user.list.cell.recharge.record.text")},{id:4,height:30,imageIcon:"../../static/images/order_l.png",name:this.$t("user.list.cell.order.record.text")},{id:5,height:26,imageIcon:"../../static/images/money_l.png",name:this.$t("user.list.cell.fund.record.text")},{id:6,height:26,imageIcon:"../../static/images/money_l.png",name:this.$t("user.list.cell.order.status.text")},{id:9,height:30,imageIcon:"../../static/images/notice_l.png",name:this.$t("user.list.cell.sys.notice.text")},{id:10,height:26,imageIcon:"../../static/images/logout_l.png",name:this.$t("user.list.cell.account.logout.text")}]},getServerAddData:function(){var e=this;return(0,n.default)(regeneratorRuntime.mark((function t(){var i;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,(0,r.servReq)(e.localLoginToken);case 2:i=t.sent,200===i.code&&(e.defaultServerAdd=i.result.CustomerServiceAddress);case 4:case"end":return t.stop()}}),t)})))()},getAppDownUrlData:function(){var e=this;return(0,n.default)(regeneratorRuntime.mark((function t(){var i,a,n;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return i=uni.getSystemInfoSync(),i.platform&&"ios"===i.platform?e.currDevOsNum=1:e.currDevOsNum=0,t.next=4,appUrlReq();case 4:a=t.sent,200===a.code&&a.data&&0!==a.data.length&&(n=a.data,n.forEach((function(t,i){t.appType===e.currDevOsNum&&(e.downUrl=t.appUrl)})));case 6:case"end":return t.stop()}}),t)})))()},initNoticeData:function(){var e=this;return(0,n.default)(regeneratorRuntime.mark((function t(){var i,a,n,r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return i={},a=uni.getStorageSync("run_lang"),i.lang="cn"==a?"zh":"en",t.next=5,noticeReq(i,e.localLoginToken);case 5:if(n=t.sent,200!==n.code){t.next=12;break}if(e.noticeNum=0,r=n.data,0!==r.length){t.next=11;break}return t.abrupt("return",!1);case 11:r.forEach((function(t,i){0===t.readStatus&&(e.noticeNum=e.noticeNum+1)}));case 12:case"end":return t.stop()}}),t)})))()},initPlayInfoData:function(){var e=this;return(0,n.default)(regeneratorRuntime.mark((function t(){var i;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return e.balance=0,e.cashPledge=0,t.next=4,(0,r.getMeReq)(e.localLoginToken);case 4:i=t.sent,200===i.code?(e.username=i.result.Username,e.balance=i.result.Balance,e.cashPledge=i.result.CashPledge,e.leverTree=i.result.LeverTree,e.initLang()):uni.showToast({icon:"none",title:i.msg});case 6:case"end":return t.stop()}}),t)})))()},orderCenterClick:function(){uni.navigateTo({url:"/pages/user/orderCenter",success:function(e){uni.setStorageSync("queryOrderCenterUrlFrom","user"),e.eventChannel.emit("eventOrderClick",{from:"user"})},animationType:"pop-in",animationDuration:200})},rwRecordClick:function(){uni.navigateTo({url:"/pages/user/withdrawRecord",animationType:"pop-in",animationDuration:200})},rechargeRecordClick:function(){uni.navigateTo({url:"/pages/user/RWrecord",success:function(e){},animationType:"pop-in",animationDuration:200})},toNoticeClick:function(){uni.navigateTo({url:"/pages/list/notice_list",success:function(e){uni.setStorageSync("queryNoticListUrlFrom","user"),e.eventChannel.emit("userNoticeClick",{from:"user"})},animationType:"pop-in",animationDuration:200})},rechargeBtnClick:function(){window.location.href=this.defaultServerAdd},withdrawBtnClick:function(){var e=this;uni.navigateTo({url:"/pages/withdraw/index",success:function(t){e.save_withdraw_from("user"),t.eventChannel.emit("withdrawEventClick",{from:"user"})},animationType:"pop-in",animationDuration:200})},wrapUserListCell_click:function(e,t){0===t?this.report_management_click():1===t?this.security_center_click():2===t?this.help_center_click():3===t?this.share_center_click():4===t&&this.down_app_click()},report_management_click:function(){uni.navigateTo({url:"/pages/user/report_management",animationType:"pop-in",animationDuration:200})},security_center_click:function(){uni.navigateTo({url:"/pages/user/securityCenter",animationType:"pop-in",animationDuration:200})},help_center_click:function(){uni.navigateTo({url:"/pages/user/help/help",animationType:"pop-in",animationDuration:200})},share_center_click:function(){uni.navigateTo({url:"/pages/refer/share",animationType:"pop-in",animationDuration:200})},down_app_click:function(){this.downUrl&&(window.location.href=this.downUrl)},listCellClick:function(e){var t=this;return(0,n.default)(regeneratorRuntime.mark((function i(){var a;return regeneratorRuntime.wrap((function(i){while(1)switch(i.prev=i.next){case 0:if(1!=e.id){i.next=4;break}uni.navigateTo({url:"/pages/user/info",animationType:"pop-in",animationDuration:200}),i.next=38;break;case 4:if(2!=e.id){i.next=8;break}uni.navigateTo({url:"/pages/bank/bank",animationType:"pop-in",animationDuration:200}),i.next=38;break;case 8:if(3!=e.id){i.next=11;break}i.next=38;break;case 11:if(4!=e.id){i.next=15;break}uni.navigateTo({url:"/pages/user/orderRecord",animationType:"pop-in",animationDuration:200}),i.next=38;break;case 15:if(5!=e.id){i.next=19;break}uni.navigateTo({url:"/pages/user/accountChange",animationType:"pop-in",animationDuration:200}),i.next=38;break;case 19:if(6!=e.id){i.next=23;break}uni.navigateTo({url:"/pages/user/receOrderSitua",animationType:"pop-in",animationDuration:200}),i.next=38;break;case 23:if(7!=e.id){i.next=26;break}i.next=38;break;case 26:if(8!=e.id){i.next=29;break}i.next=38;break;case 29:if(9!=e.id){i.next=33;break}uni.navigateTo({url:"/pages/notice/notice_list",animationType:"pop-in",animationDuration:200}),i.next=38;break;case 33:if(10!=e.id){i.next=38;break}return i.next=36,(0,r.logoutReq)(t.localLoginToken);case 36:a=i.sent,200===a.code&&(t.logout(),uni.reLaunch({url:"/pages/login/login",animationType:"pop-in",animationDuration:200}));case 38:case"end":return i.stop()}}),i)})))()},logoutClick:function(){}}};t.default=l},ec86:function(e,t,i){e.exports=i.p+"static/img/my-bg2.png"},fc55:function(e,t,i){"use strict";var a=i("19a4"),n=i.n(a);n.a}}]);