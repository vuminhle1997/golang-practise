package hello

type Gender string

const (
	Male      = "Male"
	Female    = "Female"
	Undefined = "Undefined"
)

type Person struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Age      int    `json:"age"`
	Gender   Gender `json:"gender"`
}
