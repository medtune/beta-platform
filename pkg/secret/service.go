package secret

import (
	"fmt"
)

var (
	secrets map[string][]string
)

func init() {
	secrets = make(map[string][]string, 0)

	secrets["signup"] = []string{"supersecret"}
	secrets["auth"] = make([]string, 0, 0)
}

func Register(service, key string) error {
	if _, ok := secrets[service]; !ok {
		return fmt.Errorf("Unknown service")
	}
	secrets[service] = append(secrets[service], key)
	return nil
}

func Check(service, key string) (bool, error) {
	if _, ok := secrets[service]; !ok {
		return false, fmt.Errorf("Unknown service")
	}
	for _, s := range secrets[service] {
		if key == s {
			return true, nil
		}
	}
	return false, fmt.Errorf("Wrong secret.")
}
