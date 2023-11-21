package response

import (
	"time"
)

type UserDTO struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Status    int       `json:"status"`
	Avatar    string    `json:"avatar"`
	Intro     string    `json:"intro"`
	Website   string    `json:"website"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`

	IpAddress string     `json:"ip_address"` // ip host
	IpSource  string     `json:"ip_source"`  // ip 源
	Roles     []*RoleDTO `json:"roles"`
}

type UserAreaDTO struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
