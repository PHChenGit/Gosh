package alias

import (
	"strings"
	"testing"
)

func TestAlias(t *testing.T) {
	t.Run("Expend alias", func(t *testing.T) {
		mockAlias := NewAlias()
		mockAlias.SetAlias("gts", "git status")

		act, err := mockAlias.ExpendAlias("gts")

		if err != nil {
			t.Errorf("%v", err)
		}

		if !strings.EqualFold(act, "git status") {
			t.Errorf("got %s, but excepted %s", act, "git status")
		}
	})
}