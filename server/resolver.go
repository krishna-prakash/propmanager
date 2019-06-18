package main

import (
	"context"
	"fmt"
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

func (r *landlordResolver) Properties(ctx context.Context, obj *prisma.Landlord) ([]prisma.Property, error) {
	prop, err := r.Prisma.Landlord(prisma.LandlordWhereUniqueInput{
		ID: &obj.ID,
	}).Properties(nil).Exec(ctx)

	return prop, err
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
	fmt.Println(user)
	auth0.CreateUser(&userinfo, user.ID)
	return user, err
}

func (r *mutationResolver) CreateProperty(ctx context.Context, propinfo input.PropertyInfo) (*prisma.Property, error) {
	fmt.Println(propinfo)
	prop, err := r.Prisma.CreateProperty(prisma.PropertyCreateInput{
		Address1: propinfo.Address1,
		Address2: propinfo.Address2,
		Postcode: propinfo.Postcode,
		City: propinfo.City,
		Country: propinfo.Country,
		Percentageofownership: utils.ToInt32(propinfo.Percentageofownership),
		Status: prisma.PropertyStatusCreateOneInput{
			Connect: &prisma.PropertyStatusWhereUniqueInput{
					ID: propinfo.StatusID,
				},
		},
		Type: prisma.PropertyTypeCreateOneInput{
			Connect: &prisma.PropertyTypeWhereUniqueInput{
					ID: propinfo.TypeID,
				},
		},
		Purchaseprice: utils.ToInt32(propinfo.Purchaseprice),
		Currentprice: utils.ToInt32(propinfo.Currentprice),
		MortgageAmount: utils.ToInt32(propinfo.MortgageAmount),
		MortgageInterestRate: utils.ToInt32(propinfo.MortgageInterestRate),
		Currency: propinfo.Currency,
		Landlords: &prisma.LandlordCreateOneWithoutPropertiesInput{
			Connect: &prisma.LandlordWhereUniqueInput{
					ID: &propinfo.LandlordID,
				},
		},
	}).Exec(ctx)

	return prop, err
}

func (r *mutationResolver) UpdateProperty(ctx context.Context, id string, propInfo input.PropertyInfo) (*prisma.Property, error) {
	updatedProp, err := r.Prisma.UpdateProperty(prisma.PropertyUpdateParams{
		Data: prisma.PropertyUpdateInput{
			Address1: &propInfo.Address1,
			Address2: propInfo.Address2,
			Postcode: &propInfo.Postcode,
			City: &propInfo.City,
			Country: &propInfo.Country,
			Percentageofownership: utils.ToInt32(propInfo.Percentageofownership),
			Status: &prisma.PropertyStatusUpdateOneRequiredInput{
				Connect: &prisma.PropertyStatusWhereUniqueInput{
						ID: propInfo.StatusID,
					},
			},
			Type: &prisma.PropertyTypeUpdateOneRequiredInput{
				Connect: &prisma.PropertyTypeWhereUniqueInput{
						ID: propInfo.TypeID,
					},
			},
			Purchaseprice: utils.ToInt32(propInfo.Purchaseprice),
			Currentprice: utils.ToInt32(propInfo.Currentprice),
			Currency: propInfo.Currency,
			},
			Where: prisma.PropertyWhereUniqueInput{
				ID: &id,
			},
	}).Exec(ctx)

	return updatedProp, err
}
func (r *mutationResolver) CreateTenant(ctx context.Context, tenantInfo input.TenantInfo) (*prisma.Tenant, error) {
	tenant, err := r.Prisma.CreateTenant(prisma.TenantCreateInput{
		Title: tenantInfo.Title,
		FirstName: tenantInfo.FirstName,
		MiddleName: tenantInfo.MiddleName,
		LastName: tenantInfo.LastName,
		DisplayName: tenantInfo.DisplayName,
		PersonalEmail: tenantInfo.PersonalEmail,
		WorkEmail: tenantInfo.WorkEmail,
      	TypeOfLet: prisma.TypeOfLetCreateOneInput{
			Connect: &prisma.TypeOfLetWhereUniqueInput{
					ID: tenantInfo.LetID,
				},
		},
      	RentInterval: tenantInfo.RentInterval,
      	Day: tenantInfo.Day,
      	StartDate: tenantInfo.StartDate,
      	EndDate: tenantInfo.EndDate,
      	Notes: tenantInfo.Notes,
		Homenumber: utils.ToInt32(tenantInfo.Homenumber),
		Mobilenumber: utils.ToInt32(tenantInfo.Mobilenumber),
		Property: &prisma.PropertyCreateOneWithoutTenantsInput{
			Connect: &prisma.PropertyWhereUniqueInput{
					ID: &tenantInfo.PropertyID,
				},
		},
	}).Exec(ctx)

	return tenant, err
}
func (r *mutationResolver) UpdateTenant(ctx context.Context, id string, tenantInfo input.TenantInfo) (*prisma.Tenant, error) {
	updatedTenant, err := r.Prisma.UpdateTenant(prisma.TenantUpdateParams{
		Data: prisma.TenantUpdateInput{
			Title: tenantInfo.Title,
			FirstName: &tenantInfo.FirstName,
			MiddleName: tenantInfo.MiddleName,
			LastName: tenantInfo.LastName,
			DisplayName: &tenantInfo.DisplayName,
			PersonalEmail: &tenantInfo.PersonalEmail,
			WorkEmail: &tenantInfo.WorkEmail,
			TypeOfLet: &prisma.TypeOfLetUpdateOneRequiredInput{
				Connect: &prisma.TypeOfLetWhereUniqueInput{
						ID: tenantInfo.LetID,
					},
			},
			RentInterval: &tenantInfo.RentInterval,
			Day: tenantInfo.Day,
			StartDate: tenantInfo.StartDate,
			EndDate: tenantInfo.EndDate,
			Notes: tenantInfo.Notes,
			Homenumber: utils.ToInt32(tenantInfo.Homenumber),
			Mobilenumber: utils.ToInt32(tenantInfo.Mobilenumber),
			Property: &prisma.PropertyUpdateOneWithoutTenantsInput{
				Connect: &prisma.PropertyWhereUniqueInput{
						ID: &tenantInfo.PropertyID,
					},
			},
		},
			Where: prisma.TenantWhereUniqueInput{
				ID: &id,
			},
	}).Exec(ctx)

	return updatedTenant, err
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

	auth0.CreateUser(&userinfo, user.ID)
	return user, err
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetLandlords(ctx context.Context) ([]prisma.Landlord, error) {
	// err := auth0.JwtVerification(ctx)
	// if err != nil {
	// 	return nil, err
	// }
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
	return r.Prisma.Properties(nil).Exec(ctx)
}

func (r *queryResolver) GetPropertTypes(ctx context.Context) ([]prisma.PropertyType, error) {
	return r.Prisma.PropertyTypes(nil).Exec(ctx)
}

func (r *queryResolver) GetPropertStatus(ctx context.Context) ([]prisma.PropertyStatus, error) {
	return r.Prisma.PropertyStatuses(nil).Exec(ctx)
}

func (r *queryResolver) GetProperty(ctx context.Context, id string) (*prisma.Property, error) {
	tes, err := r.Prisma.Property(prisma.PropertyWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	return tes, err
}

func (r *queryResolver) GetTenants(ctx context.Context) ([]prisma.Tenant, error) {
	panic("not implemented")
}

func (r *queryResolver) GetTenant(ctx context.Context, id string) (*prisma.Tenant, error) {
	tes, err := r.Prisma.Tenant(prisma.TenantWhereUniqueInput{
		ID: &id,
	}).Exec(ctx)

	return tes, err
}

type propertyResolver struct{ *Resolver }

func (r *propertyResolver) Status(ctx context.Context, obj *prisma.Property) (*prisma.PropertyStatus, error) {
	status, err := r.Prisma.Property(prisma.PropertyWhereUniqueInput{
		ID: &obj.ID,
	}).Status().Exec(ctx)

	return status, err
}

func (r *propertyResolver) Type(ctx context.Context, obj *prisma.Property) (*prisma.PropertyType, error) {
	ty, err := r.Prisma.Property(prisma.PropertyWhereUniqueInput{
		ID: &obj.ID,
	}).Type().Exec(ctx)

	return ty, err
}

func (r *propertyResolver) Tenants(ctx context.Context, obj *prisma.Property) (*prisma.Tenant, error) {
	tenant, err := r.Prisma.Property(prisma.PropertyWhereUniqueInput{
		ID: &obj.ID,
	}).Tenants().Exec(ctx)

	return tenant, err
}

func (r *propertyResolver) Landlords(ctx context.Context, obj *prisma.Property) (*prisma.Landlord, error) {
	prop, err := r.Prisma.Property(prisma.PropertyWhereUniqueInput{
		ID: &obj.ID,
	}).Landlords().Exec(ctx)

	return prop, err
}

type tenantResolver struct{ *Resolver }

func (r *tenantResolver) TypeOfLet(ctx context.Context, obj *prisma.Tenant) (*prisma.TypeOfLet, error) {
	let, err := r.Prisma.Tenant(prisma.TenantWhereUniqueInput{
		ID: &obj.ID,
	}).TypeOfLet().Exec(ctx)

	return let, err
}

func (r *tenantResolver) Property(ctx context.Context, obj *prisma.Tenant) (*prisma.Property, error) {
	prop, err := r.Prisma.Tenant(prisma.TenantWhereUniqueInput{
		ID: &obj.ID,
	}).Property().Exec(ctx)

	return prop, err
}