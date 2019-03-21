package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type FeeCalculate struct {
	AccountNumber string `json:"account_number,required"`
	Amount        string `json:"amount,required"`
}

func main() {
	route := gin.Default()
	route.GET("/api/v1/deposit/limit/transaction", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"limit": "30000.00",
		})
	})

	route.POST("/api/v1/deposit/fee", func(context *gin.Context) {
		var feeCalculate FeeCalculate
		err := context.ShouldBindJSON(&feeCalculate)
		if err != nil {
			context.JSON(400, gin.H{
				"error_code":    "E12002",
				"error_message": "wrong input validation",
			})
			return
		}
		amount, err := strconv.ParseFloat(feeCalculate.Amount, 64)
		if err != nil {
			context.JSON(400, gin.H{
				"error_code":    "E12002",
				"error_message": "wrong input validation",
			})
			return
		}

		if strings.HasPrefix(feeCalculate.AccountNumber, "12") {
			context.JSON(200, gin.H{
				"account_number": feeCalculate.AccountNumber,
				"amount":         feeCalculate.Amount,
				"fee":            "15.00",
				"net_amount":     fmt.Sprintf("%.2f", amount-15),
			})
			return
		}
		context.JSON(200, gin.H{
			"account_number": feeCalculate.AccountNumber,
			"amount":         feeCalculate.Amount,
			"fee":            "0.00",
			"net_amount":     fmt.Sprintf("%.2f", amount),
		})

	})
	route.Run(":3000")
}
