package main

import (
	"fmt"
	"github.com/spf13/viper"
	"thing"
)

func main() {
	viper.SetConfigName("thing")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./fixtures")

	err := viper.ReadInConfig()
	c := &thing.Config{
		ProjectName: viper.GetString("project_name"),
	}

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("Project Name is")
	fmt.Println(c.ProjectName)
}
