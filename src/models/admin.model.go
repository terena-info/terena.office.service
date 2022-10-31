package models

var (
	AdminModelName string = "admins"
)

type Admin struct {
	DefaultField `bson:",inline"`
	Email        string `validate:"required,email" json:"email" form:"email" bson:"email"`
	Password     string `validate:"required" json:"-" form:"password" bson:"password"`
	ProfileIcon  string `validate:"required" json:"profile_icon" form:"profile_icon" bson:"profile_icon"`
}

type Login struct {
	Email    string `validate:"required,email" json:"email" form:"email" bson:"email"`
	Password string `validate:"required" json:"password" form:"password" bson:"password"`
}
