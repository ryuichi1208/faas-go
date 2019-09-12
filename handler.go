package function

import (
	"io/ioutil"
	"net/http"
)

// Handle a serverless request
func Handle(req []byte) string {
	url := "http://google.co.jp"

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray)

	// return fmt.Sprintf("Go serverless on Unubo Cloud. %s", string(req))
}
