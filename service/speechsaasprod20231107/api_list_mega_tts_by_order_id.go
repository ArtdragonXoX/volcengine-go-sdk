package speechsaasprod20231107

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListMegaTTSByOrderIDCommon = "ListMegaTTSByOrderID"

// ListMegaTTSByOrderIDCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListMegaTTSByOrderIDCommon operation. The "output" return
// value will be populated with the ListMegaTTSByOrderIDCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListMegaTTSByOrderIDCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListMegaTTSByOrderIDCommon Send returns without error.
//
// See ListMegaTTSByOrderIDCommon for more information on using the ListMegaTTSByOrderIDCommon
// API call, and error handling.
//
//	// Example sending a request using the ListMegaTTSByOrderIDCommonRequest method.
//	req, resp := client.ListMegaTTSByOrderIDCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPEECHSAASPROD20231107) ListMegaTTSByOrderIDCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListMegaTTSByOrderIDCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListMegaTTSByOrderIDCommon API operation for SPEECH_SAAS_PROD.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPEECH_SAAS_PROD's
// API operation ListMegaTTSByOrderIDCommon for usage and error information.
func (c *SPEECHSAASPROD20231107) ListMegaTTSByOrderIDCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListMegaTTSByOrderIDCommonRequest(input)
	return out, req.Send()
}

// ListMegaTTSByOrderIDCommonWithContext is the same as ListMegaTTSByOrderIDCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListMegaTTSByOrderIDCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPEECHSAASPROD20231107) ListMegaTTSByOrderIDCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListMegaTTSByOrderIDCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListMegaTTSByOrderID = "ListMegaTTSByOrderID"

// ListMegaTTSByOrderIDRequest generates a "volcengine/request.Request" representing the
// client's request for the ListMegaTTSByOrderID operation. The "output" return
// value will be populated with the ListMegaTTSByOrderID request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListMegaTTSByOrderID Request to send the API call to the service.
// the "output" return value is not valid until after ListMegaTTSByOrderID Send returns without error.
//
// See ListMegaTTSByOrderID for more information on using the ListMegaTTSByOrderID
// API call, and error handling.
//
//	// Example sending a request using the ListMegaTTSByOrderIDRequest method.
//	req, resp := client.ListMegaTTSByOrderIDRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPEECHSAASPROD20231107) ListMegaTTSByOrderIDRequest(input *ListMegaTTSByOrderIDInput) (req *request.Request, output *ListMegaTTSByOrderIDOutput) {
	op := &request.Operation{
		Name:       opListMegaTTSByOrderID,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListMegaTTSByOrderIDInput{}
	}

	output = &ListMegaTTSByOrderIDOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListMegaTTSByOrderID API operation for SPEECH_SAAS_PROD.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPEECH_SAAS_PROD's
// API operation ListMegaTTSByOrderID for usage and error information.
func (c *SPEECHSAASPROD20231107) ListMegaTTSByOrderID(input *ListMegaTTSByOrderIDInput) (*ListMegaTTSByOrderIDOutput, error) {
	req, out := c.ListMegaTTSByOrderIDRequest(input)
	return out, req.Send()
}

// ListMegaTTSByOrderIDWithContext is the same as ListMegaTTSByOrderID with the addition of
// the ability to pass a context and additional request options.
//
// See ListMegaTTSByOrderID for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPEECHSAASPROD20231107) ListMegaTTSByOrderIDWithContext(ctx volcengine.Context, input *ListMegaTTSByOrderIDInput, opts ...request.Option) (*ListMegaTTSByOrderIDOutput, error) {
	req, out := c.ListMegaTTSByOrderIDRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListMegaTTSByOrderIDInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AppID *string `type:"string" json:",omitempty"`

	OrderID *string `type:"string" json:",omitempty"`

	PageNumber *int64 `type:"integer" json:",omitempty"`

	PageSize *int64 `type:"integer" json:",omitempty"`

	NextToken *string `type:"string" json:",omitempty"`

	MaxResults *int64 `type:"integer" json:",omitempty"`
}

// String returns the string representation
func (s ListMegaTTSByOrderIDInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListMegaTTSByOrderIDInput) GoString() string {
	return s.String()
}

