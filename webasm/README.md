# An implementation of the Barnsley Fern fractal algorithm in Go + WebAssembly

A [WebAssembly](https://github.com/golang/go/wiki/WebAssembly) version that renders directly onto the canvas: it's really rather slow!

It will render in your browser thus:

Copy _wasm_exec.js_ to this folder:

```shell
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

Compile the code to wasm:

```shell
GOOS=js GOARCH=wasm go build -o barnsley.wasm 
```

Then you will need to serve the _index.html_ file via a web server and visit it in your browser.