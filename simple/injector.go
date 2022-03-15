//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabaseMongoDB, NewDatabasePostgreSQL, NewDatabaseRepository)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// Wrong
//func InitializedHelloService() *HelloService {
//	wire.Build(NewHelloService, NewSayHelloImpl)
//	return nil
//}

var HelloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(HelloSet, NewHelloService)
	return nil
}

var FooBarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

func InirializedFooBar() *FooBar {
	wire.Build(FooBarSet, wire.Struct(new(FooBar), "*"))
	return nil
}

//var FooBarValueSet = wire.NewSet(
//	wire.Value(&Foo{}),
//	wire.Value(&Bar{}),
//)
var FooValue = &Foo{}
var BarValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	//wire.Build(FooBarValueSet, wire.Struct(new(FooBar), "*"))
	wire.Build(wire.Value(FooValue), wire.Value(BarValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(NewApplication, wire.FieldsOf(new(*Application), "Configuration"))
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
