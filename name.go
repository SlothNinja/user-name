package name

import (
	"errors"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

var ErrNameInUse = errors.New("Name already in use.")

type Name struct {
	Key *datastore.Key `datastore:"__key__"`
	// ID        string `gae:"$id"`
	// Kind      string `gae:"$kind"`
	GoogleID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

const kind = "UserName"

func New(name string) *Name {
	name = strings.ToLower(name)
	return &Name{Key: datastore.NameKey(kind, name, nil)}
	// return &Name{Kind: kind}
}

func ByName(c *gin.Context, name string) (*Name, error) {
	dsClient, err := datastore.NewClient(c, "")
	if err != nil {
		return nil, err
	}

	n := New(name)
	err = dsClient.Get(c, n.Key, n)
	return n, err
}

//func NewKey(c *restful.Context, name string) *datastore.Key {
//	return datastore.NewKey(c, kind, name, 0, nil)
//}

func IsUnique(c *gin.Context, name string) bool {
	_, err := ByName(c, name)
	return err == datastore.ErrNoSuchEntity
	// return ByName(c, name, New()) == datastore.ErrNoSuchEntity
}
