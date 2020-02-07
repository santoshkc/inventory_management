package login

// User used to track logged in user
type User struct {
	Name     string
	password string
}

var userInfo User = User{
	Name:     "santosh",
	password: "hello",
}

func GetUserName() string {
	return userInfo.Name
}
