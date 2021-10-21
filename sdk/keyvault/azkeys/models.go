//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azkeys

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azkeys/internal"
)

// Attributes - The object attributes managed by the KeyVault service.
type Attributes struct {
	// Determines whether the object is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Expiry date in UTC.
	Expires *time.Time `json:"exp,omitempty"`

	// Not before date in UTC.
	NotBefore *time.Time `json:"nbf,omitempty"`

	// READ-ONLY; Creation time in UTC.
	Created *time.Time `json:"created,omitempty" azure:"ro"`

	// READ-ONLY; Last updated time in UTC.
	Updated *time.Time `json:"updated,omitempty" azure:"ro"`
}

// KeyAttributes - The attributes of a key managed by the key vault service.
type KeyAttributes struct {
	Attributes
	// READ-ONLY; softDelete data retention days. Value should be >=7 and <=90 when softDelete enabled, otherwise 0.
	RecoverableDays *int32 `json:"recoverableDays,omitempty" azure:"ro"`

	// READ-ONLY; Reflects the deletion recovery level currently in effect for keys in the current vault. If it contains 'Purgeable' the key can be permanently
	// deleted by a privileged user; otherwise, only the system
	// can purge the key, at the end of the retention interval.
	RecoveryLevel *DeletionRecoveryLevel `json:"recoveryLevel,omitempty" azure:"ro"`
}

// converts a KeyAttributes to *internal.KeyAttributes
func (k KeyAttributes) toGenerated() *internal.KeyAttributes {
	return &internal.KeyAttributes{
		RecoverableDays: k.RecoverableDays,
		RecoveryLevel:   recoveryLevelToGenerated(k.RecoveryLevel),
		Attributes: internal.Attributes{
			Enabled:   k.Enabled,
			Expires:   k.Expires,
			NotBefore: k.NotBefore,
			Created:   k.Created,
			Updated:   k.Updated,
		},
	}
}

// converts *internal.KeyAttributes to *KeyAttributes
func keyAttributesFromGenerated(i *internal.KeyAttributes) *KeyAttributes {
	if i == nil {
		return &KeyAttributes{}
	}

	return &KeyAttributes{
		RecoverableDays: i.RecoverableDays,
		RecoveryLevel:   DeletionRecoveryLevel(*i.RecoveryLevel).ToPtr(),
		Attributes: Attributes{
			Enabled:   i.Enabled,
			Expires:   i.Expires,
			NotBefore: i.NotBefore,
			Created:   i.Created,
			Updated:   i.Updated,
		},
	}
}

