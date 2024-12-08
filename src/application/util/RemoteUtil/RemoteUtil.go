package RemoteUtil

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/aymanbagabas/go-pty"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-cmd/cmd"
	"io"
	"os"
	"os/exec"
	"script-go/src/application/pojo/dto/Log"
	"script-go/src/application/util/LogUtil"
	"sync"
)

func ChangeWorkFolder(home string) {
	err := os.Chdir(home)
	if err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ChangeWorkFolder", "os.Chdir", err))
	}
}

func ExecLocalCmdByPty(bin string, args ...string) {
	ptyClient, err := pty.New()
	if err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdByPty", "pty.New", err))
	}

	defer ptyClient.Close()
	c := ptyClient.Command(bin, args...)
	if err := c.Start(); err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdByPty", "ptyClient.Command", err))
	}

	go io.Copy(os.Stdout, ptyClient)

	if err = c.Wait(); err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdByPty", "c.Wait", err))
	}
}

func ExecLocalCmdByTea(bin string, args ...string) {
	if _, err := tea.NewProgram(newExecModel(bin, args)).Run(); err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdByTea", "tea.NewProgram", err))
	}
}

func ExecLocalCmd(bin string, args ...string) {
	envCmd := cmd.NewCmd(bin, args...)

	status := <-envCmd.Start()

	for _, line := range status.Stdout {
		fmt.Println(line)
	}

	for _, line := range status.Stderr {
		fmt.Println(line)
	}
}

func ExecLocalCmdByStream(bin string, args ...string) {
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	envCmd := cmd.NewCmdOptions(cmdOptions, bin, args...)
	doneChan := make(chan struct{})
	go func() {
		defer close(doneChan)
		for envCmd.Stdout != nil || envCmd.Stderr != nil {
			select {
			case line, open := <-envCmd.Stdout:
				if !open {
					envCmd.Stdout = nil
					continue
				}
				fmt.Println(line)
			case line, open := <-envCmd.Stderr:
				if !open {
					envCmd.Stderr = nil
					continue
				}
				_, _ = fmt.Fprintln(os.Stderr, line)
			}
		}
	}()

	<-envCmd.Start()

	<-doneChan
}

func ExecLocalCmdByAsyncTest(bin string, args ...string) {
	terminal := exec.Command(bin, args...)

	stdout, err := terminal.StdoutPipe()
	if err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdByAsyncTest", "terminal.StdoutPipe", err))
	}

	var wg sync.WaitGroup
	wg.Add(1)

	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			fmt.Print(scanner.Text())
		}
		wg.Done()
	}()

	if err = terminal.Start(); err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdByAsyncTest", "terminal.Start", err))
	}

	wg.Wait()

	if err = terminal.Wait(); err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdByAsyncTest", "terminal.Wait", err))
	}

}

func ExecLocalCmdByAsync(bin string, args ...string) {
	terminal := exec.Command(bin, args...)
	terminal.Stdin = os.Stdin

	var wg sync.WaitGroup
	wg.Add(2)

	stdout, err := terminal.StdoutPipe()
	if err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdByAsync", "terminal.StdoutPipe", err))
		_ = terminal.Process.Kill()
	}
	readOut := bufio.NewReader(stdout)
	go func() {
		defer wg.Done()
		GetOutput(readOut)
	}()

	stderr, err := terminal.StderrPipe()
	if err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdByAsync", "terminal.StderrPipe", err))
		_ = terminal.Process.Kill()
	}
	readErr := bufio.NewReader(stderr)
	go func() {
		defer wg.Done()
		GetOutput(readErr)
	}()

	err = terminal.Run()
	if err != nil {
		return
	}

	wg.Wait()
}

func GetOutput(reader *bufio.Reader) {
	var sumOutput string
	outputBytes := make([]byte, 200)
	for {
		n, err := reader.Read(outputBytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			sumOutput += err.Error()
		}
		output := string(outputBytes[:n])

		fmt.Print(output)
		sumOutput += output
	}
}

func ExecLocalCmdSimpleTest(bin string, args ...string) {
	terminal := exec.Command(bin, args...)
	terminal.Stdin = os.Stdin

	stdout, err := terminal.StdoutPipe()
	if err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdSimpleTest", "terminal.StdoutPipe", err))
		_ = terminal.Process.Kill()
	}
	go io.Copy(os.Stdout, stdout)

	stderr, err := terminal.StderrPipe()
	if err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdSimpleTest", "terminal.StderrPipe", err))
		_ = terminal.Process.Kill()
	}
	go io.Copy(os.Stdout, stderr)

	if err = terminal.Start(); err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdSimpleTest", "terminal.Start", err))
	}

	if err = terminal.Wait(); err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdSimpleTest", "terminal.Wait", err))
	}
}

func ExecLocalCmdSimple(bin string, args ...string) {
	terminal := exec.Command(bin, args...)
	var stdout, stderr bytes.Buffer
	terminal.Stdout = &stdout
	terminal.Stderr = &stderr
	err := terminal.Run()

	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if len(outStr) > 0 {
		fmt.Println(outStr)
	}
	if len(errStr) > 0 {
		fmt.Println(errStr)
	}
	if err != nil {
		LogUtil.LoggerLine(Log.Of("RemoteUtil", "ExecLocalCmdSimple", "terminal.Run", err))
	}
}
