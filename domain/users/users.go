package users

type UserKey struct {
	Email string
}

func NewUserKey(email string) UserKey {
	return UserKey{Email: email}
}

func (k *UserKey) String() string {
	return k.Email
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
