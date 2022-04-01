package library

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/contract"
	"talkcheap.xiaoeknow.com/xiaoetong/eframe/helpers/network"
	"time"
)

type HttpClient struct {
	*http.Client
}

func NewHttpClient(config *HttpClientConfig) (httpClient *HttpClient) {
	httpClient = &HttpClient{
		Client: &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives:   config.DialKeepAliveSecond < 0,
				MaxIdleConns:        config.MaxIdleConnections,
				MaxIdleConnsPerHost: config.MaxIdleConnectionsPerHost,
				IdleConnTimeout:     time.Duration(config.IdleConnTimeoutSecond) * time.Second,
				Proxy:               http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(config.DialTimeoutSecond) * time.Second,
					KeepAlive: time.Duration(config.DialKeepAliveSecond) * time.Second,
				}).DialContext,
			},
		},
	}
	return
}

func (c *HttpClient) Get(ctx context.Context, url string, header map[string]string, logger contract.XiaoeRequestLoggerInterface) (response []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	var clientResp *http.Response
	var paramsString string
	var beginTime = time.Now()
	defer recordLog(req, &clientResp, &paramsString, &response, &err, beginTime, logger)
	if err != nil {
		return nil, err
	}

	addXeHeader(ctx, req)
	addCustomHeader(req, header)

	clientResp, err = c.Do(req)

	if err != nil {
		err = fmt.Errorf("response is nil: %w", err)
		return nil, err
	}

	if clientResp == nil {
		err = fmt.Errorf("response is nil: %w", err)
		return nil, err
	}
	defer clientResp.Body.Close()

	if clientResp.StatusCode != http.StatusOK {
		err = &contract.HttpResponseError{
			Code: clientResp.StatusCode,
			Msg:  fmt.Sprintf("response error, code %d", clientResp.StatusCode),
		}
		return nil, err
	}

	response, err = getBytesFromHttpResponse(clientResp)

	return response, err
}

func (c *HttpClient) Post(ctx context.Context, url string, params []byte, header map[string]string, logger contract.XiaoeRequestLoggerInterface) (response []byte, err error) {
	req, err := http.NewRequest("POST", url, nil)
	var clientResp *http.Response
	var paramsString string
	var beginTime = time.Now()
	defer recordLog(req, &clientResp, &paramsString, &response, &err, beginTime, logger)
	if err != nil {
		return nil, err
	}

	if params != nil {
		req.Body = ioutil.NopCloser(bytes.NewReader(params))
		paramsString = string(params)
	}

	addXeHeader(ctx, req)
	addCustomHeader(req, header)

	clientResp, err = c.Do(req)

	if err != nil {
		err = fmt.Errorf("response is nil: %w", err)
		return nil, err
	}

	if clientResp == nil {
		err = fmt.Errorf("response is nil: %w", err)
		return nil, err
	}
	defer clientResp.Body.Close()

	if clientResp.StatusCode != http.StatusOK {
		err = &contract.HttpResponseError{
			Code: clientResp.StatusCode,
			Msg:  fmt.Sprintf("response error, code %d", clientResp.StatusCode),
		}
		return nil, err
	}

	response, err = getBytesFromHttpResponse(clientResp)

	return response, err
}

func (c *HttpClient) GetJson(ctx context.Context, url string, response interface{}, logger contract.XiaoeRequestLoggerInterface) error {
	req, err := http.NewRequest("GET", url, nil)
	var responseBytes []byte
	var clientResp *http.Response
	var params string
	var beginTime = time.Now()
	defer recordLog(req, &clientResp, &params, &responseBytes, &err, beginTime, logger)
	if err != nil {
		return err
	}

	addJsonHeader(req)
	addXeHeader(ctx, req)

	clientResp, err = c.Do(req)

	if err != nil {
		err = fmt.Errorf("response is nil: %w", err)
		return err
	}

	if clientResp == nil {
		err = fmt.Errorf("response is nil: %w", err)
		return err
	}
	defer clientResp.Body.Close()

	if clientResp.StatusCode != http.StatusOK {
		err = &contract.HttpResponseError{
			Code: clientResp.StatusCode,
			Msg:  fmt.Sprintf("response error, code %d", clientResp.StatusCode),
		}
		return err
	}

	responseBytes, err = getBytesFromHttpResponse(clientResp)
	if err != nil {
		return err
	}

	if response == nil {
		response = &(map[string]interface{}{})
	}

	if err = json.Unmarshal(responseBytes, response); err != nil {
		return err
	}

	return err
}

