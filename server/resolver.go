package main

import (
	"context"

	input "github.com/krishna/rogerapp"
	"github.com/krishna/rogerapp/utils"
	"github.com/krishna/rogerapp/auth0"
	prisma "github.com/krishna/rogerapp/generated/prisma-client"
)

type tok struct {
	Authorization string
}
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

func (r *Resolver) Property() PropertyResolver {
	return &propertyResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Tenant() TenantResolver {
	return &tenantResolver{r}
}
type agentResolver struct{ *Resolver }

func (r *agentResolver) Clients(ctx context.Context, obj *prisma.Agent) ([]prisma.Landlord, error) {
	tes, err := r.Prisma.Agent(prisma.AgentWhereUniqueInput{
		ID: &obj.ID,
	}).Clients(nil).Exec(ctx)

	return tes, err
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
		FullName: userinfo.Name,
		Email:    userinfo.Email,
		Mobile:   utils.ToInt32(&userinfo.Mobile),
		Password: &userinfo.Password,
	}).Exec(ctx)

	if err != nil {
		panic(err)
	}

	auth0.CreateUser(&userinfo)
	return user, err
}

type propertyResolver struct{ *Resolver }

func (r *propertyResolver) Tenants(ctx context.Context, obj *prisma.Property) ([]*prisma.Tenant, error) {
	panic("not implemented")
}
func (r *propertyResolver) Landlords(ctx context.Context, obj *prisma.Property) ([]*prisma.Landlord, error) {
	panic("not implemented")
}

type tenantResolver struct{ *Resolver }

func (r *tenantResolver) Property(ctx context.Context, obj *prisma.Tenant) (*prisma.Property, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateProperty(ctx context.Context, propinfo input.PropertyInfo) (*prisma.Property, error) {
	prop, err := r.Prisma.CreateProperty(prisma.PropertyCreateInput{
		Displayname : propinfo.Displayname,
		Address1: propinfo.Address1,
		Address2: propinfo.Address2,
		Postcode: propinfo.Postcode,
		City: propinfo.City,
		Country: propinfo.Country,
		Percentageofownership: utils.ToInt32(propinfo.Percentageofownership),
		Status: propinfo.Status,
		Purchaseprice: utils.ToInt32(propinfo.Purchaseprice),
		Currentprice: utils.ToInt32(propinfo.Currentprice),
		Currency: propinfo.Currency,
	}).Exec(ctx)

	return prop, err
}

func (r *mutationResolver) UpdateLandlord(ctx context.Context, id string, landlordInfo input.LandlordInfo) (*prisma.Landlord, error) {
	// err := auth0.JwtVerification(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	landlord, err := r.Prisma.UpdateLandlord(prisma.LandlordUpdateParams{
		Data: prisma.LandlordUpdateInput{
			FullName: landlordInfo.FullName,
			Email:    landlordInfo.Email,
			Mobile:   utils.ToInt32(landlordInfo.Mobile),
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
	err := auth0.JwtVerification(ctx)
	if err != nil {
		return nil, err
	}
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
		FirstName: userinfo.Name,
		Email:     userinfo.Email,
		Password:  &userinfo.Password,
	}).Exec(ctx)

	if err != nil {
		panic(err)
	}

	auth0.CreateUser(&userinfo)
	return user, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetLandlords(ctx context.Context) ([]prisma.Landlord, error) {
	err := auth0.JwtVerification(ctx)
	if err != nil {
		return nil, err
	}
	return r.Prisma.Landlords(nil).Exec(ctx)
}

func (r *queryResolver) GetLandlord(ctx context.Context, id string) (*prisma.Landlord, error) {
	// err := auth0.JwtVerification(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	tes, err := r.Prisma.Landlord(prisma.LandlordWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	return tes, err
}

func (r *queryResolver) GetAgents(ctx context.Context) ([]prisma.Agent, error) {
	err := auth0.JwtVerification(ctx)
	if err != nil {
		return nil, err
	}
	agents, err := r.Prisma.Agents(nil).Exec(ctx)

	return agents, err
}

func (r *queryResolver) GetProperties(ctx context.Context) ([]prisma.Property, error) {
	panic("not implemented")
}
