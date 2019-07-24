package function

import (
	"fmt"
	"time"
)

// Handle a serverless request
func Handle(req []byte) string {
	t := time.Now()
	return fmt.Sprintf("%s: %s", t.String(), string(req))
}
