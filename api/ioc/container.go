package ioc

//IContainer contract to built a register dependencie
type IContainer interface {
	RegisterDependencies() error
}
