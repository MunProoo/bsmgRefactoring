/************************************************
 * userList.js
 * Created at 2022. 5. 23. 오전 9:18:58.
 *
 * @author SW2Team
 ************************************************/

/**
 * UDC 컨트롤이 그리드의 뷰 모드에서 표시할 텍스트를 반환합니다.
 */
exports.getText = function(){
	// TODO: 그리드의 뷰 모드에서 표시할 텍스트를 반환하는 하는 코드를 작성해야 합니다.
	return "";
};


// 잘 모르겠음
exports.enablePageIndexer = function(enable) {
	udcUserList_enablePageIndexer = enable;
	var pageIndex = app.lookup("userListPageIndexer");
	if(udcUserList_enablePageIndexer == true){
		pageIndex.visible = false;
	} else{
		pageIndex.visible = false;
	}
}

// 잘 모르겠음
exports.getCheckedRowIndices = function() {
	var userList = app.lookup("UDC_grdUserList");
	var indices = userList.getCheckRowIndices();
	var result = [];
	indices.forEach(function(idx){
		if(userList.getRowState(idx) != cpr.data.tabledata.RowState.DELETED){
			result.push(idx);
		} else{
			userList.setCheckRowIndex(idx, false);
		}
	});
}

// 선택된 열 ID값 반환
exports.getSelectedID = function() {
	var grdUserList = app.lookup("UDC_grdUserList");
	var row = grdUserList.getSelectedRow();
	if(row) {
		return row.getValue("ID");
	}
	return null;
}

// 선택된 열 값 전부 반환
exports.getSelectedRowDate = function() {
	var grdUserList = app.lookup("UDC_grdUserList");
	var row = grdUserList.getSelectedRow();
	if(row){
		return row.getRowData();
	}
	return null;
}

// 행 이동
exports.moveColumn = function(srcIdx, targetIdx) {
	var grdUserList = app.lookup("UDC_grdUserList");
	grdUserList.moveColumn(srcIdx, targetIdx);
}

exports.deleteColumn = function(indices){
	if (indices==undefined || indices == null ){
		return;
	}
	var gridUserList = app.lookup("UDC_grdUserList");
	indices.forEach(function(index){
		gridUserList.deleteColumn(index);
	});	
};

exports.visibleColumn = function(cellIndex, visible) {
	var gridUserList = app.lookup("UDC_grdUserList");
	gridUserList.columnVisible(cellIndex, visible);			
}

exports.deleteRow = function(checkRow) {
	var userList = app.lookup("UDC_grdUserList");
	if( checkRow >= userList.getRowCount()){
		return;
	}
	userList.deleteRow(checkRow);
	userList.setCheckRowIndex(checkRow, false);
	return;
}

exports.realDeleteRow = function(index) {
	var userList = app.lookup("UserList");	
	userList.realDeleteRow(index);	
	return;
}

exports.deleteUser = function(deleteID) {
	var userList = app.lookup("UDC_grdUserList");
	var getUserInfo = userList.findFirstRow("ID == "+ deleteID);
	
	if (getUserInfo) {
		userList.deleteRow(getUserInfo.getIndex());	
	} 
	return;
}

exports.setUserList = function( /*cpr.data.DataSet*/userDataSet ){
			
	var userList = app.lookup("UserList");
	var userListSrc = app.lookup("UserListSrc");
	
	userList.clear();	
	userListSrc.clear();		
	
	userDataSet.copyToDataSet(userList);	
	userDataSet.copyToDataSet(userListSrc);
		
	userList.setRowStateAll(cpr.data.tabledata.RowState.UNCHANGED);
	
	var grdUserList = app.lookup("UDC_grdUserList");	
	grdUserList.redraw();
}

