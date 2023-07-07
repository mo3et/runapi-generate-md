package constant

var FileName = map[string]string{
	"会话相关": "conversationManagement",
	"好友管理": "friendsManagement",
	"群组管理": "groupManagement",
	"用户管理": "userManagement",

	"消息相关":               "messageManagement",
	"三方服务":               "thirdManagement",
	"撤回消息":               "revokeMessage",
	"删除用户所有消息":       "deleteUserAllMessage",
	"发送消息":               "sendMessage",
	"用户注册":               "userRegister",
	"分页获取已注册用户列表": "getUserList",
	"设置用户全局免打扰":     "updateGlobalRemind",
	"获取指定用户列表":       "getSpecifiedUser",
	"更新用户信息":           "updateUserInfo",
	"获取用户token":          "getUserToken",
	"查询用户是否注册":       "checkUserRegistered",
	"分页获取所有用户ID":     "getAllUserID",
	"获取用户在线状态":       "getUserOnlineStatus",
	"统计用户注册":           "statisticsUserRegister",

	"判断是否是好友":               "checkFriend",
	"获取收到的好友申请列表":       "getRecvApplication",
	"移除黑名单列表":               "deleteBlackList",
	"获取好友信息列表":             "getFriendList",
	"设置好友备注":                 "updateFriendRemark",
	"用户申请加好友":               "sendApplication",
	"删除好友":                     "deleteFriend",
	"获取用户黑名单列表":           "getBlackList",
	"添加黑名单":                   "addBlackList",
	"导入好友":                     "importFriend",
	"好友申请处理":                 "processApplication",
	"获取主动发出去的好友申请列表": "getSentApplication",

	"多个用户对同一会话设置字段": "setConversations	",

	"创建群":                               "createGroup",
	"邀请用户进群":                         "inviteUserToGroup",
	"用户加入群":                           "joinGroup",
	"用户退出群":                           "quitGroup",
	"转让群组":                             "transferGroup",
	"获取群组信息":                         "getGroupsInfo",
	"将用户踢出群":                         "kickGroup",
	"获取群成员":                           "getGroupMembersInfo",
	"获取群成员列表":                       "getGroupMemberList",
	"获取用户加群列表":                     "getJoinedGroupList",
	"设置群成员信息":                       "setGroupMemberInfo",
	"禁言群组":                             "muteGroup",
	"取消群禁言":                           "cancelMuteGroup",
	"解散群":                               "dismissGroup",
	"禁言群成员":                           "muteGroupMember",
	"取消禁言群成员":                       "cancelMuteGroupMember",
	"设置群组信息":                         "setGroupInfo",
	"（以管理员或群主身份）获取群的加群申请": "getRecvGroupApplicationList",
	"群主或管理员处理进群申请":             "groupApplicationResponse",
	"获取用户自己的主动加群申请":           "getUserReqGroupApplicationList",
}
var Exclude = map[string]struct{}{
	"获取最新seq":                                     {},
	"通过seq拉取消息":                                 {},
	"用户清空指定会话ID消息":                          {},
	"用户删除会话指定消息":                            {},
	"按时间物理删除消息":                              {},
	"按seq物理删除消息":                               {},
	"获取会话最大seq和用户已读最大seq":                {},
	"标记seq列表中消息已读并且发送已读回执同步未读数": {},
	"SDK-Core编译文档":                                {},
	"OpenIM文档撰写说明以及参数":                      {},
	"获取用户所有会话":                                {},
	"批量设置会话字段":                                {},
	"用户根据会话ID获取指定会话":                      {},
	"用户根据会话ID获取单个会话":                      {},
	"设置会话免打扰":                                  {},
	"调整会话字段":                                    {},
	"获取群成员哈希":                                  {},
	"OpenIM字段说明":                                  {},
}
