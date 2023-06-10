package payload

import (
	tpb "google.golang.org/protobuf/types/known/timestamppb"
	bin "github.com/0xzero/matchmaker/binary"
)

type SocketPayloads struct {}

func (p *SocketPayloads) StartTypingPayload(userId string, matchId string) ([]byte, error) {
	msg := &bin.SocketEvents{
		Data: &bin.SocketEvents_StartTyping{
			StartTyping: &bin.StartTyping{
				Timestamp: tpb.Now(),
				MatchId: matchId,
				UserId: userId,
			},
		},
	}
	payload, err := bin.EncodeProtoMessage(msg)
	return payload,err
}