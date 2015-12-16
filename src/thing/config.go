package thing

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os/user"
)

// Config represents an app configuration
type Config struct {
	CircleCIApiKey string
	ProjectName    string
	Namespace      string
}

// home directory config file path
func homeConfigDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.thing", usr.HomeDir)
}

// Configure uses viper to get the configuration
func (c *Config) Configure(configPath string) {
	viper.SetConfigName("thing")
	viper.SetConfigType("yaml")

	fmt.Println(configPath)
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	c = &Config{
		ProjectName: viper.GetString("project_name"),
	}

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
