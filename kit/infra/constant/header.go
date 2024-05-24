package constant

const (
	HeaderXRequestId = "x-request-id"
	HeaderXAuthToken = "x-auth-token"
	HeaderXUserId    = "x-user-id"

	HeaderTimezone = "timezone"
	HeaderCountry  = "country"
	HeaderLanguage = "language"
)

var HeaderFields = []string{
	HeaderXRequestId,
	HeaderXAuthToken,
	HeaderXUserId,
	HeaderTimezone,
	HeaderCountry,
	HeaderLanguage,
}
