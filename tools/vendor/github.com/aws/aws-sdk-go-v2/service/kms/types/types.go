// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Contains information about an alias.
type AliasListEntry struct {

	// String that contains the key ARN.
	AliasArn *string

	// String that contains the alias. This value begins with alias/.
	AliasName *string

	// Date and time that the alias was most recently created in the account and
	// Region. Formatted as Unix time.
	CreationDate *time.Time

	// Date and time that the alias was most recently associated with a KMS key in the
	// account and Region. Formatted as Unix time.
	LastUpdatedDate *time.Time

	// String that contains the key identifier of the KMS key associated with the
	// alias.
	TargetKeyId *string

	noSmithyDocumentSerde
}

// Contains information about each custom key store in the custom key store list.
type CustomKeyStoresListEntry struct {

	// A unique identifier for the CloudHSM cluster that is associated with an CloudHSM
	// key store. This field appears only when the CustomKeyStoreType is AWS_CLOUDHSM.
	CloudHsmClusterId *string

	// Describes the connection error. This field appears in the response only when the
	// ConnectionState is FAILED. Many failures can be resolved by updating the
	// properties of the custom key store. To update a custom key store, disconnect it
	// (DisconnectCustomKeyStore), correct the errors (UpdateCustomKeyStore), and try
	// to connect again (ConnectCustomKeyStore). For additional help resolving these
	// errors, see How to Fix a Connection Failure
	// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-failed)
	// in Key Management Service Developer Guide. All custom key stores:
	//
	// *
	// INTERNAL_ERROR — KMS could not complete the request due to an internal error.
	// Retry the request. For ConnectCustomKeyStore requests, disconnect the custom key
	// store before trying to connect again.
	//
	// * NETWORK_ERRORS — Network errors are
	// preventing KMS from connecting the custom key store to its backing key
	// store.
	//
	// CloudHSM key stores:
	//
	// * CLUSTER_NOT_FOUND — KMS cannot find the CloudHSM
	// cluster with the specified cluster ID.
	//
	// * INSUFFICIENT_CLOUDHSM_HSMS — The
	// associated CloudHSM cluster does not contain any active HSMs. To connect a
	// custom key store to its CloudHSM cluster, the cluster must contain at least one
	// active HSM.
	//
	// * INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET — At least one private
	// subnet associated with the CloudHSM cluster doesn't have any available IP
	// addresses. A CloudHSM key store connection requires one free IP address in each
	// of the associated private subnets, although two are preferable. For details, see
	// How to Fix a Connection Failure
	// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-failed)
	// in the Key Management Service Developer Guide.
	//
	// * INVALID_CREDENTIALS — The
	// KeyStorePassword for the custom key store doesn't match the current password of
	// the kmsuser crypto user in the CloudHSM cluster. Before you can connect your
	// custom key store to its CloudHSM cluster, you must change the kmsuser account
	// password and update the KeyStorePassword value for the custom key store.
	//
	// *
	// SUBNET_NOT_FOUND — A subnet in the CloudHSM cluster configuration was deleted.
	// If KMS cannot find all of the subnets in the cluster configuration, attempts to
	// connect the custom key store to the CloudHSM cluster fail. To fix this error,
	// create a cluster from a recent backup and associate it with your custom key
	// store. (This process creates a new cluster configuration with a VPC and private
	// subnets.) For details, see How to Fix a Connection Failure
	// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-failed)
	// in the Key Management Service Developer Guide.
	//
	// * USER_LOCKED_OUT — The kmsuser
	// CU account is locked out of the associated CloudHSM cluster due to too many
	// failed password attempts. Before you can connect your custom key store to its
	// CloudHSM cluster, you must change the kmsuser account password and update the
	// key store password value for the custom key store.
	//
	// * USER_LOGGED_IN — The
	// kmsuser CU account is logged into the associated CloudHSM cluster. This prevents
	// KMS from rotating the kmsuser account password and logging into the cluster.
	// Before you can connect your custom key store to its CloudHSM cluster, you must
	// log the kmsuser CU out of the cluster. If you changed the kmsuser password to
	// log into the cluster, you must also and update the key store password value for
	// the custom key store. For help, see How to Log Out and Reconnect
	// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#login-kmsuser-2)
	// in the Key Management Service Developer Guide.
	//
	// * USER_NOT_FOUND — KMS cannot
	// find a kmsuser CU account in the associated CloudHSM cluster. Before you can
	// connect your custom key store to its CloudHSM cluster, you must create a kmsuser
	// CU account in the cluster, and then update the key store password value for the
	// custom key store.
	//
	// External key stores:
	//
	// * INVALID_CREDENTIALS — One or both of
	// the XksProxyAuthenticationCredential values is not valid on the specified
	// external key store proxy.
	//
	// * XKS_PROXY_ACCESS_DENIED — KMS requests are denied
	// access to the external key store proxy. If the external key store proxy has
	// authorization rules, verify that they permit KMS to communicate with the proxy
	// on your behalf.
	//
	// * XKS_PROXY_INVALID_CONFIGURATION — A configuration error is
	// preventing the external key store from connecting to its proxy. Verify the value
	// of the XksProxyUriPath.
	//
	// * XKS_PROXY_INVALID_RESPONSE — KMS cannot interpret the
	// response from the external key store proxy. If you see this connection error
	// code repeatedly, notify your external key store proxy vendor.
	//
	// *
	// XKS_PROXY_INVALID_TLS_CONFIGURATION — KMS cannot connect to the external key
	// store proxy because the TLS configuration is invalid. Verify that the XKS proxy
	// supports TLS 1.2 or 1.3. Also, verify that the TLS certificate is not expired,
	// and that it matches the hostname in the XksProxyUriEndpoint value, and that it
	// is signed by a certificate authority included in the Trusted Certificate
	// Authorities
	// (https://github.com/aws/aws-kms-xksproxy-api-spec/blob/main/TrustedCertificateAuthorities)
	// list.
	//
	// * XKS_PROXY_NOT_REACHABLE — KMS can't communicate with your external key
	// store proxy. Verify that the XksProxyUriEndpoint and XksProxyUriPath are
	// correct. Use the tools for your external key store proxy to verify that the
	// proxy is active and available on its network. Also, verify that your external
	// key manager instances are operating properly. Connection attempts fail with this
	// connection error code if the proxy reports that all external key manager
	// instances are unavailable.
	//
	// * XKS_PROXY_TIMED_OUT — KMS can connect to the
	// external key store proxy, but the proxy does not respond to KMS in the time
	// allotted. If you see this connection error code repeatedly, notify your external
	// key store proxy vendor.
	//
	// * XKS_VPC_ENDPOINT_SERVICE_INVALID_CONFIGURATION — The
	// Amazon VPC endpoint service configuration doesn't conform to the requirements
	// for an KMS external key store.
	//
	// * The VPC endpoint service must be an endpoint
	// service for interface endpoints in the caller's Amazon Web Services account.
	//
	// *
	// It must have a network load balancer (NLB) connected to at least two subnets,
	// each in a different Availability Zone.
	//
	// * The Allow principals list must include
	// the KMS service principal for the Region, cks.kms..amazonaws.com,
	//
	// such as
	// cks.kms.us-east-1.amazonaws.com.
	//
	// * It must not require acceptance
	// (https://docs.aws.amazon.com/vpc/latest/privatelink/create-endpoint-service.html)
	// of connection requests.
	//
	// * It must have a private DNS name. The private DNS name
	// for an external key store with VPC_ENDPOINT_SERVICE connectivity must be unique
	// in its Amazon Web Services Region.
	//
	// * The domain of the private DNS name must
	// have a verification status
	// (https://docs.aws.amazon.com/vpc/latest/privatelink/verify-domains.html) of
	// verified.
	//
	// * The TLS certificate
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/network/create-tls-listener.html)
	// specifies the private DNS hostname at which the endpoint is reachable.
	//
	// *
	// XKS_VPC_ENDPOINT_SERVICE_NOT_FOUND — KMS can't find the VPC endpoint service
	// that it uses to communicate with the external key store proxy. Verify that the
	// XksProxyVpcEndpointServiceName is correct and the KMS service principal has
	// service consumer permissions on the Amazon VPC endpoint service.
	ConnectionErrorCode ConnectionErrorCodeType

	// Indicates whether the custom key store is connected to its backing key store.
	// For an CloudHSM key store, the ConnectionState indicates whether it is connected
	// to its CloudHSM cluster. For an external key store, the ConnectionState
	// indicates whether it is connected to the external key store proxy that
	// communicates with your external key manager. You can create and use KMS keys in
	// your custom key stores only when its ConnectionState is CONNECTED. The
	// ConnectionState value is DISCONNECTED only if the key store has never been
	// connected or you use the DisconnectCustomKeyStore operation to disconnect it. If
	// the value is CONNECTED but you are having trouble using the custom key store,
	// make sure that the backing key store is reachable and active. For an CloudHSM
	// key store, verify that its associated CloudHSM cluster is active and contains at
	// least one active HSM. For an external key store, verify that the external key
	// store proxy and external key manager are connected and enabled. A value of
	// FAILED indicates that an attempt to connect was unsuccessful. The
	// ConnectionErrorCode field in the response indicates the cause of the failure.
	// For help resolving a connection failure, see Troubleshooting a custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html) in the
	// Key Management Service Developer Guide.
	ConnectionState ConnectionStateType

	// The date and time when the custom key store was created.
	CreationDate *time.Time

	// A unique identifier for the custom key store.
	CustomKeyStoreId *string

	// The user-specified friendly name for the custom key store.
	CustomKeyStoreName *string

	// Indicates the type of the custom key store. AWS_CLOUDHSM indicates a custom key
	// store backed by an CloudHSM cluster. EXTERNAL_KEY_STORE indicates a custom key
	// store backed by an external key store proxy and external key manager outside of
	// Amazon Web Services.
	CustomKeyStoreType CustomKeyStoreType

	// The trust anchor certificate of the CloudHSM cluster associated with an CloudHSM
	// key store. When you initialize the cluster
	// (https://docs.aws.amazon.com/cloudhsm/latest/userguide/initialize-cluster.html#sign-csr),
	// you create this certificate and save it in the customerCA.crt file. This field
	// appears only when the CustomKeyStoreType is AWS_CLOUDHSM.
	TrustAnchorCertificate *string

	// Configuration settings for the external key store proxy (XKS proxy). The
	// external key store proxy translates KMS requests into a format that your
	// external key manager can understand. The proxy configuration includes connection
	// information that KMS requires. This field appears only when the
	// CustomKeyStoreType is EXTERNAL_KEY_STORE.
	XksProxyConfiguration *XksProxyConfigurationType

	noSmithyDocumentSerde
}

