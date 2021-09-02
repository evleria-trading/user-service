package config

type Config struct {
	PostgresUser       string `env:"POSTGRES_USER" envDefault:"postgres"`
	PostgresPass       string `env:"POSTGRES_PASSWORD" envDefault:""`
	PostgresHost       string `env:"POSTGRES_HOST" envDefault:"localhost"`
	PostgresPort       int    `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresDb         string `env:"POSTGRES_DB" envDefault:"users_db"`
	PostgresSSLDisable bool   `env:"POSTGRES_SSL_DISABLE" envDefault:"false"`
}
