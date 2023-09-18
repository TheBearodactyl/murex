# Tilde (`~`) Token

> Home directory path variable

## Description

The tilde token is used as a lazy reference to the users home directory.

## Examples

```
» out ~
/home/bob

» out ~joe
/home/joe
```

## Detail

Tilde can be expanded inside double quotes, brace quotes as well as used naked.
But it cannot be expanded inside single quotes.

```
» out ~
/home/bob

» out '~'
~

» out "~"
/home/bob

» out %(~)
/home/bob
```

## See Also

* [Array (`@`) Token](../parser/array.md):
  Expand values as an array
* [Double Quote (`"`) Token](../parser/double-quote.md):
  Initiates or terminates a string (variables expanded)
* [Reserved Variables](../user-guide/reserved-vars.md):
  Special variables reserved by Murex
* [Single Quote (`'`) Token](../parser/single-quote.md):
  Initiates or terminates a string (variables not expanded)
* [`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`ja` (mkarray)](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`out`](../commands/out.md):
  Print a string to the STDOUT with a trailing new line character
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [`string` (stringing)](../types/str.md):
  string (primitive)

<hr/>

This document was generated from [gen/parser/variables_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/parser/variables_doc.yaml).