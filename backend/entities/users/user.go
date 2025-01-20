package users

type User struct {
	id           int64  `json: "id"`
	name         string `json: "name"`
	surname      string `json: "surname"`
	email        string `json: "email"`
	phone_number string `json: "phone_number"`
	role         Role   `json: "role"`
	password     string `json: "password"`
}
