package another

import (
	"fmt"
	"os"
)

// export:
func Cat(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
