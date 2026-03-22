//go:build js && wasm

package main

import (
	"encoding/json"
	"syscall/js"
)

func main() {
	js.Global().Set("FormatJson", formatJsonWrapper())
	select {}
}

func formatJsonWrapper() js.Func {
	return js.FuncOf(func(_ js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid number of argumnets"
		}
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			return "Unable to get document object"
		}
		jsOutput := jsDoc.Call("querySelector", "#jsonOutput")
		if !jsOutput.Truthy() {
			return "Unable to get jsOutput"
		}

		prettryJson, err := preetify(args[0].String())
		if err != nil {
			return err.Error()
		}
		jsOutput.Set("value", prettryJson)
		return nil
	})
}

func preetify(input string) (string, error) {
	var data any
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return "", err
	}
	pretty_json, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return "", err
	}
	return string(pretty_json), nil
}
