package config

var (
	PORT              = ":9000"
	DATABASE_USER     = "postgres"
	DATABASE_NAME     = "postgres"
	DATABASE_PASSWORD = "kartikeya"
	DATABASE_PORT     = "5432"
	HOST              = "localhost"
	DATABASE_URL      = "host=" + HOST + " user=" + DATABASE_USER + " password=" + DATABASE_PASSWORD + " dbname=" + DATABASE_NAME + " port=" + DATABASE_PORT + " sslmode=disable"
)
