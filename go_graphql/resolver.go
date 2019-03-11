package go_graphql

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	return &User{
		ID:   "1",
		Name: "tanaka",
	}, nil
}
func (r *queryResolver) Users(ctx context.Context) ([]*User, error) {
	tanaka := &User{
		ID:   "1",
		Name: "tanaka",
	}
	yamada := &User{
		ID:   "2",
		Name: "yamada",
	}
	var users = []*User{tanaka, yamada}

	return users, nil

}
