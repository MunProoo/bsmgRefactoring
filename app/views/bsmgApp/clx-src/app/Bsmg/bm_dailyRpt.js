/************************************************
 * bm_dailyRpt.js
 * Created at 2022. 5. 19. 오전 9:48:38.
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
	momentToday();
//	app.lookup("sms_setAttr").send();

	// 싱글톤 사용해서 서버 부담 완화
	var dsList = dataManager.getDsAttrTree(); 
	dsList.copyToDataSet(app.lookup("ds_List"));
	app.lookup("lcb1").redraw();

	app.lookup("sms_chkLogin").send();
	
	makeTitle();
}

function makeTitle(){
	var mem_name = app.lookup("dm_memberInfo").getString("mem_name");
	var rpt_title = app.lookup("rpt_title");
	var rpt_date = app.lookup("rpt_date").value;
	var title = rpt_date.substring(0, 4)+"년 "+ rpt_date.substring(4,6)+"월 "+rpt_date.substring(6,8)+"일 ";
	title = title + mem_name + " 일일 업무보고";
	rpt_title.value = title;
}

// 데이트 인풋 컨트롤 기본 날짜 지정
function momentToday(){
	var vcDateInput = app.lookup("rpt_date");
	vcDateInput.format = "YYYYMMDDHHmmss";
	
	var vsToday = moment().format(vcDateInput.format);
	vcDateInput.value = vsToday;
	
}

/*
 * "선택" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var dmReport = app.lookup("dm_reportInfo");
	var srcUserList = app.lookup("Src_memberList");
	var toRptUserList = app.lookup("toRpt_memberList");
	var refUserList = app.lookup("ref_memberList");
	var toRpt = app.lookup("toRpt");
	var ref = app.lookup("ref");
	
	app.getRootAppInstance().openDialog("app/Bsmg/bm_toRptList", {
		width : 1000, height : 600
	}, function(dialog){
		dialog.ready(function(dialogApp){
			dialog.modal = true;
			dialog.headerVisible = true;
			dialog.headerMovable = true;
			dialog.resizable = true;
			dialog.headerTitle = "보고 대상, 참조 대상 선택";
			dialog.initValue = {
				srcUserList : srcUserList,
				toRptUserList : toRptUserList,
				refUserList : refUserList
			};
			dialog.addEventListener("keyup", function(e){
				if(e.keyCode == 27) { // ESC
					dialog.close();
				}
			});
		});
	}).then(function(returnValue){
		
		dmReport.setValue("rpt_toRpt", returnValue.dmReport.getString("rpt_toRpt"));
		dmReport.setValue("rpt_ref", returnValue.dmReport.getString("rpt_ref"));
		dmReport.setValue("rpt_toRptID", returnValue.dmReport.getString("rpt_toRptID"));
		dmReport.setValue("rpt_refID", returnValue.dmReport.getString("rpt_refID"));
		
		srcUserList.clear();
		toRptUserList.clear();
		refUserList.clear();
		
		returnValue.srcUserList.copyToDataSet(srcUserList);
		returnValue.toRptUserList.copyToDataSet(toRptUserList);
		returnValue.refUserList.copyToDataSet(refUserList);
		
		toRpt.value = dmReport.getString("rpt_toRpt");
		ref.value = dmReport.getString("rpt_ref");
		
//		console.log('리턴밸류 2개 가능? ', as);  안된다 ㅋ
	})
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
//		console.log(app.lookup("ds_List").getRowDataRanged());
		app.lookup("lcb1").redraw();
	} else {
		alert(getErrorString(result));
	}
}


/*
 * "보고서 저장" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick2(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var date = app.lookup("rpt_date").value;
	var toRpt = app.lookup("toRpt").value;
	var ref = app.lookup("ref").value;
	
	var attr = app.lookup("lcb1").value;
	var title = app.lookup("rpt_title").value;
	var content = app.lookup("rpt_content").value;
	app.lookup("rpt_etc").value= app.lookup("rpt_etc").value.trim();
	
//	console.log(date); // 2022-05-24 16:41:11 
//	console.log(toRpt); // "" 
//	console.log(content);
// ----------------------------- 필수 정보 확인 ----------------------------

	if(toRpt == ""){
		alert("보고대상을 선택하세요.");
		return;
	} else if(attr == null) {
		alert("업무 속성을 선택하세요.");
		app.lookup("lcb1").focus();
		return;
	} else if(title == ""){
		alert("제목을 입력하세요.");
		return;
	} else if(content.trim() == ""){
		alert("업무 내용을 입력하세요");
		app.lookup("rpt_content").focus();
		return;
	} 
	
	attr = attr.split(",");
	// 업무속성 2를 안고르고 넘길 수도 있으므로.
	if(attr[1] == undefined){
		alert("업무 속성을 전부 선택하세요.");
		app.lookup("lcb1").focus();
		return;
	}
	
//	console.log(dmReport.getDatas()); // 데이터 없는 컬럼 :  attr1,attr2 , reporter
	var dmReport = app.lookup("dm_reportInfo");
	dmReport.setValue("rpt_attr1", attr[0]);
	dmReport.setValue("rpt_attr2", attr[1]);
//	console.log(app.lookup("ds_schedule").getRowDataRanged());
	// 순차적으로 보내기 위해서 동기통신으로 이게 문제라서 idx가 안갔나?
	
	var grd = app.lookup("grdSch");
	var cnt = grd.getRowCount();
	
	for(var i=0; i<cnt; i++){
		if(grd.getRow(i).getString("sc_content").trim() < 1){
			alert("일정에 공백만 넣을 수 없습니다.");
			return;
		} 
	}
	app.lookup("sms_registRpt").send();
	
	
	
	app.close(1);
	
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_registRptSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_registRpt = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	var rptIdx = Number(app.lookup("dm_reportInfo").getValue("rpt_idx"));
	if (result == 0){
		console.log(rptIdx);
	
		var grd = app.lookup("grdSch");	
		if(grd.getRowCount() > 0){
			app.lookup("sms_registSch").send();
		} else {
			alert("보고가 성공적으로 저장되었습니다.");
		}
		
		return;
	} else {
		alert(getErrorString(result));
	}
}



/*
 * "+" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick3(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var grid = app.lookup("grdSch");
	var endRow = grid.getViewingEndRowIndex();
	grid.insertRow(endRow, true);
}


/*
 * "-" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick4(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var grid = app.lookup("grdSch");
	var endRow = grid.getViewingEndRowIndex();
	grid.deleteRow(endRow);
	grid.showDeletedRow = false;
}




/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_registSchSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_registSch = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if (result == 0){
		alert("보고가 성공적으로 저장되었습니다.");
		app.close(1);
		return;
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
	if(result != 0){
		alert(getErrorString(result));
		app.close();
	} 
}


/*
 * 데이트 인풋에서 value-change 이벤트 발생 시 호출.
 * Dateinput의 value를 변경하여 변경된 값이 저장된 후에 발생하는 이벤트.
 */
function onRpt_dateValueChange(/* cpr.events.CValueChangeEvent */ e){
	/** 
	 * @type cpr.controls.DateInput
	 */
	var rpt_date = e.control;
	makeTitle();
}




/*
 * "닫기" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick5(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	app.close();
}
