package dto

import (
	"time"
)

type UserInfoReq struct {
	Nickname string `json:"nickname" example:"nickname"` // 昵称
	Website  string `json:"website" example:"website"`   // 网站
	Intro    string `json:"intro" example:"intro"`       // 简介
	Avatar   string `json:"avatar" example:"avatar"`     // 头像
}

type UserDTO struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Intro    string `json:"intro"`
	Website  string `json:"website"`
	Email    string `json:"email"`

	Status       int64  `json:"status"`
	RegisterType string `json:"register_type"`
	IpAddress    string `json:"ip_address"` // ip host
	IpSource     string `json:"ip_source"`  // ip 源

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Roles     []*RoleDTO `json:"roles"`
}

type UserAreaDTO struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type UserMenuDTO struct {
	Id        int64         `json:"id"`
	Name      string        `json:"name"`
	Path      string        `json:"path"`
	Component string        `json:"component"`
	Icon      string        `json:"icon"`
	Rank      int64         `json:"rank"`
	IsHidden  int64         `json:"is_hidden"`
	Children  []UserMenuDTO `json:"children"`
}
