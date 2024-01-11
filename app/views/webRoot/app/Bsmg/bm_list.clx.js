/*
 * App URI: app/Bsmg/bm_list
 * Source Location: app/Bsmg/bm_list.clx
 *
 * This file was generated by eXbuilder6 compiler, Don't edit manually.
 */
(function(){
	var app = new cpr.core.App("app/Bsmg/bm_list", {
		onPrepare: function(loader){
		},
		onCreate: function(/* cpr.core.AppInstance */ app, exports){
			var linker = {};
			// Start - User Script
			/************************************************
			 * bm_list.js
			 * Created at 2022. 5. 12. 오전 10:13:38.
			 *
			 * @author SW2Team
			 ************************************************/
			
			/* 페이지 처리 */
			var RowCount = 6; // 한 페이지 행 개수
			var SearchFlag = false; // 검색해서 받은 리스트
			var AttrFlag = false; // 업무속성으로 검색
			var NewSearch = false; // 현재 페이지가 1이 아닐때, 새로 search해서 리스트를 불러올 시 1페이지로 돌아가게.
			
			var dataManager = cpr.core.Module.require("lib/DataManager"); // 싱글톤패턴
			
			/*
			 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
			 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
			 */
			function onBodyLoad(/* cpr.events.CEvent */ e){
				dataManager = getDataManager();
				
				app.lookup("sms_setTree").send();
				app.lookup("sms_getAttr1").send();
				setPaging(0, 1, RowCount, 5);
				sendRptListRequest();
			}
			
			function setPaging(totCnt,pageIdx ,RowCount, pageSize){
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
				var sms_getRptList = app.lookup("sms_getRptList");
				
				// dm_page에 setValue로 집어넣으면 서버에서 에러남... why? 뭐가 다르지?
				sms_getRptList.setParameters("offset", offset);
				sms_getRptList.setParameters("limit", RowCount);
				sms_getRptList.send();
			}
			
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_getRptListSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_getRptList = e.control;
				var pageIndexer = app.lookup("pageIndex");  
				var result = app.lookup("Result").getString("ResultCode");	
			
				if(result == 0){
					var totalCount = app.lookup("totalCount").getValue("Count");
			
					AllGridColorWhite();
					
					setPaging(Number(totalCount), pageIndexer.currentPageIndex, RowCount, 5);
					SearchFlag = false;
					AttrFlag = false;
					return;
				} else {
					alert(getErrorString(result));
				}
			}
			
			
			/*
			 * "일일 업무보고 작성" 버튼에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onButtonClick(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var button = e.control;
				app.getRootAppInstance().openDialog("app/Bsmg/bm_dailyRpt", {
					width : 1000, height : 800
				}, function(dialog){
					dialog.ready(function(dialogApp){
						dialog.modal = true;
						dialog.headerVisible = true;
						dialog.headerMovable = true;
						dialog.resizable = true;
						dialog.headerTitle = "일일 업무보고 등록";
					});
				}).then(function(returnValue){
					if (returnValue == 1){
						if(AttrFlag){
							sendAttrRptRequest();
						} else if(SearchFlag){
							sendSearchRequest();
						} else {
							sendRptListRequest();
						}
					}
				})
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
				var selectionEvent = new cpr.events.CSelectionEvent("before-pagechange", {
					oldSelection: e.oldSelection,
					newSelection: e.newSelection
				});
				app.dispatchEvent(selectionEvent);
				// 기본처리가 중단되었을 때 변경을 취소함.
				if(selectionEvent.defaultPrevented == true) {
					e.preventDefault();
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
				var selectionEvent = new cpr.events.CSelectionEvent("pagechange",{
					oldSelection: e.oldSelection,
					newSelection: e.newSelection
				});
				app.dispatchEvent(selectionEvent);
			//	console.log("선택 페이지 : " + pageIndex.currentPageIndex);
				if(SearchFlag){
					sendSearchRequest();
				} else if(AttrFlag){
					sendAttrRptRequest();
				} else{
					sendRptListRequest();
				}
			}
			
			
			function sendSearchRequest(){
				var pageIndexer = app.lookup("pageIndex");
				if(!SearchFlag){	// 카테고리에서 3페이지를 보다가 검색하는 경우, 시작이 3페이지이므로 1페이지로 초기화
					NewSearch = true;
					pageIndexer.currentPageIndex = 1;
				} else { // 이미 검색을 한 상태. 다음 페이지를 보려할 때
					NewSearch = false;
				}
				
				var pageIdx = pageIndexer.currentPageIndex;
				var offset = (pageIdx - 1) * RowCount;
				var combo = app.lookup("cmb1").value;
				var input = app.lookup("ipb1").value;	
				
				app.lookup("ds_rptList").clear();
				app.lookup("dm_search").setValue("search_combo", combo);
				app.lookup("dm_search").setValue("search_input", input);
				app.lookup("sms_getRptSearch").setParameters("offset", offset);
				app.lookup("sms_getRptSearch").setParameters("limit", RowCount);
				app.lookup("sms_getRptSearch").send();
			}
			
			/*
			 * 서치 인풋에서 search 이벤트 발생 시 호출.
			 * Searchinput의 enter키 또는 검색버튼을 클릭하여 인풋의 값이 Search될때 발생하는 이벤트
			 */
			function onSearchInputSearch(/* cpr.events.CUIEvent */ e){
				/** 
				 * @type cpr.controls.SearchInput
				 */
				var searchInput = e.control;
				sendSearchRequest();
			}
			
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_getRptSearchSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_getRptSearch = e.control;
				var result = app.lookup("Result").getString("ResultCode");
				var totalCount = app.lookup("totalCount").getValue("Count");
				var pageIndexer = app.lookup("pageIndex");
				
				if(result == 0){
					SearchFlag = true;
					AttrFlag = false;
					
					if(totalCount == 0){
						alert('자료가 없습니다.');
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
				} else {
					alert(getErrorString(result));
				}
			}
			
			
			/*
			 * '새로고침' 버튼에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onButtonClick2(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var button = e.control;
				var pageIndexer = app.lookup("pageIndex");
				pageIndexer.currentPageIndex = 1; 
				sendRptListRequest();
			}
			
			function gridColorWhite(cellIndex){
				var grd = app.lookup("grd1");
				var voCell = grd.detail.getColumn(cellIndex);
			//	var voExpress = "switch(getValue(\"color\")){\n\tcase \"Red\" : \"red\"\n\tdefault : \"transparent\"\n}";
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
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_setTreeSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_setTree = e.control;
				var result = app.lookup("Result").getString("ResultCode");
				if(result == 0){
			//		console.log(app.lookup("ds_List").getRowDataRanged());
			
					var dsAttrTree = app.lookup("ds_List");
					dataManager.setDsAttrTree(dsAttrTree);
			
					app.lookup("tre1").redraw();
				} else {
					alert(getErrorString(result));
				}
			}
			
			
			function sendAttrRptRequest(){
				var pageIndexer = app.lookup("pageIndex");
				if(!AttrFlag){
					NewSearch = true;
					pageIndexer.currentPageIndex = 1;
				} else {
					NewSearch = false;
				}
				
				var pageIdx = pageIndexer.currentPageIndex;
				var offset = (pageIdx -1) * RowCount;
				
				app.lookup("ds_rptList").clear();
				app.lookup("sms_rptAttrSearch").setParameters("offset", offset);
				app.lookup("sms_rptAttrSearch").setParameters("limit", RowCount);
				app.lookup("sms_rptAttrSearch").send();
			}
			
			
			
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_rptAttrSearchSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_rptAttrSearch = e.control;
				var result = app.lookup("Result").getString("ResultCode");
				var totalCount = app.lookup("totalCount").getValue("Count");
				var pageIndexer = app.lookup("pageIndex");
				
				
				if(result == 0) {
					SearchFlag = false;
					AttrFlag = true;
					
					if(totalCount == 0){
						alert("자료가 없습니다.");
						return;
					}
					
					AllGridColorWhite();
					gridColorImpact(4);
					
					if(NewSearch){  // 새로 검색한 거면 1페이지부터
						setPaging(Number(totalCount), 1, RowCount, 5);
					} else {
						setPaging(Number(totalCount), pageIndexer.currentPageIndex, RowCount, 5);
					}
					app.getContainer().redraw();
					return;
					
					// 이제 서버단 짜기만 하면 됨
				} else {
					alert(getErrorString(result));
				}
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
				var pageIndexer = app.lookup("pageIndex");
				pageIndexer.currentPageIndex = 1;
				
				var attrValue = tre1.getSelectionLast().value;
				var attrCategory = 999;
			//	console.log("파싱전 : "+ tre1.getSelectionLast().value);
				attrValue = attrValue.split("-")
				if(attrValue[1] == null){
					attrValue = attrValue[0];
					attrCategory = 0; // 업무속성1로 검색
				} else {
					attrValue = attrValue[1];
					attrCategory = 1; // 업무속성2로 검색
				}
			//	console.log("파싱 후 : " +attrValue);
				app.lookup("dm_attr").setValue("attrValue", attrValue);
				app.lookup("dm_attr").setValue("attrCategory", attrCategory);
				
				sendAttrRptRequest();
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
				var paramValue = row.getValue("rpt_idx");
				
			//	console.log("선택한 셀 확인 : " + paramValue);
				app.getRootAppInstance().openDialog("app/Bsmg/bm_rptView",{
					width : 1000,
					height : 800,
				}, function(dialog){
					dialog.modal = true;
					dialog.headerVisible = true;
					dialog.headerMovable = true;
					dialog.headerTitle = "일일 업무보고 확인";
					dialog.headerClose = true;
					dialog.addEventListener("keyup", function(e){
						if(e.keyCode == 27){ // ESC
							dialog.close();
						}
					});
					dialog.initValue = {
						rpt_idx : paramValue
					};
				}).then(function(returnValue){
					if (returnValue == 1){
						if(AttrFlag){
							sendAttrRptRequest();
						} else if(SearchFlag){
							sendSearchRequest();
						} else {
							sendRptListRequest();
						}
					}
				});
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
			//	app.lookup("grd1").redraw(); 속성이 변경된 경우 자동으로 그려진다.
			}
			
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_getAttr1SubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_getAttr1 = e.control;
				var result = app.lookup("Result").getString("ResultCode");
				if(result == 0){
			//		console.log(app.lookup("ds_List").getRowDataRanged());
			
					var dsAttr1 = app.lookup("ds_attr1"); // 업무 속성 1 : 카테고리
			//		dataManager.setDsAttrTree(dsAttrTree);
			
				} else {
					alert(getErrorString(result));
				}
			};
			// End - User Script
			
			// Header
			var dataSet_1 = new cpr.data.DataSet("ds_rptList");
			dataSet_1.parseData({
				"columns" : [
					{"name": "rpt_idx"},
					{"name": "rpt_title"},
					{"name": "rpt_content"},
					{"name": "rpt_reporter"},
					{"name": "rpt_date"},
					{"name": "rpt_attr1"}
				]
			});
			app.register(dataSet_1);
			
			var dataSet_2 = new cpr.data.DataSet("ds_List");
			dataSet_2.parseData({
				"columns" : [
					{"name": "label"},
					{"name": "value"},
					{"name": "parent"}
				]
			});
			app.register(dataSet_2);
			
			var dataSet_3 = new cpr.data.DataSet("ds_attr1");
			dataSet_3.parseData({
				"columns" : [
					{
						"name": "attr1_idx",
						"dataType": "string"
					},
					{
						"name": "attr1_category",
						"dataType": "string"
					}
				]
			});
			app.register(dataSet_3);
			var dataMap_1 = new cpr.data.DataMap("Result");
			dataMap_1.parseData({
				"columns" : [{"name": "ResultCode"}]
			});
			app.register(dataMap_1);
			
			var dataMap_2 = new cpr.data.DataMap("totalCount");
			dataMap_2.parseData({
				"columns" : [{"name": "Count"}]
			});
			app.register(dataMap_2);
			
			var dataMap_3 = new cpr.data.DataMap("dm_page");
			dataMap_3.parseData({
				"columns" : [
					{"name": "offset"},
					{"name": "limit"}
				]
			});
			app.register(dataMap_3);
			
			var dataMap_4 = new cpr.data.DataMap("dm_search");
			dataMap_4.parseData({
				"columns" : [
					{"name": "search_combo"},
					{"name": "search_input"}
				]
			});
			app.register(dataMap_4);
			
			var dataMap_5 = new cpr.data.DataMap("dm_attr");
			dataMap_5.parseData({
				"columns" : [
					{"name": "attrValue"},
					{"name": "attrCategory"}
				]
			});
			app.register(dataMap_5);
			var submission_1 = new cpr.protocols.Submission("sms_getRptList");
			submission_1.method = "get";
			submission_1.action = "/bsmg/report/reportList";
			submission_1.addRequestData(dataMap_3);
			submission_1.addResponseData(dataSet_1, false);
			submission_1.addResponseData(dataMap_1, false);
			submission_1.addResponseData(dataMap_2, false);
			if(typeof onSms_getRptListSubmitDone == "function") {
				submission_1.addEventListener("submit-done", onSms_getRptListSubmitDone);
			}
			app.register(submission_1);
			
			var submission_2 = new cpr.protocols.Submission("sms_getRptSearch");
			submission_2.method = "get";
			submission_2.action = "/bsmg/report/reportSearch";
			submission_2.addRequestData(dataMap_4);
			submission_2.addRequestData(dataMap_3);
			submission_2.addResponseData(dataSet_1, false);
			submission_2.addResponseData(dataMap_1, false);
			submission_2.addResponseData(dataMap_2, false);
			if(typeof onSms_getRptSearchSubmitDone == "function") {
				submission_2.addEventListener("submit-done", onSms_getRptSearchSubmitDone);
			}
			app.register(submission_2);
			
			var submission_3 = new cpr.protocols.Submission("sms_setTree");
			submission_3.method = "get";
			submission_3.action = "/bsmg/setting/attrTree";
			submission_3.addResponseData(dataSet_2, false);
			submission_3.addResponseData(dataMap_1, false);
			if(typeof onSms_setTreeSubmitDone == "function") {
				submission_3.addEventListener("submit-done", onSms_setTreeSubmitDone);
			}
			app.register(submission_3);
			
			var submission_4 = new cpr.protocols.Submission("sms_rptAttrSearch");
			submission_4.method = "get";
			submission_4.action = "/bsmg/report/reportAttrSearch";
			submission_4.addRequestData(dataMap_5);
			submission_4.addRequestData(dataMap_3);
			submission_4.addResponseData(dataSet_1, false);
			submission_4.addResponseData(dataMap_1, false);
			submission_4.addResponseData(dataMap_2, false);
			if(typeof onSms_rptAttrSearchSubmitDone == "function") {
				submission_4.addEventListener("submit-done", onSms_rptAttrSearchSubmitDone);
			}
			app.register(submission_4);
			
			var submission_5 = new cpr.protocols.Submission("sms_getAttr1");
			submission_5.method = "get";
			submission_5.action = "/bsmg/setting/attr1";
			submission_5.addResponseData(dataSet_3, false);
			submission_5.addResponseData(dataMap_1, false);
			if(typeof onSms_getAttr1SubmitDone == "function") {
				submission_5.addEventListener("submit-done", onSms_getAttr1SubmitDone);
			}
			app.register(submission_5);
			
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
				"background-color" : "#f9fbf0",
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
				output_1.value = "일일 업무보고 리스트";
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
				if(typeof onButtonClick2 == "function") {
					button_1.addEventListener("click", onButtonClick2);
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
				var grid_1 = new cpr.controls.Grid("grd1");
				grid_1.readOnly = true;
				grid_1.init({
					"dataSet": app.lookup("ds_rptList"),
					"autoRowHeight": "none",
					"wheelRowCount": 1,
					"resizableColumns": "all",
					"columns": [
						{"width": "37px"},
						{"width": "168px"},
						{"width": "269px"},
						{"width": "51px"},
						{"width": "40px"},
						{"width": "100px"}
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
									cell.targetColumnName = "rpt_title";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "제목";
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 2},
								"configurator": function(cell){
									cell.targetColumnName = "rpt_content";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "내용";
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 3},
								"configurator": function(cell){
									cell.targetColumnName = "rpt_reporter";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "보고자";
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 4},
								"configurator": function(cell){
									cell.targetColumnName = "rpt_attr1";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "업무 속성";
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 5},
								"configurator": function(cell){
									cell.targetColumnName = "rpt_date";
									cell.visible = false;
								}
							}
						]
					},
					"detail": {
						"rows": [{"height": "64px"}],
						"cells": [
							{
								"constraint": {"rowIndex": 0, "colIndex": 0},
								"configurator": function(cell){
									cell.columnType = "rowindex";
									cell.style.css({
										"background-color" : "transparent",
										"color" : "#000000",
										"font-weight" : "normal"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 1},
								"configurator": function(cell){
									cell.columnName = "rpt_title";
									cell.style.css({
										"background-color" : "transparent",
										"color" : "#000000",
										"font-weight" : "normal"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 2},
								"configurator": function(cell){
									cell.columnName = "rpt_content";
									cell.style.css({
										"background-color" : "transparent",
										"color" : "#000000",
										"font-weight" : "normal",
										"padding-left" : "5px",
										"text-align" : "left"
									});
									cell.control = (function(){
										var textArea_1 = new cpr.controls.TextArea("txa1");
										textArea_1.style.css({
											"font-weight" : "normal"
										});
										textArea_1.bind("value").toDataColumn("rpt_content");
										return textArea_1;
									})();
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 3},
								"configurator": function(cell){
									cell.columnName = "rpt_reporter";
									cell.style.css({
										"background-color" : "transparent",
										"color" : "#000000",
										"font-weight" : "normal"
									});
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 4},
								"configurator": function(cell){
									cell.columnName = "rpt_attr1";
									cell.style.css({
										"background-color" : "transparent",
										"color" : "#000000",
										"font-weight" : "normal"
									});
									cell.control = (function(){
										var comboBox_1 = new cpr.controls.ComboBox("cmb_attr1");
										comboBox_1.readOnly = true;
										(function(comboBox_1){
											comboBox_1.setItemSet(app.lookup("ds_attr1"), {
												"label": "attr1_category",
												"value": "attr1_idx"
											})
										})(comboBox_1);
										comboBox_1.bind("value").toDataColumn("rpt_attr1");
										return comboBox_1;
									})();
								}
							},
							{
								"constraint": {"rowIndex": 0, "colIndex": 5},
								"configurator": function(cell){
									cell.columnName = "rpt_date";
								}
							}
						]
					}
				});
				grid_1.style.css({
					"font-weight" : "normal"
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
				var button_2 = new cpr.controls.Button();
				button_2.value = "일일 업무보고 작성";
				if(typeof onButtonClick == "function") {
					button_2.addEventListener("click", onButtonClick);
				}
				container.addChild(button_2, {
					positions: [
						{
							"media": "all and (min-width: 1024px)",
							"top": "81px",
							"left": "825px",
							"width": "141px",
							"height": "39px"
						}, 
						{
							"media": "all and (min-width: 500px) and (max-width: 1023px)",
							"hidden": false,
							"top": "81px",
							"left": "403px",
							"width": "69px",
							"height": "39px"
						}, 
						{
							"media": "all and (max-width: 499px)",
							"hidden": false,
							"top": "81px",
							"left": "282px",
							"width": "48px",
							"height": "39px"
						}
					]
				});
				var comboBox_2 = new cpr.controls.ComboBox("cmb1");
				comboBox_2.value = "0";
				comboBox_2.fixedListWidth = true;
				comboBox_2.placeholder = "전체";
				comboBox_2.preventInput = true;
				(function(comboBox_2){
					comboBox_2.addItem(new cpr.controls.Item("전체", "0"));
					comboBox_2.addItem(new cpr.controls.Item("제목", "1"));
					comboBox_2.addItem(new cpr.controls.Item("내용", "2"));
					comboBox_2.addItem(new cpr.controls.Item("보고자", "3"));
				})(comboBox_2);
				container.addChild(comboBox_2, {
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
				if(typeof onSearchInputSearch == "function") {
					searchInput_1.addEventListener("search", onSearchInputSearch);
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
			output_2.value = "일일 업무보고 카테고리";
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
				"background-size" : "cover",
				"border-bottom-color" : "#000000",
				"border-left-color" : "#000000",
				"border-right-color" : "#000000",
				"border-top-style" : "solid",
				"border-radius" : "15px",
				"background-color" : "#f9fbf0",
				"background-repeat" : "no-repeat",
				"border-left-style" : "solid",
				"border-top-color" : "#000000",
				"border-bottom-style" : "solid",
				"background-image" : "none"
			});
			tree_1.style.item.css({
				"background-repeat" : "no-repeat",
				"background-size" : "auto",
				"background-image" : "none",
				"background-position" : "top left",
				"background-origin" : "content-box"
			});
			(function(tree_1){
				tree_1.setItemSet(app.lookup("ds_List"), {
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
	app.title = "bm_list";
	cpr.core.Platform.INSTANCE.register(app);
})();