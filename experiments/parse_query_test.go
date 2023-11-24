package experiments

import (
	"encoding/json"
	"fmt"
	"github.com/test-go/testify/require"
	"net/http"
	"net/url"
	"testing"
)

func TestParseQuery(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r, err := http.NewRequest(http.MethodGet, "localhost:8080/api?start=1&end=2", nil)
		require.NoError(t, err)

		require.NoError(t, r.ParseForm())

		u := url.Values{
			"start": []string{"1"},
			"end":   []string{"2"},
		}
		require.Equal(t, u, r.Form)

		bytes, err := json.Marshal(r.Form)
		require.NoError(t, err)

		someStr := struct {
			Start []string
			End   []string
		}{}

		require.NoError(t, json.Unmarshal(bytes, &someStr))
		fmt.Println(someStr)
	})
}
