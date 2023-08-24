package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/strernd/tipo/basic/controllers/offer"
)

func RunServer(oc *offer.OfferController) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/offer", func(c *gin.Context) {
		var offer offer.BaseOffer

		if err := c.ShouldBindJSON(&offer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		res, err := oc.CreateOffer(c, offer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, res)
	})

	r.GET("/offer/:email", func(c *gin.Context) {
		email := c.Param("email")
		res, err := oc.GetOffer(c, email)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	})

	r.Run()
}
