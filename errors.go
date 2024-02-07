package goauthultimate

type customErrors int

const (
	SUCCESS        customErrors = 200
	ERROR          customErrors = 500
	INVALID_PARAMS customErrors = 400

	ERROR_EXIST_TAG       customErrors = 10001
	ERROR_EXIST_TAG_FAIL  customErrors = 10002
	ERROR_NOT_EXIST_TAG   customErrors = 10003
	ERROR_GET_TAGS_FAIL   customErrors = 10004
	ERROR_COUNT_TAG_FAIL  customErrors = 10005
	ERROR_ADD_TAG_FAIL    customErrors = 10006
	ERROR_EDIT_TAG_FAIL   customErrors = 10007
	ERROR_DELETE_TAG_FAIL customErrors = 10008
	ERROR_EXPORT_TAG_FAIL customErrors = 10009
	ERROR_IMPORT_TAG_FAIL customErrors = 10010

	ERROR_NOT_EXIST_ARTICLE        customErrors = 10011
	ERROR_CHECK_EXIST_ARTICLE_FAIL customErrors = 10012
	ERROR_ADD_ARTICLE_FAIL         customErrors = 10013
	ERROR_DELETE_ARTICLE_FAIL      customErrors = 10014
	ERROR_EDIT_ARTICLE_FAIL        customErrors = 10015
	ERROR_COUNT_ARTICLE_FAIL       customErrors = 10016
	ERROR_GET_ARTICLES_FAIL        customErrors = 10017
	ERROR_GET_ARTICLE_FAIL         customErrors = 10018
	ERROR_GEN_ARTICLE_POSTER_FAIL  customErrors = 10019

	ERROR_AUTH_CHECK_TOKEN_FAIL    customErrors = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT customErrors = 20002
	ERROR_AUTH_TOKEN               customErrors = 20003
	ERROR_AUTH                     customErrors = 20004
	ERROR_EXIST_USER               customErrors = 20005
	ERROR_SENDING_MAIL_FAILED      customErrors = 20006
	ERROR_AUTH_CHECK_CODE_FAIL     customErrors = 20007
	ERROR_USER_DOESNT_EXIST        customErrors = 20008
	ERROR_USER_OPERATIONS_FAILED   customErrors = 20009
	ERROR_CHAT_DOESNT_EXIST        customErrors = 20010
	ERROR_CODE_GENERATE_FAILED     customErrors = 20011
	ERROR_USER_NOT_VERIFIED        customErrors = 20012

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    customErrors = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   customErrors = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT customErrors = 30003

	ERROR_FETCHING_DATA customErrors = 40001
)

var MsgFlags = map[customErrors]string{
	SUCCESS: "Success",
	ERROR:   "General error",

	INVALID_PARAMS: "Invalid parameters",

	ERROR_EXIST_TAG:       "Tag already exists",
	ERROR_EXIST_TAG_FAIL:  "Error checking tag existance",
	ERROR_NOT_EXIST_TAG:   "Tag does not exist",
	ERROR_GET_TAGS_FAIL:   "Error getting tags",
	ERROR_COUNT_TAG_FAIL:  "Error counting tags",
	ERROR_ADD_TAG_FAIL:    "Error adding tag",
	ERROR_EDIT_TAG_FAIL:   "Error editing tag",
	ERROR_DELETE_TAG_FAIL: "Error deleting tag",
	ERROR_EXPORT_TAG_FAIL: "Error exporting tags",
	ERROR_IMPORT_TAG_FAIL: "Error importing tags",

	ERROR_NOT_EXIST_ARTICLE:        "Article does not exist",
	ERROR_CHECK_EXIST_ARTICLE_FAIL: "Error checking article existance",
	ERROR_ADD_ARTICLE_FAIL:         "Error adding article",
	ERROR_DELETE_ARTICLE_FAIL:      "Error deleting article",
	ERROR_EDIT_ARTICLE_FAIL:        "Error editing article",
	ERROR_COUNT_ARTICLE_FAIL:       "Error counting articles",
	ERROR_GET_ARTICLES_FAIL:        "Error getting articles",
	ERROR_GET_ARTICLE_FAIL:         "Error getting article",
	ERROR_GEN_ARTICLE_POSTER_FAIL:  "Error generating article poster",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Error checking auth token",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Auth token expired",
	ERROR_AUTH_TOKEN:               "Invalid auth token",
	ERROR_AUTH:                     "General auth error",
	ERROR_EXIST_USER:               "User already exists",
	ERROR_SENDING_MAIL_FAILED:      "Error sending email",
	ERROR_AUTH_CHECK_CODE_FAIL:     "Error checking auth code",
	ERROR_USER_DOESNT_EXIST:        "User does not exist",
	ERROR_USER_OPERATIONS_FAILED:   "Error during user operations",
	ERROR_CHAT_DOESNT_EXIST:        "Chat does not exist",
	ERROR_CODE_GENERATE_FAILED:     "Error generating auth code",
	ERROR_USER_NOT_VERIFIED:        "User not verified",

	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "Error saving uploaded image",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "Error checking uploaded image",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "Invalid image format",

	ERROR_FETCHING_DATA: "Error fetching data",
}

type ErrorResponseMessage struct {
	Status customErrors `json:"status"`
	Error  string       `json:"error"`
}
type ResponseMessage struct {
	Status  customErrors `json:"status"`
	Message string       `json:"message"`
}

func (c customErrors) GetMsg(code customErrors) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
