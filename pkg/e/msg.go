// Copyright © 2020. Drew Lee. All rights reserved.

package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	// 1xxxx - member error
	ERROR_EXIST_MEMBER:       "已存在该用户",
	ERROR_EXIST_MEMBER_FAIL:  "获取已存在用户失败",
	ERROR_NOT_EXIST_MEMBER:   "该用户不存在",
	ERROR_GET_MEMBERS_FAIL:   "获取所有用户失败",
	ERROR_COUNT_MEMBER_FAIL:  "统计用户失败",
	ERROR_ADD_MEMBER_FAIL:    "新增用户失败",
	ERROR_EDIT_MEMBER_FAIL:   "修改用户失败",
	ERROR_DELETE_MEMBER_FAIL: "删除用户失败",
	ERROR_EXPORT_MEMBER_FAIL: "导出用户失败",
	ERROR_IMPORT_MEMBER_FAIL: "导入用户失败",
	ERROR_LOGIN_MEMBER:       "用户登录失败",
	ERROR_REGISTER_MEMBER:    "用户注册失败",
	ERROR_ADD_FRIEND:         "添加好友失败",

	// 2xxxx - group error
	ERROR_NOT_EXIST_GROUP:        "该群组不存在",
	ERROR_ADD_GROUP_FAIL:         "新增群组失败",
	ERROR_DELETE_GROUP_FAIL:      "删除群组失败",
	ERROR_CHECK_EXIST_GROUP_FAIL: "检查群组是否存在失败",
	ERROR_EDIT_GROUP_FAIL:        "修改群组失败",
	ERROR_COUNT_GROUP_FAIL:       "统计群组失败",
	ERROR_GET_GROUPS_FAIL:        "获取多个群组失败",
	ERROR_GET_GROUP_FAIL:         "获取单个群组失败",
	ERROR_GEN_GROUP_POSTER_FAIL:  "生成群组海报失败",
	ERROR_JOIN_GROUP_FAIL:        "加入群组失败",

	// 3xxxx - auth error
	ERROR_AUTH_NO_TOKEN:            "缺少token",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "token验证失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "token验证超时",
	ERROR_AUTH_GET_TOKEN_FAIL:      "token获取失败",
	ERROR_AHTH_INVALID_HEADERS:     "无效headers参数",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
