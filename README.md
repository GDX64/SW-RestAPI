# SW Golang API

## dependencies

You are going to need the mongodb driver "go.mongodb.org/mongo-driver/mongo" and "github.com/gorilla/mux"

## connecting a database

By default the api is trying to connect to a mongodb database running on **local:27017**. You just need tu run the linux binary main.bin to put it to work.
To compile this by yourself, you just need to use `go run *.go`. All the important constants are at `main.go`.

## using the api
Planets have the following fields:
- ID
- Nome
- Clima
- Terreno
- Filmes

### Adding a new planet
To add a new planet to the database, you will use the url /planet and method POST. You will need to put in the body of the request a json with data like this:
```
{
    "nome": "Tatooine",
    "clima": "Montanhoso",
    "terreno": "Árido",
}

```

Adjacent fields won't be used. The server will calculate how many movies this planet is on using this API: https://swapi.dev/about.

### Quering a planet
A list of all registered planets can be retrieved by a GET at `/planets`
A planet can be queried by name or id using this url: `/planet/?nome=[name]&id=[id]`, you can use one or the other, or both.

### Deleting a planet

You can delete a planet only using the id with this url `/planet/?id=[id]` and the DELETE method.


