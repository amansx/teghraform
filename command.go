package main

import "strings"
import "reflect"
import "teghraform/aws"

func GetInstanceFor(indexMap map[string]int, row []string) (string, interface{}) {
	var instance interface{}
	var instanceType string
	var instanceName string

	instancetypeIndex := indexMap[OBJECT_TYPE]
	if instancetypeIndex > 0 {
		instanceType = row[instancetypeIndex-1]
		instanceType = strings.TrimSpace(strings.ToLower(instanceType))
	} else {
		return instanceName, instance
	}

	switch instanceType {
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
				if fieldName == OBJECT_NAME {
					instanceName = value
				}
				strcv.FieldByName(fieldName).SetString(value)
			}
		}

	}

	return instanceName, instance
}
