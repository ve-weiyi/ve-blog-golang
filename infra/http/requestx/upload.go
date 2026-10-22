package requestx

const (
	MinUploadTokenTTL     = 60
	MaxUploadTokenTTL     = 3600
	DefaultUploadTokenTTL = 3600
)

// NormalizeUploadTokenTTL 把凭证有效期收敛到允许区间。
func NormalizeUploadTokenTTL(seconds int) int {
	if seconds <= 0 {
		return DefaultUploadTokenTTL
	}
	if seconds < MinUploadTokenTTL {
		return MinUploadTokenTTL
	}
	if seconds > MaxUploadTokenTTL {
		return MaxUploadTokenTTL
	}
	return seconds
}
