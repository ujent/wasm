# WASMification

**To compile WASM:**

cd cmd/wasm
GOOS=js GOARCH=wasm go build -o ../../assets/json.wasm

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./assets
