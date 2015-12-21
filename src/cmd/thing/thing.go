package main

import (
	"fmt"
	"thing"
)

func main() {
	c := &thing.Config{}
	c.Configure("./fixtures")

	fmt.Println("Project Name is")
	fmt.Println(c.ProjectName)
}
