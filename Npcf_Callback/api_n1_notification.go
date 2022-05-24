package Npcf_Callback

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"net/http"
	"net/url"

	"github.com/free5gc/openapi"
	omodels "github.com/free5gc/openapi/models"
	"github.com/softmurata/freeopenapi/models"
)

var (
	_ context.Context
)

type N1NotificationApiService service

func (a *N1NotificationApiService) NotifyN1Notification(ctx context.Context, supi string, n1NotificationRequest models.N1NotificationRequest) (models.N1NotificationResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  models.N1NotificationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/n1-notification/{supi}"
	localVarPath = strings.Replace(localVarPath, "{"+"supi"+"}", fmt.Sprintf("%v", supi), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if n1NotificationRequest.BinaryDataN1SmMessage != nil {
		localVarHeaderParams["Content-Type"] = "multipart/related"
		localVarPostBody = &n1NotificationRequest
	} else {
		localVarHeaderParams["Content-Type"] = "application/json"
		localVarPostBody = n1NotificationRequest.JsonData
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "multipart/related", "application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := openapi.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHttpResponse.Status,
	}

	switch localVarHttpResponse.StatusCode {
	case 204:
		return localVarReturnValue, localVarHttpResponse, nil
	case 400:
		var v omodels.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 401:
		var v omodels.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 403:
		var v omodels.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 404:
		var v omodels.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 411:
		var v omodels.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 413:
		var v omodels.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 415:
		var v omodels.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 429:
		var v omodels.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 500:
		var v omodels.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	case 503:
		var v omodels.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHttpResponse, apiError
	default:
		return localVarReturnValue, nil, nil

	}

}
