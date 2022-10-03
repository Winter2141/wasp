// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

export class StoreBlobCall {
	func: wasmlib.ScFunc;
	params: sc.MutableStoreBlobParams = new sc.MutableStoreBlobParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableStoreBlobResults = new sc.ImmutableStoreBlobResults(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScFuncCallContext) {
		this.func = new wasmlib.ScFunc(ctx, sc.HScName, sc.HFuncStoreBlob);
	}
}

export class GetBlobFieldCall {
	func: wasmlib.ScView;
	params: sc.MutableGetBlobFieldParams = new sc.MutableGetBlobFieldParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetBlobFieldResults = new sc.ImmutableGetBlobFieldResults(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScViewCallContext) {
		this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetBlobField);
	}
}

export class GetBlobInfoCall {
	func: wasmlib.ScView;
	params: sc.MutableGetBlobInfoParams = new sc.MutableGetBlobInfoParams(wasmlib.ScView.nilProxy);
	results: sc.ImmutableGetBlobInfoResults = new sc.ImmutableGetBlobInfoResults(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScViewCallContext) {
		this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetBlobInfo);
	}
}

export class ListBlobsCall {
	func: wasmlib.ScView;
	results: sc.ImmutableListBlobsResults = new sc.ImmutableListBlobsResults(wasmlib.ScView.nilProxy);
	public constructor(ctx: wasmlib.ScViewCallContext) {
		this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewListBlobs);
	}
}

export class ScFuncs {
	static storeBlob(ctx: wasmlib.ScFuncCallContext): StoreBlobCall {
		const f = new StoreBlobCall(ctx);
		f.params = new sc.MutableStoreBlobParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableStoreBlobResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static getBlobField(ctx: wasmlib.ScViewCallContext): GetBlobFieldCall {
		const f = new GetBlobFieldCall(ctx);
		f.params = new sc.MutableGetBlobFieldParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetBlobFieldResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static getBlobInfo(ctx: wasmlib.ScViewCallContext): GetBlobInfoCall {
		const f = new GetBlobInfoCall(ctx);
		f.params = new sc.MutableGetBlobInfoParams(wasmlib.newCallParamsProxy(f.func));
		f.results = new sc.ImmutableGetBlobInfoResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}

	static listBlobs(ctx: wasmlib.ScViewCallContext): ListBlobsCall {
		const f = new ListBlobsCall(ctx);
		f.results = new sc.ImmutableListBlobsResults(wasmlib.newCallResultsProxy(f.func));
		return f;
	}
}
