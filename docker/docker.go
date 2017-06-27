package docker

import (
	"encoding/json"
	"io"

	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/pkg/errors"
)

// PrintFunc represents the function to print log
type PrintFunc func(line string)

// PrintPullMessage decodes and prints docker pull message
func PrintPullMessage(out io.ReadCloser, printFunc PrintFunc) error {
	decoder := json.NewDecoder(out)

	for {
		var msg jsonmessage.JSONMessage

		err := decoder.Decode(&msg)

		if err == io.EOF {
			break
		}

		if err != nil {
			return errors.Wrap(err, "invalid message format")
		}

		if msg.Error != nil {
			return errors.Wrap(msg.Error, "error was raised during pulling image")
		}

		printFunc(msg.Status)
	}

	return nil
}
