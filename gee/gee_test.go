package gee

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRouterGroup(t *testing.T) {
	r := New()
	v1 := r.Group("/v1")
	v2 := v1.Group("/v2")
	v3 := v2.Group("/v3")
	assert.Equal(t, v2.prefix, "/v1/v2")
	assert.Equal(t, v3.prefix, "/v1/v2/v3")
}
