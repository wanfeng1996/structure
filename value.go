package structure

import (
	"reflect"
	"slices"
	"strings"
	"structure/constant"
)

// GetStructIndex 确保传入的 查找 类型为结构体
func GetStructIndex(s interface{}, field string) ([]int, bool) {
	fds := strings.Split(field, constant.Point)
	res, ok := getIndex(getPtrStruct(reflect.ValueOf(s)), make([]int, 0, len(fds)), fds)
	if !ok {
		return res, ok
	}
	slices.Reverse(res)

	return res, ok
}

// getIndex 查找索引顺序为字段倒叙
func getIndex(v reflect.Value, indexes []int, field []string) ([]int, bool) {
	if len(field) < 1 {
		return indexes, false
	}
	var found bool
	for i := 0; i < v.Type().NumField(); i++ {
		if v.Type().Field(i).Name != field[0] {
			continue
		}
		vi := v.Field(i)
		found = true
	Start:
		if vi.Kind() == reflect.Ptr {
			vi = getPtrStruct(v.Field(i))
			goto Start
		}

		if vi.Kind() == reflect.Struct && len(field) > 1 {
			indexes, found = getIndex(vi, indexes, field[1:])
		}

		indexes = append(indexes, i)
		break
	}
	return indexes, found
}

// 获取指针类型结构体的值，如果为空，创建一个当前对象
func getPtrStruct(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr {
		return v
	}

	if v.IsZero() {
		return reflect.New(v.Type().Elem()).Elem()
	}

	return v.Elem()
}

// StructField 获取结构体里 单层字段值或多层字段值
func StructField[T Type](itr interface{}, name string) Iter[T] {
	if name == "" {
		return nil
	}

	if itr == nil {
		return nil
	}

	return fields[T](itr, strings.Split(name, constant.Point))
}

func fields[T Type](data interface{}, fields []string) Iter[T] {
	var (
		value      = getPtrStruct(reflect.ValueOf(data))
		res        = make(Iter[T], value.Len())
		idx, found = getIndex(getPtrStruct(value.Index(0)), make([]int, 0, len(fields)), fields)
		vi         *reflect.Value
	)

	if !found {
		return res
	}
	slices.Reverse(idx) //查找出来是倒叙，顺序掉正
	for i := 0; i < value.Len(); i++ {
		vi = getValue(getPtrStruct(value.Index(i)), idx)
		//如果为空，跳出循环
		if vi == nil {
			continue
		}

		res[i] = Address(vi.Interface().(T)) //添加当前值到对应位置

		//if vi.IsZero() {
		//	continue
		//}
		//res = append(res, vi.Interface().(T))
	}
	return res
}

func getValue(v reflect.Value, slice []int) *reflect.Value {
	for i := 0; i < len(slice); i++ {
		v = v.Field(slice[i])

		//if v.IsZero() || !v.IsValid() {
		//	return reflect.Zero(v.Type())
		//}

		if !v.IsValid() {
			return nil
		}

		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

	}
	return &v
}
