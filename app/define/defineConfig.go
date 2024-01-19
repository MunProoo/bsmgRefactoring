package define

type Config struct {
	DBConfig       `json:"Database"`
	ScheduleConfig `json:"Cron"`
}

type DBConfig struct {
	DatabaseIP   string `json:"host"`
	DatabasePort string `json:"port"`
	DatabaseID   string `json:"user"`
	DatabasePW   string `json:"password"`
	DatabaseName string `json:"dbname"`
}

type ScheduleConfig struct {
	Spec string `json:"spec"`
}
