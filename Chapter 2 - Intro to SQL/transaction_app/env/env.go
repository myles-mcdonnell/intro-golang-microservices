package env

import (
	"fmt"
	"github.com/myles-mcdonnell/logrusx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	PostgresConfig struct {
		Host                 string
		Port                 int
		Db                   string
		User                 string
		Password             string
		Ssl                  bool
		ApplySchemaMigration bool
		ResetDB              bool
		MigrationPath        string
	}

	Config struct {
		PrettyLogOutput bool
		LogLevel        logrusx.Level
		DbConfig        PostgresConfig
	}
)

const (
	PRETTY_LOG_OUTPUT      string = "PRETTY_LOG_OUTPUT"
	LOG_LEVEL              string = "LOG_LEVEL"
	DB_HOST                string = "DB_HOST"
	DB_PORT                string = "DB_PORT"
	DB_NAME                string = "DB_NAME"
	DB_USERNAME            string = "DB_USERNAME"
	DB_PWD                 string = "DB_PWD"
	RESET_DB               string = "RESET_DB"
	APPLY_SCHEMA_MIGRATION string = "APPLY_SCHEMA_MIGRATION"
	DB_SSL                 string = "DB_SSL"
)

func (config *PostgresConfig) Address() string {
	address := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=", config.User, config.Password, config.Host, config.Port, config.Db)
	if !config.Ssl {
		address += "disable"
	} //defaults to required

	return address
}

func Parse() (*Config, error) {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")
	viper.SetDefault(PRETTY_LOG_OUTPUT, true)
	viper.SetDefault(LOG_LEVEL, "DEBUG")
	viper.SetDefault(DB_HOST, "localhost")
	viper.SetDefault(DB_PORT, 5433)
	viper.SetDefault(DB_NAME, "postgres")
	viper.SetDefault(DB_USERNAME, "postgres")
	viper.SetDefault(DB_PWD, "postgres")
	viper.SetDefault(APPLY_SCHEMA_MIGRATION, true)
	viper.SetDefault(RESET_DB, true)

	rawLogLevel := viper.GetString(LOG_LEVEL)
	logLevel, err := logrusx.ParseLevel(rawLogLevel)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse %v", rawLogLevel)
	}

	resetDb := viper.GetBool(RESET_DB)

	return &Config{
		PrettyLogOutput: viper.GetBool(PRETTY_LOG_OUTPUT),
		LogLevel:        *logLevel,
		DbConfig: PostgresConfig{
			Host:                 viper.GetString(DB_HOST),
			Port:                 viper.GetInt(DB_PORT),
			Db:                   viper.GetString(DB_NAME),
			User:                 viper.GetString(DB_USERNAME),
			Password:             viper.GetString(DB_PWD),
			Ssl:                  viper.GetBool(DB_SSL),
			ApplySchemaMigration: viper.GetBool(APPLY_SCHEMA_MIGRATION),
			MigrationPath:        "file://db/migrations",
			ResetDB:              resetDb,
		}}, nil
}
