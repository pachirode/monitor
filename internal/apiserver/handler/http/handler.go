package http

import (
	"github.com/pachirode/monitor/internal/apiserver/biz"
	"github.com/pachirode/monitor/internal/apiserver/pkg/validation"
)

type Handler struct {
	biz biz.IBiz
	val *validation.Validator
}

func NewHandler(biz biz.IBiz, val *validation.Validator) *Handler {
	return &Handler{
		biz: biz,
		val: val,
	}
}