func (c *HttpClient) PostJson(ctx context.Context, url string, params interface{}, response interface{}, logger contract.XiaoeRequestLoggerInterface) error {
	req, err := http.NewRequest("POST", url, nil)
	var responseBytes []byte
	var clientResp *http.Response
	var paramsString string
	var beginTime = time.Now()
	defer recordLog(req, &clientResp, &paramsString, &responseBytes, &err, beginTime, logger)
	if err != nil {
		return err
	}

	if params != nil {
		paramsBytes, e := json.Marshal(params)
		if e != nil {
			err = e
			return err
		}

		if paramsBytes != nil {
			req.Body = ioutil.NopCloser(bytes.NewReader(paramsBytes))
			paramsString = string(paramsBytes)
		}

	}

	addJsonHeader(req)
	addXeHeader(ctx, req)

	clientResp, err = c.Do(req)

	if err != nil {
		err = fmt.Errorf("response is nil: %w", err)
		return err
	}

	if clientResp == nil {
		err = fmt.Errorf("response is nil: %w", err)
		return err
	}
	defer clientResp.Body.Close()

	if clientResp.StatusCode != http.StatusOK {
		err = &contract.HttpResponseError{
			Code: clientResp.StatusCode,
			Msg:  fmt.Sprintf("response error, code %d", clientResp.StatusCode),
		}
		return err
	}

	responseBytes, err = getBytesFromHttpResponse(clientResp)
	if err != nil {
		return err
	}

	if response == nil {
		response = &(map[string]interface{}{})
	}

	if err = json.Unmarshal(responseBytes, response); err != nil {
		return err
	}

	return err

}

func (c *HttpClient) Close() error {
	c.CloseIdleConnections()
	return nil
}

func addCustomHeader(req *http.Request, header map[string]string) {
	for k, v := range header {
		req.Header.Add(k, v)
	}
}

func getBytesFromHttpResponse(response *http.Response) (b []byte, err error) {
	if response == nil {
		return nil, errors.New("http response is nil")
	}
	if response.ContentLength > 0 {
		b = make([]byte, response.ContentLength)
		readBytes := 0
		for {
			n, e := response.Body.Read(b[readBytes:])
			readBytes = readBytes + n
			if e != nil {
				break
			}
		}
	} else {
		b, err = io.ReadAll(response.Body)
	}
	return
}

func addJsonHeader(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

func addXeHeader(ctx context.Context, req *http.Request) {
	values, ok := ctx.Value(contract.XeCtx).(map[string]string)
	if ok {
		//Xe灰度标识
		if header := values[contract.XeTagHeader]; header != "" {
			req.Header.Add(contract.XeTagHeader, header)
		}

		//SkyWalking标识
		if header := values[contract.Sw8Header]; header != "" {
			req.Header.Add(contract.Sw8Header, header)
		}

		if header := values[contract.Sw8CorrelationHeader]; header != "" {
			req.Header.Add(contract.Sw8CorrelationHeader, header)
		}
	}

}

func recordLog(req *http.Request, resp **http.Response, params *string, response *[]byte, err *error, beginTime time.Time, logger contract.XiaoeRequestLoggerInterface) {
	if logger != nil && req != nil {
		record := contract.XiaoeHttpRequestRecord{}
		record.Sw8 = req.Header.Get(contract.Sw8Header)
		record.Sw8Correlation = req.Header.Get(contract.Sw8CorrelationHeader)
		record.XeTag = req.Header.Get(contract.XeTagHeader)
		record.TargetUrl = req.URL.String()
		record.Method = req.Method
		if params != nil {
			record.Params = *params
		}
		if err != nil && *err != nil {
			record.Msg = (*err).Error()
		}

		record.ClientIp = network.GetInternalIp()
		if resp != nil && *resp != nil {
			record.HttpStatus = (*resp).StatusCode
			record.ServerIp = (*resp).Request.RemoteAddr
			record.UserAgent = req.UserAgent()
		}
		if response != nil && *response != nil {
			record.Response = string(*response)
		}
		record.BeginTime = beginTime.Format("2006-01-02 15:04:05.000")
		end := time.Now()
		spend := end.UnixNano() - beginTime.UnixNano()
		record.EndTime = end.Format("2006-01-02 15:04:05.000")
		record.CostTime = int(spend / 1000000)
		logger.HttpRequestLog(&record)
	}
}
