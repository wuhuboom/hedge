(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-order-order"],{"06c0":function(e,t,r){"use strict";var a=r("4ea4");Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0,r("96cf");var n=a(r("1da1")),i=r("d8ab"),o=r("0a93"),c={components:{},data:function(){return{navTitle:this.$t("tabar.competition"),username:"",balance:0,workStatus:!1,receiverList:[],backgroundOBj:{background:"transparent"},scrollTop:0,workText:"...",rec_b_style:"",confirmModelBool:!1,actualAmount:"",placeholderStyle:"font-size:14px;",confirmItem:{}}},filters:{},mixins:[o.myMixin],onShow:function(){this.initLang(),this.initPlayInfoData(),this.getReceiveListData()},methods:{initLang:function(){var e=uni.getStorageSync("run_lang");this.$i18n.locale=e,uni.setTabBarItem({index:0,text:this.$t("tabar.home")}),uni.setTabBarItem({index:1,text:this.$t("tabar.competition")}),uni.setTabBarItem({index:2,text:this.$t("tabar.sportsresults")}),uni.setTabBarItem({index:3,text:this.$t("tabar.mine")})},initPlayInfoData:function(){var e=this;return(0,n.default)(regeneratorRuntime.mark((function t(){var r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return e.balance=0,t.next=3,(0,i.getMeReq)(e.localLoginToken);case 3:r=t.sent,200===r.code&&(e.username=r.result.Username,e.balance=r.result.Balance,r.result&&2==r.result.Working?(e.workText=e.$t("order.receive.btn.stop.text"),e.workStatus=!0,e.rec_b_style="background: linear-gradient(120deg, #CE9FFC, #7367F0);color: #fff"):r.result&&1==r.result.Working&&(e.workText=e.$t("order.receive.btn.start.text"),e.workStatus=!1,e.rec_b_style="background: linear-gradient(120deg, #81FBB8, #28C76F);color: #fff"));case 5:case"end":return t.stop()}}),t)})))()},getReceiveListData:function(){var e=this;return(0,n.default)(regeneratorRuntime.mark((function t(){var r,a;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return r={},r.status=1,t.next=4,(0,i.getReceiveCollectionOrderReq)(r,e.localLoginToken);case 4:a=t.sent,200===a.code&&(e.receiverList=a.result);case 6:case"end":return t.stop()}}),t)})))()},receiveClick:function(){var e=this;return(0,n.default)(regeneratorRuntime.mark((function t(){var r,a;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return r={},r.work=e.workStatus?1:2,t.next=4,(0,i.imWorkingReq)(r,e.localLoginToken);case 4:a=t.sent,200===a.code&&("2"==a.msg?(e.workText=e.$t("order.receive.btn.stop.text"),e.workStatus=!0,e.rec_b_style="background: linear-gradient(120deg, #CE9FFC, #7367F0);color: #fff"):"1"==a.msg&&(e.workText=e.$t("order.receive.btn.start.text"),e.workStatus=!1,e.rec_b_style="background: linear-gradient(120deg, #81FBB8, #28C76F);color: #fff"));case 6:case"end":return t.stop()}}),t)})))()},cofirmReceiverMoneyClick:function(e){this.confirmItem=e,this.confirmModelBool=!0,this.actualAmount=e.ActualAmount},cancelModelClick:function(){this.confirmModelBool=!1},confirmModelClick:function(){var e=this;return(0,n.default)(regeneratorRuntime.mark((function t(){var r,a;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return r={},r.id=e.confirmItem.ID,r.amount=e.actualAmount,e.confirmModelBool=!1,t.next=6,(0,i.confirmTheOrderReq)(r,e.localLoginToken);case 6:a=t.sent,200===a.code&&(uni.showToast({icon:"none",title:a.msg}),e.getReceiveListData());case 8:case"end":return t.stop()}}),t)})))()}}};t.default=c},"1a26":function(e,t,r){e.exports=r.p+"static/img/bg2.png"},"2bd0":function(e,t,r){"use strict";r.r(t);var a=r("06c0"),n=r.n(a);for(var i in a)"default"!==i&&function(e){r.d(t,e,(function(){return a[e]}))}(i);t["default"]=n.a},"5dc9":function(e,t,r){"use strict";r.d(t,"b",(function(){return n})),r.d(t,"c",(function(){return i})),r.d(t,"a",(function(){return a}));var a={uNavbar:r("56f0").default,uPopup:r("b113").default},n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("v-uni-view",{staticClass:"page-layout"},[a("v-uni-view",{staticClass:"back"},[a("v-uni-scroll-view",{staticClass:"scroll-Y",attrs:{"scroll-top":e.scrollTop,"scroll-y":"true"}},[a("v-uni-view",{staticClass:"main"},[a("u-navbar",{attrs:{"is-back":!1,background:e.backgroundOBj,"border-bottom":!1,"z-index":"1200","back-icon-color":"#fff","title-color":"#fff",title:e.navTitle}}),a("v-uni-view",{staticClass:"orderContainer"},[a("v-uni-view",{staticClass:"receive_order"},[a("v-uni-view",{staticClass:"rec_action"},[a("v-uni-view",{staticClass:"rec_body",style:e.rec_b_style,on:{click:function(t){arguments[0]=t=e.$handleEvent(t),e.receiveClick.apply(void 0,arguments)}}},[e._v(e._s(e.workText))])],1)],1),0!==e.receiverList.length?a("v-uni-view",{staticClass:"leaderContainer"},[a("v-uni-view",{staticClass:"box"},[a("v-uni-view",{staticClass:"tit"},[e._v(e._s(e.$t("order.receive.list.head.tip.text")))]),e._l(e.receiverList,(function(t,r){return a("v-uni-view",{key:t.id,staticClass:"list_d"},[a("v-uni-view",{staticClass:"top"},[a("v-uni-view",{staticClass:"content"},[a("v-uni-view",{staticClass:"title"},[e._v(e._s(e.$t("order.receive.list.cell.ordernum.text")))]),a("v-uni-view",{staticClass:"con"},[a("v-uni-view",{staticClass:"real_c"},[e._v(e._s(t.OwnOrder))])],1)],1),a("v-uni-view",{staticClass:"btn",on:{click:function(r){arguments[0]=r=e.$handleEvent(r),e.cofirmReceiverMoneyClick(t)}}},[e._v(e._s(e.$t("order.receive.list.cell.oprate.confirm2.text")))])],1),a("v-uni-view",{staticClass:"bottom"},[a("v-uni-view",{staticClass:"title"},[e._v(e._s(e.$t("order.receive.list.cell.expiretime.text")))]),a("v-uni-view",{staticClass:"time"},[e._v(e._s(e._f("timestampStr2")(t.ExpireTime)))])],1)],1)}))],2)],1):e._e()],1)],1)],1)],1),a("u-popup",{attrs:{mode:"center","border-radius":"40","mask-close-able":!1},model:{value:e.confirmModelBool,callback:function(t){e.confirmModelBool=t},expression:"confirmModelBool"}},[a("v-uni-view",{staticClass:"classModel"},[a("v-uni-image",{staticClass:"headImg",staticStyle:{height:"176px"},attrs:{src:r("8a9e"),mode:""}}),a("v-uni-view",{staticClass:"center_txt"},[e._v(e._s(e.$t("order.receive.model.center.text")))]),a("v-uni-view",{staticClass:"container"},[a("v-uni-view",{staticClass:"left"},[e._v(e._s(e.$t("order.receive.model.form.input.left.text")))]),a("v-uni-view",{staticClass:"right"},[a("v-uni-input",{attrs:{type:"text",value:"",placeholder:e.$t("order.receive.model.form.input.place.text"),"placeholder-style":e.placeholderStyle},model:{value:e.actualAmount,callback:function(t){e.actualAmount=t},expression:"actualAmount"}})],1)],1),a("v-uni-view",{staticClass:"btns"},[a("v-uni-view",{staticClass:"cancel",on:{click:function(t){arguments[0]=t=e.$handleEvent(t),e.cancelModelClick.apply(void 0,arguments)}}},[e._v(e._s(e.$t("order.receive.list.cell.oprate.cancel.text")))]),a("v-uni-view",{staticClass:"confirm",on:{click:function(t){arguments[0]=t=e.$handleEvent(t),e.confirmModelClick.apply(void 0,arguments)}}},[e._v(e._s(e.$t("order.receive.list.cell.oprate.confirm.text")))])],1)],1)],1)],1)},i=[]},"69bd":function(e,t,r){var a=r("f792");a.__esModule&&(a=a.default),"string"===typeof a&&(a=[[e.i,a,""]]),a.locals&&(e.exports=a.locals);var n=r("4f06").default;n("7f759b50",a,!0,{sourceMap:!1,shadowMode:!1})},"8a9e":function(e,t,r){e.exports=r.p+"static/img/bg.png"},db6f:function(e,t,r){"use strict";var a=r("69bd"),n=r.n(a);n.a},ef3f:function(e,t,r){"use strict";r.r(t);var a=r("5dc9"),n=r("2bd0");for(var i in n)"default"!==i&&function(e){r.d(t,e,(function(){return n[e]}))}(i);r("db6f");var o,c=r("f0c5"),l=Object(c["a"])(n["default"],a["b"],a["c"],!1,null,"31cedfe5",null,!1,a["a"],o);t["default"]=l.exports},f792:function(e,t,r){var a=r("24fb"),n=r("1de5"),i=r("4bc6"),o=r("1a26");t=a(!1);var c=n(i),l=n(o);t.push([e.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */uni-page-body[data-v-31cedfe5]{display:block;height:100%;position:relative}.page-layout[data-v-31cedfe5]{height:100%}.page-layout .back[data-v-31cedfe5]{position:relative;height:100vh;background:url('+c+") no-repeat 0 0;background-size:100%!important}.page-layout .back .scroll-Y[data-v-31cedfe5]{height:100%}.page-layout .back .scroll-Y .orderContainer .receive_order[data-v-31cedfe5]{padding-top:30px}.page-layout .back .scroll-Y .orderContainer .receive_order .rec_action[data-v-31cedfe5]{display:flex;justify-content:center;flex-wrap:wrap}.page-layout .back .scroll-Y .orderContainer .receive_order .rec_action .rec_body[data-v-31cedfe5]{border-radius:50%;text-align:center;height:168px;width:168px;line-height:168px;color:#fff;font-weight:700;font-size:22.4px}.page-layout .back .scroll-Y .orderContainer .leaderContainer[data-v-31cedfe5]{padding:25px 13px 0;padding-bottom:80px}.page-layout .back .scroll-Y .orderContainer .leaderContainer .box[data-v-31cedfe5]{background-color:#fff;border-radius:12px;height:100%;padding:25px 12px 33px;margin-bottom:10px;box-shadow:0 4px 10px rgba(0,0,0,.3)}.page-layout .back .scroll-Y .orderContainer .leaderContainer .box .tit[data-v-31cedfe5]{font-size:21px;color:#5b9eff;line-height:1;position:relative;z-index:2;font-weight:700}.page-layout .back .scroll-Y .orderContainer .leaderContainer .box .list_d[data-v-31cedfe5]{background:url("+l+") no-repeat 0;background-size:contain;padding:15px 17px;margin-top:10px}.page-layout .back .scroll-Y .orderContainer .leaderContainer .box .list_d .top[data-v-31cedfe5]{background:linear-gradient(0deg,#f5f5ff,#fffefe);box-shadow:0 .5px 2px 0 rgba(50,116,224,.5);border-radius:10px;display:flex;align-items:center;padding:21px 15px 15px}.page-layout .back .scroll-Y .orderContainer .leaderContainer .box .list_d .top .content[data-v-31cedfe5]{flex:1}.page-layout .back .scroll-Y .orderContainer .leaderContainer .box .list_d .top .content .title[data-v-31cedfe5]{font-size:15px;line-height:1;color:#000}.page-layout .back .scroll-Y .orderContainer .leaderContainer .box .list_d .top .content .con[data-v-31cedfe5]{display:flex;align-items:center;margin-top:7px}.page-layout .back .scroll-Y .orderContainer .leaderContainer .box .list_d .top .content .con .real_c[data-v-31cedfe5]{display:-webkit-box!important;overflow:hidden;text-overflow:ellipsis;word-break:break-all;-webkit-line-clamp:2;-webkit-box-orient:vertical!important;font-size:9px;color:#ff4300}.page-layout .back .scroll-Y .orderContainer .leaderContainer .box .list_d .top .btn[data-v-31cedfe5]{background-color:#ff4e3f;border-radius:11px;color:#fff;font-size:11px;height:23;line-height:23px;padding-left:7px;padding-right:8px}.page-layout .back .scroll-Y .orderContainer .leaderContainer .box .list_d .bottom[data-v-31cedfe5]{display:flex;align-items:center;flex-direction:column;margin-top:15px;flex-flow:row wrap;justify-content:space-between}.page-layout .back .scroll-Y .orderContainer .leaderContainer .box .list_d .bottom .title[data-v-31cedfe5]{font-size:11px;color:#2e74ff}.page-layout .back .scroll-Y .orderContainer .leaderContainer .box .list_d .bottom .time[data-v-31cedfe5]{font-size:9px;color:#ff4300;margin-top:5px;margin-bottom:5px}.classModel[data-v-31cedfe5]{overflow:hidden;box-sizing:border-box;border-radius:20px;width:282px;min-height:332px;background-size:contain}.classModel .headImg[data-v-31cedfe5]{width:100%;margin-top:-60px;margin-bottom:22px}.classModel .center_txt[data-v-31cedfe5]{text-align:center;font-size:16px;color:#1c1c20;font-weight:700}.classModel .container[data-v-31cedfe5]{padding:0 10px;display:flex;flex-direction:row;align-items:center;justify-content:center;font-size:14px;color:#1c1c20;margin-top:30px}.classModel .container .left[data-v-31cedfe5]{height:20px;margin-right:5px}.classModel .container .right[data-v-31cedfe5]{width:40%}.classModel .btns[data-v-31cedfe5]{padding:0 20px;display:flex;flex-direction:row;align-items:center;justify-content:space-around;margin-top:30px}.classModel .btns .cancel[data-v-31cedfe5]{width:110px;height:40px;line-height:40px;background:#676767;box-shadow:0 5px 0 #243c47;border-radius:30px;text-align:center;font-size:14px;color:#fff}.classModel .btns .confirm[data-v-31cedfe5]{width:110px;height:40px;line-height:40px;background:#0686e7;box-shadow:0 5px 0 #016fd7;border-radius:30px;text-align:center;font-size:14px;color:#fff}",""]),e.exports=t}}]);