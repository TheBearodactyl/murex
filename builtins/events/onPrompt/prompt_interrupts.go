package onprompt

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/builtins/events/onPrompt/promptops"
)

var interrupts = []string{
	promptops.Before,
	promptops.After,
	promptops.EOF,
	promptops.Cancel,
}

func isValidInterrupt(interrupt string) error {
	for i := range interrupts {
		if interrupts[i] == interrupt {
			return nil
		}
	}

	return fmt.Errorf("invalid interrupt. Expecting one of the following: %s", strings.Join(interrupts, ", "))
}

func compileInterruptKey(interrupt, name string) string {
	return fmt.Sprintf("%s_%s", interrupt, name)
}

func getInterruptFromKey(key string) []string {
	split := strings.SplitN(key, "_", 2)
	switch len(split) {
	case 2:
		return split
	case 1:
		return []string{"", split[0]}
	default:
		return []string{"", ""}
	}
}
