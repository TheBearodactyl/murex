package lang_test

import (
	"fmt"
	"os"
	"os/user"
	"testing"

	"github.com/lmorg/murex/test"
)

func TestVarSelf(t *testing.T) {
	tests := []test.MurexTest{
		// TTY
		{
			Block: `
				function TestVarSelf {
					$SELF -> [TTY]
				}
				TestVarSelf
			`,
			Stdout: "false",
		},

		// Method
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Method]
				}
				TestVarSelf
			`,
			Stdout: "false",
		},
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Method]
				}
				out foobar -> TestVarSelf
			`,
			Stdout: "true",
		},

		// Not
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Not]
				}
				TestVarSelf
			`,
			Stdout: "false",
		},
		{
			Block: `
				function !TestVarSelf {
					$SELF -> [Not]
				}
				!TestVarSelf
			`,
			Stdout: "true",
		},

		// Background
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Background]
				}
				TestVarSelf
			`,
			Stdout: "false",
		},
		{
			Block: `
				function TestVarSelf {
					$SELF -> [Background]
				}
				bg { TestVarSelf }
				sleep 1
			`,
			Stdout: "true",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestVarArgs(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				function TestVarArgs {
					out @ARGS
				}
				TestVarArgs
			`,
			Stdout: "TestVarArgs\n",
		},
		{
			Block: `
				function TestVarArgs {
					out @ARGS
				}
				TestVarArgs 1 2 3
			`,
			Stdout: "TestVarArgs 1 2 3\n",
		},
		{
			Block: `
				function TestVarArgs {
					out @ARGS
				}
				TestVarArgs 1   2   3
			`,
			Stdout: "TestVarArgs 1 2 3\n",
		},
		{
			Block: `
				function TestVarArgs {
					out $ARGS
				}
				TestVarArgs 1   2   3
			`,
			Stdout: `["TestVarArgs","1","2","3"]` + "\n",
		},
	}

	test.RunMurexTests(tests, t)
}

func TestVarParams(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block: `
				function TestVarParams {
					out $0
				}
				TestVarParams
			`,
			Stdout: "^TestVarParams\n$",
		},
		{
			Block: `
				function TestVarParams {
					out $0
				}
				TestVarParams 1 2 3
			`,
			Stdout: "^TestVarParams\n$",
		},
		{
			Block: `
				function TestVarParams {
					out @PARAMS
				}
				TestVarParams
			`,
			Stderr:  "array '@PARAMS' is empty\n",
			ExitNum: 1,
		},
		{
			Block: `
				function TestVarParams {
					out @PARAMS
				}
				TestVarParams 1 2 3
			`,
			Stdout: "^1 2 3\n$",
		},
		{
			Block: `
				function TestVarParams {
					out @PARAMS
				}
				TestVarParams 1   2   3
			`,
			Stdout: "^1 2 3\n$",
		},
		{
			Block: `
				function TestVarParams {
					out $PARAMS
				}
				TestVarParams 1   2   3
			`,
			Stdout: `\["1","2","3"\]`,
		},
	}

	test.RunMurexTestsRx(tests, t)
}

func TestVarHostname(t *testing.T) {
	hostname := func() string {
		s, err := os.Hostname()
		if err != nil {
			t.Error(err.Error())
		}
		return s
	}

	tests := []test.MurexTest{
		{
			Block:  `out $HOSTNAME`,
			Stdout: fmt.Sprintf("%s\n", hostname()),
		},
	}

	test.RunMurexTests(tests, t)
}

func TestVarPwd(t *testing.T) {
	pwd := func() string {
		s, err := os.Getwd()
		if err != nil {
			t.Error(err.Error())
		}
		return s
	}

	tests := []test.MurexTest{
		{
			Block:  `out $PWD`,
			Stdout: fmt.Sprintf("%s\n", pwd()),
		},
	}

	test.RunMurexTests(tests, t)
}

func TestVarUser(t *testing.T) {
	u, _ := user.Current()
	userName := u.Username
	tests := []test.MurexTest{
		{
			Block:  `$USER`,
			Stdout: userName,
		},
		{
			Block:  `$LOGNAME`,
			Stdout: userName,
		},
		/////
		{
			Block:  `$USER -> debug -> [[/Data-Type/Murex]]`,
			Stdout: `str`,
		},
		{
			Block:  `$LOGNAME-> debug -> [[/Data-Type/Murex]]`,
			Stdout: `str`,
		},
	}

	test.RunMurexTests(tests, t)
}

func TestVarRandom(t *testing.T) {
	tests := []test.MurexTest{
		{
			Block:  `i = $RANDOM; if { $i > 0 && $i < 32768 } then { true } else { false }`,
			Stdout: `true`,
		},
		{
			Block:  `$RANDOM -> debug -> [[/Data-Type/Murex]]`,
			Stdout: `int`,
		},
	}

	test.RunMurexTests(tests, t)
}
