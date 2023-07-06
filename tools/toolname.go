package toolname

import (
	"fmt"
	"strings"
)

type Toolname struct {
	Scope string
	Name  string

	Inner string

	IsValid bool
	Err     string
}

func checkIntent(intent string) (bool, string) {
	if len(intent) == 0 || intent == "" {
		return false, "%s can not be empty"
	}

	if strings.Contains(intent, "/") {
		return false, "%s can not contain '/'"
	}

	return true, ""
}

func New(input string) Toolname {
	var toolname Toolname

	input = strings.TrimSpace(input)

	split := strings.SplitAfter(input, "/")

	if len(split) != 2 {
		toolname.IsValid = false
		toolname.Err = "Toolname must be in the format of 'scope/name'"
		return toolname
	}

	scope := split[0][:len(split[0])-1]
	name := split[1]

	toolname.IsValid = true

	isValid, errMessage := checkIntent(scope)

	if !isValid {
		toolname.IsValid = false
		toolname.Err = fmt.Sprintf(errMessage, "Scope")

		return toolname
	}

	isValid, errMessage = checkIntent(name)

	if !isValid {
		toolname.IsValid = false
		toolname.Err = fmt.Sprintf(errMessage, "Name")

		return toolname
	}

	toolname.Name = name
	toolname.Scope = scope

	inner := fmt.Sprintf("%s/%s", scope, name)

	toolname.Inner = inner

	return toolname
}
