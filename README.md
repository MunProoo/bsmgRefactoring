Go로 만든 웹 서비스 프로젝트 입니다.   
http://141.164.54.101   
admin / admin 으로 접속 가능합니다.

* Swagger UI 접근은 http://141.164.54.101/swagger/index.html 입니다. 작업이 안된 API는 추가 예정입니다.
* ref : swaggo/echo-swagger 


## 아키텍처
![시스템 아키텍처 drawio](https://github.com/MunProoo/bsmgRefactoring/assets/52486862/8d7ec53f-3dcb-4c67-89ae-0e84a339ff98)

# 업무보고 관리 서비스 리팩토링
리팩토링의 목적은 다음과 같습니다.
- 클린 아키텍처 적용
- 모듈화
- goroutine 및 channel 사용
- JWT 사용
- Docker를 이용한 배포 경험

## 메뉴 구성
<img width="945" alt="image" src="https://github.com/MunProoo/bsmgRefactoring/assets/52486862/14f8a121-5310-4b8c-a14e-86f7e085c5c9">

## 기능 간략 요약
- 일일 업무보고를 기록합니다.
- 보고대상이 보고확인 기능을 수행하면 더는 수정이 불가합니다.
- 매주 목요일마다 각 주의 일일 업무보고들이 취합됩니다.
- 주간 보고를 Excel로 추출할 수 있습니다.

## CI/CD
CI : Github Actions  
CD : 예정







