package internal

import (
	"net/url"
	"net/http"
	"testing"

)

func TestNewHandler(t *testing.T) {
	_ = []struct {
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
				cache:          NewMemoryCache(1*MB, 1*MB),
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
				cache:        NewMemoryCache(1*MB, 1*MB),
				maxCacheableResponseBody: 10240,
				maxRequestBody:            1024,
		
				xSendfileEnabled: true,
			},
			wantCode: http.StatusBadRequest,
			wantBody: "Bad Request",
		},
		
		{
			name:"bad request body, max request body, max cacheable response body",
			options: HandlerOptions{
				targetUrl: &url.URL{},
				badGatewayPage: "514.html",
				cache:        NewMemoryCache(1*MB, 1*MB),
				
				maxCacheableResponseBody: 10240,

				maxRequestBody:            2048,
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "bad request body, bad gateway page",
			options: HandlerOptions{
				targetUrl:   &url.URL{},
				badGatewayPage: "",
				cache:        NewMemoryCache(1*MB, 1*MB),
				maxCacheableResponseBody: 10240,
				maxRequestBody:            512,
				xSendfileEnabled: true,
			},

			wantCode: http.StatusBadRequest,
			wantBody: "Bad Request",
		},
		{
			name : "bad request body, bad gateway page, max request body",
			options: HandlerOptions{
				targetUrl: &url.URL{},
				maxRequestBody: 512,
				badGatewayPage: "",
				cache:          NewMemoryCache(1*MB, 1*MB),
				maxCacheableResponseBody: 128,
			},
			wantCode: http.StatusBadGateway,
			wantBody: "503.html",
		},
		{
			name: "bad request body max request body",
			options: HandlerOptions{
				maxRequestBody: 512,
				targetUrl: &url.URL{},
				cache:          NewMemoryCache(1*MB, 1*MB),
				badGatewayPage: "",
				maxCacheableResponseBody: 128, // not used
			},
			wantCode: http.StatusBadRequest, // not used
			wantBody: "Bad Request", // not used
		},
		{
			name: "bad request body, cache",
			options: HandlerOptions{
				targetUrl: &url.URL{},
				badGatewayPage: "",
	
				cache: NewMemoryCache(1*MB, 1*MB),
				maxCacheableResponseBody: 16,
				maxRequestBody:            16,
				xSendfileEnabled: true,
			},
 
			wantCode: http.StatusBadRequest,
			wantBody: "Bad Request",
		},
		{
			name  : "bad request body, cache, max request body",
			options: HandlerOptions{
			 
				targetUrl: &url.URL{},
				maxRequestBody: 16,
				badGatewayPage: "",
				cache          : NewMemoryCache(1*MB, 1*MB),
				maxCacheableResponseBody: 17,
			},
			wantCode: http.StatusBadRequest, // not used
		},
		{
			name: "bad target url",
			options: HandlerOptions{
				targetUrl: nil,
				badGatewayPage: "",
				cache         : NewMemoryCache(1*MB, 1*MB),
				maxCacheableResponseBody: 0,
				maxRequestBody:            0,
				xSendfileEnabled: true,
			},
    
			wantCode: http.StatusBadRequest,
			wantBody: "Bad Request",
		},
		{
			name   : "bad target url, max request body",
			options: HandlerOptions{
			  
				targetUrl: nil,
				maxRequestBody: 16,
				badGatewayPage: "",
				cache:          NewMemoryCache(1*MB, 1*MB),
				xSendfileEnabled: true,
			},
	    
			wantCode: http.StatusBadRequest,
			
			wantBody: "Bad Request",
		},
		
		{
			name: "bad target url, max request body, max cacheable response body",
			options: HandlerOptions{
				targetUrl: nil,
				maxRequestBody:  16,
				badGatewayPage: "",
			
				cache: NewMemoryCache(1*MB, 1*MB),
				maxCacheableResponseBody: 17,
				xSendfileEnabled: true,
			},
       
			wantCode: http.StatusBadRequest,
			wantBody: "Bad Request",
		},
	}
}