package zcssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
)

func executeCommand(user, password, host, port, command string) (string, error) {
	// SSH客户端配置
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// SSH连接
	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		return "", fmt.Errorf("failed to dial SSH: %v", err)
	}
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {
			fmt.Printf("zcssh.executeCommand 关闭ssh客户端发生错误: %s\n", err.Error())
		}
	}(client)

	// 创建会话
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create SSH session: %v", err)
	}
	defer func(session *ssh.Session) {
		err := session.Close()
		if err != nil {
			if err == io.EOF {
				fmt.Println("zcssh.executeCommand 关闭ssh会话成功")
			} else {
				fmt.Printf("zcssh.executeCommand 关闭ssh会话发生错误: %s\n", err.Error())
			}
		}
	}(session)

	// 执行命令
	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v", err)
	}

	return string(output), nil
}

func executeCommands(user, password, host, port string, commands []string) ([]string, error) {
	// SSH客户端配置
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// SSH连接
	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial SSH: %v", err)
	}
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {
			fmt.Printf("zcssh.executeCommands 关闭ssh客户端发生错误: %s\n", err.Error())
		}
	}(client)

	// 创建会话
	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create SSH session: %v", err)
	}
	defer func(session *ssh.Session) {
		err := session.Close()
		if err != nil {
			if err == io.EOF {
				fmt.Println("zcssh.executeCommands 关闭ssh会话成功")
			} else {
				fmt.Printf("zcssh.executeCommands 关闭ssh会话发生错误: %s\n", err.Error())
			}
		}
	}(session)

	// 执行命令
	var results []string
	for _, command := range commands {
		output, err := session.CombinedOutput(command)
		if err != nil {
			return nil, fmt.Errorf("failed to execute command: %v", err)
		}
		results = append(results, string(output))
	}
	return results, nil
}
