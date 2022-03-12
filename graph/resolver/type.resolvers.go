package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/MatsuoTakuro/gqlgen-todos/graph/generated"
	"github.com/MatsuoTakuro/gqlgen-todos/graph/model"
)

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	user := &model.User{
		ID:   obj.UserID,
		Name: "user " + obj.UserID,
	}
	return user, nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
