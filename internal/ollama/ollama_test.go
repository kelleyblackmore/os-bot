// Tests for the ollama package.
package ollama

import (
	"testing"
)

func TestAsk(t *testing.T) {
	// This test requires a running Ollama service.
	prompt := "What is the color of the sky?"
	response, err := Ask(prompt)
	if err != nil {
		t.Fatalf("Ask() returned an error: %v", err)
	}
	if response == "" {
		t.Fatalf("Ask() returned an empty response")
	}
}
