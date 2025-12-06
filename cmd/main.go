package main

import (
	"fmt"
	"log"
	"os"

	"github.com/malivvan/cui/markup/qjs"
)

func must[T any](val T, err error) T {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return val
}

const script = `
// JS handlers for HTTP routes
const about = () => {
	return "QuickJS in Go - Hello World!";
};



export default { about };
`

func main() {
	rt := must(qjs.New())
	defer rt.Close()
	ctx := rt.Context()

	// Precompile the script to bytecode
	byteCode := must(ctx.Compile("script.js", qjs.Code(script), qjs.TypeModule()))

	// Use a pool of runtimes for concurrent requests
	pool := qjs.NewPool(3, qjs.Option{
		CacheDir: "/home/malivvan/.cache",
	}, func(r *qjs.Runtime) error {
		results := must(r.Context().Eval("script.js", qjs.Bytecode(byteCode), qjs.TypeModule()))
		results.ForEach(func(key *qjs.Value, value *qjs.Value) {
			if value.IsFunction() {
				// call
				log.Printf("Registered handler: %s", key.String())
				// run function to ensure it's valid
				value.Free()

			}
		})
		// Store the exported functions in the global object for easy access
		r.Context().Global().SetPropertyStr("handlers", results)
		return nil
	})

	// Register HTTP handlers based on JS functions
	val := must(ctx.Eval("script.js", qjs.Bytecode(byteCode), qjs.TypeModule()))
	methodNames := must(val.GetOwnPropertyNames())
	val.Free()

	for _, methodName := range methodNames {
		if methodName == "about" {
			runtime := must(pool.Get())
			defer pool.Put(runtime)

			// Call the corresponding JS function
			handlers := runtime.Context().Global().GetPropertyStr("handlers")
			result := must(handlers.InvokeJS(methodName))
			fmt.Fprint(os.Stdout, result.String())
			result.Free()
		}
	}

}
