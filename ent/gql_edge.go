// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (po *Post) Creator(ctx context.Context) (*User, error) {
	result, err := po.Edges.CreatorOrErr()
	if IsNotLoaded(err) {
		result, err = po.QueryCreator().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Posts(ctx context.Context) (result []*Post, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedPosts(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.PostsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryPosts().All(ctx)
	}
	return result, err
}