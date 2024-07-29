# `pretty`

> Prettifies JSON to make it human readable

## Description

Takes JSON from the stdin and reformats it to make it human readable, then
outputs that to stdout.

## Usage

```
<stdin> -> pretty -> <stdout>
```

## Examples

```
» tout json {"Array":[1,2,3],"Map":{"String": "Foobar","Number":123.456}} -> pretty 
{
    "Array": [
        1,
        2,
        3
    ],
    "Map": {
        "String": "Foobar",
        "Number": 123.456
    }
}
```

## See Also

* [`format`](../commands/format.md):
  Reformat one data-type into another data-type
* [`out`](../commands/out.md):
  Print a string to the stdout with a trailing new line character
* [`tout`](../commands/tout.md):
  Print a string to the stdout and set it's data-type

<hr/>

This document was generated from [builtins/core/pretty/pretty_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/pretty/pretty_doc.yaml).