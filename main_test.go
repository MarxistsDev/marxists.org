//go:build integration
// +build integration

package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var serverStarted bool

func TestMain(m *testing.M) {
	if !serverStarted {
		StartServer()
		duration, _ := time.ParseDuration("5s")
		time.Sleep(duration)
	}

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestE2E_GetOneAuthor(t *testing.T) {
	assert.True(t, false)
}
