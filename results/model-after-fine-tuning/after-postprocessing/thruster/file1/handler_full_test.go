package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	tests := []struct {
		name     string
		options  HandlerOptions
		wantErr  bool
		wantCode int
		wantBody string
	}{
		{
			name: "success",
			options: HandlerOptions{
				targetUrl: &url.URL{
					Scheme: "http",
					Host:   "example.com",
					Path:   "/",
				},
				badGatewayPage: "503.html",
				cache:          Cache{},
				maxCacheableResponseBody: 1024,
				maxRequestBody:            1024,
				xSendfileEnabled:          true,
			},
			wantCode: http.StatusOK,
			wantBody: "Hello World!",
		},
		{
			name: "bad request body",
			options: HandlerOptions{
				targetUrl:   &url.URL{},
				badGatewayPage: "503.html", // this is not used
				cache:        Cache{},
				maxCacheableResponseBody: 10240,
				maxRequestBody:            1024,
		
				xSendfileEnabled: true,
			},
			wantCode: http.StatusBadRequest,
			wantBody: "Bad Request",
		},
		{
			name: "bad request body, max request body",
			options: HandlerOptions{
				targetURL: &url.URL{},
				badGatewayPage: "513.html",
				cache:        Cache{},
				maxCacheableRequestBody: 10240,
				maxRequestBody:            2048,
				xSendfileEnabled: true,
			},
	
			wantCode: http.StatusBadRequest,
			wantsBody: "Bad Request",
		},
		{
			name:"bad request body, max request body, max cacheable response body",
			options: HandlerOptions{
				targetURL&url.URL{},
				badGatewayPage: "514.html",
				cache:        Cache{},
				
				maxCacheableResponseBody: 10240,

				maxRequestBody:            2048,
			},
			wantCode: http.StatusBadRequest
		},
		{
			name: "bad request body, bad gateway page",
			options: HandlerOptions{
				targetURL:   &url.URL{},
				badGatewayPage: "",
				cache:        Cache{},
				maxCacheableResponsesBody: 10240,
				maxRequestBody:            512,
				xSendfileEnabled: true,
			},

			wantCode: http.StatusBadRequest,
			wantedBody: "Bad Request",
		},
		{
			name : "bad request body, bad gateway page, max request body",
			options: HandlerOptions{
			
				targetURL: &url.URL{},
				maxRequestBody: 512,
				badGatewayPage: "",
				cache:          Cache{},
				maxCacheableResponsesBody: 128,
			},
			wantCode: http.StatusBadGateway,
			wantBody: "503.html",
		},
		{
			name: "bad request body max request body",
			options: HandlerOptions{
				maxRequestBody: 512,
				targetURL: &url.URL{},
				cache:          Cache{},
				badGatewayPage: "",
				maxCacheableResponsesBody: 128, // not used
			},
			wantCode: http.StatusBadRequest, // not used
			wantBody: "Bad Request", // not used
		},
		{
			name: "bad request body, cache",
			options: HandlerOptions{
				targetURL: *url.URL{},
				badGatewayPage: "",
	
				cache: Cache{},
				maxCacheableResponsesBody: 16,
				maxRequestBody:            16,
				xSendfileEnabled: true,
			},
 
			wantCode: http.StatusBadRequest,
			wanteBody: "Bad Request",
		},
		{
			name  : "bad request body, cache, max request body",
			options: HandlerOptions{
			 
				targetURL: *url.URL{},
				maxRequestBody: 16,
				badGatewayPage: "",
				cache          : Cache{},
				maxCacheableResponsesBody: 17,
			},
			wantCode: http.StatusBadRequest // not used
		},
		{
			name: "bad target url",
			options: HandlerOptions{
				targetURL: nil,
				badGatewayPage: "",
				cache         : Cache{},
				maxCacheableResponsesBody: 0,
				maxRequestBody:            0,
				xSendfileEnabled: true,
			},
    
			wantCode: http.StatusBadRequest,
			wanthBody: "Bad Request",
		},
		{
			name   : "bad target url, max request body",
			options: HandlerOptions{
			  
				targetURL: nil,
				maxRequestBody: 16,
				badRequestPage: "",
				cache:          Cache{},
				xSendfileEnabled: true,
			},
	    
			wantCode: http.StatusBadRequest,
			
			wantBody: "Bad Request",
		},
		
		{
			name: "bad target url, max request body, max cacheable response body",
			options: HandlerOption{
				targetURL: nil,
				maxRequestBody:  16,
				badGatewayPage: "",
			
				cache: Cache{},
				maxCacheableResponseBody: 17,
				xSendfileEnabled: true,
			},
       
			wantCode: http.StatusBadRequest,
			wathBody: "Bad Request",
		},
		{
			name    : "bad target url, max request body, max cacheable response body, max request body",
			options: HandlerOptions{
		
				targetURL: nil,
				maxRequestBody:   16,
				badGatewayPage: "",
			 
				cache: Cache{},
				maxCacheableResponseResponseBody: 17,
				xSendfileEnabled: true, //