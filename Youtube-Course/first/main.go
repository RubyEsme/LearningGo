//and this is going to be a little lets say microservice that will get albums
//first endpoint with GOOOO

package main

import "fmt"

func main(){
	//this is how you declare a structure
	type album struct {
		ID string `json:"id"`
		Title string `json:"title"`
		Artist string `json:"artist"`
		Price float64 `json:"price"`
	}

	//data I will use

	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "Jhon Coltrane", Price: 56.99 },
	}

	//getAlbums function
	func getAlbums(c *gin.Context)  {
		c.IndentedJSON(http.StatusOK, albums)
		
	}

}