package tag

import (
	"fmt"
	"testing"
)

func TestNewTagGenerator(t *testing.T) {
	generator := NewTagGenerator()

	for i := 0; i < 10200; i++ {
		tag := generator()
		fmt.Println(tag)
	}
}
