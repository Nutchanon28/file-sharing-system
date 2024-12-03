package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.57

import (
	"context"
	"fmt"

	"github.com/Nutchanon28/file-sharing-system/graph/model"
)

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Departments is the resolver for the departments field.
func (r *queryResolver) Departments(ctx context.Context) ([]*model.Department, error) {
	panic(fmt.Errorf("not implemented: Departments - departments"))
}

// Faculties is the resolver for the faculties field.
func (r *queryResolver) Faculties(ctx context.Context) ([]*model.Faculty, error) {
	panic(fmt.Errorf("not implemented: Faculties - faculties"))
}

// SharedFiles is the resolver for the sharedFiles field.
func (r *queryResolver) SharedFiles(ctx context.Context) ([]*model.SharedFile, error) {
	panic(fmt.Errorf("not implemented: SharedFiles - sharedFiles"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// GetDepartment is the resolver for the getDepartment field.
func (r *queryResolver) GetDepartment(ctx context.Context, id string) (*model.Department, error) {
	panic(fmt.Errorf("not implemented: GetDepartment - getDepartment"))
}

// GetFaculty is the resolver for the getFaculty field.
func (r *queryResolver) GetFaculty(ctx context.Context, id string) (*model.Faculty, error) {
	panic(fmt.Errorf("not implemented: GetFaculty - getFaculty"))
}

// GetSharedFile is the resolver for the getSharedFile field.
func (r *queryResolver) GetSharedFile(ctx context.Context, id string) (*model.SharedFile, error) {
	panic(fmt.Errorf("not implemented: GetSharedFile - getSharedFile"))
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }
type mutationResolver struct{ *Resolver }
*/