// Use this structure to allow cryptographic operations
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
// in the grant only when the operation request includes the specified encryption
// context
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context).
// KMS applies the grant constraints only to cryptographic operations that support
// an encryption context, that is, all cryptographic operations with a symmetric
// KMS key
// (https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-concepts.html#symmetric-cmks).
// Grant constraints are not applied to operations that do not support an
// encryption context, such as cryptographic operations with asymmetric KMS keys
// and management operations, such as DescribeKey or RetireGrant. In a
// cryptographic operation, the encryption context in the decryption operation must
// be an exact, case-sensitive match for the keys and values in the encryption
// context of the encryption operation. Only the order of the pairs can vary.
// However, in a grant constraint, the key in each key-value pair is not case
// sensitive, but the value is case sensitive. To avoid confusion, do not use
// multiple encryption context pairs that differ only by case. To require a fully
// case-sensitive encryption context, use the kms:EncryptionContext: and
// kms:EncryptionContextKeys conditions in an IAM or key policy. For details, see
// kms:EncryptionContext:
// (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-context)
// in the Key Management Service Developer Guide .
type GrantConstraints struct {

	// A list of key-value pairs that must match the encryption context in the
	// cryptographic operation
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// request. The grant allows the operation only when the encryption context in the
	// request is the same as the encryption context specified in this constraint.
	EncryptionContextEquals map[string]string

	// A list of key-value pairs that must be included in the encryption context of the
	// cryptographic operation
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// request. The grant allows the cryptographic operation only when the encryption
	// context in the request includes the key-value pairs specified in this
	// constraint, although it can include additional key-value pairs.
	EncryptionContextSubset map[string]string

	noSmithyDocumentSerde
}

