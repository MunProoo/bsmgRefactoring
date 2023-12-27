/************************************************
 * bm_list.js
 * Created at 2022. 5. 12. 오전 10:13:38.
 *
 * @author SW2Team
 ************************************************/

/* 페이지 처리 */
var RowCount = 6; // 한 페이지 행 개수
var SearchFlag = false; // 검색해서 받은 리스트
var AttrFlag = false; // 업무속성으로 검색
var NewSearch = false; // 현재 페이지가 1이 아닐때, 새로 search해서 리스트를 불러올 시 1페이지로 돌아가게.

var dataManager = cpr.core.Module.require("lib/DataManager"); // 싱글톤패턴

/*
 * 루트 컨테이너에서 load 이벤트 발생 시 호출.
 * 앱이 최초 구성된후 최초 랜더링 직후에 발생하는 이벤트 입니다.
 */
function onBodyLoad(/* cpr.events.CEvent */ e){
	dataManager = getDataManager();
	
	app.lookup("sms_setTree").send();
	setPaging(0, 1, RowCount, 5);
	sendRptListRequest();
}

function setPaging(totCnt,pageIdx ,RowCount, pageSize){
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
	var sms_getRptList = app.lookup("sms_getRptList");
	
	// dm_page에 setValue로 집어넣으면 서버에서 에러남... why? 뭐가 다르지?
	sms_getRptList.setParameters("offset", offset);
	sms_getRptList.setParameters("limit", RowCount);
	sms_getRptList.send();
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_getRptListSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getRptList = e.control;
	var pageIndexer = app.lookup("pageIndex");  
	var result = app.lookup("Result").getString("ResultCode");	

	if(result == 0){
		var totalCount = app.lookup("totalCount").getValue("Count");

		AllGridColorWhite();
		
		setPaging(Number(totalCount), pageIndexer.currentPageIndex, RowCount, 5);
		SearchFlag = false;
		AttrFlag = false;
		return;
	} else{
		alert("업무 보고 리스트를 불러오는 데 실패하였습니다.")
		return;
	}
}


/*
 * "일일 업무보고 작성" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	app.getRootAppInstance().openDialog("app/Bsmg/bm_dailyRpt", {
		width : 1000, height : 800
	}, function(dialog){
		dialog.ready(function(dialogApp){
			dialog.modal = true;
			dialog.headerVisible = true;
			dialog.headerMovable = true;
			dialog.resizable = true;
			dialog.headerTitle = "일일 업무보고 등록";
		});
	}).then(function(returnValue){
		if (returnValue == 1){
			if(AttrFlag){
				sendAttrRptRequest();
			} else if(SearchFlag){
				sendSearchRequest();
			} else {
				sendRptListRequest();
			}
		}
	})
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
	var selectionEvent = new cpr.events.CSelectionEvent("before-pagechange", {
		oldSelection: e.oldSelection,
		newSelection: e.newSelection
	});
	app.dispatchEvent(selectionEvent);
	// 기본처리가 중단되었을 때 변경을 취소함.
	if(selectionEvent.defaultPrevented == true) {
		e.preventDefault();
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
	var selectionEvent = new cpr.events.CSelectionEvent("pagechange",{
		oldSelection: e.oldSelection,
		newSelection: e.newSelection
	});
	app.dispatchEvent(selectionEvent);
//	console.log("선택 페이지 : " + pageIndex.currentPageIndex);
	if(SearchFlag){
		sendSearchRequest();
	} else if(AttrFlag){
		sendAttrRptRequest();
	} else{
		sendRptListRequest();
	}
}


function sendSearchRequest(){
	var pageIndexer = app.lookup("pageIndex");
	if(!SearchFlag){	// 카테고리에서 3페이지를 보다가 검색하는 경우, 시작이 3페이지이므로 1페이지로 초기화
		NewSearch = true;
		pageIndexer.currentPageIndex = 1;
	} else { // 이미 검색을 한 상태. 다음 페이지를 보려할 때
		NewSearch = false;
	}
	
	var pageIdx = pageIndexer.currentPageIndex;
	var offset = (pageIdx - 1) * RowCount;
	var combo = app.lookup("cmb1").value;
	var input = app.lookup("ipb1").value;	
	
	app.lookup("ds_rptList").clear();
	app.lookup("dm_search").setValue("search_combo", combo);
	app.lookup("dm_search").setValue("search_input", input);
	app.lookup("sms_getRptSearch").setParameters("offset", offset);
	app.lookup("sms_getRptSearch").setParameters("limit", RowCount);
	app.lookup("sms_getRptSearch").send();
}

/*
 * 서치 인풋에서 search 이벤트 발생 시 호출.
 * Searchinput의 enter키 또는 검색버튼을 클릭하여 인풋의 값이 Search될때 발생하는 이벤트
 */
