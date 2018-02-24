package helperF

func ConvertStringInterface(str ...string) []interface{} {
	res := make([]interface{}, len(str))
	for i, one := range str {
		res[i] = one
	}
	return res
}
