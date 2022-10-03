// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package coreblob

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type MapStringToImmutableBytes struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToImmutableBytes) GetBytes(key string) wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type ImmutableStoreBlobParams struct {
	proxy wasmtypes.Proxy
}

// set of named blobs
func (s ImmutableStoreBlobParams) Blobs() MapStringToImmutableBytes {
	return MapStringToImmutableBytes(s)
}

// description of progBinary
func (s ImmutableStoreBlobParams) Description() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamDescription))
}

// smart contract program binary code
func (s ImmutableStoreBlobParams) ProgBinary() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ParamProgBinary))
}

// VM type that must be used to run progBinary
func (s ImmutableStoreBlobParams) VMType() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamVMType))
}

type MapStringToMutableBytes struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToMutableBytes) Clear() {
	m.proxy.ClearMap()
}

func (m MapStringToMutableBytes) GetBytes(key string) wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type MutableStoreBlobParams struct {
	proxy wasmtypes.Proxy
}

// set of named blobs
func (s MutableStoreBlobParams) Blobs() MapStringToMutableBytes {
	return MapStringToMutableBytes(s)
}

// description of progBinary
func (s MutableStoreBlobParams) Description() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamDescription))
}

// smart contract program binary code
func (s MutableStoreBlobParams) ProgBinary() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ParamProgBinary))
}

// VM type that must be used to run progBinary
func (s MutableStoreBlobParams) VMType() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamVMType))
}

type ImmutableGetBlobFieldParams struct {
	proxy wasmtypes.Proxy
}

// blob name
func (s ImmutableGetBlobFieldParams) Field() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamField))
}

// blob set
func (s ImmutableGetBlobFieldParams) Hash() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamHash))
}

type MutableGetBlobFieldParams struct {
	proxy wasmtypes.Proxy
}

// blob name
func (s MutableGetBlobFieldParams) Field() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamField))
}

// blob set
func (s MutableGetBlobFieldParams) Hash() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamHash))
}

type ImmutableGetBlobInfoParams struct {
	proxy wasmtypes.Proxy
}

// blob set
func (s ImmutableGetBlobInfoParams) Hash() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamHash))
}

type MutableGetBlobInfoParams struct {
	proxy wasmtypes.Proxy
}

// blob set
func (s MutableGetBlobInfoParams) Hash() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamHash))
}