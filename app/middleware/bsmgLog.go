package middleware

import (
	"BsmgRefactoring/utils"
	"fmt"
	"os"
	"time"

	"github.com/inconshreveable/log15"
)

// 사용자 정의 로그 레벨 상수 정의
type LogArg map[string]interface{}

const (
	Debug int = iota + 1
	Info
	Warn // 주의
	Error
	Crit // 심각
)

type LogData struct {
	level int
	data  map[string]interface{}
}

var (
	bsmgLog       log15.Logger
	logFile       *os.File
	curDay        int
	logNamePrefix = "BSMG"

	logCh  chan LogData
	doneCh chan int
)

func InitLog() {

	bsmgLog = log15.New("module", "app/server")
	// 스트림 핸들러 (터미널에 출력)
	streamHandler := log15.StreamHandler(os.Stdout, log15.TerminalFormat())

	logFile, err := utils.CreateLogFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	curDay = time.Now().Day()

	// 파일 핸들러 (프로덕션 환경에서만 중요한 오류 기록 ,,,? Info도 일단은 기록하고 싶음)
	fileHandler := log15.Must.FileHandler(logFile.Name(), log15.LogfmtFormat())

	// Error, Critical 만 적용할 때 사용
	// fileHandler := log15.LvlFilterHandler(
	// 	log15.LvlError,
	// 	log15.Must.FileHandler(logFile.Name(), log15.LogfmtFormat()),
	// )

	// 로거에 여러 핸들러 추가 (디버깅 중에는 모든 로그를 스트림에, 프로덕션 환경에서는 중요한 오류만 파일에)
	bsmgLog.SetHandler(log15.MultiHandler(streamHandler, fileHandler))
	bsmgLog.Info("BSMG Server Start")

	logCh = make(chan LogData, 1000000)
	doneCh = make(chan int, 100)

	go logProcess()
}

func logProcess() {
	periodicTimer := time.NewTimer(60 * time.Second)

	// isRun := true
	for {
		select {
		case logData := <-logCh:
			validateLogFile()

			switch logData.level {
			case Debug:
				bsmgLog.Debug("Debug :", "detail", logData.data)
			case Info:
				bsmgLog.Info("Info :", "detail", logData.data)
			case Warn:
				bsmgLog.Warn("Warn :", "detail", logData.data)
			case Error:
				bsmgLog.Error("Error :", "detail", logData.data)
			case Crit:
				bsmgLog.Crit("Crit :", "detail", logData.data)
			}

		// case <-doneCh:
		// 	isRun = false
		case <-periodicTimer.C:
			// 서버 잘 살아있다는 heartbeat 할까말까
		}

	}

	defer logFile.Close()

}

// 날짜 지나거나 기록 많이쌓이면 로그파일 변경
func validateLogFile() {
	timeInfo := time.Now()
	checkDay := timeInfo.Day()

	if fi, err := logFile.Stat(); err == nil {
		if fi.Size() > 10240000 {
			logFile.Close()
			logFile, _ = utils.CreateLogFile()
			return
		}
	}

	if curDay != checkDay {
		curDay = checkDay
		logFile.Close()
		logFile, _ = utils.CreateLogFile()
		return
	}
}

// Debug 전송
func PrintD(data map[string]interface{}) {
	logCh <- LogData{level: Debug, data: data}
}

// Info 전송
func PrintI(data map[string]interface{}) {
	logCh <- LogData{level: Info, data: data}
}

// Warn 전송
func PrintW(data map[string]interface{}) {
	logCh <- LogData{level: Warn, data: data}
}

// Err 전송
func PrintE(data map[string]interface{}) {
	logCh <- LogData{level: Error, data: data}
}

// Crit 전송
func PrintC(data map[string]interface{}) {
	logCh <- LogData{level: Crit, data: data}
}
