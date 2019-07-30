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
func (r *Resolver) Category() main.CategoryResolver {
	return &categoryResolver{r}
}
func (r *Resolver) Landlord() main.LandlordResolver {
	return &landlordResolver{r}
}
func (r *Resolver) Mutation() main.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Property() main.PropertyResolver {
	return &propertyResolver{r}
}
func (r *Resolver) Query() main.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Tenant() main.TenantResolver {
	return &tenantResolver{r}
}
func (r *Resolver) Transaction() main.TransactionResolver {
	return &transactionResolver{r}
}
func (r *Resolver) Type() main.TypeResolver {
	return &typeResolver{r}
}
func (r *Resolver) UserCategory() main.UserCategoryResolver {
	return &userCategoryResolver{r}
}

type agentResolver struct{ *Resolver }

func (r *agentResolver) Clients(ctx context.Context, obj *prisma.Agent) ([]prisma.Landlord, error) {
	panic("not implemented")
}

type categoryResolver struct{ *Resolver }

func (r *categoryResolver) Type(ctx context.Context, obj *prisma.Category) (*prisma.Type, error) {
	panic("not implemented")
}

type landlordResolver struct{ *Resolver }

func (r *landlordResolver) Agent(ctx context.Context, obj *prisma.Landlord) (*prisma.Agent, error) {
	panic("not implemented")
}
func (r *landlordResolver) Properties(ctx context.Context, obj *prisma.Landlord) ([]prisma.Property, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateLandlord(ctx context.Context, userinfo rogerapp.SignupInfo) (*prisma.Landlord, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateLandlord(ctx context.Context, id string, landlordInfo rogerapp.LandlordInfo) (*prisma.Landlord, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateProperty(ctx context.Context, propinfo rogerapp.PropertyInfo) (*prisma.Property, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateProperty(ctx context.Context, id string, propInfo rogerapp.PropertyInfo) (*prisma.Property, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateTenant(ctx context.Context, tenantInfo rogerapp.TenantInfo) (*prisma.Tenant, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateTenant(ctx context.Context, id string, tenantInfo rogerapp.TenantInfo) (*prisma.Tenant, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateAgent(ctx context.Context, userinfo rogerapp.SignupInfo) (*prisma.Agent, error) {
	panic("not implemented")
}
func (r *mutationResolver) AssignAgentToLandlord(ctx context.Context, agentID string, landlordID string) (*prisma.Landlord, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateCategory(ctx context.Context, categoryInfo rogerapp.CategoryInfo) (*prisma.Category, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateUserCategory(ctx context.Context, categoryInfo rogerapp.UserCategoryInfo) (*prisma.UserCategory, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateTransaction(ctx context.Context, transactionInfo rogerapp.TransactionInfo) (*prisma.Transaction, error) {
	panic("not implemented")
}

type propertyResolver struct{ *Resolver }

func (r *propertyResolver) Status(ctx context.Context, obj *prisma.Property) (*prisma.PropertyStatus, error) {
	panic("not implemented")
}
func (r *propertyResolver) Type(ctx context.Context, obj *prisma.Property) (*prisma.PropertyType, error) {
	panic("not implemented")
}
func (r *propertyResolver) Tenants(ctx context.Context, obj *prisma.Property) (*prisma.Tenant, error) {
	panic("not implemented")
}
func (r *propertyResolver) Landlords(ctx context.Context, obj *prisma.Property) (*prisma.Landlord, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetLandlords(ctx context.Context) ([]prisma.Landlord, error) {
	panic("not implemented")
}
func (r *queryResolver) GetProperties(ctx context.Context) ([]prisma.Property, error) {
	panic("not implemented")
}
func (r *queryResolver) GetAgents(ctx context.Context) ([]prisma.Agent, error) {
	panic("not implemented")
}
func (r *queryResolver) GetTenants(ctx context.Context) ([]prisma.Tenant, error) {
	panic("not implemented")
}
func (r *queryResolver) GetTypes(ctx context.Context) ([]prisma.Type, error) {
	panic("not implemented")
}
func (r *queryResolver) GetCategories(ctx context.Context) ([]prisma.Category, error) {
	panic("not implemented")
}
func (r *queryResolver) GetUserCategories(ctx context.Context) ([]prisma.UserCategory, error) {
	panic("not implemented")
}
func (r *queryResolver) GetPropertTypes(ctx context.Context) ([]prisma.PropertyType, error) {
	panic("not implemented")
}
func (r *queryResolver) GetPropertStatus(ctx context.Context) ([]prisma.PropertyStatus, error) {
	panic("not implemented")
}
func (r *queryResolver) GetLandlord(ctx context.Context, id string) (*prisma.Landlord, error) {
	panic("not implemented")
}
func (r *queryResolver) GetProperty(ctx context.Context, id string) (*prisma.Property, error) {
	panic("not implemented")
}
func (r *queryResolver) GetTenant(ctx context.Context, id string) (*prisma.Tenant, error) {
	panic("not implemented")
}
func (r *queryResolver) GetTransactionByType(ctx context.Context, typeID string) ([]prisma.Transaction, error) {
	panic("not implemented")
}
func (r *queryResolver) GetTransactionByCategory(ctx context.Context, categoryID string) ([]prisma.Transaction, error) {
	panic("not implemented")
}

type tenantResolver struct{ *Resolver }

func (r *tenantResolver) TypeOfLet(ctx context.Context, obj *prisma.Tenant) (*prisma.TypeOfLet, error) {
	panic("not implemented")
}
func (r *tenantResolver) Property(ctx context.Context, obj *prisma.Tenant) (*prisma.Property, error) {
	panic("not implemented")
}

type transactionResolver struct{ *Resolver }

func (r *transactionResolver) Type(ctx context.Context, obj *prisma.Transaction) (*prisma.Type, error) {
	panic("not implemented")
}
func (r *transactionResolver) Category(ctx context.Context, obj *prisma.Transaction) (*prisma.Category, error) {
	panic("not implemented")
}
func (r *transactionResolver) Property(ctx context.Context, obj *prisma.Transaction) (*prisma.Property, error) {
	panic("not implemented")
}
func (r *transactionResolver) Supplier(ctx context.Context, obj *prisma.Transaction) (*prisma.Supplier, error) {
	panic("not implemented")
}

type typeResolver struct{ *Resolver }

func (r *typeResolver) Categories(ctx context.Context, obj *prisma.Type) ([]prisma.Category, error) {
	panic("not implemented")
}

type userCategoryResolver struct{ *Resolver }

func (r *userCategoryResolver) Category(ctx context.Context, obj *prisma.UserCategory) (*prisma.Category, error) {
	panic("not implemented")
}
func (r *userCategoryResolver) Landlord(ctx context.Context, obj *prisma.UserCategory) (*prisma.Landlord, error) {
	panic("not implemented")
}
func (r *userCategoryResolver) Agent(ctx context.Context, obj *prisma.UserCategory) (*prisma.Agent, error) {
	panic("not implemented")
}
