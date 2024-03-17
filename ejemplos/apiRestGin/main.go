package main

// go get -u github.com/gin-gonic/gin
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "A Night At The Opera", Artist: "Queen", Price: 56.99},
	{ID: "2", Title: "Hotel California", Artist: "Eagles", Price: 40.0},
	{ID: "3", Title: "Rumours", Artist: "Fleetwood Mac", Price: 45.0},
	{ID: "4", Title: "Dark Side of the Moon", Artist: "Pink Floyd", Price: 55.0},
	{ID: "5", Title: "Led Zeppelin IV", Artist: "Led Zeppelin", Price: 50.0},
	{ID: "6", Title: "Back in Black", Artist: "AC/DC", Price: 48.0},
	{ID: "7", Title: "The Wall", Artist: "Pink Floyd", Price: 60.0},
	{ID: "8", Title: "Thriller", Artist: "Michael Jackson", Price: 42.0},
	{ID: "9", Title: "Born to Run", Artist: "Bruce Springsteen", Price: 38.0},
	{ID: "10", Title: "Sticky Fingers", Artist: "The Rolling Stones", Price: 52.0},
	{ID: "11", Title: "Born in the USA", Artist: "Bruce Springsteen", Price: 46.0},
	{ID: "12", Title: "Bat Out of Hell", Artist: "Meat Loaf", Price: 44.0},
	{ID: "13", Title: "Boston", Artist: "Boston", Price: 47.0},
	{ID: "14", Title: "Hotel California", Artist: "Eagles", Price: 40.0},
	{ID: "15", Title: "Highway to Hell", Artist: "AC/DC", Price: 43.0},
	{ID: "16", Title: "Parallel Lines", Artist: "Blondie", Price: 41.0},
	{ID: "17", Title: "Synchronicity", Artist: "The Police", Price: 49.0},
	{ID: "18", Title: "London Calling", Artist: "The Clash", Price: 53.0},
	{ID: "19", Title: "The River", Artist: "Bruce Springsteen", Price: 39.0},
	{ID: "20", Title: "Off the Wall", Artist: "Michael Jackson", Price: 37.0},
	{ID: "21", Title: "Crime of the Century", Artist: "Supertramp", Price: 51.0},
	{ID: "22", Title: "Electric Warrior", Artist: "T. Rex", Price: 54.0},
	{ID: "23", Title: "Let It Be", Artist: "The Beatles", Price: 58.0},
	{ID: "24", Title: "Who's Next", Artist: "The Who", Price: 57.0},
	{ID: "25", Title: "Some Girls", Artist: "The Rolling Stones", Price: 59.0},
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album
	
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsById(c *gin.Context){
	id := c.Param("id")
	for _, a  := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Album no encontrado"})
}


func main() {
	fmt.Println("Welcome to this example")
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	router.GET("/albums",getAlbums)
	router.GET("/albums/:id",getAlbumsById)
	router.POST("/albums",postAlbums)

	router.Run(":8800")
}
