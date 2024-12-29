package settings

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type EnvSettings struct {
	SqlUsername    string
	SqlPassword    string
	SqlHost        string
	SqlPort        int
	SqlDatabase    string
	SqlSchema      string
	RedisDb        int
	RedisHost      string
	RedisPort      int
	RedisPassword  string
	S3Bucket       string
	S3Key          string
	S3Region       string
	S3Secret       string
	S3TimePresign  int
	JwtSecret      string
	JwtExpireAcess time.Duration
	TimeZone       string
	Port           int
	SuperName      string
	SuperUsername  string
	SuperPass      string
	SuperEmail     string
	SuperPhone     int
}

var Env EnvSettings

// Load reads and validates environment variables
func Load() {
	Env = EnvSettings{
		// SQL
		SqlUsername: getEnv("SQL_USERNAME", true),
		SqlPassword: getEnv("SQL_PASSWORD", true),
		SqlHost:     getEnv("SQL_HOST", true),
		SqlPort:     getEnvAsInt("SQL_PORT", true),
		SqlDatabase: getEnv("SQL_DATABASE", true),
		SqlSchema:   getEnv("SQL_SCHEMA", true),
		// Redis
		RedisDb:       getEnvAsInt("REDIS_DB", true),
		RedisHost:     getEnv("REDIS_HOST", true),
		RedisPort:     getEnvAsInt("REDIS_PORT", true),
		RedisPassword: getEnv("REDIS_PASSWORD", true),
		// S3 or MINIO
		S3Bucket:      getEnv("S3_BUCKET", true),
		S3Key:         getEnv("S3_KEY", true),
		S3Region:      getEnv("S3_REGION", true),
		S3Secret:      getEnv("S3_SECRET", true),
		S3TimePresign: getEnvAsInt("S3_TIME_PRESIGN", false, 300),
		// JWT
		JwtSecret:      getEnv("JWT_SECRET", true),
		JwtExpireAcess: getEnvAsTime("JWT_EXPIRE_ACCESS", false, 10080),
		// APP
		TimeZone: getEnv("TIMEZONE", false, "America/Fortaleza"),
		Port:     getEnvAsInt("PORT", false, 3000),
		// SUPER USER
		SuperName:     getEnv("SUPER_NAME", true, "Admin"),
		SuperUsername: getEnv("SUPER_USERNAME", true, "admin"),
		SuperPass:     getEnv("SUPER_PASS", true, "admin"),
		SuperEmail:    getEnv("SUPER_EMAIL", true, "ronald.ralds@gmail"),
		SuperPhone:    getEnvAsInt("SUPER_PHONE", true, 558892200365),
	}
}

func getEnv(key string, required bool, defaultValue ...string) string {
	value := os.Getenv(key)

	if value == "" {
		if required {
			panic(fmt.Sprintf("variable %s is required", key))
		}
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return value
}

func getEnvAsInt(key string, required bool, defaultValue ...int) int {
	valueStr := os.Getenv(key)

	if valueStr == "" {
		if required {
			panic(fmt.Sprintf("variable %s is required", key))
		}
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return 0
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		panic(fmt.Sprintf("failed to convert %s to integer: %v", key, err))
	}
	return value
}

func getEnvAsTime(key string, required bool, defaultValue ...int) time.Duration {
	valueStr := os.Getenv(key)

	if valueStr == "" {
		if required {
			panic(fmt.Sprintf("variable %s is required", key))
		}
		if len(defaultValue) > 0 {
			return time.Duration(defaultValue[0]) * time.Minute
		}
		return 0
	}

	// value, err := strconv.ParseInt(valueStr, 10, 64)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		panic(fmt.Sprintf("failed to convert %s to integer: %v", key, err))
	}
	return time.Duration(value) * time.Minute
}
