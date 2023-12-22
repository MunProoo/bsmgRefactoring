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
		alert("세션이 끊어졌습니다.");
		app.close();
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
}
