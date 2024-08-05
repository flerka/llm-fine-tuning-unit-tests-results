package internal

import (
    "testing"
    "net/http"
    "strings"
)

type TestVariant struct {
    name string
    req *http.Request
    want *Variant
    wantErr bool
}

func TestNewVariant(t *testing.T) {
    tests := []TestVariant{
        {
            name: "Valid request",
            req: &http.Request{
                Method: "GET",
                Host:   "example.com",
                Path:   "/valid-path",
            },
            want: NewVariant(req),
            wantErr: false,
        },
        {
            name: "'Host' header missing",
            req: &http<|end|><|assistant|> package internal

import "testing"

type TestVariant struct{
    name string
    req  *http.Request
    want Variant
    wantErr bool
    errMsg string
}

func TestNewVarient(t *testing.T) int {
    tests := []TestVarient{
        {
            name : "Valid request",
            req :  &http.Request{
                Host:   "example1.com",
                Method : "GET",
                Path:    "/valid-path",
            }
        },
        {
           name : "Host header missing",
           req :  &http.Request{}
        },
        {
          name : "Invalid method",
          req :  &http.Request{"POST", "/valid-path", nil}
        },
        {
         name : "Invalid path",
         req :  &http.Request {"GET", "/", nil}
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T)  {
            var err error
            var v *Variant
            if tt.wantErr {
                err = tt.errMsg
                v = nil
           00000000 00 00 01 00 00 10 00 00                                 |......|
00000010 00 01                                                        |.|

00000001 00 01 10 00 10                                              |.|.|.|
00000