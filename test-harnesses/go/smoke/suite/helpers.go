package suite

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func bytesAsBody(in []byte) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(in)))
}

func loadFixture(t *testing.T, operation string, scenario string) []byte {
	fileName := fmt.Sprintf("./fixtures/%s/%s.json", operation, scenario)
	file, err := os.Open(fileName)
	require.NoError(t, err)
	bytes, err := ioutil.ReadAll(file)
	require.NoError(t, err)
	return bytes
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

func prettyJSON(t *testing.T, in []byte) string {
	out := new(bytes.Buffer)
	err := json.Indent(out, in, "", "\t")
	require.NoError(t, err)
	return out.String()
}

func requireJSONSubset(t *testing.T, subsetArg, supersetArg interface{}, descr string) {
	subset := getBytes(t, subsetArg)
	superset := getBytes(t, supersetArg)
	isSubset := strings.Contains(string(superset), string(subset))
	if !isSubset {
		t.Fatalf(`
This JSON:

%s

Does not contain this subset:

%s
`, prettyJSON(t, superset), prettyJSON(t, subset))
	}
	require.Contains(t, string(superset), string(subset), descr)
}

type MockClient struct {
	mock.Mock
}

func (c *MockClient) Do(req *http.Request) (*http.Response, error) {
	args := c.Called(req)

	arg1, arg2 := args.Get(0), args.Error(1)
	if arg1 == nil {
		return nil, arg2
	}
	return arg1.(*http.Response), arg2
}
