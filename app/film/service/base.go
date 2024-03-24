package service

type BaseFilm struct {
	Name []string `json:"name"`
	Data []string `json:"data"`
}

type BaseFilmRate struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func handle(mapFilmInfo map[string]interface{}) {

}

var (
	rata = "_rata"
)
