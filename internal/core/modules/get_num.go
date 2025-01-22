package modules

import (
	"context"
	"math"

	"github.com/kevinsantana/gosolve-recruitment-task/internal/app"
	"github.com/kevinsantana/gosolve-recruitment-task/internal/core/share"
	log "github.com/sirupsen/logrus"
)

func SearchIndexByValue(ctx context.Context, value int64) (int, int64, error) {
	nums, err := app.GetSliceNums(ctx, share.INPUT_FILE)
	if err != nil {
		log.WithField("path", share.INPUT_FILE).
			Error("error to load slice of nums")
		return 0, 0, err
	}

	for i := range nums {
		var conformation = share.CONFORMATION

		percent := int64(float64(nums[i]) * float64(conformation))

		if nums[i] == value {
			log.WithField("index", i).
				Info("found index")
			return i, nums[i], nil
		}

		if int64(math.Abs(float64(value)-float64(nums[i]))) <= percent {
			log.WithField("conformation", conformation).
				WithField("value", value).
				WithField("nums[i]", nums[i]).
				WithField("percent", percent).
				Info("found index by conformation")
			return i, nums[i], nil
		}
	}

	return -1, 0, nil
}
