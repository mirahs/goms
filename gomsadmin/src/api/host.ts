import { request, IApiResponse } from "./base"


// 获取所有主机
export const getAll = (page: number, pageSize: number): Promise<IApiResponse> => {
    return request.get('hosts', { params: { _page: page, _page_size: pageSize } })
}

// 添加主机
export const add = (data: any): Promise<IApiResponse> => {
    return request.post('hosts', data)
}

// 编辑主机
export const edit = (id: number, data: any): Promise<IApiResponse> => {
    return request.put(`hosts/${id}`, data)
}

// 删除主机
export const del = (id: number): Promise<IApiResponse> => {
    return request.delete(`hosts/${id}`)
}
