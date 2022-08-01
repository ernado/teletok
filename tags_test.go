package teletok

import (
	"embed"
	_ "embed"
	"path"
	"testing"

	"github.com/stretchr/testify/require"

	"teletok/internal/oas"
)

var (
	//go:embed _testdata/responses/*.json
	responses embed.FS
)

func TestTags(t *testing.T) {
	base := "_testdata/responses"
	items, err := responses.ReadDir(base)
	require.NoError(t, err)
	for _, item := range items {
		t.Run(item.Name(), func(t *testing.T) {
			var data oas.Data
			b, err := responses.ReadFile(path.Join(base, item.Name()))
			require.NoError(t, err)
			require.NoError(t, data.UnmarshalJSON(b))
		})
	}
}
