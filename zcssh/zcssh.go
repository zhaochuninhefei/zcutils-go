package zcssh

import (
	"bytes"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"strings"
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

	// 拼接命令
	command := strings.Join(commands, fmt.Sprintf(" && echo %s && ", cmdSeprator))

	// 执行命令
	output, err := session.CombinedOutput(command)
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %v", err)
	}
	// 将output转为string
	outputStr := strings.TrimSpace(string(output))
	fmt.Println(outputStr)
	// 使用cmdSeprator对outputStr做分割,得到一个切片
	outputStrs := strings.Split(outputStr, cmdSeprator+"\n")

	return outputStrs, nil
}

const cmdSeprator = "=====Command Done====="

// RemoteRun runs multiple commands on the same session and returns the results as a slice of strings
func RemoteRun(host, port, user, password string, commands []string) ([]string, error) {
	// Create a new SSH client
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // ignore host key verification for simplicity
	}
	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// Create a buffer to store the output
	var output bytes.Buffer

	// Set IO
	session.Stdout = &output
	session.Stderr = &output

	// Get the stdin pipe
	stdin, err := session.StdinPipe()
	if err != nil {
		return nil, err
	}

	// Start remote shell
	err = session.Shell()
	if err != nil {
		return nil, err
	}

	// Send the commands
	for i, cmd := range commands {
		_, err = fmt.Fprintf(stdin, "%s\n", cmd)
		if err != nil {
			return nil, err
		}
		if i < len(commands)-1 {
			_, err = fmt.Fprintf(stdin, "echo %s\n", cmdSeprator)
			if err != nil {
				return nil, err
			}
		}
	}

	// Close the stdin pipe to send EOF signal to the session
	stdin.Close()

	// Wait for session to finish
	err = session.Wait()
	if err != nil {
		return nil, err
	}

	// 将output转为string
	outputStr := strings.TrimSpace(output.String())
	// 使用cmdSeprator对outputStr做分割,得到一个切片
	outputStrs := strings.Split(outputStr, "\n"+cmdSeprator+"\n")

	//// Split the output by newline and return as a slice of strings
	//results := bytes.Split(output.Bytes(), []byte("\n"))
	//
	//// Convert each byte slice to a string
	//var resultsStr []string
	//for _, result := range results {
	//	resultsStr = append(resultsStr, string(result))
	//}

	return outputStrs, nil
}
