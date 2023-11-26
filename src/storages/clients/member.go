package clients

import (
	"context"

	member "github.com/Hack-Hack-geek-Vol10/services/pkg/grpc/member-service/v1"
	"github.com/Hack-Hack-geek-Vol10/services/src/domain"
	"google.golang.org/grpc"
)

type memberClient struct {
	conn *grpc.ClientConn
}
type MemberClient interface {
	AddMember(ctx context.Context, arg domain.CreateMemberParam) (*member.Member, error)
}

func NewMemberClient(conn *grpc.ClientConn) MemberClient {
	return &memberClient{
		conn: conn,
	}
}

func (m *memberClient) AddMember(ctx context.Context, arg domain.CreateMemberParam) (*member.Member, error) {
	return member.NewMemberServiceClient(m.conn).AddMember(ctx, &member.MemberRequest{
		UserId:    arg.UserID,
		ProjectId: arg.ProjectID,
		Authority: string(arg.Authority),
	})
}
