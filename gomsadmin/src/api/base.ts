import axios from 'axios'
import { ElMessage } from 'element-plus'

import useAuthStore from '@/sotre/auth'


export interface IApiResponse {
    code: number,
    msg: string,
    data: any,
}


export const request = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL,
    timeout: 5 * 1000,
})


// 添加请求拦截器
request.interceptors.request.use((config) => {
    const authStore = useAuthStore()

    config.headers!['X-Token'] = authStore.token

    return config
}, (error) => {
    return Promise.reject(error)
})

// 添加响应拦截器
request.interceptors.response.use(response => {
    if (response.data.code != 0) {
        ElMessage.error(response.data.msg)
    }

    return response.data
}, error => {
    if (error.response && error.response.data.msg) {
        if (error.response.data.code == 1) {
            const authStore = useAuthStore()
            authStore.tokenInvalid()
        }

        ElMessage.error(error.response.data.msg)

        return error.response.data
    }

    ElMessage.error('网络错误')

    return { code: -1}
})
