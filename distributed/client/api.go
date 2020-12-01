package client

import (
	"github.com/nbcx/gcs/util"
)

func (s *Remote) BroadcastFd(server *util.Server, fd string, msg []byte) (userIds []string, err error) {
	return
}

func (s *Remote) BroadcastUid(server *util.Server, appId, uid string, msg []byte) (userIds []string, err error) {

	return
}

func (s *Remote) Login(server *util.Server, appId uint32) (userIds []string, err error) {
	return
}

func (s *Remote) JoinGroup(server *util.Server, groupId, appId string) (userIds []string, err error) {
	return
}

func (s *Remote) JoinGroupWithUid(appId, uid, groupId string) (userIds []string, err error) {
	return
}
