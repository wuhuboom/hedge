(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-bank-bank"],{"2ae2":function(t,n,e){var i=e("24fb");n=i(!1),n.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */uni-page-body[data-v-c7604558]{height:100%;background-color:#fff}.main[data-v-c7604558]{display:flex;flex-direction:column;height:100%;overflow:hidden}.main .content_cell[data-v-c7604558]{font-size:14px;color:#303133;margin-top:2px}.main .u-line[data-v-c7604558]{margin:0;border-bottom:1px solid #d6d7d9;width:100%;-webkit-transform:scaleY(.5);transform:scaleY(.5);border-top-color:#d6d7d9;border-right-color:#d6d7d9;border-left-color:#d6d7d9}body.?%PAGE?%[data-v-c7604558]{background-color:#fff}',""]),t.exports=n},"58af":function(t,n,e){"use strict";var i=e("dd55"),a=e.n(i);a.a},9371:function(t,n,e){"use strict";e.d(n,"b",(function(){return a})),e.d(n,"c",(function(){return r})),e.d(n,"a",(function(){return i}));var i={uNavbar:e("56f0").default,uCellGroup:e("f23f").default,uCellItem:e("1bae").default,uToast:e("3487").default},a=function(){var t=this,n=t.$createElement,i=t._self._c||n;return i("v-uni-view",{staticClass:"main"},[i("u-navbar",{attrs:{"is-back":!0,"border-bottom":!1,"z-index":"1200","title-bold":!0,"back-icon-color":"#000000E6","title-color":"#000000E6",title:t.navTitle,"custom-back":t.goBackBtnClick}}),i("u-cell-group",{attrs:{border:!0}},[i("u-cell-item",{attrs:{center:!0,arrow:!1,"border-bottom":!0,index:"1"},on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.uCellEventClick(1)}}},[i("v-uni-image",{staticStyle:{"vertical-align":"middle",width:"19px",height:"18px","margin-right":"4px"},attrs:{slot:"icon",src:e("73ea"),mode:""},slot:"icon"}),i("v-uni-view",{staticClass:"content_cell",attrs:{slot:"title"},slot:"title"},[t._v(t._s(t.$t("account.page.cell.bindupi1.text")))])],1),i("v-uni-view",{staticClass:"u-line"}),i("u-cell-item",{attrs:{center:!0,arrow:!1,"border-bottom":!0,index:"2"},on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.uCellEventClick(2)}}},[i("v-uni-image",{staticStyle:{"vertical-align":"middle",width:"19px",height:"18px","margin-right":"4px"},attrs:{slot:"icon",src:e("73ea"),mode:""},slot:"icon"}),i("v-uni-view",{staticClass:"content_cell",attrs:{slot:"title"},slot:"title"},[t._v(t._s(t.$t("account.page.cell.bindupi2.text")))])],1),i("v-uni-view",{staticClass:"u-line"}),i("u-cell-item",{attrs:{center:!0,arrow:!1,"border-bottom":!0,index:"3"},on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.uCellEventClick(3)}}},[i("v-uni-image",{staticStyle:{"vertical-align":"middle",width:"19px",height:"18px","margin-right":"4px"},attrs:{slot:"icon",src:e("73ea"),mode:""},slot:"icon"}),i("v-uni-view",{staticClass:"content_cell",attrs:{slot:"title"},slot:"title"},[t._v(t._s(t.$t("account.page.cell.bindupi3.text")))])],1),i("v-uni-view",{staticClass:"u-line"})],1),i("u-toast",{ref:"uToast"})],1)},r=[]},c257:function(t,n,e){"use strict";e.r(n);var i=e("ec4c"),a=e.n(i);for(var r in i)"default"!==r&&function(t){e.d(n,t,(function(){return i[t]}))}(r);n["default"]=a.a},dd55:function(t,n,e){var i=e("2ae2");i.__esModule&&(i=i.default),"string"===typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);var a=e("4f06").default;a("7b444930",i,!0,{sourceMap:!1,shadowMode:!1})},ec4c:function(t,n,e){"use strict";var i=e("4ea4");Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0,e("96cf");var a=i(e("1da1")),r=e("0a93"),o={data:function(){return{navTitle:"",backgroundOBj:{background:"transparent"},scrollTop:0,urlFromStr:null}},mixins:[r.myMixin,r.runOrderMixin],onShow:function(){var t=this;this.initLang(),this.initData();var n=this.getOpenerEventChannel();n.on("userNoticeClick",function(){var n=(0,a.default)(regeneratorRuntime.mark((function n(e){return regeneratorRuntime.wrap((function(n){while(1)switch(n.prev=n.next){case 0:t.urlFromStr=e.from;case 1:case"end":return n.stop()}}),n)})));return function(t){return n.apply(this,arguments)}}())},methods:{initLang:function(){var t=uni.getStorageSync("run_lang");this.$i18n.locale=t},initData:function(){this.navTitle=this.$t("user.list.cell.my.account.text")},goBackBtnClick:function(){uni.switchTab({url:"/pages/user/user",animationType:"pop-in",animationDuration:200})},uCellEventClick:function(t){var n=this;1==t?uni.navigateTo({url:"/pages/bank/bank_run",success:function(t){n.save_order_from("bank"),t.eventChannel.emit("runOrderClick",{from:"bank"})},animationType:"pop-in",animationDuration:200}):2==t?uni.navigateTo({url:"/pages/bank/bank_withdraw",animationType:"pop-in",animationDuration:200}):3==t&&uni.navigateTo({url:"/pages/bank/bank_usdt",animationType:"pop-in",animationDuration:200})}}};n.default=o},ee53:function(t,n,e){"use strict";e.r(n);var i=e("9371"),a=e("c257");for(var r in a)"default"!==r&&function(t){e.d(n,t,(function(){return a[t]}))}(r);e("58af");var o,c=e("f0c5"),l=Object(c["a"])(a["default"],i["b"],i["c"],!1,null,"c7604558",null,!1,i["a"],o);n["default"]=l.exports}}]);