package apiBuilder

import (
	"context"
	"errors"
	"github.com/charmbracelet/log"
)

type IGroupMemberInfo interface {
	ToGroupIDAndUserID(groupID int64, userID int64) IGroupMemberInfo
	DoResponse(ctx context.Context) (*Status, error)
	DoApi
}

func (b *Builder) GetGroupMemberInfo() IGroupMemberInfo {
	return b
}

func (b *Builder) ToGroupIDAndUserID(groupID int64, userID int64) IGroupMemberInfo {
	b.action = GetGroupMemberInfo
	b.Params.GroupID = groupID
	b.Params.UserID = userID
	return b
}

func (b *Builder) DoResponse(ctx context.Context) (*Status, error) {
	resp, err := b.DoAndResponse(ctx)
	if err != nil {
		return nil, err
	}
	log.Debug(string(resp.GetOrigin()))
	if !resp.Ok() {
		return nil, errors.New(resp.StatusMsg())
	}

	if err != nil {
		return nil, err
	}
	return resp.response, nil
}