// SetAppID sets the AppID field's value.
func (s *ListMegaTTSByOrderIDInput) SetAppID(v string) *ListMegaTTSByOrderIDInput {
	s.AppID = &v
	return s
}

// SetOrderID sets the OrderID field's value.
func (s *ListMegaTTSByOrderIDInput) SetOrderID(v string) *ListMegaTTSByOrderIDInput {
	s.OrderID = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListMegaTTSByOrderIDInput) SetPageNumber(v int64) *ListMegaTTSByOrderIDInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListMegaTTSByOrderIDInput) SetPageSize(v int64) *ListMegaTTSByOrderIDInput {
	s.PageSize = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListMegaTTSByOrderIDInput) SetNextToken(v string) *ListMegaTTSByOrderIDInput {
	s.NextToken = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListMegaTTSByOrderIDInput) SetMaxResults(v int64) *ListMegaTTSByOrderIDInput {
	s.MaxResults = &v
	return s
}

type ListMegaTTSByOrderIDOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	OrderBuyQuantity *int64 `type:"integer" json:",omitempty"`

	OrderBuyCurrentSuccessQuantity *int64 `type:"integer" json:",omitempty"`

	OrderBuyCurrentFailedQuantity *int64 `type:"integer" json:",omitempty"`

	IsAllProcessed *bool `type:"boolean" json:",omitempty"`

	Statuses []*StatusForListMegaTTSByOrderIDOutput `type:"list" json:",omitempty"`

	PageNumber *int64 `type:"integer" json:",omitempty"`

	PageSize *int64 `type:"integer" json:",omitempty"`

	TotalCount *int64 `type:"integer" json:",omitempty"`

	NextToken *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListMegaTTSByOrderIDOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListMegaTTSByOrderIDOutput) GoString() string {
	return s.String()
}

// SetOrderBuyQuantity sets the OrderBuyQuantity field's value.
func (s *ListMegaTTSByOrderIDOutput) SetOrderBuyQuantity(v int64) *ListMegaTTSByOrderIDOutput {
	s.OrderBuyQuantity = &v
	return s
}

// SetOrderBuyCurrentSuccessQuantity sets the OrderBuyCurrentSuccessQuantity field's value.
func (s *ListMegaTTSByOrderIDOutput) SetOrderBuyCurrentSuccessQuantity(v int64) *ListMegaTTSByOrderIDOutput {
	s.OrderBuyCurrentSuccessQuantity = &v
	return s
}

// SetOrderBuyCurrentFailedQuantity sets the OrderBuyCurrentFailedQuantity field's value.
func (s *ListMegaTTSByOrderIDOutput) SetOrderBuyCurrentFailedQuantity(v int64) *ListMegaTTSByOrderIDOutput {
	s.OrderBuyCurrentFailedQuantity = &v
	return s
}

// SetIsAllProcessed sets the IsAllProcessed field's value.
func (s *ListMegaTTSByOrderIDOutput) SetIsAllProcessed(v bool) *ListMegaTTSByOrderIDOutput {
	s.IsAllProcessed = &v
	return s
}

// SetStatuses sets the Statuses field's value.
func (s *ListMegaTTSByOrderIDOutput) SetStatuses(v []*StatusForListMegaTTSByOrderIDOutput) *ListMegaTTSByOrderIDOutput {
	s.Statuses = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListMegaTTSByOrderIDOutput) SetPageNumber(v int64) *ListMegaTTSByOrderIDOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListMegaTTSByOrderIDOutput) SetPageSize(v int64) *ListMegaTTSByOrderIDOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListMegaTTSByOrderIDOutput) SetTotalCount(v int64) *ListMegaTTSByOrderIDOutput {
	s.TotalCount = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListMegaTTSByOrderIDOutput) SetNextToken(v string) *ListMegaTTSByOrderIDOutput {
	s.NextToken = &v
	return s
}

type StatusForListMegaTTSByOrderIDOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	SpeakerID *string `type:"string" json:",omitempty"`

	InstanceNO *string `type:"string" json:",omitempty"`

	InstanceStatus *string `type:"string" json:",omitempty"`

	IsActivatable *bool `type:"boolean" json:",omitempty"`

	State *string `type:"string" json:",omitempty"`

	DemoAudio *string `type:"string" json:",omitempty"`

	Version *string `type:"string" json:",omitempty"`

	CreateTime *int64 `type:"long" json:",omitempty"`

	ExpireTime *int64 `type:"long" json:",omitempty"`

	OrderTime *int64 `type:"long" json:",omitempty"`

	Alias *string `type:"string" json:",omitempty"`

	AvailableTrainingTimes *int64 `type:"integer" json:",omitempty"`

	ResourceID *string `type:"string" json:",omitempty"`

	Avatar *string `type:"string" json:",omitempty"`

	Labels []*string `type:"list" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	ModelTypeDetails interface{} `type:"structure" json:",omitempty"`

	CloneSource *int64 `type:"integer" json:",omitempty"`
}

// String returns the string representation
func (s StatusForListMegaTTSByOrderIDOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatusForListMegaTTSByOrderIDOutput) GoString() string {
	return s.String()
}

// SetSpeakerID sets the SpeakerID field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetSpeakerID(v string) *StatusForListMegaTTSByOrderIDOutput {
	s.SpeakerID = &v
	return s
}

// SetInstanceNO sets the InstanceNO field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetInstanceNO(v string) *StatusForListMegaTTSByOrderIDOutput {
	s.InstanceNO = &v
	return s
}

// SetIsActivatable sets the IsActivatable field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetIsActivatable(v bool) *StatusForListMegaTTSByOrderIDOutput {
	s.IsActivatable = &v
	return s
}

// SetState sets the State field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetState(v string) *StatusForListMegaTTSByOrderIDOutput {
	s.State = &v
	return s
}

// SetDemoAudio sets the DemoAudio field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetDemoAudio(v string) *StatusForListMegaTTSByOrderIDOutput {
	s.DemoAudio = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetVersion(v string) *StatusForListMegaTTSByOrderIDOutput {
	s.Version = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetCreateTime(v int64) *StatusForListMegaTTSByOrderIDOutput {
	s.CreateTime = &v
	return s
}

// SetExpireTime sets the ExpireTime field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetExpireTime(v int64) *StatusForListMegaTTSByOrderIDOutput {
	s.ExpireTime = &v
	return s
}

// SetOrderTime sets the OrderTime field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetOrderTime(v int64) *StatusForListMegaTTSByOrderIDOutput {
	s.OrderTime = &v
	return s
}

// SetAlias sets the Alias field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetAlias(v string) *StatusForListMegaTTSByOrderIDOutput {
	s.Alias = &v
	return s
}

// SetAvailableTrainingTimes sets the AvailableTrainingTimes field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetAvailableTrainingTimes(v int64) *StatusForListMegaTTSByOrderIDOutput {
	s.AvailableTrainingTimes = &v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetInstanceStatus(v string) *StatusForListMegaTTSByOrderIDOutput {
	s.InstanceStatus = &v
	return s
}

// SetResourceID sets the ResourceID field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetResourceID(v string) *StatusForListMegaTTSByOrderIDOutput {
	s.ResourceID = &v
	return s
}

// SetAvatar sets the Avatar field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetAvatar(v string) *StatusForListMegaTTSByOrderIDOutput {
	s.Avatar = &v
	return s
}

// SetLabels sets the Labels field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetLabels(v []*string) *StatusForListMegaTTSByOrderIDOutput {
	s.Labels = v
	return s
}

// SetDescription sets the Description field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetDescription(v string) *StatusForListMegaTTSByOrderIDOutput {
	s.Description = &v
	return s
}

// SetModelTypeDetails sets the ModelTypeDetails field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetModelTypeDetails(v interface{}) *StatusForListMegaTTSByOrderIDOutput {
	s.ModelTypeDetails = v
	return s
}

// SetCloneSource sets the CloneSource field's value.
func (s *StatusForListMegaTTSByOrderIDOutput) SetCloneSource(v int64) *StatusForListMegaTTSByOrderIDOutput {
	s.CloneSource = &v
	return s
}
