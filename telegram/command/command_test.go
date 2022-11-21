package command

import (
	"fmt"
	"strings"
	"testing"
)

func TestDisableCommands(t *testing.T) {
	cleanupForNextTest()
	registerCommands("help", "version")

	if IsCommandDisabled("help") || IsCommandDisabled("version") {
		t.Errorf("IsCommandDisabled() working wrongly, no command was disabled")
	}

	DisableCommands("help")

	if !IsCommandDisabled("help") {
		t.Errorf("IsCommandDisabled() working wrongly, command was disabled")
	}

	if IsCommandDisabled("version") {
		t.Errorf("IsCommandDisabled() working wrongly, command was not disabled")
	}

	// DisableCommands automatically register commands

	DisableCommands("me")

	if !IsSupportCommand("me") {
		t.Errorf("IsCommandDisabled() working wrongly, command should be registered automatically")
	}

	if !IsCommandDisabled("me") {
		t.Errorf("IsCommandDisabled() working wrongly, command was disabled")
	}
}

func TestGetCommandInfo(t *testing.T) {
	cleanupForNextTest()

	RegisterCommand("help", "h", "desc", "argDesc")
	RegisterCommand("version", "v", "", "")
	RegisterCommand("me", "", "d2", "ad2")

	tests := []struct {
		command             string
		wantOriginalCommand string
		wantCommandAlias    string
		wantDesc            string
		wantArgDesc         string
		wantPanic           bool
	}{
		{
			command:             "help",
			wantOriginalCommand: "help",
			wantCommandAlias:    "h",
			wantDesc:            "desc",
			wantArgDesc:         "<argDesc>",
		},
		{
			command:             "h",
			wantOriginalCommand: "help",
			wantCommandAlias:    "h",
			wantDesc:            "desc",
			wantArgDesc:         "<argDesc>",
		},
		{
			command:             "v",
			wantOriginalCommand: "version",
			wantCommandAlias:    "v",
			wantDesc:            "",
			wantArgDesc:         "",
		},
		{
			command:             "me",
			wantOriginalCommand: "me",
			wantCommandAlias:    "",
			wantDesc:            "d2",
			wantArgDesc:         "<ad2>",
		},
		{
			command:   "nonexists",
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.command, func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil && tt.wantPanic {
					t.Errorf("The code did not panic")
				} else if r != nil && !tt.wantPanic {
					t.Errorf("The code should panic")
				}
			}()
			gotOriginalCommand, gotCommandAlias, gotDesc, gotArgDesc := GetCommandInfo(tt.command)
			if gotOriginalCommand != tt.wantOriginalCommand {
				t.Errorf("GetCommandInfo() gotOriginalCommand = %v, want %v", gotOriginalCommand, tt.wantOriginalCommand)
			}
			if gotCommandAlias == nil && len(tt.wantCommandAlias) > 0 {
				t.Errorf("GetCommandInfo() gotCommandAlias = nil, want %v", tt.wantCommandAlias)
			} else if gotCommandAlias != nil && len(tt.wantCommandAlias) < 1 {
				t.Errorf("GetCommandInfo() gotCommandAlias = %v, want nil", gotCommandAlias)
			} else if gotCommandAlias != nil && *gotCommandAlias != tt.wantCommandAlias {
				t.Errorf("GetCommandInfo() gotCommandAlias = %v, want %v", *gotCommandAlias, tt.wantCommandAlias)
			}
			if gotDesc == nil && len(tt.wantDesc) > 0 {
				t.Errorf("GetCommandInfo() gotDesc = nil, want %v", tt.wantDesc)
			} else if gotDesc != nil && len(tt.wantDesc) < 1 {
				t.Errorf("GetCommandInfo() gotDesc = %v, want nil", gotDesc)
			} else if gotDesc != nil && *gotDesc != tt.wantDesc {
				t.Errorf("GetCommandInfo() gotDesc = %v, want %v", *gotDesc, tt.wantDesc)
			}
			if gotArgDesc == nil && len(tt.wantArgDesc) > 0 {
				t.Errorf("GetCommandInfo() gotArgDesc = nil, want %v", tt.wantArgDesc)
			} else if gotArgDesc != nil && len(tt.wantArgDesc) < 1 {
				t.Errorf("GetCommandInfo() gotArgDesc = %v, want nil", gotArgDesc)
			} else if gotArgDesc != nil && *gotArgDesc != tt.wantArgDesc {
				t.Errorf("GetCommandInfo() gotArgDesc = %v, want %v", *gotArgDesc, tt.wantArgDesc)
			}
		})
	}
}

