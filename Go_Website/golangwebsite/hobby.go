package main
import "hithub.com/gin-gonic/gin"

function main()
{
	r : gin.Default()

	r.GET("/hobby", func(c *gin.Context)
{
	c.String(200, "Hobby Page.")

	r.Run(":3008")
})

}