# `*` (generic)

> generic (primitive)

## Description

This is the default data type used when STDOUT is returned from any external
executables.

## Supported Hooks

* `Marshal()`
    Supported. Tables columns are aligned
* `ReadArray()`
    Treats each new line as a new array element
* `ReadArrayWithType()`
    Treats each new line as a new array element, each element is `*` 
* `ReadIndex()`
    Indexes treated as table coordinates
* `ReadMap()`
    Works against tables such as the output from `ps -fe` 
* `ReadNotIndex()`
    Indexes treated as table coordinates
* `Unmarshal()`
    Supported
* `WriteArray()`
    Writes a new line per array element - tabs are treated as columns

## See Also

* [`int`](../types/int.md):
  Whole number (primitive)
* [`num` (number)](../types/num.md):
  Floating point number (primitive)
* [`str` (string)](../types/str.md):
  string (primitive)
* [cast](../types/cast.md):
  
* [element](../types/element.md):
  
* [format](../types/format.md):
  
* [index](../types/index.md):
  
* [open](../types/open.md):
  
* [runtime](../types/runtime.md):
  

### Read more about type hooks

- [`ReadIndex()` (type)](../apis/ReadIndex.md): Data type handler for the index, `[`, builtin
- [`ReadNotIndex()` (type)](../apis/ReadNotIndex.md): Data type handler for the bang-prefixed index, `![`, builtin
- [`ReadArray()` (type)](../apis/ReadArray.md): Read from a data type one array element at a time
- [`WriteArray()` (type)](../apis/WriteArray.md): Write a data type, one array element at a time
- [`ReadMap()` (type)](../apis/ReadMap.md): Treat data type as a key/value structure and read its contents
- [`Marshal()` (type)](../apis/Marshal.md): Converts structured memory into a structured file format (eg for stdio)
- [`Unmarshal()` (type)](../apis/Unmarshal.md): Converts a structured file format into structured memory

<hr/>

This document was generated from [builtins/types/generic/generic_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/types/generic/generic_doc.yaml).