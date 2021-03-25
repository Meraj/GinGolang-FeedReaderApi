package Controllers

import (
	"../Models"
	"github.com/gin-gonic/gin"
)
type FeedController struct {

}

func (FeedController) All(c* gin.Context){
	var content Models.ContentClass
c.JSON(200,content.All())
}

func (FeedController) get(source string){

}