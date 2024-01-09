/************************************************
 * bm_regist.js
 * Created at 2022. 5. 12. 오전 11:16:49.
 *
 * @author SW2Team
 ************************************************/
var idCheck = false;

/*
 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
 */
function onBodyLoad(/* cpr.events.CEvent */ e){
	app.lookup("sms_setRankPart").send();
	
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_overlapCheckSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_overlapCheck = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
		alert("사용 가능한 아이디입니다.");
		idCheck = true;
		app.lookup("mem_pw").focus();
		return;
	} else {
		alert(getErrorString(result));
		app.lookup("mem_id").focus();
		return;
	}
}



/*
 * 직급, 부서 정보 받아오기
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_getRankPartSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getRankPart = e.control;
	var result = app.lookup("Result").getValue("ResultCode");
	if(result == 0){
		//console.log(app.lookup("ds_rank").getRowCount())
		app.lookup("mem_rank").redraw();
		app.lookup("mem_part").redraw();
	} else {
		alert(getErrorString(result));
	}
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_registUserSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_registUser = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if (result == 0) {
		alert("사용자를 등록하였습니다.");
		app.close();
	} else {
		alert(getErrorString(result));
	}
	
}





/*
 * "사용자 등록" 버튼(regist)에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onRegistClick(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var regist = e.control;
	var id = app.lookup("mem_id");
	var pw = app.lookup("mem_pw");
	var name = app.lookup("mem_name");
	var rank = app.lookup("mem_rank");
	var part = app.lookup("mem_part");
	
	//console.log(id, pw, name, rank, part);
	if(idCheck == false){
		alert("아이디 중복 검사를 해야합니다.");
	} 
	else if(pw.value.trim().length < 1){
		alert("비밀번호를 입력해야 합니다."); 
		pw.focus();	
	} 
	else if(name.value.trim().length < 1) {
		alert("이름을 입력해야 합니다.");
		name.focus();	
	} 
	else if(rank.value == "") {
		alert("직급을 선택해야 합니다.");	
	}
	else if(part.value == "") {
		alert("부서를 선택해야 합니다.");	
	}
	else{
		var memInfo = app.lookup("dm_memberInfo");
		memInfo.setValue("mem_rank", rank.value);
		memInfo.setValue("mem_part", part.value);
		//console.log(rank.value);
		//console.log(app.lookup("dm_memberInfo").getValue("mem_rank"))
		app.lookup("sms_registUser").send();
//		console.log(app.lookup("sms_registUser").method);
	}
}


/*
 * "중복 검사" 버튼(overlapCh)에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onOverlapChClick(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var overlapCh = e.control;
	var id = app.lookup("mem_id").value;
	
	if(id.trim().length < 1){
		alert("아이디를 입력하세요.(※공백만으로 아이디를 지정할 순 없습니다.)");
		return;
	}
	else{
		app.lookup("sms_idCheck").send();
		return;
	}
}


/*
 * 인풋 박스에서 input 이벤트 발생 시 호출.
 * 입력상자에 보여주는 텍스트가 키보드로부터 입력되어 변경되었을때 발생하는 이벤트.
 */
function onMem_idInput(/* cpr.events.CKeyboardEvent */ e){
	/** 
	 * @type cpr.controls.InputBox
	 */
	var mem_id = e.control;
	idCheck = false;
}
