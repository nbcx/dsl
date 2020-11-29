package client

import (
	"github.com/nbcx/gcs/model"
)

func (s *Remote) BroadcastFd(server *model.Server, fd string, msg []byte) (userIds []string, err error) {
	return
}

func (s *Remote) BroadcastUid(appId, uid string, msg []byte) (userIds []string, err error) {

	return
}

func (s *Remote) Login(server *model.Server, appId uint32) (userIds []string, err error) {
	return
}

func (s *Remote) JoinGroup(server *model.Server, groupId, appId string) (userIds []string, err error) {
	return
}

func (s *Remote) JoinGroupWithUid(appId, uid, groupId string) (userIds []string, err error) {
	return
}
