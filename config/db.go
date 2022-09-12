package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type databaseConfig struct {
	driver          string
	host            string
	name            string
	user            string
	password        string
	port            int
	maxPoolSize     int
	maxOpenCons     int
	maxLifeTimeMins int
}

func Database() databaseConfig {
	return appConfig.db
}

func (c databaseConfig) Driver() string {
	return c.driver
}
func (c databaseConfig) ConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.user, c.password, c.host, c.port, c.name)
}

func (c databaseConfig) MaxPoolSize() int {
	return c.maxPoolSize
}

func (c databaseConfig) MaxOpenCons() int {
	return c.maxOpenCons
}

func (c databaseConfig) MaxLifeTimeMins() int {
	return c.maxLifeTimeMins
}

func newDatabaseConfig() databaseConfig {
	return databaseConfig{
		driver:          readEnvString("DS_DRIVER"),
		host:            readEnvString("DB_HOST"),
		name:            readEnvString("DB_NAME"),
		user:            readEnvString("DB_USER"),
		password:        readEnvString("DB_PASSWORD"),
		port:            readEnvInt("DB_PORT"),
		maxPoolSize:     readEnvInt("DB_MAX_POOL_SIZE"),
		maxLifeTimeMins: readEnvInt("DB_MAX_LIFE_TIME_MINS"),
	}
}
