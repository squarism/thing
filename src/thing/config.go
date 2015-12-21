package thing

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os/user"
)

// Config represents an app configuration
type Config struct {
	ProjectName string
}

// home directory config file path
func homeConfigDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/.thing", usr.HomeDir)
}

func setViperConfigDefaults() {
	viper.SetConfigName("thing")
	viper.SetConfigType("yaml")
}

func populateConfig(c *Config) {
	err := viper.ReadInConfig()
	c.ProjectName = viper.GetString("project_name")

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

// Configure uses viper to get the configuration
func (c *Config) Configure(configPaths ...string) {
	setViperConfigDefaults()

	// loop through optionally configured paths
	for _, path := range configPaths {
		viper.AddConfigPath(path)
	}
	viper.AddConfigPath(homeConfigDir())

	populateConfig(c)
}
