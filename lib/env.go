package lib

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	LogLevel    string `mapstructure:"LOG_LEVEL"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENV"`

	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBType     string `mapstructure:"DB_TYPE"`

	MailClientID     string `mapstructure:"MAIL_CLIENT_ID"`
	MailClientSecret string `mapstructure:"MAIL_CLIENT_SECRET"`
	MailTokenType    string `mapstructure:"MAIL_TOKEN_TYPE"`

	SentryDSN          string `mapstructure:"SENTRY_DSN"`
	MaxMultipartMemory int64  `mapstructure:"MAX_MULTIPART_MEMORY"`
	StorageBucketName  string `mapstructure:"STORAGE_BUCKET_NAME"`

	TimeZone      string `mapstructure:"TIMEZONE"`
	AdminEmail    string `mapstructure:"ADMIN_EMAIL"`
	AdminPassword string `mapstructure:"ADMIN_PASSWORD"`
}

var globalEnv = Env{}

func GetEnv() Env {
	return globalEnv
}

func NewEnv() *Env {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Cannot Read Configuration", err)
	}
	if err := viper.Unmarshal(&globalEnv); err != nil {
		log.Println("Environmen Cannot Be Loaded", err)
	}

	return &globalEnv
}
