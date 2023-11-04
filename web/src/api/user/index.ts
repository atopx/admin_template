import { request } from "@/utils/service"
import type * as Types from "./types"
import { NullResponse } from "types/api"

export function loginApi(data: Types.ILoginRequestData) {
    return request<Types.LoginResponseData>({
        url: "/user/login",
        method: "post",
        data
    })
}

export function refreshTokenApi(data: Types.IRefreshRequestData) {
    return request<Types.RefreshResponseData>({
        url: "/user/refresh",
        method: "post",
        data
    })
}

export function getUserInfoApi() {
    return request<Types.UserInfoResponseData>({
        url: "/user/info",
        method: "get"
    })
}

export function getUserListApi(data: Types.UserListRequestData) {
    return request<Types.UserListResponseData>({
        url: "/user/list",
        method: "post",
        data
    })
}

export function deleteUserApi(userId: number) {
    return request<NullResponse>({
        url: "/user/delete",
        method: "delete",
        data: { userId: userId }
    })
}


export function disableUserApi(userId: number) {
    return request<NullResponse>({
        url: "/user/disable",
        method: "patch",
        data: { userId: userId }
    })
}

export function createUserApi(data: Types.UserInfo) {
    return request<NullResponse>({
        url: "/user/create",
        method: "post",
        data
    })
}

export function updateUserApi(data: Types.UserInfo) {
    return request<NullResponse>({
        url: "/user/update",
        method: "post",
        data
    })
}
