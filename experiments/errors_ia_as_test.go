package experiments

import (
	"errors"
	"fmt"
	"github.com/test-go/testify/assert"
	"github.com/test-go/testify/require"
	"testing"
)

type customErrorWithValueReceiver struct {
	message string
}

func (e customErrorWithValueReceiver) Error() string {
	return e.message
}

func (e customErrorWithValueReceiver) Is(target error) bool {
	_, ok := target.(customErrorWithValueReceiver)
	return ok
}

type customErrorWithPtrReceiver struct {
	message string
}

func (e *customErrorWithPtrReceiver) Error() string {
	return e.message
}

func (e *customErrorWithPtrReceiver) Is(target error) bool {
	_, ok := target.(*customErrorWithPtrReceiver)
	return ok
}

func TestErrorsIs(t *testing.T) {
	t.Run("value receiver, without wrapping", func(t *testing.T) {
		err := customErrorWithValueReceiver{message: "test"}
		require.True(t, errors.Is(err, &customErrorWithValueReceiver{message: "test"}))
	})
	t.Run("value receiver, with wrapping", func(t *testing.T) {
		err := fmt.Errorf("some error: %w", customErrorWithValueReceiver{message: "test"})
		require.True(t, errors.Is(err, &customErrorWithValueReceiver{message: "test"}))
	})

	t.Run("ptr receiver, without wrapping", func(t *testing.T) {
		err := &customErrorWithPtrReceiver{message: "test"}
		require.True(t, errors.Is(err, &customErrorWithPtrReceiver{message: "test"}))
	})
	t.Run("ptr receiver, with wrapping", func(t *testing.T) {
		err := fmt.Errorf("some error: %w", &customErrorWithPtrReceiver{message: "test"})
		require.True(t, errors.Is(err, &customErrorWithPtrReceiver{message: "test"}))
	})
}

func TestErrorsAs(t *testing.T) {
	t.Run("without wrapping, ptr-ptr", func(t *testing.T) {
		expected := &customErrorWithPtrReceiver{message: "test"}
		var actual *customErrorWithPtrReceiver

		require.True(t, errors.As(expected, &actual))
		assert.Equal(t, expected, actual)
	})
	t.Run("with wrapping, ptr-noptr", func(t *testing.T) {
		expected := fmt.Errorf("some error: %w", &customErrorWithPtrReceiver{message: "test"})
		var actual *customErrorWithPtrReceiver

		require.True(t, errors.As(expected, &actual))
		assert.Equal(t, &customErrorWithPtrReceiver{message: "test"}, actual)
	})
	t.Run("with wrapping", func(t *testing.T) {
		expected := fmt.Errorf("some error: %w", &customErrorWithPtrReceiver{message: "test"})
		var actual *customErrorWithPtrReceiver
		require.True(t, errors.As(expected, &actual))
		assert.Equal(t, &customErrorWithPtrReceiver{message: "test"}, actual)
	})
}