function onSearchInputSearch(/* cpr.events.CUIEvent */ e){
	/** 
	 * @type cpr.controls.SearchInput
	 */
	var searchInput = e.control;
	sendSearchRequest();
}


/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_getRptSearchSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_getRptSearch = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	var totalCount = app.lookup("totalCount").getValue("Count");
	var pageIndexer = app.lookup("pageIndex");
	
	if(result == 0){
		SearchFlag = true;
		AttrFlag = false;
		
		if(totalCount == 0){
			alert('자료가 없습니다.');
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
	}
}


/*
 * '새로고침' 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick2(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var pageIndexer = app.lookup("pageIndex");
	pageIndexer.currentPageIndex = 1; 
	sendRptListRequest();
}

function gridColorWhite(cellIndex){
	var grd = app.lookup("grd1");
	var voCell = grd.detail.getColumn(cellIndex);
//	var voExpress = "switch(getValue(\"color\")){\n\tcase \"Red\" : \"red\"\n\tdefault : \"transparent\"\n}";
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
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_setTreeSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_setTree = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	if(result == 0){
//		console.log(app.lookup("ds_List").getRowDataRanged());

		var dsAttrTree = app.lookup("ds_List");
		dataManager.setDsAttrTree(dsAttrTree);

		app.lookup("tre1").redraw();
	} else{
		alert("tree 갱신 실패");
		return;
	}
}


function sendAttrRptRequest(){
	var pageIndexer = app.lookup("pageIndex");
	if(!AttrFlag){
		NewSearch = true;
		pageIndexer.currentPageIndex = 1;
	} else {
		NewSearch = false;
	}
	
	var pageIdx = pageIndexer.currentPageIndex;
	var offset = (pageIdx -1) * RowCount;
	
	app.lookup("ds_rptList").clear();
	app.lookup("sms_rptAttrSearch").setParameters("offset", offset);
	app.lookup("sms_rptAttrSearch").setParameters("limit", RowCount);
	app.lookup("sms_rptAttrSearch").send();
}




/*
 * 서브미션에서 submit-done 이벤트 발생 시 호출.
 * 응답처리가 모두 종료되면 발생합니다.
 */
function onSms_rptAttrSearchSubmitDone(/* cpr.events.CSubmissionEvent */ e){
	/** 
	 * @type cpr.protocols.Submission
	 */
	var sms_rptAttrSearch = e.control;
	var result = app.lookup("Result").getString("ResultCode");
	var totalCount = app.lookup("totalCount").getValue("Count");
	var pageIndexer = app.lookup("pageIndex");
	
	
	if(result == 0) {
		SearchFlag = false;
		AttrFlag = true;
		
		if(totalCount == 0){
			alert("자료가 없습니다.");
			return;
		}
		
		AllGridColorWhite();
		gridColorImpact(4);
		
		if(NewSearch){  // 새로 검색한 거면 1페이지부터
			setPaging(Number(totalCount), 1, RowCount, 5);
		} else {
			setPaging(Number(totalCount), pageIndexer.currentPageIndex, RowCount, 5);
		}
		app.getContainer().redraw();
		return;
		
		// 이제 서버단 짜기만 하면 됨
	}
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
	var pageIndexer = app.lookup("pageIndex");
	pageIndexer.currentPageIndex = 1;
	
	var attrValue = tre1.getSelectionLast().value;
	var attrCategory = 999;
//	console.log("파싱전 : "+ tre1.getSelectionLast().value);
	attrValue = attrValue.split("-")
	if(attrValue[1] == null){
		attrValue = attrValue[0]-1;
		attrCategory = 0; // 업무속성1로 검색
	} else {
		attrValue = attrValue[1];
		attrCategory = 1; // 업무속성2로 검색
	}
//	console.log("파싱 후 : " +attrValue);
	app.lookup("dm_attr").setValue("attrValue", attrValue);
	app.lookup("dm_attr").setValue("attrCategory", attrCategory);
	
	sendAttrRptRequest();
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
	var paramValue = row.getValue("rpt_idx");
	
//	console.log("선택한 셀 확인 : " + paramValue);
	app.getRootAppInstance().openDialog("app/Bsmg/bm_rptView",{
		width : 1000,
		height : 800,
	}, function(dialog){
		dialog.modal = true;
		dialog.headerVisible = true;
		dialog.headerMovable = true;
		dialog.headerTitle = "일일 업무보고 확인";
		dialog.headerClose = true;
		dialog.addEventListener("keyup", function(e){
			if(e.keyCode == 27){ // ESC
				dialog.close();
			}
		});
		dialog.initValue = {
			rpt_idx : paramValue
		};
	}).then(function(returnValue){
		if (returnValue == 1){
			if(AttrFlag){
				sendAttrRptRequest();
			} else if(SearchFlag){
				sendSearchRequest();
			} else {
				sendRptListRequest();
			}
		}
	});
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
//	app.lookup("grd1").redraw(); 속성이 변경된 경우 자동으로 그려진다.
}
