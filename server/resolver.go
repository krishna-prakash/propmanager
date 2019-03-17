package main

import (
	"context"

	input "github.com/krishna/rogerapp"
	"github.com/krishna/rogerapp/auth0"
	prisma "github.com/krishna/rogerapp/generated/prisma-client"
)

type Resolver struct {
	Prisma *prisma.Client
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, userinfo *input.UserInfo) (*prisma.User, error) {
	user, err := r.Prisma.CreateUser(prisma.UserCreateInput{
		Name:     *userinfo.Name,
		Email:    *userinfo.Email,
		Password: userinfo.Password,
	}).Exec(ctx)

	if err != nil {
		panic(err)
	}

	auth0.CreateUser(userinfo)
	return user, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetUers(ctx context.Context) ([]prisma.User, error) {
	users, err := r.Prisma.Users(nil).Exec(ctx)

	return users, err
}

type userResolver struct{ *Resolver }
