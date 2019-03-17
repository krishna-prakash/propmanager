package tmp

import (
	"context"

	"github.com/krishna/rogerapp"
	"github.com/krishna/rogerapp/generated/prisma-client"
	"github.com/krishna/rogerapp/server"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() main.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() main.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, userinfo *rogerapp.UserInfo) (*prisma.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetUers(ctx context.Context) ([]prisma.User, error) {
	panic("not implemented")
}
