package docker

import (
	"testing"
)

func TestGetRegistry(t *testing.T) {
	testcases := []struct {
		image    string
		expected string
	}{
		{
			image:    "nginx",
			expected: "https://index.docker.io/v1/",
		},
		{
			image:    "foo/bar",
			expected: "https://index.docker.io/v1/",
		},
		{
			image:    "quay.io/dtan4/imageup",
			expected: "https://quay.io",
		},
		{
			image:    "012345678901.dkr.ecr.ap-northeast-1.amazonaws.com/foo",
			expected: "https://012345678901.dkr.ecr.ap-northeast-1.amazonaws.com",
		},
	}

	for _, tc := range testcases {
		if got := getRegistry(tc.image); got != tc.expected {
			t.Errorf("expected: %q, got: %q", tc.expected, got)
		}
	}
}
