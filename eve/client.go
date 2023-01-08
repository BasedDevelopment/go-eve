package eve

// Authenticate uses an email and password to login via Eve and returns a token on success or error on failure
func Authenticate(email string, password string) (string, error) {
	return "", nil
}

// Equivalent to the Authenticate function but panics on errors
func MustAuthenticate(email string, password string) string {
	return ""
}

type Client struct {
	Token string
}
