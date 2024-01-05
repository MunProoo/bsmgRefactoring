package define

type Config struct {
	DBConfig       `json:"Database"`
	ScheduleConfig `json:"Cron"`
}

type ScheduleConfig struct {
	Spec string `json:"spec"`
}
