package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"syscall"
)

func main() {
	lfd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := syscall.Bind(lfd, &syscall.SockaddrInet4{Port: 9999}); err != nil {
		fmt.Println(err)
		return
	}
	if err := syscall.Listen(lfd, 1024); err != nil {
		fmt.Println(err)
		return
	}
	e, err := syscall.EpollCreate1(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := syscall.EpollCtl(e, syscall.EPOLL_CTL_ADD, lfd, &syscall.EpollEvent{
		Events: syscall.EPOLLIN | syscall.EPOLLPRI | syscall.EPOLLERR | syscall.EPOLLHUP | unix.EPOLLET,
		Fd:     int32(lfd),
	}); err != nil {
		fmt.Println(err)
		return
	}
	events := make([]syscall.EpollEvent, 100)
	eventQueue := make(chan syscall.EpollEvent, 1000)
	go func() {
		for {
			n, err := syscall.EpollWait(e, events, -1)
			if err != nil {
				fmt.Println(err)
				return
			}
			for i := 0; i < n; i++ {
				if events[i].Fd == int32(lfd) {
					eventQueue <- events[i]

				} else {
					fmt.Println("message from: ", events[i].Fd)

				}
			}

		}
	}()
	go func() {
		for value := range eventQueue {
			fmt.Println(value)
			nfd, sa, err2 := syscall.Accept(int(value.Fd))
			if err2 != nil {
				fmt.Println(err2)
				return
			}
			fmt.Println(nfd)
			fmt.Println(sa)
			if err := syscall.EpollCtl(e, syscall.EPOLL_CTL_ADD, nfd, &syscall.EpollEvent{
				Events: syscall.EPOLLIN | syscall.EPOLLPRI | syscall.EPOLLERR | syscall.EPOLLHUP | unix.EPOLLET,
				Fd:     int32(nfd),
			}); err != nil {
				fmt.Println(err)
				return
			}
		}
	}()
	select {}
}
