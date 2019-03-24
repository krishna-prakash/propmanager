package tmp

import (
	"context"

	"github.com/krishna/rogerapp"
	"github.com/krishna/rogerapp/generated/prisma-client"
	"github.com/krishna/rogerapp/server"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Agent() main.AgentResolver {
	return &agentResolver{r}
}
func (r *Resolver) Landlord() main.LandlordResolver {
	return &landlordResolver{r}
}
func (r *Resolver) Mutation() main.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() main.QueryResolver {
	return &queryResolver{r}
}

type agentResolver struct{ *Resolver }

func (r *agentResolver) Clients(ctx context.Context, obj *prisma.Agent) ([]*prisma.Landlord, error) {
	panic("not implemented")
}

type landlordResolver struct{ *Resolver }

func (r *landlordResolver) Agent(ctx context.Context, obj *prisma.Landlord) (*prisma.Agent, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateLandlord(ctx context.Context, userinfo rogerapp.SignupInfo) (*prisma.Landlord, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateAgent(ctx context.Context, userinfo rogerapp.SignupInfo) (*prisma.Agent, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateLandlord(ctx context.Context, id string, landlordInfo rogerapp.LandlordInfo) (*prisma.Landlord, error) {
	panic("not implemented")
}
func (r *mutationResolver) AssignAgentToLandlord(ctx context.Context, agentID string, landlordID string) (*prisma.Landlord, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetLandlords(ctx context.Context) ([]prisma.Landlord, error) {
	panic("not implemented")
}
func (r *queryResolver) GetAgents(ctx context.Context) ([]prisma.Agent, error) {
	panic("not implemented")
}
func (r *queryResolver) GetLandlord(ctx context.Context, id string) (*prisma.Landlord, error) {
	panic("not implemented")
}
