/************************************************
 * bm_rptView.js
 * Created at 2022. 5. 27. 오전 9:43:58.
 *
 * @author SW2Team
 ************************************************/

var dataManager = cpr.core.Module.require("lib/DataManager");



/*
 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
 */
function onBodyLoad(/* cpr.events.CEvent */ e){
	
	app.lookup("sms_chkLogin").send();
}




/*
 * 루트 컨테이너에서 init 이벤트 발생 시 호출.
 * 앱이 최초 구성될 때 발생하는 이벤트 입니다.
 */
function onBodyInit(/* cpr.events.CEvent */ e){
	dataManager = getDataManager();
	var initValue = app.getHost().initValue;
	var rpt_idx = initValue.rpt_idx;
//	console.log("rptIdx : " +rpt_idx);
	app.lookup("dm_rptIdx").setValue("rpt_idx", Number(rpt_idx));
	
	// 싱글톤 사용해서 서버 부담 완화
	var dsList = dataManager.getDsAttrTree(); 
	dsList.copyToDataSet(app.lookup("ds_List"));
	
	
	app.lookup("sms_getRptInfo").send();
//	app.lookup("sms_setAttr").send();
	app.lookup("sms_getRptSchedule").send();
}

/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_getRptInfoSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getRptInfo = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	var src = app.lookup("dm_reportInfoSrc"); // 원본 복제
	var dmReportInfo = app.lookup("dm_reportInfo");
	var rpt_confirm = dmReportInfo.getString("rpt_confirm");
	if (result == 0){
		dmReportInfo.copyToDataMap(src);
		if(rpt_confirm == 'true'){ // 보고가 확인되었으면.
//			app.lookup("rpt_content").style.bind("background") = "url('../../images/sogood.png') no-repeat / contain"
			app.lookup("rpt_content").style.css({
				"background-image" : "linear-gradient(rgba(255,255,255,0.5),rgba(255, 255, 255, 0.5)),url('../../images/sogood.png')",
//				"opacity" : "0.5"
			});
			
		}
		setAttr();
		dateFormat();
	} else {
		alert(getErrorString(result));
	}
}

function dateFormat(){
	var rpt_date = app.lookup("rpt_date");
	var dmReportInfo = app.lookup("dm_reportInfo");
	var date = dmReportInfo.getString("rpt_date").substring(0,8);
	date = date.substring(0,4)+ "-"+date.substring(4,6)+"-"+date.substring(6);
	rpt_date.value = date;
}

/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_getRptScheduleSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getRptSchedule = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	var src = app.lookup("ds_scheduleSrc");
	var dsSchedule = app.lookup("ds_schedule");
	
	if (result == 0){
		dsSchedule.copyToDataSet(src);
		app.getContainer().redraw();
	} else {
		alert(getErrorString(result));
	}
}



/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_setAttrSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_setAttr = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
		setAttr();
	} else {
		alert(getErrorString(result));
	}
}

function setAttr(){ // 원본으로 돌리기
	var lcb = app.lookup("lcb1");
	var src = app.lookup("dm_reportInfoSrc");
	var value = src.getValue("rpt_attr1").toString() + "-" + src.getValue("rpt_attr2").toString();
	
	lcb.selectItemByValue(value);
	
//	app.lookup("lcb1").redraw();
}

function setSchedule(){ // 원본으로 돌리기
	var src = app.lookup("ds_scheduleSrc");
	var dsSchedule = app.lookup("ds_schedule");
	dsSchedule.clear();
	var grd = app.lookup("grdSch");
	src.copyToDataSet(dsSchedule);
//	grd.redraw();
}
function setRpt(){
	var src = app.lookup("dm_reportInfoSrc");
	var dmReportInfo = app.lookup("dm_reportInfo");
	dmReportInfo.clear();
	src.copyToDataMap(dmReportInfo);
	dateFormat();
	app.getContainer().redraw();
}

/*
 * "+" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick2(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var update = app.lookup("update").value;
	
	if(update == '수정'){
		return;
	} else {
		var grid = app.lookup("grdSch");
		var endRow = grid.getViewingEndRowIndex();
		grid.insertRow(endRow, true);
		return;
	}
}


/*
 * "-" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick3(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var update = app.lookup("update").value;
	if(update == '수정'){
		return;
	} else {
		var grid = app.lookup("grdSch");
//		var endRow = grid.getViewingEndRowIndex();
		var delRow = grid.getSelectedRowIndex();
		var dsSchedule = app.lookup("ds_schedule");
		console.log("delRow : ",grid.getRow(delRow).getRowData());
		dsSchedule.realDeleteRow(delRow);
		/*
		console.log("endRow : ", grid.getRow(endRow).getRowData());
		console.log("지워지는 값 : ", dsSchedule.getRow(endRow).getRowData());
		console.log(grid.getExportData());
		console.log(dsSchedule.getRowDataRanged());
		*/ 
		grid.deleteRow(delRow);
		for (var i=0; i<grid.rowCount; i++){ // 자꾸 삭제한 row의 바로 밑 row는 DELETED로 상태가 변경됨
			grid.setRowState(i, cpr.data.tabledata.RowState.UNCHANGED );
		}		
