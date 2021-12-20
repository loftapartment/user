package user

import "context"

type UserService interface {
	GetUser(ctx context.Context, id int) (*User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
}

type UserRepository interface {
	Create(ctx context.Context, req *User) error
	Update(ctx context.Context, req *User) error
	Get(ctx context.Context, id int) (*User, error)
	GetAll(ctx context.Context) ([]User, error)
}

type User struct {
	ID          int           `json:"id,omitempty"`
	Account     string        `json:"account"`
	Password    string        `json:"password,omitempty"`
	FirstName   string        `json:"firstName,omitempty"`
	LastName    string        `json:"lastName,omitempty"`
	DisplayName string        `json:"displayName,omitempty"`
	Email       string        `json:"email"`
	Third       ThirdProvider `json:"thirdProvider"`
	Gender      string        `json:"gender,omitempty"`
}

type ThirdProvider struct {
	Name  string `json:"name,omitempty"` // ex. google
	Token string `json:"token,omitempty"`
}
