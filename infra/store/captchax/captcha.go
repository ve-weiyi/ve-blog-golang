package captchax

import (
	"context"
	"crypto/subtle"
	"time"

	"github.com/mojocn/base64Captcha"

	"github.com/ve-weiyi/blog-cloud/infra/store/storex"
)

// 默认过期时间
const defaultExpiration = 15 * time.Minute

// Store 图形验证码仓库
type Store struct {
	DefaultHeight  int     // 默认高度 80
	DefaultWidth   int     // 默认宽度 240
	DefaultLength  int     // 默认长度,位数 6
	DefaultMaxSkew float64 // 默认倾斜因子 0.7
	DefaultDotRate float64 // 默认干扰点比率 20%

	keyPrefix string
	kv        storex.KVStore      // 校验路径直接访问，以便传入调用方 ctx
	store     base64Captcha.Store // 生成路径经适配器，供 base64Captcha 内部回调

	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

// New 创建图形验证码仓库。
// keyPrefix 用于隔离不同服务的验证码命名空间；expiration <= 0 时使用默认 15 分钟。
func New(store storex.KVStore, keyPrefix string, expiration time.Duration) *Store {
	if expiration <= 0 {
		expiration = defaultExpiration
	}

	return &Store{
		DefaultHeight:  80,
		DefaultWidth:   240,
		DefaultLength:  6,
		DefaultMaxSkew: 0.7,
		DefaultDotRate: 0.20,
		keyPrefix:      keyPrefix,
		kv:             store,
		store:          newDriverStore(store, keyPrefix, expiration),
		DriverAudio:    base64Captcha.DefaultDriverAudio,
		DriverString:   base64Captcha.NewDriverString(80, 240, 0, 0, 5, "1234567890abcdefghijklmnopqrstuvwxyz", nil, nil, nil),
		DriverChinese:  base64Captcha.NewDriverChinese(80, 240, 0, 0, 5, "1234567890abcdefghijklmnopqrstuvwxyz", nil, nil, nil),
		DriverMath:     base64Captcha.NewDriverMath(80, 240, 0, 0, nil, nil, nil),
		DriverDigit:    base64Captcha.NewDriverDigit(40, 80, 6, 0.7, 10),
	}
}

// GetMathImageCaptcha 生成算术图形验证码，返回 captchaKey、base64 图片
const (
	// API 层允许的请求尺寸范围；内部另有防御性 clamp。
	MinRequestDimension = 80
	MaxRequestDimension = 400

	minDimension = 40
	maxDimension = 400
)

// normalizeDimension 在入口处收敛尺寸，防止超大值触发巨量内存分配。
func normalizeDimension(value, def int) int {
	if value <= 0 {
		return def
	}
	if value < minDimension {
		return minDimension
	}
	if value > maxDimension {
		return maxDimension
	}
	return value
}

func (s *Store) GetMathImageCaptcha(height int, width int) (string, string, string, error) {
	height = normalizeDimension(height, s.DefaultHeight)
	width = normalizeDimension(width, s.DefaultWidth)
	driver := base64Captcha.NewDriverMath(height, width, 0, 0, nil, nil, nil)

	c := base64Captcha.NewCaptcha(driver, s.store)
	return c.Generate()
}

// GetImageCaptcha 按类型生成图形验证码，返回 captchaKey、base64 图片
func (s *Store) GetImageCaptcha(CaptchaType string, height int, width int, length int) (string, string, error) {
	var driver base64Captcha.Driver

	height = normalizeDimension(height, s.DefaultHeight)
	width = normalizeDimension(width, s.DefaultWidth)
	if length == 0 {
		length = s.DefaultLength
	}

	var dotCount = int(float64(height)*s.DefaultDotRate + float64(width)*s.DefaultDotRate)

	//create base64 encoding captcha
	switch CaptchaType {
	case "audio":
		driver = s.DriverAudio
	case "string":
		driver = s.DriverString.ConvertFonts()
	case "math":
		driver = s.DriverMath.ConvertFonts()
	case "chinese":
		driver = s.DriverChinese.ConvertFonts()
	case "digit":
		driver = base64Captcha.NewDriverDigit(height, width, length, s.DefaultMaxSkew, dotCount)
	default:
		driver = base64Captcha.NewDriverDigit(height, width, length, s.DefaultMaxSkew, dotCount)
	}

	c := base64Captcha.NewCaptcha(driver, s.store)
	id, b64s, _, err := c.Generate()

	return id, b64s, err
}

// VerifyCaptcha 校验图形验证码。
// 无论答案对错都立即删除，验证码一次性使用，防止同一 key 被反复试错。
// 未经由 base64Captcha.Store，以便传入调用方 ctx。
func (s *Store) VerifyCaptcha(ctx context.Context, id string, answer string) (bool, error) {
	if id == "" || answer == "" {
		return false, nil
	}

	key := s.keyPrefix + id

	stored, found, err := s.kv.Get(ctx, key)
	if err != nil {
		return false, err
	}
	if !found {
		return false, nil
	}

	if err := s.kv.Delete(ctx, key); err != nil {
		return false, err
	}

	// constant-time 比较，避免时序侧信道
	return subtle.ConstantTimeCompare([]byte(stored), []byte(answer)) == 1, nil
}
