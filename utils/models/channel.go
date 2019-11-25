package models

type Channel struct {
	ID            uint64   `json:"id"`
	Name          string   `json:"name"`
	TotalMSGCount int64    `json:"-"`
	Members       []uint64 `json:"members"`
	Admins        []uint64 `json:"admins"`
	WorkspaceID   uint64   `json:"-"`
	CreatorID     uint64   `json:"-"`
}
