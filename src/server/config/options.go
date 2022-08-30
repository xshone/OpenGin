package config

type Server struct {
	Port string `json:"port"`
}

type Database struct {
	Host                 string `json:"host"`
	Port                 string `json:"port"`
	User                 string `json:"user"`
	Password             string `json:"password"`
	ConnMaxLifetimeHours int    `json:"conn_max_lifetime_hours"`
}

type JWT struct {
	SecretKey   string `json:"secret_key"`
	ExpireHours int    `json:"expire_hours"`
}

type RabbitMQ struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	User        string `json:"user"`
	Password    string `json:"password"`
	VirtualHost string `json:"virtual_host"`
}

type Redis struct {
	Address  string `json:"address"`
	Password string `json:"password"`
	Db       int    `json:"db"`
	PoolSize int    `json:"pool_size"`
}

type Options struct {
	Server   *Server   `json:"server"`
	Database *Database `json:"database"`
	JWT      *JWT      `json:"jwt"`
	RabbitMQ *RabbitMQ `json:"rabbit_mq"`
	Redis    *Redis    `json:"redis"`
}

var (
	Settings *Options
)
