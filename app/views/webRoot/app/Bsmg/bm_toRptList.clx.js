/*
 * App URI: app/Bsmg/bm_toRptList
 * Source Location: app/Bsmg/bm_toRptList.clx
 *
 * This file was generated by eXbuilder6 compiler, Don't edit manually.
 */
(function(){
	var app = new cpr.core.App("app/Bsmg/bm_toRptList", {
		onPrepare: function(loader){
		},
		onCreate: function(/* cpr.core.AppInstance */ app, exports){
			var linker = {};
			// Start - User Script
			/************************************************
			 * bm_toRptList.js
			 * Created at 2022. 5. 19. 오전 10:49:11.
			 *
			 * @author SW2Team
			 ************************************************/
			
			var selectUserIDMap; // 선택된 유저객체 map
			var dataManager = cpr.core.Module.require("lib/DataManager");
			
			/*
			 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
			 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
			 */
			function onBodyLoad(/* cpr.events.CEvent */ e){
				// 보고대상, 참조대상에 다시 들어왔을 때 기존 선택지 유지하기 위해 dm 받아와서 저장
				
				app.lookup("sms_getUserList").send();
				app.lookup("sms_chkLogin").send();
				
				selectUserIDMap = new Map();
				
				dataManager = getDataManager();
				
				var dsRank = dataManager.getRankList();
				var dsRankList = app.lookup("ds_rank");
				dsRank.copyToDataSet(dsRankList);
				
				var dsPart = dataManager.getPartList();
				var dsPartList = app.lookup("ds_part");
				dsPart.copyToDataSet(dsPartList);
				
				app.getContainer().redraw();
				
				
			}
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_getUserListSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_getUserList = e.control;
				var result = app.lookup("Result").getString("ResultCode");
				var grd = app.lookup("userList");
				
				var initValue = app.getHost().initValue;
				var srcUserList = app.lookup("Src_memberList");	
				var toRptUserList = app.lookup("toRpt_memberList");
				var refUserList = app.lookup("ref_memberList");
				
				if(result == 0){
					// 이미 선택한 정보가 있을 경우
					if((srcUserList.getRowCount() != initValue.srcUserList.getRowCount()) && initValue.toRptUserList.getRowCount() != 0){
						srcUserList.clear();
			//			toRptUserList.clear();
			//			refUserList.clear();
						initValue.srcUserList.copyToDataSet(srcUserList);
						initValue.toRptUserList.copyToDataSet(toRptUserList);
						initValue.refUserList.copyToDataSet(refUserList);
					}
					grd.sort("mem_part ASC");
					grd.redraw();
				} else {
					alert(getErrorString(result));
				}
				 
			}
			
			
			/*
			 * ▷버튼에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onButtonClick(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var button = e.control;
				
				var grdUserList = app.lookup("userList");
				var dsSrcUserList = app.lookup("Src_memberList");
				var grdToRptUserList = app.lookup("toRptUserList");
				var dsDesUserList = app.lookup("toRpt_memberList");
				
				if(dsDesUserList.getRowCount() != 0){ // 이미 1명을 넣어놓은 상태.
					alert("보고자는 한명만 가능합니다.");
					return;
				}
				
				var indices = grdUserList.getCheckRowIndices();
				
				if (indices.length == 0) {
					return ;
				}
				
				// 선택한 user들에 대한 map을 만든다. (src 그리드에서 삭제를 위한 준비)
			//	indices.forEach(function(index){
				var row = grdUserList.getRow(indices[0]);
				var mem_id = row.getValue("mem_id");
				if(selectUserIDMap.get(mem_id) == undefined) { 
					dsDesUserList.pushRowData(row.getRowData());
					selectUserIDMap.set(mem_id, 1);
				} 
			//	});
			
				grdUserList.deleteRow(indices[0], false);
				grdUserList.showDeletedRow = false;
				
				// 선택해서 넘긴건 원본ds에서 삭제처리하고 그리드 redraw
				var total = dsSrcUserList.getRowCount();
				for (var i=0; i<total; i++){
					var row = dsSrcUserList.getRow(i);
					
					if(row){
						var mem_id = row.getValue("mem_id");
						if(selectUserIDMap.get(mem_id) != undefined){
							dsSrcUserList.realDeleteRow(i);
						} else {
							dsSrcUserList.setRowState(i, cpr.data.tabledata.RowState.UNCHANGED);
						}
					}
				}
			//	console.log(dsSrcUserList.getRowDatasByState(cpr.data.tabledata.RowState.UNCHANGED));
			//	console.log(dsDesUserList.getRowDatasByState(cpr.data.tabledata.RowState.INSERTED));
				grdUserList.redraw();
				grdToRptUserList.redraw();
				
			}
			
			
			
			/*
			 * 버튼(btnUserRemove)에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onBtnUserRemoveClick(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var btnUserRemove = e.control;
				
				var dsSrcUserList = app.lookup("Src_memberList");
				var grdSrcUserList = app.lookup("userList");
				var dsDesUserList = app.lookup("toRpt_memberList");
				var grdDesUserList = app.lookup("toRptUserList");
				
				var indices = grdDesUserList.getCheckRowIndices();
				if (indices.length == 0){
					return;
				}
				
				var idList = [];
				indices.forEach(function(index){
					var row = grdDesUserList.getRow(index);
					var mem_id = row.getValue("mem_id");
					
					selectUserIDMap.delete(mem_id);
					idList.push(mem_id);
				});
				
				idList.forEach(function(mem_id){
			//		var delRow = dsDesUserList.findFirstRowBoundlessly("mem_id=="+mem_id);
					var delRow = dsDesUserList.findFirstRow("mem_id==\'"+mem_id+"\'");
					dsSrcUserList.addRowData(delRow.getRowData());
					dsDesUserList.realDeleteRow(delRow.getIndex());
				});
				
				
				/*
				var total = dsSrcUserList.getRowCount();
				// 원본 userList에서 해당하는 row 복구
				for(var i=0; i<total; i++){
					var row = dsSrcUserList.getRow(i);
					if(row){
						var mem_id = row.getValue("mem_id");
						if(selectUserIDMap.get(mem_id) != undefined){
							dsSrcUserList.setRowState(i, cpr.data.tabledata.RowState.DELETED);
						} else {
							dsSrcUserList.setRowState(i, cpr.data.tabledata.RowState.UNCHANGED);
						}
					}
				}
				* 	*/
				grdSrcUserList.redraw();
				grdDesUserList.redraw();
			}
			
			
			/*
			 * 버튼에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onButtonClick2(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var button = e.control;
				
				var grdUserList = app.lookup("userList");
				var dsSrcUserList = app.lookup("Src_memberList");
				var grdRefUserList = app.lookup("refUserList");
				var dsDesUserList = app.lookup("ref_memberList");
				
				var indices = grdUserList.getCheckRowIndices();
				
				if (indices.length == 0) {
					return ;
				}
				
				// 선택한 user들에 대한 map을 만든다. (src 그리드에서 삭제를 위한 준비)
				indices.forEach(function(index){
					var row = grdUserList.getRow(index);
					var mem_id = row.getValue("mem_id");
					if(selectUserIDMap.get(mem_id) == undefined) { 
						dsDesUserList.pushRowData(row.getRowData());
						selectUserIDMap.set(mem_id, 1);
					} 
				});
				grdUserList.deleteRow(indices, false);
				grdUserList.showDeletedRow = false;
			
			//	dsDesUserList.setSort("mem_id");
				
				// 선택해서 넘긴건 원본ds에서 삭제처리하고 그리드 redraw
				var total = dsSrcUserList.getRowCount();
				for (var i=0; i<total; i++){
					var row = dsSrcUserList.getRow(i);
					
					if(row){
						var mem_id = row.getValue("mem_id");
						if(selectUserIDMap.get(mem_id) != undefined){
							dsSrcUserList.realDeleteRow(i);
						} else {
							dsSrcUserList.setRowState(i, cpr.data.tabledata.RowState.UNCHANGED);
						}
					}
				}
			//	console.log(dsSrcUserList.getRowDatasByState(cpr.data.tabledata.RowState.UNCHANGED));
			//	console.log(dsDesUserList.getRowDatasByState(cpr.data.tabledata.RowState.INSERTED));
				grdUserList.redraw();
				grdRefUserList.redraw();
				
			}
			
			
			/*
			 * 버튼에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onButtonClick3(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var btnUserRemove = e.control;
				
				var dsSrcUserList = app.lookup("Src_memberList");
				var grdSrcUserList = app.lookup("userList");
				var dsDesUserList = app.lookup("ref_memberList");
				var grdDesUserList = app.lookup("refUserList");
				
				var indices = grdDesUserList.getCheckRowIndices();
				if (indices.length == 0){
					return;
				}
				
				var idList = [];
				indices.forEach(function(index){
					var row = grdDesUserList.getRow(index);
					var mem_id = row.getValue("mem_id");
					
					selectUserIDMap.delete(mem_id);
					idList.push(mem_id);
				});
				
				idList.forEach(function(mem_id){
			//		var delRow = dsDesUserList.findFirstRowBoundlessly("mem_id=="+mem_id);
					var delRow = dsDesUserList.findFirstRow("mem_id==\'"+mem_id+"\'");
					dsSrcUserList.addRowData(delRow.getRowData());
					dsDesUserList.realDeleteRow(delRow.getIndex());
				});
				
				grdSrcUserList.redraw();
				grdDesUserList.redraw();
			}
			
			
			
			
			
			
			/*
			 * 서브미션에서 submit-done 이벤트 발생 시 호출.
			 * 응답처리가 모두 종료되면 발생합니다.
			 */
			function onSms_getUserListSearchSubmitDone(/* cpr.events.CSubmissionEvent */ e){
				/** 
				 * @type cpr.protocols.Submission
				 */
				var sms_getUserListSearch = e.control;
				var result = app.lookup("Result").getString("ResultCode");
				var ds = app.lookup("Src_memberList");
				if(result == 0){
					app.lookup("userList").redraw();
			//		console.log(ds.getRowDatasByState(cpr.data.tabledata.RowState.UNCHANGED));
			//		console.log(ds.getRowDatasByState(cpr.data.tabledata.RowState.INSERTED));
				} else {
					alert(getErrorString(result));
				}
			}
			
			/*
			 * 서치 인풋에서 keydown 이벤트 발생 시 호출. 검색버튼 눌렀을 때 send되게는 안되니?
			 * 사용자가 키를 누를 때 발생하는 이벤트.
			 */
			//function onIpb1Keydown(/* cpr.events.CKeyboardEvent */ e){
			//	/** 
			//	 * @type cpr.controls.SearchInput
			//	 */
			//	var ipb1 = e.control;
			//	if(e.keyCode == 13) { // enter
			//		smsSearchSend();
			//	}
			//}
			
			
			/*
			 * 서치 인풋에서 search 이벤트 발생 시 호출.
			 * Searchinput의 enter키 또는 검색버튼을 클릭하여 인풋의 값이 Search될때 발생하는 이벤트
			 */
			function onIpb1Search(/* cpr.events.CUIEvent */ e){
				/** 
				 * @type cpr.controls.SearchInput
				 */
				var ipb1 = e.control;
				smsSearchSend();
			}
			
			function smsSearchSend(){
				app.lookup("Src_memberList").clear();
					
				var checked = app.lookup("cmb1").value;
				var input = app.lookup("ipb1").value;
				app.lookup("dm_search").setValue("search_combo", checked);
				app.lookup("dm_search").setValue("search_input", input);
				
				app.lookup("sms_getUserListSearch").send();	
			}
			
			/*
			 * "선택" 버튼에서 click 이벤트 발생 시 호출.
			 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
			 */
			function onButtonClick4(/* cpr.events.CMouseEvent */ e){
				/** 
				 * @type cpr.controls.Button
				 */
				var button = e.control;
				
				if(confirm("저장하시겠습니까?")){
					var srcUserList = app.lookup("Src_memberList"); 
					var toRptUserList = app.lookup("toRpt_memberList");
					var refUserList = app.lookup("ref_memberList");
					if(toRptUserList.getRowCount() == 0){
						alert("보고 대상을 선택하셔야 합니다.");
						return;
					}
					
					var dmReport = app.lookup("dm_Report");
					
					var total = toRptUserList.getRowCount();
					var toRpt = toRptUserList.getRow(0).getValue("mem_name");
					var toRptID = toRptUserList.getRow(0).getValue("mem_id");
					
					dmReport.setValue("rpt_toRpt", toRpt);
					dmReport.setValue("rpt_toRptID", toRptID);
					
					total = refUserList.getRowCount();
					var ref = "";
					var refID = "";
					for(var i=0; i<total; i++){
						var row = refUserList.getRow(i);
						if(row){
							var mem_id = row.getValue("mem_id");
							var mem_name = row.getValue("mem_name");
							ref = ref + "," + mem_name;
							refID = refID + ","+ mem_id;
						}
					}
					ref = ref.substring(1,ref.length);
					refID = refID.substring(1,refID.length);
					dmReport.setValue("rpt_ref", ref);
					dmReport.setValue("rpt_refID", refID);
					
					var returnValue = {
						dmReport : dmReport,
						srcUserList : srcUserList,
						toRptUserList : toRptUserList,
						refUserList : refUserList
					}
			//		console.log('이거 왜 안뜸');
					app.close(returnValue);
					return;
				} else{
					return;
				}
				
			//	console.log('보내기전 rpt : ' + toRptUserList.getRowDataRanged());
			//	console.log(refUserList.getRowDatasByState(cpr.data.tabledata.RowState.UNCHANGED));
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
				var dmMemInfo = app.lookup("dm_memberInfo");
				var grd = app.lookup("userList");
				
				if(result == 0){
					var memPart = dmMemInfo.getString("mem_part");
					var memRank = "팀장"
					var cnt = grd.getRowCount();
					
					for(var i=0; i<cnt; i++){
						if(grd.getRow(i).getValue("mem_part") == memPart && grd.getRow(i).getValue("mem_rank") == memRank){
							grd.setEditRowIndex(i)
						}
					}
					
				} else {
					alert(getErrorString(result));
				}
			}
			
			
			/*
			 * 그리드에서 row-dblclick 이벤트 발생 시 호출.
			 * detail이 row를 더블클릭 한 경우 발생하는 이벤트.
			 */
			function onUserListRowDblclick(/* cpr.events.CGridMouseEvent */ e){
				/** 
				 * @type cpr.controls.Grid
				 */
				var userList = e.control;
				var rowIndex = userList.getSelectedRowIndex();
				if(rowIndex != null){
					if(!userList.getSelectedRow().checked){
						userList.setCheckRowIndex(rowIndex, true);
						return;		
					} else {
						userList.setCheckRowIndex(rowIndex, false);
						return;		
					}
				}
			}
			
			
			/*
			 * 그리드에서 row-dblclick 이벤트 발생 시 호출.
			 * detail이 row를 더블클릭 한 경우 발생하는 이벤트.
			 */
			function onToRptUserListRowDblclick(/* cpr.events.CGridMouseEvent */ e){
				/** 
				 * @type cpr.controls.Grid
				 */
				var toRptUserList = e.control;
				var rowIndex = toRptUserList.getSelectedRowIndex();
				if(rowIndex != null){
					if(!toRptUserList.getSelectedRow().checked){
						toRptUserList.setCheckRowIndex(rowIndex, true);
						return;		
					} else {
						toRptUserList.setCheckRowIndex(rowIndex, false);
						return;		
					}
				}
			}
			
			
			/*
			 * 그리드에서 row-dblclick 이벤트 발생 시 호출.
			 * detail이 row를 더블클릭 한 경우 발생하는 이벤트.
			 */
			function onRefUserListRowDblclick(/* cpr.events.CGridMouseEvent */ e){
				/** 
				 * @type cpr.controls.Grid
				 */
				var refUserList = e.control;
				var rowIndex = refUserList.getSelectedRowIndex();
				if(rowIndex != null){
					if(!refUserList.getSelectedRow().checked){
						refUserList.setCheckRowIndex(rowIndex, true);
						return;		
					} else {
						refUserList.setCheckRowIndex(rowIndex, false);
						return;		
					}	
				}
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
				var dmMemInfo = app.lookup("dm_memberInfo");
				var grd = app.lookup("userList");
				
				if(result == 0){
					var memPart = dmMemInfo.getString("mem_part");
					var memRank = "팀장"
					var cnt = grd.getRowCount();
					
					for(var i=0; i<cnt; i++){
						if(grd.getRow(i).getValue("mem_part") == memPart && grd.getRow(i).getValue("mem_rank") == memRank){
							grd.setEditRowIndex(i)
						}
					}
					
				} else {
					alert("세션이 끊어졌습니다.");
					app.close();
				}
			};
			// End - User Script
			
			// Header
			var dataSet_1 = new cpr.data.DataSet("Src_memberList");
			dataSet_1.parseData({
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
			app.register(dataSet_1);
			
			var dataSet_2 = new cpr.data.DataSet("toRpt_memberList");
			dataSet_2.parseData({
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
			app.register(dataSet_2);
			
			var dataSet_3 = new cpr.data.DataSet("ref_memberList");
			dataSet_3.parseData({
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
			app.register(dataSet_3);
			
			var dataSet_4 = new cpr.data.DataSet("ds_rank");
			dataSet_4.parseData({
				"columns" : [
					{"name": "rank_name"},
					{
						"name": "rank_idx",
						"dataType": "number"
					}
				]
			});
			app.register(dataSet_4);
			
			var dataSet_5 = new cpr.data.DataSet("ds_part");
			dataSet_5.parseData({
				"columns" : [
					{"name": "part_name"},
					{
						"name": "part_idx",
						"dataType": "number"
					}
				]
			});
			app.register(dataSet_5);
			var dataMap_1 = new cpr.data.DataMap("Result");
			dataMap_1.parseData({
				"columns" : [{"name": "ResultCode"}]
			});
			app.register(dataMap_1);
			
			var dataMap_2 = new cpr.data.DataMap("dm_search");
			dataMap_2.parseData({
				"columns" : [
					{"name": "search_combo"},
					{"name": "search_input"}
				]
			});
			app.register(dataMap_2);
			
			var dataMap_3 = new cpr.data.DataMap("dm_Report");
			dataMap_3.parseData({
				"columns" : [
					{"name": "rpt_toRpt"},
					{"name": "rpt_ref"},
					{"name": "rpt_toRptID"},
					{"name": "rpt_refID"}
				]
			});
			app.register(dataMap_3);
			
			var dataMap_4 = new cpr.data.DataMap("dm_memberInfo");
			dataMap_4.parseData({
				"columns" : [
					{"name": "mem_id"},
					{"name": "mem_name"},
					{"name": "mem_rank"},
					{"name": "mem_part"}
				]
			});
			app.register(dataMap_4);
			var submission_1 = new cpr.protocols.Submission("sms_getUserList");
			submission_1.method = "get";
			submission_1.action = "/bsmg/user/userList";
			submission_1.addResponseData(dataSet_1, false);
			submission_1.addResponseData(dataMap_1, false);
			if(typeof onSms_getUserListSubmitDone == "function") {
				submission_1.addEventListener("submit-done", onSms_getUserListSubmitDone);
			}
			app.register(submission_1);
			
			var submission_2 = new cpr.protocols.Submission("sms_getUserListSearch");
			submission_2.method = "get";
			submission_2.action = "/bsmg/user/userSearch";
			submission_2.addRequestData(dataMap_2);
			submission_2.addResponseData(dataSet_1, false);
			submission_2.addResponseData(dataMap_1, false);
			if(typeof onSms_getUserListSearchSubmitDone == "function") {
				submission_2.addEventListener("submit-done", onSms_getUserListSearchSubmitDone);
			}
			app.register(submission_2);
			
			var submission_3 = new cpr.protocols.Submission("sms_chkLogin");
			submission_3.method = "get";
			submission_3.action = "/bsmg/login/chkLogin";
			submission_3.addResponseData(dataMap_1, false);
			submission_3.addResponseData(dataMap_4, false);
			if(typeof onSms_chkLoginSubmitDone == "function") {
				submission_3.addEventListener("submit-done", onSms_chkLoginSubmitDone);
			}
			app.register(submission_3);
			
			app.supportMedia("all and (min-width: 1024px)", "default");
			app.supportMedia("all and (min-width: 500px) and (max-width: 1023px)", "tablet");
			app.supportMedia("all and (max-width: 499px)", "mobile");
			
			// Configure root container
			var container = app.getContainer();
			container.style.setClasses(["cl-form-group"]);
			container.style.css({
				"width" : "100%",
				"top" : "0px",
				"height" : "100%",
				"left" : "0px"
			});
			
			// Layout
			var formLayout_1 = new cpr.controls.layouts.FormLayout();
			formLayout_1.topMargin = "0px";
			formLayout_1.rightMargin = "0px";
			formLayout_1.bottomMargin = "0px";
			formLayout_1.leftMargin = "0px";
			formLayout_1.horizontalSpacing = "0px";
			formLayout_1.verticalSpacing = "0px";
			formLayout_1.horizontalSeparatorWidth = 1;
			formLayout_1.verticalSeparatorWidth = 1;
			formLayout_1.setColumns(["2fr", "32px", "1fr", "1fr"]);
			formLayout_1.setUseColumnShade(0, true);
			formLayout_1.setUseColumnShade(2, true);
			formLayout_1.setRows(["30px", "30px", "30px", "1fr", "40px"]);
			container.setLayout(formLayout_1);
			
			// UI Configuration
			var group_1 = new cpr.controls.Container();
			group_1.style.setClasses(["cl-form-group"]);
			// Layout
			var formLayout_2 = new cpr.controls.layouts.FormLayout();
			formLayout_2.topMargin = "0px";
			formLayout_2.rightMargin = "0px";
			formLayout_2.bottomMargin = "0px";
			formLayout_2.leftMargin = "0px";
			formLayout_2.horizontalSpacing = "0px";
			formLayout_2.verticalSpacing = "0px";
			formLayout_2.horizontalSeparatorWidth = 1;
			formLayout_2.verticalSeparatorWidth = 1;
			formLayout_2.setColumns(["1fr", "2fr", "2fr", "100px"]);
			formLayout_2.setUseColumnShade(0, true);
			formLayout_2.setUseColumnShade(2, true);
			formLayout_2.setRows(["1fr"]);
			group_1.setLayout(formLayout_2);
			(function(container){
				var button_1 = new cpr.controls.Button();
				button_1.value = "선택";
				if(typeof onButtonClick4 == "function") {
					button_1.addEventListener("click", onButtonClick4);
				}
				container.addChild(button_1, {
					"colIndex": 3,
					"rowIndex": 0
				});
				var group_2 = new cpr.controls.Container();
				// Layout
				var xYLayout_1 = new cpr.controls.layouts.XYLayout();
				group_2.setLayout(xYLayout_1);
				(function(container){
				})(group_2);
				container.addChild(group_2, {
					"colIndex": 0,
					"rowIndex": 0,
					"colSpan": 3,
					"rowSpan": 1
				});
			})(group_1);
			container.addChild(group_1, {
				"colIndex": 0,
				"rowIndex": 0,
				"colSpan": 4,
				"rowSpan": 1
			});
			
			var output_1 = new cpr.controls.Output();
			output_1.value = "보고 대상";
			output_1.style.css({
				"background-color" : "#f8e4e4",
				"font-weight" : "bolder",
				"text-align" : "center"
			});
			container.addChild(output_1, {
				"colIndex": 2,
				"rowIndex": 1
			});
			
			var output_2 = new cpr.controls.Output();
			output_2.value = "참조 대상";
			output_2.style.css({
				"background-color" : "#e2f1cf",
				"font-weight" : "bolder",
				"text-align" : "center"
			});
			container.addChild(output_2, {
				"colIndex": 3,
				"rowIndex": 1
			});
			
			var group_3 = new cpr.controls.Container();
			group_3.style.setClasses(["cl-form-group"]);
			// Layout
			var formLayout_3 = new cpr.controls.layouts.FormLayout();
			formLayout_3.topMargin = "0px";
			formLayout_3.rightMargin = "0px";
			formLayout_3.bottomMargin = "0px";
			formLayout_3.leftMargin = "0px";
			formLayout_3.horizontalSpacing = "0px";
			formLayout_3.verticalSpacing = "0px";
			formLayout_3.horizontalSeparatorWidth = 1;
			formLayout_3.verticalSeparatorWidth = 1;
			formLayout_3.setColumns(["1fr", "3fr"]);
			formLayout_3.setUseColumnShade(0, true);
			formLayout_3.setRows(["1fr"]);
			group_3.setLayout(formLayout_3);
			(function(container){
				var searchInput_1 = new cpr.controls.SearchInput("ipb1");
				if(typeof onIpb1Search == "function") {
					searchInput_1.addEventListener("search", onIpb1Search);
				}
				container.addChild(searchInput_1, {
					"colIndex": 1,
					"rowIndex": 0,
					"colSpan": 1,
					"rowSpan": 1
				});
				var comboBox_1 = new cpr.controls.ComboBox("cmb1");
				comboBox_1.value = "0";
				comboBox_1.fixedListWidth = true;
				comboBox_1.placeholder = "전체";
				comboBox_1.preventInput = true;
				(function(comboBox_1){
					comboBox_1.addItem(new cpr.controls.Item("전체", "0"));
					comboBox_1.addItem(new cpr.controls.Item("이름", "1"));
					comboBox_1.addItem(new cpr.controls.Item("직급", "2"));
					comboBox_1.addItem(new cpr.controls.Item("부서", "3"));
				})(comboBox_1);
				container.addChild(comboBox_1, {
					"colIndex": 0,
					"rowIndex": 0,
					"colSpan": 1,
					"rowSpan": 1
				});
			})(group_3);
			container.addChild(group_3, {
				"colIndex": 0,
				"rowIndex": 2
			});
			
			var grid_1 = new cpr.controls.Grid("toRptUserList");
			grid_1.init({
				"dataSet": app.lookup("toRpt_memberList"),
				"columnMovable": true,
				"resizableColumns": "all",
				"columns": [
					{"width": "25px"},
					{"width": "100px"},
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
								cell.columnType = "checkbox";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 1},
							"configurator": function(cell){
								cell.targetColumnName = "mem_name";
								cell.filterable = false;
								cell.sortable = false;
								cell.text = "이름";
								cell.style.css({
									"background-color" : "#ffffff"
								});
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 2},
							"configurator": function(cell){
								cell.targetColumnName = "mem_rank";
								cell.filterable = false;
								cell.sortable = false;
								cell.text = "직급";
								cell.style.css({
									"background-color" : "#ffffff"
								});
							}
						}
					]
				},
				"detail": {
					"rows": [{"height": "24px"}],
					"cells": [
						{
							"constraint": {"rowIndex": 0, "colIndex": 0},
							"configurator": function(cell){
								cell.columnType = "checkbox";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 1},
							"configurator": function(cell){
								cell.columnName = "mem_name";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 2},
							"configurator": function(cell){
								cell.columnName = "mem_rank";
								cell.control = (function(){
									var comboBox_2 = new cpr.controls.ComboBox("cmb2");
									comboBox_2.readOnly = true;
									(function(comboBox_2){
										comboBox_2.setItemSet(app.lookup("ds_rank"), {
											"label": "rank_name",
											"value": "rank_idx"
										})
									})(comboBox_2);
									comboBox_2.bind("value").toDataColumn("mem_rank");
									return comboBox_2;
								})();
							}
						}
					]
				}
			});
			grid_1.style.css({
				"background-color" : "#f8e4e4"
			});
			if(typeof onToRptUserListRowDblclick == "function") {
				grid_1.addEventListener("row-dblclick", onToRptUserListRowDblclick);
			}
			container.addChild(grid_1, {
				"colIndex": 2,
				"rowIndex": 2,
				"colSpan": 1,
				"rowSpan": 3
			});
			
			var group_4 = new cpr.controls.Container();
			group_4.style.setClasses(["cl-form-group"]);
			// Layout
			var formLayout_4 = new cpr.controls.layouts.FormLayout();
			formLayout_4.topMargin = "0px";
			formLayout_4.rightMargin = "0px";
			formLayout_4.bottomMargin = "0px";
			formLayout_4.leftMargin = "0px";
			formLayout_4.horizontalSpacing = "0px";
			formLayout_4.verticalSpacing = "0px";
			formLayout_4.horizontalSeparatorWidth = 1;
			formLayout_4.verticalSeparatorWidth = 1;
			formLayout_4.setColumns(["1fr"]);
			formLayout_4.setRows(["1fr", "32px", "32px", "32px", "25px", "32px", "32px", "32px", "1fr"]);
			group_4.setLayout(formLayout_4);
			(function(container){
				var button_2 = new cpr.controls.Button("btnUserAdd");
				button_2.style.css({
					"background-color" : "#f8e4e4",
					"background-repeat" : "no-repeat",
					"background-image" : "url('theme/images/arrow_icon/foward.png')",
					"background-position" : "center"
				});
				if(typeof onButtonClick == "function") {
					button_2.addEventListener("click", onButtonClick);
				}
				container.addChild(button_2, {
					"colIndex": 0,
					"rowIndex": 2
				});
				var button_3 = new cpr.controls.Button("btnUserRemove");
				button_3.style.css({
					"background-color" : "#f8e4e4",
					"background-repeat" : "no-repeat",
					"background-position" : "center",
					"background-image" : "url('theme/images/arrow_icon/rewind.png')"
				});
				if(typeof onBtnUserRemoveClick == "function") {
					button_3.addEventListener("click", onBtnUserRemoveClick);
				}
				container.addChild(button_3, {
					"colIndex": 0,
					"rowIndex": 3
				});
				var button_4 = new cpr.controls.Button();
				button_4.style.css({
					"background-color" : "#e2f1cf",
					"background-repeat" : "no-repeat",
					"background-position" : "center",
					"background-image" : "url('theme/images/arrow_icon/foward.png')"
				});
				if(typeof onButtonClick2 == "function") {
					button_4.addEventListener("click", onButtonClick2);
				}
				container.addChild(button_4, {
					"colIndex": 0,
					"rowIndex": 6
				});
				var button_5 = new cpr.controls.Button();
				button_5.style.css({
					"background-color" : "#e2f1cf",
					"background-repeat" : "no-repeat",
					"background-image" : "url('theme/images/arrow_icon/rewind.png')",
					"background-position" : "center"
				});
				if(typeof onButtonClick3 == "function") {
					button_5.addEventListener("click", onButtonClick3);
				}
				container.addChild(button_5, {
					"colIndex": 0,
					"rowIndex": 7
				});
				var output_3 = new cpr.controls.Output();
				output_3.value = "보고";
				output_3.style.css({
					"background-color" : "#ffffff",
					"color" : "#000000",
					"font-weight" : "bold"
				});
				container.addChild(output_3, {
					"colIndex": 0,
					"rowIndex": 1
				});
				var output_4 = new cpr.controls.Output();
				output_4.value = "참조";
				output_4.style.css({
					"background-color" : "#ffffff",
					"color" : "#63abf3",
					"font-weight" : "bold"
				});
				container.addChild(output_4, {
					"colIndex": 0,
					"rowIndex": 5
				});
			})(group_4);
			container.addChild(group_4, {
				"colIndex": 1,
				"rowIndex": 2,
				"colSpan": 1,
				"rowSpan": 3
			});
			
			var grid_2 = new cpr.controls.Grid("refUserList");
			grid_2.init({
				"dataSet": app.lookup("ref_memberList"),
				"columnMovable": true,
				"resizableColumns": "all",
				"columns": [
					{"width": "25px"},
					{"width": "100px"},
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
								cell.columnType = "checkbox";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 1},
							"configurator": function(cell){
								cell.targetColumnName = "mem_name";
								cell.filterable = false;
								cell.sortable = false;
								cell.text = "이름";
								cell.style.css({
									"background-color" : "#ffffff"
								});
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 2},
							"configurator": function(cell){
								cell.targetColumnName = "mem_rank";
								cell.filterable = false;
								cell.sortable = false;
								cell.text = "직급";
								cell.style.css({
									"background-color" : "#ffffff"
								});
							}
						}
					]
				},
				"detail": {
					"rows": [{"height": "24px"}],
					"cells": [
						{
							"constraint": {"rowIndex": 0, "colIndex": 0},
							"configurator": function(cell){
								cell.columnType = "checkbox";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 1},
							"configurator": function(cell){
								cell.columnName = "mem_name";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 2},
							"configurator": function(cell){
								cell.columnName = "mem_rank";
								cell.control = (function(){
									var comboBox_3 = new cpr.controls.ComboBox("cmb3");
									comboBox_3.readOnly = true;
									(function(comboBox_3){
										comboBox_3.setItemSet(app.lookup("ds_rank"), {
											"label": "rank_name",
											"value": "rank_idx"
										})
									})(comboBox_3);
									comboBox_3.bind("value").toDataColumn("mem_rank");
									return comboBox_3;
								})();
							}
						}
					]
				}
			});
			grid_2.style.css({
				"background-color" : "#e2f1cf"
			});
			if(typeof onRefUserListRowDblclick == "function") {
				grid_2.addEventListener("row-dblclick", onRefUserListRowDblclick);
			}
			container.addChild(grid_2, {
				"colIndex": 3,
				"rowIndex": 2,
				"colSpan": 1,
				"rowSpan": 3
			});
			
			var grid_3 = new cpr.controls.Grid("userList");
			grid_3.init({
				"dataSet": app.lookup("Src_memberList"),
				"columnMovable": true,
				"resizableColumns": "all",
				"columns": [
					{"width": "25px"},
					{"width": "100px"},
					{"width": "100px"},
					{"width": "100px"},
					{"width": "100px"}
				],
				"header": {
					"rows": [{"height": "24px"}],
					"cells": [
						{
							"constraint": {"rowIndex": 0, "colIndex": 0},
							"configurator": function(cell){
								cell.targetColumnName = "chk";
								cell.filterable = false;
								cell.sortable = false;
								cell.columnType = "checkbox";
								cell.text = "chk";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 1},
							"configurator": function(cell){
								cell.targetColumnName = "mem_id";
								cell.filterable = false;
								cell.sortable = false;
								cell.text = "아이디";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 2},
							"configurator": function(cell){
								cell.targetColumnName = "mem_name";
								cell.filterable = false;
								cell.sortable = false;
								cell.text = "이름";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 3},
							"configurator": function(cell){
								cell.targetColumnName = "mem_rank";
								cell.filterable = false;
								cell.sortable = true;
								cell.sortColumnName = "mem_rank";
								cell.text = "직급";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 4},
							"configurator": function(cell){
								cell.targetColumnName = "mem_part";
								cell.filterable = false;
								cell.sortable = true;
								cell.sortColumnName = "mem_part";
								cell.text = "부서";
							}
						}
					]
				},
				"detail": {
					"rows": [{"height": "24px"}],
					"cells": [
						{
							"constraint": {"rowIndex": 0, "colIndex": 0},
							"configurator": function(cell){
								cell.columnType = "checkbox";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 1},
							"configurator": function(cell){
								cell.columnName = "mem_id";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 2},
							"configurator": function(cell){
								cell.columnName = "mem_name";
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 3},
							"configurator": function(cell){
								cell.columnName = "mem_rank";
								cell.control = (function(){
									var comboBox_4 = new cpr.controls.ComboBox("cmb_rank");
									comboBox_4.readOnly = true;
									(function(comboBox_4){
										comboBox_4.setItemSet(app.lookup("ds_rank"), {
											"label": "rank_name",
											"value": "rank_idx"
										})
									})(comboBox_4);
									comboBox_4.bind("value").toDataColumn("mem_rank");
									return comboBox_4;
								})();
							}
						},
						{
							"constraint": {"rowIndex": 0, "colIndex": 4},
							"configurator": function(cell){
								cell.columnName = "mem_part";
								cell.control = (function(){
									var comboBox_5 = new cpr.controls.ComboBox("cmb_part");
									comboBox_5.readOnly = true;
									(function(comboBox_5){
										comboBox_5.setItemSet(app.lookup("ds_part"), {
											"label": "part_name",
											"value": "part_idx"
										})
									})(comboBox_5);
									comboBox_5.bind("value").toDataColumn("mem_part");
									return comboBox_5;
								})();
							}
						}
					]
				}
			});
			if(typeof onUserListRowDblclick == "function") {
				grid_3.addEventListener("row-dblclick", onUserListRowDblclick);
			}
			container.addChild(grid_3, {
				"colIndex": 0,
				"rowIndex": 3
			});
			
			var pageIndexer_1 = new cpr.controls.PageIndexer("pageIndex");
			pageIndexer_1.visible = false;
			pageIndexer_1.init(1, 1, 1);
			container.addChild(pageIndexer_1, {
				"colIndex": 0,
				"rowIndex": 4
			});
			
			var output_5 = new cpr.controls.Output();
			output_5.value = "사용자 목록";
			output_5.style.css({
				"font-weight" : "bolder",
				"text-align" : "center"
			});
			container.addChild(output_5, {
				"colIndex": 0,
				"rowIndex": 1
			});
			if(typeof onBodyLoad == "function"){
				app.addEventListener("load", onBodyLoad);
			}
		}
	});
	app.title = "bm_toRptList";
	cpr.core.Platform.INSTANCE.register(app);
})();
