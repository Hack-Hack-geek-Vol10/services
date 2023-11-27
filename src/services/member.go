package services

import (
	"context"

	member "github.com/Hack-Hack-geek-Vol10/services/pkg/grpc/member-service/v1"
	"github.com/Hack-Hack-geek-Vol10/services/src/domain"
	"github.com/Hack-Hack-geek-Vol10/services/src/storages"
)

type memberService struct {
	member.UnimplementedMemberServiceServer
	memberRepo storages.MemberRepo
}

func NewMemberService(memberRepo storages.MemberRepo) member.MemberServiceServer {
	return &memberService{
		memberRepo: memberRepo,
	}
}

func (m *memberService) AddMember(ctx context.Context, in *member.MemberRequest) (*member.Member, error) {
	result, err := m.memberRepo.Create(ctx, domain.CreateMemberParam{
		ProjectID: in.ProjectId,
		UserID:    in.UserId,
		Authority: domain.Authority(in.Authority),
	})

	if err != nil {
		return nil, err
	}
	return &member.Member{
		ProjectId: result.ProjectID,
		UserId:    result.UserID,
		Authority: member.Auth(member.Auth_value[string(result.Authority)]),
	}, nil
}

func (m *memberService) ReadMembers(ctx context.Context, in *member.GetMembersRequest) (*member.ListMembers, error) {
	members, err := m.memberRepo.ReadAll(ctx, in.ProjectId)
	if err != nil {
		return nil, err
	}

	var listMembers member.ListMembers
	for _, m := range members {
		listMembers.Members = append(listMembers.Members, &member.Member{
			ProjectId: m.ProjectID,
			UserId:    m.UserID,
			Authority: member.Auth(member.Auth_value[string(m.Authority)]),
		})
	}

	return &listMembers, nil
}
func (m *memberService) UpdateAuthority(ctx context.Context, in *member.MemberRequest) (*member.Member, error) {
	result, err := m.memberRepo.UpdateAuthority(ctx, domain.UpdateAuthorityParam{
		ProjectID: in.ProjectId,
		UserID:    in.UserId,
		Authority: domain.Authority(in.Authority),
	})
	if err != nil {
		return nil, err
	}

	return &member.Member{
		ProjectId: result.ProjectID,
		UserId:    result.UserID,
		Authority: member.Auth(member.Auth_value[string(result.Authority)]),
	}, nil
}

func (m *memberService) DeleteMember(ctx context.Context, in *member.DeleteMemberRequest) (*member.DeleteMemberResponse, error) {
	if err := m.memberRepo.Delete(ctx, domain.DeleteMemberParam{
		ProjectID: in.ProjectId,
		UserID:    in.UserId,
	}); err != nil {
		return nil, err
	}

	return &member.DeleteMemberResponse{
		Message: "successful",
	}, nil
}
