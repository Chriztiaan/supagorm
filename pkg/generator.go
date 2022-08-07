package pkg

import (
	"reflect"

	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func GenerateTypeScriptModel(entity interface{}) {
	converter := typescriptify.New().Add(entity)

	// Not interested in backing these files up
	converter.BackupDir = ""

	err := converter.ConvertToFile("./models/" + DeriveEntityName(entity) + ".ts")

	if err != nil {
		panic(err.Error())
	}
}

func DeriveEntityName(entity interface{}) string {
	return reflect.TypeOf(entity).Elem().Name()
}
