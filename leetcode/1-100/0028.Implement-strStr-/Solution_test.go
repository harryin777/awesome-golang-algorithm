package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := strStr("hello", "ll")
		want := 2
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		got := strStr2("mississippi", "issip")
		want := 4
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}
