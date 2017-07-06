package handler

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type DummyDockerClient struct{}

func (*DummyDockerClient) PullImage(image, tag string) (io.ReadCloser, error) {
	return ioutil.NopCloser(strings.NewReader("test")), nil
}

func TestWebhooksQuayHandler(t *testing.T) {
	testcases := []struct {
		reqBody        string
		imageWhitelist []string
		code           int
		resBody        string
		expectErr      bool
	}{
		{
			reqBody:        `{"build_id": "fake-build-id", "trigger_kind": "GitHub", "name": "imageup", "repository": "dtan4/imageup", "namespace": "dtan4", "docker_url": "quay.io/dtan4/imageup", "trigger_id": "1245634", "docker_tags": ["latest", "foo", "bar"], "build_name": "some-fake-build", "image_id": "1245657346", "trigger_metadata": {"default_branch": "master", "ref": "refs/heads/somebranch", "commit": "361babb16f96bcf8499194b4962d841bbb3629d9"}, "homepage": "https://quay.io/repository/dtan4/imageup/build/fake-build-id"}`,
			imageWhitelist: nil,
			code:           http.StatusAccepted,
			resBody:        "accepted",
			expectErr:      false,
		},
		{
			reqBody:        `{"build_id": "fake-build-id", "trigger_kind": "GitHub", "name": "imageup", "repository": "dtan4/imageup", "namespace": "dtan4", "docker_url": "quay.io/dtan4/imageup", "trigger_id": "1245634", "docker_tags": ["latest", "foo", "bar"], "build_name": "some-fake-build", "image_id": "1245657346", "trigger_metadata": {"default_branch": "master", "ref": "refs/heads/somebranch", "commit": "361babb16f96bcf8499194b4962d841bbb3629d9"}, "homepage": "https://quay.io/repository/dtan4/imageup/build/fake-build-id"}`,
			imageWhitelist: []string{"quay.io/dtan4/imageup"},
			code:           http.StatusAccepted,
			resBody:        "accepted",
			expectErr:      false,
		},
		{
			reqBody:        `{"build_id": "fake-build-id", "trigger_kind": "GitHub", "name": "imageup", "repository": "dtan4/imageup", "namespace": "dtan4", "docker_url": "quay.io/dtan4/imageup", "trigger_id": "1245634", "docker_tags": ["latest", "foo", "bar"], "build_name": "some-fake-build", "image_id": "1245657346", "trigger_metadata": {"default_branch": "master", "ref": "refs/heads/somebranch", "commit": "361babb16f96bcf8499194b4962d841bbb3629d9"}, "homepage": "https://quay.io/repository/dtan4/imageup/build/fake-build-id"}`,
			imageWhitelist: []string{"foobar"},
			code:           http.StatusForbidden,
			expectErr:      true,
		},
		{
			reqBody:        `invalid body`,
			imageWhitelist: nil,
			expectErr:      true,
		},
	}

	e := echo.New()
	e.Logger.SetLevel(log.OFF)

	for _, tc := range testcases {
		req := httptest.NewRequest(echo.POST, "/", strings.NewReader(tc.reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("DOCKER", &DummyDockerClient{})
		c.Set("IMAGE_WHITELIST", tc.imageWhitelist)

		err := WebhooksQuayHandler(c)

		if tc.expectErr {
			if err == nil {
				t.Errorf("error should be raised")
			}

			continue
		}

		if err != nil {
			t.Errorf("error should not be raised: %s", err)
		}

		if rec.Code != tc.code {
			t.Errorf("status code expected: %d, got: %d", http.StatusAccepted, tc.code)
		}

		b, err := ioutil.ReadAll(rec.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %s", err)
		}

		if string(b) != tc.resBody {
			t.Errorf("body expected: %q, got: %q", tc.resBody, string(b))
		}
	}
}
