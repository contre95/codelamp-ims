package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"codelamp-ims/pkg/domain/clients"
	"codelamp-ims/pkg/presenters/graphql/graph/generated"
	"codelamp-ims/pkg/presenters/graphql/graph/model"
	"context"
	"fmt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (string, error) {
	panic(fmt.Errorf("not implemented: CreateProject - createProject"))
}

// CreateClient is the resolver for the createClient field.
func (r *mutationResolver) CreateClient(ctx context.Context, input model.NewClient) (string, error) {
	panic(fmt.Errorf("not implemented: CreateClient - createClient"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// Clients is the resolver for the clients field.
func (r *queryResolver) Clients(ctx context.Context) ([]*model.Client, error) {
	req := clients.ListRequest{
		// Maybe no filter, so bring every client at first and let the client filter with graphql
		Filters:  clients.Filter{},
		Page:     0,
		PageSize: 0,
	}
	resp, _ := r.Resolver.clientService.ListUseCase.List(req)
	fmt.Println(resp.Clients)
	var clients []*model.Client
	for _, c := range resp.Clients {
		var projects []*model.Project
		for _, p := range c.Projects {
			projects = append(projects, &model.Project{
				ID:   string(rune(p.ID)),
				Name: p.Name,
				Type: string(p.Type),
			})
		}
		client := model.Client{
			ID:       string(rune(c.ID)),
			Name:     c.Name,
			Country:  c.Country,
			Tag:      c.Tag,
			Projects: projects,
		}
		clients = append(clients, &client)
	}
	return clients, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
