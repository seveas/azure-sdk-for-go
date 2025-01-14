//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package shared

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

// TokenRequestOptions contain specific parameter that may be used by credentials types when attempting to get a token.
type TokenRequestOptions struct {
	// Scopes contains the list of permission scopes required for the token.
	Scopes []string
	// TenantID contains the tenant ID to use in a multi-tenant authentication scenario, if TenantID is set
	// it will override the tenant ID that was added at credential creation time.
	TenantID string
}

// TokenCredential represents a credential capable of providing an OAuth token.
type TokenCredential interface {
	// GetToken requests an access token for the specified set of scopes.
	GetToken(ctx context.Context, options TokenRequestOptions) (*AccessToken, error)
}

// AccessToken represents an Azure service bearer access token with expiry information.
type AccessToken struct {
	Token     string
	ExpiresOn time.Time
}

// CtxWithHTTPHeaderKey is used as a context key for adding/retrieving http.Header.
type CtxWithHTTPHeaderKey struct{}

// CtxWithRetryOptionsKey is used as a context key for adding/retrieving RetryOptions.
type CtxWithRetryOptionsKey struct{}

// CtxIncludeResponseKey is used as a context key for retrieving the raw response.
type CtxIncludeResponseKey struct{}

type nopCloser struct {
	io.ReadSeeker
}

func (n nopCloser) Close() error {
	return nil
}

// NopCloser returns a ReadSeekCloser with a no-op close method wrapping the provided io.ReadSeeker.
func NopCloser(rs io.ReadSeeker) io.ReadSeekCloser {
	return nopCloser{rs}
}

// Delay waits for the duration to elapse or the context to be cancelled.
func Delay(ctx context.Context, delay time.Duration) error {
	select {
	case <-time.After(delay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ErrNoBody is returned if the response didn't contain a body.
var ErrNoBody = errors.New("the response did not contain a body")

// GetJSON reads the response body into a raw JSON object.
// It returns ErrNoBody if there was no content.
func GetJSON(resp *http.Response) (map[string]interface{}, error) {
	body, err := Payload(resp)
	if err != nil {
		return nil, err
	}
	if len(body) == 0 {
		return nil, ErrNoBody
	}
	// unmarshall the body to get the value
	var jsonBody map[string]interface{}
	if err = json.Unmarshal(body, &jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// RetryAfter returns non-zero if the response contains a Retry-After header value.
func RetryAfter(resp *http.Response) time.Duration {
	if resp == nil {
		return 0
	}
	ra := resp.Header.Get(HeaderRetryAfter)
	if ra == "" {
		return 0
	}
	// retry-after values are expressed in either number of
	// seconds or an HTTP-date indicating when to try again
	if retryAfter, _ := strconv.Atoi(ra); retryAfter > 0 {
		return time.Duration(retryAfter) * time.Second
	} else if t, err := time.Parse(time.RFC1123, ra); err == nil {
		return time.Until(t)
	}
	return 0
}

// HasStatusCode returns true if the Response's status code is one of the specified values.
func HasStatusCode(resp *http.Response, statusCodes ...int) bool {
	if resp == nil {
		return false
	}
	for _, sc := range statusCodes {
		if resp.StatusCode == sc {
			return true
		}
	}
	return false
}

// Payload reads and returns the response body or an error.
// On a successful read, the response body is cached.
// Subsequent reads will access the cached value.
func Payload(resp *http.Response) ([]byte, error) {
	// r.Body won't be a nopClosingBytesReader if downloading was skipped
	if buf, ok := resp.Body.(*NopClosingBytesReader); ok {
		return buf.Bytes(), nil
	}
	bytesBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	resp.Body = &NopClosingBytesReader{s: bytesBody, i: 0}
	return bytesBody, nil
}

// NopClosingBytesReader is an io.ReadSeekCloser around a byte slice.
// It also provides direct access to the byte slice to avoid rereading.
type NopClosingBytesReader struct {
	s []byte
	i int64
}

// NewNopClosingBytesReader creates a new NopClosingBytesReader around the specified byte slice.
func NewNopClosingBytesReader(data []byte) *NopClosingBytesReader {
	return &NopClosingBytesReader{s: data}
}

// Bytes returns the underlying byte slice.
func (r *NopClosingBytesReader) Bytes() []byte {
	return r.s
}

// Close implements the io.Closer interface.
func (*NopClosingBytesReader) Close() error {
	return nil
}

// Read implements the io.Reader interface.
func (r *NopClosingBytesReader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

// Set replaces the existing byte slice with the specified byte slice and resets the reader.
func (r *NopClosingBytesReader) Set(b []byte) {
	r.s = b
	r.i = 0
}

// Seek implements the io.Seeker interface.
func (r *NopClosingBytesReader) Seek(offset int64, whence int) (int64, error) {
	var i int64
	switch whence {
	case io.SeekStart:
		i = offset
	case io.SeekCurrent:
		i = r.i + offset
	case io.SeekEnd:
		i = int64(len(r.s)) + offset
	default:
		return 0, errors.New("nopClosingBytesReader: invalid whence")
	}
	if i < 0 {
		return 0, errors.New("nopClosingBytesReader: negative position")
	}
	r.i = i
	return i, nil
}

// TypeOfT returns the type of the generic type param.
func TypeOfT[T any]() reflect.Type {
	// you can't, at present, obtain the type of
	// a type parameter, so this is the trick
	return reflect.TypeOf((*T)(nil)).Elem()
}
