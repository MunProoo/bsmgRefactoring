/************************************************
 * bm_weekRpt.js
 * Created at 2022. 6. 3. 오후 3:44:36.
 *
 * @author SW2Team
 ************************************************/

/* 페이지처리 */
var RowCount = 6; // 한페이지 로우 개수
var SearchFlag = false;
var AttrFlag = false;
var NewSearch = false; // 현재 페이지가 1이 아닌 상태. 새롭게 검색한다면 1페이지로 돌아가게
var REFRESH = false; // 새로 검색하는건지 페이지 이동인지 구분자


/*
 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
 */
function onBodyLoad(/* cpr.events.CEvent */ e){
	app.lookup("sms_getCategoryList").send();
	setPaging(0, 1, RowCount, 5);
	sendRptListRequest();
}

function setPaging(totCnt, pageIdx, RowCount, pageSize){
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
	var sms_getWeekRptList = app.lookup("sms_getWeekRptList");
	
	sms_getWeekRptList.setParameters("offset", offset);
	sms_getWeekRptList.setParameters("limit", RowCount);
	sms_getWeekRptList.send();
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_getWeekRptListSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getWeekRptList = e.control;
	var pageIndexer = app.lookup("pageIndex");
	var result = app.lookup("Result").getString("ResultCode");
	
	if(result == 0){
		var totalCount = app.lookup("totalCount").getValue("Count");
		
		AllGridColorWhite();
		
		if(REFRESH){
			setPaging(Number(totalCount), 1, RowCount, 5);
			REFRESH = false;
		} else {
			setPaging(Number(totalCount), pageIndexer.currentPageIndex, RowCount, 5);
		}
		SearchFlag = false;
		AttrFlag = false;
		return;
	} else {
		alert(getErrorString(result));
	}
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_getCategoryListSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getCategoryList = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
		app.lookup("tre1").redraw();
		return;
	} else {
		alert(getErrorString(result));
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
	var selectionEvent = new cpr.events.CSelectionEvent("pagechange", {
		oldSelection: e.oldSelection,
		newSelection: e.newSelection
	});
	app.dispatchEvent(selectionEvent);
	// 기본처리가 중단되었을 때 변경을 취소함.
	if(SearchFlag){
		sendSearchRequest();
	} else if(AttrFlag){
		sendPartRptRequest();
	} else{
		sendRptListRequest();
	}
	
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
	var selectionEvent = new cpr.events.CSelectionEvent("before-pagechange",{
		oldSelection: e.oldSelection,
		newSelection: e.newSelection
	});
	app.dispatchEvent(selectionEvent);
	if(selectionEvent.defaultPrevented == true) {
		e.preventDefault();
	}

}

function sendSearchRequest(){
	var pageIndexer = app.lookup("pageIndex");
	if (!SearchFlag){ // 검색중이 아니라면 페이지인덱서 새롭게 세팅
		NewSearch = true; 
		pageIndexer.currentPageIndex = 1;
	} else {
		NewSearch = false;
	}
	
	var pageIdx = pageIndexer.currentPageIndex;
	var offset = (pageIdx - 1) * RowCount;
	var combo = app.lookup("cmb1").value;
	var input = app.lookup("ipb1").value;
	
	app.lookup("ds_weekRptList").clear();
	app.lookup("dm_search").setValue("search_combo", combo);
	app.lookup("dm_search").setValue("search_input", input);
	var smsGetWeekRptSearch = app.lookup("sms_getWeekRptSearch");
	smsGetWeekRptSearch.setParameters("offset", offset);
	smsGetWeekRptSearch.setParameters("limit", RowCount);
	smsGetWeekRptSearch.send();
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
	sendSearchRequest();
}



/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_getWeekRptSearchSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getWeekRptSearch = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	var totalCount = app.lookup("totalCount").getValue("Count");
	var pageIndexer = app.lookup("pageIndex");
	if(result == 0){
		SearchFlag = true;
		AttrFlag = false;
		
		if(totalCount == 0){
			alert("자료가 없습니다.");
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
 * 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	REFRESH = true;
	sendRptListRequest();
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
	var grid = app.lookup("grd1");
	var pageIndexer = app.lookup("pageIndex");
	pageIndexer.currentPageIndex = 1;
	var partValue = tre1.getSelectionLast().value;
	partValue = partValue.split("-")[1]
//	console.log("partValue : ", partValue)
	if(partValue == undefined){
		partValue = 0;
	}
	
	app.lookup("dm_part").setValue("part_value", partValue);
	sendPartRptRequest();

}

function sendPartRptRequest(){
	var pageIndexer = app.lookup("pageIndex");
	if(!AttrFlag){
		NewSearch = true;
		pageIndexer.currentPageIndex = 1;
	} else {
		NewSearch = false;
	} 
	
	var pageIdx = pageIndexer.currentPageIndex;
	var offset = (pageIdx - 1) * RowCount;
	
	app.lookup("ds_weekRptList").clear();
	app.lookup("sms_getWeekRptCategory").setParameters("offset", offset);
	app.lookup("sms_getWeekRptCategory").setParameters("limit", RowCount);
	app.lookup("sms_getWeekRptCategory").send();
}

/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_getWeekRptCategorySubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getWeekRptCategory = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	var totalCount = app.lookup("totalCount").getValue("Count");
	var pageIndexer = app.lookup("pageIndex");
	
	if(result == 0){
		SearchFlag = false;
		AttrFlag = true;
		
		if(totalCount == 0){
			alert("자료가 없습니다.");
			return;
		}
		AllGridColorWhite();
		gridColorImpact(3);
		
		if(NewSearch){
			setPaging(Number(totalCount),1,RowCount, 5);
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
	var paramValue = row.getValue("wRpt_idx");
	
	app.getRootAppInstance().openDialog("app/Bsmg/bm_weekRptView", {
		width : 1000, height : 800
	}, function(dialog){
		dialog.modal = true;
		dialog.headerVisible = true;
		dialog.headerMovable = true;
		dialog.headerTitle = "주간 업무보고 상세 확인";
		dialog.headerClose = true;
		dialog.addEventListener("keyup", function(e){
			if(e.keyCode == 27){ // ESC
				dialog.close();
			}
		});
		dialog.initValue = {
			wRpt_idx : paramValue
		}
	}).then(function(returnValue){
		if(returnValue == 1){
			REFRESH = true;
			sendRptListRequest();
		}
	});
}

function gridColorWhite(cellIndex){
	var grd = app.lookup("grd1");
	var voCell = grd.detail.getColumn(cellIndex);
//	var voExpress = "switch(getValue(\"color\")){\n\tcase \"Red\" : \"red\"\n\tdefault : \"#transparent\"\n}";
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
}
