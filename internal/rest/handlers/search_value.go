package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinsantana/gosolve-recruitment-task/internal/core/modules"
	"github.com/kevinsantana/gosolve-recruitment-task/internal/share"
	log "github.com/sirupsen/logrus"
)

func SearchByValueHandler(ctx *fiber.Ctx) error {
	value, errParse := strconv.ParseInt(ctx.Params("value"), 10, 64)
	if errParse != nil {
		log.WithField("ctx", ctx).
			WithError(errParse).
			Error("error to parse value")

		return Error(ctx, errParse)
	}

	index, num, err := modules.SearchIndexByValue(ctx.Context(), value)
	if index == -1 {
		log.WithField("index", index).
			WithField("value", value).
			Error("Index not found")

		return Error(ctx, share.ClientError{
			Domain:      "search",
			Module:      "get_nums",
			Err:         "index_not_found",
			Description: "The value {value} don't have a match index",
		})
	}
	if index == 0 {
		log.WithField("value", value).
			WithError(err).
			Error("Slice of nums couldn't be loaded")
		return Error(ctx, share.ClientError{
			Domain:      "search",
			Module:      "get_nums",
			Err:         "nums_not_loaded",
			Description: "Slice of nums couldn't be loaded",
		})
	}
	if err != nil {
		log.WithError(err)

		return Error(ctx, err)
	}
	return ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"index": index,
		"value": num,
	})
}
