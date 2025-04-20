package config

type DBConfig struct {
	Protocol string
	Username string
	Password string
	Host     string //For now give both url and port
	Appname  string
}

type Config struct {
	DB *DBConfig
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Protocol: "mongodb+srv",
			Username: "mongo",
			Password: "mongo",
			Host:     "rewardsrecognition.dpmkpzc.mongodb.net",
			Appname:  "rewards_and_recognition",
		},
	}
}
