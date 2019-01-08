package main

import "fmt"
import "reflect"
import "teghraform/aws"

func GetInstanceFor(instancetype string, indexMap map[string]int, row []string) interface{} {

	var params []string
	var instance interface{}
	instance = aws.Bucket{}

	val := reflect.ValueOf(instance) // could be any underlying type

	// if its a pointer, resolve its value
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	// now we grab our values as before (note: I assume table name should come from the struct type)
	structType := val.Type()
	tableName := structType.Name()
	params = append(params, tableName)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)

		fmt.Println(field)
		// tag := field.Tag

		// fieldName := field.Name
		// fieldType := tag.Get("sql_type")
		// fieldTags := tag.Get("sql_tag")

		// paramstring := fieldName + " " + fieldType + " " + fieldTags
		// params = append(params, paramstring)
	}

	// switch instancetype {
	// case "aws.s3.bucket":
	// 	instance = aws.Bucket{}
	// }

	// fields := reflect.ValueOf(&instance)
	// //.Elem()
	// fieldTypes := fields.Type()

	// for f := 0; f < fields.NumField(); f++ {

	// 	// field := fields.Field(f).Interface()
	// 	fieldName := fieldTypes.Field(f).Name

	// 	if indexMap[fieldName] > 0 {
	// 		if f := fields.FieldByName(fieldName); f.IsValid() && f.CanSet() {
	// 			f.SetString("Aman")
	// 		}
	// 	}
	// }

	// for _, item := range row {
	// 	fmt.Println(item)
	// }

	// return instance

	return nil

}
