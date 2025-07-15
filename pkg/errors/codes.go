package errors

var (
	// General
	ErrInvalidPayload  = New(701, "Invalid payload", 400)
	ErrInternalError   = New(799, "Internal server error", 500)
	ErrUnauthorized    = New(703, "Unauthorized", 401)
	ErrForbidden       = New(704, "Forbidden", 403)
	ErrNotFound        = New(705, "Resource not found", 404)
	ErrConflict        = New(706, "Conflict detected", 409)
	ErrTooManyRequests = New(707, "Too many requests", 429)

	// Auth
	ErrInvalidCredentials = New(710, "Invalid username or password", 401)
	ErrTokenExpired       = New(711, "Token has expired", 401)
	ErrTokenInvalid       = New(712, "Invalid token", 401)

	// User
	ErrUserExists   = New(720, "User already exists", 409)
	ErrUserNotFound = New(721, "User not found", 404)
	ErrUserDisabled = New(722, "User account is disabled", 403)

	// Validation
	ErrMissingFields = New(730, "Required fields are missing", 400)
	ErrInvalidEmail  = New(731, "Invalid email format", 400)
	ErrWeakPassword  = New(732, "Password does not meet requirements", 400)
	ErrHashFailed    = New(733, "Failed to Hash Password", 400)

	// Database
	ErrDBConnectionFailed = New(740, "Database connection failed", 500)
	ErrDBQueryFailed      = New(741, "Database query failed", 500)
	ErrDBTransactionFail  = New(742, "Database transaction failed", 500)

	// File & Uploads
	ErrFileTooLarge     = New(750, "File too large", 413)
	ErrFileUnsupported  = New(751, "Unsupported file type", 415)
	ErrFileUploadFailed = New(752, "File upload failed", 500)

	// Payment
	ErrPaymentFailed = New(760, "Payment failed", 402)
	ErrCardDeclined  = New(761, "Card was declined", 402)
)
