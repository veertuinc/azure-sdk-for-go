//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azkeys

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azkeys/internal"
)

// Client is the struct for interacting with a KeyVault Keys instance
type Client struct {
	kvClient *internal.KeyVaultClient
	vaultUrl string
}

// ClientOptions are the configurable options on a Client.
type ClientOptions struct {
	// Transport sets the transport for making HTTP requests.
	Transport policy.Transporter

	// Retry configures the built-in retry policy behavior.
	Retry policy.RetryOptions

	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry policy.TelemetryOptions

	// Logging configures the built-in logging policy behavior.
	Logging policy.LogOptions

	// PerCallPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCallPolicies []policy.Policy

	// PerRetryPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry request.
	PerTryPolicies []policy.Policy
}

// converts ClientOptions to generated *internal.ConnectionOptions
func (c *ClientOptions) toConnectionOptions() *internal.ConnectionOptions {
	if c == nil {
		return nil
	}

	return &internal.ConnectionOptions{
		HTTPClient:       c.Transport,
		Retry:            c.Retry,
		Telemetry:        c.Telemetry,
		Logging:          c.Logging,
		PerCallPolicies:  c.PerCallPolicies,
		PerRetryPolicies: c.PerTryPolicies,
	}
}

// NewClient returns a pointer to a Client object affinitized to a vaultUrl.
func NewClient(vaultUrl string, credential azcore.TokenCredential, options *ClientOptions) (*Client, error) {
	if options == nil {
		options = &ClientOptions{}
	}

	conn := internal.NewConnection(credential, options.toConnectionOptions())

	return &Client{
		kvClient: &internal.KeyVaultClient{
			Con: conn,
		},
		vaultUrl: vaultUrl,
	}, nil
}

// CreateKeyOptions contains the optional parameters for the KeyVaultClient.CreateKey method.
type CreateKeyOptions struct {
	// Elliptic curve name. For valid values, see JsonWebKeyCurveName.
	Curve *JSONWebKeyCurveName `json:"crv,omitempty"`

	// The attributes of a key managed by the key vault service.
	KeyAttributes *KeyAttributes         `json:"attributes,omitempty"`
	KeyOps        []*JSONWebKeyOperation `json:"key_ops,omitempty"`

	// The key size in bits. For example: 2048, 3072, or 4096 for RSA.
	KeySize *int32 `json:"key_size,omitempty"`

	// The public exponent for a RSA key.
	PublicExponent *int32 `json:"public_exponent,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

// convert CreateKeyOptions to *internal.KeyVaultClientCreateKeyOptions
func (c *CreateKeyOptions) toGenerated() *internal.KeyVaultClientCreateKeyOptions {
	return &internal.KeyVaultClientCreateKeyOptions{}
}

// convert CreateKeyOptions to internal.KeyCreateParameters
func (c *CreateKeyOptions) toKeyCreateParameters(keyType KeyType) internal.KeyCreateParameters {
	var attribs *internal.KeyAttributes
	if c.KeyAttributes != nil {
		attribs = c.KeyAttributes.toGenerated()
	}

	ops := make([]*internal.JSONWebKeyOperation, 0)
	for _, o := range c.KeyOps {
		ops = append(ops, (*internal.JSONWebKeyOperation)(o))
	}

	return internal.KeyCreateParameters{
		Kty:            keyType.toGenerated(),
		Curve:          (*internal.JSONWebKeyCurveName)(c.Curve),
		KeyAttributes:  attribs,
		KeyOps:         ops,
		KeySize:        c.KeySize,
		PublicExponent: c.PublicExponent,
		Tags:           c.Tags,
	}
}

// KeyVaultClientCreateKeyResponse contains the response from method KeyVaultClient.CreateKey.
type CreateKeyResponse struct {
	KeyBundle
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// creates CreateKeyResponse from internal.KeyVaultClient.CreateKeyResponse
func createKeyResponseFromGenerated(g internal.KeyVaultClientCreateKeyResponse) CreateKeyResponse {
	return CreateKeyResponse{
		RawResponse: g.RawResponse,
		KeyBundle: KeyBundle{
			Attributes: keyAttributesFromGenerated(g.Attributes),
			Key:        jsonWebKeyFromGenerated(g.Key),
			Tags:       g.Tags,
			Managed:    g.Managed,
		},
	}
}

// CreateKey - The create key operation can be used to create any key type in Azure Key Vault.
// If the named key already exists, Azure Key Vault creates
// a new version of the key. It requires the keys/create  permission.
func (c *Client) CreateKey(ctx context.Context, name string, keyType KeyType, options *CreateKeyOptions) (CreateKeyResponse, error) {
	if options == nil {
		options = &CreateKeyOptions{}
	}

	resp, err := c.kvClient.CreateKey(ctx, c.vaultUrl, name, options.toKeyCreateParameters(keyType), options.toGenerated())
	if err != nil {
		return CreateKeyResponse{}, err
	}

	return createKeyResponseFromGenerated(resp), nil
}

// CreateECKeyOptions contains the optional parameters for the KeyVaultClient.CreateECKey method
type CreateECKeyOptions struct {
	// Elliptic curve name. For valid values, see JsonWebKeyCurveName.
	CurveName *JSONWebKeyCurveName `json:"crv,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// Whether to create an EC key with HSM protection
	HardwareProtected bool
}

// convert CreateECKeyOptions to internal.KeyCreateParameters
func (c *CreateECKeyOptions) toKeyCreateParameters(keyType KeyType) internal.KeyCreateParameters {
	return internal.KeyCreateParameters{
		Kty:   keyType.toGenerated(),
		Curve: (*internal.JSONWebKeyCurveName)(c.CurveName),
		Tags:  c.Tags,
	}
}

