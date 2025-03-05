package tools

import (
	"fmt"
	"reflect"
	"strings"
)

/*
	type MDB[T any] struct {
		DB    string
		Table string
	}

	func (x *MDB[T]) SetDB(db string) {
		x.DB = db
		x.Table = "users"
	}

	type User struct {
		MDB[User] `json:"-" bson:"-"`
		Phone     string    `json:"phone" bson:"phone"`           //手机号码
		Passwd    string    `json:"passwd" bson:"passwd"`         //密码
		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"` //更新时间
	}

	func main() {
		model := User{}
		model.SetDB("hotel")
		model.Phone = "136"

		// Demonstrate setting fields using the unified SetField function
		if err := SetField(&model, "Phone", "1234567890"); err != nil {
			fmt.Println("Error:", err)
		}

		model.Find("1")

		println("DB=>", model.DB)
		println("Table=>", model.Table)
	}
*/
func SetField(source interface{} /*must be a interface*/, fieldName string, fieldValue string) {
	v := reflect.ValueOf(source).Elem()
	field := v.FieldByName(fieldName)

	if field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
		field.SetString(fieldValue)
	}
}

/*
	if err := SetField2(&model, "MDB.DB", "exampleDB"); err != nil {
		fmt.Println("Error:", err)
	}
*/
func SetFieldNest(source interface{}, fieldPath string, fieldValue string) error {
	v := reflect.ValueOf(source).Elem()

	// Split the field path to handle nested fields
	fields := strings.Split(fieldPath, ".")

	for i, fieldName := range fields {
		if v.Kind() == reflect.Struct {
			v = v.FieldByName(fieldName)
			if !v.IsValid() {
				return fmt.Errorf("field %s not found", fieldName)
			}
			if i < len(fields)-1 {
				// If it's not the last field, ensure it's a struct
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				if v.Kind() != reflect.Struct {
					return fmt.Errorf("field %s is not a struct", fieldName)
				}
			}
		} else {
			return fmt.Errorf("field %s is not a struct", fieldName)
		}
	}

	if v.CanSet() && v.Kind() == reflect.String {
		v.SetString(fieldValue)
		return nil
	}

	return fmt.Errorf("field %s cannot be set", fieldPath)
}
