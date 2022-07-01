import { IMenu, menus } from "@/data/menu"
import { RouteRecordRaw } from "vue-router";


const views = import.meta.glob('@/views/*/*.vue')


/**
 * 生成用户菜单和路由数据
 * @param userType 用户类型
 * @returns 菜单和路由数据
 */
export const gen = (userType: number): [IMenu[], RouteRecordRaw[]] => {
    return gen2(userType, menus, '/', [], []);
}

const gen2 = (userType: number, menus: IMenu[], parentPath: string, genMenus: IMenu[], genRoutes: RouteRecordRaw[]): [IMenu[], RouteRecordRaw[]] => {
    for (let i = 0; i < menus.length; i++) {
        const menu = menus[i];
        if (menu.children) {
            const [genMenusChild, genRoutesChild] = gen2(userType, menu.children, parentPath + menu.name + '/', [], []);
            if (genMenusChild.length > 0) {
                genMenus.push({ name: menu.name, desc: menu.desc, children: genMenusChild });
                genRoutes = genRoutes.concat(genRoutesChild);
            }
        } else {
            if (menu.key!.indexOf(userType) >= 0) {
                menu.route = parentPath + menu.name
                genMenus.push(menu);

                const genRoute = {
                    path: parentPath + menu.name,
                    name: menu.name,
                    meta: { auth: true },
                    component: views['../views' + parentPath + menu.name + '.vue'],

                }
                genRoutes.push(genRoute);
            }
        }
    }
    return [genMenus, genRoutes];
}


/**
 * 获取顶部菜单对应的初始路由
 * @param menus 菜单数据
 * @returns 顶部菜单对应的初始路由
 */
export const getHeaderRoutes = (menus: IMenu[]) => {
    const headerRoutes: { [key: string]: string; } = {}

    for (let i = 0; i < menus.length; i++) {
        const menu = menus[i]

        const header = menu.name
        const route = getHeaderRouteFirst(menu.children![0], '/' + header)

        headerRoutes[header] = route
    }

    return headerRoutes
}

const getHeaderRouteFirst = (menu: IMenu, path: string): string => {
    path += "/" + menu.name;

    if (menu.children) {
        return getHeaderRouteFirst(menu.children[0], path);
    } else {
        return path;
    }
}


/**
 * 获取菜单面包屑数据
 * @param menus 菜单数据
 * @returns 菜单面包屑数据
 */
export const getBreadcrumbs = (menus: IMenu[]) => {
    const bcs: { [key: string]: string[]; } = {}

    for (let i = 0; i < menus.length; i++) {
        const menu = menus[i];
        getBreadcrumbs2(bcs, menu.children!, '/' + menu.name, [menu.desc]);
    }

    return bcs;
}

const getBreadcrumbs2 = (bcs: { [key: string]: string[]; }, menus: IMenu[], key: string, value0: string[]) => {
    for (let i = 0; i < menus.length; i++) {
        const menu = menus[i];

        const value = value0.concat([])
        value.push(menu.desc)

        if (menu.children) {
            getBreadcrumbs2(bcs, menu.children, key + '/' + menu.name, value)
        } else {
            bcs[key + '/' + menu.name] = value
        }
    }
}
