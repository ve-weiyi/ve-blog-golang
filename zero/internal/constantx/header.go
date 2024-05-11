package constantx

const (
	HeaderXRequestID = "x-request-id"
	HeaderXAuthToken = "x-auth-token"
	HeaderXUserID    = "x-user-id"

	HeaderTimezone = "timezone"
	HeaderCountry  = "country"
	HeaderLanguage = "language"
)

var HeaderFields = []string{
	HeaderXRequestID,
	HeaderXAuthToken,
	HeaderXUserID,
	HeaderTimezone,
	HeaderCountry,
	HeaderLanguage,
}
