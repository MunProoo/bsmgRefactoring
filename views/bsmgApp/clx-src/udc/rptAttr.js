/************************************************
 * rptAttr.js
 * Created at 2022. 5. 27. 오후 6:09:53.
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


/*
 * "+" 버튼에서 click 이벤트 발생 시 호출.
 * 사용자가 컨트롤을 클릭할 때 발생하는 이벤트.
 */
function onButtonClick5(/* cpr.events.CMouseEvent */ e){
	/** 
	 * @type cpr.controls.Button
	 */
	var button = e.control;
	var vcFormLayout = app.lookup("grdEx");
	
	vcFormLayout.getLayout().insertRows(["30 px"]);
	var idx = vcFormLayout.getLayout().getRows().length;
	var ipbName = "sch"+idx;
	var input = new cpr.controls.InputBox(ipbName);
	
	vcFormLayout.addChild(input, {
		rowIndex:idx-1
	});
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
	var vcFormLayout = app.lookup("grdEx");
	var idx = vcFormLayout.getLayout().getRows().length;
	var child = vcFormLayout.getLayout().removeRows(idx);
	vcFormLayout.removeChild(child);
}