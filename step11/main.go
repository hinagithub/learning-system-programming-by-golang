package main
import (
	"fmt"
	"time"
	"bufio"
	"os/exec"
	"io"
	"os"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
	"github.com/kr/pty"
)

func main () {

	//  Main11()
	Main11_7_2()
	
}

// カウントアップexecファイル
func Main11(){
	for i := 0; i <10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// カウントアップexecファイルの実行
func Main11_2(){
	count := exec.Command("./count")
	stdout, _ := count.StdoutPipe()
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan(){
			fmt.Printf("(stdout) %s\n", scanner.Text())
		}
	}()
	err := count.Run()
	if err != nil {
		panic(err)
	}
}

var data = "\033[34m\033[47m\033[4mB\033[31me\n\033[24m\033[30mOS\033[49m\033[m\n"
func Main11_3(){
	var stdOut io.Writer
	if isatty.IsTerminal(os.Stdout.Fd()){
		stdOut = colorable.NewColorableStdout()
	}else{
		stdOut = colorable.NewNonColorable(os.Stdout)
	}
	fmt.Fprintln(stdOut, data)
}

func Main11_7(){
	var out io.Writer
	if isatty.IsTerminal(os.Stdout.Fd()){
		out = colorable.NewColorableStdout()

	}else{
		 out = colorable.NewNonColorable(os.Stdout)
	}

	if isatty.IsTerminal(os.Stdin.Fd()){
		fmt.Fprintln(out, "stdin: terminal")
	} else {
		fmt.Println("stdin: pipe") 
	}

	 if isatty.IsTerminal(os.Stdout.Fd()) {
		 fmt.Fprintln(out, "stdout: terminal")
	 } else {
		 fmt.Println("stdout: pipe")
	 }

	 if isatty.IsTerminal(os.Stdout.Fd()) {
		 fmt.Fprintln(out, "stderr: terminal")
	 } else {
		 fmt.Println("stdout: pipe")
	 }
}

func Main11_7_2(){
	cmd := exec.Command("./check")
	stdpty, stdtty, _ := pty.Open()
	defer stdtty.Close()
	cmd.Stdin = stdpty
	cmd.Stdout = stdpty
	errpty, errtty, _ := pty.Open()
	defer errtty.Close()　
	cmd.Stderr = errtty 
	go func(){
		io.Copy(os.Stdout, stdpty)
	}()
	go func(){
		io.Copy(os.Stdout, errpty)
	}()
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	 
}
