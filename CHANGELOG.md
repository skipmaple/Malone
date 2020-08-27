# CHANGE LOG

## 20200818
- jwt token 的获取需要在headers里传入member_id
- token 有效期三小时
- 登录和注册不需要token（v0接口都不需要）
- 请求v1接口内容需要携带token

## 20200825
- 添加登出接口

## 20200827
- token获取传入的memberId需要base64加密
- 返回token成功的状态码由200改为201
- 成功登出后清除token
