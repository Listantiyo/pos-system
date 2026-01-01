package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName   				string	`mapstructure:"APP_NAME"`
	AppPort   				string	`mapstructure:"APP_PORT"`

	DBHost    				string	`mapstructure:"DB_HOST"`
	DBPort    				string	`mapstructure:"DB_PORT"`
	DBUser    				string	`mapstructure:"DB_USER"`
	DBPass    				string	`mapstructure:"DB_PASSWORD"`
	DBName    				string	`mapstructure:"DB_NAME"`
	DBSSLMode 				string	`mapstructure:"DB_SSLMODE"`

	JWTSecret 				string	`mapstructure:"JWT_SECRET"`
	JWTExpiryHours 			string	`mapstructure:"JWT_EXPIRY_HOURS"`
	RefreshTokenExpiryDays 	string	`mapstructure:"REFRESH_TOKEN_EXPIRY_DAYS"`

	RedisHost				string	`mapstructure:"REDIS_HOST"`
	RedisPort				string	`mapstructure:"REDIS_PORT"`
	RedisPass				string	`mapstructure:"REDIS_PASSWORD"`
	RedisDB					int		`mapstructure:"REDIS_DB"`

	TaxRate					float64	`mapstructure:"TAX_RATE"`
	DefaultCurrency			string	`mapstructure:"DEFAULT_CURRENCY"`
	Timezone				string	`mapstructure:"TIMEZONE"`

	EnableRateLimiting		bool	`mapstructure:"ENABLE_RATE_LIMITING"`
	EnableCaching			bool	`mapstructure:"ENABLE_CACHING"`
}

func LoadConfig() *Config {
	var config Config

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetDefault("APP_NAME", "App Name")
	viper.SetDefault("APP_PORT", "8080")

	// Default Postgres config
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASS", "123123")
	viper.SetDefault("DB_NAME", "book_api")
	viper.SetDefault("DB_SSLMODE", "disable")

	// Default JWT config
	viper.SetDefault("JWT_SECRET", "secret")
	viper.SetDefault("JWT_EXPIRY_HOURS", "24")
	viper.SetDefault("REFRESH_TOKEN_EXPIRY_DAYS", "7")

	// Default Redis config
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", "6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)

	// Default POS setting
	viper.SetDefault("TAX_RATE", 0.10)
	viper.SetDefault("DEFAULT_CURRENCY", "IDR")
	viper.SetDefault("TIMEZONE", "Asia/Jakarta")

	// Default Featured Flags
	viper.SetDefault("ENABLE_RATE_LIMITING", true)
	viper.SetDefault("ENABLE_CACHING", true)

	if err := viper.ReadInConfig(); err != nil {
		 log.Println("No .env file found, using environment variables or defaults")
	}else{
		log.Println("✅ Configuration loaded successfully.")
	}

	// Unmarshal
	err := viper.Unmarshal(&config)
	if err != nil{
		log.Fatalf("❌ Unable to decode into struct: %v", err)
	}

	log.Println("✅ Configuration successfully unmarshaled.")
    return &config
}