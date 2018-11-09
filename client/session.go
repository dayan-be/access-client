package client

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

type Option func(*Options)

type Options struct {
	timeout time.Duration
}

const (
	MSG_READ_SIZE   = 4096
	MSG_BUFFER_SIZE = 10240
)

var (
	DefaultDialTimeout = time.Second * 5
)

func Timeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.timeout = timeout
	}
}

type Session struct {
	conn net.Conn
	messageCallback func(context.Context, []byte) error
}

func NewSession(addr string, opts ...Option) (*Session, error) {
	opt := &Options{}
	for _, o := range opts {
		o(opt)
	}
	if opt.timeout == 0 {
		opt.timeout = DefaultDialTimeout
	}

	conn, err := net.DialTimeout("tcp", addr, opt.timeout)
	if err != nil {
		return nil, err
	}

	return &Session{
		conn:conn,
	}, nil
}

func (s *Session) Auth(phoneNum, password, token string) error {
	if token == "" {

	} else {

	}

	return nil
}

func (s *Session) Run() {
	ctx := context.Background()
	msgBuf := bytes.NewBuffer(make([]byte, 0, MSG_BUFFER_SIZE))
	// 数据缓冲
	dataBuf := make([]byte, MSG_READ_SIZE)
	// 消息长度
	length := 0
	// 消息长度uint32
	uLen := uint32(0)
	msgFlag := ""

	for {
		// 读取数据
		n, err := s.conn.Read(dataBuf)
		if err == io.EOF {
			fmt.Printf("Session exit: %s\n", ss.socket.RemoteAddr())
		}
		if err != nil {
			fmt.Printf("Read error: %s\n", err)
			return
		}
		//fmt.Println(dataBuf[:n])
		// 数据添加到消息缓冲
		n, err = msgBuf.Write(dataBuf[:n])
		if err != nil {
			fmt.Printf("Buffer write error: %s\n", err)
			return
		}

		// 消息分割循环
		for {
			// 消息头
			if length == 0 && msgBuf.Len() >= 6 {
				msgFlag = string(msgBuf.Next(2))
				if msgFlag != "DY" {
					fmt.Printf("invalid message")
					return
				}
				binary.Read(msgBuf, binary.LittleEndian, &uLen)
				length = int(uLen)
				// 检查超长消息
				if length > MSG_BUFFER_SIZE {
					fmt.Printf("Message too length: %d\n", length)
					return
				}
			}
			// 消息体
			if length > 0 && msgBuf.Len() >= length {
				length = 0
				go s.messageCallback(ctx, msgBuf.Next(length))
			} else {
				break
			}
		}
	}
}

func (s *Session) SetMessageCallback(fun func(ctx context.Context, msg []byte) error ){
	s.messageCallback = fun
}