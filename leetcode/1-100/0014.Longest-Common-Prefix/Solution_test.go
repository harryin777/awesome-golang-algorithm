package Solution

import "testing"

func TestSolution1(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		data := []string{
			"flower",
			"flow",
			"flight",
		}
		got := longestCommonPrefix(data)
		want := "fl"
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		data := []string{
			"aaa",
			"aa",
			"aaa",
		}
		got := longestCommonPrefix2(data)
		want := "aa"
		if got != want {
			t.Error("GOT:", got, "WANT:", want)
		}
	})
}
