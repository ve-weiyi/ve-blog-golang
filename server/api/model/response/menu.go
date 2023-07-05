package response

type UserMenu struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	Component string     `json:"component"`
	Icon      string     `json:"icon"`
	Rank      int        `json:"rank"`
	IsHidden  bool       `json:"is_hidden"`
	Children  []UserMenu `json:"children"`
}
