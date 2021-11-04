package recommendation_system_movie_store

type Movie struct {
	Id           int64   `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Photo        string  `json:"photo,omitempty"`
	Description  string  `json:"description,omitempty"`
	Genre        string  `json:"genre,omitempty"`
	Year         string  `json:"year,omitempty"`
	CountEpisode string  `json:"count_episode,omitempty"`
	Score        float64 `json:"score,omitempty"`
}

type MovieUpdate struct {
	Id           int64    `json:"id,omitempty"`
	Name         *string  `json:"name,omitempty"`
	Photo        *string  `json:"photo,omitempty"`
	Description  *string  `json:"description,omitempty"`
	Genre        *string  `json:"genre,omitempty"`
	Year         *string  `json:"year,omitempty"`
	CountEpisode *string  `json:"count_episode,omitempty"`
	Score        *float64 `json:"score,omitempty"`
}

type MovieStore interface {
	List(count int64) ([]Movie, error)
	Create(movie *Movie) (*Movie, error)
	GetById(id int64) (*Movie, error)
	Update(movie *MovieUpdate) (*Movie, error)
	Delete(id int64) error
	GetByName(name string) (*Movie, error)
}
