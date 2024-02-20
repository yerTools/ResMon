package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"
)

func (r *activityMutationResolver) SetActive(ctx context.Context, obj *ActivityMutation, active bool) (bool, error) {
	panic(fmt.Errorf("not implemented: SetActive - setActive"))
}

func (r *Resolver) ActivityMutation() ActivityMutationResolver { return &activityMutationResolver{r} }

type activityMutationResolver struct{ *Resolver }
