package api

type Login struct {
	Email	string		`json:"email"`
	Pass	string		`json:"password"`
}

type UserImage struct {
	Id		string		`json:"id"`
	Email		string		`json:"email"`
	Name		string		`json:"name"`
	Role		string		`json:"role"`
	Created		string		`json:"created"`
}
