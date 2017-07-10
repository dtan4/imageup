package docker

import (
	"reflect"
	"testing"
)

func TestGetRegistries(t *testing.T) {
	testcases := []struct {
		image    string
		expected []string
	}{
		{
			image: "nginx",
			expected: []string{
				"https://index.docker.io/v1/",
			},
		},
		{
			image: "foo/bar",
			expected: []string{
				"https://index.docker.io/v1/",
			},
		},
		{
			image: "quay.io/dtan4/imageup",
			expected: []string{
				"quay.io",
				"https://quay.io",
			},
		},
		{
			image: "012345678901.dkr.ecr.ap-northeast-1.amazonaws.com/foo",
			expected: []string{
				"012345678901.dkr.ecr.ap-northeast-1.amazonaws.com",
				"https://012345678901.dkr.ecr.ap-northeast-1.amazonaws.com",
			},
		},
	}

	for _, tc := range testcases {
		if got := getRegistries(tc.image); !reflect.DeepEqual(got, tc.expected) {
			t.Errorf("expected: %q, got: %q", tc.expected, got)
		}
	}
}