exports.setFilter = function(category,keyword){
	if (dataManager.getOemVersion() == OEM_OMAN_TERMINAL_UPDATEUSER) {
		if (category != "id" && category != "name" && category != "UpdateX" && category != "UpdateO"){
			return;
		}
	}
	else {
		if (category != "id" && category != "name"){return;}	
	}
	
	
	var userList = app.lookup("UserList");
	userList.clear();
	
	keyword = keyword.toLowerCase();
	
	var userListSrc = app.lookup("UserListSrc");
	var count = userListSrc.getRowCount();
	for( var i = 0; i < count; i++ ){
		var user = userListSrc.getRow(i);
		var srcData;
		
		if (category == "UpdateX" || category == "UpdateO") {
			if (category == "UpdateX" && user.getValue("UpdateFlag") == 0){
				userList.addRowData(user.getRowData());
			} else if (category == "UpdateO" && user.getValue("UpdateFlag") == 1){
				userList.addRowData(user.getRowData());
			}
			continue;
		}
		
		if (category == "id") {
			srcData = user.getValue("ID");
		} else if (category == "name") {
			srcData = user.getValue("Name");
		}
		srcData = srcData.toLowerCase();
		if (srcData.indexOf(keyword) != -1) {
  			userList.addRowData(user.getRowData());
		}		
	}
	userList.commit();
}
exports.clearFilter = function(){
	var userList = app.lookup("UserList");
	var userListSrc = app.lookup("UserListSrc");	
	userList.clear();
	userListSrc.copyToDataSet(userList);
	userList.commit();
}

exports.clearUserList = function(  ){
			
	var pageIndex = app.lookup("userListPageIndexer");
	pageIndex.totalRowCount = 0;
	pageIndex.visible = false;	
	pageIndex.redraw();
	
	var userListSet = app.lookup("UserList");
	userListSet.clear();
				
	var userList = app.lookup("UDC_grdUserList");	
	userList.redraw();
}

exports.setUserListRows = function( /*cpr.data.RowConfigInfo[]*/userData ){
			
	var userListSet = app.lookup("UserList"); 
	userListSet.clear();	
	userListSet.build(userData);	
	userListSet.setRowStateAll(cpr.data.tabledata.RowState.UNCHANGED);
	
	var userList = app.lookup("UDC_grdUserList");	
	userList.redraw();	
}

exports.getUserID = function( index ){
	
	var userList = app.lookup("UDC_grdUserList");
	var userID = userList.getRow(index).getString("ID");
	return userID;
}

exports.getRowData = function( index ){
	
	var userList = app.lookup("UDC_grdUserList");
	return userList.getRow(index).getRowData();	
}

exports.updateUserInfo = function( userInfoData ){
	
	var groupList = dataManager.getGroup();
	if( groupList && groupList.getRowCount()>0){
		var cmbGroup = app.lookup("userListGrid_cmbGroup");
		var count = groupList.getRowCount();
		for ( var i = 0; i < count; i++ ){			
			var groupInfo = groupList.getRow(i);						
			cmbGroup.addItem(new cpr.controls.Item(groupInfo.getValue("Name"),groupInfo.getValue("GroupID")));
		}	
	}
	
	var positionList = dataManager.getPositionList();
	if( positionList && positionList.getRowCount()>0){
		var cmbPosition = app.lookup("userListGrid_cmbPosition");
		var count = positionList.getRowCount();
		for ( var i = 0; i < count; i++ ){			
			var positionInfo = positionList.getRow(i);						
			cmbPosition.addItem(new cpr.controls.Item(positionInfo.getValue("Name"),positionInfo.getValue("PositionID")));
		}	
	}
	
	var accessGroupList = dataManager.getAccessGroup();
	if( accessGroupList && accessGroupList.getRowCount()>0){
		var cmbAccessGroup = app.lookup("userListGrid_cmbAccessGroup");
		var count = accessGroupList.getRowCount();
		for ( var i = 0; i < count; i++ ){			
			var accessGroupInfo = accessGroupList.getRow(i);						
			cmbAccessGroup.addItem(new cpr.controls.Item(accessGroupInfo.getValue("Name"),accessGroupInfo.getValue("ID")));
		}
	}
	
	
	var dsUserList = app.lookup("UserList");
	var userInfo = dsUserList.findFirstRow("ID == '"+userInfoData.ID+"'");
	if(userInfo){
		
		userInfo.setRowData(userInfoData);
		
		var AuthType = userInfo.getValue("AuthInfo").split(',');
				
		var setCount = 0;
		var andAuth = "";
		for( var idx=0; idx < AuthType[7]; idx++ ){		
			if(AuthType[idx]!="0"){
				andAuth += getAuthTypeString( parseInt(AuthType[idx],10))+" ";
				setCount++;
			}	
		}
		var orAuth = "";	
		for( var idx=AuthType[7]; idx< AuthType.length-1; idx++ ){		
			if(AuthType[idx]!="0"){
				orAuth += getAuthTypeString( parseInt(AuthType[idx],10))+" ";
				setCount++;
			}
		}
			
		if( setCount > 1 ){
			userInfo.setValue("AuthInfo",andAuth+"/ "+orAuth);
		} else {
			userInfo.setValue("AuthInfo",andAuth+orAuth);
		}		
	}	
}

