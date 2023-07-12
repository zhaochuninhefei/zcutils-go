package zcssh

import (
	"fmt"
	"testing"
)

func Test_executeCommand(t *testing.T) {
	type args struct {
		user     string
		password string
		host     string
		port     string
		command  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				user:     "zhaochun",
				password: "password",
				host:     "localhost",
				port:     "22",
				command:  "free",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExecuteCommand(tt.args.user, tt.args.password, tt.args.host, tt.args.port, tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("executeCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("远程执行结果:\n%s", got)
		})
	}
}

func Test_executeCommands(t *testing.T) {
	type args struct {
		user     string
		password string
		host     string
		port     string
		commands []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				user:     "zhaochun",
				password: "password",
				host:     "localhost",
				port:     "22",
				commands: []string{"free", "cat /asdfasdfasdf", "id", "pwd"},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExecuteCommands(tt.args.user, tt.args.password, tt.args.host, tt.args.port, tt.args.commands)

			if (err != nil) != tt.wantErr {
				t.Errorf("executeCommands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for i, v := range got {
				fmt.Printf("执行命令[%s]结果:\n%s\n", tt.args.commands[i], v)
			}
		})
	}
}
