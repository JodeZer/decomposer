package helper

import "testing"

func TestStringSlice(t *testing.T) {
	strs := []string{"hello", "world"}
	slice := MakeStringSlice(0, 5)
	slice.AppendIf(nil, strs...)
	slice.AppendIf(nil, "!")
	slice.AppendIf(func(str string) bool {
		if str == "" {
			return true
		}
		return false
	}, "")

	t.Logf("%+v", slice.GetRaw())

	slice = slice.GetMendedSlice(func(str string) string {
		return "haha+" + str
	})

	t.Logf("%+v", slice.GetRaw())
}
