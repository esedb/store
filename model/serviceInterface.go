package model

import (
	"estore/common"
)

type Service interface {
	GetCollection(context *common.Context, collection string)
}
