# shell.rescan.path: `murex-update-exe-list`

> Forces Murex to rescan $PATH looking for executables

## Description

On application launch, Murex scans and caches all the executables found in
$PATH on your host. Murex then does regular scans there after. However if
you want to force a new scan (for example you've just installed a new
program and you want it to appear in tab completion) then you can run `murex-update-exe-list`.

## Usage

```
murex-update-exe-list
```

## Examples

```
» murex-update-exe-list
```

## Detail

Murex will automatically update the exe list each time tab completion is
invoked for command name completion via the REPL shell.

## Synonyms

* `murex-update-exe-list`
* `shell.rescan.path`


## See Also

* [sys.cpu.arch: `cpuarch`](../commands/cpuarch.md):
  Output the hosts CPU architecture
* [sys.cpu.count: `cpucount`](../commands/cpucount.md):
  Output the number of CPU cores available on your host
* [sys.os: `os`](../commands/os.md):
  Output the auto-detected OS name

<hr/>

This document was generated from [builtins/core/management/functions_doc.yaml](https://github.com/lmorg/murex/blob/master/builtins/core/management/functions_doc.yaml).