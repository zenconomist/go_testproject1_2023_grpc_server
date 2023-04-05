package logstream

import (
	logevents "logstream1_proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type LogEventServer struct {
	logevents.UnimplementedLogEventServiceServer
}

func (s *LogEventServer) StreamLogEvents(req *logevents.StreamLogEventsRequest, stream logevents.LogEventService_StreamLogEventsServer) error {
	// Replace this with your actual implementation to generate or fetch log events.
	// This example sends a few dummy log events.
	logEvents := []*logevents.LogEvent{
		{Id: 1, Msg: "First log event", Timestamp: timestamppb.Now()},
		{Id: 2, Msg: "Second log event", Timestamp: timestamppb.Now()},
		{Id: 3, Msg: "Third log event", Timestamp: timestamppb.Now()},
	}

	for _, logEvent := range logEvents {
		if err := stream.Send(logEvent); err != nil {
			return err
		}
	}

	return nil
}
