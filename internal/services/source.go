package services

import "context"

type Sourcer interface {
	//TODO: set return values
	Collect(ctx context.Context) error
}
