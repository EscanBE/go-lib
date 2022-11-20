package types

//goland:noinspection GoSnakeCaseUsage,GoUnusedConst
const (
	// log level

	// LOG_LEVEL_DEBUG is constant for debug level of logger
	LOG_LEVEL_DEBUG = "debug"

	// LOG_LEVEL_INFO is constant for info level of logger
	LOG_LEVEL_INFO = "info"

	// LOG_LEVEL_ERROR is constant for error level of logger
	LOG_LEVEL_ERROR = "error"

	// LOG_LEVEL_DEFAULT is constant for default log level of logger (info)
	LOG_LEVEL_DEFAULT = LOG_LEVEL_INFO

	// log format

	// LOG_FORMAT_TEXT is constant for output log with text format of logger
	LOG_FORMAT_TEXT = "text"

	// LOG_FORMAT_JSON is constant for output log with json format of logger
	LOG_FORMAT_JSON = "json"

	// LOG_FORMAT_DEFAULT is constant for default output log format of logger (json)
	LOG_FORMAT_DEFAULT = LOG_FORMAT_JSON
)
