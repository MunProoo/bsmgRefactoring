/************************************************
 * bm_login.js
 * Created at 2022. 5. 10. 오전 9:12:18.
 *
 * @author SW2Team
 ************************************************/



/*
 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
 */
function onBodyLoad(/* cpr.events.CEvent */ e){
	// 여기선 초기 세팅이 필요하지 않을듯? Nope. 로그인체크를 해야함
	app.lookup("sms_chkLogin").send();
}



/*
 * "Login" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	
	if(app.lookup("mem_id").value.length < 1) {
		alert("ID를 입력하세요.");
		app.lookup("mem_id").focus();
		return;
	}
	else if(app.lookup("mem_pw").value.length < 1) {
		alert("비밀번호를 입력하세요.");
		app.lookup("mem_pw").focus();
		return;
	}
	else{
		app.lookup("sms_login").send();
	}
	
	
}

/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_loginSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_login = e.control;
	var result = app.lookup("Result").getString("ResultCode");

	var dmMember = app.lookup("dm_memberInfo");
	if(result == 0){
		var mem_name = app.lookup("dm_memberInfo").getString("mem_name");
		alert(mem_name + "님 반갑습니다.");
		
		cpr.core.App.load("app/Bsmg/bm_Main", function(newapp){
			var applicationInfo = app.lookup("dm_memberInfo").getDatas();
			app.close();
			var newInst = newapp.createNewInstance();
			newInst.run().callAppMethod("setMemberInfo", applicationInfo);
		});
	} else {
		alert(getErrorString(result));
		app.lookup("mem_id").value = "";
		app.lookup("mem_pw").value = "";
		app.lookup("mem_id").focus();
		return;
	}
	
	
}




/*
 * 인풋 박스에서 keyup 이벤트 발생 시 호출.
 * 사용자가 키에서 손을 뗄 때 발생하는 이벤트.
 */
function onMem_pwKeyup(/* cpr.events.CKeyboardEvent */ e){
	/** 
	 * @type cpr.controls.InputBox
	 */
	var mem_pw = e.control;
	// 엔터키로 클릭
	if(e.keyCode == "13") app.lookup("login").click();
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
	
	if (result == 0){
		cpr.core.App.load("app/Bsmg/bm_Main", function(newapp){
			var applicationInfo = app.lookup("dm_memberInfo").getDatas();
			app.close();
			var newInst = newapp.createNewInstance();
			newInst.run().callAppMethod("setMemberInfo", applicationInfo);
		});
	} else {
		console.log("로그인 안됨");
	}
}


/*
 * 서브미션에서 submit-error 이벤트 발생 시 호출.
 * 통신 중 문제가 생기면 발생합니다.
 */
function onSms_chkLoginSubmitError(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_chkLogin = e.control;
	
}


/*
 * 서브미션에서 error-status 이벤트 발생 시 호출.
 * 서버로 부터 에러로 분류되는 HTTP상태 코드를 전송받았을 때 발생합니다.
 */
function onSms_loginErrorStatus(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_login = e.control;
	alert("권한이 없습니다. (토큰 Error)");
}


/*
 * 서브미션에서 error-status 이벤트 발생 시 호출.
 * 서버로 부터 에러로 분류되는 HTTP상태 코드를 전송받았을 때 발생합니다.
 */
function onSms_chkLoginErrorStatus(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_chkLogin = e.control;
	alert("권한이 없습니다. (토큰 Error)");
}
