package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) string {
	return fmt.Sprintf("Go serverless on Unubo Cloud. %s", string(req))
}
