package config

import (
	"testing"

	"github.com/variadico/noti/cmd/noti/run"
)

func TestEvalFields(t *testing.T) {
	st := run.Stats{
		Cmd: "testing",
	}

	s := struct {
		Title string
		Num   int
	}{
		Title: "{{.Cmd}}",
		Num:   42,
	}

	ptrs := []interface{}{
		&s.Title,
		&s.Num,
	}

	if err := EvalFields(ptrs, st); err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if s.Title != st.Cmd {
		t.Error("Unexpected eval")
		t.Errorf("got: %v; want: %v", s.Title, st.Cmd)
	}
	if s.Num != 42 {
		t.Error("Unexpected eval")
		t.Errorf("got: %v; want: %v", s.Num, 42)
	}
}
