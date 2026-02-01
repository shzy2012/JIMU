package example

import (
	"tolo/src/tools"
)

type Params struct {
	tools.Paging
}

type ReqDto struct {
	Field1 string `json:"filed1"`
	Field2 string `json:"filed2"`
	Field3 string `json:"filed3"`
}

type RespDto struct {
	Field1 string `json:"filed1"`
	Field2 string `json:"filed2"`
	Field3 string `json:"filed3"`
}
