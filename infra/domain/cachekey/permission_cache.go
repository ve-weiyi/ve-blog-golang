package cachekey

import "fmt"

// Permission模块缓存Key定义

const (
	PolicyInvalidateChannel   = "blog:permissionx:invalidate:policy"
	UserRoleInvalidateChannel = "blog:permissionx:invalidate:user"
	userRoleCachePrefix       = "blog:permissionx:user_roles:"
	userRoleGenPrefix         = "blog:permissionx:user_roles_gen:"
)

func UserRoleCacheKey(userId string) string {
	return fmt.Sprintf("%s%s", userRoleCachePrefix, userId)
}

// UserRoleGenKey 是用户角色缓存的代际号，撤权时递增，阻止在途加载把旧角色写回。
func UserRoleGenKey(userId string) string {
	return fmt.Sprintf("%s%s", userRoleGenPrefix, userId)
}
