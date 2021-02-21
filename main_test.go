package main

import (
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	t.Run("Alias", func(t *testing.T) {
		givenInput := "alias gts='git status'"
		act := parseArgs(givenInput)
		excepted := []string{"alias", "gts='git status'"}

		if !reflect.DeepEqual(act, excepted) {
			t.Errorf("got %v, but excepted %v", act, excepted)
		}
	})

	t.Run("Export command", func(t *testing.T) {
		givenInput := "export google=Google fb=Facebook"
		act := parseArgs(givenInput)
		excepted := []string{"export", "google=Google fb=Facebook"}

		if !reflect.DeepEqual(act, excepted) {
			t.Errorf("got %v, but excepted %v", act, excepted)
		}
	})

	t.Run("No args", func(t *testing.T) {
		givenInput := "ls"
		act := parseArgs(givenInput)
		exceptedInputSepLen := 1
		excepted := "ls"

		if len(act) != exceptedInputSepLen {
			t.Errorf("got %d, but excepted len %d", len(act), exceptedInputSepLen)
		}

		if act[0] != excepted {
			t.Errorf("got %s, but excepted %s", act[0], excepted)
		}
	})

	t.Run("Has args but not set it", func(t *testing.T) {
		givenInput := "ls -a -l"
		act := parseArgs(givenInput)
		excepted := []string{"ls", "-a", "-l"}

		if !reflect.DeepEqual(act, excepted) {
			t.Errorf("got %v, but excepted %v", act, excepted)
		}
	})
}
