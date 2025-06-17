package initialize

import (
	"fmt"

	"github.com/hiimlamxung/go-ecommerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/") // path to config
	viper.SetConfigType("yaml")      // type of file

	// Đọc config local
	viper.SetConfigName("local") // ten file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	// configure struct cho config chính
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("failed to unmarshal configuration %w", err))
	}
}
