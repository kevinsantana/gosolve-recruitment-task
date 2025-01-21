package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinsantana/gosolve-recruitment-task/internal/core/modules"
	log "github.com/sirupsen/logrus"
)

func SearchByValueHandler(ctx *fiber.Ctx) error {
	value, errParse := strconv.ParseInt(ctx.Params("value"), 10, 64)
	if errParse != nil {
		log.WithField("ctx", ctx).
			WithError(errParse)

		return Error(ctx, errParse)
	}

	index, err := modules.GetIndexByValue(ctx.Context(), value)
	if err != nil {
		log.WithError(err)

		return err
	}
	return ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"index": index,
	})
}
