package command

import (
	"testing"
)

func cleanupForNextTest() {
	registeredCommands.Clear()
	registeredAliases.Clear()
	registeredCommandsInOrder = make([]string, 0)
}

func TestTranslateCommandIfAlias(t *testing.T) {
	cleanupForNextTest()
	RegisterCommand("help", "h", "", "")
	RegisterCommand("version", "v", "", "")

	tests := []struct {
		maybeCommandAlias string
		want              string
	}{
		{
			maybeCommandAlias: "/h",
			want:              "help",
		},
		{
			maybeCommandAlias: "h",
			want:              "help",
		},
		{
			maybeCommandAlias: "v",
			want:              "version",
		},
		{
			maybeCommandAlias: "d",
			want:              "d",
		},
		{
			maybeCommandAlias: "help",
			want:              "help",
		},
	}
	for _, tt := range tests {
		t.Run(tt.maybeCommandAlias, func(t *testing.T) {
			if got := TranslateCommandIfAlias(tt.maybeCommandAlias); got != tt.want {
				t.Errorf("TranslateCommandIfAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_registerCommandAlias(t *testing.T) {
	cleanupForNextTest()

	extraCommand := "version"
	tests := []struct {
		name                 string
		command              string
		alias                string
		register             bool
		registerExtraCommand bool
		wantPanic            bool
	}{
		{
			name:      "not registered command can not register for alias",
			command:   "help",
			alias:     "h",
			register:  false,
			wantPanic: true,
		},
		{
			name:      "success",
			command:   "help",
			alias:     "h",
			register:  true,
			wantPanic: false,
		},
		{
			name:      "success",
			command:   extraCommand,
			alias:     "v",
			register:  true,
			wantPanic: false,
		},
		{
			name:      "invalid alias format",
			command:   "help",
			alias:     "1",
			register:  true,
			wantPanic: true,
		},
		{
			name:      "invalid alias format",
			command:   "help",
			alias:     "h@",
			register:  true,
			wantPanic: true,
		},
		{
			name:                 "registered command can not be alias",
			command:              "help",
			alias:                extraCommand,
			register:             true,
			registerExtraCommand: true,
			wantPanic:            true,
		},
		{
			name:      "alias and command can not be the same",
			command:   extraCommand,
			alias:     extraCommand,
			register:  true,
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if r == nil && tt.wantPanic {
					t.Errorf("The code did not panic")
				} else if r != nil && !tt.wantPanic {
					t.Errorf("The code should panic")
				}
			}()
			if tt.register {
				registerCommands(tt.command)
			}

			if tt.registerExtraCommand {
				registerCommands(extraCommand)
			}

			registerCommandAlias(tt.command, tt.alias)

			cleanupForNextTest()
		})
	}
}

func Test_registerCommandAlias2(t *testing.T) {
	registerCommands("help", "version")
	registerCommandAlias("help", "h")

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	// expect panic because "h" was previously registered as alias for "help"
	registerCommandAlias("version", "h")
}
