package main

import "testing"

func TestFind(t *testing.T) {
	want_word := "A-posteriori"
	want_definition := "\n\nA-POSTERIORI, [L. posterior, after.]\nArguments a posteriori, are drawn from effect, consequences or facts; in opposition to reasoning a priori, or from causes previously known."
	got, err := find(want_word)
	if err != nil {
		t.Errorf("error in find()")
	}
	if got.Word != want_word {
		t.Errorf("find() = %q, word %q", got.Word, want_word)
	}
	if got.Definition != want_definition {
		t.Errorf("find() = %q, def %q", got.Definition, want_definition)
	}
}

func TestNotFound(t *testing.T) {
	got, err := find("NotFound")
	if err != nil {
		t.Errorf("error in find()")
	}
	if got.Word != "" {
		t.Errorf("find() = %q, word", got.Word)
	}
}

func TestFormatArgument(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "Empty string", input: "", want: ""},
		{name: "String with lowercase first letter", input: "hello", want: "Hello"},
		{name: "String with uppercase first letter", input: "WORLD", want: "World"},
		{name: "Title case", input: "World", want: "World"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := formatArgument(tc.input)
			if err != nil {
				t.Fail()
			}
			if got != tc.want {
				t.Errorf("wrong output: in: %q got: %q want: %q", tc.input, got, tc.want)
			}
		})
	}
}

func FuzzFormatArgument(f *testing.F) {
    f.Add("hello")
    f.Fuzz(func(t *testing.T, arg string) {
        _, err := formatArgument(arg)
        if err != nil {
            t.Fatalf("Formatted argument different than expected %v", arg)
        }
    })
}

