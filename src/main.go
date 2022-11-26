package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Sample Data 4", Artist: "Sample User", Price: 60.00},
}

var env_v1 = os.Getenv("env_variables")
var env_v2 = os.Getenv("GIN_DEMO_VALUE1")
var env_v3 = os.Getenv("GIN_DEMO_VALUE2")

var env_v4 = os.Getenv("env_param")
var env_v5 = os.Getenv("test_ps_value1")

var albums1 = []album{
	{ID: "1", Title: "Blue 11111", Artist: "John 11111", Price: 56.99},
	{ID: "env_variables1", Title: env_v1, Artist: env_v2, Price: 0.0},
	{ID: "env_variables2", Title: env_v1, Artist: env_v3, Price: 0.0},
}

var albums2 = []album{
	{ID: "2", Title: "Jeru 2222", Artist: "Gerry 22222", Price: 17.99},
	{ID: "env_param", Title: env_v4, Artist: env_v5, Price: 0.0},
}

func main() {
	router := gin.Default()
	router.GET("/", getAlbums)
	router.GET("/test", getAlbums1)
	router.GET("/test2", getAlbums2)
	router.GET("/api", getResponse)
	router.GET("/api2", getResponse)
	router.GET("/api3", getResponse)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run(":80")
}

//get api response
func getResponse(c *gin.Context) {
	response, err := http.Get("https://holidays-jp.github.io/api/v1/date.json")

	if err != nil {
		c.IndentedJSON(http.StatusOK, "error1")
	}

	var j interface{}
	err = json.NewDecoder(response.Body).Decode(&j)
	if err != nil {
		c.IndentedJSON(http.StatusOK, "error2")
	}
	c.IndentedJSON(http.StatusOK, j)
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func getAlbums1(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums1)
}
func getAlbums2(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums2)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
