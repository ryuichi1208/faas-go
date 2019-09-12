package function

import (
	"fmt"
	"os/exec"
)

func ExecCmd() string {
	var cmd [12]string
	cmd[0] = "ps"
	cmd[1] = "ls"
	cmd[2] = "df"

	ret := ""
	for i := 0; i < len(cmd); i++ {
		out, err := exec.Command(cmd[i]).Output()
		if err != nil {
			fmt.Println("Command Exec Error.")
		}
		ret = fmt.Sprintf("%s\n%s", ret, string(out))
	}

	return ret
}

// Handle a serverless request
func Handle(req []byte) string {
	return ExecCmd()
}
