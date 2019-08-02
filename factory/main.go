package main

import (
	"fmt"
	"reflect"

	"github.com/tensor-programming/pattern-tutorial/factory"
)

func SetupConstructors(env string) (factory.Database, factory.FileSystem) {
	fs := factory.AbstractFactory("filesystem")
	db := factory.AbstractFactory("database")

	return db(env).(factory.Database),
		fs(env).(factory.FileSystem)
}

func main() {
	env1 := "production"
	env2 := "development"

	db1, fs1 := SetupConstructors(env1)
	db2, fs2 := SetupConstructors(env2)

	// db1 := factory.DatabaseFactory(env1)
	// db2 := factory.DatabaseFactory(env2)

	db1.PutData("test", "this is mongodb")
	fmt.Println(db1.GetData("test"))

	db2.PutData("test", "this is sqlite")
	fmt.Println(db2.GetData("test"))

	fs1.CreateFile("../example/testntfs.txt")
	fmt.Println(fs1.FindFile("../example/testntfs.txt"))

	fs2.CreateFile("../example/testext4.txt")
	fmt.Println(fs2.FindFile("../example/testext4.txt"))

	fmt.Println(reflect.TypeOf(db1).Name())
	fmt.Println(reflect.TypeOf(&db1).Elem())
	fmt.Println(reflect.TypeOf(db2).Name())
	fmt.Println(reflect.TypeOf(&db2).Elem())

	fmt.Println(reflect.TypeOf(fs1).Name())
	fmt.Println(reflect.TypeOf(&fs1).Elem())
	fmt.Println(reflect.TypeOf(fs2).Name())
	fmt.Println(reflect.TypeOf(&fs2).Elem())
}
