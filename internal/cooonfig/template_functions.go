package cooonfig

import (
	"fmt"
)

func templateFunctionRequired(message string, value any) (any, error) {
	if value == nil || value == "" {
		return nil, fmt.Errorf("%s", message)
	}
	return value, nil
}