// Contains information about a grant.
type GrantListEntry struct {

	// A list of key-value pairs that must be present in the encryption context of
	// certain subsequent operations that the grant allows.
	Constraints *GrantConstraints

	// The date and time when the grant was created.
	CreationDate *time.Time

	// The unique identifier for the grant.
	GrantId *string

	// The identity that gets the permissions in the grant. The GranteePrincipal field
	// in the ListGrants response usually contains the user or role designated as the
	// grantee principal in the grant. However, when the grantee principal in the grant
	// is an Amazon Web Services service, the GranteePrincipal field contains the
	// service principal
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html#principal-services),
	// which might represent several different grantee principals.
	GranteePrincipal *string

	// The Amazon Web Services account under which the grant was issued.
	IssuingAccount *string

	// The unique identifier for the KMS key to which the grant applies.
	KeyId *string

	// The friendly name that identifies the grant. If a name was provided in the
	// CreateGrant request, that name is returned. Otherwise this value is null.
	Name *string

	// The list of operations permitted by the grant.
	Operations []GrantOperation

	// The principal that can retire the grant.
	RetiringPrincipal *string

	noSmithyDocumentSerde
}

// Contains information about each entry in the key list.
type KeyListEntry struct {

	// ARN of the key.
	KeyArn *string

	// Unique identifier of the key.
	KeyId *string

	noSmithyDocumentSerde
}

