package services

import (
	"encoding/json"
	"net"

	"github.com/anthdm/hollywood/actor"
	"github.com/thankala/gregor_chair_common/configuration"
	"github.com/thankala/gregor_chair_common/enums"
	"github.com/thankala/gregor_chair_common/events"
	"github.com/thankala/gregor_chair_common/logger"
)

// TCPServer is a service that listens for incoming TCP connections
type TCPServer struct {
	listenAddr string
}

func NewTCPServer(opts ...configuration.TcpOptFunc) *TCPServer {
	options := configuration.DefaultTcpOpts()
	for _, opt := range opts {
		opt(options)
	}
	return &TCPServer{
		listenAddr: options.Address,
	}
}

func (s *TCPServer) Receive(ctx *actor.Context) {
	switch event := ctx.Message().(type) {
	case actor.Initialized:
		// Do nothing
	case actor.Started:
		s.Accept(ctx)
	case *events.AssemblyTaskEvent:
		s.Send(event.Source.String(), event.Destination.String(), enums.AssemblyTaskEvent, event)
	case *events.OrchestratorEvent:
		s.Send(event.Source.String(), event.Destination.String(), enums.OrchestratorEvent, event)
	default:
		logger.Get().Error("Unknown message", "Event", event)
	}
}

func (s *TCPServer) Accept(ctx *actor.Context) {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		panic(err)
	}
	for {
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

				var baseEvent events.BaseEvent
				err = json.Unmarshal(msg, &baseEvent)
				if err != nil {
					logger.Get().Error("Unable to serialize base event", "Event", msg, "Error", err)
					continue
				}
				switch baseEvent.Event {
				case enums.OrchestratorEvent:
					var orchestratorEvent events.OrchestratorEvent
					err = json.Unmarshal(baseEvent.Data, &orchestratorEvent)
					if err == nil {
						// send it to the parent actor
						ctx.Send(ctx.Parent(), &orchestratorEvent)
						continue
					}
				case enums.AssemblyTaskEvent:
					var assemblyTaskEvent events.AssemblyTaskEvent
					err = json.Unmarshal(baseEvent.Data, &assemblyTaskEvent)
					if err == nil {
						// send it to the parent actor
						ctx.Send(ctx.Parent(), &assemblyTaskEvent)
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

func (s *TCPServer) Send(from string, to string, event enums.Event, msg any) {
	jsonMessage, err := json.Marshal(msg)
	if err != nil {
		logger.Get().Error("Unable to marshal event", "Event", msg)
		return
	}

	// Create a BaseEvent struct with the marshaled message as RawMessage
	baseEvent := &events.BaseEvent{
		Event: event,
		Data:  json.RawMessage(jsonMessage), // Assign marshaled message to Data
	}

	// Serialize the message to JSON
	serializedEvent, err := json.Marshal(baseEvent)
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
	_, err = recipientConn.Write(serializedEvent)
	if err != nil {
		logger.Get().Error("Failed to send event", "Error", err)
	}
}

func (s *TCPServer) GetProducer() actor.Producer {
	return func() actor.Receiver {
		return s
	}
}
