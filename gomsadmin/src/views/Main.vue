<template>
    <el-container>
        <el-header>
            <el-row>
                <el-col :xs="19" :sm="20" :md="21" :lg="21" :xl="22">
                    <el-menu :default-active="menuStore.activeHeader" mode="horizontal" background-color="#23262E"
                        text-color="#fff" active-text-color="#ffd04b" @select="handleHeaderSelect">
                        <el-menu-item v-for="menu in menuStore.menus" :key="menu.name" :index="menu.name">{{ menu.desc
                        }}</el-menu-item>
                    </el-menu>
                </el-col>
                <el-col :xs="5" :sm="4" :md="3" :lg="3" :xl="2" class="user-info">
                    <template v-if="authStore.userInfo">
                        <el-dropdown trigger="click" class="user-info-dd">
                            <span class="el-dropdown-link">
                                <el-icon style="vertical-align: middle">
                                    <icon-user />
                                </el-icon><span style="vertical-align: middle">{{ ' ' + authStore.userInfo.account + ' '
                                }}</span>
                                <el-icon style="vertical-align: middle">
                                    <icon-caret-bottom />
                                </el-icon>
                            </span>
                            <template #dropdown>
                                <el-dropdown-menu>
                                    <el-dropdown-item @click.native="jumpTo('/home/Welcome')">个人信息</el-dropdown-item>
                                    <el-dropdown-item @click.native="jumpTo('/sys/AdmUserPassword')">修改密码
                                    </el-dropdown-item>
                                    <el-dropdown-item divided @click.native="logout">退出登录</el-dropdown-item>
                                </el-dropdown-menu>
                            </template>
                        </el-dropdown>
                    </template>
                </el-col>
            </el-row>
        </el-header>
        <el-container>
            <el-aside width="200px">
                <the-side-bar :menus="menuStore.menusAside" @select="handleAsideSelect"></the-side-bar>
            </el-aside>
            <el-main>
                <el-breadcrumb class="main-head" :separator-icon="ArrowRight">
                    <el-breadcrumb-item v-for="(item, index) in menuStore.breadcrumbs[$route.path]" :key="index">{{ item
                    }}
                    </el-breadcrumb-item>
                </el-breadcrumb>
                <router-view />
            </el-main>
        </el-container>
    </el-container>
</template>


<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

import { Action, ElMessageBox } from 'element-plus'
import { User as IconUser, CaretBottom as IconCaretBottom, ArrowRight } from '@element-plus/icons-vue'

import useAuthStore from '@/sotre/auth'
import useMenuStore from '@/sotre/menu'

import TheSideBar from '@/components/TheSideBar.vue'


const router = useRouter()

const authStore = useAuthStore()
const menuStore = useMenuStore()


onMounted(() => {
    handleHeaderSelect(menuStore.activeHeader, undefined)
})


const handleHeaderSelect = (key: string, keyPath?: string[]) => {
    // console.log('handleHeaderSelect', key, keyPath)

    if (keyPath != undefined && key == menuStore.activeHeader) return

    menuStore.onHeaderTap(key)

    router.replace({ path: menuStore.headerRoutes[key] })
}

const handleAsideSelect = (key: string, keyPath: string[]) => {
    // console.log('handleAsideSelect', key, keyPath)

    menuStore.onAsideTap(key)
}

const jumpTo = (path: string) => {
    router.replace({ path: path })
}

const logout = () => {
    ElMessageBox.alert('确认退出吗?', '提示', {
        type: 'warning',
        confirmButtonText: '确定',
        showCancelButton: true,
        cancelButtonText: '取消',
        callback: async (action: Action) => {
            if (action != 'confirm') return

            await authStore.logout()
            menuStore.destory()
            router.replace({ name: 'Login' })
        },
    })
}
</script>


<style scoped>
.el-header {
    padding: 0px;
}

.el-aside {
    background-color: #393d49;
}

.el-main {
    padding: 0px;
}

.user-info {
    padding-right: 15px;
    margin-bottom: 1px;
    background-color: #23262e;

    display: -webkit-flex;
    display: flex;
    justify-content: flex-end;
    /*子元素在水平轴上的对齐方式*/
    align-items: center;
    /*子元素在垂直轴上的对齐方式*/
}

.user-info-dd {
    /* align-self: center; */
    /*如果你元素设置了aligh-items为center就不用设置了*/
    color: #ffffff;
}

.main-head {
    height: 32px;
    line-height: 32px;
    padding: 0 10px;
    background-color: #eee;
}

.el-main>.container {
    padding: 9px 15px 15px;
    text-align: left;
}
</style>
