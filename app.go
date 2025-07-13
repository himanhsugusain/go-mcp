package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"go.lsp.dev/jsonrpc2"
)

type Backend interface{
	GetCapabilities() Capabilities 
	ListTools() ListToolResponse
	ToolsCall(*jsonrpc2.Call) map[string]any
}
type App struct{
	backend Backend
}

func NewApp(backend Backend) *App {
	return &App{
		backend : backend,
	}
}

func (a *App) GetInitResponse() InitResponse{
	return InitResponse{
		ProtocolVersion: protocolVersion,
		Capabilities: a.backend.GetCapabilities(),
		ServerInfo: serverInfo,
		Instructions: "",
	}
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request){
	call := jsonrpc2.Call{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(data))
	if err := json.Unmarshal(data, &call); err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := a.mcpHandler(&call, w); err != nil{
		log.Println(err)
		// TODO: return http error
	}
}

func (a *App)mcpHandler(call *jsonrpc2.Call,w http.ResponseWriter) error {
	method := call.Method()
	var err error
	var resp *jsonrpc2.Response
	
	switch method {
		case "initialize":
			resp, err = jsonrpc2.NewResponse(call.ID(), a.GetInitResponse(), nil)
		case "tools/list":
			resp, err = jsonrpc2.NewResponse(call.ID(), a.backend.ListTools(), nil)
		case "tools/call":
			resp, err = jsonrpc2.NewResponse(call.ID(), a.backend.ToolsCall(call), nil) 
		default:
			resp, err = jsonrpc2.NewResponse(call.ID(),nil,errors.New("method Not Found")) 
	}
	if err != nil {
		return err
	}
	respJSON, err := resp.MarshalJSON()
	if err != nil {
		return err
	}
	if method == "initialized" {
		// a hack to make it work with mcp inspectore
		fmt.Fprintf(w,"[%s]", respJSON)
		return nil
	}
	fmt.Fprint(w,string(respJSON))
	return nil
}
