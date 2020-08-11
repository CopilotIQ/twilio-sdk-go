// This is an autogenerated file. DO NOT MODIFY
package items

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/utils"
)

// SyncMapItemsPageOptions defines the query options for the api operation
type SyncMapItemsPageOptions struct {
	PageSize  *int
	Page      *int
	PageToken *string
	Order     *string
	From      *string
	Bounds    *string
}

type PageMetaResponse struct {
	FirstPageURL    string  `json:"first_page_url"`
	Key             string  `json:"key"`
	NextPageURL     *string `json:"next_page_url,omitempty"`
	Page            int     `json:"page"`
	PageSize        int     `json:"page_size"`
	PreviousPageURL *string `json:"previous_page_url,omitempty"`
	URL             string  `json:"url"`
}

type PageSyncMapItemResponse struct {
	AccountSid  string                 `json:"account_sid"`
	CreatedBy   string                 `json:"created_by"`
	Data        map[string]interface{} `json:"data"`
	DateCreated time.Time              `json:"date_created"`
	DateExpires *time.Time             `json:"date_expires,omitempty"`
	DateUpdated *time.Time             `json:"date_updated,omitempty"`
	Key         string                 `json:"key"`
	MapSid      string                 `json:"map_sid"`
	Revision    string                 `json:"revision"`
	ServiceSid  string                 `json:"service_Sid"`
	URL         string                 `json:"url"`
}

// SyncMapItemsPageResponse defines the response fields for the map items page
type SyncMapItemsPageResponse struct {
	Meta         PageMetaResponse          `json:"meta"`
	SyncMapItems []PageSyncMapItemResponse `json:"items"`
}

// Page retrieves a page of map items
// See https://www.twilio.com/docs/sync/api/map-item-resource#read-all-mapitem-resources for more details
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (c Client) Page(options *SyncMapItemsPageOptions) (*SyncMapItemsPageResponse, error) {
	return c.PageWithContext(context.Background(), options)
}

// PageWithContext retrieves a page of map items
// See https://www.twilio.com/docs/sync/api/map-item-resource#read-all-mapitem-resources for more details
func (c Client) PageWithContext(context context.Context, options *SyncMapItemsPageOptions) (*SyncMapItemsPageResponse, error) {
	op := client.Operation{
		Method: http.MethodGet,
		URI:    "/Services/{serviceSid}/Maps/{syncMapSid}/Items",
		PathParams: map[string]string{
			"serviceSid": c.serviceSid,
			"syncMapSid": c.syncMapSid,
		},
		QueryParams: utils.StructToStringMap(options),
	}

	response := &SyncMapItemsPageResponse{}
	if err := c.client.Send(context, op, nil, response); err != nil {
		return nil, err
	}
	return response, nil
}

// SyncMapItemsPaginator defines the fields for makings paginated api calls
// SyncMapItems is an array of syncmapitems that have been returned from all of the page calls
type SyncMapItemsPaginator struct {
	options      *SyncMapItemsPageOptions
	Page         *SyncMapItemsPage
	SyncMapItems []PageSyncMapItemResponse
}

// NewSyncMapItemsPaginator creates a new instance of the paginator for Page.
func (c *Client) NewSyncMapItemsPaginator() *SyncMapItemsPaginator {
	return c.NewSyncMapItemsPaginatorWithOptions(nil)
}

// NewSyncMapItemsPaginatorWithOptions creates a new instance of the paginator for Page with options.
func (c *Client) NewSyncMapItemsPaginatorWithOptions(options *SyncMapItemsPageOptions) *SyncMapItemsPaginator {
	return &SyncMapItemsPaginator{
		options: options,
		Page: &SyncMapItemsPage{
			CurrentPage: nil,
			Error:       nil,
			client:      c,
		},
		SyncMapItems: make([]PageSyncMapItemResponse, 0),
	}
}

// SyncMapItemsPage defines the fields for the page
// The CurrentPage and Error fields can be used to access the PageSyncMapItemResponse or error that is returned from the api call(s)
type SyncMapItemsPage struct {
	client *Client

	CurrentPage *SyncMapItemsPageResponse
	Error       error
}

// CurrentPage retrieves the results for the current page
func (p *SyncMapItemsPaginator) CurrentPage() *SyncMapItemsPageResponse {
	return p.Page.CurrentPage
}

// Error retrieves the error returned from the page
func (p *SyncMapItemsPaginator) Error() error {
	return p.Page.Error
}

// Next retrieves the next page of results.
// Next will return false when either an error occurs or there are no more pages to iterate
// Context is defaulted to Background. See https://golang.org/pkg/context/#Background for more information
func (p *SyncMapItemsPaginator) Next() bool {
	return p.NextWithContext(context.Background())
}

// NextWithContext retrieves the next page of results.
// NextWithContext will return false when either an error occurs or there are no more pages to iterate
func (p *SyncMapItemsPaginator) NextWithContext(context context.Context) bool {
	options := p.options

	if options == nil {
		options = &SyncMapItemsPageOptions{}
	}

	if p.CurrentPage() != nil {
		nextPageURL := p.CurrentPage().Meta.NextPageURL

		if nextPageURL == nil {
			return false
		}

		parsedURL, err := url.Parse(*nextPageURL)
		if err != nil {
			p.Page.Error = err
			return false
		}

		options.PageToken = utils.String(parsedURL.Query().Get("PageToken"))

		page, pageErr := strconv.Atoi(parsedURL.Query().Get("Page"))
		if pageErr != nil {
			p.Page.Error = pageErr
			return false
		}
		options.Page = utils.Int(page)

		pageSize, pageSizeErr := strconv.Atoi(parsedURL.Query().Get("PageSize"))
		if pageSizeErr != nil {
			p.Page.Error = pageSizeErr
			return false
		}
		options.PageSize = utils.Int(pageSize)
	}

	resp, err := p.Page.client.PageWithContext(context, options)
	p.Page.CurrentPage = resp
	p.Page.Error = err

	if p.Page.Error == nil {
		p.SyncMapItems = append(p.SyncMapItems, resp.SyncMapItems...)
	}

	return p.Page.Error == nil
}