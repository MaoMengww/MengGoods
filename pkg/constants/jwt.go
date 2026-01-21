package constants

const (
	TypeAccess = 1 + iota
	TypeRefresh
	TypeLogin

	AuthHeader         = "Authorization"
	AccessTokenHeader  = "Access-Token"
	RefreshTokenHeader = "Refresh-Token"
)
