export interface IMenu {
    name: string
    desc: string

    key?: number[] //菜单权限
    children?: IMenu[] //多级菜单

    route?: string //Element Plus 菜单路由
}


export const menus: IMenu[] = [
    {
        name: 'home',
        desc: '首页',
        children: [
            {
                name: 'Welcome',
                desc: '欢迎',
                key: [1, 2]
            }
        ],
    },
    {
        name: 'sys',
        desc: '系统',
        children: [
            {
                name: 'AdmUserPassword',
                desc: '密码更新',
                key: [1, 2]
            },
            {
                name: 'AdmUser',
                desc: '管理员',
                key: [1]
            },
        ],
    },
    {
        name: 'log',
        desc: '日志',
        children: [
            {
                name: 'AdmUserLogin',
                desc: '管理员登录',
                key: [1],
            },
        ],
    },
    {
        name: 'ms',
        desc: 'MS',
        children: [
            {
                name: 'Host',
                desc: '主机',
                key: [1],
            },
        ],
    },
]
