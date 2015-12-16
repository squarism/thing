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

	c := new(Config)
	c.Configure(fmt.Sprintf("%s/../../fixtures", pwd))

	assert.Equal(t, "testtest", c.ProjectName)
}
