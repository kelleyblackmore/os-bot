// Package ollama handles interactions with the Ollama service.
package ollama

import (
	"github.com/xyproto/ollamaclient/v2"
	"github.com/xyproto/usermodel"
)

// Ask sends a prompt to the Ollama service and returns the response.
func Ask(prompt string) (string, error) {
	oc := ollamaclient.New(usermodel.GetTextGenerationModel())
	if err := oc.PullIfNeeded(); err != nil {
		return "", err
	}
	output, err := oc.GetOutput(prompt)
	if err != nil {
		return "", err
	}
	return output, nil
}
