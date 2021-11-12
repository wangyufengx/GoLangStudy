package docker

import (
	"fmt"
	"testing"
)

func TestImageOperation(t *testing.T) {
	err := ImageOperation()
	fmt.Println(err)
}
