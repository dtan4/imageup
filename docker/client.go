package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strings"

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

func getRegistry(image string) string {
	ss := strings.Split(image, "/")

	if len(ss) == 3 {
		return "https://" + ss[0]
	}

	return "https://index.docker.io/v1/"
}

func (c *RealClient) getRegistryAuth(image string) (string, error) {
	var store credentials.Store

	if c.cfg.CredentialsStore == "" {
		store = credentials.NewFileStore(c.cfg)
	} else {
		store = credentials.NewNativeStore(c.cfg, c.cfg.CredentialsStore)
	}

	newAuths, err := store.GetAll()
	if err != nil {
		return "", errors.Wrap(err, "failed to get Docker credentials")
	}

	registry := getRegistry(image)

	if v, ok := newAuths[registry]; ok {
		buf, err := json.Marshal(v)
		if err != nil {
			return "", errors.Wrap(err, "failed to marshal authConfig to JSON")
		}

		return base64.URLEncoding.EncodeToString(buf), nil
	}

	return "", nil
}

// PullImage pulls image from Docker Image Registry
func (c *RealClient) PullImage(image, tag string) (io.ReadCloser, error) {
	imageRef := fmt.Sprintf("%s:%s", image, tag)

	registryAuth, err := c.getRegistryAuth(image)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get Docker registry credential for %s", image)
	}

	out, err := c.cli.ImagePull(context.Background(), imageRef, types.ImagePullOptions{
		RegistryAuth: registryAuth,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot pull image %q", imageRef)
	}

	return out, nil
}
