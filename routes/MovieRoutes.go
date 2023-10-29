package MovieRoutes

import (
	// Custom
	MovieHandler "github.com/Idoarizs/crud-api-mux/controller"
	MovieType "github.com/Idoarizs/crud-api-mux/model"
)

var MovieRoutes = []MovieType.Route{
	{Type: "GET", Path: "/movies", Handler: MovieHandler.GetMovies, Methods: []string{"GET"}},
	{Type: "GET", Path: "/movies/{id}", Handler: MovieHandler.GetMovie, Methods: []string{"GET"}},
	{Type: "POST", Path: "/movies", Handler: MovieHandler.CreateMovie, Methods: []string{"POST"}},
	{Type: "PUT", Path: "/movies/{id}", Handler: MovieHandler.UpdateMovie, Methods: []string{"PUT"}},
	{Type: "DELETE", Path: "/movies/{id}", Handler: MovieHandler.DeleteMovie, Methods: []string{"DELETE"}},
}
