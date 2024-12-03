package constant

// General
const (
	MsgSuccess      = "operation completed successfully"
	MsgFailed       = "operation failed"
	MsgNotFound     = "resource not found"
	MsgUnauthorized = "unauthorized access"
	MsgForbidden    = "you do not have permission to perform this action"
	MsgBadRequest   = "invalid request parameters"
	MsgServerError  = "internal server error please try again later"
	MsgConflict     = "conflict detected resource already exists"
	MsgInvalidInput = "invalid input provided"
)

// Api
const (
	MsgAPIInvalidRequest       = "invalid request parameters please check and try again"
	MsgAPIResourceNotFound     = "requested resource not found"
	MsgAPIForbidden            = "you do not have permission to access this resource"
	MsgAPIUnauthorized         = "authorization required please log in to continue"
	MsgAPIConflict             = "conflict detected the resource already exists"
	MsgAPIBadRequest           = "bad request please verify the request data"
	MsgAPIRequestTooLarge      = "request entity too large please reduce the data size and try again"
	MsgAPIRateLimitExceeded    = "rate limit exceeded please wait and try again"
	MsgAPIUnsupportedMediaType = "unsupported media type please check the content type and try again"
)

// Attendance
const (
	MsgAttendanceGetCacheFailure = "failed to get attendances cache"
	MsgAttendanceSetCacheFailure = "failed to set attendances cache"
)

// Candidate
const (
	CandidatesCache = "candidates"

	MsgCandidateGetCacheFailure    = "failed to get candidates cache"
	MsgCandidateSetCacheFailure    = "failed to set candidates cache"
	MsgCandidateDeleteCacheFailure = "failed to delete candidates cache"
	MsgCandidateCannotUpdate       = "cannot update a candidate who has completed the process"

	MsgCandidateStatusApplied       = "Applied"
	MsgCandidateStatusInterviewing  = "Interviewing"
	MsgCandidateStatusRejected      = "Rejected"
	MsgCandidateStatusOnHold        = "On Hold"
	MsgCandidateStatusOfferSent     = "Offer Sent"
	MsgCandidateStatusOfferAccepted = "Offer Accepted"
	MsgCandidateStatusOfferDeclined = "Offer Declined"
	MsgCandidateStatusOnboarding    = "Onboarding"
)

// User
const (
	MsgUserRegistrationSuccess  = "user account has been successfully created"
	MsgUserRegistrationFailure  = "failed to create user account please check the input and try again"
	MsgUserLoginSuccess         = "login successful welcome back"
	MsgUserLoginFailure         = "invalid username or password please try again"
	MsgUserAccountLocked        = "your account has been locked due to multiple failed login attempts contact support for assistance"
	MsgUserPasswordResetSuccess = "password has been reset successfully please login with your new password"
	MsgUserPasswordResetFailure = "failed to reset password please try again"
	MsgUserLogoutSuccess        = "logout successful"
	MsgUserUnauthorizedAccess   = "unauthorized access please log in to continue"
	MsgUserForbiddenAccess      = "you do not have permission to access this resource"
	MsgUserNotFound             = "user not found please check the credentials and try again"
	MsgUserProfileUpdateSuccess = "profile updated successfully"
	MsgUserProfileUpdateFailure = "failed to update profile please try again"
	MsgUserAccountDeletion      = "user account has been deleted successfully"
)

// Product
const (
	MsgProductAdded    = "product has been added successfully"
	MsgProductUpdated  = "product has been updated"
	MsgProductDeleted  = "product has been deleted"
	MsgProductNotFound = "product not found"
)

// Order
const (
	MsgOrderPlaced    = "your order has been placed successfully"
	MsgOrderCancelled = "your order has been cancelled"
	MsgOrderNotFound  = "order not found"
)

// Action media type
const (
	MsgActionUploadSuccess     = "file uploaded successfully"
	MsgActionUploadFailure     = "file upload failed please try again"
	MsgActionDownloadSuccess   = "file downloaded successfully"
	MsgActionDownloadFailure   = "file download failed please try again"
	MsgActionSaveSuccess       = "changes saved successfully"
	MsgActionSaveFailure       = "failed to save changes please try again"
	MsgActionSendSuccess       = "request sent successfully"
	MsgActionSendFailure       = "failed to send request please try again"
	MsgActionProcessInProgress = "your request is being processed please wait"
	MsgActionProcessCompleted  = "processing completed successfully"
)

// Data
const (
	MsgDataCreationSuccess = "data has been created successfully"
	MsgDataCreationFailure = "failed to create data please try again"
	MsgDataUpdateSuccess   = "data has been updated successfully"
	MsgDataUpdateFailure   = "failed to update data please try again"
	MsgDataDeletionSuccess = "data has been deleted successfully"
	MsgDataDeletionFailure = "failed to delete data please try again"
	MsgDataNotFound        = "requested data not found"
	MsgDataDuplicateEntry  = "duplicate entry the data already exists"
	MsgDataInvalidFormat   = "invalid data format please verify and try again"
	MsgDataProcessingError = "error processing data please contact support"
)

// Department
const (
	MsgDepartmentParentNotFound = "parent department not found"
	MsgDepartmentNotFound       = "department not found"
)

// Leave Request
const (
	MsgLeaveRequestApprove = "approved"
	MsgLeaveRequestReject  = "rejected"
	MsgLeaveRequestPending = "pending"
)
