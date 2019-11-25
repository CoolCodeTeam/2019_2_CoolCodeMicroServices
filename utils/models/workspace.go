package models

type Workspace struct {
	ID        uint64     `json:"id"`
	Name      string     `json:"name"`
	Channels  []*Channel `json:"-"`
	Members   []uint64   `json:"members"`
	Admins    []uint64   `json:"admins"`
	CreatorID uint64     `json:"-"`
}
