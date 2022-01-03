package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host     string = "api-gamesdb-db.mysql.database.azure.com"
	database string = "gamesdb"
	user     string = "sagnax@api-gamesdb-db"
	password string = "a8FWYrLQW3mjW6B"
)

type Game struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Developer   string  `json:"developer"`
	Publisher   string  `json:"publisher"`
	Overview    string  `json:"overview"`
	Genre       string  `json:"genre"`
	Platform    string  `json:"platform"`
	ReleaseDate string  `json:"releaseDate"`
	AddedDate   string  `json:"addedDate"`
	Price       float64 `json:"price"`
	Cover       string  `json:"cover"`
}

// getGames responds with the list of all games as JSON
func getGames(c *gin.Context) {
	var games = []Game{}
	var game = Game{}
	rows, err := db.Query("SELECT * FROM games;")
	checkError(err)
	defer rows.Close()
	//fmt.Println("Reading data:")
	for rows.Next() {
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Publisher, &game.Overview, &game.Genre, &game.Platform, &game.ReleaseDate, &game.AddedDate, &game.Price, &game.Cover)
		checkError(err)
		//fmt.Println(game)
		games = append(games, game)
	}

	c.IndentedJSON(http.StatusOK, games)
}

// getGameByTitle responds with the game matching the title given as JSON
func getGameByTitle(c *gin.Context) {
	var game = Game{}
	title := c.Param("title")

	rows, err := db.Query("SELECT * FROM games WHERE title = ?;", title)
	checkError(err)
	defer rows.Close()
	//fmt.Println("Reading data:")
	for rows.Next() {
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Publisher, &game.Overview, &game.Genre, &game.Platform, &game.ReleaseDate, &game.AddedDate, &game.Price, &game.Cover)
		checkError(err)
		//fmt.Println(game)
	}

	c.IndentedJSON(http.StatusOK, game)
}

// getGamesByTitle responds with a list of games that contains the title given as JSON
func getGamesByTitle(c *gin.Context) {
	var games = []Game{}
	var game = Game{}
	title := c.Param("title")

	query := "SELECT * FROM games WHERE title LIKE '%" + title + "%';"
	rows, err := db.Query(query)
	checkError(err)
	defer rows.Close()
	//fmt.Println("Reading data:")
	for rows.Next() {
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Publisher, &game.Overview, &game.Genre, &game.Platform, &game.ReleaseDate, &game.AddedDate, &game.Price, &game.Cover)
		checkError(err)
		//fmt.Println(game)
		games = append(games, game)
	}

	c.IndentedJSON(http.StatusOK, games)
}

// getGameByID responds with the data of the game which id matches the given id as JSON
func getGameByID(c *gin.Context) {
	var game = Game{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	//println(id)

	rows, err := db.Query("SELECT * FROM games WHERE id = ?;", id)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Publisher, &game.Overview, &game.Genre, &game.Platform, &game.ReleaseDate, &game.AddedDate, &game.Price, &game.Cover)
		checkError(err)
	}

	c.IndentedJSON(http.StatusOK, game)
}

// getGamesByPlatform responds with the list of all games matching given platform as JSON
func getGamesByPlatform(c *gin.Context) {
	var games = []Game{}
	var game = Game{}
	platform := c.Param("platform")
	if platform == "" {
		return
	}
	//println(platform)

	query := "SELECT * FROM games WHERE platform LIKE '%" + platform + "%';"

	rows, err := db.Query(query)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Publisher, &game.Overview, &game.Genre, &game.Platform, &game.ReleaseDate, &game.AddedDate, &game.Price, &game.Cover)
		checkError(err)
		games = append(games, game)
	}

	c.IndentedJSON(http.StatusOK, games)
}

// getGamesByGenre responds with the list of all games matching given genre as JSON
func getGamesByGenre(c *gin.Context) {
	var games = []Game{}
	var game = Game{}
	genre := c.Param("genre")
	if genre == "" {
		return
	}
	//println(genre)

	query := "SELECT * FROM games WHERE genre LIKE '%" + genre + "%';"

	rows, err := db.Query(query)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&game.ID, &game.Title, &game.Developer, &game.Publisher, &game.Overview, &game.Genre, &game.Platform, &game.ReleaseDate, &game.AddedDate, &game.Price, &game.Cover)
		checkError(err)
		games = append(games, game)
	}

	c.IndentedJSON(http.StatusOK, games)
}