exports.getRow = function( index ){
	
	var userList = app.lookup("UDC_grdUserList");
	return userList.getRow(index);	
}

exports.getRowState = function( index ){
	
	var userList = app.lookup("UDC_grdUserList");
	return userList.getRowState(index);	
}

exports.setRowState = function(index, state){
	var dsUserList = app.lookup("UserList");
	dsUserList.setRowState(index, state);
}
/*
 * make bisangoo
 */
exports.setUnCheckAll = function(idx,checked){
	var userList = app.lookup("UDC_grdUserList");
	var indices = userList.getCheckRowIndices();
	
	indices.forEach(function(idx){
		userList.setCheckRowIndex(idx, false);		
	});
}

exports.setCheckAll = function(checked){
	var userList = app.lookup("UDC_grdUserList");
	var total = userList.getRowCount();
	
	for (var i = 0; i < total; i++) {
		userList.setCheckRowIndex(i, checked);
	}
} 
 

exports.findInnerUserList = function(category,keyword){
	var grdUserList = app.lookup("UDC_grdUserList");
	var user;
	if (category == "Name") {
		user = grdUserList.findFirstRow(category + " == '" + keyword + "'");
	} else if (category == "UniqueID") {
		user = grdUserList.findFirstRow(category + " == '" + keyword + "'");
	}
	
	if (user) {
		var idx = user.getIndex();
		grdUserList.selectRows(idx);
		grdUserList.focusCell(idx, 0);
	} else {
		if (category == "Name") {
			dialogAlert(app, "Waning", dataManager.getString("Str_ErrorNoNameFound"));
		} else if (category == "UniqueID") {
			dialogAlert(app, "Waning", dataManager.getString("Str_ErrorNoUniqueIDFound"));
		}
		
	}
}



/**
 * 사용자 리스트 컨트롤의 페이징 정보를 설정합니다.
 */
exports.setPaging = function( totalCount, currentPageIndex, pageRowCount, viewPageCount ) {
	var pageIndex = app.lookup("userListPageIndexer");
	
	pageIndex.totalRowCount = totalCount;//전체 데이터 수.
	pageIndex.currentPageIndex = currentPageIndex;//현재 선택된 페이지의 인덱스
	pageIndex.pageRowCount = pageRowCount;//한 페이지에 보여 줄 행의 수
	pageIndex.viewPageCount = viewPageCount;// 보여지는 페이지 수(하단 부 인덱스 수)
	
	if( udcUserList_enablePageIndexer == true ){
		if(totalCount == 0) {
			pageIndex.visible = false;
		} else {
			pageIndex.visible = true;
		}
	}else {
		pageIndex.visible = false;
	}
	
	pageIndex.redraw();
}

exports.setPaging = function( totalCount, pageRowCount, viewPageCount ) {
	var pageIndex = app.lookup("userListPageIndexer");
	
	pageIndex.totalRowCount = totalCount;//전체 데이터 수.
	//pageIndex.currentPageIndex = currentPageIndex;//현재 선택된 페이지의 인덱스
	pageIndex.pageRowCount = pageRowCount;//한 페이지에 보여 줄 행의 수
	pageIndex.viewPageCount = viewPageCount;// 보여지는 페이지 수(하단 부 인덱스 수)
		
	if( udcUserList_enablePageIndexer == true ){
		if(totalCount == 0) {
			pageIndex.visible = false;
		} else {
			pageIndex.visible = true;
		}
	}else {
		pageIndex.visible = false;
	}
	pageIndex.redraw();
}

exports.setTotalCount = function(totalCount) {
	
	var pageIndex = app.lookup("userListPageIndexer");
	pageIndex.totalRowCount = totalCount
	
	if( udcUserList_enablePageIndexer == true ){
		if(totalCount == 0) {
			pageIndex.visible = false;
		} else {
			pageIndex.visible = true;
		}
	}else {
		pageIndex.visible = false;
	}
	
	pageIndex.redraw();
}

