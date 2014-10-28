package game

var users = InitUsers()

type User struct {
	Id        int64
	Name      string
	anonymous bool
}

func NewUser(id int64, name string) *User {
	return &User{id, name, false}
}

func NewAnonymousUser(randomName string) *User {
	return &User{0, randomName, true}
}

func InitUsers() map[string]*User {
	us := make(map[string]*User)
	us["simon"] = NewUser(1, "simon")
	us["valor"] = NewUser(2, "valor")
	return us
}

func AuthUser(name string, password string) (*User, bool) {
	user, ok := users[name]
	return user, ok
}
