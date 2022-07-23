package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const (
	defaultHTTPHost           = "127.0.0.1"
	defaultHTTPPort           = ":80"
	defaultHTTPRWTimeout      = 10 * time.Second
	defaultHTTPMaxHeaderBytes = 1

	// mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority
	defaultMongoHost     = "cluster.rwtax.mongodb.net"
	defaultMongoUsername = "admin"
	defaultMongoPassword = "admin"
	defaultMongoDBName   = "authDB"

	defaultJWTKey   = "hk24k52uy5gk"
	defaultTokenTTL = 24 * time.Hour

	defaultAuthSalt = "jk3l4nlj34n34l5"
)

type (
	Config struct {
		Mongo MongoCfg
		HTTP  HTTPCfg
		Auth  AuthCfg
	}

	MongoCfg struct {
		Host     string
		Username string
		Password string
		DBName   string
	}

	HTTPCfg struct {
		Host           string
		Port           string
		ReadTimeout    time.Duration
		WriteTimeout   time.Duration
		MaxHeaderBytes int
	}

	JWTCfg struct {
		TokenTTL   time.Duration
		SigningKey string
	}

	AuthCfg struct {
		JWT          JWTCfg
		PasswordSalt string
	}
)

func InitConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(err.Error())
	}

	var cfg Config

	setDefaults(&cfg)
	setFromEnv(&cfg)

	return &cfg
}

func setDefaults(cfg *Config) {
	cfg.HTTP.Host = defaultHTTPHost
	cfg.HTTP.Port = defaultHTTPPort
	cfg.HTTP.ReadTimeout = defaultHTTPRWTimeout
	cfg.HTTP.WriteTimeout = defaultHTTPRWTimeout
	cfg.HTTP.MaxHeaderBytes = defaultHTTPMaxHeaderBytes

	cfg.Mongo.Host = defaultMongoHost
	cfg.Mongo.Username = defaultMongoUsername
	cfg.Mongo.Password = defaultMongoPassword
	cfg.Mongo.DBName = defaultMongoDBName

	cfg.Auth.JWT.SigningKey = defaultJWTKey
	cfg.Auth.JWT.TokenTTL = defaultTokenTTL

	cfg.Auth.PasswordSalt = defaultAuthSalt
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = os.Getenv("HTTP_HOST")
	cfg.HTTP.Port = os.Getenv("HTTP_PORT")

	cfg.Mongo.Host = os.Getenv("MONGO_HOST")
	cfg.Mongo.Username = os.Getenv("MONGO_USER")
	cfg.Mongo.Password = os.Getenv("MONGO_PASSWORD")
	cfg.Mongo.DBName = os.Getenv("MONGO_DBNAME")

	cfg.Auth.JWT.SigningKey = os.Getenv("JWT_KEY")
	cfg.Auth.PasswordSalt = os.Getenv("AUTH_SALT")
}
