package docker

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

// Client represents the interface of Docker Engine API client
type Client interface {
	PullImage(image, tag string) (io.ReadCloser, error)
}

// RealClient represents the wrapper of Docker API client
type RealClient struct {
	cli *client.Client
}

// NewClient creates new RealClient object
func NewClient() (*RealClient, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, errors.Wrap(err, "cannot create Docker Engine API client")
	}

	return &RealClient{
		cli: cli,
	}, nil
}

// PullImage pulls image from Docker Image Registry
func (c *RealClient) PullImage(image, tag string) (io.ReadCloser, error) {
	imageRef := fmt.Sprintf("%s:%s", image, tag)

	out, err := c.cli.ImagePull(context.Background(), imageRef, types.ImagePullOptions{})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot pull image %q", imageRef)
	}

	return out, nil
}
