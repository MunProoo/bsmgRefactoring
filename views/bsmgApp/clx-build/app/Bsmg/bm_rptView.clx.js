/*
 * App URI: app/Bsmg/bm_rptView
 * Source Location: app/Bsmg/bm_rptView.clx
 *
 * This file was generated by eXbuilder6 compiler, Don't edit manually.
 */
(function(){
	var app = new cpr.core.App("app/Bsmg/bm_rptView", {
		onPrepare: function(loader){
		},
		onCreate: function(/* cpr.core.AppInstance */ app, exports){
			var linker = {};
			// Start - User Script
			/************************************************
			 * bm_rptView.js
			 * Created at 2022. 5. 27. 오전 9:43:58.
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
				var rpt_idx = initValue.rpt_idx;
			//	console.log("rptIdx : " +rpt_idx);
				app.lookup("dm_rptIdx").setValue("rpt_idx", Number(rpt_idx));
				
				app.lookup("sms_getRptInfo").send();
				app.lookup("sms_setAttr").send();
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
					dateFormat();
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
				} else{
					alert("속성 갱신 실패");
					return;
				}
			}
			
			function setAttr(){ // 원본으로 돌리기
				var lcb = app.lookup("lcb1");
				var src = app.lookup("dm_reportInfoSrc");
				lcb.selectItemByLabel(src.getString("rpt_attr2"));
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
					app.lookup("sms_putShcedule").send();
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
					var mem_rank = app.lookup("dm_memberInfo").getString("mem_rank")
					var mem_name = app.lookup("dm_memberInfo").getString("mem_name");	
					var rpt_reporter = app.lookup("dm_reportInfo").getString("rpt_reporter");
					var rpt_toRpt = app.lookup("dm_reportInfo").getString("rpt_toRpt");
					var rpt_confirm = app.lookup("dm_reportInfo").getString("rpt_confirm");
					
					if(mem_name == rpt_reporter && rpt_confirm == 'false'){
						app.lookup("update").visible = true;
						app.lookup("cancel").visible = true;
						app.lookup("delete").visible = true;
					} else if(mem_name == rpt_toRpt || (mem_rank == '관리자' || mem_rank == '연구소장' || mem_rank == '부소장' )){
						app.lookup("confirm").visible = true;
					} 
					app.getContainer().redraw();
				} else {
					alert("세션이 끊어졌습니다.");
					app.close();
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
				}
			};
			// End - User Script
			
			// Header
			var dataSet_1 = new cpr.data.DataSet("ds_schedule");
			dataSet_1.parseData({
				"columns" : [{"name": "sc_content"}]
			});
			app.register(dataSet_1);
			
			var dataSet_2 = new cpr.data.DataSet("ds_List");
			dataSet_2.parseData({
				"columns" : [
					{"name": "label"},
					{"name": "value"},
					{"name": "parent"}
				]
			});
			app.register(dataSet_2);
			
			var dataSet_3 = new cpr.data.DataSet("ds_scheduleSrc");
			dataSet_3.parseData({
				"columns" : [{"name": "sc_content"}]
			});
			app.register(dataSet_3);
			var dataMap_1 = new cpr.data.DataMap("dm_reportInfo");
			dataMap_1.parseData({
				"columns" : [
					{"name": "rpt_idx"},
					{"name": "rpt_reporter"},
					{"name": "rpt_date"},
					{"name": "rpt_toRpt"},
					{"name": "rpt_ref"},
					{"name": "rpt_title"},
					{"name": "rpt_content"},
					{"name": "rpt_etc"},
					{"name": "rpt_attr1"},
					{"name": "rpt_attr2"},
					{
						"name": "rpt_confirm",
						"dataType": "string"
					}
				]
			});
			app.register(dataMap_1);
			
			var dataMap_2 = new cpr.data.DataMap("Result");
			dataMap_2.parseData({
				"columns" : [{"name": "ResultCode"}]
			});
			app.register(dataMap_2);
			
			var dataMap_3 = new cpr.data.DataMap("dm_rptIdx");
			dataMap_3.parseData({
				"columns" : [{"name": "rpt_idx"}]
			});
			app.register(dataMap_3);
			
			var dataMap_4 = new cpr.data.DataMap("dm_reportInfoSrc");
			dataMap_4.parseData({
				"columns" : [
					{"name": "rpt_idx"},
					{"name": "rpt_reporter"},
					{"name": "rpt_date"},
					{"name": "rpt_toRpt"},
					{"name": "rpt_ref"},
					{"name": "rpt_title"},
					{"name": "rpt_content"},
					{"name": "rpt_etc"},
					{"name": "rpt_attr1"},
					{"name": "rpt_attr2"}
				]
			});
			app.register(dataMap_4);
			
			var dataMap_5 = new cpr.data.DataMap("dm_memberInfo");
			dataMap_5.parseData({
				"columns" : [
					{"name": "mem_id"},
					{"name": "mem_pw"},
					{"name": "mem_name"},
					{"name": "mem_rank"},
					{"name": "mem_part"}
				]
			});
			app.register(dataMap_5);
			var submission_1 = new cpr.protocols.Submission("sms_getRptInfo");
			submission_1.async = false;
			submission_1.method = "get";
			submission_1.action = "/bsmg/report/reportInfo";
			submission_1.addRequestData(dataMap_3);
			submission_1.addResponseData(dataMap_2, false);
			submission_1.addResponseData(dataMap_1, false);
			if(typeof onSms_getRptInfoSubmitDone == "function") {
				submission_1.addEventListener("submit-done", onSms_getRptInfoSubmitDone);
			}
			app.register(submission_1);
			
			var submission_2 = new cpr.protocols.Submission("sms_getRptSchedule");
			submission_2.async = false;
			submission_2.method = "get";
			submission_2.action = "/bsmg/report/getSchdule";
			submission_2.addRequestData(dataMap_3);
			submission_2.addResponseData(dataSet_1, false);
			submission_2.addResponseData(dataMap_2, false);
			if(typeof onSms_getRptScheduleSubmitDone == "function") {
				submission_2.addEventListener("submit-done", onSms_getRptScheduleSubmitDone);
			}
			app.register(submission_2);
			
			var submission_3 = new cpr.protocols.Submission("sms_setAttr");
			submission_3.async = false;
			submission_3.method = "get";
			submission_3.action = "/bsmg/setting/attrTree";
			submission_3.addResponseData(dataSet_2, false);
			submission_3.addResponseData(dataMap_2, false);
			if(typeof onSms_setAttrSubmitDone == "function") {
				submission_3.addEventListener("submit-done", onSms_setAttrSubmitDone);
			}
			app.register(submission_3);
			
			var submission_4 = new cpr.protocols.Submission("sms_putRpt");
			submission_4.async = false;
			submission_4.method = "put";
			submission_4.action = "/bsmg/report/putRpt";
			submission_4.mediaType = "application/json";
			submission_4.addRequestData(dataMap_1);
			submission_4.addResponseData(dataMap_2, false);
			if(typeof onSms_putRptSubmitDone == "function") {
				submission_4.addEventListener("submit-done", onSms_putRptSubmitDone);
			}
			app.register(submission_4);
			
			var submission_5 = new cpr.protocols.Submission("sms_putShcedule");
			submission_5.async = false;
			submission_5.method = "put";
			submission_5.action = "/bsmg/report/putSchedule";
			submission_5.mediaType = "application/json";
			submission_5.addRequestData(dataMap_3);
			submission_5.addRequestData(dataSet_1, cpr.protocols.PayloadType.all);
			submission_5.addResponseData(dataMap_2, false);
			if(typeof onSms_putShceduleSubmitDone == "function") {
				submission_5.addEventListener("submit-done", onSms_putShceduleSubmitDone);
			}
			if(typeof onSms_putShceduleBeforeSubmit == "function") {
				submission_5.addEventListener("before-submit", onSms_putShceduleBeforeSubmit);
			}
			app.register(submission_5);
			
			var submission_6 = new cpr.protocols.Submission("sms_deleteRpt");
			submission_6.method = "delete";
			submission_6.action = "/bsmg/report/deleteRpt";
			submission_6.addRequestData(dataMap_3);
			submission_6.addResponseData(dataMap_2, false);
			if(typeof onSms_deleteRptSubmitDone == "function") {
				submission_6.addEventListener("submit-done", onSms_deleteRptSubmitDone);
			}
			app.register(submission_6);
			
			var submission_7 = new cpr.protocols.Submission("sms_chkLogin");
			submission_7.method = "get";
			submission_7.action = "/bsmg/login/chkLogin";
			submission_7.addResponseData(dataMap_2, false);
			submission_7.addResponseData(dataMap_5, false);
			if(typeof onSms_chkLoginSubmitDone == "function") {
				submission_7.addEventListener("submit-done", onSms_chkLoginSubmitDone);
			}
			app.register(submission_7);
			
			var submission_8 = new cpr.protocols.Submission("sms_confirmRpt");
			submission_8.method = "get";
			submission_8.action = "/bsmg/report/confirmRpt";
			submission_8.addRequestData(dataMap_3);
			submission_8.addResponseData(dataMap_2, false);
			if(typeof onSms_confirmRptSubmitDone == "function") {
				submission_8.addEventListener("submit-done", onSms_confirmRptSubmitDone);
			}
			app.register(submission_8);
			
			app.supportMedia("all and (min-width: 1024px)", "default");
			app.supportMedia("all and (min-width: 500px) and (max-width: 1023px)", "tablet");
			app.supportMedia("all and (max-width: 499px)", "mobile");
			
			// Configure root container
			var container = app.getContainer();
			container.style.css({
				"background-size" : "auto",
				"background-origin" : "padding-box",
				"width" : "100%",
				"top" : "0px",
				"height" : "100%",
				"left" : "0px"
			});
			
			// Layout
			var xYLayout_1 = new cpr.controls.layouts.XYLayout();
			container.setLayout(xYLayout_1);
			
			// UI Configuration
			var group_1 = new cpr.controls.Container("gr1");
			group_1.readOnly = true;
			group_1.style.setClasses(["cl-form-group"]);
			// Layout
			var formLayout_1 = new cpr.controls.layouts.FormLayout();
			formLayout_1.topMargin = "0px";
			formLayout_1.rightMargin = "0px";
			formLayout_1.bottomMargin = "0px";
			formLayout_1.leftMargin = "0px";
			formLayout_1.horizontalSpacing = "0px";
			formLayout_1.verticalSpacing = "0px";
			formLayout_1.horizontalSeparatorWidth = 1;
			formLayout_1.verticalSeparatorWidth = 1;
			formLayout_1.setColumns(["100px", "1fr"]);
			formLayout_1.setUseColumnShade(0, true);
			formLayout_1.setRows(["100px", "35px", "35px", "285px", "35px"]);
			group_1.setLayout(formLayout_1);
			(function(container){
				var output_1 = new cpr.controls.Output();
				output_1.value = "주요 일정";
				output_1.style.css({
					"font-weight" : "bolder",
					"text-align" : "center"
				});
				container.addChild(output_1, {
					"colIndex": 0,
					"rowIndex": 0
				});
				var group_2 = new cpr.controls.Container();
				group_2.style.setClasses(["cl-form-group"]);
				// Layout
				var formLayout_2 = new cpr.controls.layouts.FormLayout();
				formLayout_2.topMargin = "0px";
				formLayout_2.rightMargin = "0px";
				formLayout_2.bottomMargin = "0px";
				formLayout_2.leftMargin = "0px";
				formLayout_2.horizontalSpacing = "0px";
				formLayout_2.verticalSpacing = "0px";
				formLayout_2.horizontalSeparatorWidth = 1;
				formLayout_2.verticalSeparatorWidth = 1;
				formLayout_2.setColumns(["1fr", "80px"]);
				formLayout_2.setUseColumnShade(0, true);
				formLayout_2.setRows(["1fr"]);
				group_2.setLayout(formLayout_2);
				(function(container){
					var grid_1 = new cpr.controls.Grid("grdSch");
					grid_1.init({
						"dataSet": app.lookup("ds_schedule"),
						"autoRowHeight": "all",
						"resizableColumns": "all",
						"columns": [{"width": "100px"}],
						"header": {
							"rows": [{"height": "24px"}],
							"cells": [{
								"constraint": {"rowIndex": 0, "colIndex": 0},
								"configurator": function(cell){
									cell.targetColumnName = "sc_content";
									cell.filterable = false;
									cell.sortable = false;
									cell.text = "주요 일정";
								}
							}]
						},
						"detail": {
							"rows": [{"height": "24px"}],
							"cells": [{
								"constraint": {"rowIndex": 0, "colIndex": 0},
								"configurator": function(cell){
									cell.columnName = "sc_content";
									cell.control = (function(){
										var inputBox_1 = new cpr.controls.InputBox("sc_content");
										inputBox_1.style.css({
											"padding-left" : "2px"
										});
										inputBox_1.bind("value").toDataColumn("sc_content");
										return inputBox_1;
									})();
								}
							}]
						}
					});
					grid_1.bind("fieldLabel").toDataSet(app.lookup("ds_schedule"), "sc_content", 0);
					grid_1.style.css({
						"background-color" : "#ffffff"
					});
					container.addChild(grid_1, {
						"colIndex": 0,
						"rowIndex": 0
					});
					var group_3 = new cpr.controls.Container();
					group_3.style.setClasses(["cl-form-group"]);
					// Layout
					var formLayout_3 = new cpr.controls.layouts.FormLayout();
					formLayout_3.topMargin = "0px";
					formLayout_3.rightMargin = "0px";
					formLayout_3.bottomMargin = "0px";
					formLayout_3.leftMargin = "0px";
					formLayout_3.horizontalSpacing = "0px";
					formLayout_3.verticalSpacing = "0px";
					formLayout_3.horizontalSeparatorWidth = 1;
					formLayout_3.verticalSeparatorWidth = 1;
					formLayout_3.setColumns(["1fr", "1fr"]);
					formLayout_3.setUseColumnShade(0, true);
					formLayout_3.setRows(["1fr"]);
					group_3.setLayout(formLayout_3);
					(function(container){
						var button_1 = new cpr.controls.Button();
						button_1.value = "-";
						if(typeof onButtonClick3 == "function") {
							button_1.addEventListener("click", onButtonClick3);
						}
						container.addChild(button_1, {
							"colIndex": 0,
							"rowIndex": 0,
							"colSpan": 1,
							"rowSpan": 1
						});
						var button_2 = new cpr.controls.Button();
						button_2.value = "+";
						if(typeof onButtonClick2 == "function") {
							button_2.addEventListener("click", onButtonClick2);
						}
						container.addChild(button_2, {
							"colIndex": 1,
							"rowIndex": 0
						});
					})(group_3);
					container.addChild(group_3, {
						"colIndex": 1,
						"rowIndex": 0
					});
				})(group_2);
				container.addChild(group_2, {
					"colIndex": 1,
					"rowIndex": 0
				});
				var output_2 = new cpr.controls.Output();
				output_2.value = "업무 속성";
				output_2.style.css({
					"font-weight" : "bolder",
					"text-align" : "center"
				});
				container.addChild(output_2, {
					"colIndex": 0,
					"rowIndex": 1
				});
				var linkedComboBox_1 = new cpr.controls.LinkedComboBox("lcb1");
				linkedComboBox_1.preventInput = true;
				linkedComboBox_1.style.css({
					"background-color" : "#ffffff"
				});
				linkedComboBox_1.style.combo.css({
					"background-color" : "#e2e5e4"
				});
				(function(linkedComboBox_1){
					linkedComboBox_1.setItemSet(app.lookup("ds_List"), {
						"label": "label",
						"value": "value",
						"parentValue": "parent"
					});
				})(linkedComboBox_1);
				linkedComboBox_1.placeholders = [
				];
				container.addChild(linkedComboBox_1, {
					"colIndex": 1,
					"rowIndex": 1
				});
				var output_3 = new cpr.controls.Output();
				output_3.value = "주요업무제목";
				output_3.style.css({
					"font-weight" : "bolder",
					"text-align" : "center"
				});
				container.addChild(output_3, {
					"colIndex": 0,
					"rowIndex": 2
				});
				var inputBox_2 = new cpr.controls.InputBox("rpt_title");
				inputBox_2.placeholder = "제목을 입력하세요.";
				inputBox_2.style.css({
					"padding-left" : "2px"
				});
				inputBox_2.bind("value").toDataMap(app.lookup("dm_reportInfo"), "rpt_title");
				container.addChild(inputBox_2, {
					"colIndex": 1,
					"rowIndex": 2
				});
				var output_4 = new cpr.controls.Output();
				output_4.value = "주요업무 내용";
				output_4.style.css({
					"font-weight" : "bolder",
					"text-align" : "center"
				});
				container.addChild(output_4, {
					"colIndex": 0,
					"rowIndex": 3
				});
				var output_5 = new cpr.controls.Output();
				output_5.value = "기타 특이사항";
				output_5.style.css({
					"font-weight" : "bolder",
					"text-align" : "center"
				});
				container.addChild(output_5, {
					"colIndex": 0,
					"rowIndex": 4
				});
				var inputBox_3 = new cpr.controls.InputBox("rpt_etc");
				inputBox_3.style.css({
					"padding-left" : "2px"
				});
				inputBox_3.bind("value").toDataMap(app.lookup("dm_reportInfo"), "rpt_etc");
				container.addChild(inputBox_3, {
					"colIndex": 1,
					"rowIndex": 4
				});
				var textArea_1 = new cpr.controls.TextArea("rpt_content");
				textArea_1.placeholder = "주요업무 내용을 입력하세요.";
				textArea_1.style.css({
					"background-repeat" : "no-repeat",
					"background-size" : "contain",
					"padding-left" : "2px",
					"background-image" : "none",
					"background-position" : "center"
				});
				textArea_1.bind("value").toDataMap(app.lookup("dm_reportInfo"), "rpt_content");
				container.addChild(textArea_1, {
					"colIndex": 1,
					"rowIndex": 3
				});
			})(group_1);
			container.addChild(group_1, {
				"top": "89px",
				"left": "10px",
				"width": "820px",
				"height": "494px"
			});
			
			var button_3 = new cpr.controls.Button("update");
			button_3.visible = false;
			button_3.value = "수정";
			button_3.style.css({
				"background-color" : "#52c183",
				"color" : "#ffffff",
				"background-image" : "none"
			});
			if(typeof onButtonClick4 == "function") {
				button_3.addEventListener("click", onButtonClick4);
			}
			container.addChild(button_3, {
				"top": "593px",
				"left": "48px",
				"width": "149px",
				"height": "43px"
			});
			
			var button_4 = new cpr.controls.Button();
			button_4.value = "닫기";
			if(typeof onButtonClick == "function") {
				button_4.addEventListener("click", onButtonClick);
			}
			container.addChild(button_4, {
				"top": "593px",
				"left": "648px",
				"width": "149px",
				"height": "43px"
			});
			
			var group_4 = new cpr.controls.Container();
			group_4.style.setClasses(["cl-form-group"]);
			// Layout
			var formLayout_4 = new cpr.controls.layouts.FormLayout();
			formLayout_4.topMargin = "0px";
			formLayout_4.rightMargin = "0px";
			formLayout_4.bottomMargin = "0px";
			formLayout_4.leftMargin = "0px";
			formLayout_4.horizontalSpacing = "0px";
			formLayout_4.verticalSpacing = "0px";
			formLayout_4.horizontalSeparatorWidth = 1;
			formLayout_4.verticalSeparatorWidth = 1;
			formLayout_4.setColumns(["100px", "1fr", "100px", "1fr"]);
			formLayout_4.setUseColumnShade(0, true);
			formLayout_4.setUseColumnShade(2, true);
			formLayout_4.setRows(["1fr", "1fr"]);
			formLayout_4.setRowMinHeight(0, 10);
			group_4.setLayout(formLayout_4);
			(function(container){
				var output_6 = new cpr.controls.Output();
				output_6.value = "구분";
				output_6.style.css({
					"font-weight" : "bolder",
					"text-align" : "center"
				});
				container.addChild(output_6, {
					"colIndex": 0,
					"rowIndex": 0
				});
				var output_7 = new cpr.controls.Output();
				output_7.value = "보고일자";
				output_7.style.css({
					"font-weight" : "bolder",
					"text-align" : "center"
				});
				container.addChild(output_7, {
					"colIndex": 2,
					"rowIndex": 0
				});
				var output_8 = new cpr.controls.Output();
				output_8.readOnly = true;
				output_8.value = "일일 업무보고";
				output_8.style.css({
					"padding-left" : "2px",
					"text-align" : "left"
				});
				container.addChild(output_8, {
					"colIndex": 1,
					"rowIndex": 0
				});
				var group_5 = new cpr.controls.Container();
				group_5.style.setClasses(["cl-form-group"]);
				// Layout
				var formLayout_5 = new cpr.controls.layouts.FormLayout();
				formLayout_5.topMargin = "0px";
				formLayout_5.rightMargin = "0px";
				formLayout_5.bottomMargin = "0px";
				formLayout_5.leftMargin = "0px";
				formLayout_5.horizontalSpacing = "0px";
				formLayout_5.verticalSpacing = "0px";
				formLayout_5.horizontalSeparatorWidth = 1;
				formLayout_5.verticalSeparatorWidth = 1;
				formLayout_5.setColumns(["100px", "1fr", "100px", "1fr", "100px", "1fr"]);
				formLayout_5.setUseColumnShade(0, true);
				formLayout_5.setUseColumnShade(2, true);
				formLayout_5.setUseColumnShade(4, true);
				formLayout_5.setRows(["1fr"]);
				group_5.setLayout(formLayout_5);
				(function(container){
					var output_9 = new cpr.controls.Output();
					output_9.value = "보고대상";
					output_9.style.css({
						"font-weight" : "bold",
						"text-align" : "center"
					});
					container.addChild(output_9, {
						"colIndex": 2,
						"rowIndex": 0
					});
					var output_10 = new cpr.controls.Output("toRpt");
					output_10.readOnly = true;
					output_10.style.css({
						"background-color" : "#f8e4e4",
						"padding-left" : "2px",
						"text-align" : "left"
					});
					output_10.bind("value").toDataMap(app.lookup("dm_reportInfo"), "rpt_toRpt");
					container.addChild(output_10, {
						"colIndex": 3,
						"rowIndex": 0
					});
					var output_11 = new cpr.controls.Output();
					output_11.value = "참조대상";
					output_11.style.css({
						"font-weight" : "bold",
						"text-align" : "center"
					});
					container.addChild(output_11, {
						"colIndex": 4,
						"rowIndex": 0
					});
					var output_12 = new cpr.controls.Output("ref");
					output_12.readOnly = true;
					output_12.style.css({
						"background-color" : "#e2f1cf",
						"padding-left" : "2px",
						"text-align" : "left"
					});
					output_12.bind("value").toDataMap(app.lookup("dm_reportInfo"), "rpt_ref");
					container.addChild(output_12, {
						"colIndex": 5,
						"rowIndex": 0
					});
					var output_13 = new cpr.controls.Output();
					output_13.value = "보고자";
					output_13.style.css({
						"font-weight" : "bold",
						"text-align" : "center"
					});
					container.addChild(output_13, {
						"colIndex": 0,
						"rowIndex": 0
					});
					var output_14 = new cpr.controls.Output();
					output_14.readOnly = true;
					output_14.style.css({
						"background-color" : "#e4eff8",
						"padding-left" : "2px",
						"text-align" : "left"
					});
					output_14.bind("value").toDataMap(app.lookup("dm_reportInfo"), "rpt_reporter");
					container.addChild(output_14, {
						"colIndex": 1,
						"rowIndex": 0
					});
				})(group_5);
				container.addChild(group_5, {
					"colIndex": 0,
					"rowIndex": 1,
					"colSpan": 4,
					"rowSpan": 1
				});
				var output_15 = new cpr.controls.Output("rpt_date");
				output_15.readOnly = true;
				output_15.style.css({
					"padding-left" : "2px"
				});
				container.addChild(output_15, {
					"colIndex": 3,
					"rowIndex": 0
				});
			})(group_4);
			container.addChild(group_4, {
				"top": "10px",
				"left": "10px",
				"width": "820px",
				"height": "80px"
			});
			
			var button_5 = new cpr.controls.Button("cancel");
			button_5.visible = false;
			button_5.value = "취소";
			button_5.style.css({
				"background-color" : "#63abf3",
				"color" : "#ffffff",
				"background-image" : "none"
			});
			if(typeof onButtonClick5 == "function") {
				button_5.addEventListener("click", onButtonClick5);
			}
			container.addChild(button_5, {
				"top": "593px",
				"left": "248px",
				"width": "149px",
				"height": "43px"
			});
			
			var button_6 = new cpr.controls.Button("delete");
			button_6.visible = false;
			button_6.value = "삭제";
			button_6.style.css({
				"background-color" : "#de004b",
				"color" : "#fff7f7",
				"background-image" : "none"
			});
			if(typeof onButtonClick6 == "function") {
				button_6.addEventListener("click", onButtonClick6);
			}
			container.addChild(button_6, {
				"top": "593px",
				"left": "448px",
				"width": "149px",
				"height": "43px"
			});
			
			var button_7 = new cpr.controls.Button("confirm");
			button_7.visible = false;
			button_7.value = "보고 확인";
			button_7.style.css({
				"border-radius" : "25px",
				"border-right-style" : "dashed",
				"background-color" : "#ff9336",
				"color" : "#ffffff",
				"border-left-style" : "dashed",
				"border-bottom-style" : "dashed",
				"background-image" : "none",
				"border-top-style" : "dashed"
			});
			if(typeof onConfirmClick == "function") {
				button_7.addEventListener("click", onConfirmClick);
			}
			container.addChild(button_7, {
				"top": "593px",
				"left": "345px",
				"width": "149px",
				"height": "43px"
			});
			if(typeof onBodyLoad == "function"){
				app.addEventListener("load", onBodyLoad);
			}
			if(typeof onBodyInit == "function"){
				app.addEventListener("init", onBodyInit);
			}
		}
	});
	app.title = "bm_rptView";
	cpr.core.Platform.INSTANCE.register(app);
})();
