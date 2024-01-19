/************************************************
 * bm_weekRptView.js
 * Created at 2022. 6. 7. 오후 4:28:47.
 *
 * @author SW2Team
 ************************************************/


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
	var initValue = app.getHost().initValue;
	var wRpt_idx = initValue.wRpt_idx;
	app.lookup("dm_wRptIdx").setValue("wRpt_idx", Number(wRpt_idx));
	app.lookup("sms_getWeekRptInfo").send();
	app.lookup("sms_setPart").send();
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_getWeekRptInfoSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getWeekRptInfo = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	var src = app.lookup("dm_weekRptInfoSrc");
	var dmWeekRptInfo = app.lookup("dm_weekRptInfo");
	
	if(result == 0){
		dmWeekRptInfo.copyToDataMap(src);
		dateFormat();
	} else {
		alert(getErrorString(result));
	}
}

function dateFormat(){
	var dmWeekRptInfo = app.lookup("dm_weekRptInfo");
	var wRpt_date = app.lookup("wRpt_date");
	var date = dmWeekRptInfo.getString("wRpt_date").substring(0,8);
	date = date.substring(0,4)+ "-"+date.substring(4,6)+"-"+date.substring(6);
	wRpt_date.value = date;
}

function setPart(){
	var cmb1 = app.lookup("cmb1");
	var src = app.lookup("dm_weekRptInfoSrc");
	cmb1.selectItemByValue(src.getString("wRpt_part"));
}

/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_setPartSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_setPart = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
		setPart();
	} else {
		alert(getErrorString(result));
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
 * "수정" 버튼(update)에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onUpdateClick(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var update = e.control;
	var group = app.lookup("gr1");
	if(group.readOnly){
		group.readOnly = false;
		update.style.css("background-color", "#045729");
		update.value = "완료";
		return;
	} else {
		group.readOnly = true;
		update.style.css("background-color", "#52c183");
		update.value = "수정";
		
		var title = app.lookup("wRpt_title").value;
		var content = app.lookup("wRpt_content").value;
		var part = app.lookup("cmb1").value;
		
		console.log(part);
		if(title == ""){
			alert("제목을 입력하세요.");
			return;
		} else if(content == "") {
			alert("업무 내용을 입력하세요");
			return;
		} else if(part == null){
			alert("보고할 부서를 선택하세요");
			return;
		} 
		
		var dmWeekRpt = app.lookup("dm_weekRptInfo");
		dmWeekRpt.setValue("wRpt_part", part);
		var dmWeekRptSrc = app.lookup("dm_weekRptInfoSrc");
		app.lookup("sms_putWeekRpt").send();
		
		// 수정한 값을 Src에도 저장
		dmWeekRpt.copyToDataMap(dmWeekRptSrc);
		app.getContainer().redraw();
	}
}



/*
 * "취소" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick2(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
//	var group = app.lookup("gr1");
	setPart();
	setWeekRpt();
}

function setWeekRpt(){
	var src = app.lookup("dm_weekRptInfoSrc");
	var dmWeekRptInfo = app.lookup("dm_weekRptInfo");
	dmWeekRptInfo.clear();
	src.copyToDataMap(dmWeekRptInfo);
	dateFormat();
	app.getContainer().redraw();
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_putWeekRptSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_putWeekRpt = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
		alert('주간 업무보고가 수정되었습니다.');
		return;
	} else {
		alert(getErrorString(result));
	}
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_getToRptSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getToRpt = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
		var teamLeader = app.lookup("dm_part").getString("team_leader");
		app.lookup("teamLeader").value = teamLeader;
		console.log('teamLeader : ',teamLeader);
	} else {
		alert(getErrorString(result));
	}
}


/*
 * 콤보 박스에서 item-click 이벤트 발생 시 호출.
 * 아이템 클릭시 발생하는 이벤트.
 */
function onCmb1ItemClick(/* cpr.events.CItemEvent */ e){
	/** 
	 * @type cpr.controls.ComboBox
	 */
	var cmb1 = e.control;
	var combo = cmb1.value;
	app.lookup("dm_part").setValue("part_idx", combo);
	app.lookup("sms_getToRpt").send();
}


/*
 * "삭제" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick3(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var wRpt_idx = app.lookup("dm_wRptIdx").getString("wRpt_idx");
	if(confirm('정말 삭제하시겠습니까?')){
		app.lookup("sms_deleteWeekRpt").action = "/bsmg/report/deleteWeekRpt/"+wRpt_idx;
		app.lookup("sms_deleteWeekRpt").send();
	} else {
		return;
	}
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_deleteWeekRptSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_deleteWeekRpt = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
		alert("정상적으로 삭제되었습니다.");
		app.close(1);
	} else {
		alert(getErrorString(result));
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
	
	if(result == 0){
		var mem_rank = app.lookup("dm_memberInfo").getString("mem_rank")
		var mem_name = app.lookup("dm_memberInfo").getString("mem_name");	
		var wRpt_reporter = app.lookup("dm_weekRptInfo").getValue("wRpt_reporter_name");
		var wRpt_toRpt = app.lookup("dm_weekRptInfo").getString("wRpt_toRpt");
		if(mem_name == wRpt_reporter){
			app.lookup("update").visible = true;
			app.lookup("cancel").visible = true;
			app.lookup("delete").visible = true;
		} else if(mem_name == wRpt_toRpt || (mem_rank < Rank3)){
			
		}
		app.getContainer().redraw();
	} else {
		alert(getErrorString(result));
		app.close();
	}
}

