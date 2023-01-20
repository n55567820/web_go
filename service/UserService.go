package service

import (
	"github.com/gin-gonic/gin"

	"golangAPI/pojo"
	"net/http"
	"strconv"
	"log"

)

var userList = []pojo.User{}

func FindAllUsers(c *gin.Context){
	c.JSON(http.StatusOK, userList)
}

func PostUser(c *gin.Context){
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	userList = append(userList, user)
	c.JSON(http.StatusOK, "Successfully posted")
}

func DeleteUser(c *gin.Context){
	userId, _:= strconv.Atoi(c.Param("id"))
	for _, user := range userList {
		log.Println(user)
		if user.Id == userId {
			userList = append(userList[:userId], userList[userId+1:]...)
			c.JSON(http.StatusOK, "Successfully deleted")
			return
		}
	}
	c.JSON(http.StatusNotFound, "Error")
}

func PutUser(c *gin.Context){
	berforeUser := pojo.User{}
	err := c.BindJSON(&berforeUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
	}
	userId, _ := strconv.Atoi(c.Param("id"))
	for key, user:= range userList {
		if userId ==user.Id {
			userList[key] = berforeUser
			log.Println(userList[key])
			c.JSON(http.StatusOK, "Successfully")
			return
		}
	}
	c.JSON(http.StatusNotFound, "Error")
}