// Contains metadata about a KMS key. This data type is used as a response element
// for the CreateKey, DescribeKey, and ReplicateKey operations.
type KeyMetadata struct {

	// The globally unique identifier for the KMS key.
	//
	// This member is required.
	KeyId *string

	// The twelve-digit account ID of the Amazon Web Services account that owns the KMS
	// key.
	AWSAccountId *string

	// The Amazon Resource Name (ARN) of the KMS key. For examples, see Key Management
	// Service (KMS)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kms)
	// in the Example ARNs section of the Amazon Web Services General Reference.
	Arn *string

	// The cluster ID of the CloudHSM cluster that contains the key material for the
	// KMS key. When you create a KMS key in an CloudHSM custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html),
	// KMS creates the key material for the KMS key in the associated CloudHSM cluster.
	// This field is present only when the KMS key is created in an CloudHSM key store.
	CloudHsmClusterId *string

	// The date and time when the KMS key was created.
	CreationDate *time.Time

	// A unique identifier for the custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
	// that contains the KMS key. This field is present only when the KMS key is
	// created in a custom key store.
	CustomKeyStoreId *string

	// Instead, use the KeySpec field. The KeySpec and CustomerMasterKeySpec fields
	// have the same value. We recommend that you use the KeySpec field in your code.
	// However, to avoid breaking changes, KMS supports both fields.
	//
	// Deprecated: This field has been deprecated. Instead, use the KeySpec field.
	CustomerMasterKeySpec CustomerMasterKeySpec

	// The date and time after which KMS deletes this KMS key. This value is present
	// only when the KMS key is scheduled for deletion, that is, when its KeyState is
	// PendingDeletion. When the primary key in a multi-Region key is scheduled for
	// deletion but still has replica keys, its key state is PendingReplicaDeletion and
	// the length of its waiting period is displayed in the PendingDeletionWindowInDays
	// field.
	DeletionDate *time.Time

	// The description of the KMS key.
	Description *string

	// Specifies whether the KMS key is enabled. When KeyState is Enabled this value is
	// true, otherwise it is false.
	Enabled bool

	// The encryption algorithms that the KMS key supports. You cannot use the KMS key
	// with other encryption algorithms within KMS. This value is present only when the
	// KeyUsage of the KMS key is ENCRYPT_DECRYPT.
	EncryptionAlgorithms []EncryptionAlgorithmSpec

	// Specifies whether the KMS key's key material expires. This value is present only
	// when Origin is EXTERNAL, otherwise this value is omitted.
	ExpirationModel ExpirationModelType

	// The manager of the KMS key. KMS keys in your Amazon Web Services account are
	// either customer managed or Amazon Web Services managed. For more information
	// about the difference, see KMS keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#kms_keys)
	// in the Key Management Service Developer Guide.
	KeyManager KeyManagerType

	// Describes the type of key material in the KMS key.
	KeySpec KeySpec

	// The current status of the KMS key. For more information about how key state
	// affects the use of a KMS key, see Key states of KMS keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
	// Key Management Service Developer Guide.
	KeyState KeyState

	// The cryptographic operations
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// for which you can use the KMS key.
	KeyUsage KeyUsageType

	// The message authentication code (MAC) algorithm that the HMAC KMS key supports.
	// This value is present only when the KeyUsage of the KMS key is
	// GENERATE_VERIFY_MAC.
	MacAlgorithms []MacAlgorithmSpec

	// Indicates whether the KMS key is a multi-Region (True) or regional (False) key.
	// This value is True for multi-Region primary and replica keys and False for
	// regional KMS keys. For more information about multi-Region keys, see
	// Multi-Region keys in KMS
	// (https://docs.aws.amazon.com/kms/latest/developerguide/multi-region-keys-overview.html)
	// in the Key Management Service Developer Guide.
	MultiRegion *bool

	// Lists the primary and replica keys in same multi-Region key. This field is
	// present only when the value of the MultiRegion field is True. For more
	// information about any listed KMS key, use the DescribeKey operation.
	//
	// *
	// MultiRegionKeyType indicates whether the KMS key is a PRIMARY or REPLICA key.
	//
	// *
	// PrimaryKey displays the key ARN and Region of the primary key. This field
	// displays the current KMS key if it is the primary key.
	//
	// * ReplicaKeys displays
	// the key ARNs and Regions of all replica keys. This field includes the current
	// KMS key if it is a replica key.
	MultiRegionConfiguration *MultiRegionConfiguration

	// The source of the key material for the KMS key. When this value is AWS_KMS, KMS
	// created the key material. When this value is EXTERNAL, the key material was
	// imported or the KMS key doesn't have any key material. When this value is
	// AWS_CLOUDHSM, the key material was created in the CloudHSM cluster associated
	// with a custom key store.
	Origin OriginType

	// The waiting period before the primary key in a multi-Region key is deleted. This
	// waiting period begins when the last of its replica keys is deleted. This value
	// is present only when the KeyState of the KMS key is PendingReplicaDeletion. That
	// indicates that the KMS key is the primary key in a multi-Region key, it is
	// scheduled for deletion, and it still has existing replica keys. When a
	// single-Region KMS key or a multi-Region replica key is scheduled for deletion,
	// its deletion date is displayed in the DeletionDate field. However, when the
	// primary key in a multi-Region key is scheduled for deletion, its waiting period
	// doesn't begin until all of its replica keys are deleted. This value displays
	// that waiting period. When the last replica key in the multi-Region key is
	// deleted, the KeyState of the scheduled primary key changes from
	// PendingReplicaDeletion to PendingDeletion and the deletion date appears in the
	// DeletionDate field.
	PendingDeletionWindowInDays *int32

	// The signing algorithms that the KMS key supports. You cannot use the KMS key
	// with other signing algorithms within KMS. This field appears only when the
	// KeyUsage of the KMS key is SIGN_VERIFY.
	SigningAlgorithms []SigningAlgorithmSpec

	// The time at which the imported key material expires. When the key material
	// expires, KMS deletes the key material and the KMS key becomes unusable. This
	// value is present only for KMS keys whose Origin is EXTERNAL and whose
	// ExpirationModel is KEY_MATERIAL_EXPIRES, otherwise this value is omitted.
	ValidTo *time.Time

	// Information about the external key that is associated with a KMS key in an
	// external key store. For more information, see External key
	// (https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html#concept-external-key)
	// in the Key Management Service Developer Guide.
	XksKeyConfiguration *XksKeyConfigurationType

	noSmithyDocumentSerde
}

