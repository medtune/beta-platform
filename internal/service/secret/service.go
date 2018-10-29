package secret

import (
	"fmt"
)

var (
	secrets map[string][]string
)

func init() {
	secrets = make(map[string][]string, 0)
	secrets["signup"] = make([]string, 0, 0)
	secrets["auth"] = make([]string, 0, 0)
}

// Register a new secret key
func Register(service, key string) error {
	if _, ok := secrets[service]; !ok {
		return fmt.Errorf("unknown service")
	}
	secrets[service] = append(secrets[service], key)
	return nil
}

// Check if service exists then if secret exist
// within service
func Check(service, key string) (bool, error) {
	if _, ok := secrets[service]; !ok {
		return false, fmt.Errorf("unknown service")
	}
	for _, s := range secrets[service] {
		if key == s {
			return true, nil
		}
	}
	return false, fmt.Errorf("wrong secret")
}
