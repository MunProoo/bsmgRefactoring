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
	
}