// Describes the configuration of this multi-Region key. This field appears only
// when the KMS key is a primary or replica of a multi-Region key. For more
// information about any listed KMS key, use the DescribeKey operation.
type MultiRegionConfiguration struct {

	// Indicates whether the KMS key is a PRIMARY or REPLICA key.
	MultiRegionKeyType MultiRegionKeyType

	// Displays the key ARN and Region of the primary key. This field includes the
	// current KMS key if it is the primary key.
	PrimaryKey *MultiRegionKey

	// displays the key ARNs and Regions of all replica keys. This field includes the
	// current KMS key if it is a replica key.
	ReplicaKeys []MultiRegionKey

	noSmithyDocumentSerde
}

// Describes the primary or replica key in a multi-Region key.
type MultiRegionKey struct {

	// Displays the key ARN of a primary or replica key of a multi-Region key.
	Arn *string

	// Displays the Amazon Web Services Region of a primary or replica key in a
	// multi-Region key.
	Region *string

	noSmithyDocumentSerde
}

// A key-value pair. A tag consists of a tag key and a tag value. Tag keys and tag
// values are both required, but tag values can be empty (null) strings. For
// information about the rules that apply to tag keys and tag values, see
// User-Defined Tag Restrictions
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html)
// in the Amazon Web Services Billing and Cost Management User Guide.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	TagKey *string

	// The value of the tag.
	//
	// This member is required.
	TagValue *string

	noSmithyDocumentSerde
}

