package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func CreatePDF(ctx *gin.Context) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	ctx.JSON(http.StatusCreated, nil)
}
