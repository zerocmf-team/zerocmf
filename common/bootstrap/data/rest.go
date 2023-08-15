/**
** @创建时间: 2022/3/13 18:03
** @作者　　: return
** @描述　　:
 */

package data

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

type Rest struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	StatusCode *int        `json:"-"`
	okBytes    bool        `json:"-"`
}

type result struct {
	Data       interface{}
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"update_time"`
}

type H map[string]interface{}

func hasKey(arr []string, key string) bool {
	for _, t := range arr {
		if t == key {
			return true
		}
	}
	return false
}

var timeMap = map[string]string{
	"CreatedAt": "createdTime",
	"UpdatedAt": "updatedTime",
}

//func formatData(data interface{}) interface{} {
//	t := reflect.TypeOf(data)
//	json := make(map[string]interface{}, t.Size())
//	if t.Kind() == reflect.Struct {
//		v := reflect.ValueOf(data)
//		for i := 0; i < t.NumField(); i++ {
//			field := t.Field(i)
//			value := v.Field(i)
//			name := field.Name
//			fmt.Println("name", name)
//			tag := field.Tag.Get("json")
//			tagArr := strings.Split(tag, ",")
//			if field.IsExported() {
//				if len(tagArr) > 0 {
//					name = tagArr[0]
//				}
//				if field.Name == "CreatedAt" || field.Name == "UpdatedAt" {
//					// 设置 CreateTime 和 UpdateTime 字段值为当前时间
//					currentTime := time.Now().Format("2006-01-02 15:04:05")
//					_name := timeMap[field.Name]
//					json[_name] = currentTime
//				}
//				if !(hasKey(tagArr, "omitempty") && value.IsZero()) {
//					json[name] = value.Interface()
//				}
//			}
//		}
//	}
//	return json
//}

func formatData(data interface{}) interface{} {

	if data != nil {

		objValue := reflect.ValueOf(data)
		objType := objValue.Type()

		if objType.Kind() == reflect.Slice && objValue.Len() == 0 {
			return []string{}
		}

		if !(objType.Kind() == reflect.Ptr || objType.Kind() == reflect.Struct) {
			return data
		}

		if objType.Kind() == reflect.Ptr {
			objValue = objValue.Elem()
			objType = objValue.Type()
		}

		json := make(map[string]interface{}, 0)

		for i := 0; i < objValue.NumField(); i++ {
			field := objValue.Field(i)
			fieldName := objType.Field(i).Name
			if objType.Field(i).PkgPath != "" {
				continue // 如果字段的 PkgPath 不为空，则表示字段未导出，忽略该字段
			}
			fieldTag := objType.Field(i).Tag.Get("json") // 获取字段的标签值
			if fieldTag == "-" {
				continue // 如果标签值为空或为"-"，则忽略该字段
			}
			tags := strings.Split(fieldTag, ",")
			tag := tags[0]

			if field.Kind() == reflect.Ptr {
				if field.Elem().Kind() == reflect.Slice {
					if field.Len() == 0 && strings.Contains(fieldTag, "omitempty") {
						continue // 如果切片字段为空且标签包含"omitempty"，则忽略该字段
					}

					// 处理指向切片的指针类型
					sliceValue := field.Elem()
					sliceLen := sliceValue.Len()
					sliceValues := make([]interface{}, sliceLen)

					for j := 0; j < sliceLen; j++ {
						sliceValues[j] = sliceValue.Index(j).Interface()
					}
					json[fieldName] = sliceValues
				} else if field.Elem().Kind() == reflect.Struct {
					if tag != "" {
						json[tag] = field.Interface()
					} else {
						nestedMap := formatData(field.Interface())
						childJson, ok := nestedMap.(map[string]interface{})
						if ok {
							for key, value := range childJson {
								json[key] = value
							}
						}
					}
				}
			} else if field.Kind() == reflect.Struct {
				// 处理嵌套结构体
				nestedMap := formatData(field.Interface())
				childJson, ok := nestedMap.(map[string]interface{})
				if ok {
					for key, value := range childJson {
						json[key] = value
					}
				}
			} else {
				name := fieldName
				if tag != "" {
					name = tag
				}
				if field.Interface() != nil {
					json[name] = field.Interface()
					if fieldName == "CreatedAt" || fieldName == "UpdatedAt" {
						// 设置 CreateTime 和 UpdateTime 字段值为当前时间
						key := timeMap[fieldName]
						currentTime := time.Now().Format("2006-01-02 15:04:05")
						json[key] = currentTime
					}
				}

			}
		}
		return json
	}
	return nil
}

func (r *Rest) Success(msg string, data interface{}) (resp *Rest) {
	r.Code = 1
	r.Msg = msg
	r.Data = formatData(data)
	resp = r
	return
}

func (r *Rest) Error(msg string, data interface{}) (resp *Rest) {
	r.Code = 0
	r.Msg = msg
	r.Data = formatData(data)
	resp = r
	return
}

func (r *Rest) ToBytes(msg string, data interface{}) (bytes []byte) {
	r.Code = 1
	r.Msg = msg
	r.Data = formatData(data)
	bytes, _ = json.Marshal(r)
	return
}

func (r *Rest) OkBytes() bool {
	return r.okBytes
}

func (r *Rest) String(msg string) (resp *Rest) {
	r.Msg = msg
	r.okBytes = true
	resp = r
	return
}
