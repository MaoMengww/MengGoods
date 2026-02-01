namespace go user 

include "model.thrift"


//注册
struct RegisterReq {
    1: required string username;
    2: required string password;
    4: required string avatarName;
    5: required string email;
}

struct RegisterResp {
    1: required model.BaseResp base;
    2: required i64 userId;
}

//上传头像
struct UploadAvatarReq {
    1: binary avatarData;
    2: string avatarName;
}

struct UploadAvatarResp {
    1: model.BaseResp base;
    2: string avatarURL;
}

//登录
struct LoginReq {
    1: required string username;
    2: required string password;
}

struct LoginResp {
    1: required model.BaseResp base;
    2: required model.UserInfo userInfo;
}

//增加地址
struct AddAddressReq {
    1: required string province;
    2: required string city;
    3: required string detail;
}

struct AddAddressResp {
    1: required model.BaseResp base;
    2: required i64 addressId;
}

//获取地址(userID由context隐式透传)
struct GetAddressesReq {
}

struct GetAddressesResp {
    1: required model.BaseResp base;
    2: required list<model.AddressInfo> address;
}

struct GetAddressReq {
    1: required i64 addressId;
}

struct GetAddressResp {
    1: required model.BaseResp base;
    2: required model.AddressInfo address;
}


//封禁用户
struct BanUserReq {
    1: required i64 userId;
}

struct BanUserResp {
    1: required model.BaseResp base;
}

//解除封禁用户
struct UnBanUserReq {
    1: required i64 userId;
}

struct UnBanUserResp {
    1: required model.BaseResp base;
}

//设置管理员
struct SetAdminReq {
    1: required i64 userId;        
    2: required string password;
}

 struct SetAdminResp {
    1: required model.BaseResp base;
}

//获取用户信息
struct GetUserInfoReq {
    1: required i64 userId;
}

struct GetUserInfoResp {
    1: required model.BaseResp base;
    2: required model.UserInfo userInfo;
}


//登出
struct LogoutReq {
}

struct LogoutResp {
    1: required model.BaseResp base;
}

//发送验证码
struct SendCodeReq {
    1: required string email;
    2: required i64 action;
}

struct SendCodeResp {
    1: required model.BaseResp base;
}

//重置密码
struct ResetPwdReq {
    1: required string email;
    2: required string code;
    3: required string password;
}

struct ResetPwdResp {
    1: required model.BaseResp base;
}





service UserService {
    RegisterResp Register(1: RegisterReq req);
    LoginResp Login(1: LoginReq req);
    AddAddressResp AddAddress(1: AddAddressReq req);
    GetAddressesResp GetAddresses(1: GetAddressesReq req);
    GetAddressResp GetAddress(1: GetAddressReq req);
    BanUserResp BanUser(1: BanUserReq req);
    UnBanUserResp UnBanUser(1: UnBanUserReq req);
    SetAdminResp SetAdmin(1: SetAdminReq req);
    GetUserInfoResp GetUserInfo(1: GetUserInfoReq req);
    LogoutResp logout(1: LogoutReq req);
    SendCodeResp SendCode(1: SendCodeReq req);
    ResetPwdResp ResetPwd(1: ResetPwdReq req);
    UploadAvatarResp UploadAvatar(1: UploadAvatarReq req);
}