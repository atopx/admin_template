import { type } from "os"
import { PageIface, ResponseIface, TimeRangeIface } from "types/api"

export interface ILoginRequestData {
    username: string
    password: string
}

export interface IRefreshRequestData {
    refreshToken: string
}

export type UserListRequestData = {
    pageInfo: PageIface;
    filter: {
        keyword: string;
        timeRange: TimeRangeIface;
    }
}

export interface Token {
    userId: number
    expireTime: number
    accessToken: string
    refreshToken: string
}

export type UserInfo = {
    userId: number;
    username?: string;
    password?: string;
    status?: string;
    level?: number;
    avatar?: string;
    email?: string;
    phone?: string;
    create_time?: string;
    update_time?: string;
}


export type LoginResponseData = ResponseIface<Token>

export type RefreshResponseData = ResponseIface<Token>

export type UserInfoResponseData = ResponseIface<UserInfo>

export type UserListResponseData = ResponseIface<{
    pager: any
    title: string
    content: string
    pageInfo: PageIface;
    list: UserInfo[];
}>