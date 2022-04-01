package contract

type XiaoeRequestLoggerInterface interface {
	HttpRequestLog(record *XiaoeHttpRequestRecord)
}
