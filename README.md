# A Game Database API
---

A RESTful API for a Game Database Written with **Go (Golang)**, **Gin Web Framework** and using **Microsoft Azure DB** as the database.

### Endpoints:

- "/gamesdb/api" GET - returns a list with all the games
- "/gamesdb/api/id/:id" GET - returns a list with the game matching the given id
- "/gamesdb/api/title/:title" GET - returns a list with the game matching the given title
- "/gamesdb/api/search/:title" GET - returns a list with all the games containing the given title
- "/gamesdb/api/platform/:platform" GET - returns a list with all the games from the given platform
- "/gamesdb/api/genre/:genre" GET - returns a list with all the games matching the given genre

- "/gamesdb/api/add" POST - Adds a game to the database<br>
Example data:

```json
{
  "title": "Dishonored 2",
  "developer": "Arkane Studios",
  "publisher": "Bethesda Softworks",
  "overview": "Reprise your role as a supernatural assassin in Dishonored 2. Declared a “masterpiece” by Eurogamer and hailed “a must-play revenge tale” by Game Informer, Dishonored 2 is the follow up to Arkane’s 1st-person action blockbuster & winner of 100+ 'Game of the Year' awards, Dishonored.",
  "genre": "Action, Adventure",
  "platform": "Steam",
  "releaseDate": "11 Nov, 2016",
  "addedDate": "28 Dec, 2021",
  "price": 89.99,
  "cover": "https://cdn.akamai.steamstatic.com/steam/apps/403640/header.jpg?t=1603889340"
}
```

- "/gamesdb/api/delete/:id" GET - deletes the game with the given id from the database

- "/gamesdb/api/update/:id" POST - update the game with the given id in the database<br>
Example data:

```json
{
  "title": "Dishonored 2",
  "developer": "Arkane Studios",
  "publisher": "Bethesda Softworks",
  "overview": "Reprise your role as a supernatural assassin in Dishonored 2. Declared a “masterpiece” by Eurogamer and hailed “a must-play revenge tale” by Game Informer, Dishonored 2 is the follow up to Arkane’s 1st-person action blockbuster & winner of 100+ 'Game of the Year' awards, Dishonored.",
  "genre": "Action, Adventure",
  "platform": "Steam",
  "releaseDate": "11 Nov, 2016",
  "addedDate": "28 Dec, 2021",
  "price": 89.99,
  "cover": "https://cdn.akamai.steamstatic.com/steam/apps/403640/header.jpg?t=1603889340"
}
```
