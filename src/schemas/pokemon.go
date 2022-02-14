package schemas

//Pokemon -> represents a json pokemon
type Pokemon struct {
	Order   int    `json:"order"`
	Name    string `json:"name"`
	Sprites Sprite `json:"sprites"`
}

//Sprite -> represent a sprite of the pokemon
type Sprite struct {
	Default string `json:"front_default"`
}
