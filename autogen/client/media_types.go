// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "feedpushr": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/ncarlier/feedpushr/v2/design
// --out=/home/nicolas/workspace/feedpushr/autogen
// --version=v1.4.3

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

// A RSS feed (default view)
//
// Identifier: application/vnd.feedpushr.feed.v1+json; view=default
type Feed struct {
	// Date of creation
	Cdate time.Time `form:"cdate" json:"cdate" yaml:"cdate" xml:"cdate"`
	// Number of consecutive aggregation errors
	ErrorCount *int `form:"errorCount,omitempty" json:"errorCount,omitempty" yaml:"errorCount,omitempty" xml:"errorCount,omitempty"`
	// Last aggregation error
	ErrorMsg *string `form:"errorMsg,omitempty" json:"errorMsg,omitempty" yaml:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// URL of the feed website
	HTMLURL *string `form:"htmlUrl,omitempty" json:"htmlUrl,omitempty" yaml:"htmlUrl,omitempty" xml:"htmlUrl,omitempty"`
	// URL of the PubSubHubbud hub
	HubURL *string `form:"hubUrl,omitempty" json:"hubUrl,omitempty" yaml:"hubUrl,omitempty" xml:"hubUrl,omitempty"`
	// ID of feed (MD5 of the xmlUrl)
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Last aggregation pass
	LastCheck *time.Time `form:"lastCheck,omitempty" json:"lastCheck,omitempty" yaml:"lastCheck,omitempty" xml:"lastCheck,omitempty"`
	// Date of modification
	Mdate time.Time `form:"mdate" json:"mdate" yaml:"mdate" xml:"mdate"`
	// Total number of processed items
	NbProcessedItems *int `form:"nbProcessedItems,omitempty" json:"nbProcessedItems,omitempty" yaml:"nbProcessedItems,omitempty" xml:"nbProcessedItems,omitempty"`
	// Next aggregation pass
	NextCheck *time.Time `form:"nextCheck,omitempty" json:"nextCheck,omitempty" yaml:"nextCheck,omitempty" xml:"nextCheck,omitempty"`
	// Aggregation status
	Status *string `form:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty" xml:"status,omitempty"`
	// List of tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" yaml:"tags,omitempty" xml:"tags,omitempty"`
	// Text attribute of the Feed
	Text *string `form:"text,omitempty" json:"text,omitempty" yaml:"text,omitempty" xml:"text,omitempty"`
	// Title of the Feed
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
	// URL of the XML feed
	XMLURL string `form:"xmlUrl" json:"xmlUrl" yaml:"xmlUrl" xml:"xmlUrl"`
}

// Validate validates the Feed media type instance.
func (mt *Feed) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.XMLURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "xmlUrl"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	if mt.Status != nil {
		if !(*mt.Status == "running" || *mt.Status == "stopped") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.status`, *mt.Status, []interface{}{"running", "stopped"}))
		}
	}
	return
}

// A RSS feed (link view)
//
// Identifier: application/vnd.feedpushr.feed.v1+json; view=link
type FeedLink struct {
	// ID of feed (MD5 of the xmlUrl)
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// URL of the XML feed
	XMLURL string `form:"xmlUrl" json:"xmlUrl" yaml:"xmlUrl" xml:"xmlUrl"`
}

// Validate validates the FeedLink media type instance.
func (mt *FeedLink) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.XMLURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "xmlUrl"))
	}
	return
}

// A RSS feed (tiny view)
//
// Identifier: application/vnd.feedpushr.feed.v1+json; view=tiny
type FeedTiny struct {
	// Date of creation
	Cdate time.Time `form:"cdate" json:"cdate" yaml:"cdate" xml:"cdate"`
	// ID of feed (MD5 of the xmlUrl)
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// List of tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" yaml:"tags,omitempty" xml:"tags,omitempty"`
	// Title of the Feed
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
	// URL of the XML feed
	XMLURL string `form:"xmlUrl" json:"xmlUrl" yaml:"xmlUrl" xml:"xmlUrl"`
}

// Validate validates the FeedTiny media type instance.
func (mt *FeedTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.XMLURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "xmlUrl"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	return
}

