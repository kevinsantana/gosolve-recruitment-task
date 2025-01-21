package modules

import (
	"context"

	"github.com/kevinsantana/gosolve-recruitment-task/internal/app"
	"github.com/kevinsantana/gosolve-recruitment-task/internal/core/share"
	log "github.com/sirupsen/logrus"
)

func GetIndexByValue(ctx context.Context, value int64) ([]int, error) {
	nums, err := app.GetSliceNums(ctx, share.INPUT_FILE)
	if err != nil {
		log.WithError(err).Error("Error to get nums")

		return nil, err
	}

	log.WithField("value", value).Info("received value")

	return nums, nil
}
