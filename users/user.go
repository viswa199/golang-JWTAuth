package users

type User struct {
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Autheticate struct{
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}