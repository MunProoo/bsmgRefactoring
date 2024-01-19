/*
 * App URI: app/Bsmg/bm_weekRptList
 * Source Location: app/Bsmg/bm_weekRptList.clx
 *
 * This file was generated by eXbuilder6 compiler, Don't edit manually.
 */
(function(){
	var app = new cpr.core.App("app/Bsmg/bm_weekRptList", {
		onPrepare: function(loader){
		},
		onCreate: function(/* cpr.core.AppInstance */ app, exports){
			var linker = {};
			// Start - User Script
			/************************************************
			 * bm_weekRpt.js
			 * Created at 2022. 6. 3. 오후 3:44:36.
			 *
			 * @author SW2Team
			 ************************************************/
			
			/* 페이지처리 */
			var RowCount = 6; // 한페이지 로우 개수
			var SearchFlag = false;
			var AttrFlag = false;
			var NewSearch = false; // 현재 페이지가 1이 아닌 상태. 새롭게 검색한다면 1페이지로 돌아가게
			var REFRESH = false; // 새로 검색하는건지 페이지 이동인지 구분자
			
			
			/*
			 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
			 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
			 */
			function onBodyLoad(/* cpr.events.CEvent */ e){
				app.lookup("sms_getCategoryList").send();
				setPaging(0, 1, RowCount, 5);
				sendRptListRequest();
			}
			
			function setPaging(totCnt, pageIdx, RowCount, pageSize){
				var pageIndexer = app.lookup("pageIndex");
				pageIndexer.totalRowCount = totCnt;
				pageIndexer.currentPageIndex = pageIdx;
				pageIndexer.pageRowCount = RowCount;
				pageIndexer.viewPageCount = pageSize;
				if(totCnt == 0){
					pageIndexer.visible = false;
				} else {
					pageIndexer.visible = true;
				}
				app.getContainer().redraw();
			}
			
			function sendRptListRequest(){
				var pageIndexer = app.lookup("pageIndex");
				var pageIdx = pageIndexer.currentPageIndex;
				var offset = (pageIdx - 1) * RowCount;
				var sms_getWeekRptList = app.lookup("sms_getWeekRptList");
				
				sms_getWeekRptList.setParameters("offset", offset);
				sms_getWeekRptList.setParameters("limit", RowCount);
				sms_getWeekRptList.send();
			}
			
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_getWeekRptListSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_getWeekRptList = e.control;
				var pageIndexer = app.lookup("pageIndex");
				var result = app.lookup("Result").getString("ResultCode");
				
				if(result == 0){
					var totalCount = app.lookup("totalCount").getValue("Count");
					
					AllGridColorWhite();
					
					if(REFRESH){
						setPaging(Number(totalCount), 1, RowCount, 5);
						REFRESH = false;
					} else {
						setPaging(Number(totalCount), pageIndexer.currentPageIndex, RowCount, 5);
					}
					SearchFlag = false;
					AttrFlag = false;
					return;
				} else {
					alert("주간업무보고 리스트를 불러오는데 실패하였습니다.");
					return;
				}
			}
			
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_getCategoryListSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_getCategoryList = e.control;
				var result = app.lookup("Result").getString("ResultCode");
				if(result == 0){
					app.lookup("tre1").redraw();
					return;
				} else {
					alert("카테고리 트리 갱신 실패");
					return;
				}
			}
			
			
			
			
			/*
			 * 페이지 인덱서에서 selection-change 이벤트 발생 시 호출.
			 * Page index를 선택하여 선택된 페이지가 변경된 후에 발생하는 이벤트.
			 */
			function onPageIndexSelectionChange(/* cpr.events.CSelectionEvent */ e){
				/** 
				 * @type cpr.controls.PageIndexer
				 */
				var pageIndex = e.control;
				var selectionEvent = new cpr.events.CSelectionEvent("pagechange", {
					oldSelection: e.oldSelection,
					newSelection: e.newSelection
				});
				app.dispatchEvent(selectionEvent);
				// 기본처리가 중단되었을 때 변경을 취소함.
				if(SearchFlag){
					sendSearchRequest();
				} else if(AttrFlag){
					sendPartRptRequest();
				} else{
					sendRptListRequest();
				}
				
			}
			
			
			/*
			 * 페이지 인덱서에서 before-selection-change 이벤트 발생 시 호출.
			 * Page index를 선택하여 선택된 페이지가 변경되기 전에 발생하는 이벤트. 다음 이벤트로 selection-change를 발생합니다.
			 */
			function onPageIndexBeforeSelectionChange(/* cpr.events.CSelectionEvent */ e){
				/** 
				 * @type cpr.controls.PageIndexer
				 */
				var pageIndex = e.control;
				var selectionEvent = new cpr.events.CSelectionEvent("before-pagechange",{
					oldSelection: e.oldSelection,
					newSelection: e.newSelection
				});
				app.dispatchEvent(selectionEvent);
				if(selectionEvent.defaultPrevented == true) {
					e.preventDefault();
				}
			
			}
			
			function sendSearchRequest(){
				var pageIndexer = app.lookup("pageIndex");
				if (!SearchFlag){ // 검색중이 아니라면 페이지인덱서 새롭게 세팅
					NewSearch = true; 
					pageIndexer.currentPageIndex = 1;
				} else {
					NewSearch = false;
				}
				
				var pageIdx = pageIndexer.currentPageIndex;
				var offset = (pageIdx - 1) * RowCount;
				var combo = app.lookup("cmb1").value;
				var input = app.lookup("ipb1").value;
				
				app.lookup("ds_weekRptList").clear();
				app.lookup("dm_search").setValue("search_combo", combo);
				app.lookup("dm_search").setValue("search_input", input);
				var smsGetWeekRptSearch = app.lookup("sms_getWeekRptSearch");
				smsGetWeekRptSearch.setParameters("offset", offset);
				smsGetWeekRptSearch.setParameters("limit", RowCount);
				smsGetWeekRptSearch.send();
			}
			
			
			
			/*
			 * 서치 인풋에서 search 이벤트 발생 시 호출.
			 * Searchinput의 enter키 또는 검색버튼을 클릭하여 인풋의 값이 Search될때 발생하는 이벤트
			 */
			function onIpb1Search(/* cpr.events.CUIEvent */ e){
				/** 
				 * @type cpr.controls.SearchInput
				 */
				var ipb1 = e.control;
				sendSearchRequest();
			}
			
			
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_getWeekRptSearchSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_getWeekRptSearch = e.control;
				var result = app.lookup("Result").getString("ResultCode");
				var totalCount = app.lookup("totalCount").getValue("Count");
				var pageIndexer = app.lookup("pageIndex");
				if(result == 0){
					SearchFlag = true;
					AttrFlag = false;
					
					if(totalCount == 0){
						alert("자료가 없습니다.");
						return;
					}
					
					AllGridColorWhite();
					var combo = app.lookup("cmb1").value;
					if(combo != 0){
						gridColorImpact(combo);
					}
					
					if(NewSearch){  // 새로 검색한 거면 1페이지부터
						setPaging(Number(totalCount), 1, RowCount, 5);
					} else {
						setPaging(Number(totalCount), pageIndexer.currentPageIndex, RowCount, 5);
					}
					app.getContainer().redraw();
					return;
				}
			}
			
			
			
			
			/*
			 * 버튼에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onButtonClick(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var button = e.control;
				REFRESH = true;
				sendRptListRequest();
			}
			
			
			/*
			 * 트리에서 item-dblclick 이벤트 발생 시 호출.
			 * 아이템 더블 클릭시 발생하는 이벤트.
			 */
			function onTre1ItemDblclick(/* cpr.events.CItemEvent */ e){
				/** 
				 * @type cpr.controls.Tree
				 */
				var tre1 = e.control;
				var grid = app.lookup("grd1");
				var pageIndexer = app.lookup("pageIndex");
				pageIndexer.currentPageIndex = 1;
				var partValue = tre1.getSelectionLast().value;
				partValue = partValue.split("-")[1]
			//	console.log("partValue : ", partValue)
				if(partValue == undefined){
					partValue = 0;
				}
				
				app.lookup("dm_part").setValue("part_value", partValue);
				sendPartRptRequest();
			
			}
			
			function sendPartRptRequest(){
				var pageIndexer = app.lookup("pageIndex");
				if(!AttrFlag){
					NewSearch = true;
					pageIndexer.currentPageIndex = 1;
				} else {
					NewSearch = false;
				} 
				
				var pageIdx = pageIndexer.currentPageIndex;
				var offset = (pageIdx - 1) * RowCount;
				
				app.lookup("ds_weekRptList").clear();
				app.lookup("sms_getWeekRptCategory").setParameters("offset", offset);
				app.lookup("sms_getWeekRptCategory").setParameters("limit", RowCount);
				app.lookup("sms_getWeekRptCategory").send();
			}
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_getWeekRptCategorySubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_getWeekRptCategory = e.control;
				var result = app.lookup("Result").getString("ResultCode");
				var totalCount = app.lookup("totalCount").getValue("Count");
				var pageIndexer = app.lookup("pageIndex");
				
				if(result == 0){
					SearchFlag = false;
					AttrFlag = true;
					
					if(totalCount == 0){
						alert("자료가 없습니다.");
						return;
					}
					AllGridColorWhite();
					gridColorImpact(3);
					
					if(NewSearch){
						setPaging(Number(totalCount),1,RowCount, 5);
					} else {
						setPaging(Number(totalCount), pageIndexer.currentPageIndex, RowCount, 5);
					}
					app.getContainer().redraw();
					return;
				}
			}
			
			
			
			
			/*
			 * 그리드에서 row-dblclick 이벤트 발생 시 호출.
			 * detail이 row를 더블클릭 한 경우 발생하는 이벤트.
			 */
			function onGrd1RowDblclick(/* cpr.events.CGridMouseEvent */ e){
				/** 
				 * @type cpr.controls.Grid
				 */
				var grd1 = e.control;
				var rowIndex = grd1.getSelectedRowIndex();
				var row = grd1.getRow(rowIndex);
				var paramValue = row.getValue("wRpt_idx");
				
				app.getRootAppInstance().openDialog("app/Bsmg/bm_weekRptView", {
					width : 1000, height : 800
				}, function(dialog){
					dialog.modal = true;
					dialog.headerVisible = true;
					dialog.headerMovable = true;
					dialog.headerTitle = "주간 업무보고 상세 확인";
					dialog.headerClose = true;
					dialog.addEventListener("keyup", function(e){
						if(e.keyCode == 27){ // ESC
							dialog.close();
						}
					});
					dialog.initValue = {
						wRpt_idx : paramValue
					}
				}).then(function(returnValue){
					if(returnValue == 1){
						REFRESH = true;
						sendRptListRequest();
					}
				});
			}
			
			function gridColorWhite(cellIndex){
				var grd = app.lookup("grd1");
				var voCell = grd.detail.getColumn(cellIndex);
			//	var voExpress = "switch(getValue(\"color\")){\n\tcase \"Red\" : \"red\"\n\tdefault : \"#transparent\"\n}";
			//	voCell.style.bind("background-color").toExpression(voExpress);
				voCell.style.bind("background-color").toExpression("transparent");
			}
			
			function AllGridColorWhite(){
				for(var i=1; i<5; i++){
					gridColorWhite(i);
				}
			}
			
			function gridColorImpact(cellIndex){
				var grd = app.lookup("grd1");
				var voCell = grd.detail.getColumn(cellIndex);
			//	var voExpress = "switch(getValue(\"color\")){\n\tcase \"Red\" : \"red\"\n\tdefault : \"#ffdede\"\n}";
			//	voCell.style.bind("background-color").toExpression(voExpress);
				voCell.style.bind("background-color").toExpression("\"#ffdede\"");
			}
			
			
			
			
			/*
			 * 라디오 버튼에서 selection-change 이벤트 발생 시 호출.
			 * 라디오버튼 아이템을 선택하여 선택된 값이 저장된 후에 발생하는 이벤트.
			 */
			function onRdb1SelectionChange(/* cpr.events.CSelectionEvent */ e){
				/** 
				 * @type cpr.controls.RadioButton
				 */
				var rdb1 = e.control;
				if(rdb1.value == "0"){
					app.lookup("grd1").autoRowHeight = "none";
				} else {
					app.lookup("grd1").autoRowHeight = "all";
				}
			};
			// End - User Script
			
			// Header
			var dataSet_1 = new cpr.data.DataSet("ds_weekRptList");
			dataSet_1.parseData({
				"columns" : [
					{"name": "wRpt_idx"},
					{"name": "wRpt_title"},
					{"name": "wRpt_content"},
					{"name": "wRpt_toRpt"},
					{
						"name": "wRpt_reporter",
						"dataType": "string"
					}
				]
			});
			app.register(dataSet_1);
			
			var dataSet_2 = new cpr.data.DataSet("ds_partTree");
			dataSet_2.parseData({
				"columns" : [
					{"name": "label"},
					{"name": "value"},
					{"name": "parent"}
				]
			});
			app.register(dataSet_2);
			var dataMap_1 = new cpr.data.DataMap("Result");
			dataMap_1.parseData({
				"columns" : [{"name": "ResultCode"}]
			});
			app.register(dataMap_1);
			
			var dataMap_2 = new cpr.data.DataMap("dm_page");
			dataMap_2.parseData({
				"columns" : [
					{"name": "offset"},
					{"name": "limit"}
				]
			});
			app.register(dataMap_2);
			
			var dataMap_3 = new cpr.data.DataMap("totalCount");
			dataMap_3.parseData({
				"columns" : [{"name": "Count"}]
			});
			app.register(dataMap_3);
			
			var dataMap_4 = new cpr.data.DataMap("dm_search");
			dataMap_4.parseData({
				"columns" : [
					{"name": "search_input"},
					{"name": "search_combo"}
				]
			});
			app.register(dataMap_4);
			
			var dataMap_5 = new cpr.data.DataMap("dm_part");
			dataMap_5.parseData({
				"columns" : [{"name": "part_value"}]
			});
			app.register(dataMap_5);
			var submission_1 = new cpr.protocols.Submission("sms_getWeekRptList");
			submission_1.method = "get";
			submission_1.action = "/bsmg/report/getWeekRptList";
			submission_1.addRequestData(dataMap_2);
			submission_1.addResponseData(dataSet_1, false);
			submission_1.addResponseData(dataMap_3, false);
			submission_1.addResponseData(dataMap_1, false);
			if(typeof onSms_getWeekRptListSubmitDone == "function") {
				submission_1.addEventListener("submit-done", onSms_getWeekRptListSubmitDone);
			}
			app.register(submission_1);
			
			var submission_2 = new cpr.protocols.Submission("sms_getCategoryList");
			submission_2.method = "get";
			submission_2.action = "/bsmg/setting/weekRptCategory";
			submission_2.addResponseData(dataMap_1, false);
			submission_2.addResponseData(dataSet_2, false);
			if(typeof onSms_getCategoryListSubmitDone == "function") {
				submission_2.addEventListener("submit-done", onSms_getCategoryListSubmitDone);
			}
			app.register(submission_2);
			
			var submission_3 = new cpr.protocols.Submission("sms_getWeekRptSearch");
			submission_3.async = true;
			submission_3.method = "get";
			submission_3.action = "/bsmg/report/getWeekRptSearch";
			submission_3.addRequestData(dataMap_4);
			submission_3.addRequestData(dataMap_2);
			submission_3.addResponseData(dataSet_1, false);
			submission_3.addResponseData(dataMap_3, false);
			submission_3.addResponseData(dataMap_1, false);
			if(typeof onSms_getWeekRptSearchSubmitDone == "function") {
				submission_3.addEventListener("submit-done", onSms_getWeekRptSearchSubmitDone);
			}
			app.register(submission_3);
			
			var submission_4 = new cpr.protocols.Submission("sms_getWeekRptCategory");
			submission_4.method = "get";
			submission_4.action = "/bsmg/report/getWeekRptCategory";
			submission_4.addRequestData(dataMap_5);
			submission_4.addRequestData(dataMap_2);
			submission_4.addResponseData(dataSet_1, false);
			submission_4.addResponseData(dataMap_3, false);
			submission_4.addResponseData(dataMap_1, false);
			if(typeof onSms_getWeekRptCategorySubmitDone == "function") {
				submission_4.addEventListener("submit-done", onSms_getWeekRptCategorySubmitDone);
			}
			app.register(submission_4);
			
			app.supportMedia("all and (min-width: 1024px)", "default");
			app.supportMedia("all and (min-width: 500px) and (max-width: 1023px)", "tablet");
			app.supportMedia("all and (max-width: 499px)", "mobile");
			
			// Configure root container
			var container = app.getContainer();
			container.style.css({
				"width" : "100%",
				"top" : "0px",
				"height" : "100%",
				"left" : "0px"
			});
			
			// Layout
			var xYLayout_1 = new cpr.controls.layouts.XYLayout();
			container.setLayout(xYLayout_1);
			
			// UI Configuration
			var group_1 = new cpr.controls.Container();
			group_1.style.css({
				"background-color" : "#fbfce9",
				"border-radius" : "15px",
				"border-right-style" : "solid",
				"border-left-style" : "solid",
				"border-bottom-style" : "solid",
				"border-top-style" : "solid"
			});
			// Layout
			var responsiveXYLayout_1 = new cpr.controls.layouts.ResponsiveXYLayout();
			group_1.setLayout(responsiveXYLayout_1);
			(function(container){
				var output_1 = new cpr.controls.Output();
				output_1.value = "주간 업무보고 리스트";
				output_1.style.css({
					"color" : "#000000",
					"font-weight" : "bold",
					"font-size" : "20px",
					"font-style" : "normal"
				});
				container.addChild(output_1, {
					positions: [
						{
							"media": "all and (min-width: 1024px)",
							"top": "5px",
							"left": "5px",
							"width": "217px",
							"height": "43px"
						}, 
						{
							"media": "all and (min-width: 500px) and (max-width: 1023px)",
							"hidden": false,
							"top": "5px",
							"left": "2px",
							"width": "106px",
							"height": "43px"
						}, 
						{
							"media": "all and (max-width: 499px)",
							"hidden": false,
							"top": "5px",
							"left": "2px",
							"width": "74px",
							"height": "43px"
						}
					]
				});
				var button_1 = new cpr.controls.Button();
				button_1.style.css({
					"border-radius" : "45px",
					"background-repeat" : "no-repeat",
					"background-position" : "center",
					"background-image" : "url('images/arrow-clockwise.svg')"
				});
				if(typeof onButtonClick == "function") {
					button_1.addEventListener("click", onButtonClick);
				}
				container.addChild(button_1, {
					positions: [
						{
							"media": "all and (min-width: 1024px)",
							"top": "83px",
							"left": "738px",
							"width": "41px",
							"height": "35px"
						}, 
						{
							"media": "all and (min-width: 500px) and (max-width: 1023px)",
							"hidden": false,
							"top": "83px",
							"left": "360px",
							"width": "20px",
							"height": "35px"
						}, 
						{
							"media": "all and (max-width: 499px)",
							"hidden": false,
							"top": "83px",
							"left": "252px",
							"width": "14px",
							"height": "35px"
						}
					]
				});
				var comboBox_1 = new cpr.controls.ComboBox("cmb1");
				comboBox_1.value = "0";
				comboBox_1.fixedListWidth = true;
				comboBox_1.placeholder = "전체";
				comboBox_1.preventInput = true;
				(function(comboBox_1){
					comboBox_1.addItem(new cpr.controls.Item("전체", "0"));
					comboBox_1.addItem(new cpr.controls.Item("제목", "1"));
					comboBox_1.addItem(new cpr.controls.Item("내용", "2"));
					comboBox_1.addItem(new cpr.controls.Item("보고자", "4"));
				})(comboBox_1);
				container.addChild(comboBox_1, {
					positions: [
						{
							"media": "all and (min-width: 1024px)",
							"top": "92px",
							"left": "221px",
							"width": "120px",
							"height": "28px"
						}, 
						{
							"media": "all and (min-width: 500px) and (max-width: 1023px)",
							"hidden": false,
							"top": "92px",
							"left": "108px",
							"width": "59px",
							"height": "28px"
						}, 
						{
							"media": "all and (max-width: 499px)",
							"hidden": false,
							"top": "92px",
							"left": "76px",
							"width": "41px",
							"height": "28px"
						}
					]
				});
				var searchInput_1 = new cpr.controls.SearchInput("ipb1");
				if(typeof onIpb1Search == "function") {
					searchInput_1.addEventListener("search", onIpb1Search);
				}
				container.addChild(searchInput_1, {
					positions: [
						{
							"media": "all and (min-width: 1024px)",
							"top": "92px",
							"left": "341px",
							"width": "361px",
							"height": "28px"
						}, 
						{
							"media": "all and (min-width: 500px) and (max-width: 1023px)",
							"hidden": false,
							"top": "92px",
							"left": "167px",
							"width": "176px",
							"height": "28px"
						}, 
						{
							"media": "all and (max-width: 499px)",
							"hidden": false,
							"top": "92px",
							"left": "117px",
							"width": "123px",
							"height": "28px"
						}
					]
				});
				var pageIndexer_1 = new cpr.controls.PageIndexer("pageIndex");
				pageIndexer_1.init(1, 1, 1);
				if(typeof onPageIndexSelectionChange == "function") {
					pageIndexer_1.addEventListener("selection-change", onPageIndexSelectionChange);
				}
				if(typeof onPageIndexBeforeSelectionChange == "function") {
					pageIndexer_1.addEventListener("before-selection-change", onPageIndexBeforeSelectionChange);
				}
				container.addChild(pageIndexer_1, {
					positions: [
						{
							"media": "all and (min-width: 1024px)",
							"top": "614px",
							"left": "269px",
							"width": "505px",
							"height": "37px"
						}, 
						{
							"media": "all and (min-width: 500px) and (max-width: 1023px)",
							"hidden": false,
							"top": "614px",
							"left": "131px",
							"width": "247px",
							"height": "37px"
						}, 
						{
							"media": "all and (max-width: 499px)",
							"hidden": false,
							"top": "614px",
							"left": "92px",
							"width": "173px",
							"height": "37px"
						}
					]
				});
				var grid_1 = new cpr.controls.Grid("grd1");
				grid_1.readOnly = true;
				grid_1.init({
					"dataSet": app.lookup("ds_weekRptList"),
					"autoRowHeight": "none",
					"wheelRowCount": 1,
					"resizableColumns": "all",
					"columns": [
						{"width": "34px"},
						{"width": "154px"},
						{"width": "371px"},
						{"width": "44px"},
						{"width": "46px"}
					],
					"header": {
						"rows": [{"height": "24px"}],
						"cells": [
							{
								"constraint": {"rowIndex": 0, "colIndex": 0},
								"configurator": function(cell){
									cell.filterable = false;
									cell.sortable = false;
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 1},
								"configurator": function(cell){
									cell.targetColumnName = "wRpt_title";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "제목";
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 2},
								"configurator": function(cell){
									cell.targetColumnName = "wRpt_content";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "내용";
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 3},
								"configurator": function(cell){
									cell.targetColumnName = "wRpt_toRpt";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "보고대상";
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 4},
								"configurator": function(cell){
									cell.targetColumnName = "wRpt_reporter";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "보고자";
								}
							}
						]
					},
					"detail": {
						"rows": [{"height": "146px"}],
						"cells": [
							{
								"constraint": {"rowIndex": 0, "colIndex": 0},
								"configurator": function(cell){
									cell.columnType = "rowindex";
									cell.style.css({
										"background-color" : "transparent",
										"color" : "#000000"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 1},
								"configurator": function(cell){
									cell.columnName = "wRpt_title";
									cell.style.css({
										"background-color" : "transparent",
										"color" : "#000000"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 2},
								"configurator": function(cell){
									cell.columnName = "wRpt_content";
									cell.style.css({
										"background-color" : "transparent",
										"color" : "#000000"
									});
									cell.control = (function(){
										var textArea_1 = new cpr.controls.TextArea("txa1");
										textArea_1.style.css({
											"padding-left" : "5px"
										});
										textArea_1.bind("value").toDataColumn("wRpt_content");
										return textArea_1;
									})();
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 3},
								"configurator": function(cell){
									cell.columnName = "wRpt_toRpt";
									cell.style.css({
										"background-color" : "transparent",
										"color" : "#000000"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 4},
								"configurator": function(cell){
									cell.columnName = "wRpt_reporter";
									cell.style.css({
										"background-color" : "transparent",
										"color" : "#000000"
									});
								}
							}
						]
					}
				});
				if(typeof onGrd1RowDblclick == "function") {
					grid_1.addEventListener("row-dblclick", onGrd1RowDblclick);
				}
				container.addChild(grid_1, {
					positions: [
						{
							"media": "all and (min-width: 1024px)",
							"top": "128px",
							"left": "20px",
							"width": "984px",
							"height": "487px"
						}, 
						{
							"media": "all and (min-width: 500px) and (max-width: 1023px)",
							"hidden": false,
							"top": "128px",
							"left": "10px",
							"width": "480px",
							"height": "487px"
						}, 
						{
							"media": "all and (max-width: 499px)",
							"hidden": false,
							"top": "128px",
							"left": "7px",
							"width": "336px",
							"height": "487px"
						}
					]
				});
				var radioButton_1 = new cpr.controls.RadioButton("rdb1");
				radioButton_1.value = "0";
				(function(radioButton_1){
					radioButton_1.addItem(new cpr.controls.Item("작게", "0"));
					radioButton_1.addItem(new cpr.controls.Item("크게", "1"));
				})(radioButton_1);
				if(typeof onRdb1SelectionChange == "function") {
					radioButton_1.addEventListener("selection-change", onRdb1SelectionChange);
				}
				container.addChild(radioButton_1, {
					positions: [
						{
							"media": "all and (min-width: 1024px)",
							"top": "81px",
							"left": "20px",
							"width": "152px",
							"height": "40px"
						}, 
						{
							"media": "all and (min-width: 500px) and (max-width: 1023px)",
							"hidden": false,
							"top": "81px",
							"left": "10px",
							"width": "74px",
							"height": "40px"
						}, 
						{
							"media": "all and (max-width: 499px)",
							"hidden": false,
							"top": "81px",
							"left": "7px",
							"width": "52px",
							"height": "40px"
						}
					]
				});
			})(group_1);
			container.addChild(group_1, {
				"top": "4px",
				"left": "364px",
				"width": "1024px",
				"height": "668px"
			});
			
			var output_2 = new cpr.controls.Output();
			output_2.value = "주간 업무보고 카테고리";
			output_2.style.css({
				"font-weight" : "bold"
			});
			container.addChild(output_2, {
				"top": "16px",
				"left": "10px",
				"width": "259px",
				"height": "45px"
			});
			
			var tree_1 = new cpr.controls.Tree("tre1");
			tree_1.style.css({
				"border-right-style" : "solid",
				"border-radius" : "15px",
				"background-color" : "#fbfce9",
				"border-left-style" : "solid",
				"border-bottom-style" : "solid",
				"border-top-style" : "solid"
			});
			(function(tree_1){
				tree_1.setItemSet(app.lookup("ds_partTree"), {
					"label": "label",
					"value": "value",
					"parentValue": "parent"
				});
			})(tree_1);
			if(typeof onTre1ItemDblclick == "function") {
				tree_1.addEventListener("item-dblclick", onTre1ItemDblclick);
			}
			container.addChild(tree_1, {
				"top": "60px",
				"left": "10px",
				"width": "344px",
				"height": "606px"
			});
			if(typeof onBodyLoad == "function"){
				app.addEventListener("load", onBodyLoad);
			}
		}
	});
	app.title = "bm_weekRptList";
	cpr.core.Platform.INSTANCE.register(app);
})();
