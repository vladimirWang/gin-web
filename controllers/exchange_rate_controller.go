package controllers

import (
	"exchange_app/global"
	"exchange_app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateExchangeRate(ctx *gin.Context) {
	var exchangeRate models.ExchangeRate
	if err := ctx.ShouldBindJSON(&exchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exchangeRate.Date = time.Now()
	if err := global.Db.AutoMigrate(exchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.Create(&exchangeRate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, exchangeRate)
}
