// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// GetAllJobsURL generates an URL for the get all jobs operation
type GetAllJobsURL struct {
	Description *string
	FullTime    *string
	Location    *string
	Page        *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAllJobsURL) WithBasePath(bp string) *GetAllJobsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetAllJobsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetAllJobsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/job"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var descriptionQ string
	if o.Description != nil {
		descriptionQ = *o.Description
	}
	if descriptionQ != "" {
		qs.Set("description", descriptionQ)
	}

	var fullTimeQ string
	if o.FullTime != nil {
		fullTimeQ = *o.FullTime
	}
	if fullTimeQ != "" {
		qs.Set("full_time", fullTimeQ)
	}

	var locationQ string
	if o.Location != nil {
		locationQ = *o.Location
	}
	if locationQ != "" {
		qs.Set("location", locationQ)
	}

	var pageQ string
	if o.Page != nil {
		pageQ = *o.Page
	}
	if pageQ != "" {
		qs.Set("page", pageQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetAllJobsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetAllJobsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetAllJobsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetAllJobsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetAllJobsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetAllJobsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}