package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"project_alterra/graph/generated"
	"project_alterra/graph/model"
)

// CreateCast is the resolver for the createCast field.
func (r *mutationResolver) CreateCast(ctx context.Context, input model.CastInput) (*model.Cast, error) {
	var cast model.Cast = r.castService.CreateCast(input)

	return &cast, nil
}

// EditCast is the resolver for the editCast field.
func (r *mutationResolver) EditCast(ctx context.Context, id string, input model.CastInput) (*model.Cast, error) {
	var cast model.Cast = r.castService.EditCast(id, input)

	return &cast, nil
}

// Casts is the resolver for the casts field.
func (r *queryResolver) Casts(ctx context.Context) ([]*model.Cast, error) {
	var casts []*model.Cast = r.castService.GetAllCasts()

	return casts, nil
}

// Cast is the resolver for the cast field.
func (r *queryResolver) Cast(ctx context.Context, id string) (*model.Cast, error) {
	cast, err := r.castService.GetCastById(id)

	if err != nil {
		return nil, err
	}

	return &cast, nil
}

// CastDelete is the resolver for the castDelete field.
func (r *queryResolver) CastDelete(ctx context.Context, id string) (*model.Cast, error) {
	cast, err := r.castService.DeleteCast(id)

	if err != nil {
		return nil, err
	}

	return &cast, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
