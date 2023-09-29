// Code generated by goa v3.13.1, DO NOT EDIT.
//
// indexes HTTP client types
//
// Command:
// $ goa gen github.com/arduino/arduino-create-agent/design

package client

import (
	indexes "github.com/arduino/arduino-create-agent/gen/indexes"
	indexesviews "github.com/arduino/arduino-create-agent/gen/indexes/views"
	goa "goa.design/goa/v3/pkg"
)

// AddRequestBody is the type of the "indexes" service "add" endpoint HTTP
// request body.
type AddRequestBody struct {
	// The url of the index file
	URL string `form:"url" json:"url" xml:"url"`
}

// RemoveRequestBody is the type of the "indexes" service "remove" endpoint
// HTTP request body.
type RemoveRequestBody struct {
	// The url of the index file
	URL string `form:"url" json:"url" xml:"url"`
}

// AddResponseBody is the type of the "indexes" service "add" endpoint HTTP
// response body.
type AddResponseBody struct {
	// The status of the operation
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// RemoveResponseBody is the type of the "indexes" service "remove" endpoint
// HTTP response body.
type RemoveResponseBody struct {
	// The status of the operation
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
}

// ListInvalidURLResponseBody is the type of the "indexes" service "list"
// endpoint HTTP response body for the "invalid_url" error.
type ListInvalidURLResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// AddInvalidURLResponseBody is the type of the "indexes" service "add"
// endpoint HTTP response body for the "invalid_url" error.
type AddInvalidURLResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// RemoveInvalidURLResponseBody is the type of the "indexes" service "remove"
// endpoint HTTP response body for the "invalid_url" error.
type RemoveInvalidURLResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// NewAddRequestBody builds the HTTP request body from the payload of the "add"
// endpoint of the "indexes" service.
func NewAddRequestBody(p *indexes.IndexPayload) *AddRequestBody {
	body := &AddRequestBody{
		URL: p.URL,
	}
	return body
}

// NewRemoveRequestBody builds the HTTP request body from the payload of the
// "remove" endpoint of the "indexes" service.
func NewRemoveRequestBody(p *indexes.IndexPayload) *RemoveRequestBody {
	body := &RemoveRequestBody{
		URL: p.URL,
	}
	return body
}

// NewListInvalidURL builds a indexes service list endpoint invalid_url error.
func NewListInvalidURL(body *ListInvalidURLResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewAddOperationOK builds a "indexes" service "add" endpoint result from a
// HTTP "OK" response.
func NewAddOperationOK(body *AddResponseBody) *indexesviews.OperationView {
	v := &indexesviews.OperationView{
		Status: body.Status,
	}

	return v
}

// NewAddInvalidURL builds a indexes service add endpoint invalid_url error.
func NewAddInvalidURL(body *AddInvalidURLResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewRemoveOperationOK builds a "indexes" service "remove" endpoint result
// from a HTTP "OK" response.
func NewRemoveOperationOK(body *RemoveResponseBody) *indexesviews.OperationView {
	v := &indexesviews.OperationView{
		Status: body.Status,
	}

	return v
}

// NewRemoveInvalidURL builds a indexes service remove endpoint invalid_url
// error.
func NewRemoveInvalidURL(body *RemoveInvalidURLResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateListInvalidURLResponseBody runs the validations defined on
// list_invalid_url_response_body
func ValidateListInvalidURLResponseBody(body *ListInvalidURLResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateAddInvalidURLResponseBody runs the validations defined on
// add_invalid_url_response_body
func ValidateAddInvalidURLResponseBody(body *AddInvalidURLResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateRemoveInvalidURLResponseBody runs the validations defined on
// remove_invalid_url_response_body
func ValidateRemoveInvalidURLResponseBody(body *RemoveInvalidURLResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}
