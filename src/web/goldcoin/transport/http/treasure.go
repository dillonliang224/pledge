package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTreasureInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func getTreasureUserRecords(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func postTreasureCodes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func getTreasureParticipation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func getTreasureWinners(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func getTreasureMyParticipation(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func getTreasureDetailInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func getTreasureRecords(c *gin.Context) {
	activityId := c.Query("activityId")
	productId := c.Query("productId")
	period := c.Query("period")
	records, err := svc.GetTreasureRecords(c, activityId, productId, period)
	if err != nil {
		_ = c.Error(err)
		return
	}

	// var d struct {
	// 	nickname  string
	// 	avatar    string
	// 	period    string
	// 	codeCount int32
	// 	created   *time.Time
	// }

	// get user info

	c.JSON(http.StatusOK, gin.H{
		"records": records,
	})
}

func getTreasurePop(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func getTreasureClientNotification(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func getTreasureStatistics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
