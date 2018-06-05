package features

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// Client is the client for the Features methods of the Features service.
type Client struct {
	ManagementClient
}

// NewClient creates an instance of the Client client.
func NewClient(subscriptionID string) Client {
	return NewClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClientWithBaseURI creates an instance of the Client client.
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return Client{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get all features under the subscription.
//
// resourceProviderNamespace is namespace of the resource provider.
// featureName is previewed feature name in the resource provider.
func (client Client) Get(resourceProviderNamespace string, featureName string) (result FeatureResult, err error) {
	req, err := client.GetPreparer(resourceProviderNamespace, featureName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "features.Client", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "features.Client", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client Client) GetPreparer(resourceProviderNamespace string, featureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"featureName":               autorest.Encode("path", featureName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features/{featureName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client Client) GetResponder(resp *http.Response) (result FeatureResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of previewed features of a resource provider.
//
// resourceProviderNamespace is the namespace of the resource provider.
func (client Client) List(resourceProviderNamespace string) (result FeatureOperationsListResult, err error) {
	req, err := client.ListPreparer(resourceProviderNamespace)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "features.Client", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "features.Client", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client Client) ListPreparer(resourceProviderNamespace string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client Client) ListResponder(resp *http.Response) (result FeatureOperationsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client Client) ListNextResults(lastResults FeatureOperationsListResult) (result FeatureOperationsListResult, err error) {
	req, err := lastResults.FeatureOperationsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "features.Client", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "features.Client", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "List", resp, "Failure responding to next results request request")
	}

	return
}

// ListAll gets a list of previewed features for all the providers in the
// current subscription.
func (client Client) ListAll() (result FeatureOperationsListResult, err error) {
	req, err := client.ListAllPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "features.Client", "ListAll", nil, "Failure preparing request")
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "features.Client", "ListAll", resp, "Failure sending request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client Client) ListAllPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/features", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client Client) ListAllResponder(resp *http.Response) (result FeatureOperationsListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllNextResults retrieves the next set of results, if any.
func (client Client) ListAllNextResults(lastResults FeatureOperationsListResult) (result FeatureOperationsListResult, err error) {
	req, err := lastResults.FeatureOperationsListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "features.Client", "ListAll", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "features.Client", "ListAll", resp, "Failure sending next results request request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "ListAll", resp, "Failure responding to next results request request")
	}

	return
}

// Register registers for a previewed feature of a resource provider.
//
// resourceProviderNamespace is namespace of the resource provider.
// featureName is previewed feature name in the resource provider.
func (client Client) Register(resourceProviderNamespace string, featureName string) (result FeatureResult, err error) {
	req, err := client.RegisterPreparer(resourceProviderNamespace, featureName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "features.Client", "Register", nil, "Failure preparing request")
	}

	resp, err := client.RegisterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "features.Client", "Register", resp, "Failure sending request")
	}

	result, err = client.RegisterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "features.Client", "Register", resp, "Failure responding to request")
	}

	return
}

// RegisterPreparer prepares the Register request.
func (client Client) RegisterPreparer(resourceProviderNamespace string, featureName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"featureName":               autorest.Encode("path", featureName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Features/providers/{resourceProviderNamespace}/features/{featureName}/register", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// RegisterSender sends the Register request. The method will close the
// http.Response Body if it receives an error.
func (client Client) RegisterSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// RegisterResponder handles the response to the Register request. The method always
// closes the http.Response Body.
func (client Client) RegisterResponder(resp *http.Response) (result FeatureResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
