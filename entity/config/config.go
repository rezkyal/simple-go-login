package config

type Config struct {
	Database Database `json:"database"`
	Redis    Redis    `json:"redis"`
	Password Password `json:"password"`
	Token    Token    `json:"token"`
	TTL      TTL      `json:"ttl"`
}

type Database struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
}

type Redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type Password struct {
	Cost int `json:"cost"`
}

type Token struct {
	Lifespan int64  `json:"lifespan"`
	Secret   string `json:"secret"`
}

type TTL struct {
	GetUserData int64 `json:"GetUserData"`
}
