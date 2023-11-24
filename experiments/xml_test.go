package experiments

import (
	"encoding/xml"
	"fmt"
	"github.com/test-go/testify/require"
	"os"
	"testing"
)

type someTypeForXml struct {
	Hello  string `xml:"hello"`
	Name   string `xml:"name"`
	Nested struct {
		Number      int `xml:"number,omitempty"`
		notImported int
	}
}

type nested2 struct {
	XMLName xml.Name `xml:"Response"`
	*someTypeForXml
}

func TestXml(t *testing.T) {
	t.Run("header", func(t *testing.T) {
		xml.NewEncoder(os.Stdout).Encode(xml.Header)
	})

	t.Run("someTypeForXml", func(t *testing.T) {
		s := &someTypeForXml{
			Hello: "world",
			Name:  "byzuka",
			Nested: struct {
				Number      int `xml:"number,omitempty"`
				notImported int
			}{
				notImported: 100,
			},
		}

		nested := &nested2{
			someTypeForXml: s,
		}

		require.NoError(t, xml.NewEncoder(os.Stdout).Encode(nested))
		fmt.Println()

		errXml := struct {
			XMLName xml.Name `xml:"response"`
			Result  int      `xml:"result"`
		}{
			Result: 0,
		}

		b, err := xml.Marshal(errXml)
		require.NoError(t, err)

		fmt.Println(string(b))

		//require.NoError(t, xml.NewEncoder(os.Stdout).Encode(xml.Header))
		//require.NoError(t, xml.NewEncoder(os.Stdout).Encode(s))
		//fmt.Println()
		//require.NoError(t, xml.NewEncoder(os.Stdout).Encode(nested))
	})
}
