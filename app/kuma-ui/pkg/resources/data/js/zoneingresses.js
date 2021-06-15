(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["zoneingresses"],{"28aa":function(t,e,a){"use strict";var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"loader-card"},[t.isReady?a("div",{staticClass:"loader-card-content"},[t.isLoading||t.isEmpty?t._e():a("KCard",{attrs:{"border-variant":"noBorder"}},[a("template",{slot:"body"},[t._t("default")],2)],2)],1):t._e(),t.isLoading?a("KEmptyState",{attrs:{"cta-is-hidden":""}},[a("template",{slot:"title"},[a("div",{staticClass:"card-icon mb-3"},[a("KIcon",{attrs:{icon:"spinner",color:"rgba(0, 0, 0, 0.1)",size:"42"}})],1),t._v(" Data Loading... ")])],2):t._e(),t.isEmpty&&!t.isLoading?a("KEmptyState",{attrs:{"cta-is-hidden":""}},[a("template",{slot:"title"},[a("div",{staticClass:"card-icon mb-3"},[a("KIcon",{staticClass:"kong-icon--centered",attrs:{color:"var(--yellow-200)",icon:"warning",size:"42"}})],1),t._v(" There is no data to display. ")])],2):t._e(),t.hasError?a("KEmptyState",{attrs:{"cta-is-hidden":""}},[a("template",{slot:"title"},[a("div",{staticClass:"card-icon mb-3"},[a("KIcon",{staticClass:"kong-icon--centered",attrs:{color:"var(--yellow-200)",icon:"warning",size:"42"}})],1),t._v(" An error has occurred while trying to load this data. ")])],2):t._e()],1)},s=[],i={name:"LoaderCard",props:{title:{type:String,default:null},isLoading:{type:Boolean,default:!1},hasError:{type:Boolean,default:!1},isEmpty:{type:Boolean,default:!1}},computed:{isReady:function(){return!this.isEmpty&&!this.hasError&&!this.isLoading}}},r=i,o=(a("b068"),a("2877")),l=Object(o["a"])(r,n,s,!1,null,null,null);e["a"]=l.exports},"38ba":function(t,e,a){},"43c3":function(t,e,a){"use strict";var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("header",{staticClass:"page-header",class:{"flex justify-between items-center my-6":!t.noflex}},[t._t("default")],2)},s=[],i={props:{noflex:{type:Boolean,default:!1}}},r=i,o=(a("e234"),a("2877")),l=Object(o["a"])(r,n,s,!1,null,null,null);e["a"]=l.exports},4566:function(t,e,a){},"4bba":function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"zoneingresses"},[a("page-header",{attrs:{noflex:""}},[a("breadcrumbs")],1),!1===t.multicluster?a("KEmptyState",{staticClass:"global-api-status"},[a("template",{slot:"title"},[a("KIcon",{staticClass:"kong-icon--centered",attrs:{icon:"dangerCircle",size:"64"}}),t._v(" "+t._s(t.$productName)+" is running in Standalone mode. ")],1),a("template",{slot:"message"},[a("p",[t._v(" To access this page, you must be running in "),a("strong",[t._v("Multi-Zone")]),t._v(" mode. ")])]),a("template",{slot:"cta"},[a("KButton",{attrs:{to:"https://kuma.io/docs/0.6.0/documentation/deployments/",target:"_blank",appearance:"primary"}},[t._v(" Learn More ")])],1)],2):a("FrameSkeleton",[a("DataOverview",{attrs:{"page-size":t.pageSize,"has-error":t.hasError,"is-loading":t.isLoading,"empty-state":t.empty_state,"display-data-table":!0,"table-data":t.tableData,"table-data-is-empty":t.tableDataIsEmpty,"show-warnings":t.tableData.data.some((function(t){return t.withWarnings})),"table-data-function-text":"View","table-data-row":"name"},on:{tableAction:t.tableAction,reloadData:t.loadData}},[a("template",{slot:"pagination"},[a("Pagination",{attrs:{"has-previous":t.previous.length>0,"has-next":t.hasNext},on:{next:t.goToNextPage,previous:t.goToPreviousPage}})],1)],2),!1===t.isEmpty?a("Tabs",{attrs:{"has-error":t.hasError,"is-loading":t.isLoading,tabs:t.tabs,"initial-tab-override":"overview"}},[a("template",{slot:"tabHeader"},[a("div",[a("h3",[t._v(t._s(t.tabGroupTitle))])])]),a("template",{slot:"overview"},[a("LabelList",{attrs:{"has-error":t.entityHasError,"is-loading":t.entityIsLoading,"is-empty":t.entityIsEmpty}},[a("div",[a("ul",t._l(t.entity,(function(e,n){return a("li",{key:n},[e?a("h4",[t._v(" "+t._s(n)+" ")]):t._e(),"status"===n?a("p",[a("KBadge",{attrs:{appearance:"Offline"===e?"danger":"success"}},[t._v(" "+t._s(e)+" ")])],1):a("p",[t._v(" "+t._s(e)+" ")])])})),0)])])],1),a("template",{slot:"insights"},[a("LoaderCard",{attrs:{"has-error":t.entityHasError,"is-loading":t.entityIsLoading,"is-empty":t.entityIsEmpty}},[t.rawEntity?a("div",t._l(t.rawEntity.zoneIngressInsight.subscriptions,(function(e,n){return a("div",{key:n,staticClass:"overview-stack"},[a("h4",{staticClass:"overview-title"},[t._v(" ID: "),a("span",{staticClass:"mono"},[t._v(t._s(e.id))])]),e.globalInstanceId||e.connectTime||e.disconnectTime?a("div",[a("h5",{staticClass:"overview-tertiary-title"},[t._v(" General Information: ")]),a("ul",[e.globalInstanceId?a("li",[a("strong",[t._v("Global Instance ID:")]),t._v(" "),a("span",{staticClass:"mono"},[t._v(t._s(e.globalInstanceId))])]):t._e(),e.connectTime?a("li",[a("strong",[t._v("Last Connected:")]),t._v(" "+t._s(t._f("readableDate")(e.connectTime))+" ")]):t._e(),e.disconnectTime?a("li",[a("strong",[t._v("Last Disconnected:")]),t._v(" "+t._s(t._f("readableDate")(e.disconnectTime))+" ")]):t._e()])]):t._e(),e.status?a("div",[e.status.stat?a("ul",{staticClass:"overview-stat-grid"},t._l(e.status.stat,(function(e,n){return a("li",{key:n},[a("h6",{staticClass:"overview-tertiary-title"},[t._v(" "+t._s(t._f("humanReadable")(n))+": ")]),a("ul",t._l(e,(function(e,n){return a("li",{key:n},[a("strong",[t._v(t._s(t._f("humanReadable")(n))+":")]),t._v(" "),a("span",{staticClass:"mono"},[t._v(t._s(t._f("formatError")(t._f("formatValue")(e))))])])})),0)])})),0):t._e()]):a("KAlert",{staticClass:"mt-4",attrs:{appearance:"info"}},[a("template",{slot:"alertIcon"},[a("KIcon",{attrs:{icon:"portal"}})],1),a("template",{slot:"alertMessage"},[t._v(" There are no Policy statistics for "),a("strong",[t._v(t._s(e.id))])])],2)],1)})),0):t._e()])],1),a("template",{slot:"yaml"},[a("YamlView",{attrs:{title:t.entityOverviewTitle,"has-error":t.entityHasError,"is-loading":t.entityIsLoading,"is-empty":t.entityIsEmpty,content:t.yamlEntity}})],1)],2):t._e()],1)],1)},s=[],i=(a("99af"),a("d81d"),a("b0c0"),a("d3b7"),a("25f0"),a("15fd")),r=(a("96cf"),a("1da1")),o=a("2909"),l=a("5530"),c=a("2f62"),u=a("bc1e"),m=a("b912"),d=a("1d10"),p=a("43c3"),h=a("b6c6"),b=a("1799"),f=a("2778"),v=a("251b"),g=a("ff9d"),y=a("0ada"),_=a("28aa"),E=a("dbf3"),I={name:"ZoneIngresses",metaInfo:{title:"ZoneIngresses"},components:{FrameSkeleton:d["a"],PageHeader:p["a"],Breadcrumbs:h["a"],Pagination:b["a"],DataOverview:f["a"],Tabs:v["a"],YamlView:g["a"],LabelList:y["a"],LoaderCard:_["a"]},filters:{formatValue:function(t){return t?parseInt(t).toLocaleString("en").toString():0},readableDate:function(t){return Object(u["g"])(t)},humanReadable:function(t){return Object(u["b"])(t)},formatError:function(t){return"--"===t?"error calculating":t}},mixins:[m["a"]],data:function(){return{isLoading:!0,isEmpty:!1,hasError:!1,entityIsLoading:!0,entityIsEmpty:!1,entityHasError:!1,tableDataIsEmpty:!1,empty_state:{title:"No Data",message:"There are no Zone Ingresses present."},tableData:{headers:[{key:"actions",hideLabel:!0},{label:"Status",key:"status"},{label:"Name",key:"name"}],data:[]},tabs:[{hash:"#overview",title:"Overview"},{hash:"#insights",title:"Zone Ingress Insights"}],entity:[],rawEntity:null,yamlEntity:null,firstEntity:null,pageSize:this.$pageSize,pageOffset:null,next:null,hasNext:!1,previous:[],tabGroupTitle:null,entityOverviewTitle:null,itemsPerCol:3}},computed:Object(l["a"])(Object(l["a"])({},Object(c["c"])({multicluster:"getMulticlusterStatus"})),{},{pageTitle:function(){return this.$route.meta.title},shareUrl:function(){var t=this,e="".concat(window.location.origin,"#"),a=this.entity,n=function(){return t.$route.query.ns?t.$route.fullPath:"".concat(e).concat(t.$route.fullPath,"?ns=").concat(a.name)};return n()}}),watch:{$route:function(){this.init()}},beforeMount:function(){this.init()},methods:{init:function(){this.multicluster&&this.loadData()},goToPreviousPage:function(){this.pageOffset=this.previous.pop(),this.next=null,this.loadData()},goToNextPage:function(){this.previous.push(this.pageOffset),this.pageOffset=this.next,this.next=null,this.loadData()},tableAction:function(t){var e=t;this.$store.dispatch("updateSelectedTab",this.tabs[0].hash),this.$store.dispatch("updateSelectedTableRow",e.name),this.getEntity(e)},loadData:function(){var t=this;this.isLoading=!0,this.isEmpty=!1;var e=this.$api.getAllZoneIngressOverviews(),a=function(){return e.then((function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},a=!!e.next;a?(t.next=Object(u["e"])(e.next),t.hasNext=!0):t.hasNext=!1;var n=e.items,s=void 0===n?[]:n;s.length>0?(s=s.map((function(t){var e=t.zoneIngressInsight,a=void 0===e?{}:e;return Object(l["a"])(Object(l["a"])({},t),Object(E["n"])(a))})),t.sortEntities(s),t.firstEntity=s[0].name,t.getEntity(s[0]),t.$store.dispatch("updateSelectedTableRow",t.firstEntity),t.tableData.data=Object(o["a"])(s),t.tableDataIsEmpty=!1,t.isEmpty=!1):(t.tableData.data=[],t.tableDataIsEmpty=!0,t.isEmpty=!0,t.getEntity(null))})).catch((function(e){t.hasError=!0,t.isEmpty=!0,console.error(e)})).finally((function(){setTimeout((function(){t.isLoading=!1}),"500")}))};a()},getEntity:function(t){var e=this;return Object(r["a"])(regeneratorRuntime.mark((function a(){var n,s,r,o,c;return regeneratorRuntime.wrap((function(a){while(1)switch(a.prev=a.next){case 0:if(e.entityIsLoading=!0,e.entityIsEmpty=!0,n=["type","name","mesh"],s=setTimeout((function(){e.entityIsEmpty=!0,e.entityIsLoading=!1}),"500"),!t){a.next=26;break}return e.entityIsEmpty=!1,a.prev=6,a.next=9,e.$api.getZoneIngressOverview(t.name);case 9:r=a.sent,o=r.name,r.zoneIngressInsight,c=Object(i["a"])(r,["name","zoneIngressInsight"]),e.tabGroupTitle="Zone Ingress: ".concat(o),e.entityOverviewTitle="Zone Ingress Overview for ".concat(o),e.entity=Object(u["f"])(r,n),e.rawEntity=Object(u["k"])(r),e.yamlEntity=Object(l["a"])({name:o},c),a.next=23;break;case 18:a.prev=18,a.t0=a["catch"](6),e.entity=null,e.entityHasError=!0,e.entityIsEmpty=!0;case 23:return a.prev=23,clearTimeout(s),a.finish(23);case 26:e.entityIsLoading=!1;case 27:case"end":return a.stop()}}),a,null,[[6,18,23,26]])})))()}}},w=I,x=(a("6c462"),a("2877")),O=Object(x["a"])(w,n,s,!1,null,"1725c9da",null);e["default"]=O.exports},"5f76":function(t,e,a){},"6c462":function(t,e,a){"use strict";a("e1c3")},b068:function(t,e,a){"use strict";a("4566")},b6c6:function(t,e,a){"use strict";var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[t.hideBreadcrumbs?t._e():a("Krumbs",{attrs:{items:t.routes}})],1)},s=[],i=(a("4de4"),a("c975"),a("d81d"),a("b0c0"),a("07ac"),a("ac1f"),a("1276"),a("498a"),a("15fd")),r=a("5530"),o=a("bc1e"),l={computed:{pageMesh:function(){return this.$route.params.mesh},routes:function(){var t=this,e=[];this.$route.matched.map((function(a){t.calculateRouteText(a),t.calculateRouteTitle(a);var n=void 0!==a.redirect&&void 0!==a.redirect.name?a.redirect.name:a.name;t.isCurrentRoute(a)&&t.pageMesh&&e.push({key:t.pageMesh,to:{path:"/meshes/".concat(t.pageMesh)},title:"Mesh Overview for ".concat(t.pageMesh),text:t.pageMesh}),t.isCurrentRoute(a)&&a.meta.parent&&"undefined"!==a.meta.parent?e.push({key:a.meta.parent,to:{name:a.meta.parent},title:a.meta.title,text:a.meta.breadcrumb||a.meta.title}):t.isCurrentRoute(a)&&!a.meta.excludeAsBreadcrumb?e.push({key:n,to:{name:n},title:a.meta.title,text:a.meta.breadcrumb||a.meta.title}):a.meta.parent&&"undefined"!==a.meta.parent&&e.push({key:a.meta.parent,to:{name:a.meta.parent},title:a.meta.title,text:a.meta.breadcrumb||a.meta.title})}));var a=this.calculateRouteTextAdvanced(this.$route);return a&&e.push({title:a,text:a}),e},hideBreadcrumbs:function(){return this.$route.query.hide_breadcrumb}},methods:{getBreadcrumbItem:function(t,e,a,n){return{key:t,to:e,title:a,text:n}},isCurrentRoute:function(t){return t.name&&t.name===this.$router.currentRoute.name||t.redirect===this.$router.currentRoute.name},calculateRouteFromQuery:function(t){var e=t.entity_id,a=t.entity_type;if(e&&a){var n=this.$router.resolve({name:"show-".concat(a.split("_")[0]),params:{id:e.split(",")[0]}}).normalizedTo,s=Object(r["a"])(Object(r["a"])({},n),{},{meta:Object(r["a"])({},n.meta)}),i=s.params.id.split("-")[0];return e.split(",").length>1&&e.split(",")[1]&&(i=e.split(",")[1]),s.meta.breadcrumb=i,[Object(r["a"])({},this.getBreadcrumbItem(s.name,s,this.calculateRouteTitle(s),this.calculateRouteText(s)))]}},calculateRouteText:function(t){if(t.path&&t.path.indexOf(":mesh")>-1){var e=this.$router.currentRoute.params;return(e&&e.mesh&&Object(o["h"])(e.mesh)?e.mesh.split("-")[0].trim():e.mesh)||t.meta.breadcrumb||t.meta.title}return t.meta&&(t.meta.breadcrumb||t.meta.title)||t.name||t.meta.breadcrumb||t.meta.title},calculateRouteTitle:function(t){return t.params&&t.params.mesh||t.path.indexOf(":mesh")>-1&&this.$router.currentRoute.params&&this.$router.currentRoute.params.mesh},calculateRouteTextAdvanced:function(t){var e=t.params,a=(e.expandSidebar,Object(i["a"])(e,["expandSidebar"])),n="mesh-overview"===t.name,s=Object.assign({},a,{mesh:null});return n?e.mesh:Object.values(s).filter((function(t){return t}))[0]}}},c=l,u=(a("e7ab"),a("2877")),m=Object(u["a"])(c,n,s,!1,null,null,null);e["a"]=m.exports},e1c3:function(t,e,a){},e234:function(t,e,a){"use strict";a("38ba")},e7ab:function(t,e,a){"use strict";a("5f76")}}]);