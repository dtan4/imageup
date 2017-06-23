package server

import (
	"net/http"

	"github.com/labstack/echo"
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

func webhooksQuayHandler(c echo.Context) error {
	r := new(QuayRequest)

	if err := c.Bind(r); err != nil {
		// TODO: print error in log
		return echo.NewHTTPError(http.StatusBadRequest, "invalid json payload")
	}

	return c.String(http.StatusAccepted, "accepted")
}
