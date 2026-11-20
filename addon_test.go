package main

import (
	"testing"

	"github.com/theDataCamp/ptz-controller/sdk/addon/addontest"
)

func TestOverThePipe(t *testing.T) {
	addontest.Start(t, spec()) // manifest valid and it answers initialize
}
