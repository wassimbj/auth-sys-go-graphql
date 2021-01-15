package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"auth-sys/graph/generated"
	"auth-sys/graph/model"
	"auth-sys/services"
	"context"
	"errors"
	"fmt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, data *model.CreateAccData) (bool, error) {
	// panic(fmt.Errorf("not implemented"))
	// fmt.Println(data)
	success := services.CreateAccount(data.Name, data.Email, data.Password)

	if success == true {
		return true, nil
	}

	return false, errors.New("Something went wrong")
}

func (r *mutationResolver) LoginUser(ctx context.Context, data *model.LoginData) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context, id *int) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
