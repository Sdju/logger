package logger

import (
	"strings"
	"fmt"
)

func tprintf(format string, params map[string]interface{})string {
	for key, val := range params {
		format = strings.Replace(format, "$"+key+";", fmt.Sprintf("%s", val), -1)
	}
	return fmt.Sprintf(format)
}