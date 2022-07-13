package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	projectName := "go-practise"
	getConfig(projectName)
}

func getConfig(projectName string) {
	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath(".")                                                                          // optionally look for config in the working directory
	viper.AddConfigPath(fmt.Sprintf("D:/go/%s/chapt05/", projectName))                                // call multiple times to add many search paths
	viper.AddConfigPath(fmt.Sprintf("/$HOME/GoProjects/go-practise/config/%s/chapt05/", projectName)) // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

// GetMysqlConnectingString func
func GetMysqlConnectingString() string {
	usr := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	port := viper.GetString("mysql.port")
	host := viper.GetString("mysql.host")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", usr, pwd, host, port, db, charset)
}
