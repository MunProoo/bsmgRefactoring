/**
 * Control별 기본 값 및 공통 상수 정의
 */
var CPR_DEFAULTS = {
	controls: {
		/* 각 개별 Control 관련 기본 설정. 속성, Style,... */
		/*
		linkedlistbox: {
			space: "1px",
			direction: "right",
			style: {
				root: {
					customClasses: ["kmu-linkedlist"],
					css: {
					}
				},
				list: {
					customClasses: ["kmu-linkedlist-list"],
					css: {
						border: "1px solid yellow",
						borderRadius: "3px"
					}
				},
				header: {

				},
				item: {

				},
			}
		}
		*/
	},
	protocol: {
		/* Submission 관련 기본 설정. URL, Progressive Image,... */
		/*
		submission : {
			contextPath : "/"
		}
		*/
	},
	data: {
		/* DataSet, DataMap 관련 기본 설정. */
	},
	environment: {
		/* 구동 환경 설정. 로깅, Plugin 관련, 전역이미지, ... */
		/* 활성화 AppInstance가 없을 때 App을 정의하는 JavaScript 객체를 캐시할지 여부를 설정합니다.
		 * false로 설정되면 활성화된 AppInsance가 없을 때 App을 삭제하고 
		 * 다음 App의 요청이 있을 때 App 정의 객체를 서버로 다시 요청합니다.
		 * default는 true입니다.
		 */
		/*
		appcache: false
		*/
		/* 다국어 바인딩 시 다국어 키에 매핑된 다국어 메시지가 없을 때 빈값 대신 요청된 다국어 키를 리턴할지를 설정합니다.
		 * true로 설정되면 요청된 다국어 키 값에 배핑된 다국어 메시지가 없을 때 요청된 다국어 키를 리턴합니다.
		 * default는 false 입니다.
		 */
		/*
		useKeyInsteadOfNullI18NMessage: false
		*/
		/*
		 * 컨트롤 및 그룹에서 브라우저 내장 스크롤바 대신 스타일 적용이 가능한 커스텀 스크롤바 적용 여부를 설정합니다.
		 * 강제로 브라우저 스크롤이 보이도록 스타일 되어있거나 HTMLObject, UIControlShell, HTMLSnippet을 통해 생성되는 스크롤바에 대해서는 지원하지 않습니다.
		 * 실험적인 기능이라 true로 설정했을 때 기존 스타일에 문제가 있을 수 있습니다.
		 * default는 false 입니다.
		 */
		/*
		useCustomScrollbar: false
		*/
		/**
		 * 생성되는 UIControl에 대응되는 HTMLElement의 ID 속성을 UIControl의 uuid를 사용할 지 여부를 설정합니다.
		 * true로 설정되면 UIControl에 대응되는 HTMLElement의 ID 속성이 UIControl의 uuid를 사용하여 매크로를 통한 위협을 방어할 수 있습니다.
		 * false로 설정되면 UIControl에 대응되는 HTMLElement의 ID 속성이 UIContorl의 경로로 생성되어 
		 * Web UI 자동화 테스트툴을 사용하여 테스트할 때 테스트툴이 테스트 시 동일한 HTMLElement를 찾을 수 있도록 지원합니다.
		 * default는 true 입니다.
		 */
		/*
		useControlUUIDasHTMLId: true
		*/
		/**
		 * 가능한 경우, 루트 앱이 브라우저의 기본 스크롤 동작을 활용하도록 합니다.
		 */
		/*
		preferNativeScroll : true
		*/
	},
	variables: {
		/* 전역상수 */
	}

};