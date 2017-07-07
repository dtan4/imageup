package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/cli/config"
	"github.com/docker/docker/cli/config/configfile"
	"github.com/docker/docker/cli/config/credentials"
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
	cfg *configfile.ConfigFile
}

// NewClient creates new RealClient object
func NewClient() (*RealClient, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, errors.Wrap(err, "cannot create Docker Engine API client")
	}

	cfg, err := config.Load("")
	if err != nil {
		return nil, errors.Wrap(err, "cannot load docker config")
	}

	return &RealClient{
		cli: cli,
		cfg: cfg,
	}, nil
}

// PullImage pulls image from Docker Image Registry
func (c *RealClient) PullImage(image, tag string) (io.ReadCloser, error) {
	imageRef := fmt.Sprintf("%s:%s", image, tag)

	store := credentials.NewNativeStore(c.cfg, c.cfg.CredentialsStore)

	newAuths, err := store.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Docker credentials")
	}

	buf, err := json.Marshal(newAuths["https://quay.io"])
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal authConfig to JSON")
	}

	fmt.Println(string(buf))

	out, err := c.cli.ImagePull(context.Background(), imageRef, types.ImagePullOptions{
		RegistryAuth: base64.URLEncoding.EncodeToString(buf),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot pull image %q", imageRef)
	}

	return out, nil
}
