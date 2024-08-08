package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type GetPrResponse struct {
	Status string `json:"status"`
}

type PostPrRequest struct {
	Name string `json:"name"`
}

type PostPrResponse struct {
	Message string `json:"message"`
}

type PrHandler struct{}

func NewPrHandler(e *echo.Echo) {
	handler := &PrHandler{}
	e.GET("/getPr", handler.getPr)
	e.POST("/postPr", handler.postPr)
}

func (h *PrHandler) getPr(e echo.Context) error {
	return e.JSON(http.StatusOK, GetPrResponse{Status: "OKE BNGT"})
}

func (h *PrHandler) postPr(e echo.Context) error {
	var reqBody = PostPrRequest{}

	err := e.Bind(&reqBody)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	log.Println(StructToJson(reqBody))
	message := fmt.Sprintf("HELLO %s!", reqBody.Name)
	log.Println(message)
	return e.JSON(http.StatusOK, PostPrResponse{Message: message})
}

func StructToJson(val interface{}) string {
	b, err := json.Marshal(val)
	if err != nil {
		return err.Error()
	}

	return string(b)
}
