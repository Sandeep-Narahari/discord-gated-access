package bot

import (
	"context"
	"fmt"

	botpb "github.com/AutonomyNetwork/iam/types/bot/v1/bot"
)

type Server struct {
	botpb.UnimplementedCommunityServicesServer
	botpb.UnimplementedUserServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) CreateCommunity(ctx context.Context, r *botpb.CreateCommunityRequest) (*botpb.CreateCommunityResponse, error) {

	fmt.Println("CREATE COMMUNITY")
	fmt.Println(&botpb.CreateCommunityResponse{GatedCollectionId: r.GatedCollectionId})
	fmt.Println(r.AccessRoleName, r.AccountAddress)
	return &botpb.CreateCommunityResponse{GatedCollectionId: "collectionid"}, nil
}
func (s *Server) UpdateCommunity(ctx context.Context, r *botpb.UpdateCommunityRequest) (*botpb.UpdateCommunityResponse, error) {
	return &botpb.UpdateCommunityResponse{Id: 0}, nil
}
func (s *Server) Communities(ctx context.Context, r *botpb.GetCommunitiesRequest) (*botpb.GetCommunitiesResponse, error) {
	return &botpb.GetCommunitiesResponse{}, nil
}
func (s *Server) CommunityById(ctx context.Context, r *botpb.GetCommunityByIDRequest) (*botpb.GetCommunityResponse, error) {
	return &botpb.GetCommunityResponse{}, nil
}
func (s *Server) CommunityByCollectionId(ctx context.Context, r *botpb.GetCommunityByCollectionIdRequest) (*botpb.GetCommunityResponse, error) {
	return &botpb.GetCommunityResponse{}, nil
}
func (s *Server) CommunityByNFTRoleId(ctx context.Context, r *botpb.GetCommunityByNFTRoleIdRequest) (*botpb.GetCommunityResponse, error) {
	return &botpb.GetCommunityResponse{}, nil
}
func (s *Server) CommunityByDiscordCategoryID(ctx context.Context, r *botpb.GetCommunityByDiscordCategoryIDRequest) (*botpb.GetCommunityResponse, error) {
	return &botpb.GetCommunityResponse{}, nil
}

func (s *Server) CreateUser(ctx context.Context, r *botpb.CreateUserRequest) (*botpb.CreateUserResponse, error) {
	return &botpb.CreateUserResponse{}, nil
}
