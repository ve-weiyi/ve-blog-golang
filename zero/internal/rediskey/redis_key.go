package rediskey

func GetUserLogoutKey(uid string) string {
	return "user:logout:" + uid
}
