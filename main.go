package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	apiRoot string = "/gamesdb"
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
	Platform    string  `json:"platform"`
	ReleaseDate string  `json:"releaseDate"`
	AddedDate   string  `json:"addedDate"`
	Price       float64 `json:"price"`
	Cover       string  `json:"cover"`
}

// getGames responds with the list of all games as JSON
func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

// getGameByID responds with the data of the game which id matches the given id as JSON
func getGameByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	// query db for game with id
	println(id)
	c.IndentedJSON(http.StatusOK, games)
}

// postGames adds a game from JSON received in the request body
func postGames(c *gin.Context) {
	var newGame Game

	// call BindJSON to bind the received JSON to newGame
	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	// Add the new game to the slice
	games = append(games, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}

var games = []Game{
	{strconv.Itoa(rand.Int()), "Back 4 Blood", "Turtle Rock Studios", "Warner Bros. Games", "Back 4 Blood is a thrilling cooperative first-person shooter from the creators of the critically acclaimed Left 4 Dead franchise. Experience the intense 4 player co-op narrative campaign, competitive multiplayer as human or Ridden, and frenetic gameplay that keeps you in the action.", "Steam", "12 Oct, 2021", "28 Dec, 2021", 279.99, "https://cdn.akamai.steamstatic.com/steam/apps/924970/header.jpg?t=1639522452"},
	{strconv.Itoa(rand.Int()), "Ghostrunner", "One More Level, 3D Realms, Slipgate Ironworks™, All in! Games", "505 Games", "Ghostrunner offers a unique single-player experience: fast-paced, violent combat, and an original setting that blends science fiction with post-apocalyptic themes. It tells the story of a world that has already ended and its inhabitants who fight to survive.", "Steam", "27 Oct, 2020", "28 Dec, 2021", 99.99, "https://cdn.akamai.steamstatic.com/steam/apps/1139900/header.jpg?t=1635496307"},
	{strconv.Itoa(rand.Int()), "Hollow Knight", "Team Cherry", "Team Cherry", "Forge your own path in Hollow Knight! An epic action adventure through a vast ruined kingdom of insects and heroes. Explore twisting caverns, battle tainted creatures and befriend bizarre bugs, all in a classic, hand-drawn 2D style.", "Steam", "24 Feb, 2017", "28 Dec, 2021", 27.99, "https://cdn.akamai.steamstatic.com/steam/apps/367520/header.jpg?t=1625363925"},
	{strconv.Itoa(rand.Int()), "Death's Door", "Acid Nerve", "Devolver Digital", "Reaping souls of the dead and punching a clock might get monotonous but it's honest work for a Crow. The job gets lively when your assigned soul is stolen and you must track down a desperate thief to a realm untouched by death - where creatures grow far past their expiry.", "Steam", "20 Jul, 2021", "28 Dec, 2021", 49.95, "https://cdn.akamai.steamstatic.com/steam/apps/894020/header.jpg?t=1629235525"},
}

func main() {
	router := gin.Default()

	router.GET(apiRoot, getGames)
	router.GET(apiRoot+"/allgames", getGames)
	router.GET(apiRoot+"/:id", getGameByID)

	router.POST(apiRoot+"/addgame", postGames)

	router.Run("localhost:3000")
}
