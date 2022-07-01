import { request, IApiResponse } from "./base"


// 登录
export const login = (data: any): Promise<IApiResponse> => {
    return request.post('auth/login', data)
};

// 登录(token)
export const loginToken = (): Promise<IApiResponse> => {
    return request.post('auth/login_token')
};

// 登出
export const logout = () => {
    return request.post('auth/logout')
}
