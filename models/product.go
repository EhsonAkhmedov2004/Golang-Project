package Models

type Product struct {
	Name string `json:"name"`
	Price int `json:"price"`
	Range []int `json:"range"`
	Percent int `json:"percent"`
}
type Db struct {
	Db []Product `json:"db"` 
}