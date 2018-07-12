package access

var (
	signupSecretKeys []string
)

func IsValideSecret(secret string) bool {
	for _, s := range signupSecretKeys {
		if secret == s {
			return true
		}
	}
	return false
}
