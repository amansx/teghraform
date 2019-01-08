package main

import "reflect"
import "teghraform/aws"

func GetInstanceFor(instancetype string, indexMap map[string]int, row []string) (string, interface{}) {
	var instance interface{}
	var instanceName string

	switch instancetype {
	case "aws.s3.bucket":
		instance = &aws.Bucket{}
	}

	if instance != nil {

		strct := reflect.ValueOf(instance)
		if strct.Kind() == reflect.Ptr {
			strct = reflect.Indirect(strct)
		}

		strctType := strct.Type()
		strcv := reflect.ValueOf(instance).Elem()

		for i := 0; i < strct.NumField(); i++ {
			fieldName := strctType.Field(i).Name
			if vindex := indexMap[fieldName]; vindex > 0 {
				value := row[vindex-1]
				if fieldName == "Name" {
					instanceName = value
				}
				strcv.FieldByName(fieldName).SetString(value)
			}
		}

	}

	return instanceName, instance
}
