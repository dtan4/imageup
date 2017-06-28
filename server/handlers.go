package server

import (
	"fmt"
	"net/http"

	"github.com/dtan4/imageup/docker"
	"github.com/dtan4/imageup/server/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

func rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "ImageUp")
}

func pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

type QuayRequest struct {
	BuildID         string                      `json:"build_id"`
	TriggerKind     string                      `json:"trigger_kind"`
	Name            string                      `json:"name"`
	Repository      string                      `json:"repository"`
	Namespace       string                      `json:"namespace"`
	DockerURL       string                      `json:"docker_url"`
	TriggerID       string                      `json:"trigger_id"`
	DockerTags      []string                    `json:"docker_tags"`
	BuildName       string                      `json:"build_name"`
	ImageID         string                      `json:"image_id"`
	TriggerMetadata *QuayRequestTriggerMetadata `json:"trigger_metadata"`
	Homepage        string                      `json:"homepage"`
}

type QuayRequestTriggerMetadata struct {
	DefaultBranch string `json:"default_branch"`
	Ref           string `json:"ref"`
	Commit        string `json:"commit"`
}

func contains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}

	return false
}

func webhooksQuayHandler(c echo.Context) error {
	r := new(QuayRequest)

	if err := c.Bind(r); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errors.Wrap(err, "invalid json payload").Error())
	}

	cli, err := middleware.GetDockerClient(c)
	if err != nil {
		c.Logger().Error(errors.Wrap(err, "Docker client is not set"))
		return echo.NewHTTPError(http.StatusInternalServerError, "Docker client is not set")
	}

	imageWhitelist, err := middleware.GetImageWhitelist(c)
	if err != nil {
		c.Logger().Error(errors.Wrap(err, "invalid image whitelist"))
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid image whitelist")
	}

	if imageWhitelist != nil {
		if !contains(imageWhitelist, r.DockerURL) {
			return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("pulling %s is not authorized", r.DockerURL))
		}
	}

	for _, tag := range r.DockerTags {
		go func(t string) {
			out, err := cli.PullImage(r.DockerURL, t)
			if err != nil {
				c.Logger().Error(err)
				return
			}
			defer out.Close()

			pullID := uuid.NewV4().String()

			docker.PrintPullMessage(out, func(line string) {
				c.Logger().Printj(log.JSON{
					"message": line,
					"pull-id": pullID,
				})
			})

			c.Logger().Printj(log.JSON{
				"message": fmt.Sprintf("pulling %s:%s completed successfully", r.DockerURL, t),
				"pull-id": pullID,
			})

		}(tag)
	}

	return c.String(http.StatusAccepted, "accepted")
}