//		grid.showDeletedRow = false;
		
		return;
	}
}


/*
 * "수정" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick4(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var group = app.lookup("gr1");
	
	if(group.readOnly){
		group.readOnly = false;
		app.lookup("rpt_title").readOnly = true; // 제목은 변경X
		button.style.css("background-color", "#045729");
		button.value = "완료"
		return;
	} else {
		
		var attr = app.lookup("lcb1").value;
		var title = app.lookup("rpt_title").value;
		var content = app.lookup("rpt_content").value;
		var etc = app.lookup("rpt_etc").value;
		
		if(title == ""){
			alert("제목을 입력하세요.");
			return;
		} else if(content.trim() == "") {
			alert("업무 내용을 입력하세요");
			return;
		} else if(attr == null){
			alert("업무 속성을 선택하세요");
			return;
		} 
		
		attr = attr.split(",");
		if(attr[1] == undefined){
			alert("업무 속성을 전부 선택하세요.");
			return;
		}
		
		var grd = app.lookup("grdSch");
		var cnt = grd.getRowCount();
		for(var i=0; i<cnt; i++){
			if(grd.getRow(i).getString("sc_content").trim() < 1){
				alert("일정에 공백만 넣을 수 없습니다.");
				return;
			} 
		}
		
		group.readOnly = true;
		button.style.css("background-color", "#52c183");
		button.value = "수정";
		var dmReport = app.lookup("dm_reportInfo");
		var dmReportSrc = app.lookup("dm_reportInfoSrc");
		var dsSchedule = app.lookup("ds_schedule");
		var dsScheduleSrc = app.lookup("ds_scheduleSrc");
		app.lookup("rpt_etc").value = app.lookup("rpt_etc").value.trim();
		dmReport.setValue("rpt_attr1", attr[0]);
		dmReport.setValue("rpt_attr2", attr[1]);
		
//		console.log(app.lookup("ds_schedule").getRowDataRanged());
		app.lookup("sms_putRpt").send();
		
	}
}


/*
 * "닫기" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	app.close(1);
}


/*
 * "취소" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick5(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var group = app.lookup("gr1");
	
	setAttr();
	setSchedule();
	setRpt();
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_putShceduleSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_putShcedule = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
		alert("수정되었습니다.");
	} else {
		alert(getErrorString(result));
	}
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_putRptSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_putRpt = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
//		console.log("보고 수정 완료");
		app.lookup("sms_putShcedule").send();
	} else {
		alert(getErrorString(result));
	}
}


/*
 * "삭제" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick6(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var rptIdx = app.lookup("dm_rptIdx").getString("rpt_idx");
	if(confirm('정말 삭제하시곘습니까?')){
		app.lookup("sms_deleteRpt").action = "/bsmg/report/deleteRpt/"+rptIdx;
		app.lookup("sms_deleteRpt").send();
	} else {
		return;
	}
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_deleteRptSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_deleteRpt = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
		alert("보고가 삭제되었습니다.");
		app.close(1);
	} else {
		alert(getErrorString(result));
	}
}



/*
 * 서브미션에서 before-submit 이벤트 발생 시 호출.
 * 통신을 시작하기전에 발생합니다.
 */
function onSms_putShceduleBeforeSubmit(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	// dsSchedule에 데이터는 잘 들어가는데, 서브미션으로 보내지는 데이터셋에는 기존 데이터가 안들어간다 -> 추가한 일정만 들어간다.
	var sms_putShcedule = e.control;
//	var gridSch = app.lookup("grdSch");
	var dsSchedule = app.lookup("ds_schedule");
	
//	gridSch.commitData();
	
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
	if(result == 0){
		var mem_rank = app.lookup("dm_memberInfo").getValue("mem_rank")
		var mem_name = app.lookup("dm_memberInfo").getValue("mem_name");	
		var rpt_reporter_name = app.lookup("dm_reportInfo").getValue("rpt_reporter_name");
		var rpt_toRpt_name = app.lookup("dm_reportInfo").getValue("rpt_toRpt_name");
		var rpt_confirm = app.lookup("dm_reportInfo").getValue("rpt_confirm");
		
		if(mem_name == rpt_reporter_name && rpt_confirm == 'false'){
			app.lookup("update").visible = true;
			app.lookup("cancel").visible = true;
			app.lookup("delete").visible = true;
		} else if(mem_name == rpt_toRpt_name || mem_rank < Rank3){
			app.lookup("confirm").visible = true;
		} 
		app.getContainer().redraw();
	} else {
		alert(getErrorString(result));
	}
}


/*
 * "보고 확인" 버튼(confirm)에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onConfirmClick(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var confirm12 = e.control;
	var rpt_confirm = app.lookup("dm_reportInfo").getString("rpt_confirm");
	if(rpt_confirm == 'true'){
		alert('이미 확인된 보고입니다.');
		return;
	} else{
		if(confirm('보고를 확인하시겠습니까?')){
			app.lookup("sms_confirmRpt").send();
		}
	}
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_confirmRptSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_confirmRpt = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
		alert('정상적으로 확인되었습니다.');
		app.lookup("sms_getRptInfo").send();
	} else {
		alert(getErrorString(result));
	}
}
