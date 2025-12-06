package assets_test

import (
	"testing"

	"github.com/malivvan/cui/editor/assets"
)

func TestThemes(t *testing.T) {
	t.Log(len(assets.Assets.Themes))
	t.Log(len(assets.Assets.Syntax))

}
