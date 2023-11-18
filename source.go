package main

import (
	"compress/gzip"
	"io"
	"os"

	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/integrations"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/consts"
)

func diskSource(filename string) ([]byte, error) {
	var b []byte

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	if len(filename) > 3 && filename[len(filename)-3:] == ".gz" {
		gz, err := gzip.NewReader(file)
		if err != nil {
			file.Close()
			return nil, err
		}
		b, err = io.ReadAll(gz)

		file.Close()
		gz.Close()

		if err != nil {
			return nil, err
		}

	} else {
		b, err = io.ReadAll(file)
		file.Close()
		if err != nil {
			return nil, err
		}
	}

	return b, nil
}

func execSource(source []rune, sourceRef *ref.Source, exitOnError bool) {
	if sourceRef == nil {
		panic("sourceRef is not defined")
	}

	if debug.Enabled {
		os.Stderr.WriteString("Loading profile `" + sourceRef.Module + "`" + utils.NewLineString)
	}

	var stdin int
	if os.Getenv(consts.EnvMethod) != consts.EnvTrue {
		stdin = lang.F_NO_STDIN
	}
	fork := lang.ShellProcess.Fork(lang.F_PARENT_VARTABLE | stdin)
	fork.Stdout = new(term.Out)
	fork.Stderr = term.NewErr(ansi.IsAllowed())
	fork.FileRef.Source = sourceRef
	fork.RunMode = lang.ShellProcess.RunMode
	exitNum, err := fork.Execute(source)

	if err != nil {
		if exitNum == 0 {
			exitNum = 1
		}
		os.Stderr.WriteString(err.Error() + utils.NewLineString)
		lang.Exit(exitNum)
	}

	if exitNum != 0 && exitOnError {
		lang.Exit(exitNum)
	}
}

func defaultProfile() {
	defaults.AddMurexProfile()

	for _, profile := range defaults.DefaultProfiles {
		ref := ref.History.AddSource("(builtin)", "builtin/"+profile.Name, profile.Block)
		execSource([]rune(string(profile.Block)), ref, false)
	}

	for _, profile := range integrations.Profiles() {
		ref := ref.History.AddSource("(builtin)", "builtin/integrations_"+profile.Name, profile.Block)
		execSource([]rune(string(profile.Block)), ref, false)
	}
}
