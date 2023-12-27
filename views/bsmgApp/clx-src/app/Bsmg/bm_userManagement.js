/************************************************
 * bm_userManagement.js
 * Created at 2022. 6. 2. 오후 3:09:29.
 *
 * @author SW2Team
 ************************************************/
var dataManager = cpr.core.Module.require("lib/DataManager");


/*
 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
 */
function onBodyLoad(/* cpr.events.CEvent */ e){
	dataManager = getDataManager();
	var dsRank = dataManager.getRankList();
	var dsPart = dataManager.getPartList();
	
	dsRank.copyToDataSet(app.lookup("ds_rank"));
	dsPart.copyToDataSet(app.lookup("ds_part"));

	
	app.lookup("sms_getUserList").send();
//	app.lookup("sms_setRankPart").send();
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
	var src = app.lookup("Src_memberList");
	var copy = app.lookup("ds_memberListCopy");
	if (result == 0){
		src.copyToDataSet(copy);
		
		insertRankPartValue();
	}
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
 * "삭제" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var userList = app.lookup("userList");
	var indices = userList.getCheckRowIndices();
	
	if(indices.length == 0){
		alert("삭제할 사용자를 선택하세요");
		return;
	}
	
	var row = userList.getRow(indices[0]);
	var memID = row.getValue("mem_id");
//	console.log(memID);
	
	if(confirm("정말 선택한 사용자를 삭제하시겠습니까?")){
		app.lookup("sms_delUser").action = "/bsmg/user/deleteUser/"+memID;
		app.lookup("sms_delUser").send();
	}
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_delUserSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_delUser = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if (result == 0){
		app.lookup("sms_getUserList").send();
		app.lookup("sms_setRankPart").send();
		alert("정상적으로 삭제되었습니다.");
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
	var result = app.lookup("Result").getString("ResultCode");
	
	if(result==0){
		insertRankPartValue();
	}
}

function insertRankPartValue(){
	// init RankPart
	var dsRank = dataManager.getRankList();
	var dsPart = dataManager.getPartList();
	
	dsRank.copyToDataSet(app.lookup("ds_rank"));
	dsPart.copyToDataSet(app.lookup("ds_part"));
	
	
	// insertValue
	var dsUserList = app.lookup("Src_memberList");
	var grd = app.lookup("userList");
	for(var i=0; i<grd.getRowCount(); i++){
		var cmb4 = app.lookup("cmb2");
		var cmb5 = app.lookup("cmb3");
		cmb4.selectItemByValue(dsUserList.getRow(i).getValue("mem_rank"));
		cmb5.selectItemByValue(dsUserList.getRow(i).getValue("mem_part"));
		
	}
	grd.sort("mem_part ASC");
	grd.redraw();
}

/*
 * "수정" 버튼(update)에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onUpdateClick(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var update = e.control;
	var grd = app.lookup("userList");
	var indices = grd.getCheckRowIndices();
	var msg = "";
	var spaceFlag = true;
	
	if(indices.length == 0){
		alert("수정할 사용자가 없습니다.");
		return;
	} else if(indices.length == 1){
		msg = "선택한 1명의 사용자를 수정하시겠습니까?"
	} else {
		msg = "선택한 사용자들을 수정하시겠습니까?"
	}
	var putMemberList = app.lookup("ds_putMember");
	putMemberList.clear();
	
	if(confirm(msg)){
		indices.forEach(function(index){
			if(grd.getRow(index).getString("mem_name").trim().length < 1){
				alert("공백은 입력할 수 없습니다.");
				spaceFlag = false;
				return;
			} else {
				var rowData = grd.getRow(index).getRowData();
				putMemberList.addRowData(rowData);		
			}
		});
		if(spaceFlag){
			app.lookup("sms_putUserList").send();
		}
	} else {
		return;
	}
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_putUserListSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_putUserList = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
		alert('성공적으로 수정되었습니다.');
		app.lookup("sms_getUserList").send();
		app.lookup("sms_setRankPart").send();
	}
}


/*
 * "취소" 버튼(cancel)에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onCancelClick(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var cancel = e.control;
	var copy = app.lookup("ds_memberListCopy");
	var src = app.lookup("Src_memberList");
	var grd = app.lookup("userList");
	
	src.clear();
	copy.copyToDataSet(src);
	insertRankPartValue();
	app.lookup("userList").redraw();
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
 * 서브미션에서 before-send 이벤트 발생 시 호출.
 * XMLHttpRequest가 open된 후 send 함수가 호출되기 직전에 발생합니다.
 */
function onSms_getUserListBeforeSend(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getUserList = e.control;
	dsClear();
}

function dsClear(){
	var src = app.lookup("Src_memberList");
	var copy = app.lookup("ds_memberListCopy");
	var rank = app.lookup("ds_rank");
	var part = app.lookup("ds_part");	
	src.clear();
	copy.clear();
	rank.clear();
	part.clear();
}


/*
 * "닫기" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick2(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	app.close();
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
	if(result == 0){
		var src = app.lookup("Src_memberList");
		var copy = app.lookup("ds_memberListCopy");
		src.copyToDataSet(copy);
//		app.lookup("sms_setRankPart").send();

		insertRankPartValue();
	}
}


/*
 * 서브미션에서 before-send 이벤트 발생 시 호출.
 * XMLHttpRequest가 open된 후 send 함수가 호출되기 직전에 발생합니다.
 */
function onSms_getUserListSearchBeforeSend(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getUserListSearch = e.control;
	dsClear();
}
