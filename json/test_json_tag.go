package json

type Person struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
}

type Human struct {
	Name   string
	Person interface{}
}
