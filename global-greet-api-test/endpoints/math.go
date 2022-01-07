package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetMath responds with about/help page
func GetMath(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"found": "Welcome, you found this hidden easter egg endpoint! Congrats, smarty pants ;)",

		"about": "This endpoint receives two integer values, and returns the addition of such integers.",

		"help": "PUT /math/:x/:y -> responds with x+y",
	})
}

// PutMath receives two integers and responds with their addition
func PutMath(c *gin.Context) {
	x := c.Param("x")
	y := c.Param("y")

	intX, err := strconv.Atoi(x)
	intY, err1 := strconv.Atoi(y)

	if err != nil || err1 != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "user gave non-integer values that cannot be added"})
	} else {
		result := intX + intY
		c.IndentedJSON(http.StatusCreated, gin.H{"addition": result})
	}
}
