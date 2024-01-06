package sdk

import (
	"context"

	"github.com/kluff-com/kluff-go/db"
	"google.golang.org/grpc"
)

type Interactor struct {
	cl db.DbClient
}

type ObjectSchema struct {
	Name          string      `json:"name"`
	CreatedByID   float64     `json:"created_by_id"`
	Fields        []*db.Field `json:"fields"`
	Owner         string      `json:"owner"`
	PluralLabel   string      `json:"plural_label"`
	SingularLabel string      `json:"singular_label"`
	TotalCount    float64     `json:"total_count"`
}

func NewDBInteractor(conn grpc.ClientConnInterface) Interactor {
	client := db.NewDbClient(conn)
	return Interactor{
		cl: client,
	}
}

// New Methods
func (i *Interactor) Object(name string) *Object {
	return &Object{
		cl: i.cl,
		Object: &db.Object{
			Base: &db.BaseObject{
				Name: name,
			},
		},
	}
}

func (i *Interactor) NewObject(obj *db.Object) *Object {
	return &Object{
		cl:     i.cl,
		Object: obj,
	}
}

func (i *Interactor) SendPing(context context.Context, ping *db.Ping) error {
	_, err := i.cl.SendPing(context, ping)
	return err
}

func (i *Interactor) GetObjects(context context.Context, q *db.ObjectQuery) ([]map[string]any, error) {
	data, err := i.cl.GetObjects(context, q)
	if err != nil {
		return nil, err
	}
	return data.ParseMapSlice()
}
