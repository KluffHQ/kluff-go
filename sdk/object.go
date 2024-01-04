package sdk

import "github.com/kluff-com/kluff-go/dt"

type Object struct {
	cl     *dt.DbClient
	Base   dt.ObjectBase
	Fields []dt.Field
}
