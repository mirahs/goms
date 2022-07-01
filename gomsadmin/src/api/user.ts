import { request, IApiResponse } from "./base"


// 获取所有用户
export const getAll = (page: number, pageSize: number): Promise<IApiResponse> => {
    return request.get('adm_users', { params: { _page: page, _page_size: pageSize } })
}

// 添加用户
export const add = (data: any): Promise<IApiResponse> => {
    return request.post('adm_users', data)
}

// 编辑用户
export const edit = (id: number, data: any): Promise<IApiResponse> => {
    return request.put(`adm_users/${id}`, data)
}

// 删除用户
export const del = (id: number): Promise<IApiResponse> => {
    return request.delete(`adm_users/${id}`)
}

// 锁住用户
export const lock = (id: number): Promise<IApiResponse> => {
    return request.patch(`adm_users/${id}/lock`)
}

// 重置用户密码
export const reset = (id: number): Promise<IApiResponse> => {
    return request.patch(`adm_users/${id}/reset`)
}

// 修改密码
export const password = (data: any): Promise<IApiResponse> => {
    return request.patch('adm_users/password', data)
}
