package services

import (
	"encoding/json"
	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/gregor_chair_common/configuration"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/logger"
	"github.com/thankala/gregor_chair_common/messages"
	"net"
)

// TCPServer is a service that listens for incoming TCP connections
type TCPServer struct {
	listenAddr string
	stopCh     chan struct{}
}

func NewTCPServer(opts ...configuration.TcpOptFunc) *TCPServer {
	options := configuration.DefaultTcpOpts()
	for _, opt := range opts {
		opt(options)
	}
	return &TCPServer{
		listenAddr: options.Address,
		stopCh:     make(chan struct{}),
	}
}

func (s *TCPServer) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Initialized:
		// Do nothing
	case actor.Started:
		go s.Accept(ctx, s.stopCh)
	case actor.Stopped:
		close(s.stopCh)
	case *messages.AssemblyTaskMessage:
		s.Send(msg.Source, msg.Destination, msg.Event, msg)
	case *messages.CoordinatorMessage:
		s.Send(msg.Source, msg.Destination, msg.Event, msg)
	default:
		logger.Get().Error("Unknown message", "Message", msg)
	}
}

func (s *TCPServer) Accept(ctx *actor.Context, stopCh <-chan struct{}) {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		panic(err)
	}
	for {
		select {
		case <-stopCh:
			return // Exit the loop if something is received on stopCh
		default:
			conn, err := ln.Accept()
			if err != nil {
				logger.Get().Error("Unable to accept connection", "Error", err)
				break
			}
			go func(ctx *actor.Context, conn net.Conn) {
				buf := make([]byte, 1024)
				for {
					n, err := conn.Read(buf)
					if err != nil {
						//slog.Error("conn read error", "err", err)
						break
					}
					// copy shared buffer, to prevent race conditions.
					msg := make([]byte, n)
					copy(msg, buf[:n])

					var baseEvent messages.BaseEvent
					err = json.Unmarshal(msg, &baseEvent)
					if err != nil {
						logger.Get().Error("Unable to serialize base event", "Event", msg, "Error", err)
						continue
					}
					switch baseEvent.Event {
					case enums.CoordinatorEvent:
						var coordinatorMessage messages.CoordinatorMessage
						err = json.Unmarshal(baseEvent.Data, &coordinatorMessage)
						if err == nil {
							// send it to the parent actor
							ctx.Send(ctx.Parent(), &coordinatorMessage)
							continue
						}
					case enums.AssemblyTaskEvent:
						var assemblyTaskMessage messages.AssemblyTaskMessage
						err = json.Unmarshal(baseEvent.Data, &assemblyTaskMessage)
						if err == nil {
							// send it to the parent actor
							ctx.Send(ctx.Parent(), &assemblyTaskMessage)
							continue
						}
					default:
						logger.Get().Warn("Unable to serialize event", "Event", msg)
					}
					err = conn.Close()
					if err != nil {
						return
					}
				}
			}(ctx, conn)
		}
	}
}

func (s *TCPServer) Send(from string, to string, event enums.Event, msg any) {
	jsonMessage, err := json.Marshal(msg)
	if err != nil {
		logger.Get().Error("Unable to marshal event", "Event", msg)
		return
	}

	// Create a BaseEvent struct with the marshaled message as RawMessage
	baseEvent := &messages.BaseEvent{
		Event: event,
		Data:  json.RawMessage(jsonMessage), // Assign marshaled message to Data
	}

	// Serialize the message to JSON
	serializedMsg, err := json.Marshal(baseEvent)
	if err != nil {
		logger.Get().Error("Failed to serialize event", "Error", err)
		return
	}

	// Get the connection for the recipient
	var recipientConn net.Conn
	recipientConn, err = net.Dial("tcp", to)
	if err != nil {
		logger.Get().Error("Failed to connect to recipient", "Error", err)
		return
	}
	// Send the serialized message over the connection
	_, err = recipientConn.Write(serializedMsg)
	if err != nil {
		logger.Get().Error("Failed to send event", "Error", err)
	}
}

func (s *TCPServer) GetProducer() actor.Producer {
	return func() actor.Receiver {
		return s
	}
}
