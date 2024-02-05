package postgres

// Config is POSTGRES structure config.
type Config struct {
	Host string `json:"host" env:"POSTGRES_HOST" env-default:"0.0.0.0"`
	Port string `json:"port" env:"POSTGRES_PORT"`

	User     string `json:"user" env:"POSTGRES_USER"`
	Password string `json:"password" env:"POSTGRES_PASSWORD"`

	DBName string `json:"db_name" env:"POSTGRES_DB_NAME"`
}
