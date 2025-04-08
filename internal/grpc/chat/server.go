package chat

import (
	desc "chat-server/pkg/chat_v1"
)

type Server struct {
	desc.UnimplementedChatServer
}

// func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
// 	panic("implement me")
// }

// func (s *Server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
// 	panic("implement me")
// }

// func (s *Server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
// 	panic("implement me")
// }
