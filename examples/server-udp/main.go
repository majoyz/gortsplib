package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/majoyz/gortsplib"
	"github.com/majoyz/gortsplib/pkg/base"
)

// This example shows how to
// 1. create a RTSP server which accepts plain connections
// 2. allow a single client to publish a stream with TCP or UDP
// 3. allow multiple clients to read that stream with TCP or UDP

var mutex sync.Mutex
var publisher *gortsplib.ServerConn
var readers = make(map[*gortsplib.ServerConn]struct{})
var sdp []byte

// this is called for each incoming connection
func handleConn(conn *gortsplib.ServerConn) {
	defer conn.Close()

	log.Printf("client connected")

	// called after receiving a DESCRIBE request.
	onDescribe := func(ctx *gortsplib.ServerConnDescribeCtx) (*base.Response, []byte, error) {
		mutex.Lock()
		defer mutex.Unlock()

		// no one is publishing yet
		if publisher == nil {
			return &base.Response{
				StatusCode: base.StatusNotFound,
			}, nil, nil
		}

		return &base.Response{
			StatusCode: base.StatusOK,
		}, sdp, nil
	}

	// called after receiving an ANNOUNCE request.
	onAnnounce := func(ctx *gortsplib.ServerConnAnnounceCtx) (*base.Response, error) {
		mutex.Lock()
		defer mutex.Unlock()

		if publisher != nil {
			return &base.Response{
				StatusCode: base.StatusBadRequest,
			}, fmt.Errorf("someone is already publishing")
		}

		publisher = conn
		sdp = ctx.Tracks.Write()

		return &base.Response{
			StatusCode: base.StatusOK,
			Header: base.Header{
				"Session": base.HeaderValue{"12345678"},
			},
		}, nil
	}

	// called after receiving a SETUP request.
	onSetup := func(ctx *gortsplib.ServerConnSetupCtx) (*base.Response, error) {
		return &base.Response{
			StatusCode: base.StatusOK,
			Header: base.Header{
				"Session": base.HeaderValue{"12345678"},
			},
		}, nil
	}

	// called after receiving a PLAY request.
	onPlay := func(ctx *gortsplib.ServerConnPlayCtx) (*base.Response, error) {
		mutex.Lock()
		defer mutex.Unlock()

		readers[conn] = struct{}{}

		return &base.Response{
			StatusCode: base.StatusOK,
			Header: base.Header{
				"Session": base.HeaderValue{"12345678"},
			},
		}, nil
	}

	// called after receiving a RECORD request.
	onRecord := func(ctx *gortsplib.ServerConnRecordCtx) (*base.Response, error) {
		mutex.Lock()
		defer mutex.Unlock()

		if conn != publisher {
			return &base.Response{
				StatusCode: base.StatusBadRequest,
			}, fmt.Errorf("someone is already publishing")
		}

		return &base.Response{
			StatusCode: base.StatusOK,
			Header: base.Header{
				"Session": base.HeaderValue{"12345678"},
			},
		}, nil
	}

	// called after receiving a frame.
	onFrame := func(trackID int, typ gortsplib.StreamType, buf []byte) {
		mutex.Lock()
		defer mutex.Unlock()

		// if we are the publisher, route frames to readers
		if conn == publisher {
			for r := range readers {
				r.WriteFrame(trackID, typ, buf)
			}
		}
	}

	err := <-conn.Read(gortsplib.ServerConnReadHandlers{
		OnDescribe: onDescribe,
		OnAnnounce: onAnnounce,
		OnSetup:    onSetup,
		OnPlay:     onPlay,
		OnRecord:   onRecord,
		OnFrame:    onFrame,
	})
	log.Printf("client disconnected (%s)", err)

	mutex.Lock()
	defer mutex.Unlock()

	if conn == publisher {
		publisher = nil
		sdp = nil
	} else {
		delete(readers, conn)
	}
}

func main() {
	// create configuration
	conf := gortsplib.ServerConf{
		UDPRTPAddress:  ":8000",
		UDPRTCPAddress: ":8001",
	}

	// create server
	s, err := conf.Serve(":8554")
	if err != nil {
		panic(err)
	}
	log.Printf("server is ready")

	// accept connections
	for {
		conn, err := s.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(conn)
	}
}
