package cmd

import (
	"bytes"
	"context"
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Session struct {
	ShowLog    bool
	ShowStdOut bool
	dir        string
	pid        chan int
	logWriter  io.Writer
}

func Run(command string) (string, error) {
	return NewSession().run(context.Background(), command, true)
}

func NewSession() *Session             { return &Session{pid: make(chan int, 1)} }
func (s *Session) SetDir(dir string)   { s.dir = strings.TrimSpace(dir) }
func (s *Session) SetLog(wr io.Writer) { s.logWriter = wr }
func (s *Session) GetPid() <-chan int  { return s.pid }
func (s *Session) run(ctx context.Context, command string, disableStybel bool) (string, error) {
	if s.ShowLog {
		log.SetPrefix("go-command: ")
		if s.logWriter != nil {
			log.SetOutput(s.logWriter)
		}
	}

	var cmdSlice []string
	if disableStybel {
		cmdSlice = append(cmdSlice, command)
	} else {
		cmdSlice = append(cmdSlice, strings.Split(command, " ")...)
	}

	//
	var cmd *exec.Cmd
	var err error
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", command)
	case "linux", "darwin", "freebsd":
		cmd = exec.Command("bash", "-c", command)
	default:

	}
	cmd.Dir = s.dir
	log.Println(cmd.String())

	outputErr := &bytes.Buffer{}
	outputOut := &bytes.Buffer{}

	if s.ShowStdOut {
		cmd.Stderr = io.MultiWriter(outputErr, os.Stderr)
		cmd.Stdout = io.MultiWriter(outputOut, os.Stdout)
	} else {
		cmd.Stderr = io.MultiWriter(outputErr)
		cmd.Stdout = io.MultiWriter(outputOut)
	}
	err = cmd.Start()
	if err != nil {
		return "", err
	}
	s.pid <- cmd.Process.Pid

	//
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	done := make(chan struct{}, 1)
	go func() {
		select {
		case <-ctx.Done():
			log.Printf("%s , err = %v", cmd.String(), ctx.Err())
			cmd.Process.Kill()
		case <-done:
		}
	}()

	//
	err = cmd.Wait()
	done <- struct{}{}
	if err != nil {
		return "", errors.New(ConvertByte2String(outputErr.Bytes(), GB18030))
	}
	return ConvertByte2String(outputOut.Bytes(), GB18030), nil
}

func Kill(pid int) error {
	return nil
	//return syscall.Kill(pid, syscall.SIGKILL)
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
