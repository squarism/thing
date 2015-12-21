package thing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fixturesDirectory := fmt.Sprintf("%s/../../fixtures", pwd)

	c := new(Config)
	c.Configure(fixturesDirectory)

	assert.Equal(t, "thing_project", c.ProjectName)
}
