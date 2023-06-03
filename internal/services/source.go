package services

import (
	"context"
	"weather-collector/internal/models"
)

type Sourcer interface {
	Collect(ctx context.Context) (*models.Forecast, error)
}