// CreateECKeyResponse contains the response from method Client.CreateECKey.
type CreateECKeyResponse struct {
	KeyBundle
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// convert the internal.KeyVaultClientCreateKeyResponse to CreateECKeyResponse
func createECKeyResponseFromGenerated(g internal.KeyVaultClientCreateKeyResponse) CreateECKeyResponse {
	return CreateECKeyResponse{
		RawResponse: g.RawResponse,
		KeyBundle: KeyBundle{
			Attributes: keyAttributesFromGenerated(g.Attributes),
			Key:        jsonWebKeyFromGenerated(g.Key),
			Tags:       g.Tags,
			Managed:    g.Managed,
		},
	}
}

// CreateKey - The create key operation can be used to create a new elliptic key curve in Azure Key Vault.
// If the named key already exists, Azure Key Vault creates
// a new version of the key. It requires the keys/create  permission.
func (c *Client) CreateECKey(ctx context.Context, name string, options *CreateECKeyOptions) (CreateECKeyResponse, error) {
	keyType := EC

	if options != nil && options.HardwareProtected {
		keyType = ECHSM
	} else if options == nil {
		options = &CreateECKeyOptions{}
	}

	resp, err := c.kvClient.CreateKey(ctx, c.vaultUrl, name, options.toKeyCreateParameters(keyType), &internal.KeyVaultClientCreateKeyOptions{})
	if err != nil {
		return CreateECKeyResponse{}, nil
	}

	return createECKeyResponseFromGenerated(resp), nil
}

// CreateOCTKeyOptions contains the optional parameters for the Client.CreateOCTKey method
type CreateOCTKeyOptions struct {
	// Hardware Protected OCT Key
	HardwareProtected bool

	// The key size in bits. For example: 2048, 3072, or 4096 for RSA.
	KeySize *int32 `json:"key_size,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

// conver the CreateOCTKeyOptions to internal.KeyCreateParameters
func (c *CreateOCTKeyOptions) toKeyCreateParameters(keyType KeyType) internal.KeyCreateParameters {
	return internal.KeyCreateParameters{
		Kty:     keyType.toGenerated(),
		KeySize: c.KeySize,
		Tags:    c.Tags,
	}
}

// CreateOCTKeyResponse contains the response from method Client.CreateOCTKey.
type CreateOCTKeyResponse struct {
	KeyBundle
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// convert generated response to CreateOCTKeyResponse
func createOCTKeyResponseFromGenerated(i internal.KeyVaultClientCreateKeyResponse) CreateOCTKeyResponse {
	return CreateOCTKeyResponse{
		RawResponse: i.RawResponse,
		KeyBundle: KeyBundle{
			Attributes: keyAttributesFromGenerated(i.Attributes),
			Key:        jsonWebKeyFromGenerated(i.Key),
			Tags:       i.Tags,
			Managed:    i.Managed,
		},
	}
}

// CreateKey - The create key operation can be used to create a new octet sequence (symmetric) key in Azure Key Vault.
// If the named key already exists, Azure Key Vault creates
// a new version of the key. It requires the keys/create  permission.
func (c *Client) CreateOCTKey(ctx context.Context, name string, options *CreateOCTKeyOptions) (CreateOCTKeyResponse, error) {
	keyType := Oct

	if options != nil && options.HardwareProtected {
		keyType = OctHSM
	} else if options == nil {
		options = &CreateOCTKeyOptions{}
	}

	resp, err := c.kvClient.CreateKey(ctx, c.vaultUrl, name, options.toKeyCreateParameters(keyType), &internal.KeyVaultClientCreateKeyOptions{})

	return createOCTKeyResponseFromGenerated(resp), err
}

// CreateRSAKeyOptions contains the optional parameters for the Client.CreateRSAKey method.
type CreateRSAKeyOptions struct {
	// Hardware Protected OCT Key
	HardwareProtected bool

	// The key size in bits. For example: 2048, 3072, or 4096 for RSA.
	KeySize *int32 `json:"key_size,omitempty"`

	// The public exponent for a RSA key.
	PublicExponent *int32 `json:"public_exponent,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

// convert CreateRSAKeyOptions to internal.KeyCreateParameters
func (c CreateRSAKeyOptions) toKeyCreateParameters(k KeyType) internal.KeyCreateParameters {
	return internal.KeyCreateParameters{
		Kty:            k.toGenerated(),
		KeySize:        c.KeySize,
		PublicExponent: c.PublicExponent,
		Tags:           c.Tags,
	}
}

// CreateRSAKeyResponse contains the response from method Client.CreateRSAKey.
type CreateRSAKeyResponse struct {
	KeyBundle
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// convert internal response to CreateRSAKeyResponse
func createRSAKeyResponseFromGenerated(i internal.KeyVaultClientCreateKeyResponse) CreateRSAKeyResponse {
	return CreateRSAKeyResponse{
		RawResponse: i.RawResponse,
		KeyBundle: KeyBundle{
			Attributes: keyAttributesFromGenerated(i.Attributes),
			Key:        jsonWebKeyFromGenerated(i.Key),
			Tags:       i.Tags,
			Managed:    i.Managed,
		},
	}
}

// CreateKey - The create key operation can be used to create a new RSA key in Azure Key Vault.
// If the named key already exists, Azure Key Vault creates
// a new version of the key. It requires the keys/create  permission.
func (c *Client) CreateRSAKey(ctx context.Context, name string, options *CreateRSAKeyOptions) (CreateRSAKeyResponse, error) {
	keyType := RSA

	if options != nil && options.HardwareProtected {
		keyType = RSAHSM
	} else if options == nil {
		options = &CreateRSAKeyOptions{}
	}

	resp, err := c.kvClient.CreateKey(ctx, c.vaultUrl, name, options.toKeyCreateParameters(keyType), &internal.KeyVaultClientCreateKeyOptions{})
	if err != nil {
		return CreateRSAKeyResponse{}, err
	}

	return createRSAKeyResponseFromGenerated(resp), nil
}

// ListKeysPager is a Pager for the Client.ListSecrets operation
type ListKeysPager interface {
	// PageResponse returns the current ListKeysPage
	PageResponse() ListKeysPage

	// Err returns true if there is another page of data available, false if not
	Err() error

	// NextPage returns true if there is another page of data available, false if not
	NextPage(context.Context) bool
}

// listKeysPager implements the ListKeysPager interface
type listKeysPager struct {
	genPager internal.KeyVaultClientGetKeysPager
}

// PageResponse returns the results from the page most recently fetched from the service
func (l *listKeysPager) PageResponse() ListKeysPage {
	return listKeysPageFromGenerated(l.genPager.PageResponse())
}

// Err returns an error value if the most recent call to NextPage was not successful, else nil
func (l *listKeysPager) Err() error {
	return l.genPager.Err()
}

// NextPage fetches the next available page of results from the service. If the fetched page
// contains results, the return value is true, else false. Results fetched from the service
// can be evaluated by calling PageResponse on this Pager.
func (l *listKeysPager) NextPage(ctx context.Context) bool {
	return l.genPager.NextPage(ctx)
}

// ListKeysOptions contains the optional parameters for the Client.ListKeys method
type ListKeysOptions struct {
	MaxResults *int32
}

// convert ListKeysOptions to generated options
func (l ListKeysOptions) toGenerated() *internal.KeyVaultClientGetKeysOptions {
	return &internal.KeyVaultClientGetKeysOptions{Maxresults: l.MaxResults}
}

// ListKeysPage contains the current page of results for the Client.ListSecrets operation
type ListKeysPage struct {
	// READ-ONLY; The URL to get the next set of keys.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; A response message containing a list of keys in the key vault along with a link to the next page of keys.
	Keys []*KeyItem `json:"value,omitempty" azure:"ro"`
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// convert internal Response to ListKeysPage
func listKeysPageFromGenerated(i internal.KeyVaultClientGetKeysResponse) ListKeysPage {
	var keys []*KeyItem
	for _, k := range i.Value {
		keys = append(keys, keyItemFromGenerated(k))
	}
	return ListKeysPage{
		RawResponse: i.RawResponse,
		NextLink:    i.NextLink,
		Keys:        keys,
	}
}

// ListKeys retrieves a list of the keys in the Key Vault as JSON Web Key structures that contain the
// public part of a stored key. The LIST operation is applicable to all key types, however only the
// base key identifier, attributes, and tags are provided in the response. Individual versions of a
// key are not listed in the response. This operation requires the keys/list permission.
func (c *Client) ListKeys(options *ListKeysOptions) ListKeysPager {
	if options == nil {
		options = &ListKeysOptions{}
	}
	p := c.kvClient.GetKeys(c.vaultUrl, options.toGenerated())

	return &listKeysPager{
		genPager: *p,
	}
}

// GetKeyOptions contains the options for the Client.GetKey method
type GetKeyOptions struct {
	Version string
}

// GetKeyResponse contains the response for the Client.GetResponse method
type GetKeyResponse struct {
	KeyBundle
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// convert internal response to GetKeyResponse
func getKeyResponseFromGenerated(i internal.KeyVaultClientGetKeyResponse) GetKeyResponse {
	return GetKeyResponse{
		RawResponse: i.RawResponse,
		KeyBundle: KeyBundle{
			Attributes: keyAttributesFromGenerated(i.Attributes),
			Key:        jsonWebKeyFromGenerated(i.Key),
			Tags:       i.Tags,
			Managed:    i.Managed,
		},
	}
}

// GetKey - The get key operation is applicable to all key types. If the requested key is symmetric, then
// no key material is released in the response. This operation requires the keys/get permission.
func (c *Client) GetKey(ctx context.Context, keyName string, options *GetKeyOptions) (GetKeyResponse, error) {
	if options == nil {
		options = &GetKeyOptions{}
	}

	resp, err := c.kvClient.GetKey(ctx, c.vaultUrl, keyName, options.Version, &internal.KeyVaultClientGetKeyOptions{})
	if err != nil {
		return GetKeyResponse{}, err
	}

	return getKeyResponseFromGenerated(resp), err
}

// GetDeletedKeyOptions contains the optional parameters for the Client.GetDeletedKey method
type GetDeletedKeyOptions struct{}

// convert the GetDeletedKeyOptions to the internal representation
func (g GetDeletedKeyOptions) toGenerated() *internal.KeyVaultClientGetDeletedKeyOptions {
	return &internal.KeyVaultClientGetDeletedKeyOptions{}
}

// GetDeletedKeyResponse contains the response from a Client.GetDeletedKey
type GetDeletedKeyResponse struct {
	DeletedKeyBundle
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// convert generated response to GetDeletedKeyResponse
func getDeletedKeyResponseFromGenerated(i internal.KeyVaultClientGetDeletedKeyResponse) GetDeletedKeyResponse {
	return GetDeletedKeyResponse{
		RawResponse: i.RawResponse,
		DeletedKeyBundle: DeletedKeyBundle{
			KeyBundle: KeyBundle{
				Attributes: keyAttributesFromGenerated(i.Attributes),
				Key:        jsonWebKeyFromGenerated(i.Key),
				Tags:       i.Tags,
				Managed:    i.Managed,
			},
			RecoveryID:         i.RecoveryID,
			DeletedDate:        i.DeletedDate,
			ScheduledPurgeDate: i.ScheduledPurgeDate,
		},
	}
}

// GetDeletedKey - The Get Deleted Key operation is applicable for soft-delete enabled vaults.
// While the operation can be invoked on any vault, it will return an error if invoked on a non
// soft-delete enabled vault. This operation requires the keys/get permission.
func (c *Client) GetDeletedKey(ctx context.Context, keyName string, options *GetDeletedKeyOptions) (GetDeletedKeyResponse, error) {
	if options == nil {
		options = &GetDeletedKeyOptions{}
	}

	resp, err := c.kvClient.GetDeletedKey(ctx, c.vaultUrl, keyName, options.toGenerated())
	if err != nil {
		return GetDeletedKeyResponse{}, err
	}

	return getDeletedKeyResponseFromGenerated(resp), nil
}

// PurgeDeletedKeyOptions is the struct for any future options for Client.PurgeDeletedKey.
type PurgeDeletedKeyOptions struct{}

// convert options to internal options
func (p *PurgeDeletedKeyOptions) toGenerated() *internal.KeyVaultClientPurgeDeletedKeyOptions {
	return &internal.KeyVaultClientPurgeDeletedKeyOptions{}
}

// PurgeDeletedKeyResponse contains the response from method Client.PurgeDeletedKey.
type PurgeDeletedKeyResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// Converts the generated response to the publicly exposed version.
func purgeDeletedKeyResponseFromGenerated(i internal.KeyVaultClientPurgeDeletedKeyResponse) PurgeDeletedKeyResponse {
	return PurgeDeletedKeyResponse{
		RawResponse: i.RawResponse,
	}
}

// PurgeDeletedKey deletes the specified key. The purge deleted key operation removes the key permanently, without the possibility of recovery.
// This operation can only be enabled on a soft-delete enabled vault. This operation requires the key/purge permission.
func (c *Client) PurgeDeletedKey(ctx context.Context, keyName string, options *PurgeDeletedKeyOptions) (PurgeDeletedKeyResponse, error) {
	if options == nil {
		options = &PurgeDeletedKeyOptions{}
	}
	resp, err := c.kvClient.PurgeDeletedKey(ctx, c.vaultUrl, keyName, options.toGenerated())
	return purgeDeletedKeyResponseFromGenerated(resp), err
}

// DeletedKeyResponse contains the response for a Client.BeginDeleteKey operation.
type DeleteKeyResponse struct {
	DeletedKeyBundle
	// RawResponse holds the underlying HTTP response
	RawResponse *http.Response
}

// convert interal response to DeleteKeyResponse
func deleteKeyResponseFromGenerated(i *internal.KeyVaultClientDeleteKeyResponse) *DeleteKeyResponse {
	if i == nil {
		return nil
	}
	return &DeleteKeyResponse{
		RawResponse: i.RawResponse,
	}
}

// BeginDeleteKeyOptions contains the optional parameters for the Client.BeginDeleteKey method.
type BeginDeleteKeyOptions struct{}

// convert public options to generated options struct
func (b *BeginDeleteKeyOptions) toGenerated() *internal.KeyVaultClientDeleteKeyOptions {
	return &internal.KeyVaultClientDeleteKeyOptions{}
}

// DeleteKeyPoller is the interface for the Client.DeleteKey operation.
type DeleteKeyPoller interface {
	// Done returns true if the LRO has reached a terminal state
	Done() bool

	// Poll fetches the latest state of the LRO. It returns an HTTP response or error.
	// If the LRO has completed successfully, the poller's state is updated and the HTTP response is returned.
	// If the LRO has completed with failure or was cancelled, the poller's state is updated and the error is returned.
	Poll(context.Context) (*http.Response, error)

	// FinalResponse returns the final response after the operations has finished
	FinalResponse(context.Context) (DeleteKeyResponse, error)
}

// The poller returned by the Client.StartDeleteKey operation
type startDeleteKeyPoller struct {
	keyName        string // This is the key to Poll for in GetDeletedKey
	vaultUrl       string
	client         *internal.KeyVaultClient
	deleteResponse internal.KeyVaultClientDeleteKeyResponse
	lastResponse   internal.KeyVaultClientGetDeletedKeyResponse
	RawResponse    *http.Response
}

// Done returns true if the LRO has reached a terminal state
func (s *startDeleteKeyPoller) Done() bool {
	return s.lastResponse.RawResponse != nil
}

// Poll fetches the latest state of the LRO. It returns an HTTP response or error.(
// If the LRO has completed successfully, the poller's state is updated and the HTTP response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is updated and the error is returned.)
func (s *startDeleteKeyPoller) Poll(ctx context.Context) (*http.Response, error) {
	resp, err := s.client.GetDeletedKey(ctx, s.vaultUrl, s.keyName, nil)
	if err == nil {
		// Service recognizes DeletedKey, operation is done
		s.lastResponse = resp
		return resp.RawResponse, nil
	} else if err != nil {
		return s.deleteResponse.RawResponse, nil
	}
	s.lastResponse = resp
	return resp.RawResponse, nil
}

// FinalResponse returns the final response after the operations has finished
func (s *startDeleteKeyPoller) FinalResponse(ctx context.Context) (DeleteKeyResponse, error) {
	return *deleteKeyResponseFromGenerated(&s.deleteResponse), nil
}

// pollUntilDone continually calls the Poll operation until the operation is completed. In between each
// Poll is a wait determined by the t parameter.
func (s *startDeleteKeyPoller) pollUntilDone(ctx context.Context, t time.Duration) (DeleteKeyResponse, error) {
	for {
		resp, err := s.Poll(ctx)
		if err != nil {
			return DeleteKeyResponse{}, err
		}
		s.RawResponse = resp
		if s.Done() {
			break
		}
		time.Sleep(t)
	}
	return DeleteKeyResponse{}, nil
}

// DeleteKeyPollerResponse contains the response from the Client.BeginDeleteKey method
type DeleteKeyPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error occurs
	PollUntilDone func(context.Context, time.Duration) (DeleteKeyResponse, error)

	// Poller contains an initialized WidgetPoller
	Poller DeleteKeyPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BeginDeleteKey deletes a key from the keyvault. Delete cannot be applied to an individual version of a key. This operation
// requires the key/delete permission. This response contains a Poller struct that can be used to Poll for a response, or the
// response PollUntilDone function can be used to poll until completion.
func (c *Client) BeginDeleteKey(ctx context.Context, keyName string, options *BeginDeleteKeyOptions) (DeleteKeyPollerResponse, error) {
	if options == nil {
		options = &BeginDeleteKeyOptions{}
	}
	resp, err := c.kvClient.DeleteKey(ctx, c.vaultUrl, keyName, options.toGenerated())
	if err != nil {
		return DeleteKeyPollerResponse{}, err
	}

	getResp, err := c.kvClient.GetDeletedKey(ctx, c.vaultUrl, keyName, nil)
	var httpErr azcore.HTTPResponse
	if errors.As(err, &httpErr) {
		if httpErr.RawResponse().StatusCode != http.StatusNotFound {
			return DeleteKeyPollerResponse{}, err
		}
	}

	s := &startDeleteKeyPoller{
		vaultUrl:       c.vaultUrl,
		keyName:        keyName,
		client:         c.kvClient,
		deleteResponse: resp,
		lastResponse:   getResp,
	}

	return DeleteKeyPollerResponse{
		Poller:        s,
		RawResponse:   resp.RawResponse,
		PollUntilDone: s.pollUntilDone,
	}, nil
}

// BackupKeyOptions contains the optional parameters for the Client.BackupKey method
type BackupKeyOptions struct{}

// convert Options to generated version
func (b BackupKeyOptions) toGenerated() *internal.KeyVaultClientBackupKeyOptions {
	return &internal.KeyVaultClientBackupKeyOptions{}
}

// BackupKeyResponse contains the response from the Client.BackupKey method
type BackupKeyResponse struct {
	// READ-ONLY; The backup blob containing the backed up key.
	Value []byte `json:"value,omitempty" azure:"ro"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// convert internal reponse to BackupKeyResponse
func backupKeyResponseFromGenerated(i internal.KeyVaultClientBackupKeyResponse) BackupKeyResponse {
	return BackupKeyResponse{
		RawResponse: i.RawResponse,
		Value:       i.Value,
	}
}

// BackupKey - The Key Backup operation exports a key from Azure Key Vault in a protected form.
// Note that this operation does NOT return key material in a form that can be used outside the
// Azure Key Vault system, the returned key material is either protected to a Azure Key Vault
// HSM or to Azure Key Vault itself. The intent of this operation is to allow a client to GENERATE
// a key in one Azure Key Vault instance, BACKUP the key, and then RESTORE it into another Azure
// Key Vault instance. The BACKUP operation may be used to export, in protected form, any key
// type from Azure Key Vault. Individual versions of a key cannot be backed up. BACKUP / RESTORE
// can be performed within geographical boundaries only; meaning that a BACKUP from one geographical
// area cannot be restored to another geographical area. For example, a backup from the US
// geographical area cannot be restored in an EU geographical area. This operation requires the key/backup permission.
func (c *Client) BackupKey(ctx context.Context, keyName string, options *BackupKeyOptions) (BackupKeyResponse, error) {
	if options == nil {
		options = &BackupKeyOptions{}
	}

	resp, err := c.kvClient.BackupKey(ctx, c.vaultUrl, keyName, options.toGenerated())
	if err != nil {
		return BackupKeyResponse{}, err
	}

	return backupKeyResponseFromGenerated(resp), nil
}

// RecoverDeletedKeyPoller is the interface for the Client.RecoverDeletedKey operation
type RecoverDeletedKeyPoller interface {
	// Done returns true if the LRO has reached a terminal state
	Done() bool

	// Poll fetches the latest state of the LRO. It returns an HTTP response or error.
	// If the LRO has completed successfully, the poller's state is updated and the HTTP response is returned.
	// If the LRO has completed with failure or was cancelled, the poller's state is updated and the error is returned.
	Poll(context.Context) (*http.Response, error)

	// FinalResponse returns the final response after the operations has finished
	FinalResponse(context.Context) (RecoverDeletedKeyResponse, error)
}

// beginRecoverPoller implements the RecoverDeletedKeyPoller interface
type beginRecoverPoller struct {
	keyName         string
	vaultUrl        string
	client          *internal.KeyVaultClient
	recoverResponse internal.KeyVaultClientRecoverDeletedKeyResponse
	lastResponse    internal.KeyVaultClientGetKeyResponse
	RawResponse     *http.Response
}

// Done returns true when the polling operation is completed
func (b *beginRecoverPoller) Done() bool {
	return b.RawResponse.StatusCode == http.StatusOK
}

// Poll fetches the latest state of the LRO. It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is updated and the error is returned.
func (b *beginRecoverPoller) Poll(ctx context.Context) (*http.Response, error) {
	resp, err := b.client.GetKey(ctx, b.vaultUrl, b.keyName, "", nil)
	b.lastResponse = resp
	var httpErr azcore.HTTPResponse
	if errors.As(err, &httpErr) {
		return httpErr.RawResponse(), err
	}
	return resp.RawResponse, nil
}

// FinalResponse returns the final response after the operations has finished
func (b *beginRecoverPoller) FinalResponse(ctx context.Context) (RecoverDeletedKeyResponse, error) {
	return recoverDeletedKeyResponseFromGenerated(b.recoverResponse), nil
}

// pollUntilDone is the method for the Response.PollUntilDone struct
func (b *beginRecoverPoller) pollUntilDone(ctx context.Context, t time.Duration) (RecoverDeletedKeyResponse, error) {
	for {
		resp, err := b.Poll(ctx)
		if err != nil {
			b.RawResponse = resp
		}
		if b.Done() {
			break
		}
		b.RawResponse = resp
		time.Sleep(t)
	}
	return recoverDeletedKeyResponseFromGenerated(b.recoverResponse), nil
}

// BeginRecoverDeletedKeyOptions contains the optional parameters for the Client.BeginRecoverDeletedKey operation
type BeginRecoverDeletedKeyOptions struct{}

// Convert the publicly exposed options object to the generated version
func (b BeginRecoverDeletedKeyOptions) toGenerated() *internal.KeyVaultClientRecoverDeletedKeyOptions {
	return &internal.KeyVaultClientRecoverDeletedKeyOptions{}
}

// RecoverDeletedKeyResponse is the response object for the Client.RecoverDeletedKey operation.
type RecoverDeletedKeyResponse struct {
	KeyBundle
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// change recover deleted key reponse to the generated version.
func recoverDeletedKeyResponseFromGenerated(i internal.KeyVaultClientRecoverDeletedKeyResponse) RecoverDeletedKeyResponse {
	return RecoverDeletedKeyResponse{
		RawResponse: i.RawResponse,
		KeyBundle: KeyBundle{
			Attributes: keyAttributesFromGenerated(i.Attributes),
			Key:        jsonWebKeyFromGenerated(i.Key),
			Tags:       i.Tags,
			Managed:    i.Managed,
		},
	}
}

// RecoverDeletedKeyPollerResponse contains the response of the Client.BeginRecoverDeletedKey operations
type RecoverDeletedKeyPollerResponse struct {
	// PollUntilDone will poll the service endpoint until a terminal state is reached or an error occurs
	PollUntilDone func(context.Context, time.Duration) (RecoverDeletedKeyResponse, error)

	// Poller contains an initialized RecoverDeletedKeyPoller
	Poller RecoverDeletedKeyPoller

	// RawResponse cotains the underlying HTTP response
	RawResponse *http.Response
}

// BeginRecoverDeletedKey recovers the deleted key in the specified vault to the latest version.
// This operation can only be performed on a soft-delete enabled vault. This operation requires the keys/recover permission.
func (c *Client) BeginRecoverDeletedKey(ctx context.Context, keyName string, options *BeginRecoverDeletedKeyOptions) (RecoverDeletedKeyPollerResponse, error) {
	if options == nil {
		options = &BeginRecoverDeletedKeyOptions{}
	}
	resp, err := c.kvClient.RecoverDeletedKey(ctx, c.vaultUrl, keyName, options.toGenerated())
	if err != nil {
		return RecoverDeletedKeyPollerResponse{}, err
	}

	getResp, err := c.kvClient.GetKey(ctx, c.vaultUrl, keyName, "", nil)
	var httpErr azcore.HTTPResponse
	if errors.As(err, &httpErr) {
		if httpErr.RawResponse().StatusCode != http.StatusNotFound {
			return RecoverDeletedKeyPollerResponse{}, err
		}
	}

	b := &beginRecoverPoller{
		lastResponse:    getResp,
		keyName:         keyName,
		client:          c.kvClient,
		vaultUrl:        c.vaultUrl,
		recoverResponse: resp,
		RawResponse:     getResp.RawResponse,
	}

	return RecoverDeletedKeyPollerResponse{
		PollUntilDone: b.pollUntilDone,
		Poller:        b,
		RawResponse:   getResp.RawResponse,
	}, nil
}

// UpdateKeyPropertiesOptions contains the optional parameters for the Client.UpdateKeyProperties method
type UpdateKeyPropertiesOptions struct {
	// The version of a key to update
	Version string

	// The attributes of a key managed by the key vault service.
	KeyAttributes *KeyAttributes `json:"attributes,omitempty"`

	// Json web key operations. For more information on possible key operations, see JsonWebKeyOperation.
	KeyOps []*JSONWebKeyOperation `json:"key_ops,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

// convert the options to internal.KeyUpdateParameters struct
func (u UpdateKeyPropertiesOptions) toKeyUpdateParameters() internal.KeyUpdateParameters {
	var attribs *internal.KeyAttributes
	if u.KeyAttributes != nil {
		attribs = u.KeyAttributes.toGenerated()
	}

	ops := make([]*internal.JSONWebKeyOperation, 0)
	for _, o := range u.KeyOps {
		ops = append(ops, (*internal.JSONWebKeyOperation)(o))
	}

	return internal.KeyUpdateParameters{
		KeyOps:        ops,
		KeyAttributes: attribs,
		Tags:          u.Tags,
	}
}

// convert options to generated options
func (u UpdateKeyPropertiesOptions) toGeneratedOptions() *internal.KeyVaultClientUpdateKeyOptions {
	return &internal.KeyVaultClientUpdateKeyOptions{}
}

// UpdateKeyPropertiesResponse contains the response for the Client.UpdateKeyProperties method
type UpdateKeyPropertiesResponse struct {
	KeyBundle
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// convert the internal response to UpdateKeyPropertiesResponse
func updateKeyPropertiesFromGenerated(i internal.KeyVaultClientUpdateKeyResponse) UpdateKeyPropertiesResponse {
	return UpdateKeyPropertiesResponse{
		RawResponse: i.RawResponse,
		KeyBundle: KeyBundle{
			Attributes: keyAttributesFromGenerated(i.Attributes),
			Key:        jsonWebKeyFromGenerated(i.Key),
			Tags:       i.Tags,
			Managed:    i.Managed,
		},
	}
}

// UpdateKey - In order to perform this operation, the key must already exist in the Key Vault.
// Note: The cryptographic material of a key itself cannot be changed. This operation requires the keys/update permission.
func (c *Client) UpdateKeyProperties(ctx context.Context, keyName string, options *UpdateKeyPropertiesOptions) (UpdateKeyPropertiesResponse, error) {
	if options == nil {
		options = &UpdateKeyPropertiesOptions{}
	}
	resp, err := c.kvClient.UpdateKey(
		ctx,
		c.vaultUrl,
		keyName,
		options.Version,
		options.toKeyUpdateParameters(),
		options.toGeneratedOptions(),
	)
	if err != nil {
		return UpdateKeyPropertiesResponse{}, err
	}

	return updateKeyPropertiesFromGenerated(resp), nil
}

// ListDeletedKeys is the interface for the Client.ListDeletedKeys operation
type ListDeletedKeysPager interface {
	// PageResponse returns the current ListDeletedKeysPage
	PageResponse() ListDeletedKeysPage

	// Err returns true if there is another page of data available, false if not
	Err() error

	// NextPage returns true if there is another page of data available, false if not
	NextPage(context.Context) bool
}

// listDeletedKeysPager is the pager returned by Client.ListDeletedKeys
type listDeletedKeysPager struct {
	genPager *internal.KeyVaultClientGetDeletedKeysPager
}

// PageResponse returns the current page of results
func (l *listDeletedKeysPager) PageResponse() ListDeletedKeysPage {
	resp := l.genPager.PageResponse()

	var values []*DeletedKeyItem
	for _, d := range resp.Value {
		values = append(values, deletedKeyItemFromGenerated(d))
	}

	return ListDeletedKeysPage{
		RawResponse: resp.RawResponse,
		NextLink:    resp.NextLink,
		DeletedKeys: values,
	}
}

// Err returns an error if the last operation resulted in an error.
func (l *listDeletedKeysPager) Err() error {
	return l.genPager.Err()
}

// NextPage fetches the next page of results.
func (l *listDeletedKeysPager) NextPage(ctx context.Context) bool {
	return l.genPager.NextPage(ctx)
}

// ListDeletedKeysPage holds the data for a single page.
type ListDeletedKeysPage struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// READ-ONLY; The URL to get the next set of deleted keys.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; A response message containing a list of the deleted keys in the vault along with a link to the next page of deleted keys
	DeletedKeys []*DeletedKeyItem `json:"value,omitempty" azure:"ro"`
}

// ListDeletedKeysOptions contains the optional parameters for the Client.ListDeletedKeys operation.
type ListDeletedKeysOptions struct {
	// Maximum number of results to return in a page. If not specified the service will return up to 25 results.
	MaxResults *int32
}

// Convert publicly exposed options to the generated version.a
func (l *ListDeletedKeysOptions) toGenerated() *internal.KeyVaultClientGetDeletedKeysOptions {
	return &internal.KeyVaultClientGetDeletedKeysOptions{
		Maxresults: l.MaxResults,
	}
}

// ListDeletedKeys lists all versions of the specified key. The full key identifier and attributes are provided
// in the response. No values are returned for the keys. This operation requires the keys/list permission.
func (c *Client) ListDeletedKeys(options *ListDeletedKeysOptions) ListDeletedKeysPager {
	if options == nil {
		options = &ListDeletedKeysOptions{}
	}

	return &listDeletedKeysPager{
		genPager: c.kvClient.GetDeletedKeys(c.vaultUrl, options.toGenerated()),
	}
}

// ListKeyVersionsPager is a Pager for Client.ListKeyVersions results
type ListKeyVersionsPager interface {
	// PageResponse returns the current ListKeyVersionsPage
	PageResponse() ListKeyVersionsPage

	// Err returns true if there is another page of data available, false if not
	Err() error

	// NextPage returns true if there is another page of data available, false if not
	NextPage(context.Context) bool
}

// listKeyVersionsPager implements the ListKeyVersionsPager interface
type listKeyVersionsPager struct {
	genPager *internal.KeyVaultClientGetKeyVersionsPager
}

// PageResponse returns the results from the page most recently fetched from the service.
func (l *listKeyVersionsPager) PageResponse() ListKeyVersionsPage {
	return listKeyVersionsPageFromGenerated(l.genPager.PageResponse())
}

// Err returns an error value if the most recent call to NextPage was not successful, else nil.
func (l *listKeyVersionsPager) Err() error {
	return l.genPager.Err()
}

// NextPage fetches the next available page of results from the service. If the fetched page
// contains results, the return value is true, else false. Results fetched from the service
// can be evaluated by calling PageResponse on this Pager.
func (l *listKeyVersionsPager) NextPage(ctx context.Context) bool {
	return l.genPager.NextPage(ctx)
}

// ListKeyVersionsOptions contains the options for the ListKeyVersions operations
type ListKeyVersionsOptions struct {
	// Maximum number of results to return in a page. If not specified the service will return up to 25 results.
	MaxResults *int32
}

// convert the public ListKeyVersionsOptions to the generated version
func (l *ListKeyVersionsOptions) toGenerated() *internal.KeyVaultClientGetKeyVersionsOptions {
	if l == nil {
		return &internal.KeyVaultClientGetKeyVersionsOptions{}
	}
	return &internal.KeyVaultClientGetKeyVersionsOptions{
		Maxresults: l.MaxResults,
	}
}

// ListKeyVersionsPage contains the current page from a ListKeyVersionsPager.PageResponse method
type ListKeyVersionsPage struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// READ-ONLY; The URL to get the next set of keys.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; A response message containing a list of keys in the key vault along with a link to the next page of keys.
	Keys []KeyItem `json:"value,omitempty" azure:"ro"`
}

// create ListKeysPage from generated pager
func listKeyVersionsPageFromGenerated(i internal.KeyVaultClientGetKeyVersionsResponse) ListKeyVersionsPage {
	var keys []KeyItem
	for _, s := range i.Value {
		if s != nil {
			keys = append(keys, *keyItemFromGenerated(s))
		}
	}
	return ListKeyVersionsPage{
		RawResponse: i.RawResponse,
		NextLink:    i.NextLink,
		Keys:        keys,
	}
}

// ListKeyVersions lists all versions of the specified key. The full key identifer and
// attributes are provided in the response. No values are returned for the keys. This operation
// requires the keys/list permission.
func (c *Client) ListKeyVersions(keyName string, options *ListKeyVersionsOptions) ListKeyVersionsPager {
	if options == nil {
		options = &ListKeyVersionsOptions{}
	}

	return &listKeyVersionsPager{
		genPager: c.kvClient.GetKeyVersions(
			c.vaultUrl,
			keyName,
			options.toGenerated(),
		),
	}
}

// RestoreKeyBackupOptions contains the optional parameters for the Client.RestoreKey method.
type RestoreKeyBackupOptions struct{}

func (r RestoreKeyBackupOptions) toGenerated() *internal.KeyVaultClientRestoreKeyOptions {
	return &internal.KeyVaultClientRestoreKeyOptions{}
}

// RestoreKeyBackupResponse contains the response object for the Client.RestoreKeyBackup operation.
type RestoreKeyBackupResponse struct {
	KeyBundle
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// converts the generated response to the publicly exposed version.
func restoreKeyBackupResponseFromGenerated(i internal.KeyVaultClientRestoreKeyResponse) RestoreKeyBackupResponse {
	return RestoreKeyBackupResponse{
		RawResponse: i.RawResponse,
		KeyBundle: KeyBundle{
			Attributes: keyAttributesFromGenerated(i.Attributes),
			Key:        jsonWebKeyFromGenerated(i.Key),
			Tags:       i.Tags,
			Managed:    i.Managed,
		},
	}
}

// RestoreKeyBackup restores a backed up key, and all its versions, to a vault. This operation requires the keys/restore permission.
// The backup parameter is a blob of the key to restore, this can be received from the Client.BackupKey function.
func (c *Client) RestoreKeyBackup(ctx context.Context, keyBackup []byte, options *RestoreKeyBackupOptions) (RestoreKeyBackupResponse, error) {
	if options == nil {
		options = &RestoreKeyBackupOptions{}
	}

	resp, err := c.kvClient.RestoreKey(ctx, c.vaultUrl, internal.KeyRestoreParameters{KeyBundleBackup: keyBackup}, options.toGenerated())
	if err != nil {
		return RestoreKeyBackupResponse{}, err
	}

	return restoreKeyBackupResponseFromGenerated(resp), nil
}

// ImportKeyOptions contains the optional parameters for the Client.ImportKeyOptions parameter
type ImportKeyOptions struct {
	// Whether to import as a hardware key (HSM) or software key.
	Hsm *bool `json:"Hsm,omitempty"`

	// The key management attributes.
	KeyAttributes *KeyAttributes `json:"attributes,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`
}

func (i ImportKeyOptions) toImportKeyParameters(key JSONWebKey) internal.KeyImportParameters {
	var attribs *internal.KeyAttributes
	if i.KeyAttributes != nil {
		attribs = i.KeyAttributes.toGenerated()
	}
	return internal.KeyImportParameters{
		Key:           key.toGenerated(),
		Hsm:           i.Hsm,
		KeyAttributes: attribs,
		Tags:          i.Tags,
	}
}

// ImportKeyResponse contains the response of the Client.ImportKey method
type ImportKeyResponse struct {
	KeyBundle
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// convert the generated response to the ImportKeyResponse
func importKeyResponseFromGenerated(i internal.KeyVaultClientImportKeyResponse) ImportKeyResponse {
	return ImportKeyResponse{
		RawResponse: i.RawResponse,
		KeyBundle: KeyBundle{
			Attributes: keyAttributesFromGenerated(i.Attributes),
			Key:        jsonWebKeyFromGenerated(i.Key),
			Tags:       i.Tags,
			Managed:    i.Managed,
		},
	}
}

// ImportKey - The import key operation may be used to import any key type into an Azure Key Vault.
// If the named key already exists, Azure Key Vault creates a new version of the key. This operation
// requires the  keys/import permission.
func (c *Client) ImportKey(ctx context.Context, keyName string, key JSONWebKey, options *ImportKeyOptions) (ImportKeyResponse, error) {
	if options == nil {
		options = &ImportKeyOptions{}
	}

	resp, err := c.kvClient.ImportKey(ctx, c.vaultUrl, keyName, options.toImportKeyParameters(key), &internal.KeyVaultClientImportKeyOptions{})
	if err != nil {
		return ImportKeyResponse{}, err
	}

	return importKeyResponseFromGenerated(resp), nil
}
