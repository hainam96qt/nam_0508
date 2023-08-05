package configs

import (
	"os"
	"time"

	"github.com/go-yaml/yaml"

	"nam_0508/pkg/db/mysql_db"
)

type Config struct {
	Mysqldb mysql_db.DatabaseConfig `yaml:"mysql"`
	Server  Server                  `yaml:"server"`
	Jwt     Jwt                     `yaml:"jwt"`
}

type Jwt struct {
	SecretKey       string        `yaml:"secretKey"`
	AccessTokenTLL  time.Duration `yaml:"accessTokenTLL"`
	RefreshTokenTLL time.Duration `yaml:"refreshTokenTLL"`
}

type Server struct {
	Address         string        `yaml:"address"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	CORS            struct {
		AllowOrigins []string `yaml:"allowOrigins"`
		AllowMethods []string `yaml:"allowMethods"`
		AllowHeaders []string `yaml:"allowHeaders"`
	} `yaml:"cors"`
	SkipLogPaths         []string `yaml:"skipLoggingPaths"`
	ResponseLogLimitByte int      `yaml:"responseLogLimitByte"`
	UploadSizeLimitByte  int64    `yaml:"uploadSizeLimitByte"`
}

// NewConfig returns a new decoded Config struct
func NewConfig() (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open("./config.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
