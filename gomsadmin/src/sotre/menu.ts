import { defineStore } from "pinia"

import * as utilMenu from '@/util/menu'
import { IMenu } from "@/data/menu"
import { RouteRecordRaw } from "vue-router"


const useMenuStore = defineStore('menu', {
    state: () => {
        let menus: IMenu[] = [] //所有菜单
        let routes: RouteRecordRaw[] = [] //所有路由

        let headerMenus: { [key: string]: IMenu; } = {} //顶部 key 对应的菜单数据
        let headerRoutes: { [key: string]: string; } = {} //顶部 key 对应的路由数据

        let menusAside: IMenu[] = [] //左侧菜单数据
        let activeHeader = '' //顶部 key
        let breadcrumbs: { [key: string]: string[]; } = {} //菜单面包屑数据

        return {
            menus,
            routes,
            headerMenus,
            headerRoutes,
            menusAside,
            activeHeader,
            breadcrumbs,
        }
    },

    getters: {
        isInit(): boolean {
            return this.menus.length > 0
        },

        dynamicRoutes(): RouteRecordRaw[] {
            return this.routes
        }
    },

    actions: {
        init(userType: number) {
            const [genMenus, genRoutes] = utilMenu.gen(userType)

            this.menus = genMenus
            this.routes = genRoutes

            this.activeHeader = genMenus[0].name

            for (let i = 0; i < genMenus.length; i++) {
                this.headerMenus[genMenus[i].name] = genMenus[i]
            }
            this.headerRoutes = utilMenu.getHeaderRoutes(genMenus)

            this.breadcrumbs = utilMenu.getBreadcrumbs(genMenus)
        },

        destory() {
            this.$reset()
        },

        onHeaderTap(key: string) {
            const menu =  this.headerMenus[key]

            this.activeHeader = key
            this.menusAside = menu.children!
        },

        onAsideTap(key: string) {
            this.headerRoutes[this.activeHeader] = key
        }
    },
})


export default useMenuStore
