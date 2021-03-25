package Classes

import "regexp"

type Validation struct {

}
func (Validation) IsEmail(email string) bool {
		if len(email) < 3 && len(email) > 254 {
		return false
	}
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return !emailRegex.MatchString(email)
}

func (Validation) IsUsername(username string) bool {
	var usernameRegex = regexp.MustCompile("^[a-zA-Z0-9_]+$")
	return !usernameRegex.MatchString(username)
}