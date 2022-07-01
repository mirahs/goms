import { request, IApiResponse } from "./base"
import { IUserSearch } from "@/intfs"



// 获取管理员登录日志
export const getLoginAdmUser = (page: number, pageSize: number, search: IUserSearch): Promise<IApiResponse> => {
    return request.get('log/login/adm_user', { params: { _page: page, _page_size: pageSize, ...search } })
}

// 删除管理员登录日志
export const delLoginAdmUser = (id: number): Promise<IApiResponse> => {
    return request.delete(`log/login/adm_user/${id}`)
}
