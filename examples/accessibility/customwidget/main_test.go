package main

import (
	"testing"

	"github.com/belluikit/bellui/core"
)

//docs:start test
func TestRatingIsAccessible(t *testing.T) {
	r := &Rating{Label: "Food", Value: 3}

	node, ok := core.SnapshotAccessible(r)
	if !ok {
		t.Fatal("Rating does not implement core.Accessible")
	}
	if node.Role != core.RoleSlider || node.Label != "Food" || node.Value != "3 of 5 stars" {
		t.Errorf("got role %v, label %q, value %q", node.Role, node.Label, node.Value)
	}
}

//docs:end test
