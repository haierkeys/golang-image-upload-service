package convert

import (
	"encoding/json"
	"reflect"
)

// CopyStruct
// dst 目标结构体，src 源结构体
// 它会把src与dst的相同字段名的值，复制到dst中
func StructAssign(src interface{}, dst interface{}) interface{} {
	bVal := reflect.ValueOf(dst).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(src).Elem() //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
	return dst
}

/**
 * @Description: 结构体map互转
 * @param param interface{} 需要被转的数据
 * @param data interface{} 转换完成后的数据  需要用引用传进来
 * @return []string{}
 */
func StrucToMap(param any, data any) any {
	str, _ := json.Marshal(param)
	_ = json.Unmarshal(str, data)
	return data
}
