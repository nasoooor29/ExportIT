package another

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// export: get response from website
func Curl(method string, link string, payloadPath string) error {
	f, err := os.Open(payloadPath)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(method, link, f)
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
