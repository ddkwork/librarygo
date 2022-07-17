package main

import (
	"github.com/ddkwork/librarygo/src/mycheck"
	"os"
	"path/filepath"
	"strings"

	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream"
	"github.com/ddkwork/librarygo/src/stream/tool"
)

//todo gen sudo sh file
type (
	Interface interface {
		UserHomeDir() string         //账户目录
		GoBin() string               //第一次装系统之后goPath goBin都没被env识别，手动识别了只能
		MakeEnv() string             //生成env.sh
		UpdateBash() string          //俩个用户bash文件每次写入第一行刷新
		LocalEnvFineName() string    //env.sh
		SystemEnvFineName() []string //".zshrc", ".bashrc"
		OpenWineRep() (ok bool)      //开启wine仓库
	}
	object struct{}
)

func New() Interface { return &object{} }

func main() {
	e := New()
	mylog.Info("UserHomeDir", e.UserHomeDir())
	mylog.Info("GoBin", e.GoBin())
	mylog.Info("MakeEnv", e.MakeEnv())
	mylog.Info("env first line", e.UpdateBash())
	e.OpenWineRep()
}

func (o *object) UserHomeDir() string {
	s, err := os.UserHomeDir()
	if !mycheck.Error(err) {
		return ""
	}
	return s
}
func (o *object) GoBin() string { return filepath.Join(o.UserHomeDir(), "go/bin/") }
func (o *object) MakeEnv() string {
	env := "export PATH=${PATH}:" + o.GoBin()
	f, err2 := os.Create(o.LocalEnvFineName())
	if !mycheck.Error(err2) {
		return ""
	}
	if !mycheck.Error2(f.WriteString(env)) {
		return ""
	}
	return env
}
func (o *object) UpdateBash() string {
	abs, err := filepath.Abs("env.sh")
	if !mycheck.Error(err) {
		return ""
	}
	bash := "source  " + abs
	for _, s := range o.SystemEnvFineName() {
		path := o.UserHomeDir() + "/" + s
		buf, err := os.ReadFile(path)
		if err == nil {
			mylog.Info("path", path)
			b := stream.NewBytes(buf)
			lines, ok := tool.File().ToLines(buf)
			if !ok {
				return ""
			}
			if strings.Contains(lines[0], "source") {
				b.Reset()
				if !mycheck.Error2(b.WriteString(strings.Replace(string(buf), lines[0], bash, 1))) {
					return ""
				}
				if !tool.File().WriteTruncate(path, b.Bytes()) {
					return ""
				}
				return lines[0]
			} else {
				NewBuffer := stream.New()
				NewBuffer.WriteStringLn(bash)
				if !mycheck.Error2(NewBuffer.Write(b.Bytes())) {
					return ""
				}
				tool.File().WriteTruncate(path, NewBuffer.Bytes())
				return lines[0]
			}
		}
	}
	return ""
}
func (o *object) LocalEnvFineName() string    { return "env.sh" }
func (o *object) SystemEnvFineName() []string { return []string{".zshrc", ".bashrc"} }

func (o *object) OpenWineRep() (ok bool) {
	pacmanConfName := "/etc/pacman.conf"
	pacmanConfBody, err := os.ReadFile(pacmanConfName)
	if !mycheck.Error(err) {
		return
	}
	lines, b := tool.File().ToLines(pacmanConfBody)
	if !b {
		return
	}
	for i, line := range lines {
		if strings.Contains(line, "#[multilib]") {
			if strings.Contains(lines[i+1], "#Include = /etc/pacman.d/mirrorlist") {
				lines[i] = "[multilib]"
				lines[i+1] = "Include = /etc/pacman.d/mirrorlist"
			}
		}
	}
	body := stream.New()
	for _, line := range lines {
		body.WriteStringLn(line)
	}
	tool.File().WriteTruncate(pacmanConfName, body.Bytes())
	install := `
now you can install wine use theme commands
			sudo pacman -Sy
			yay -S bottles
`
	mylog.Info("install", install)
	return true
}