// addGame adds a game from JSON received in the request body
func addGame(c *gin.Context) {
	newGame := Game{}

	// call BindJSON to bind the received JSON to newGame
	err := c.BindJSON(&newGame)
	checkError(err)

	// these are not allowed to be null
	if newGame.Title == "" || newGame.Developer == "" || newGame.AddedDate == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	//query := "INSERT INTO games (title, developer, publisher, overview, genre, platform, releaseDate, addedDate, price, cover) VALUES  ("Dishonored 2", "Arkane Studios", "Bethesda Softworks", "Reprise your role as a supernatural assassin in Dishonored 2. Declared a “masterpiece” by Eurogamer and hailed “a must-play revenge tale” by Game Informer, Dishonored 2 is the follow up to Arkane’s 1st-person action blockbuster & winner of 100+ 'Game of the Year' awards, Dishonored.", "Action, Adventure", "Steam", "11 Nov, 2016", "28 Dec, 2021", 89.99, "https://cdn.akamai.steamstatic.com/steam/apps/403640/header.jpg?t=1603889340")"
	res, err := db.Exec("INSERT INTO games (title, developer, publisher, overview, genre, platform, releaseDate, addedDate, price, cover) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", newGame.Title, newGame.Developer, newGame.Publisher, newGame.Overview, newGame.Genre, newGame.Platform, newGame.ReleaseDate, newGame.AddedDate, newGame.Price, newGame.Cover)
	checkError(err)
	rows, err := res.RowsAffected()
	checkError(err)
	fmt.Printf("Insert affected %d rows\n", rows)

	c.IndentedJSON(http.StatusCreated, newGame)
}

func updateGame(c *gin.Context) {
	updatedGame := Game{}
	id := c.Param("id")
	if id == "" {
		return
	}

	// call BindJSON to bind the received JSON to newGame
	err := c.BindJSON(&updatedGame)
	checkError(err)

	// these are not allowed to be null
	if updatedGame.Title == "" || updatedGame.Developer == "" || updatedGame.AddedDate == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	res, err := db.Exec("UPDATE games SET games.title=?, games.developer=?, games.publisher=?, overview=?, games.genre=?, games.platform=?, games.releaseDate=?, games.addedDate=?, games.price=?, games.cover=? WHERE id = ?;", updatedGame.Title, updatedGame.Developer, updatedGame.Publisher, updatedGame.Overview, updatedGame.Genre, updatedGame.Platform, updatedGame.ReleaseDate, updatedGame.AddedDate, updatedGame.Price, updatedGame.Cover, id)
	checkError(err)
	rows, err := res.RowsAffected()
	checkError(err)
	fmt.Printf("Update affected %d rows\n", rows)

	c.Status(http.StatusOK)

}

func deleteGame(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		return
	}

	res, err := db.Exec("DELETE FROM games WHERE id = ?;", id)
	checkError(err)
	rows, err := res.RowsAffected()
	checkError(err)
	fmt.Printf("Delete affected %d rows\n", rows)

	c.Status(http.StatusOK)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("error occured: " + err.Error())
		panic(err)
	}
}

var (
	db  *sql.DB
	err error
)

const apiRoot string = "/gamesdb/api"

func main() {

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true&tls=true", user, password, host, database)

	db, err = sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	err = db.Ping()
	checkError(err)
	fmt.Println("Ping database successfully.")

	router := gin.Default()

	router.GET(apiRoot, getGames)
	router.GET(apiRoot+"/", getGames)
	router.GET(apiRoot+"/id/:id", getGameByID)
	router.GET(apiRoot+"/title/:title", getGameByTitle)
	router.GET(apiRoot+"/search/:title", getGamesByTitle)
	router.GET(apiRoot+"/platform/:platform", getGamesByPlatform)
	router.GET(apiRoot+"/genre/:genre", getGamesByGenre)

	router.POST(apiRoot+"/add", addGame)
	router.GET(apiRoot+"/delete/:id", deleteGame)
	router.POST(apiRoot+"/update/:id", updateGame)

	router.Run("0.0.0.0:3000")
}
