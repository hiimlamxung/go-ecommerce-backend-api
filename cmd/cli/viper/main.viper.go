package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// tag mapstructure là một tag được sử dụng bởi thư viện mapstructure (được Viper sử dụng)
// Dùng để map (ánh xạ) giữa cấu trúc YAML/JSON và struct Go
// Chỉ định tên của trường trong file config sẽ map vào trường struct.
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		DBName   string `mapstructure:"dbName"`
	} `mapstructure:"databases"`
	Security struct {
		JWT struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"jwt"`
	} `mapstructure:"security"`
}

type PaymentConfig struct {
	Methods []struct {
		Name        string `mapstructure:"name"`
		Description string `mapstructure:"description"`
	} `mapstructure:"methods"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigType("yaml")      // type of file

	// Đọc config local
	viper.SetConfigName("local") // ten file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	// Đọc config payment
	viper.SetConfigName("payment") // Tìm file payment.yaml
	if err := viper.MergeInConfig(); err != nil {
		panic(fmt.Errorf("failed to merge payment configuration: %w", err))
	}

	// configure struct cho config chính
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("failed to unmarshal configuration %w", err))
	}

	// configure struct cho payment config
	var paymentConfig PaymentConfig
	if err := viper.Unmarshal(&paymentConfig); err != nil {
		panic(fmt.Errorf("failed to unmarshal payment configuration %w", err))
	}

	// read config from struct
	fmt.Println("Đọc config từ file local.yaml")
	fmt.Println("Server port::", config.Server.Port)
	fmt.Println("JWT key::", config.Security.JWT.Key)

	for _, db := range config.Databases {
		fmt.Printf("Database user: %s, password: %s, host: %s, port: %d \n", db.User, db.Password, db.Host, db.Port)
	}

	// read payment config
	fmt.Println("\nĐọc config từ file payment.yaml")
	for _, method := range paymentConfig.Methods {
		fmt.Printf("- %s: %s\n", method.Name, method.Description)
	}
}
