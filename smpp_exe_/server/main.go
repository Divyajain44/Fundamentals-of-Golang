package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ajankovic/smpp"
	"github.com/ajankovic/smpp/pdu"
)

var (
	serverAddr string
	dstAddr    string
	srcAddr    string
	systemID   string
	msg        string
	msgID      int
)

func main() {
	flag.StringVar(&serverAddr, "addr", "localhost:2775", "server will listen on this address.")
	flag.StringVar(&dstAddr, "dst_addr", "111111", "destination to which you are sending the message.")
	flag.StringVar(&srcAddr, "src_addr", "222222", "source from which the message is comming from.")
	flag.StringVar(&msg, "msg", "example", "contents of the message.")
	flag.StringVar(&systemID, "systemid", "ExampleServer", "descriptive server identification.")
	flag.Parse()

	sessConf := smpp.SessionConf{
		Handler: smpp.HandlerFunc(func(ctx *smpp.Context) {
			switch ctx.CommandID() {
			case pdu.BindTransceiverID:
				btrx, err := ctx.BindTRx()
				if err != nil {
					fail("Invalid PDU in context error: %+v", err)
				}
				resp := btrx.Response(systemID)
				if err := ctx.Respond(resp, pdu.StatusOK); err != nil {
					fail("Server can't respond to the Binding request: %+v", err)
				}
			case pdu.BindTransmitterID:
				btx, err := ctx.BindTx()
				if err != nil {
					fail("Invalid PDU in context error: %+v", err)
				}
				resp := btx.Response(systemID)
				if err := ctx.Respond(resp, pdu.StatusOK); err != nil {
					fail("Server can't respond to the Binding request: %+v", err)
				}
			case pdu.BindReceiverID:
				ctx1 := context.Background()
				brx, err := ctx.BindRx()
				if err != nil {
					fail("Invalid PDU in context error: %+v", err)
				}
				resp := brx.Response(systemID)
				if err := ctx.Respond(resp, pdu.StatusOK); err != nil {
					fail("Server can't respond to the Binding request: %+v", err)
				}
				sm := &pdu.DeliverSm{
					SourceAddr:      srcAddr,
					DestinationAddr: dstAddr,
					ShortMessage:    msg,
				}
				go DeliverSmSend(ctx1, sm, ctx.Sess)
				time.Sleep(1 * time.Minute)
			case pdu.SubmitSmID:
				sm, err := ctx.SubmitSm()
				if err != nil {
					fail("Invalid PDU in context error: %+v", err)
				}
				fmt.Fprintf(os.Stdout, "UPPER: %s\n", strings.ToUpper(sm.ShortMessage))
				msgID++
				resp := sm.Response(fmt.Sprintf("msgID_%d", msgID))
				if err := ctx.Respond(resp, pdu.StatusOK); err != nil {
					fail("Server can't respond to the submit_sm request: %+v", err)
				}
			case pdu.UnbindID:
				unb, err := ctx.Unbind()
				if err != nil {
					fail("Invalid PDU in context error: %+v", err)
				}
				resp := unb.Response()
				if err := ctx.Respond(resp, pdu.StatusOK); err != nil {
					fail("Server can't respond to the submit_sm request: %+v", err)
				}
				ctx.CloseSession()
			}
		}),
	}
	srv := smpp.NewServer(serverAddr, sessConf)

	fmt.Fprintf(os.Stderr, "'%s' is listening on '%s'\n", systemID, serverAddr)
	err := srv.ListenAndServe()
	if err != nil {
		fail("Serving exited with error: %+v", err)
	}
	fmt.Fprintf(os.Stderr, "Server closed\n")
}

func fail(msg string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", params...)
	os.Exit(1)
}

func DeliverSmSend(ctx context.Context, req pdu.PDU, sess *smpp.Session) {
	timeout := time.Duration(30 * time.Second)
	ctx1, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, err := sess.Send(ctx1, req)
	if err != nil {
		fail("Can't send message: %+v", err)
	}
	fmt.Fprintf(os.Stderr, "Deliver_sm_sent\n")
}