exports.getCurrentPageIndex = function() {	
	var pageIndex = app.lookup("userListPageIndexer");
	return pageIndex.currentPageIndex
}

exports.setCurrentPageIndex = function(index) {	
	var pageIndex = app.lookup("userListPageIndexer");	
	pageIndex.currentPageIndex = index;	
}

exports.setPageRowCount = function(count) {	
	var pageIndex = app.lookup("userListPageIndexer");	
	pageIndex.pageRowCount = count;	
}

exports.getPageRowCount = function() {	
	var pageIndex = app.lookup("userListPageIndexer");	
	return pageIndex.pageRowCount;	
}

exports.refreshUserList = function(idMap){
	var dsUserList = app.lookup("UserList");
	
	var total = dsUserList.getRowCount();
	for ( var i = 0; i < total; i++){		
		var row = dsUserList.getRow(i);				
		if (row){
			var userID = row.getValue("ID");			
									
			if( idMap.get(userID) != undefined ){
				dsUserList.setRowState(i,cpr.data.tabledata.RowState.DELETED);	
			} else {				
				dsUserList.setRowState(i,cpr.data.tabledata.RowState.UNCHANGED);
			}
		} 
	}
	
	var userList = app.lookup("UDC_grdUserList");
	userList.redraw();
}

/*
 * make bisangoo
 */
exports.getIsCheckedRow = function(rowIndex) {
	var userList = app.lookup("UDC_grdUserList");
	return userList.isCheckedRow(rowIndex);
}

/*
 * make bisangoo
 */
exports.refreshCheckboxStatus = function(idMap) {
	var userList = app.lookup("UDC_grdUserList");
	var total = userList.getRowCount();
	for(var i =0; i < total; i++) {
		var row = userList.getRow(i);
		if(row) {
			var userID = row.getValue("ID");
			if( idMap.get(userID) != undefined ){
				userList.setCheckRowIndex(i, true);	
			} else {				
				userList.setCheckRowIndex(i, false);
			}
		}
	}
	userList.redraw();
}

/*
 * 페이지 인덱서에서 before-selection-change 이벤트 발생 시 호출.
 * Page index를 선택하여 선택된 페이지가 변경되기 전에 발생하는 이벤트. 다음 이벤트로 selection-change를 발생합니다.
 */
function onUserListPageIndexerBeforeSelectionChange(/* cpr.events.CSelectionEvent */ e){
	/** 
	 * @type cpr.controls.PageIndexer
	 */
	var userListPageIndexer = e.control;
	
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
function onUserListPageIndexerSelectionChange(/* cpr.events.CSelectionEvent */ e){
	/** 
	 * @type cpr.controls.PageIndexer
	 */
	var userListPageIndexer = e.control;
	
	var selectionEvent = new cpr.events.CSelectionEvent("pagechange", {
		oldSelection: e.oldSelection,
		newSelection: e.newSelection
	});
	
	app.dispatchEvent(selectionEvent);
	
}


/*
 * 그리드에서 row-dblclick 이벤트 발생 시 호출.
 * detail이 row를 더블클릭 한 경우 발생하는 이벤트.
 */
function onUserListGridRowDblclick(/* cpr.events.CGridEvent */ e){
	/** 
	 * @type cpr.controls.Grid
	 */
	var userListGrid = e.control;
	
	var gridEvent = new cpr.events.CGridEvent("userListDblclick", {
		 row:e.row
	});
	
	app.dispatchEvent(gridEvent);
}


/*
 * 그리드에서 selection-change 이벤트 발생 시 호출.
 * detail의 cell 클릭하여 설정된 selectionunit에 해당되는 단위가 선택될 때 발생하는 이벤트.
 */
function onUDC_grdUserListSelectionChange(/* cpr.events.CSelectionEvent */ e){
	/** 
	 * @type cpr.controls.Grid
	 */
	var uDC_grdUserList = e.control;
	var gridRow = uDC_grdUserList.getRow(e.newSelection[0]);
	
	var selectionEvent = new cpr.events.CSelectionEvent("userListClick", {		
		oldSelection: e.oldSelection,
		newSelection: gridRow
	});
		
	app.dispatchEvent(selectionEvent);
}







