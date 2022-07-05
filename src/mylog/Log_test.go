package mylog

import (
	"fmt"
	"testing"
)

//https://gitee.com/inhere/color/
func TestLog(t *testing.T) {
	p := New()
	p.Hex("hex value", 0x888)
	p.HexDump("buf", []byte{0x11, 0x22, 0x33, 0x44})
	p.Struct(struct {
		a int
		b string
		c []byte
	}{
		a: 89,
		b: "jhjsbdd",
		c: []byte{0x11, 0x22, 0x33, 0x44},
	})
	p.Info("infomation", "tttttttttttttttttttttttt")
	p.Trace("trace", "kkkkkkkkkkkkkkkkkkkk")
	//check.Error("this is a error")
	p.Warning("warnning", "mmmmmmmmm")
	p.Success("Success", "vgoTest pass")
	p.Success("中文 Success", "vgoTest pass")
	p.Success("我是中文 Success", "vgoTest pass")
	p.Hex("firstEnd xor 0x72B8,机器码还有一个字节是丢弃的", 0x72B8)
	return
	co()
	hico()
}

func co() {
	fmt.Printf("\x1b[30m%s\x1b[0m", "Black") //?
	fmt.Println()
	fmt.Printf("\x1b[31m%s\x1b[0m", "RedS")
	fmt.Println()
	fmt.Printf("\x1b[32m%s\x1b[0m", "Green")
	fmt.Println()
	fmt.Printf("\x1b[33m%s\x1b[0m", "Yellow") //?
	fmt.Println()
	fmt.Printf("\x1b[34m%s\x1b[0m", "Blue") //?
	fmt.Println()
	fmt.Printf("\x1b[35m%s\x1b[0m", "Magenta")
	fmt.Println()
	fmt.Printf("\x1b[36m%s\x1b[0m", "Cyan")
	fmt.Println()
	fmt.Printf("\x1b[37m%s\x1b[0m", "White") //?
	fmt.Println()
}

func hico() {
	fmt.Println()
	fmt.Println()
	fmt.Printf("\x1b[90m%s\x1b[0m", "HiBlack")
	fmt.Println()
	fmt.Printf("\x1b[91m%s\x1b[0m", "HiRedS")
	fmt.Println()
	fmt.Printf("\x1b[92m%s\x1b[0m", "HiGreen") //?
	fmt.Println()
	fmt.Printf("\x1b[93m%s\x1b[0m", "HiYellow") //?
	fmt.Println()
	fmt.Printf("\x1b[94m%s\x1b[0m", "HiBlue") //? 明显是浅蓝色
	fmt.Println()
	fmt.Printf("\x1b[95m%s\x1b[0m", "HiMagenta")
	fmt.Println()
	fmt.Printf("\x1b[96m%s\x1b[0m", "HiCyan")
	fmt.Println()
	fmt.Printf("\x1b[97m%s\x1b[0m", "HiWhite") //?
}

const (
	Black = iota + 30
	RedS
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

const (
	HiBlack = iota + 90
	HiRed
	HiGreen
	HiYellow
	HiBlue
	HiMagenta
	HiCyan
	HiWhite
)
