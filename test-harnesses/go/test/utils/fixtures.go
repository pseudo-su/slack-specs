package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func LoadFixture(t *testing.T, operation, schema, scenario string) []byte {
	wd, err := os.Getwd()
	require.NoError(t, err)
	name := fmt.Sprintf("%s.%s.json", schema, scenario)
	fileName := filepath.Join(wd, "fixtures", operation, name)
	// fileName := fmt.Sprintf("./fixtures/%s/%s.json", operation, scenario)
	file, err := os.Open(fileName)
	require.NoError(t, err)
	bytes, err := ioutil.ReadAll(file)
	require.NoError(t, err)
	return bytes
}
