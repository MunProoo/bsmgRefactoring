package server

import (
	"BsmgRefactoring/database"
	"BsmgRefactoring/define"
	"BsmgRefactoring/repository"
	"errors"
	"sync"
	"time"

	"github.com/robfig/cron"
)

type ServerProcessor struct {
	DBManager    database.DatabaseManagerInterface
	State        uint16 // 서버의 상태
	Mutex        sync.RWMutex
	WeekRptMaker *cron.Cron // 주간보고 스케쥴러
	Config       *define.Config
	ReqCh        chan interface{}
	ResCh        chan interface{}
}

func InitServer() (server *ServerProcessor) {
	server = &ServerProcessor{}
	server.SetState(define.StateInit)
	// 음 채널링 써보고싶은데
	server.ReqCh = make(chan interface{}, 1000)
	server.ResCh = make(chan interface{}, 1000)
	server.Config = &define.Config{}
	server.LoadConfig()

	return
}

// Server 시작 (서버상태 별 프로세스 실행)
func (server *ServerProcessor) StartServer(repo repository.BsmgRepository) {
	// TODO : DB 연결상태 확인 다른방법 없을까? Ping은 성능에 문제 준다는데
	var err error
	for {
		switch server.State {
		case define.StateInit:
			// err = server.LoadConfig()
			// if err != nil {
			// 	// server.log.Error("Load config Failed", "error", err)
			// 	time.Sleep(100 * time.Millisecond)
			// 	continue
			// }

			// err = server.ConnectDataBase()
			// if err != nil {
			// 	// server.log.Error("Connect Database Failed", "error", err)
			// 	time.Sleep(100 * time.Millisecond)
			// 	continue
			// }

			server.SetState(define.StateConnected)

		case define.StateConnected:
			// go server.DBQuery() // 1.DB 쿼리용 고루틴 생성 (미정)
			// server.CreateCron() // 2.주간보고 스케쥴러 생성

			server.SetState(define.StateRunning)

		case define.StateRunning: // DB 연결중임을 확인
			err = server.IsConnected()
			if err != nil {
				// server.log.Error("DB is not connected", "error", err)
				server.SetState(define.StateDisconnected)
			}
			repo.ConnectDB(server.DBManager)
			time.Sleep(1 * time.Second)

		case define.StateDisconnected: // DB 연결 재시도
			err = server.ConnectDataBase()
			if err != nil {
				// server.log.Error("Connect Database Failed", "error", err)
				time.Sleep(1 * time.Second)
				continue
			}
			server.SetState(define.StateInit)
		default:
			time.Sleep(15 * time.Millisecond)
		}
	}
}

// 서버 상태 Set
func (server *ServerProcessor) SetState(state uint16) {
	server.Mutex.Lock()
	defer server.Mutex.Unlock()
	server.State = state
}

// DB 연결 확인
func (server *ServerProcessor) IsConnected() (err error) {
	// ping날리는 고루틴을 따로 두고 db가 nil인지 아닌지를 판별할지, 아니면 ping을 StartServer에서 날릴지는 고민좀 해보자
	server.Mutex.Lock()
	defer server.Mutex.Unlock()

	if server.DBManager != nil {
		err = server.DBManager.IsConnected()
		return
	}
	err = errors.New("database is not connected")
	return
}

// // DB 쿼리 작업
// func (server *ServerProcessor) DBQuery() {
// 	// 호출하는 곳에서 Mutex Lock 하도록
// 	for {
// 		select {
// 		case msg := <-server.ReqCh:
// 			fmt.Printf("%v \n", msg)
// 		}
// 	}
// }
