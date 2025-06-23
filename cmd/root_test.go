package cmd

import (
	"testing"
)

func TestInit(t *testing.T) {
	if rootCmd == nil {
		t.Fatal("serverCmd should be defined")
	}
	if rootCmd.Use != "k8s-controller" {
		t.Errorf("expected command use 'k8s-controller', got %s", rootCmd.Use)
	}
	logLevelFlag := rootCmd.Flags().Lookup("log-level")
	if logLevelFlag == nil {
		t.Error("expected 'log-level' flag to be defined")
	}
	portFlag := rootCmd.Flags().Lookup("port")
	if portFlag == nil {
		t.Error("expected 'port' flag to be defined")
	}
}
