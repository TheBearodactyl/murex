{{ if env "DOCGEN_TARGET=" }}<h2>Table of Contents</h2>
<div id="toc">

- [Command Prompt](#command-prompt)
  - [Autocomplete](#autocomplete)
  - [Fuzzy Find Autocomplete](#fuzzy-find-autocomplete)
  - [Autocomplete Preview](#autocomplete-preview)
  - [Command Line Preview](#command-line-preview)
  - [Search Shell History](#search-shell-history)
  - [Line Editing](#line-editing)
    - [Navigation](#navigation)
    - [General Hotkeys](#general-hotkeys)
  - [Vim Keys](#vim-keys)
    - [Supported Keys](#supported-keys)
  - [Recalling Previous Words](#recalling-previous-words)
- [Job Control](#job-control)
- [Miscellaneous](#miscellaneous)
  - [Cancel Prompt](#cancel-prompt)
  - [End Of File](#end-of-file)
  - [Alternative Cancel Key](#alternative-cancel-key)
  - [Clear Screen](#clear-screen)
  - [EDITOR](#editor)

</div>

{{ end }}
## Command Prompt

### Autocomplete

Pressing `tab` provides autocompletion suggestions. Suggestions can come in one
of two formats:

1. a gridded view where the hint text (the, typically blue, text under the
   prompt) provides the description
2. a list view where the description is printed alongside the completion
   suggestion.

While the autocompletion suggestions are open, the following keys are assigned
roles:

* arrow keys (`left`, `right`, `up`, `down`): highlight different suggestions
  
* `tab`: highlight the next suggestion
  
* `shift`+`tab`: highlight the previous suggestion
  
* `enter` / `return`: this selects the highlighted autocompletion
  
* `esc`: closes the suggestions without selecting one
  
* `ctrl`+`f`: fuzzy find in the suggestions

* `f1`: show / hide autocomplete preview box. This will hide your terminal
  output while enabled. The preview box supports additional key bindings
  ([see below](#autocomplete-preview))

* `f9`: show the command line preview box. This will output the contents of
  your command pipeline ([see below](#command-line-preview))

### Fuzzy Find Autocomplete

Pressing `ctrl`+`f` either from the prompt, or while the autocomplete
suggestions are open, will open up the fuzzy find dialog to search through
available suggestions. This can also be used to quickly jump to specific
sub-directories.

Your typed search terms will appear in the hint text.

By default the fuzzy finder will look for any item that includes _all_ of the
search words. However the search behavior can be changed if the first search
term is any of the following:

* `or`: show results that match _any_ of the search terms. eg `or .md .txt`
  will match both markdown and txt files (when finding files in completion
  suggestions).

* `!`: only show suggestions that do not match any of the search terms. eg
  `! .md .txt` will match all files except markdown and txt files (when finding
  files in completion suggestions).

* `g`: show only results that match a shell glob. eg `*.txt`. This mode is
  automatically assumed if you include an abstricts in your search term.

* `rx`: use a regexp pattern matcher instead of any fuzzy search. Expressions
  will be case insensitive and non-greedy by default.

Aside from globbing matches, searching in fuzzy finder is not case sensitive.

While the fuzzy finder is open, the following keys are assigned roles:

* arrow keys (`left`, `right`, `up`, `down`): highlight different suggestions
 
* `tab`: highlight the next suggestion
  
* `shift`+`tab`: highlight the previous suggestion
  
* `enter` / `return`: this selects the highlighted autocompletion
  
* `esc`: cancel search

* `f1`: show / hide preview box. This will hide your terminal output while
  enabled. The preview box supports additional key bindings ([see below](#autocomplete-preview))

* `f9`: show the command line preview box. This will output the contents of
  your command pipeline ([see below](#command-line-preview))

### Autocomplete Preview

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img src="/screenshot-preview-man-rsync.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
<!-- markdownlint-restore -->
{{ end }}

The autocomplete preview is a way of quickly examining the contents of a
function, man page, text file or even image, based on what autocomplete
suggestion is highlighted. ([read more](interactive-shell.md#autocomplete-preview))

While the preview box is open, the rest of your terminal output will be hidden.
However once you close it, that output will reappear.

While the preview box is open, the following keys are assigned roles:

* `f1` or `enter`: closes the preview box

* `f9` switches to command line preview
  
* `page up` scroll up the contents of the preview box, one page at a time
* `ctrl`+`arrow up` scroll up the contents of the preview box, one page at a
  time (IBM keyboard layouts)
* `option`+`arrow up` scroll up the contents of the preview box, one page at a
  time (Apple keyboard layouts)

* `page down` scroll down the contents of the preview box, one page at a time
* `ctrl`+`arrow down` scroll down the contents of the preview box, one page at
  a time (IBM keyboard layouts)
* `option`+`arrow down` scroll down the contents of the preview box, one page
  at a time (Apple keyboard layouts)

* `home` scroll up to start of previous section or start of document
* `end` scroll down to start of next section or end of document

### Command Line Preview

{{ if env "DOCGEN_TARGET=vuepress" }}
<!-- markdownlint-disable -->
<img src="/screenshot-preview-command-line.png?v={{ env "COMMITHASHSHORT" }}" class="centre-image"/>
<!-- markdownlint-restore -->
{{ end }}

The command line preview enables you to view the output of a command pipeline
interactively while you type it. ([read more](interactive-shell.md#command-line-preview))

While the preview box is open, the rest of your terminal output will be hidden.
However once you close it, that output will reappear.

While the preview box is open, the following keys are assigned roles:

* `f1` or `enter`: closes the preview box

* `f9` re-runs the command line and thus updates the contents in the preview
  frame
  
* `page up` scroll up the contents of the preview box, one page at a time
* `ctrl`+`arrow up` scroll up the contents of the preview box, one page at a
  time (IBM keyboard layouts)
* `option`+`arrow up` scroll up the contents of the preview box, one page at a
  time (Apple keyboard layouts)

* `page down` scroll down the contents of the preview box, one page at a time
* `ctrl`+`arrow down` scroll down the contents of the preview box, one page at
  a time (IBM keyboard layouts)
* `option`+`arrow down` scroll down the contents of the preview box, one page
  at a time (Apple keyboard layouts)

* `home` scroll up to the previous section
* `end` scroll down to the next section

### Search Shell History

This displays up your timestamped shell history as an autocomplete list with
fuzzy find activated. Using `ctrl`+`r` you can rapidly rerun previous
command lines.

From here, the usual autocomplete / fuzzy find hotkeys apply. Such as pressing
`esc` to cancel history completion.

If the prompt line is not empty, then the current line is included in the
history search.

### Line Editing

These are the various hotkeys and editing modes available in Murex's
interactive command prompt.

#### Navigation

* Arrow keys, `left` and `right`: move the cursor forwards or backwards in line
  
* Arrow keys, `up` and `down`:
  - If the command line spans multiple lines on the screen then this will jump
    up or down to the next/previous line.
  - When at the top or bottom line, or the command line is only one line long,
    the `up` or `down` keys will search through your history of past command
    lines that are similar to your current command line.
  - If your command line is empty, then the `up` or `down` keys will search
    through every command line in your history.

* `alt`+`b`: jump backwards a word at a time (Emacs compatibility)
* `ctrl`+`left`: jump backwards a word at a time (IBM keyboard layouts)
* `option`+`left`: jump backwards a word at a time (Apple keyboard layouts)
  
* `alt`+`f`: jump forwards a word at a time (Emacs compatibility)
* `ctrl`+`right`: jump forwards a word at a time (IBM keyboard layouts)
* `option`+`right`: jump forwards a word at a time (Apple keyboard layouts)

* `ctrl`+`a`: jump to beginning of line
* `home`: jump to beginning of line

* `ctrl`+`e`: jump to end of line
* `end`: jump to end of line

* `ctrl`+`z`: while readline is open will undo the previous key strokes

#### General Hotkeys

* `ctrl`+`k`: clears line after cursor

* `ctrl`+`n`: next line from history

* `ctrl`+`p`: previous line from history
  
* `ctrl`+`u`: clears the whole line

### Vim Keys

Pressing `esc` while no autocomplete suggestions are shown will switch the
line editor into **vim keys** mode.

Press `i` to return to normal editing mode.

#### Supported Keys

* `a`: insert after current character
* `A`: insert at end of line
* `b`: jump to beginning of word
* `B`: jump to previous whitespace
* `d`: delete mode
* `D`: delete characters
* `e`: jump to end of word
* `E`: jump to next whitespace
* `h`: previous character (like `left`)
* `i`: insert mode
* `I`: insert at beginning of line
* `l`: next character (like `right`)
* `p`: paste after
* `P`: paste before
* `r`: replace character (replace once)
* `R`: replace many characters
* `u`: undo
* `v`: visual editor (opens line in `$EDITOR`)
* `w`: jump to end of word
* `W`: jump to next whitespace
* `x`: delete character
* `y`: yank (copy line)
* `Y`: same as `y`
* `[`: jump to previous brace
* `]`: jump to next brace
* `$`: jump to end of line
* `%`: jump to either end of matching bracket
* `0` to `9`: repeat action _n_ times. eg `5x` would delete (`x`) five (`5`)
  characters

### Recalling Previous Words

* `shift`+`f1` recalls the first word
* `shift`+`f2` recalls the second word
* ...
* `shift`+`f12` recalls the twelfth word

In the following example, code inside square brackets represent key presses
rather than text:

```
» echo two three four five six seven eight nine
two three four five six seven eight nine
» [shift+f1]echo [shift+f5]five
```

## Job Control

While processes are running, the following keys are assigned roles:

* `ctrl`+`c`: kill foreground process. Pressing this will send a kill (SIGINT)
  request to the foreground process

* `ctrl`+`\`: kill all running processes in current shell session, including
  any background processes too. This hotkey is a effectively an emergency kill
  switch to bring you back to the command prompt should `ctrl`+`c` prove
  ineffective. Use this sparingly because it doesn't allow processes to end
  gracefully

* `ctrl`+`z`: suspend foreground process. This will take you back to the prompt
  and from there you can then use job control to resume execution in either the
  foreground or background. ([read more](../commands/fid-list.md))

## Miscellaneous

### Cancel Prompt

Pressing `ctrl`+`c` while on the prompt will clear the prompt. This is similar
to `ctrl`+`u`.

### End Of File

Pressing `ctrl`+`d` on an empty prompt will send EOF (end of file). This will
exit that running shell session.

### Alternative Cancel Key

`ctrl`+`g` performs the same action as `esc` at all states of the interactive
shell.

### Clear Screen

Pressing `ctrl`+`l` will clear the screen.  

### EDITOR

Sometimes you might want to type your command line in a different editor. You
can do via via `esc` followed by `v`.

You will need to have an environmental variable named `$EDITOR` set to the file
name and path of your preferred editor, otherwise Murex will default to `vi`.

(this feature is not currently available on Windows)