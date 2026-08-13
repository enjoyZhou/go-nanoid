package gonanoid_test

import (
	"strings"
	"testing"
	"unicode/utf8"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func requireNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGenerate(t *testing.T) {
	t.Run("empty alphabet", func(t *testing.T) {
		alphabet := ""
		_, err := gonanoid.Generate(alphabet, 32)
		if err == nil {
			t.Fatal("expected an error for an empty alphabet")
		}
	})

	t.Run("256 character alphabet", func(t *testing.T) {
		alphabet := strings.Repeat("a", 256)
		id, err := gonanoid.Generate(alphabet, 32)
		requireNoError(t, err)
		if want := strings.Repeat("a", 32); id != want {
			t.Fatalf("got %q, want %q", id, want)
		}
	})

	t.Run("alphabet longer than 256 characters", func(t *testing.T) {
		alphabet := strings.Repeat("a", 257)
		_, err := gonanoid.Generate(alphabet, 32)
		if err == nil {
			t.Fatal("expected an error for an alphabet longer than 256 characters")
		}
	})

	t.Run("256 unicode character alphabet", func(t *testing.T) {
		alphabet := strings.Repeat("🚀", 256)
		id, err := gonanoid.Generate(alphabet, 6)
		requireNoError(t, err)
		if got := utf8.RuneCountInString(id); got != 6 {
			t.Fatalf("got %d characters, want 6", got)
		}
	})

	t.Run("negative ID length", func(t *testing.T) {
		_, err := gonanoid.Generate("abcdef", -1)
		if err == nil {
			t.Fatal("expected an error for a negative ID length")
		}
	})

	t.Run("zero ID length", func(t *testing.T) {
		id, err := gonanoid.Generate("abcdef", 0)
		requireNoError(t, err)
		if id != "" {
			t.Fatalf("got %q, want an empty ID", id)
		}
	})

	t.Run("happy path", func(t *testing.T) {
		alphabet := "abcdef"
		id, err := gonanoid.Generate(alphabet, 6)
		requireNoError(t, err)
		if len(id) != 6 {
			t.Fatalf("got ID length %d, want 6", len(id))
		}
		for _, r := range id {
			if !strings.ContainsRune(alphabet, r) {
				t.Fatalf("ID contains %q outside alphabet %q", r, alphabet)
			}
		}
	})

	t.Run("works with unicode", func(t *testing.T) {
		alphabet := "🚀💩🦄🤖"
		id, err := gonanoid.Generate(alphabet, 6)
		requireNoError(t, err)
		if got := utf8.RuneCountInString(id); got != 6 {
			t.Fatalf("got ID length %d, want 6", got)
		}
		for _, r := range id {
			if !strings.ContainsRune(alphabet, r) {
				t.Fatalf("ID contains %q outside alphabet %q", r, alphabet)
			}
		}
	})
}

func TestNew(t *testing.T) {
	t.Run("matches Nano ID 6 defaults", func(t *testing.T) {
		if gonanoid.DefaultSize != 21 {
			t.Fatalf("got default size %d, want 21", gonanoid.DefaultSize)
		}
		const wantAlphabet = "useandom-26T198340PX75pxJACKVERYMINDBUSHWOLF_GQZbfghjklqvwyzrict"
		if gonanoid.URLAlphabet != wantAlphabet {
			t.Fatalf("got URL alphabet %q, want %q", gonanoid.URLAlphabet, wantAlphabet)
		}
		if len(gonanoid.URLAlphabet) != 64 {
			t.Fatalf("got URL alphabet length %d, want 64", len(gonanoid.URLAlphabet))
		}
	})

	t.Run("negative ID length", func(t *testing.T) {
		_, err := gonanoid.New(-1)
		if err == nil {
			t.Fatal("expected an error for a negative ID length")
		}
	})

	t.Run("happy path", func(t *testing.T) {
		id, err := gonanoid.New()
		requireNoError(t, err)
		if len(id) != gonanoid.DefaultSize {
			t.Fatalf("got ID length %d, want %d", len(id), gonanoid.DefaultSize)
		}
		for _, char := range id {
			if !strings.ContainsRune(gonanoid.URLAlphabet, char) {
				t.Fatalf("ID contains %q outside the URL alphabet", char)
			}
		}
	})

	t.Run("custom length", func(t *testing.T) {
		id, err := gonanoid.New(6)
		requireNoError(t, err)
		if len(id) != 6 {
			t.Fatalf("got ID length %d, want 6", len(id))
		}
	})

	t.Run("zero length", func(t *testing.T) {
		id, err := gonanoid.New(0)
		requireNoError(t, err)
		if id != "" {
			t.Fatalf("got %q, want an empty ID", id)
		}
	})

	t.Run("unexpected parameters", func(t *testing.T) {
		_, err := gonanoid.New(1, 2)
		if err == nil {
			t.Fatal("expected an error for multiple size parameters")
		}
	})
}
