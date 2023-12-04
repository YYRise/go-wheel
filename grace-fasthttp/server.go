package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	gracefulKey        = "GRACEFUL"
	gracefulListenerFd = 3
	serverReadTimeout  = 40 * time.Second
	serverWriteTimeout = 40 * time.Second
)

type serverManager struct {
	sigChan         chan os.Signal
	listener        net.Listener
	server          fasthttp.Server
	err             error
	isGraceFulStart bool
	port            string
}

var (
	sManager = serverManager{}
	err      error
)

func (s *serverManager) serv(port string) {
	fmt.Println("serv start……")
	s.init(port)
	s.initListener()
	s.initServer()

	go s.handleSignals()
	s.closeOldPid()

	fmt.Println("listening on 端口:", port)
	s.server.Serve(s.listener)
}

func (s *serverManager) initServer() {
	s.server = fasthttp.Server{
		Handler:      getRouter().Handler,
		ReadTimeout:  serverReadTimeout,
		WriteTimeout: serverWriteTimeout,
	}
}

func (s *serverManager) initListener() {
	if s.isGraceFulStart {
		file := os.NewFile(gracefulListenerFd, "")
		fmt.Println("initListener:file=", file)
		s.listener, err = net.FileListener(file)
	} else {
		s.listener, err = net.Listen("tcp", ":"+s.port)
	}
	if err != nil {
		panic("监听端口报错:" + err.Error())
	}
}

func (s *serverManager) init(port string) {
	s.sigChan = make(chan os.Signal)
	s.port = port

	for i := range os.Args {
		fmt.Println("init: start", os.Args[i])
		if os.Args[i] == gracefulKey {
			s.isGraceFulStart = true
			break
		}
	}
}

func (s *serverManager) handleSignals() {
	// 监听
	signal.Notify(s.sigChan, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)

	var (
		sig os.Signal
	)
	pid := syscall.Getpid()
	for {
		sig = <-s.sigChan
		switch sig {
		case syscall.SIGHUP:
			fmt.Println(pid, "Received SIGHUP. forking.")
			err := s.fork()
			if err != nil {
				panic("Fork err:" + err.Error())
			}
		case syscall.SIGTERM:
			fmt.Println(pid, "Received SIGTERM. 开始关闭")
			err := s.server.Shutdown()
			if err != nil {
				panic("shutdown err:" + err.Error())
			}
			fmt.Println(pid, "Received SIGTERM. 结束关闭")
		default:
			fmt.Printf("Received %v: nothing i care about...\n", sig)
		}
	}
}

func (s *serverManager) fork() (err error) {
	listenerFd, err := s.listener.(*net.TCPListener).File()
	if err != nil {
		panic("failed to get socket file descriptor: " + err.Error())
	}

	var args []string
	for _, arg := range os.Args {
		if arg == gracefulKey {
			break
		}
		fmt.Println("arg:", arg)
		args = append(args, arg)
	}
	args = append(args, gracefulKey)

	envList := []string{
		`CGO_LDFLAGS="-g -O2"`,
		`LD_LIBRARY_PATH=` + filepath.Dir(os.Args[0]) + "/../lib",
	}

	allFiles := []*os.File{os.Stdin, os.Stdout, os.Stderr, listenerFd}
	execSpec := &os.ProcAttr{
		Env:   envList,
		Files: allFiles,
	}

	var (
		pros *os.Process
		//prosState *os.ProcessState
	)
	pros, err = os.StartProcess(os.Args[0], args, execSpec)
	if err != nil {
		fmt.Println("os.StartProcess err: " + err.Error())
		return fmt.Errorf("failed to fork exec: %v", err)
	}

	fmt.Println(fmt.Sprintf("pros: %#v\n", *pros))

	return nil
}

func (s *serverManager) closeOldPid() {
	if s.isGraceFulStart {
		process, err := os.FindProcess(os.Getppid())
		if err != nil {
			panic(err)
		}
		err = process.Signal(syscall.SIGTERM)
		fmt.Println("closeOldPid 发送信号:", process.Pid, syscall.SIGTERM.String())
		if err != nil {
			fmt.Println("closeOldPid 信号错误:", err.Error())
			panic(err)
		}
	}
}
