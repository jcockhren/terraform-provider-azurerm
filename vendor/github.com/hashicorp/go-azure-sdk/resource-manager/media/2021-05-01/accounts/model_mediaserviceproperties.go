package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaServiceProperties struct {
	Encryption            *AccountEncryption     `json:"encryption,omitempty"`
	KeyDelivery           *KeyDelivery           `json:"keyDelivery,omitempty"`
	MediaServiceId        *string                `json:"mediaServiceId,omitempty"`
	StorageAccounts       *[]StorageAccount      `json:"storageAccounts,omitempty"`
	StorageAuthentication *StorageAuthentication `json:"storageAuthentication,omitempty"`
}
