package pkg

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/chalermridp/go-template-rest/internal/app/constant"
	"github.com/gin-gonic/gin"
)

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseKey constant.ResponseStatus) {
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}

func PanicHandler(ctx *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		switch key {
		case
			constant.DataNotFound.GetResponseStatus():
			ctx.JSON(http.StatusBadRequest, BuildResponse_(key, msg, Null()))
			ctx.Abort()
		case
			constant.Unauthorized.GetResponseStatus():
			ctx.JSON(http.StatusUnauthorized, BuildResponse_(key, msg, Null()))
			ctx.Abort()
		default:
			ctx.JSON(http.StatusInternalServerError, BuildResponse_(key, msg, Null()))
			ctx.Abort()
		}
	}
}
