package log_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/foundation-13/gpr/pkg/log"
)

func TestLogWithContext(t *testing.T) {
	ctx := context.Background()

	assert.Equal(t, log.L, log.FromContext(ctx))

	newEntry := log.L.WithField("test", "test")
	ctx = log.WithLogger(ctx, newEntry)

	assert.Equal(t, newEntry, log.FromContext(ctx))
}