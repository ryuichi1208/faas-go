package function

import (
	"fmt"
	"os/exec"
)

// Handle a serverless request
func Handle(req []byte) string {
	// url := "http://google.co.jp"

	// resp, err := http.Get(url)
	// if err != nil {
	// 	return ""
	// }
	// defer resp.Body.Close()

	// _, _ := ioutil.ReadAll(resp.Body)

	// return fmt.Sprintf("Go serverless on Unubo Cloud. %s", string(req))

	out_ls, err := exec.Command("ls -l").Output()
	if err != nil {
		fmt.Println("Command Exec Error.")
	}
	out_ps, err := exec.Command("ps aux").Output()
	if err != nil {
		fmt.Println("Command Exec Error.")
	}

	// 実行したコマンドの結果を出力
	fmt.Sprint("%s\n%s", string(out_ls), string(out_ps))
}
