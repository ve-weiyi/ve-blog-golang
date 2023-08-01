package response

type UserMenu struct {
	Id        int
	Name      string
	Path      string
	Component string
	Icon      string
	Rank      int
	IsHidden  bool
	Children  []UserMenu
}
