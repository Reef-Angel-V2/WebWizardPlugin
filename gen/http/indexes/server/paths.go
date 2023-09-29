// Code generated by goa v3.13.1, DO NOT EDIT.
//
// HTTP request path constructors for the indexes service.
//
// Command:
// $ goa gen github.com/arduino/arduino-create-agent/design

package server

// ListIndexesPath returns the URL path to the indexes service list HTTP endpoint.
func ListIndexesPath() string {
	return "/v2/pkgs/indexes"
}

// AddIndexesPath returns the URL path to the indexes service add HTTP endpoint.
func AddIndexesPath() string {
	return "/v2/pkgs/indexes/add"
}

// RemoveIndexesPath returns the URL path to the indexes service remove HTTP endpoint.
func RemoveIndexesPath() string {
	return "/v2/pkgs/indexes/delete"
}
