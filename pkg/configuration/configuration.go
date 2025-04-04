package configuration

import "github.com/Netflix/go-env"

type Server struct {
	HttpSocket string `env:"SERVER_HTTP_SOCKET"`
	CdnBaseUrl string `env:"SERVER_CDN_BASE_URL"`
}

type DB struct {
	Host string `env:"DB_HOST"`
	Name string `env:"DB_NAME"`
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASS"`
}

type S3 struct {
	Endpoint        string `env:"S3_ENDPOINT"`
	AccessKeyID     string `env:"S3_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"S3_SECRET_ACCESS_KEY"`
	BucketName      string `env:"S3_BUCKET_NAME"`
	UseSSL          bool   `env:"S3_USE_SSL"`
}

type Logger struct {
	OutputPaths      string `env:"LOGGER_OUTPUT_PATHS"`
	ErrorOutputPaths string `env:"LOGGER_ERROR_OUTPUT_PATHS"`
}

type Config struct {
	Server Server
	DB     DB
	Logger Logger
	S3     S3
}

func Load() (*Config, error) {
	var cfg Config
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
