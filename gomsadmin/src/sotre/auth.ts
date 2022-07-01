import { defineStore } from "pinia"

import * as auth from "@/api/auth"

import { IUser } from '@/intfs'


const useAuthStore = defineStore('auth', {
    state: () => {
        return {
            isLogin: false,
            token: localStorage.getItem('token') ?? '',
            _userInfo: null,
        }
    },

    persist: {
        enabled: true,
    },

    getters: {
        userInfo(): IUser {
            return this._userInfo!
        }
    },

    actions: {
        async login(data: any) {
            var res = await auth.login(data)
            if (0 != res.code) return false

            this.isLogin = true
            this.token = res.data.token
            this._userInfo = res.data.user

            localStorage.setItem('token', this.token)

            return true
        },

        async logout() {
            this.destory()

            await auth.logout()
        },

        // 使用 token 登录成功回调
        tokenLogin(userInfo: any) {
            this.isLogin = true
            this._userInfo = userInfo
        },

        // 使用 token 登录时失效回调
        tokenInvalid() {
            this.destory()
        },

        destory() {
            this.$reset()

            this.isLogin = false
            this.token = ''
            this._userInfo = null

            localStorage.removeItem('token')
            sessionStorage.clear()
        }
    },
})


export default useAuthStore
