package util

import "reflect"

func Insert(slice interface{}, index int, value interface{}) interface{} {

	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic("Interface slice not slice")
	}

	result := reflect.MakeSlice(v.Type(), v.Len()+1, v.Cap()+1)
	reflect.Copy(result, v.Slice(0, index))
	reflect.Copy(result.Slice(index+1, result.Len()), v.Slice(index, v.Len()))
	result.Index(index).Set(reflect.ValueOf(value))

	return result.Interface()
}

func InsertRune(s string, pos int, r rune) string {
	rs := []rune(s)
	out := append(rs[:pos], append([]rune{r}, rs[pos:]...)...)
	return string(out)
}
