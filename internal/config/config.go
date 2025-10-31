package config

type App struct {
	Name     string
	Env      string
	LogLevel string
}

type SMTP struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Config struct {
	AppConfig  App
	SMTPConfig SMTP
}
