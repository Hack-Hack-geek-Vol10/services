package usecase

import (
	"context"
	"log"

	member "github.com/schema-creator/services/member-service/api/v1"
	"github.com/schema-creator/services/member-service/internal/domain"
	"github.com/schema-creator/services/member-service/internal/infra"
)

type memberService struct {
	member.UnimplementedMemberServer
	memberRepo infra.MemberRepo
}

func NewMemberService(memberRepo infra.MemberRepo) member.MemberServer {
	return &memberService{
		memberRepo: memberRepo,
	}
}

func (m *memberService) CreateMember(ctx context.Context, in *member.MemberRequest) (*member.MemberResponse, error) {
	var (
		result *domain.ProjectMember
		err    error
	)
	members, err := m.memberRepo.ReadAll(ctx, in.ProjectId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if isExitsMember(members, in.UserId) {
		result, err = m.memberRepo.UpdateAuthority(ctx, domain.UpdateAuthorityParam{
			ProjectID: in.ProjectId,
			UserID:    in.UserId,
			Authority: domain.Authority(in.Authority),
		})
	} else {
		result, err = m.memberRepo.Create(ctx, domain.CreateMemberParam{
			ProjectID: in.ProjectId,
			UserID:    in.UserId,
			Authority: domain.Authority(in.Authority),
		})
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &member.MemberResponse{
		ProjectId: result.ProjectID,
		UserId:    result.UserID,
		Authority: member.Auth(member.Auth_value[string(result.Authority)]),
	}, nil
}

func isExitsMember(members []*domain.ProjectMember, userId string) bool {
	for _, mem := range members {
		if mem.UserID == userId {
			return true
		}
	}
	return false
}

func (m *memberService) GetMembers(ctx context.Context, in *member.GetMembersRequest) (*member.ListMembers, error) {
	members, err := m.memberRepo.ReadAll(ctx, in.ProjectId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var listMembers member.ListMembers
	for _, m := range members {
		listMembers.Members = append(listMembers.Members, &member.MemberResponse{
			ProjectId: m.ProjectID,
			UserId:    m.UserID,
			Authority: member.Auth(member.Auth_value[string(m.Authority)]),
		})
	}

	return &listMembers, nil
}
func (m *memberService) UpdateAuthority(ctx context.Context, in *member.MemberRequest) (*member.MemberResponse, error) {
	result, err := m.memberRepo.UpdateAuthority(ctx, domain.UpdateAuthorityParam{
		ProjectID: in.ProjectId,
		UserID:    in.UserId,
		Authority: domain.Authority(in.Authority),
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &member.MemberResponse{
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
		log.Println(err)
		return nil, err
	}

	return &member.DeleteMemberResponse{
		Message: "successful",
	}, nil
}

func (m *memberService) DeleteAllMembers(ctx context.Context, in *member.DeleteAllMemberRequest) (*member.DeleteMemberResponse, error) {
	if err := m.memberRepo.DeleteAll(ctx, in.ProjectId); err != nil {
		log.Println(err)
		return nil, err
	}

	return &member.DeleteMemberResponse{
		Message: "successful",
	}, nil
}

func (m *memberService) mustEmbedUnimplementedMemberServer() {
	panic("implement me")
}
