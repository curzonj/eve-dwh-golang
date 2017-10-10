package sqlh

import (
	"fmt"
	"strings"
)

func BuildColumnsValues(m map[string]interface{}) string {
	keys := make([]string, 0, len(m))
	values := make([]string, 0, len(m))

	for k, _ := range m {
		keys = append(keys, k)
		values = append(values, ":"+k)
	}

	return fmt.Sprintf("(%s) VALUES (%s)", strings.Join(keys, ","), strings.Join(values, ","))
}
