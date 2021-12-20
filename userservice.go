package user

import "context"

type ManageUserService struct {
	repo UserRepository
}

func NewUserService(ur UserRepository) *ManageUserService {
	return &ManageUserService{
		repo: ur,
	}
}

func (us *ManageUserService) GetUser(ctx context.Context, id int) (*User, error) {
	return us.repo.Get(ctx, id)
}

func (us *ManageUserService) GetAllUsers(ctx context.Context) ([]User, error) {
	return us.repo.GetAll(ctx)
}

func (us *ManageUserService) CreateUser(ctx context.Context, user *User) (*User, error) {
	err := us.repo.Create(ctx, user)
	if err != nil {
		return nil, err

	}
	return user, nil
}

func (us *ManageUserService) UpdateUser(ctx context.Context, user *User) (*User, error) {
	err := us.repo.Update(ctx, user)
	if err != nil {
		return nil, err

	}
	return user, nil
}
