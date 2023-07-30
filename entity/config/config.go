package config

type Config struct {
	Database Database `json:"database"`
	Password Password `json:"password"`
	Token    Token    `json:"token"`
}

type Database struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
}

type Password struct {
	Cost int `json:"cost"`
}

type Token struct {
	Lifespan int64  `json:"lifespan"`
	Secret   string `json:"secret"`
}
