package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func bytesAsBody(in []byte) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(in)))
}

func getBytes(t *testing.T, in interface{}) []byte {
	var b []byte
	if val, ok := in.([]byte); ok {
		b = val
	} else {
		val, err := json.Marshal(in)
		require.NoError(t, err)
		b = val
	}

	out := new(bytes.Buffer)
	err := json.Compact(out, b)
	require.NoError(t, err)
	return out.Bytes()
}

func PrettyJSON(t *testing.T, in []byte) string {
	out := new(bytes.Buffer)
	err := json.Indent(out, in, "", "\t")
	require.NoError(t, err)
	return out.String()
}

func RequireJSONSubset(t *testing.T, subsetArg, supersetArg interface{}, descr string) {
	subset := getBytes(t, subsetArg)
	superset := getBytes(t, supersetArg)
	isSubset := strings.Contains(string(superset), string(subset))
	if !isSubset {
		t.Fatalf(`
This JSON:

%s

Does not contain this subset:

%s
`, PrettyJSON(t, superset), PrettyJSON(t, subset))
	}
	require.Contains(t, string(superset), string(subset), descr)
}
