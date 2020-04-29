(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-8aeae28e"],{a62d:function(t,e,a){"use strict";a.r(e);var i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"fault-injections"},[a("FrameSkeleton",[a("DataOverview",{attrs:{"page-size":6,"has-error":t.hasError,"is-loading":t.isLoading,"is-empty":t.isEmpty,"empty-state":t.empty_state,"display-data-table":!0,"table-data":t.tableData,"table-data-is-empty":t.tableDataIsEmpty,"table-data-function-text":"View","table-data-row":"name"},on:{tableAction:t.tableAction,reloadData:t.bootstrap}}),a("Tabs",{attrs:{"has-error":t.hasError,"is-loading":t.isLoading,"is-empty":t.isEmpty,tabs:t.tabs,"tab-group-title":t.tabGroupTitle,"initial-tab-override":"overview"}},[a("template",{slot:"overview"},[a("LabelList",{attrs:{"has-error":t.entityHasError,"is-loading":t.entityIsLoading,"is-empty":t.entityIsEmpty,items:t.entity}})],1),a("template",{slot:"yaml"},[a("YamlView",{attrs:{lang:"yaml",title:t.entityOverviewTitle,"has-error":t.entityHasError,"is-loading":t.entityIsLoading,"is-empty":t.entityIsEmpty,content:t.rawEntity}})],1)],2)],1)],1)},n=[],s=a("75fc"),o=(a("7f7f"),a("d7c2")),r=a("8218"),l=a("ad2f"),c=a("1d10"),y=a("2778"),u=a("251b"),m=a("ff9d"),h=a("0ada"),p={name:"FaultInjections",metaInfo:{title:"Fault Injections"},components:{FrameSkeleton:c["a"],DataOverview:y["a"],Tabs:u["a"],YamlView:m["a"],LabelList:h["a"]},mixins:[l["a"],r["a"]],data:function(){return{isLoading:!0,isEmpty:!1,hasError:!1,entityIsLoading:!0,entityIsEmpty:!1,entityHasError:!1,tableDataIsEmpty:!1,empty_state:{title:"No Data",message:"There are no Fault Injections present."},tableData:{headers:[{key:"actions",hideLabel:!0},{label:"Name",key:"name"},{label:"Mesh",key:"mesh"},{label:"Type",key:"type"}],data:[]},tabs:[{hash:"#overview",title:"Overview"},{hash:"#yaml",title:"YAML"}],entity:null,rawEntity:null,firstEntity:null}},computed:{tabGroupTitle:function(){var t=this.entity;return t?"Fault Injection: ".concat(t.name):null},entityOverviewTitle:function(){var t=this.entity;return t?"Entity Overview for ".concat(t.name):null},formattedRawEntity:function(){var t=this.formatForCLI(this.rawEntity);return t}},watch:{$route:function(t,e){this.bootstrap()}},beforeMount:function(){this.bootstrap()},methods:{tableAction:function(t){var e=t;this.$store.dispatch("updateSelectedTab",this.tabs[0].hash),this.$store.dispatch("updateSelectedTableRow",e.name),this.getEntity(e)},bootstrap:function(){var t=this;this.isLoading=!0,this.isEmpty=!1;var e=this.$route.params.mesh,a="all"===e?this.$api.getAllFaultInjections():this.$api.getAllFaultInjectionsFromMesh(e),i=function(){return a.then((function(e){if(e.items.length>0){var a=e.items;t.sortEntities(a),t.firstEntity=a[0].name,t.getEntity(a[0]),t.$store.dispatch("updateSelectedTableRow",t.firstEntity),t.tableData.data=Object(s["a"])(a),t.tableDataIsEmpty=!1}else t.tableData.data=[],t.tableDataIsEmpty=!0,t.getEntity(null)})).catch((function(e){t.hasError=!0,console.error(e)})).finally((function(){setTimeout((function(){t.isLoading=!1}),"500")}))};i()},getEntity:function(t){var e=this;this.entityIsLoading=!0,this.entityIsEmpty=!1;var a=this.$route.params.mesh;if(t&&null!==t){var i="all"===a?t.mesh:a;return this.$api.getFaultInjection(i,t.name).then((function(t){if(t){var a=["type","name","mesh"];e.entity=Object(o["a"])(t,a),e.rawEntity=t}else e.entity=null,e.entityIsEmpty=!0})).catch((function(t){e.entityHasError=!0,console.error(t)})).finally((function(){setTimeout((function(){e.entityIsLoading=!1}),"500")}))}setTimeout((function(){e.entityIsEmpty=!0,e.entityIsLoading=!1}),"500")}}},d=p,f=a("2877"),b=Object(f["a"])(d,i,n,!1,null,null,null);e["default"]=b.exports},ad2f:function(t,e,a){"use strict";var i=a("e80b"),n=a.n(i);e["a"]={methods:{formatForCLI:function(t){var e='echo "',a='" | kumactl apply -f -',i=n()(t);return"".concat(e).concat(i).concat(a)}}}}}]);