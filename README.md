I haven't experimented with various edge cases as in C yet, so I'm not entirely sure whether this transpiler works perfectly or if it might break the source code targets.  

If you'd like to try it out, feel free to do so.  

However, my personal recommendation is to use the C implementation of the transpiler if you need a well-tested and reliable solution. You can find it here:  
[VMK Transpiler (C Version)](https://github.com/nginrsw/vmk-lang/tree/main/transpiler/src)  

### How to Build  

To reduce the binary size, it's recommended to build the transpiler using the following command:  

```bash
go build -ldflags "-s -w" -o <program-name> <transpiler-file.go>
```

This will generate a smaller executable by stripping debug and symbol information.  
