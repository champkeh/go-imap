package tag

import "fmt"

// NewTagGenerator return a new generator for generate
// command's and response's tag
// The tag is from A0001 to A9999, and then restart to
// A0000, and A0001,...
func NewTagGenerator() func() string {
	var c uint16 = 0
	return func() string {
		c = (c + 1) % 10000
		return fmt.Sprintf("A%04d", c)
	}
}
