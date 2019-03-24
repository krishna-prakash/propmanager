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

func (r *Resolver) Agent() AgentResolver {
	return &agentResolver{r}
}
func (r *Resolver) Landlord() LandlordResolver {
	return &landlordResolver{r}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type agentResolver struct{ *Resolver }

func (r *agentResolver) Clients(ctx context.Context, obj *prisma.Agent) ([]*prisma.Landlord, error) {
	// return r.Prisma.Agent(prisma.AgentWhereUniqueInput{
	// 	ID: &obj.ID,
	// }).Clients(nil).Exec(ctx)
	panic("not implemented")
}

type landlordResolver struct{ *Resolver }

func (r *landlordResolver) Agent(ctx context.Context, obj *prisma.Landlord) (*prisma.Agent, error) {
	agent, err := r.Prisma.Landlord(prisma.LandlordWhereUniqueInput{
		ID: &obj.ID,
	}).Agent().Exec(ctx)

	return agent, err
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateLandlord(ctx context.Context, userinfo input.SignupInfo) (*prisma.Landlord, error) {
	user, err := r.Prisma.CreateLandlord(prisma.LandlordCreateInput{
		FullName: *userinfo.Name,
		Email:    *userinfo.Email,
		Password: userinfo.Password,
	}).Exec(ctx)

	if err != nil {
		panic(err)
	}

	auth0.CreateUser(&userinfo)
	return user, err
}

func (r *mutationResolver) UpdateLandlord(ctx context.Context, id string, landlordInfo input.LandlordInfo) (*prisma.Landlord, error) {
	landlord, err := r.Prisma.UpdateLandlord(prisma.LandlordUpdateParams{
		Data: prisma.LandlordUpdateInput{
			FullName: landlordInfo.FullName,
			Email:    landlordInfo.Email,
			Address1: landlordInfo.Address1,
			Address2: landlordInfo.Address2,
			Country:  landlordInfo.Country,
			Postcode: landlordInfo.Postcode,
		},
		Where: prisma.LandlordWhereUniqueInput{
			ID: &id,
		},
	}).Exec(ctx)

	return landlord, err
}

func (r *mutationResolver) AssignAgentToLandlord(ctx context.Context, agentID string, landlordID string) (*prisma.Landlord, error) {
	landlord, err := r.Prisma.UpdateLandlord(prisma.LandlordUpdateParams{
		Data: prisma.LandlordUpdateInput{
			Agent: &prisma.AgentUpdateOneWithoutClientsInput{
				Connect: &prisma.AgentWhereUniqueInput{
					ID: &agentID,
				},
			},
		},
		Where: prisma.LandlordWhereUniqueInput{
			ID: &landlordID,
		},
	}).Exec(ctx)

	return landlord, err
}

func (r *mutationResolver) CreateAgent(ctx context.Context, userinfo input.SignupInfo) (*prisma.Agent, error) {
	user, err := r.Prisma.CreateAgent(prisma.AgentCreateInput{
		FirstName: *userinfo.Name,
		Email:     *userinfo.Email,
		Password:  userinfo.Password,
	}).Exec(ctx)

	if err != nil {
		panic(err)
	}

	auth0.CreateUser(&userinfo)
	return user, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetLandlords(ctx context.Context) ([]prisma.Landlord, error) {
	return r.Prisma.Landlords(nil).Exec(ctx)
}

func (r *queryResolver) GetLandlord(ctx context.Context, id string) (*prisma.Landlord, error) {
	tes, err := r.Prisma.Landlord(prisma.LandlordWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	return tes, err
}

func (r *queryResolver) GetAgents(ctx context.Context) ([]prisma.Agent, error) {
	agents, err := r.Prisma.Agents(nil).Exec(ctx)

	return agents, err
}