func TestGetRegisteredCommands(t *testing.T) {
	cleanupForNextTest()

	RegisterCommand("help", "", "", "")
	RegisterCommand("version", "", "", "")

	registered := GetRegisteredCommands()
	if len(registered) != 2 {
		t.Errorf("GetRegisteredCommands() returns wrong number of elements %d, want 2", len(registered))
	}
	if registered[0] != "help" || registered[1] != "version" {
		t.Errorf("GetRegisteredCommands() returns wrong order of registered elements")
	}
}

func TestIsCommandDisabled(t *testing.T) {
	cleanupForNextTest()
	registerCommands("help")
	registerCommands("version")
	DisableCommands("help")

	tests := []struct {
		command   string
		want      bool
		wantPanic bool
	}{
		{
			command: "help",
			want:    true,
		},
		{
			command: "version",
			want:    false,
		},
		{
			command:   "me", // panic because me was not registered
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.command, func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil && tt.wantPanic {
					t.Errorf("The code did not panic")
				} else if r != nil && !tt.wantPanic {
					t.Errorf("The code should panic")
				}
			}()
			if got := IsCommandDisabled(tt.command); got != tt.want {
				t.Errorf("IsCommandDisabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSupportCommand(t *testing.T) {
	cleanupForNextTest()

	registerCommands("help", "version")

	tests := []struct {
		command string
		want    bool
	}{
		{
			command: "help",
			want:    true,
		},
		{
			command: "version",
			want:    true,
		},
		{
			command: "me",
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.command, func(t *testing.T) {
			if got := IsSupportCommand(tt.command); got != tt.want {
				t.Errorf("IsSupportCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegisterCommand(t *testing.T) {
	cleanupForNextTest()

	type args struct {
		command string
		alias   string
		desc    string
		argDesc string
	}
	tests := []struct {
		args      args
		wantPanic bool
	}{
		{
			args: args{
				command: "help",
				alias:   "h",
				desc:    "d",
				argDesc: "ad",
			},
			wantPanic: false,
		},
		{
			args: args{
				command: "help",
				alias:   "h",
				desc:    "d",
				argDesc: "ad",
			},
			wantPanic: false, // possible duplicate
		},
		{
			args: args{
				command: "version",
				alias:   "h",
				desc:    "d",
				argDesc: "ad",
			},
			wantPanic: true, // "h" was registered as alias for "help"
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.command, func(t *testing.T) {
			disablePanic := false
			defer func() {
				if !disablePanic {
					r := recover()
					if r == nil && tt.wantPanic {
						t.Errorf("The code did not panic")
					} else if r != nil && !tt.wantPanic {
						t.Errorf("The code should panic")
					}
				}
			}()
			RegisterCommand(tt.args.command, tt.args.alias, tt.args.desc, tt.args.argDesc)
			disablePanic = true
			if len(tt.args.argDesc) > 0 {
				argDesc := tt.args.argDesc
				argDesc = strings.TrimPrefix(argDesc, "<")
				argDesc = strings.TrimSuffix(argDesc, ">")
				_, _, _, ptrArgDesc := GetCommandInfo(tt.args.command)
				if ptrArgDesc == nil {
					t.Errorf("arg desc must exists here")
				}
				if *ptrArgDesc != fmt.Sprintf("<%s>", argDesc) {
					t.Errorf("RegisterCommand does not set prefix and suffix for argDesc quoted with <>")
				}
			}
		})
	}
}

func Test_registerCommands(t *testing.T) {
	cleanupForNextTest()

	registerCommands("help", "version")

	if !IsSupportCommand("help") || !IsSupportCommand("version") {
		t.Errorf("registerCommands did not register command")
	}

	RegisterCommand("me", "m", "", "")

	tests := []string{"m", "1", " 1"}
	for _, command := range tests {
		t.Run(command, func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil {
					t.Errorf("The code did not panic")
				}
			}()
			registerCommands(command)
		})
	}
}
