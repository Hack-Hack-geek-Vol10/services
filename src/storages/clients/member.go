package clients

import "github.com/Hack-Hack-geek-Vol10/services/src/domain"

type memberClient struct {
}

type MemberClient interface {
}

func NewMemberClient() MemberClient {
	return &memberClient{}
}

func (m *memberClient) AddMember(arg domain.CreateMemberParam) {

}
