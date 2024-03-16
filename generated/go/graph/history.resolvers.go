package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/yerTools/ResMon/src/go/api/scalar"
)

func (r *historyQueryResolver) HistoryItems(ctx context.Context, obj *HistoryQuery, from *scalar.Timestamp, to *scalar.Timestamp, limit *int) ([]*HistoryItem, error) {
	panic(fmt.Errorf("not implemented: HistoryItems - historyItems"))
}

func (r *Resolver) HistoryQuery() HistoryQueryResolver { return &historyQueryResolver{r} }

type historyQueryResolver struct{ *Resolver }