// KeyBundle - A KeyBundle consisting of a WebKey plus its attributes.
type KeyBundle struct {
	// The key management attributes.
	Attributes *KeyAttributes `json:"attributes,omitempty"`

	// The Json web key.
	Key *JSONWebKey `json:"key,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; True if the key's lifetime is managed by key vault. If this is a key backing a certificate, then managed will be true.
	Managed *bool `json:"managed,omitempty" azure:"ro"`
}

// JSONWebKey - As of http://tools.ietf.org/html/draft-ietf-jose-json-web-key-18
type JSONWebKey struct {
	// Elliptic curve name. For valid values, see JsonWebKeyCurveName.
	Crv *JSONWebKeyCurveName `json:"crv,omitempty"`

	// RSA private exponent, or the D component of an EC private key.
	D []byte `json:"d,omitempty"`

	// RSA private key parameter.
	DP []byte `json:"dp,omitempty"`

	// RSA private key parameter.
	DQ []byte `json:"dq,omitempty"`

	// RSA public exponent.
	E []byte `json:"e,omitempty"`

	// Symmetric key.
	K      []byte    `json:"k,omitempty"`
	KeyOps []*string `json:"key_ops,omitempty"`

	// Key identifier.
	ID *string `json:"kid,omitempty"`

	// JsonWebKey Key Type (kty), as defined in https://tools.ietf.org/html/draft-ietf-jose-json-web-algorithms-40.
	KeyType *KeyType `json:"kty,omitempty"`

	// RSA modulus.
	N []byte `json:"n,omitempty"`

	// RSA secret prime.
	P []byte `json:"p,omitempty"`

	// RSA secret prime, with p < q.
	Q []byte `json:"q,omitempty"`

	// RSA private key parameter.
	QI []byte `json:"qi,omitempty"`

	// Protected Key, used with 'Bring Your Own Key'.
	T []byte `json:"key_hsm,omitempty"`

	// X component of an EC public key.
	X []byte `json:"x,omitempty"`

	// Y component of an EC public key.
	Y []byte `json:"y,omitempty"`
}

// converts internal.JSONWebKey to publicly exposed version
func jsonWebKeyFromGenerated(i *internal.JSONWebKey) *JSONWebKey {
	if i == nil {
		return &JSONWebKey{}
	}

	return &JSONWebKey{
		Crv:     (*JSONWebKeyCurveName)(i.Crv),
		D:       i.D,
		DP:      i.DP,
		DQ:      i.DQ,
		E:       i.E,
		K:       i.K,
		KeyOps:  i.KeyOps,
		ID:      i.Kid,
		KeyType: (*KeyType)(i.Kty),
		N:       i.N,
		P:       i.P,
		Q:       i.Q,
		QI:      i.QI,
		T:       i.T,
		X:       i.X,
		Y:       i.Y,
	}
}

// converts JSONWebKey to *internal.JSONWebKey
func (j JSONWebKey) toGenerated() *internal.JSONWebKey {
	return &internal.JSONWebKey{
		Crv:    (*internal.JSONWebKeyCurveName)(j.Crv),
		D:      j.D,
		DP:     j.DP,
		DQ:     j.DQ,
		E:      j.E,
		K:      j.K,
		KeyOps: j.KeyOps,
		Kid:    j.ID,
		Kty:    (*internal.JSONWebKeyType)(j.KeyType),
		N:      j.N,
		P:      j.P,
		Q:      j.Q,
		QI:     j.QI,
		T:      j.T,
		X:      j.X,
		Y:      j.Y,
	}
}

// KeyType - JsonWebKey Key Type (kty), as defined in https://tools.ietf.org/html/draft-ietf-jose-json-web-algorithms-40.
type KeyType string

const (
	// EC - Elliptic Curve.
	EC KeyType = "EC"

	// ECHSM - Elliptic Curve with a private key which is not exportable from the HSM.
	ECHSM KeyType = "EC-HSM"

	// Oct - Octet sequence (used to represent symmetric keys)
	Oct KeyType = "oct"

	// OctHSM - Octet sequence (used to represent symmetric keys) which is not exportable from the HSM.
	OctHSM KeyType = "oct-HSM"

	// RSA - RSA (https://tools.ietf.org/html/rfc3447)
	RSA KeyType = "RSA"

	// RSAHSM - RSA with a private key which is not exportable from the HSM.
	RSAHSM KeyType = "RSA-HSM"
)

// convert KeyType to *internal.JSONWebKeyType
func (j KeyType) toGenerated() *internal.JSONWebKeyType {
	return internal.JSONWebKeyType(j).ToPtr()
}

// KeyItem - The key item containing key metadata.
type KeyItem struct {
	// The key management attributes.
	Attributes *KeyAttributes `json:"attributes,omitempty"`

	// Key identifier.
	KID *string `json:"kid,omitempty"`

	// Application specific metadata in the form of key-value pairs.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; True if the key's lifetime is managed by key vault. If this is a key backing a certificate, then managed will be true.
	Managed *bool `json:"managed,omitempty" azure:"ro"`
}

// convert *internal.KeyItem to *KeyItem
func keyItemFromGenerated(i *internal.KeyItem) *KeyItem {
	if i == nil {
		return nil
	}

	return &KeyItem{
		Attributes: keyAttributesFromGenerated(i.Attributes),
		KID:        i.Kid,
		Tags:       i.Tags,
		Managed:    i.Managed,
	}
}

// DeletedKeyBundle - A DeletedKeyBundle consisting of a WebKey plus its Attributes and deletion info
type DeletedKeyBundle struct {
	KeyBundle
	// The url of the recovery object, used to identify and recover the deleted key.
	RecoveryID *string `json:"recoveryId,omitempty"`

	// READ-ONLY; The time when the key was deleted, in UTC
	DeletedDate *time.Time `json:"deletedDate,omitempty" azure:"ro"`

	// READ-ONLY; The time when the key is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time `json:"scheduledPurgeDate,omitempty" azure:"ro"`
}

// DeletedKeyItem - The deleted key item containing the deleted key metadata and information about deletion.
type DeletedKeyItem struct {
	KeyItem
	// The url of the recovery object, used to identify and recover the deleted key.
	RecoveryID *string `json:"recoveryId,omitempty"`

	// READ-ONLY; The time when the key was deleted, in UTC
	DeletedDate *time.Time `json:"deletedDate,omitempty" azure:"ro"`

	// READ-ONLY; The time when the key is scheduled to be purged, in UTC
	ScheduledPurgeDate *time.Time `json:"scheduledPurgeDate,omitempty" azure:"ro"`
}

// convert *internal.DeletedKeyItem to *DeletedKeyItem
func deletedKeyItemFromGenerated(i *internal.DeletedKeyItem) *DeletedKeyItem {
	if i == nil {
		return nil
	}

	return &DeletedKeyItem{
		RecoveryID:         i.RecoveryID,
		DeletedDate:        i.DeletedDate,
		ScheduledPurgeDate: i.ScheduledPurgeDate,
		KeyItem:            *keyItemFromGenerated(&i.KeyItem),
	}
}
