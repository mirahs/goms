import { RouteRecordRaw, createWebHashHistory, createRouter, NavigationGuardNext } from 'vue-router'

import useAuthStore from '@/sotre/auth'
import useMenuStore from '@/sotre/menu'

import * as auth from "@/api/auth"


const routes: RouteRecordRaw[] = [
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/Login.vue'),
    },
    {
        path: '/',
        name: 'Main',
        component: () => import('@/views/Main.vue'),
        meta: { auth: true },
    }
]

const router = createRouter({
    routes,
    history: createWebHashHistory()
})


const addDynamicRoute = (dynamicRoutes: RouteRecordRaw[]) => {
    const mainRoute = router.options.routes.find((v) => v.name === "Main");
    if (mainRoute == undefined) return

    if (mainRoute.children != undefined && mainRoute.children.length > 0) return
  
    mainRoute.children = dynamicRoutes;

    router.addRoute(mainRoute);
  }

// 使用 token 登录
const tokenLogin = async (authStore: any, next: NavigationGuardNext) => {
    const res = await auth.loginToken()

    if (res && res.code == 0) {
        console.error('router tokenLogin 登录成功, 进入主页')

        authStore.tokenLogin(res.data)

        next({ name: 'Main' })
    } else {
        console.error('router tokenLogin 清除所有数据 返回错误信息:' + res.msg)

        authStore.tokenInvalid()

        next({ name: 'Login' })
    }
}


router.beforeEach((to, from, next) => {
    const authStore = useAuthStore()
    // console.error('router.beforeEach to:', to.name, ',from:', from.name, ',isLogin:', authStore.isLogin, ',token:', authStore.token)

    if (to.name == null) {
        // 路由不存在, 说明是动态加载的路由, 可能是浏览器刷新后路由数据丢失
        next({ name: 'Main' })
        return
    }

    if (authStore.isLogin) {
        const menuStore = useMenuStore()
        if (!menuStore.isInit) {
            menuStore.init(authStore.userInfo.type)

            addDynamicRoute(menuStore.dynamicRoutes)

            next({ ...to, replace: true })
        } else {
            if (to.name == 'Login' && from != undefined && from.name != 'Login') {
                // 已经登录不跳转到登录路由直接返回原来的路由
                next(from)
            } else {
                next();
            }
        }
    } else {
        if (to.meta.auth) {
            if (authStore.token != '') {
                tokenLogin(authStore, next)
            } else {
                next({ name: 'Login' })
            }
        } else {
            if (to.name == 'Login' && authStore.token != '') {
                // 如果是跳转登录路由则先用 token 登录
                tokenLogin(authStore, next)
            } else {
                next()
            }
        }
    }
})


export default router
