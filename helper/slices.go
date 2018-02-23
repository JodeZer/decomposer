package helper

import "github.com/JodeZer/decomposer/helper/helperF"

type StringSlice struct {
	raw []string
}

func MakeStringSlice(param ...int) *StringSlice {
	if len(param) > 2 || len(param) == 0 {
		panic("fuck u")
	}
	p1, p2 := param[0], param[0]
	if len(param) == 2 {
		p2 = param[1]
	}
	return &StringSlice{
		raw: make([]string, p1, p2),
	}
}

func MakeStringSliceFromRaw(raw []string) *StringSlice {
	if raw == nil {
		raw = make([]string, 0)
	}
	return &StringSlice{
		raw: raw,
	}
}

func (ss *StringSlice) Append(strs ...string) {
	ss.AppendIf(nil, strs...)
}

func (ss *StringSlice) AppendSlice(slice *StringSlice) {
	slice.Range(func(str string) {
		ss.Append(str)
	})
}

func (ss *StringSlice) AppendIf(filter helperF.StringFilter, strs ...string) {
	MakeStringSliceFromRaw(strs).Range(func(str string) {
		if !filter.Something()(str) {
			ss.raw = append(ss.raw, str)
		}
	})
}

func (ss *StringSlice) Range(rangers ...helperF.StringRanger) {
	for _, ranger := range rangers {
		for _, str := range ss.GetRaw() {
			ranger.Something()(str)
		}
	}
}

func (ss *StringSlice) GetRaw() []string {
	return ss.raw
}

// func (ss *StringSlice) getMendedRaw(menders ...helperF.StringMender) []string {
// 	res := make()
// 	return nil
// }

// func (ss *StringSlice) getFilteredRaw(filters ...helperF.StringFilter) []string {

// }

func (ss *StringSlice) GetFilteredSlice(filters ...helperF.StringFilter) *StringSlice {

	res := MakeStringSlice(0, len(ss.GetRaw()))

	ss.Range(func(str string) {
		for _, f := range filters {
			if !f.Something()(str) {
				res.Append(str)
			}
		}
	})

	return res
}

func (ss *StringSlice) GetMendedSlice(menders ...helperF.StringMender) *StringSlice {

	res := MakeStringSlice(0, len(ss.GetRaw()))

	ss.Range(func(str string) {
		for _, f := range menders {
			res.Append(f.Something()(str))
		}
	})

	return res
}

func (ss *StringSlice) Len() int {
	return len(ss.GetRaw())
}

func (ss *StringSlice) Cap() int {
	return cap(ss.GetRaw())
}
