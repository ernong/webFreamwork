package errmsg

// Error ...
type Error interface {
	Code() int
	Error() string
}

type ErrInfo struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
}

// ErrorInfo ...
func (err ErrInfo) Code() int {
	return err.code
}

// ErrorInfo ...
func (err ErrInfo) Error() string {
	return err.msg
}

// NewErrInfo ...
func NewErrInfo(code int, msg string) Error {
	return ErrInfo{
		code: code,
		msg:  msg,
	}
}

// NewInternalErr ...
func NewInternalErr(err error) Error {
	return ErrInfo{
		code: 3000,
		msg:  err.Error(),
	}
}

// interface error msg
var (
	NoErr                      = ErrInfo{code: 0, msg: ""}
	ErrInvalidLoginInfo        = ErrInfo{code: 1000, msg: "Invalid login info"}
	ErrInvalidTokenInfo        = ErrInfo{code: 1001, msg: "Token info did not match request user"}
	ErrorInvalidCommentID      = ErrInfo{code: 1002, msg: "Invalid CommentID"}
	ErrorInvalidFriendID       = ErrInfo{code: 1003, msg: "Invalid FriendID"}
	ErrorInvalidUID            = ErrInfo{code: 1004, msg: "Invalid UID info"}
	ErrorInvalidSession        = ErrInfo{code: 1005, msg: "Session Invalid"}
	ErrInvalidRequest          = ErrInfo{code: 1006, msg: "Invalid Request"}
	ErrTokenInvalidRequest     = ErrInfo{code: 1007, msg: "Token invalid"}
	ErrLossParamUID            = ErrInfo{code: 1008, msg: "Loss parameter [platform_user_id]"}
	ErrLossParamQueryFMUID     = ErrInfo{code: 1009, msg: "Loss parameter [platform_user_id] when query friend cycle by id"}
	ErrLossContent             = ErrInfo{code: 1010, msg: "Loss parameter [content] when deploy friend cycle"}
	ErrLossCommentID           = ErrInfo{code: 1011, msg: "Loss parameter [comment_id] when cancel friend comment"}
	ErrLossCommentParameter    = ErrInfo{code: 1012, msg: "Loss parameter [platform_user_id] or [friend_cycle_id] when deploy friend comment"}
	ErrorInvalidRequestSession = ErrInfo{code: 1013, msg: "invalid parameter for request session"}
	ErrorLossCoverID           = ErrInfo{code: 1014, msg: "Loss parameter [cover_id] or [uid] when praise cover"}

	ErrInvalidComment        = ErrInfo{code: 2000, msg: "Invalid Comment SensitiveWord"}
	ErrInvalidCommentForbid  = ErrInfo{code: 2001, msg: "Comment Forbidden"}
	ErrInvalidCommentTooMany = ErrInfo{code: 2002, msg: "Comment frequently, Please wait a little while"}
	ErrInvalidPraise         = ErrInfo{code: 2003, msg: "Already Praised"}

	ErrInvalidDBOperator      = ErrInfo{code: 3000, msg: "DB operator err"}
	ErrInvalidQueryComment    = ErrInfo{code: 3001, msg: "Query Comment Info Error"}
	ErrCancelComment          = ErrInfo{code: 3002, msg: "Cancel Comment Info Error"}
	ErrCommitPraiseComment    = ErrInfo{code: 3003, msg: "Deploy Praise Comment Info Error"}
	ErrCommitCover            = ErrInfo{code: 3004, msg: "Deploy Front Cover Info Error"}
	ErrCommentNoFm            = ErrInfo{code: 3005, msg: "the friend ship message not exit"}
	ErrCommentCancelFm        = ErrInfo{code: 3006, msg: "the friend ship message already canceled"}
	ErrCommentNoFmComment     = ErrInfo{code: 3007, msg: "the friend ship message comment not exit"}
	ErrCommentCancelFmComment = ErrInfo{code: 3008, msg: "the friend ship message comment already canceled"}
	ErrCoverCancel            = ErrInfo{code: 3009, msg: "the frontCover already canceled"}

	ErrRedisSaveCoverPraise = ErrInfo{code: 4001, msg: "Save praise cover in redis err"}

	ErrLoginInvalidPhone        = ErrInfo{code: 1001, msg: "Can't find use with phone"}
	ErrAccountLogin             = &ErrInfo{code: 20101, msg: "Account or password error."}
	ErrLoginInvalidPasswd       = ErrInfo{code: 1002, msg: "Invalid Password"}
	ErrLoginInvalidUserStatus   = ErrInfo{code: 1003, msg: "Invalid user status"}
	ErrLoginInvalidUserRole     = ErrInfo{code: 1004, msg: "Invalid user role"}
	ErrLoginInvalidMiner        = ErrInfo{code: 1011, msg: "Invalid minerID or minerName"}
	ErrLoginInvalidMinerName    = ErrInfo{code: 1012, msg: "Miner name is already in use"}
	ErrLoginInvalidName         = ErrInfo{code: 1013, msg: "Miner name nil"}
	ErrInvalidVerificationCode  = ErrInfo{code: 1004, msg: "Invalid verification codes"}
	ErrInvalidTotalWithdrawCode = ErrInfo{code: 1004, msg: "Too many withdraw this week"}
	ErrTooManyRequests          = ErrInfo{code: 1005, msg: "Too many requests"}
	ErrRequireLogin             = ErrInfo{code: 1006, msg: "Require login"}
	ErrRegisterPhoneUsed        = ErrInfo{code: 1010, msg: "Phone number is already in use"}
	ErrRegisterPhone            = ErrInfo{code: 1011, msg: "Invalid phone number"}
	ErrRoleInvalidMiner         = ErrInfo{code: 1012, msg: "No role for operating"}
	ErrCreateInvalidMiner       = ErrInfo{code: 1013, msg: "Can't create miner user"}

	ErrInvalidInviteCode      = ErrInfo{code: 2001, msg: "Invalid InviteCode"}
	ErrUserNotExist           = ErrInfo{code: 2002, msg: "User not exist"}
	ErrInternalServer         = ErrInfo{code: 3000, msg: "Internal Server Error"}
	ErrNotImplement           = ErrInfo{code: 10001, msg: "Not Implement"}
	ErrInvalidMinerID         = ErrInfo{code: 20015, msg: "Invalid minerID"}
	ErrNullParentMinerNode    = ErrInfo{code: 20002, msg: "nil Parent MinerNode"}
	ErrInvalidAmount          = ErrInfo{code: 20003, msg: "Invalid Amount"}
	ErrInvalidRewardLevel     = ErrInfo{code: 20004, msg: "Invalid Reward Level"}
	ErrInvalidRewardType      = ErrInfo{code: 20005, msg: "Invalid Reward Type"}
	ErrInvalidStringAmount    = ErrInfo{code: 20006, msg: "Invalid String Amount"}
	ErrInvalidDescendantLevel = ErrInfo{code: 20007, msg: "Invalid Descendant Level"}
	ErrInvalidRewardInfo      = ErrInfo{code: 20008, msg: "Invalid Reward Info"}
	ErrMinerExists            = ErrInfo{code: 20009, msg: "Miner already exists"}
	ErrWithdrawFrequencyLimit = ErrInfo{code: 20010, msg: "Withdraw day frequency limit"}
	ErrInvalidWithdrawAddr    = ErrInfo{code: 20011, msg: "Invalid Withdraw Address"}
	ErrInvalidWithdrawTrx     = ErrInfo{code: 20014, msg: "Invalid Withdraw transaction"}
	ErrInvalidUser            = ErrInfo{code: 20012, msg: "Invalid User"}
	ErrInvalidUserStatus      = ErrInfo{code: 20013, msg: "Invalid User Status"}
)
