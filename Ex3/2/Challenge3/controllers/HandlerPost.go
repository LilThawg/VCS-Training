package controllers

import (

	"github.com/challenge3/database"
	"github.com/challenge3/models"

	"net/http"

	"github.com/gin-gonic/gin"
)
// AddPost godoc
// @Summary 	Add a post
// @Description	Thêm 1 bài post
// @Tags		Handler Post 
// @Param 		Token header string true "Insert your access token" default(<Add access token here>)
// @Accept		json
// @Produce 	json	
// @Param 		post body models.PostAdd true "Add Post"
// @Success		200 {object} models.Post
// @Failure     400 {object} models.ExampleFailure
// @Router      /api/addPost [post]
func AddPost(c *gin.Context) {
	email := c.MustGet("email").(string)
	var newPost models.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 
	newPost.OwnerPost = email

	message := validatePost(newPost)
	
	if message != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	
	c.JSON(http.StatusOK, newPost)
	connection.Create(&newPost)
}

// GetAllPosts godoc
// @Summary 	Get All Posts
// @Description	Lấy thông tin toàn bộ Post
// @Tags		Handler Post 
// @Param 		Token header string true "Insert your access token" default(<Add access token here>)
// @Param 		page query string false "page" 
// @Accept		json
// @Produce 	json	
// @Success		200 {array} models.Post
// @Failure     400 {object} models.ExampleFailure
// @Router      /api/getAllPosts [get]
func GetAllPosts(c *gin.Context) {
	pagination := GeneratePaginationFromRequest(c)
	var post models.Post
	postLists, totalRows, err := GetPaginationPost(&post, &pagination)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data":      postLists,
		"totalList": totalRows,
	})
}

// DeletePost godoc
// @Summary 	Delete A Post
// @Description	Xoá 1 Post theo id
// @Tags		Handler Post
// @Param 		Token header string true "Insert your access token" default(<Add access token here>)
// @Accept		json
// @Produce 	json	
// @Param 		id query string true "Post's ID" 
// @Success		200 {object} models.ExampleSuccess
// @Failure     400 {object} models.ExampleFailure
// @Router      /api/deletePost [delete]
func DeletePost(c *gin.Context) {
	role := c.MustGet("role").(string)
	email := c.MustGet("email").(string)
	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)
	id := c.Query("id")
	var post models.Post
	
	if result := connection.First(&post, id); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	
	if role == "user" {
		if post.OwnerPost == email {
			connection.Unscoped().Where("id = ?", id).Delete(&post)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bạn không phải chủ post"})
			return
		}
	}

	if role == "admin" {
		connection.Unscoped().Where("id = ?", id).Delete(&post)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted!",
	})
	
}


// UpdatePost godoc
// @Summary 	Update a post 
// @Description	Chỉnh thông tin post theo id
// @Tags		Handler Post
// @Param 		Token header string true "Insert your access token" default(<Add access token here>)
// @Accept		json
// @Produce 	json	
// @Param 		id query string true "Post's ID" 
// @Param 		post body models.PostAdd true "Update Post"
// @Success		200 {object} models.ExampleSuccess
// @Failure     400 {object} models.ExampleFailure
// @Router      /api/updatePost [put]
func UpdatePost(c *gin.Context) {
	var updatedPost models.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 
	
	message := validatePost(updatedPost)
	
	if message != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": message})
		return
	}

	role := c.MustGet("role").(string)
	email := c.MustGet("email").(string)

	connection := database.GetDatabase()
	defer database.CloseDatabase(connection)

	var post models.Post
	id := c.Query("id")

	if result := connection.First(&post, id); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}
	
	if role == "user" {
		if post.OwnerPost == email {
			connection.Model(&post).Where("id = ?", id).Update(&updatedPost)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bạn không phải chủ post"})
			return
		}
	}

	if role == "admin" {
		connection.Model(&post).Where("id = ?", id).Update(&updatedPost)
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "Updated!",
	})
	
}