/************************************************
 * DataManager.module.js
 * Created at 2023. 12. 22. 오후 4:48:09.
 *
 * @author MJY
 ************************************************/

exports.id = "DataManager.module.js";
{
	
	var _dataManager = null;
	
	var DataManager = function() {
		
		// 직급
		/** #type cpr.data.DataSet */
		this._rankList = null;
		
		// 부서 
		/** #type cpr.data.DataSet */
		this._partList = null;
		
		// 업무 속성 트리
		/** #type cpr.data.DataSet */
		this._dsAttrTree = null;
	}
	
	DataManager.prototype.setRankList = function( /* cpr.data.DataSet */ dsRankList) {
		this._rankList = dsRankList;
	}
	
	DataManager.prototype.getRankList = function() {
		return this._rankList;
	}
	
	DataManager.prototype.setPartList = function( /* cpr.data.DataSet */ dsPartList) {
		this._partList = dsPartList;
	}
	
	DataManager.prototype.getPartList = function() {
		return this._partList;
	}
	
	DataManager.prototype.setDsAttrTree = function( /*cpr.data.DataSet*/ dsAttrTree) {
		this._dsAttrTree = dsAttrTree;
	}
	
	DataManager.prototype.getDsAttrTree = function() {
		return this._dsAttrTree;
	}
		
//		DataManager.prototype.insertPosition = function(rowPosition) {
//			var columns = Object.keys(rowPosition);
//			console.log(rowPosition[columns[0]]);
//			
//			var positions = this._positionList.findFirstRow("PositionID == " + rowPosition[columns[0]]);
//			
//			if (positions == null) {
//				var insertedRow = this._positionList.addRowData(rowPosition);
//				this._positionList.commit();
//			}
//		
//		}
	
//		DataManager.prototype.updatePosition = function(rowPosition) {
//			var updateRow = this._positionList.findFirstRow("PositionID == " + rowPosition.PositionID);
//			if (updateRow) {
//				updateRow.setRowData(rowPosition)
//				this._positionList.commit();
//			}
//		}
//		
//		DataManager.prototype.deletePosition = function(positionID) {
//			var delRow = this._positionList.findFirstRow("PositionID == " + positionID);
//			if (delRow) {
//				this._positionList.deleteRow(delRow.getIndex());
//				this._positionList.commit();
//			}
//		}
		
	
	globals.getDataManager = function() {
		if (_dataManager == null) _dataManager = new DataManager();
		return _dataManager;
	}
}