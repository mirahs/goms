<template>
    <el-form ref="formRef" :model="formModel" :rules="rules" status-icon class="login-page">
        <el-form-item prop="account">
            <el-input v-model="formModel.account" type="text" autocomplete="off" placeholder="账号" />
        </el-form-item>
        <el-form-item prop="password">
            <el-input @keyup.enter.native="onClickLogin(formRef)" v-model="formModel.password" type="password" show-password autocomplete="off" placeholder="密码" />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" style="width:100%" @click="onClickLogin(formRef)" :loading="logining">登录
            </el-button>
        </el-form-item>
    </el-form>
</template>


<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'

import type { FormInstance, FormRules } from 'element-plus'

import useAuthStore from '@/sotre/auth'


const router = useRouter()

const authStore = useAuthStore()

const formRef = ref<FormInstance>()

const formModel = reactive({
    account: '',
    password: '',
})

let logining = ref(false)

const rules = reactive<FormRules>({
    account: [
        { required: true, message: "请输入账号", trigger: "blur" },
        { min: 2, message: "账号不能少于 2 个字符", trigger: "blur" },
    ],
    password: [
        { required: true, message: "请输入密码", trigger: "blur" },
        { min: 5, message: "密码不能少于 5 个字符", trigger: "blur" },
    ],
})


const onClickLogin = (formEl: FormInstance | undefined) => {
    if (!formEl) return

    formEl.validate(async (valid) => {
        if (!valid) return false

        logining.value = true
        var res = await authStore.login(formModel)
        logining.value = false

        if (res) {
          router.replace({ path: "/" }).catch((err) => {});
        }
    })
}
</script>


<style>
.login-page {
    margin: 180px auto;
    width: 300px;
}
</style>
