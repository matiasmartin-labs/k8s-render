package utils

import (
	"fmt"
	"strings"
)

type VarsFlag map[string]string

func (v *VarsFlag) String() string {
	return "vars"
}

func (v *VarsFlag) Set(value string) error {
	parts := strings.SplitN(value, "=", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid format: %s, expected key=value", value)
	}

	key := strings.TrimSpace(parts[0])
	val := strings.TrimSpace(parts[1])
	if key == "" || val == "" {
		return fmt.Errorf("key and value must not be empty")
	}

	(*v)[key] = val

	return nil
}
