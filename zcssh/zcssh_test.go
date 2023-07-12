package zcssh

import (
	"fmt"
	"log"
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
				password: "asdfzxcv123",
				host:     "localhost",
				port:     "22",
				command:  "test",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := executeCommand(tt.args.user, tt.args.password, tt.args.host, tt.args.port, tt.args.command)
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
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				user:     "zhaochun",
				password: "asdfzxcv123",
				host:     "localhost",
				port:     "22",
				commands: []string{"free", "whoami", "pwd"},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := executeCommands(tt.args.user, tt.args.password, tt.args.host, tt.args.port, tt.args.commands)
			//got, err := executeCommands(tt.args.user, tt.args.password, tt.args.host, tt.args.commands)

			if (err != nil) != tt.wantErr {
				t.Errorf("executeCommands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//fmt.Printf("远程执行结果:\n%s", got)
			for i, v := range got {
				//fmt.Println(v)
				fmt.Printf("执行命令[%s]结果:\n%s\n", tt.args.commands[i], v)
			}
		})
	}
}

func TestRemoteRun(t *testing.T) {
	// Run multiple commands on the same session
	commands := []string{
		//"ls -l",
		//"test",
		"pwd",
		"whoami",
		"echo hello world",
	}

	// Call the RemoteRun function with the target server information and the command list
	results, err := RemoteRun("localhost", "22", "zhaochun", "asdfzxcv123", commands)

	if err != nil {
		log.Fatal(err)
	}

	// Print the results
	for _, result := range results {
		fmt.Println(result)
	}
}
