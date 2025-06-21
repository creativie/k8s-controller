package cmd

import (
	"testing"
)

func TestServerCommandDefined(t *testing.T) {

	portFlag := rootCmd.Flags().Lookup("port")
	if portFlag == nil {
		t.Error("expected 'port' flag to be defined")
	}
}
