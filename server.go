package main

import (
	"BsmgRefactoring/define"
	"fmt"
	"time"
)

// Server 시작 (서버상태 별 프로세스 실행)
func (server *ServerProcessor) StartServer() {
	// TODO : DB 연결상태 확인 다른방법 없을까? Ping은 성능에 문제 준다는데
	var err error
	for {
		switch server.State {
		case define.StateInit:
			err = server.ConnectDataBase()
			if err != nil {
				log.Printf("%v \n ", err)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			server.mutex.Lock()
			server.SetState(define.StateConnected)
			server.mutex.Unlock()

		case define.StateConnected: // DB 쿼리용 고루틴 생성
			go server.DBQuery() //
			server.mutex.Lock()
			server.SetState(define.StateRunning)
			server.mutex.Unlock()

		case define.StateRunning: // DB 연결중임을 확인
			err = server.IsConnected()
			if err != nil {
				log.Printf("%v \n ", err)
				server.mutex.Lock()
				server.SetState(define.StateDisconnected)
				server.mutex.Unlock()
			}
			time.Sleep(1 * time.Second)

		case define.StateDisconnected: // DB 연결 재시도
			err = server.ConnectDataBase()
			if err != nil {
				log.Printf("%v \n ", err)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			server.mutex.Lock()
			server.SetState(define.StateInit)
			server.mutex.Unlock()

		}
	}
}

// 서버 상태 Set
func (server *ServerProcessor) SetState(state uint16) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	server.State = state
}

// DB 연결 확인
func (server *ServerProcessor) IsConnected() (err error) {
	// ping날리는 고루틴을 따로 두고 db가 nil인지 아닌지를 판별할지, 아니면 ping을 StartServer에서 날릴지는 고민좀 해보자
	server.mutex.Lock()
	defer server.mutex.Unlock()

	err = server.dbManager.DBGorm.IsConnected()
	return
}

// DB 쿼리 작업
func (server *ServerProcessor) DBQuery() {
	// 호출하는 곳에서 Mutex Lock 하도록
	for {
		select {
		case msg := <-server.reqCh:
			fmt.Printf("%v \n", msg)
		}
	}
}
