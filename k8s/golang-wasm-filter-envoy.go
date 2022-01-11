package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
}

// Override types.DefaultPluginContext.
func (*pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpHeaders{contextID: contextID}
}

type httpHeaders struct {
	// Embed the default http context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultHttpContext
	contextID uint32
}

// Override types.DefaultHttpContext.
func (ctx *httpHeaders) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	token, err := proxywasm.GetHttpRequestHeader("authorization")
	if err != nil {
		proxywasm.LogCriticalf("failed to get 'authorization' header: %v", err)
		return types.ActionContinue
	}

	if token == "" {
		return types.ActionContinue
	}

	proxywasm.LogInfof("request authorization token: %v", token)

	path, err := proxywasm.GetHttpRequestHeader(":path")
	if err != nil {
		proxywasm.LogCriticalf("failed to get 'path' header: %v", err)
	}

	proxywasm.LogInfof("request path: %v", path)

	if path == "" {
		return types.ActionContinue
	}

	headers := [][2]string{
		{":method", "GET"}, {":authority", "lua_cluster"}, {":path", "/auth"}, {"authorization", token},
	}

	if _, err := proxywasm.DispatchHttpCall("lua_cluster", headers, nil, nil, 5000, callBack); err != nil {
		proxywasm.LogCriticalf("'lua_cluster' call failed: %v", err)
	}

	return types.ActionPause
}

func callBack(numHeaders, bodySize, numTrailers int) {
	headers, err := proxywasm.GetHttpCallResponseHeaders()
	if err != nil && err != types.ErrorStatusNotFound {
		proxywasm.ResumeHttpRequest()
		return
	}

	identity := ""
	for _, h := range headers {
		if h[0] == "x-auth-identity" {
			identity = h[1]
		}
	}

	proxywasm.LogInfof("received 'x-auth-identity': %v", identity)

	if identity == "" {
		proxywasm.ResumeHttpRequest()
		return
	}

	err = proxywasm.AddHttpRequestHeader("x-auth-identity", identity)
	if err != nil {
		proxywasm.LogCriticalf("failed to set 'x-auth-identity' header: %v", err)
		proxywasm.ResumeHttpRequest()
		return
	}

	proxywasm.ResumeHttpRequest()
	return
}
