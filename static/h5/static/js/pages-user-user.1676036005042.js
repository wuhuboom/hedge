(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-user-user"],{"1de5":function(t,e,a){"use strict";t.exports=function(t,e){return e||(e={}),t=t&&t.__esModule?t.default:t,"string"!==typeof t?t:(/^['"].*['"]$/.test(t)&&(t=t.slice(1,-1)),e.hash&&(t+=e.hash),/["'() \t\n]/.test(t)||e.needQuotes?'"'.concat(t.replace(/"/g,'\\"').replace(/\n/g,"\\n"),'"'):t)}},"332f":function(t,e,a){"use strict";a.r(e);var i=a("847b"),n=a("4e76");for(var r in n)"default"!==r&&function(t){a.d(e,t,(function(){return n[t]}))}(r);a("de20");var o,s=a("f0c5"),l=Object(s["a"])(n["default"],i["b"],i["c"],!1,null,"1fbc87f1",null,!1,i["a"],o);e["default"]=l.exports},"392e":function(t,e,a){var i=a("24fb");e=i(!1),e.push([t.i,".wrapListCell[data-v-3171b00c]{margin-top:10px;height:45px;background:#132736;border-radius:5px;display:flex;justify-content:space-between;align-items:center;padding:0 10px}.cell_arrowRight[data-v-3171b00c]{width:12px;height:17px}.cell_iconAndTitle[data-v-3171b00c]{display:flex;align-items:center}.cell_icon[data-v-3171b00c]{width:25px;height:25px}.cell_title[data-v-3171b00c]{margin-left:5px;font-size:14px;color:#fff}",""]),t.exports=e},"3d21":function(t,e,a){t.exports=a.p+"static/img/userimg.png"},"4b92":function(t,e,a){var i=a("9752");i.__esModule&&(i=i.default),"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var n=a("4f06").default;n("64e2e004",i,!0,{sourceMap:!1,shadowMode:!1})},"4bc6":function(t,e,a){t.exports=a.p+"static/img/back.png"},"4e0e":function(t,e,a){"use strict";var i;a.d(e,"b",(function(){return n})),a.d(e,"c",(function(){return r})),a.d(e,"a",(function(){return i}));var n=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("v-uni-view",{staticClass:"wrapListCell",style:t.itemCellStyle},[i("v-uni-view",{staticClass:"cell_iconAndTitle"},[i("v-uni-image",{staticClass:"cell_icon",attrs:{src:t.itemCell.imageIcon,mode:""}}),i("v-uni-view",{staticClass:"cell_title"},[t._v(t._s(t.itemCell.name))])],1),t.showImgBool?i("v-uni-image",{staticClass:"cell_arrowRight",attrs:{src:a("4fe3"),mode:""}}):t._e()],1)},r=[]},"4e76":function(t,e,a){"use strict";a.r(e);var i=a("e9eb"),n=a.n(i);for(var r in i)"default"!==r&&function(t){a.d(e,t,(function(){return i[t]}))}(r);e["default"]=n.a},"4fe3":function(t,e,a){t.exports=a.p+"static/img/arrow.svg"},"518c":function(t,e,a){var i=a("5e34");i.__esModule&&(i=i.default),"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var n=a("4f06").default;n("3599d2a2",i,!0,{sourceMap:!1,shadowMode:!1})},"56f0":function(t,e,a){"use strict";a.r(e);var i=a("9170"),n=a("de48");for(var r in n)"default"!==r&&function(t){a.d(e,t,(function(){return n[t]}))}(r);a("e697");var o,s=a("f0c5"),l=Object(s["a"])(n["default"],i["b"],i["c"],!1,null,"1d7f90d0",null,!1,i["a"],o);e["default"]=l.exports},"5e34":function(t,e,a){var i=a("24fb"),n=a("1de5"),r=a("4bc6"),o=a("ec86"),s=a("749d");e=i(!1);var l=n(r),c=n(o),u=n(s);e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */uni-page-body[data-v-1fbc87f1]{display:block;height:100%;position:relative}.page-layout[data-v-1fbc87f1]{height:100%}.page-layout .back[data-v-1fbc87f1]{position:relative;height:100vh;background:url('+l+") no-repeat 0 0;background-size:100%!important}.page-layout .back .scroll-Y[data-v-1fbc87f1]{height:100%}.page-layout .back .scroll-Y .user_box[data-v-1fbc87f1]{margin:35px 14px 0;display:flex;align-items:center}.page-layout .back .scroll-Y .user_box .user_Image[data-v-1fbc87f1]{box-sizing:border-box;border:.5px solid #88a1d0;width:69px;height:69px;border-radius:50%;margin-right:12px}.page-layout .back .scroll-Y .user_box .user_Image uni-image[data-v-1fbc87f1]{width:69px;height:69px;vertical-align:middle;display:block;white-space:nowrap}.page-layout .back .scroll-Y .user_box .userMain[data-v-1fbc87f1]{flex:1;overflow:hidden}.page-layout .back .scroll-Y .user_box .userMain .name[data-v-1fbc87f1]{display:flex;align-items:center}.page-layout .back .scroll-Y .user_box .userMain .name uni-text[data-v-1fbc87f1]{display:inline-block;max-width:115px;font-size:13px;color:#fff;font-weight:700}.page-layout .back .scroll-Y .user_box .userMain .desc[data-v-1fbc87f1]{margin-top:9px;font-size:18px;color:#ffd40a;font-weight:400;display:flex;align-items:center}.page-layout .back .scroll-Y .user_detail[data-v-1fbc87f1]{height:81px;margin:11px 15px 7px;background:url("+c+") no-repeat;background-size:contain;display:flex;align-items:center;padding:0 15px}.page-layout .back .scroll-Y .user_detail .left[data-v-1fbc87f1]{flex:1;overflow:hidden}.page-layout .back .scroll-Y .user_detail .left .col1[data-v-1fbc87f1]{display:flex;align-items:center;font-size:12px;color:#e5eeff}.page-layout .back .scroll-Y .user_detail .left .col1 uni-image[data-v-1fbc87f1]{width:20px;margin-left:2px;vertical-align:middle;margin-left:10px}.page-layout .back .scroll-Y .user_detail .left .col2[data-v-1fbc87f1]{font-size:31px;color:#e5eeff}.page-layout .back .scroll-Y .user_detail .left .u-line-1[data-v-1fbc87f1]{display:-webkit-box!important;overflow:hidden;text-overflow:ellipsis;word-break:break-all;-webkit-line-clamp:1;-webkit-box-orient:vertical!important}.page-layout .back .scroll-Y .user_detail .right[data-v-1fbc87f1]{color:hsla(0,0%,100%,.0352941);display:flex;align-items:center}.page-layout .back .scroll-Y .user_detail .right .btn uni-text[data-v-1fbc87f1]{white-space:nowrap;font-size:17px;color:#fff;margin:0 9px}.page-layout .back .scroll-Y .lists[data-v-1fbc87f1]{padding-bottom:80px}.page-layout .back .scroll-Y .lists .list[data-v-1fbc87f1]{box-sizing:border-box;padding:0 24px 0 30px;display:flex;align-items:center;height:46px;background:url("+u+") no-repeat 50%;background-size:contain}.page-layout .back .scroll-Y .lists .list .lf[data-v-1fbc87f1]{width:26px;vertical-align:middle;margin-right:22px}.page-layout .back .scroll-Y .lists .list .txt[data-v-1fbc87f1]{flex:1;color:#fff;font-size:15px}",""]),t.exports=e},"6e31":function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i={name:"wrapListCellView",props:{itemCellStyle:{type:String,default:"border-radius: 5px;"},itemCell:{type:Object,require:!0},showImgBool:{type:Boolean,default:!1}},mounted:function(){},data:function(){return{errCode:0,errMsg:""}}};e.default=i},"749d":function(t,e,a){t.exports=a.p+"static/img/myListBg.png"},"785c":function(t,e,a){"use strict";a("a9e3"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i=uni.getSystemInfoSync(),n={},r={name:"u-navbar",props:{height:{type:[String,Number],default:""},backIconColor:{type:String,default:"#606266"},backIconName:{type:String,default:"nav-back"},backIconSize:{type:[String,Number],default:"44"},backText:{type:String,default:""},backTextStyle:{type:Object,default:function(){return{color:"#606266"}}},title:{type:String,default:""},titleWidth:{type:[String,Number],default:"250"},titleColor:{type:String,default:"#606266"},titleBold:{type:Boolean,default:!1},titleSize:{type:[String,Number],default:32},isBack:{type:[Boolean,String],default:!0},background:{type:Object,default:function(){return{background:"#ffffff"}}},isFixed:{type:Boolean,default:!0},immersive:{type:Boolean,default:!1},borderBottom:{type:Boolean,default:!0},zIndex:{type:[String,Number],default:""},customBack:{type:Function,default:null}},data:function(){return{menuButtonInfo:n,statusBarHeight:i.statusBarHeight}},computed:{navbarInnerStyle:function(){var t={};return t.height=this.navbarHeight+"px",t},navbarStyle:function(){var t={};return t.zIndex=this.zIndex?this.zIndex:this.$u.zIndex.navbar,Object.assign(t,this.background),t},titleStyle:function(){var t={};return t.left=(i.windowWidth-uni.upx2px(this.titleWidth))/2+"px",t.right=(i.windowWidth-uni.upx2px(this.titleWidth))/2+"px",t.width=uni.upx2px(this.titleWidth)+"px",t},navbarHeight:function(){return this.height?this.height:44}},created:function(){},methods:{goBack:function(){"function"===typeof this.customBack?this.customBack.bind(this.$u.$parent.call(this))():uni.navigateBack()}}};e.default=r},"847b":function(t,e,a){"use strict";a.d(e,"b",(function(){return n})),a.d(e,"c",(function(){return r})),a.d(e,"a",(function(){return i}));var i={uNavbar:a("56f0").default,uIcon:a("79ab").default},n=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("v-uni-view",{staticClass:"page-layout"},[i("v-uni-view",{staticClass:"back"},[i("v-uni-scroll-view",{staticClass:"scroll-Y",attrs:{"scroll-top":t.scrollTop,"scroll-y":"true"}},[i("v-uni-view",{staticClass:"main"},[i("u-navbar",{attrs:{"is-back":!1,background:t.backgroundOBj,"border-bottom":!1,"z-index":"1200","back-icon-color":"#fff","title-color":"#fff",title:t.navTitle}}),i("v-uni-view",{staticClass:"user_box"},[i("v-uni-view",{staticClass:"user_Image"},[i("v-uni-image",{staticStyle:{height:"70px"},attrs:{src:a("3d21"),mode:""}})],1),i("v-uni-view",{staticClass:"userMain"},[i("v-uni-view",{staticClass:"name"},[i("v-uni-text",[t._v(t._s(t.username))])],1),i("v-uni-view",{staticClass:"desc"},[t._v(t._s(t.$t("home.balance"))+":"+t._s(t.cashPledge))])],1)],1),i("v-uni-view",{staticClass:"user_detail"},[i("v-uni-view",{staticClass:"left"},[i("v-uni-view",{staticClass:"col1 u-line-1"},[t._v(t._s(t.$t("user.page.money.text"))),i("v-uni-image",{staticStyle:{height:"16px"},attrs:{src:a("e21f"),mode:""},on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.initPlayInfoData.apply(void 0,arguments)}}})],1),i("v-uni-view",{staticClass:"col2 u-line-1"},[t._v(t._s(t.balance))])],1),i("v-uni-view",{staticClass:"right"},[i("v-uni-view",{staticClass:"btn"},[i("v-uni-text",{staticStyle:{color:"rgb(255, 212, 10)"}},[t._v(t._s(t.$t("home.recharge")))])],1),i("v-uni-view",{staticClass:"btn",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.withdrawBtnClick.apply(void 0,arguments)}}},[i("v-uni-text",[t._v(t._s(t.$t("home.withdraw")))])],1)],1)],1),i("v-uni-view",{staticClass:"lists"},t._l(t.dataUserList,(function(e,a){return i("v-uni-view",{key:e.id,staticClass:"list",on:{click:function(a){arguments[0]=a=t.$handleEvent(a),t.listCellClick(e)}}},[i("v-uni-image",{staticClass:"lf",style:"height:"+e.height+"px;",attrs:{src:e.imageIcon,mode:""}}),i("v-uni-view",{staticClass:"txt"},[t._v(t._s(e.name))]),i("u-icon",{attrs:{name:"arrow-right",color:"#fff"}})],1)})),1)],1)],1)],1)],1)},r=[]},9170:function(t,e,a){"use strict";a.d(e,"b",(function(){return n})),a.d(e,"c",(function(){return r})),a.d(e,"a",(function(){return i}));var i={uIcon:a("79ab").default},n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("v-uni-view",{},[a("v-uni-view",{staticClass:"u-navbar",class:{"u-navbar-fixed":t.isFixed,"u-border-bottom":t.borderBottom},style:[t.navbarStyle]},[a("v-uni-view",{staticClass:"u-status-bar",style:{height:t.statusBarHeight+"px"}}),a("v-uni-view",{staticClass:"u-navbar-inner",style:[t.navbarInnerStyle]},[t.isBack?a("v-uni-view",{staticClass:"u-back-wrap",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.goBack.apply(void 0,arguments)}}},[a("v-uni-view",{staticClass:"u-icon-wrap"},[a("u-icon",{attrs:{name:t.backIconName,color:t.backIconColor,size:t.backIconSize}})],1),t.backText?a("v-uni-view",{staticClass:"u-icon-wrap u-back-text u-line-1",style:[t.backTextStyle]},[t._v(t._s(t.backText))]):t._e()],1):t._e(),t.title?a("v-uni-view",{staticClass:"u-navbar-content-title",style:[t.titleStyle]},[a("v-uni-view",{staticClass:"u-title u-line-1",style:{color:t.titleColor,fontSize:t.titleSize+"rpx",fontWeight:t.titleBold?"bold":"normal"}},[t._v(t._s(t.title))])],1):t._e(),a("v-uni-view",{staticClass:"u-slot-content"},[t._t("default")],2),a("v-uni-view",{staticClass:"u-slot-right"},[t._t("right")],2)],1)],1),t.isFixed&&!t.immersive?a("v-uni-view",{staticClass:"u-navbar-placeholder",style:{width:"100%",height:Number(t.navbarHeight)+t.statusBarHeight+"px"}}):t._e()],1)},r=[]},9752:function(t,e,a){var i=a("24fb");e=i(!1),e.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */.u-navbar[data-v-1d7f90d0]{width:100%}.u-navbar-fixed[data-v-1d7f90d0]{position:fixed;left:0;right:0;top:0;z-index:991}.u-status-bar[data-v-1d7f90d0]{width:100%}.u-navbar-inner[data-v-1d7f90d0]{display:flex;flex-direction:row;justify-content:space-between;position:relative;align-items:center}.u-back-wrap[data-v-1d7f90d0]{display:flex;flex-direction:row;align-items:center;flex:1;flex-grow:0;padding:%?14?% %?14?% %?14?% %?24?%}.u-back-text[data-v-1d7f90d0]{padding-left:%?4?%;font-size:%?30?%}.u-navbar-content-title[data-v-1d7f90d0]{display:flex;flex-direction:row;align-items:center;justify-content:center;flex:1;position:absolute;left:0;right:0;height:%?60?%;text-align:center;flex-shrink:0}.u-navbar-centent-slot[data-v-1d7f90d0]{flex:1}.u-title[data-v-1d7f90d0]{line-height:%?60?%;font-size:%?32?%;flex:1}.u-navbar-right[data-v-1d7f90d0]{flex:1;display:flex;flex-direction:row;align-items:center;justify-content:flex-end}.u-slot-content[data-v-1d7f90d0]{flex:1;display:flex;flex-direction:row;align-items:center}',""]),t.exports=e},b05b:function(t,e,a){var i=a("392e");i.__esModule&&(i=i.default),"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var n=a("4f06").default;n("1fa629cb",i,!0,{sourceMap:!1,shadowMode:!1})},b15f:function(t,e,a){"use strict";var i=a("b05b"),n=a.n(i);n.a},b88f:function(t,e,a){"use strict";a.r(e);var i=a("4e0e"),n=a("c223");for(var r in n)"default"!==r&&function(t){a.d(e,t,(function(){return n[t]}))}(r);a("b15f");var o,s=a("f0c5"),l=Object(s["a"])(n["default"],i["b"],i["c"],!1,null,"3171b00c",null,!1,i["a"],o);e["default"]=l.exports},c223:function(t,e,a){"use strict";a.r(e);var i=a("6e31"),n=a.n(i);for(var r in i)"default"!==r&&function(t){a.d(e,t,(function(){return i[t]}))}(r);e["default"]=n.a},de20:function(t,e,a){"use strict";var i=a("518c"),n=a.n(i);n.a},de48:function(t,e,a){"use strict";a.r(e);var i=a("785c"),n=a.n(i);for(var r in i)"default"!==r&&function(t){a.d(e,t,(function(){return i[t]}))}(r);e["default"]=n.a},e21f:function(t,e,a){t.exports=a.p+"static/img/reload.png"},e697:function(t,e,a){"use strict";var i=a("4b92"),n=a.n(i);n.a},e9eb:function(t,e,a){"use strict";var i=a("4ea4");a("4160"),a("159b"),Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0,a("96cf");var n=i(a("1da1")),r=a("d8ab"),o=a("0a93"),s=i(a("b88f")),l={components:{wrapListCell:s.default},data:function(){return{navTitle:this.$t("tabar.mine"),backgroundOBj:{background:"transparent"},scrollTop:0,dataUserList:[],username:"",balance:0,cashPledge:0,vipNum:0,reateAmount:0,dataUserListCellStyle:"border-radius: 10px;",noticeNum:0,currDevOsNum:null,downUrl:"",realSymbolName:"",realSymbol:""}},filters:{saveTwoStr:function(t){return t.toFixed(2)}},mixins:[o.myMixin,o.withdrawMinxi,o.rechargeMinxi],onShow:function(){this.initLang(),this.initPlayInfoData()},methods:{initLang:function(){var t=uni.getStorageSync("run_lang");this.$i18n.locale=t,uni.setTabBarItem({index:0,text:this.$t("tabar.home")}),uni.setTabBarItem({index:1,text:this.$t("tabar.competition")}),uni.setTabBarItem({index:2,text:this.$t("tabar.sportsresults")}),uni.setTabBarItem({index:3,text:this.$t("tabar.mine")}),this.dataUserList=[{id:1,height:20,imageIcon:"../../static/images/info_l.png",name:this.$t("user.list.cell.my.info.text")},{id:2,height:24,imageIcon:"../../static/images/account_l.png",name:this.$t("user.list.cell.my.account.text")},{id:3,height:30,imageIcon:"../../static/images/order_l.png",name:this.$t("user.list.cell.recharge.record.text")},{id:4,height:30,imageIcon:"../../static/images/order_l.png",name:this.$t("user.list.cell.order.record.text")},{id:5,height:26,imageIcon:"../../static/images/money_l.png",name:this.$t("user.list.cell.fund.record.text")},{id:6,height:26,imageIcon:"../../static/images/money_l.png",name:this.$t("user.list.cell.order.status.text")},{id:7,height:26,imageIcon:"../../static/images/server_l.png",name:this.$t("user.list.cell.online.server.text")},{id:8,height:26,imageIcon:"../../static/images/agent_l.png",name:this.$t("user.list.cell.agent.top.text")},{id:9,height:30,imageIcon:"../../static/images/notice_l.png",name:this.$t("user.list.cell.sys.notice.text")},{id:10,height:26,imageIcon:"../../static/images/logout_l.png",name:this.$t("user.list.cell.account.logout.text")}]},getAppDownUrlData:function(){var t=this;return(0,n.default)(regeneratorRuntime.mark((function e(){var a,i,n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return a=uni.getSystemInfoSync(),a.platform&&"ios"===a.platform?t.currDevOsNum=1:t.currDevOsNum=0,e.next=4,appUrlReq();case 4:i=e.sent,200===i.code&&i.data&&0!==i.data.length&&(n=i.data,n.forEach((function(e,a){e.appType===t.currDevOsNum&&(t.downUrl=e.appUrl)})));case 6:case"end":return e.stop()}}),e)})))()},initNoticeData:function(){var t=this;return(0,n.default)(regeneratorRuntime.mark((function e(){var a,i,n,r;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return a={},i=uni.getStorageSync("run_lang"),a.lang="cn"==i?"zh":"en",e.next=5,noticeReq(a,t.localLoginToken);case 5:if(n=e.sent,200!==n.code){e.next=12;break}if(t.noticeNum=0,r=n.data,0!==r.length){e.next=11;break}return e.abrupt("return",!1);case 11:r.forEach((function(e,a){0===e.readStatus&&(t.noticeNum=t.noticeNum+1)}));case 12:case"end":return e.stop()}}),e)})))()},initPlayInfoData:function(){var t=this;return(0,n.default)(regeneratorRuntime.mark((function e(){var a;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return t.balance=0,t.cashPledge=0,e.next=4,(0,r.getMeReq)(t.localLoginToken);case 4:a=e.sent,200===a.code?(t.username=a.result.Username,t.balance=a.result.Balance,t.cashPledge=a.result.CashPledge):uni.showToast({icon:"none",title:a.msg});case 6:case"end":return e.stop()}}),e)})))()},orderCenterClick:function(){uni.navigateTo({url:"/pages/user/orderCenter",success:function(t){uni.setStorageSync("queryOrderCenterUrlFrom","user"),t.eventChannel.emit("eventOrderClick",{from:"user"})},animationType:"pop-in",animationDuration:200})},rwRecordClick:function(){uni.navigateTo({url:"/pages/user/withdrawRecord",animationType:"pop-in",animationDuration:200})},rechargeRecordClick:function(){uni.navigateTo({url:"/pages/user/RWrecord",success:function(t){},animationType:"pop-in",animationDuration:200})},toNoticeClick:function(){uni.navigateTo({url:"/pages/list/notice_list",success:function(t){uni.setStorageSync("queryNoticListUrlFrom","user"),t.eventChannel.emit("userNoticeClick",{from:"user"})},animationType:"pop-in",animationDuration:200})},rechargeBtnClick:function(){var t=this;uni.navigateTo({url:"/pages/recharge/index",success:function(e){t.save_recharge_from("user"),e.eventChannel.emit("rechargeEventClick",{from:"user"})},animationType:"pop-in",animationDuration:200})},withdrawBtnClick:function(){var t=this;uni.navigateTo({url:"/pages/withdraw/index",success:function(e){t.save_withdraw_from("user"),e.eventChannel.emit("withdrawEventClick",{from:"user"})},animationType:"pop-in",animationDuration:200})},wrapUserListCell_click:function(t,e){0===e?this.report_management_click():1===e?this.security_center_click():2===e?this.help_center_click():3===e?this.share_center_click():4===e&&this.down_app_click()},report_management_click:function(){uni.navigateTo({url:"/pages/user/report_management",animationType:"pop-in",animationDuration:200})},security_center_click:function(){uni.navigateTo({url:"/pages/user/securityCenter",animationType:"pop-in",animationDuration:200})},help_center_click:function(){uni.navigateTo({url:"/pages/user/help/help",animationType:"pop-in",animationDuration:200})},share_center_click:function(){uni.navigateTo({url:"/pages/refer/share",animationType:"pop-in",animationDuration:200})},down_app_click:function(){this.downUrl&&(window.location.href=this.downUrl)},listCellClick:function(t){var e=this;return(0,n.default)(regeneratorRuntime.mark((function a(){var i;return regeneratorRuntime.wrap((function(a){while(1)switch(a.prev=a.next){case 0:if(1!=t.id){a.next=3;break}a.next=38;break;case 3:if(2!=t.id){a.next=7;break}uni.navigateTo({url:"/pages/bank/bank",animationType:"pop-in",animationDuration:200}),a.next=38;break;case 7:if(3!=t.id){a.next=10;break}a.next=38;break;case 10:if(4!=t.id){a.next=14;break}uni.navigateTo({url:"/pages/user/orderRecord",animationType:"pop-in",animationDuration:200}),a.next=38;break;case 14:if(5!=t.id){a.next=18;break}uni.navigateTo({url:"/pages/user/accountChange",animationType:"pop-in",animationDuration:200}),a.next=38;break;case 18:if(6!=t.id){a.next=22;break}uni.navigateTo({url:"/pages/user/receOrderSitua",animationType:"pop-in",animationDuration:200}),a.next=38;break;case 22:if(7!=t.id){a.next=26;break}uni.navigateTo({url:"/pages/customer/index",animationType:"pop-in",animationDuration:200}),a.next=38;break;case 26:if(8!=t.id){a.next=29;break}a.next=38;break;case 29:if(9!=t.id){a.next=33;break}uni.navigateTo({url:"/pages/notice/notice_list",animationType:"pop-in",animationDuration:200}),a.next=38;break;case 33:if(10!=t.id){a.next=38;break}return a.next=36,(0,r.logoutReq)(e.localLoginToken);case 36:i=a.sent,200===i.code&&(e.logout(),uni.reLaunch({url:"/pages/login/login",animationType:"pop-in",animationDuration:200}));case 38:case"end":return a.stop()}}),a)})))()},logoutClick:function(){}}};e.default=l},ec86:function(t,e,a){t.exports=a.p+"static/img/my-bg2.png"}}]);