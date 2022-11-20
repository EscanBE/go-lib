package command

import (
	"fmt"
	"github.com/EscanBE/go-lib/utils"
	cmap "github.com/orcaman/concurrent-map/v2"
	"regexp"
	"strings"
)

var registeredCommands = cmap.New[bool]()
var commandRegex = regexp.MustCompile("^[a-z][a-z\\d_]*$")
var registeredCommandsInOrder = make([]string, 0)

// registerCommands puts the registration command into registry
func registerCommands(commands ...string) {
	for _, command := range commands {
		if !commandRegex.MatchString(command) {
			panic(fmt.Errorf("command [%s] format is not well-formed", command))
		}
		if registeredAliases.Has(command) {
			panic(fmt.Errorf("[%s] had been registered as a command alias thus can not be a command", command))
		}
		registeredCommands.Set(command, true)
		registeredCommandsInOrder = append(registeredCommandsInOrder, command)
	}
}

// GetRegisteredCommands returns a list of registered commands
func GetRegisteredCommands() []string {
	return registeredCommandsInOrder
}

var disabledCommands = cmap.New[bool]()

// DisableCommands disables a specific or a list of command
func DisableCommands(commands ...string) {
	registerCommands(commands...)

	for _, command := range commands {
		disabledCommands.Set(command, true)
	}
}

// IsSupportCommand returns true if the command (or alias) was registered
func IsSupportCommand(command string) bool {
	return registeredCommands.Has(command) || registeredAliases.Has(command)
}

// IsCommandDisabled returns true of the command (or alias) was disabled
func IsCommandDisabled(command string) bool {
	if !IsSupportCommand(command) {
		panic(fmt.Errorf("[%s] is not a supported command", command))
	}

	cmd := TranslateCommandIfAlias(command)

	if disabled, found := disabledCommands.Get(cmd); found {
		return disabled
	}

	return false
}

var commandsDescription = cmap.New[string]()
var commandsArgDesc = cmap.New[string]()

// RegisterCommand performs registration the command with associated alias, description and argument description
func RegisterCommand(command, alias, desc, argDesc string) {
	registerCommands(command)
	if !utils.IsBlank(alias) {
		registerCommandAlias(command, strings.TrimSpace(alias))
	}
	if !utils.IsBlank(desc) {
		commandsDescription.Set(command, strings.TrimSpace(desc))
	}
	if !utils.IsBlank(argDesc) {
		argDesc = strings.TrimSpace(argDesc)
		if !strings.HasPrefix(argDesc, "<") {
			argDesc = "<" + argDesc
		}
		if !strings.HasSuffix(argDesc, ">") {
			argDesc += ">"
		}
		commandsArgDesc.Set(command, argDesc)
	}
}

// GetCommandInfo returns original command with associated alias, description, argument description if the command was registered
func GetCommandInfo(command string) (originalCommand string, commandAlias *string, desc *string, argDesc *string) {
	var _originalCommand string
	var _commandAlias, _desc, _argDesc *string
	_originalCommand = TranslateCommandIfAlias(command)
	if !IsSupportCommand(_originalCommand) {
		panic(fmt.Errorf("[%s] is not a supported command", _originalCommand))
	}
	for alias, cmd := range registeredAliases.Items() {
		if cmd == _originalCommand {
			_commandAlias = &alias
			break
		}
	}
	if d, found := commandsDescription.Get(_originalCommand); found {
		_desc = &d
	}
	if ad, found := commandsArgDesc.Get(_originalCommand); found {
		_argDesc = &ad
	}
	return _originalCommand, _commandAlias, _desc, _argDesc
}
