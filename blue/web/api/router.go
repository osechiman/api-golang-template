package api

import (
	"net/http"
	"vspro/green/controllers"
	"vspro/green/gateways"
	"vspro/green/presenters"

	"github.com/gin-gonic/gin"
)

func Listen() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("questions", listQuestion)
		v1.GET("questions/:id", getQuestionByID)
		v1.POST("questions", postQuestion)
		v1.DELETE("questions/:id", deleteQuestionByID)
	}

	router.Run(":8080")
}

func listQuestion(c *gin.Context) {
	qr := gateways.NewInMemoryRepository()
	qp := presenters.NewQuestionPresenter()
	qc := controllers.NewQuestionController(qr)
	ql, err := qc.ListQuestion()
	if err != nil {
		c.JSON(http.StatusNotFound, qp.ConvertToHttpErrorResponse(http.StatusNotFound, err))
		return
	}

	res := qp.ConvertToHttpQuestionListResponse(ql)
	c.JSON(http.StatusOK, res)
	return
}

func getQuestionByID(c *gin.Context) {
	qid := c.Param("id")
	qr := gateways.NewInMemoryRepository()
	qp := presenters.NewQuestionPresenter()
	qc := controllers.NewQuestionController(qr)
	q, err := qc.GetQuestionByID(qid)
	if err != nil {
		c.JSON(http.StatusNotFound, qp.ConvertToHttpErrorResponse(http.StatusNotFound, err))
		return
	}

	res := qp.ConvertToHttpQuestionResponse(q)
	c.JSON(http.StatusOK, res)
	return
}

func postQuestion(c *gin.Context) {
	qr := gateways.NewInMemoryRepository()
	qp := presenters.NewQuestionPresenter()
	qc := controllers.NewQuestionController(qr)
	q, err := qc.AddQuestion(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, qp.ConvertToHttpErrorResponse(http.StatusInternalServerError, err))
		return
	}

	res := qp.ConvertToHttpQuestionResponse(q)
	c.JSON(http.StatusCreated, res)
	return
}

func deleteQuestionByID(c *gin.Context) {
	qid := c.Param("id")
	qr := gateways.NewInMemoryRepository()
	qp := presenters.NewQuestionPresenter()
	qc := controllers.NewQuestionController(qr)
	err := qc.DeleteQuestionByID(qid)
	if err != nil {
		c.JSON(http.StatusNotFound, qp.ConvertToHttpErrorResponse(http.StatusNotFound, err))
		return
	}

	res := qp.ConvertToHttpDeleteQuestionResponse(http.StatusOK, qid)
	c.JSON(http.StatusOK, res)
	return
}
