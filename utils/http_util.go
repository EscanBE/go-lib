package utils

// IsHttpStatusCodeIndicatesInformational returns true if Http Status Code is 1xx
func IsHttpStatusCodeIndicatesInformational(httpStatusCode int) bool {
	return httpStatusCode >= 100 && httpStatusCode <= 199
}

// IsHttpStatusCodeIndicatesSuccess returns true if Http Status Code is 2xx
func IsHttpStatusCodeIndicatesSuccess(httpStatusCode int) bool {
	return httpStatusCode >= 200 && httpStatusCode <= 299
}

// IsHttpStatusCodeIndicatesRedirection returns true if Http Status Code is 3xx
func IsHttpStatusCodeIndicatesRedirection(httpStatusCode int) bool {
	return httpStatusCode >= 300 && httpStatusCode <= 399
}

// IsHttpStatusCodeIndicatesClientError returns true if Http Status Code is 4xx
func IsHttpStatusCodeIndicatesClientError(httpStatusCode int) bool {
	return httpStatusCode >= 400 && httpStatusCode <= 499
}

// IsHttpStatusCodeIndicatesServerError returns true if Http Status Code is 5xx
func IsHttpStatusCodeIndicatesServerError(httpStatusCode int) bool {
	return httpStatusCode >= 500 && httpStatusCode <= 599
}
