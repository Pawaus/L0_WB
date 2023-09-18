package handlers

import (
	"github.com/gin-gonic/gin"
)

type Usecase interface {
	GetOrderByID(uid string) (string, error)
}

type Handlers struct {
	usecase Usecase
}

type Uid_get struct {
	uid string `uri:"uid" binding:"required"`
}

func NewHandler(u Usecase) *Handlers {
	return &Handlers{usecase: u}
}

func (h *Handlers) GetUid(ctx *gin.Context) {
	var uid Uid_get
	err := ctx.ShouldBindUri(&uid)
	if err != nil {
		ctx.JSON(400, gin.H{"msg": err.Error()})
		return
	} else {
		uu, res := ctx.Params.Get("uid")
		if res {
			val, err_get_order := h.usecase.GetOrderByID(uu)
			if err_get_order != nil {
				ctx.JSON(400, gin.H{"msg": err_get_order.Error()})
				return
			} else {
				if val == "" {
					ctx.JSON(404, gin.H{"msg": "No data"})
					return
				}
				ctx.JSON(200, val)
				return
			}

		} else {
			ctx.JSON(400, gin.H{"msg": "Not found uid"})
			return
		}
	}
}
