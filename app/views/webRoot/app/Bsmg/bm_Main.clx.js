/*
 * App URI: app/Bsmg/bm_Main
 * Source Location: app/Bsmg/bm_Main.clx
 *
 * This file was generated by eXbuilder6 compiler, Don't edit manually.
 */
(function(){
	var app = new cpr.core.App("app/Bsmg/bm_Main", {
		onPrepare: function(loader){
		},
		onCreate: function(/* cpr.core.AppInstance */ app, exports){
			var linker = {};
			// Start - User Script
			/************************************************
			 * bm_Main.js
			 * Created at 2022. 5. 10. 오전 9:48:13.
			 *
			 * @author SW2Team
			 ************************************************/
			
			var dataManager = cpr.core.Module.require("lib/DataManager");
			
			exports.setMemberInfo = function(dm_memberInfo){
				var dmMemberInfo = app.lookup("dm_memberInfo");
				dmMemberInfo.build(dm_memberInfo);
			}
			
			/*
			 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
			 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
			 */
			function onBodyLoad(/* cpr.events.CEvent */ e){
				dataManager = getDataManager();
				app.lookup("sms_chkLogin").send();
			}
			
			
			/*
			 * "사용자 등록" 버튼(user_regist)에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onUser_registClick(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var user_regist = e.control;
				app.getRootAppInstance().openDialog("app/Bsmg/bm_regist", {
					width : 800, height : 600
				}, function(dialog){
					dialog.ready(function(dialogApp){
						dialog.modal = true;
						dialog.headerVisible = true;
						dialog.headerClose = true;
						dialog.headerMovable = true;
						dialog.resizable = true;
						dialog.headerTitle = "사용자 등록";
						dialog.addEventListener("keyup", function(e){
							if (e.keyCode == 27){
								dialog.close();
							}
						});
					});
				})
				
			}
			
			
			/*
			 * "로그아웃" 버튼(logout)에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onLogoutClick(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				
				var logout = e.control;
			//	console.log(app.lookup("Result").getString("ResultCode"));
				
				if(confirm("로그아웃 하시겠습니까?")){
					app.lookup("sms_logout").send();
					
				}
				
			}
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_logoutSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_logout = e.control;
				var result = app.lookup("Result").getString("ResultCode");
				
				if(result == 0){
					alert("정상적으로 로그아웃 되었습니다.");
					cpr.core.App.load("app/Bsmg/bm_login", function(newapp){
						app.close();
						var newInst = newapp.createNewInstance();
						newInst.run();
					});
					return; 
				} else {
					alert(getErrorString(result));
				}
			}
			
			
			/*
			 * "사용자 관리" 버튼(userManagement)에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onUserManagementClick(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var userManagement = e.control;
				app.getRootAppInstance().openDialog("app/Bsmg/bm_userManagement", {
					width : 800, height : 600
				}, function(dialog){
					dialog.ready(function(dialogApp){
						dialog.modal = true;
						dialog.headerVisible = true;
						dialog.headerClose = true;
						dialog.headerMovable = true;
						dialog.resizable = true;
						dialog.headerTitle = "사용자 관리";
						dialog.addEventListener("keyup", function(e){
							if (e.keyCode == 27){
								dialog.close();
							}
						});
					});
				})
			}
			
			
			
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_chkLoginSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_chkLogin = e.control;
				var result = app.lookup("Result").getString("ResultCode");
				if(result == 0) {
					app.lookup("sms_setRankPart").send();
				} else {
					alert(getErrorString(result));
					cpr.core.App.load("app/Bsmg/bm_login", function(newapp){
						app.close();
						var newInst = newapp.createNewInstance();
						newInst.run();
					});
					return; 
				}
			}
			
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_setRankPartSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_setRankPart = e.control;
				
				var result = app.lookup("Result").getValue("ResultCode");
				if(result == 0) {
					var memInfo = app.lookup("dm_memberInfo");
					var mem_rank = memInfo.getString("mem_rank");	
					var mem_part = memInfo.getString("mem_part");	
					if(mem_rank < 3){
						app.lookup("user_regist").visible = true;
						app.lookup("userManagement").visible = true;
					}
					
					
					var dsRankList = app.lookup("ds_rank");
					var rankRow = dsRankList.findFirstRow("rank_idx == " + mem_rank);
					var rankName = rankRow.getValue("rank_name");
					app.lookup("Main_RankOpb").value = rankName;
					
					var dsPartList = app.lookup("ds_part");
					var partRow = dsPartList.findFirstRow("part_idx == " + mem_part);
					var partName = partRow.getValue("part_name");
					app.lookup("Main_PartOpb").value = partName;
					
					
					// 직급, 부서 dataManager에 저장
					dataManager.setRankList(dsRankList);
					dataManager.setPartList(dsPartList);
					
				} else {
					alert(getErrorString(result));
				} 
				app.getContainer().redraw();
				
			};
			// End - User Script
			
			// Header
			var dataSet_1 = new cpr.data.DataSet("ds_List");
			dataSet_1.parseData({
				"columns" : [
					{"name": "label"},
					{"name": "value"},
					{"name": "parent"}
				]
			});
			app.register(dataSet_1);
			
			var dataSet_2 = new cpr.data.DataSet("ds_rank");
			dataSet_2.parseData({
				"columns" : [
					{"name": "rank_name"},
					{
						"name": "rank_idx",
						"dataType": "number"
					}
				]
			});
			app.register(dataSet_2);
			
			var dataSet_3 = new cpr.data.DataSet("ds_part");
			dataSet_3.parseData({
				"columns" : [
					{"name": "part_name"},
					{
						"name": "part_idx",
						"dataType": "number"
					}
				]
			});
			app.register(dataSet_3);
			var dataMap_1 = new cpr.data.DataMap("dm_memberInfo");
			dataMap_1.parseData({
				"columns" : [
					{"name": "mem_id"},
					{"name": "mem_name"},
					{
						"name": "mem_rank",
						"dataType": "number"
					},
					{
						"name": "mem_part",
						"dataType": "number"
					}
				]
			});
			app.register(dataMap_1);
			
			var dataMap_2 = new cpr.data.DataMap("Result");
			dataMap_2.parseData({
				"columns" : [{"name": "ResultCode"}]
			});
			app.register(dataMap_2);
			var submission_1 = new cpr.protocols.Submission("sms_logout");
			submission_1.action = "/bsmg/login/logout";
			submission_1.addResponseData(dataMap_2, false);
			if(typeof onSms_logoutSubmitDone == "function") {
				submission_1.addEventListener("submit-done", onSms_logoutSubmitDone);
			}
			app.register(submission_1);
			
			var submission_2 = new cpr.protocols.Submission("sms_chkLogin");
			submission_2.method = "get";
			submission_2.action = "/bsmg/login/chkLogin";
			submission_2.addResponseData(dataMap_2, false);
			submission_2.addResponseData(dataMap_1, false);
			if(typeof onSms_chkLoginSubmitDone == "function") {
				submission_2.addEventListener("submit-done", onSms_chkLoginSubmitDone);
			}
			app.register(submission_2);
			
			var submission_3 = new cpr.protocols.Submission("sms_setRankPart");
			submission_3.async = true;
			submission_3.method = "get";
			submission_3.action = "/bsmg/setting/rankPart";
			submission_3.addResponseData(dataSet_2, false);
			submission_3.addResponseData(dataSet_3, false);
			submission_3.addResponseData(dataMap_2, false);
			if(typeof onSms_setRankPartSubmitDone == "function") {
				submission_3.addEventListener("submit-done", onSms_setRankPartSubmitDone);
			}
			app.register(submission_3);
			
			app.supportMedia("all and (min-width: 1024px)", "default");
			app.supportMedia("all and (min-width: 500px) and (max-width: 1023px)", "tablet");
			app.supportMedia("all and (max-width: 499px)", "mobile");
			
			// Configure root container
			var container = app.getContainer();
			container.style.css({
				"border-bottom-style" : "none",
				"border-top-style" : "none",
				"border-right-style" : "none",
				"width" : "100%",
				"top" : "0px",
				"height" : "100%",
				"left" : "0px",
				"border-left-style" : "none"
			});
			
			// Layout
			var responsiveXYLayout_1 = new cpr.controls.layouts.ResponsiveXYLayout();
			container.setLayout(responsiveXYLayout_1);
			
			// UI Configuration
			var output_1 = new cpr.controls.Output("out1");
			output_1.style.css({
				"border-right-style" : "dashed",
				"font-weight" : "bold",
				"border-left-style" : "dashed",
				"border-bottom-style" : "dashed",
				"border-top-style" : "dashed",
				"text-align" : "center"
			});
			output_1.bind("value").toDataMap(app.lookup("dm_memberInfo"), "mem_name");
			container.addChild(output_1, {
				positions: [
					{
						"media": "all and (min-width: 1024px)",
						"top": "20px",
						"left": "430px",
						"width": "92px",
						"height": "45px"
					}, 
					{
						"media": "all and (min-width: 500px) and (max-width: 1023px)",
						"hidden": false,
						"top": "20px",
						"left": "210px",
						"width": "45px",
						"height": "45px"
					}, 
					{
						"media": "all and (max-width: 499px)",
						"hidden": false,
						"top": "20px",
						"left": "147px",
						"width": "31px",
						"height": "45px"
					}
				]
			});
			
			var output_2 = new cpr.controls.Output();
			output_2.value = "사용자";
			output_2.style.css({
				"background-color" : "#ddecd9",
				"font-weight" : "bold"
			});
			container.addChild(output_2, {
				positions: [
					{
						"media": "all and (min-width: 1024px)",
						"top": "20px",
						"left": "383px",
						"width": "48px",
						"height": "45px"
					}, 
					{
						"media": "all and (min-width: 500px) and (max-width: 1023px)",
						"hidden": false,
						"top": "20px",
						"left": "187px",
						"width": "23px",
						"height": "45px"
					}, 
					{
						"media": "all and (max-width: 499px)",
						"hidden": false,
						"top": "20px",
						"left": "131px",
						"width": "16px",
						"height": "45px"
					}
				]
			});
			
			var output_3 = new cpr.controls.Output();
			output_3.value = "직급";
			output_3.style.css({
				"background-color" : "#ddecd9",
				"font-weight" : "bold",
				"text-align" : "center"
			});
			container.addChild(output_3, {
				positions: [
					{
						"media": "all and (min-width: 1024px)",
						"top": "20px",
						"left": "521px",
						"width": "48px",
						"height": "45px"
					}, 
					{
						"media": "all and (min-width: 500px) and (max-width: 1023px)",
						"hidden": false,
						"top": "20px",
						"left": "254px",
						"width": "23px",
						"height": "45px"
					}, 
					{
						"media": "all and (max-width: 499px)",
						"hidden": false,
						"top": "20px",
						"left": "178px",
						"width": "16px",
						"height": "45px"
					}
				]
			});
			
			var output_4 = new cpr.controls.Output();
			output_4.value = "부서";
			output_4.style.css({
				"background-color" : "#ddecd9",
				"font-weight" : "bold",
				"text-align" : "center"
			});
			container.addChild(output_4, {
				positions: [
					{
						"media": "all and (min-width: 1024px)",
						"top": "20px",
						"left": "663px",
						"width": "48px",
						"height": "45px"
					}, 
					{
						"media": "all and (min-width: 500px) and (max-width: 1023px)",
						"hidden": false,
						"top": "20px",
						"left": "324px",
						"width": "23px",
						"height": "45px"
					}, 
					{
						"media": "all and (max-width: 499px)",
						"hidden": false,
						"top": "20px",
						"left": "227px",
						"width": "16px",
						"height": "45px"
					}
				]
			});
			
			var output_5 = new cpr.controls.Output("Main_RankOpb");
			output_5.style.css({
				"border-right-style" : "dashed",
				"font-weight" : "bold",
				"border-left-style" : "dashed",
				"border-bottom-style" : "dashed",
				"border-top-style" : "dashed",
				"text-align" : "center"
			});
			container.addChild(output_5, {
				positions: [
					{
						"media": "all and (min-width: 1024px)",
						"top": "20px",
						"left": "568px",
						"width": "96px",
						"height": "45px"
					}, 
					{
						"media": "all and (min-width: 500px) and (max-width: 1023px)",
						"hidden": false,
						"top": "20px",
						"left": "277px",
						"width": "47px",
						"height": "45px"
					}, 
					{
						"media": "all and (max-width: 499px)",
						"hidden": false,
						"top": "20px",
						"left": "194px",
						"width": "33px",
						"height": "45px"
					}
				]
			});
			
			var output_6 = new cpr.controls.Output("Main_PartOpb");
			output_6.style.css({
				"border-right-style" : "dashed",
				"border-bottom-color" : "#000000",
				"font-weight" : "bolder",
				"border-left-style" : "dashed",
				"border-left-color" : "#000000",
				"border-top-color" : "#000000",
				"border-right-color" : "#000000",
				"border-bottom-style" : "dashed",
				"border-top-style" : "dashed",
				"text-align" : "center"
			});
			container.addChild(output_6, {
				positions: [
					{
						"media": "all and (min-width: 1024px)",
						"top": "20px",
						"left": "710px",
						"width": "94px",
						"height": "45px"
					}, 
					{
						"media": "all and (min-width: 500px) and (max-width: 1023px)",
						"hidden": false,
						"top": "20px",
						"left": "347px",
						"width": "46px",
						"height": "45px"
					}, 
					{
						"media": "all and (max-width: 499px)",
						"hidden": false,
						"top": "20px",
						"left": "243px",
						"width": "32px",
						"height": "45px"
					}
				]
			});
			
			var button_1 = new cpr.controls.Button("user_regist");
			button_1.visible = false;
			button_1.value = "사용자 등록";
			button_1.style.css({
				"background-color" : "#5497da",
				"background-image" : "linear-gradient(#fcfeff,#e0e1e2)"
			});
			if(typeof onUser_registClick == "function") {
				button_1.addEventListener("click", onUser_registClick);
			}
			container.addChild(button_1, {
				positions: [
					{
						"media": "all and (min-width: 1024px)",
						"top": "20px",
						"left": "1108px",
						"width": "126px",
						"height": "31px"
					}, 
					{
						"media": "all and (min-width: 500px) and (max-width: 1023px)",
						"hidden": false,
						"top": "20px",
						"left": "541px",
						"width": "62px",
						"height": "31px"
					}, 
					{
						"media": "all and (max-width: 499px)",
						"hidden": false,
						"top": "20px",
						"left": "379px",
						"width": "43px",
						"height": "31px"
					}
				]
			});
			
			var button_2 = new cpr.controls.Button("logout");
			button_2.value = "로그아웃";
			if(typeof onLogoutClick == "function") {
				button_2.addEventListener("click", onLogoutClick);
			}
			container.addChild(button_2, {
				positions: [
					{
						"media": "all and (min-width: 1024px)",
						"top": "20px",
						"left": "916px",
						"width": "126px",
						"height": "31px"
					}, 
					{
						"media": "all and (min-width: 500px) and (max-width: 1023px)",
						"hidden": false,
						"top": "20px",
						"left": "447px",
						"width": "62px",
						"height": "31px"
					}, 
					{
						"media": "all and (max-width: 499px)",
						"hidden": false,
						"top": "20px",
						"left": "313px",
						"width": "43px",
						"height": "31px"
					}
				]
			});
			
			var output_7 = new cpr.controls.Output();
			output_7.style.css({
				"background-repeat" : "no-repeat",
				"background-size" : "contain",
				"color" : "darkGreen",
				"font-weight" : "bold",
				"font-size" : "20px",
				"font-style" : "normal",
				"background-position" : "center",
				"background-origin" : "padding-box",
				"background-image" : "url('images/reportImg.png')"
			});
			container.addChild(output_7, {
				positions: [
					{
						"media": "all and (min-width: 1024px)",
						"top": "0px",
						"right": "1800px",
						"bottom": "840px",
						"left": "0px"
					}, 
					{
						"media": "all and (min-width: 500px) and (max-width: 1023px)",
						"hidden": false,
						"top": "0px",
						"right": "879px",
						"bottom": "840px",
						"left": "0px"
					}, 
					{
						"media": "all and (max-width: 499px)",
						"hidden": false,
						"top": "0px",
						"right": "615px",
						"bottom": "840px",
						"left": "0px"
					}
				]
			});
			
			var button_3 = new cpr.controls.Button("userManagement");
			button_3.visible = false;
			button_3.value = "사용자 관리";
			button_3.style.css({
				"background-color" : "#5497da",
				"background-image" : "linear-gradient(#fcfeff,#e0e1e2)"
			});
			if(typeof onUserManagementClick == "function") {
				button_3.addEventListener("click", onUserManagementClick);
			}
			container.addChild(button_3, {
				positions: [
					{
						"media": "all and (min-width: 1024px)",
						"top": "20px",
						"left": "1244px",
						"width": "126px",
						"height": "31px"
					}, 
					{
						"media": "all and (min-width: 500px) and (max-width: 1023px)",
						"hidden": false,
						"top": "20px",
						"left": "607px",
						"width": "62px",
						"height": "31px"
					}, 
					{
						"media": "all and (max-width: 499px)",
						"hidden": false,
						"top": "20px",
						"left": "425px",
						"width": "43px",
						"height": "31px"
					}
				]
			});
			
			var tabFolder_1 = new cpr.controls.TabFolder();
			
			var tabItem_1 = (function(tabFolder){
				var tabItem_1 = new cpr.controls.TabItem();
				tabItem_1.text = "일일 업무보고";
				var group_1 = new cpr.controls.Container();
				// Layout
				var xYLayout_1 = new cpr.controls.layouts.XYLayout();
				group_1.setLayout(xYLayout_1);
				(function(container){
					var embeddedApp_1 = new cpr.controls.EmbeddedApp("ea1");
					cpr.core.App.load("app/Bsmg/bm_list", function(app) {
						if(app){
							embeddedApp_1.app = app;
						}
					});
					container.addChild(embeddedApp_1, {
						"top": "1px",
						"left": "2px",
						"width": "1393px",
						"height": "800px"
					});
				})(group_1);
				tabItem_1.content = group_1;
				return tabItem_1;
			})(tabFolder_1);
			tabFolder_1.addTabItem(tabItem_1);
			
			var tabItem_2 = (function(tabFolder){
				var tabItem_2 = new cpr.controls.TabItem();
				tabItem_2.text = "주간 업무보고";
				var group_2 = new cpr.controls.Container();
				// Layout
				var xYLayout_2 = new cpr.controls.layouts.XYLayout();
				group_2.setLayout(xYLayout_2);
				(function(container){
					var embeddedApp_2 = new cpr.controls.EmbeddedApp("ea2");
					cpr.core.App.load("app/Bsmg/bm_weekRptList", function(app) {
						if(app){
							embeddedApp_2.app = app;
						}
					});
					container.addChild(embeddedApp_2, {
						"top": "1px",
						"left": "2px",
						"width": "1393px",
						"height": "800px"
					});
				})(group_2);
				tabItem_2.content = group_2;
				return tabItem_2;
			})(tabFolder_1);
			tabFolder_1.addTabItem(tabItem_2);
			tabFolder_1.setSelectedTabItem(tabItem_1);
			container.addChild(tabFolder_1, {
				positions: [
					{
						"media": "all and (min-width: 1024px)",
						"top": "85px",
						"left": "0px",
						"width": "1403px",
						"height": "839px"
					}, 
					{
						"media": "all and (min-width: 500px) and (max-width: 1023px)",
						"hidden": false,
						"top": "85px",
						"left": "0px",
						"width": "685px",
						"height": "839px"
					}, 
					{
						"media": "all and (max-width: 499px)",
						"hidden": false,
						"top": "85px",
						"left": "0px",
						"width": "480px",
						"height": "839px"
					}
				]
			});
			if(typeof onBodyLoad == "function"){
				app.addEventListener("load", onBodyLoad);
			}
		}
	});
	app.title = "bm_Main";
	cpr.core.Platform.INSTANCE.register(app);
})();