// DecodeFeed decodes the Feed instance encoded in resp body.
func (c *Client) DecodeFeed(resp *http.Response) (*Feed, error) {
	var decoded Feed
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeFeedLink decodes the FeedLink instance encoded in resp body.
func (c *Client) DecodeFeedLink(resp *http.Response) (*FeedLink, error) {
	var decoded FeedLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeFeedTiny decodes the FeedTiny instance encoded in resp body.
func (c *Client) DecodeFeedTiny(resp *http.Response) (*FeedTiny, error) {
	var decoded FeedTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// FeedCollection is the media type for an array of Feed (default view)
//
// Identifier: application/vnd.feedpushr.feed.v1+json; type=collection; view=default
type FeedCollection []*Feed

// Validate validates the FeedCollection media type instance.
func (mt FeedCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// FeedCollection is the media type for an array of Feed (link view)
//
// Identifier: application/vnd.feedpushr.feed.v1+json; type=collection; view=link
type FeedLinkCollection []*FeedLink

// Validate validates the FeedLinkCollection media type instance.
func (mt FeedLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// FeedCollection is the media type for an array of Feed (tiny view)
//
// Identifier: application/vnd.feedpushr.feed.v1+json; type=collection; view=tiny
type FeedTinyCollection []*FeedTiny

// Validate validates the FeedTinyCollection media type instance.
func (mt FeedTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeFeedCollection decodes the FeedCollection instance encoded in resp body.
func (c *Client) DecodeFeedCollection(resp *http.Response) (FeedCollection, error) {
	var decoded FeedCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeFeedLinkCollection decodes the FeedLinkCollection instance encoded in resp body.
func (c *Client) DecodeFeedLinkCollection(resp *http.Response) (FeedLinkCollection, error) {
	var decoded FeedLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeFeedTinyCollection decodes the FeedTinyCollection instance encoded in resp body.
func (c *Client) DecodeFeedTinyCollection(resp *http.Response) (FeedTinyCollection, error) {
	var decoded FeedTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A pagignated list of feeds (default view)
//
// Identifier: application/vnd.feedpushr.feeds-page.v1+json; view=default
type FeedsPage struct {
	// Current page number
	Current int `form:"current" json:"current" yaml:"current" xml:"current"`
	// List of feeds
	Data FeedCollection `form:"data" json:"data" yaml:"data" xml:"data"`
	// Max number of feeds by page
	Limit int `form:"limit" json:"limit" yaml:"limit" xml:"limit"`
	// Total number of feeds
	Total int `form:"total" json:"total" yaml:"total" xml:"total"`
}

// Validate validates the FeedsPage media type instance.
func (mt *FeedsPage) Validate() (err error) {

	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if err2 := mt.Data.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeFeedsPage decodes the FeedsPage instance encoded in resp body.
func (c *Client) DecodeFeedsPage(resp *http.Response) (*FeedsPage, error) {
	var decoded FeedsPage
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// The filter specification (default view)
//
// Identifier: application/vnd.feedpushr.filter-spec.v1+json; view=default
type FilterSpec struct {
	// Description of the filter
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// Name of the filter
	Name  string             `form:"name" json:"name" yaml:"name" xml:"name"`
	Props PropSpecCollection `form:"props" json:"props" yaml:"props" xml:"props"`
}

// Validate validates the FilterSpec media type instance.
func (mt *FilterSpec) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.Props == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "props"))
	}
	if err2 := mt.Props.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeFilterSpec decodes the FilterSpec instance encoded in resp body.
func (c *Client) DecodeFilterSpec(resp *http.Response) (*FilterSpec, error) {
	var decoded FilterSpec
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// FilterSpecCollection is the media type for an array of FilterSpec (default view)
//
// Identifier: application/vnd.feedpushr.filter-spec.v1+json; type=collection; view=default
type FilterSpecCollection []*FilterSpec

// Validate validates the FilterSpecCollection media type instance.
func (mt FilterSpecCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeFilterSpecCollection decodes the FilterSpecCollection instance encoded in resp body.
func (c *Client) DecodeFilterSpecCollection(resp *http.Response) (FilterSpecCollection, error) {
	var decoded FilterSpecCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// A filter (default view)
//
// Identifier: application/vnd.feedpushr.filter.v1+json; view=default
type Filter struct {
	// Alias of the filter
	Alias string `form:"alias" json:"alias" yaml:"alias" xml:"alias"`
	// Conditional expression of the filter
	Condition string `form:"condition" json:"condition" yaml:"condition" xml:"condition"`
	// Description of the filter
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// Status
	Enabled bool `form:"enabled" json:"enabled" yaml:"enabled" xml:"enabled"`
	// ID of the filter
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// Name of the filter
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Filter properties
	Props map[string]interface{} `form:"props,omitempty" json:"props,omitempty" yaml:"props,omitempty" xml:"props,omitempty"`
}

// Validate validates the Filter media type instance.
func (mt *Filter) Validate() (err error) {

	if mt.Alias == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "alias"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.Condition == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "condition"))
	}
	return
}

// DecodeFilter decodes the Filter instance encoded in resp body.
func (c *Client) DecodeFilter(resp *http.Response) (*Filter, error) {
	var decoded Filter
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// FilterCollection is the media type for an array of Filter (default view)
//
// Identifier: application/vnd.feedpushr.filter.v1+json; type=collection; view=default
type FilterCollection []*Filter

// Validate validates the FilterCollection media type instance.
func (mt FilterCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeFilterCollection decodes the FilterCollection instance encoded in resp body.
func (c *Client) DecodeFilterCollection(resp *http.Response) (FilterCollection, error) {
	var decoded FilterCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// HAL link (default view)
//
// Identifier: application/vnd.feedpushr.hal-links.v1+json; view=default
type HALLink struct {
	// Link's destination
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
}

// Validate validates the HALLink media type instance.
func (mt *HALLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// DecodeHALLink decodes the HALLink instance encoded in resp body.
func (c *Client) DecodeHALLink(resp *http.Response) (*HALLink, error) {
	var decoded HALLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// API info (default view)
//
// Identifier: application/vnd.feedpushr.info.v1+json; view=default
type Info struct {
	// HAL links
	Links map[string]*HALLink `form:"_links" json:"_links" yaml:"_links" xml:"_links"`
	// Service description
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// Service name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Service version
	Version string `form:"version" json:"version" yaml:"version" xml:"version"`
}

// Validate validates the Info media type instance.
func (mt *Info) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	if mt.Links == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "_links"))
	}
	for _, e := range mt.Links {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeInfo decodes the Info instance encoded in resp body.
func (c *Client) DecodeInfo(resp *http.Response) (*Info, error) {
	var decoded Info
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// The output channel specification (default view)
//
// Identifier: application/vnd.feedpushr.output-spec.v1+json; view=default
type OutputSpec struct {
	// Description of the output channel
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// Name of the output channel
	Name  string             `form:"name" json:"name" yaml:"name" xml:"name"`
	Props PropSpecCollection `form:"props" json:"props" yaml:"props" xml:"props"`
}

// Validate validates the OutputSpec media type instance.
func (mt *OutputSpec) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.Props == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "props"))
	}
	if err2 := mt.Props.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeOutputSpec decodes the OutputSpec instance encoded in resp body.
func (c *Client) DecodeOutputSpec(resp *http.Response) (*OutputSpec, error) {
	var decoded OutputSpec
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// OutputSpecCollection is the media type for an array of OutputSpec (default view)
//
// Identifier: application/vnd.feedpushr.output-spec.v1+json; type=collection; view=default
type OutputSpecCollection []*OutputSpec

// Validate validates the OutputSpecCollection media type instance.
func (mt OutputSpecCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeOutputSpecCollection decodes the OutputSpecCollection instance encoded in resp body.
func (c *Client) DecodeOutputSpecCollection(resp *http.Response) (OutputSpecCollection, error) {
	var decoded OutputSpecCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// The output channel (default view)
//
// Identifier: application/vnd.feedpushr.output.v1+json; view=default
type Output struct {
	// Alias of the output channel
	Alias string `form:"alias" json:"alias" yaml:"alias" xml:"alias"`
	// Conditional expression of the filter
	Condition string `form:"condition" json:"condition" yaml:"condition" xml:"condition"`
	// Description of the output channel
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// Status
	Enabled bool `form:"enabled" json:"enabled" yaml:"enabled" xml:"enabled"`
	// ID of the output
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// Name of the output channel
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Output channel properties
	Props map[string]interface{} `form:"props,omitempty" json:"props,omitempty" yaml:"props,omitempty" xml:"props,omitempty"`
}

// Validate validates the Output media type instance.
func (mt *Output) Validate() (err error) {

	if mt.Alias == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "alias"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.Condition == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "condition"))
	}
	return
}

// DecodeOutput decodes the Output instance encoded in resp body.
func (c *Client) DecodeOutput(resp *http.Response) (*Output, error) {
	var decoded Output
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// OutputCollection is the media type for an array of Output (default view)
//
// Identifier: application/vnd.feedpushr.output.v1+json; type=collection; view=default
type OutputCollection []*Output

// Validate validates the OutputCollection media type instance.
func (mt OutputCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeOutputCollection decodes the OutputCollection instance encoded in resp body.
func (c *Client) DecodeOutputCollection(resp *http.Response) (OutputCollection, error) {
	var decoded OutputCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// The specification of a property (default view)
//
// Identifier: application/vnd.feedpushr.prop-spec.v1+json; view=default
type PropSpec struct {
	// Description of the output channel
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// Name of the property
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Property options
	Options map[string]string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
	// Property type ('text', 'url', ...)
	Type string `form:"type" json:"type" yaml:"type" xml:"type"`
}

// Validate validates the PropSpec media type instance.
func (mt *PropSpec) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	return
}

// DecodePropSpec decodes the PropSpec instance encoded in resp body.
func (c *Client) DecodePropSpec(resp *http.Response) (*PropSpec, error) {
	var decoded PropSpec
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// PropSpecCollection is the media type for an array of PropSpec (default view)
//
// Identifier: application/vnd.feedpushr.prop-spec.v1+json; type=collection; view=default
type PropSpecCollection []*PropSpec

// Validate validates the PropSpecCollection media type instance.
func (mt PropSpecCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodePropSpecCollection decodes the PropSpecCollection instance encoded in resp body.
func (c *Client) DecodePropSpecCollection(resp *http.Response) (PropSpecCollection, error) {
	var decoded PropSpecCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
