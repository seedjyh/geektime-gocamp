package uuid

import (
	"fmt"
	"testing"
)

func TestNewGenerator(t *testing.T) {
	c := NewGenerator().Channel()
	fmt.Println(<-c)
	fmt.Println(<-c)
}
