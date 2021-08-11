package configuration

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type api_config struct {
	Port     int
	Database struct {
		Host  string
		Port  int
		User  string
		Passw string
		Name  string
	}

	Defaults struct {
		wallet string
	}
}

var Config api_config

func setDefaults() {
	viper.SetDefault("Port", 5002)
	viper.SetDefault("Database.Host", "localhost")
	viper.SetDefault("Database.Port", 5002)
	viper.SetDefault("Database.User", "root")
	viper.SetDefault("Database.Passw", "root")
	viper.SetDefault("Database.Name", "Mak")

}

func InitConfig() {
	setDefaults()

	// map viper structure to default ENV_FOO structure to read ENV variables
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Printf("Error Unmarshal: %s", err)
	}
}
