package command

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"strings"
)

var registeredAliases = cmap.New[string]()

// registerCommandAlias registers command with corresponding alias
func registerCommandAlias(command string, alias string) {
	if !registeredCommands.Has(command) {
		panic(fmt.Errorf("command [%s] has not been registerd", command))
	}
	if !commandRegex.MatchString(alias) {
		panic(fmt.Errorf("alias [%s] format is not well-formed", alias))
	}
	if registeredCommands.Has(alias) {
		panic(fmt.Errorf("[%s] had been registered as a command thus can not be an alias", alias))
	}
	if alias == command {
		panic(fmt.Errorf("can not register [%s] as alias for [%s]", alias, command))
	}
	if mappedCommand, found := registeredAliases.Get(alias); found {
		if mappedCommand != command {
			panic(fmt.Errorf("alias [%s] had been registered as alias for command [%s]", alias, mappedCommand))
		}
	} else {
		registeredAliases.Set(alias, command)
	}
}

// TranslateCommandIfAlias returns original full-sized command if the input command is an alias
func TranslateCommandIfAlias(maybeCommandAlias string) string {
	if strings.HasPrefix(maybeCommandAlias, "/") {
		maybeCommandAlias = strings.TrimPrefix(maybeCommandAlias, "/")
	}
	if command, found := registeredAliases.Get(maybeCommandAlias); found {
		return command
	}
	return maybeCommandAlias
}
