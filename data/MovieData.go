package MovieData

import (
	MovieType "github.com/Idoarizs/crud-api-mux/model"
)

var (
	MovieList = []MovieType.Movie{
		{
			ID:    "1",
			Isbn:  "28471",
			Title: "The Last Heroes",
			Director: MovieType.Director{
				Firstname: "Jonathan",
				Lastname:  "Alexander",
			},
		},
	}
)