// Information about the external key
// (https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html#concept-external-key)that
// is associated with a KMS key in an external key store. These fields appear in a
// CreateKey or DescribeKey response only for a KMS key in an external key store.
// The external key is a symmetric encryption key that is hosted by an external key
// manager outside of Amazon Web Services. When you use the KMS key in an external
// key store in a cryptographic operation, the cryptographic operation is performed
// in the external key manager using the specified external key. For more
// information, see External key
// (https://docs.aws.amazon.com/kms/latest/developerguide/keystore-external.html#concept-external-key)
// in the Key Management Service Developer Guide.
type XksKeyConfigurationType struct {

	// The ID of the external key in its external key manager. This is the ID that the
	// external key store proxy uses to identify the external key.
	Id *string

	noSmithyDocumentSerde
}

// KMS uses the authentication credential to sign requests that it sends to the
// external key store proxy (XKS proxy) on your behalf. You establish these
// credentials on your external key store proxy and report them to KMS. The
// XksProxyAuthenticationCredential includes two required elements.
type XksProxyAuthenticationCredentialType struct {

	// A unique identifier for the raw secret access key.
	//
	// This member is required.
	AccessKeyId *string

	// A secret string of 43-64 characters. Valid characters are a-z, A-Z, 0-9, /, +,
	// and =.
	//
	// This member is required.
	RawSecretAccessKey *string

	noSmithyDocumentSerde
}

// Detailed information about the external key store proxy (XKS proxy). Your
// external key store proxy translates KMS requests into a format that your
// external key manager can understand. These fields appear in a
// DescribeCustomKeyStores response only when the CustomKeyStoreType is
// EXTERNAL_KEY_STORE.
type XksProxyConfigurationType struct {

	// The part of the external key store proxy authentication credential
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateCustomKeyStore.html#KMS-CreateCustomKeyStore-request-XksProxyAuthenticationCredential)
	// that uniquely identifies the secret access key.
	AccessKeyId *string

	// Indicates whether the external key store proxy uses a public endpoint or an
	// Amazon VPC endpoint service to communicate with KMS.
	Connectivity XksProxyConnectivityType

	// The URI endpoint for the external key store proxy. If the external key store
	// proxy has a public endpoint, it is displayed here. If the external key store
	// proxy uses an Amazon VPC endpoint service name, this field displays the private
	// DNS name associated with the VPC endpoint service.
	UriEndpoint *string

	// The path to the external key store proxy APIs.
	UriPath *string

	// The Amazon VPC endpoint service used to communicate with the external key store
	// proxy. This field appears only when the external key store proxy uses an Amazon
	// VPC endpoint service to communicate with KMS.
	VpcEndpointServiceName *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
