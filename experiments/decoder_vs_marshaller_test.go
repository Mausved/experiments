package experiments

import (
	"bytes"
	"encoding/json"
	"github.com/test-go/testify/require"
	"testing"
)

var mans = map[string]string{
	"Vladislav": "Maitsvetov",
	"Yalchyn":   "Shukyurov",
}

func BenchmarkMansEncode(b *testing.B) {
	b.Run("encode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			mansInBytes, err := json.Marshal(mans)
			require.NoError(b, err)
			writer := bytes.NewBuffer(mansInBytes)
			b.StartTimer()

			json.NewEncoder(writer).Encode(mans)
		}
	})

	b.Run("marshal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			json.Marshal(mans)
		}
	})
}
