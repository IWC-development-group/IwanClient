package iwanCore

var (
	BaseURLS = []string{
		"http://localhost:8080",
	}
)

type Configurator struct {
	URLS []string `json:"URLS"`
}

type IwanResponse struct {
	Status    string `json: "status"`
	Name      string `json: "name"`
	Namespace string `json: "namespace"`
	Content   string `json: "content"`
}
