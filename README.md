# Platform-Invariant-WASM-VM
A Platform Invariant Webassembly virtual machine written in GO.

Does not rely on any native dependencies in interpreter-only mode, and should be easily cross-compiled for running WebAssembly modules on most popular platforms (Windows/Linux/Mac/Android/iOS/etc).
Implements WebAssembly execution semantics and passes most of the [official test suite](https://github.com/WebAssembly/testsuite)

## Executing WebAssembly Modules

For this example asssume we have already loaded our *.wasm module's bytecode into the variable `var input []byte`.

Lets pass the bytecode into a newly instantiated virtual machine:
```go
vm, err := exec.NewVirtualMachine(input, exec.VMConfig{}, &exec.NopResolver{}, nil)
if err != nil { // if the wasm bytecode is invalid
    panic(err)
}
```

Lookup the function ID to a desired entry-point function titled `app_main`:
```go
entryID, ok := vm.GetFunctionExport("app_main") // can be changed to your own exported function
if !ok {
    panic("entry function not found")
}
```

And startup the VM; printing out the result of the entry-point function:
```go
ret, err := vm.Run(entryID)
if err != nil {
    vm.PrintStackTrace()
    panic(err)
}
fmt.Printf("return value = %d\n", ret)
```

## notes on interpreter updates and fundamental rewrite.

As of 2020/5/11 the main webassembly go interpreter (https://github.com/go-interpreter/wagon)
has been converted to a read-only repository. This naturally requires a fundamental rewrite of this virtual machine
in to a better maintained WASM interpreter. In this case I have selected to rewrite this virtual machine
utilizing (https://github.com/tetratelabs/wazero) the wazero intepreter library. The previous source code
passed 65/72 of the WASM test suite with only a single failure related to execution semantics.

I have set a much higher standard and would like to see a score of 67/72 or higher with zero execution semantic errors. current testing has resulted in a score of 64/72.
some fine tuning is required to bring the source code to previous standards.
