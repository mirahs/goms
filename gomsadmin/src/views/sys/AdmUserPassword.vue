<template>
  <div class="container">
    <el-form ref="formRef" :model="formModel" :rules="rules" status-icon label-width="100px">
      <el-form-item label="账号">
        <el-col :span="6">
          <el-input v-model="formModel.account" readonly></el-input>
        </el-col>
      </el-form-item>

      <el-form-item label="新密码" prop="password">
        <el-col :span="6">
          <el-input type="password" v-model="formModel.password" placeholder="请输入新密码" show-password
            autocomplete="new-password">
          </el-input>
        </el-col>
      </el-form-item>

      <el-form-item label="确认新密码" prop="repassword">
        <el-col :span="6">
          <el-input @keyup.enter.native="onSubmit(formRef)" type="password" v-model="formModel.repassword" placeholder="请输入确认新密码" show-password
            autocomplete="new-password">
          </el-input>
        </el-col>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit(formRef)" :loading="loading">立即提交</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>


<script lang="ts" setup>
import { reactive, ref } from 'vue'

import { FormInstance, FormRules, ElMessage } from 'element-plus'

import useAuthStore from '@/sotre/auth'
import * as userApi from '@/api/user'


const authStore = useAuthStore()

const formRef = ref<FormInstance>()

const formModel = reactive({
  account: authStore.userInfo.account,
  password: '',
  repassword: '',
})

let loading = ref(false)


const ruleValidatePasswordSame = (rule: any, value: any, callback: any) => {
  if (value !== formModel.password) {
    callback(new Error('密码不一致'))
  } else {
    callback()
  }
}

const rules = reactive<FormRules>({
  password: [
    { required: true, message: "请输入新密码", trigger: "blur" },
    { min: 5, message: "密码不能少于 5 个字符", trigger: "blur" },
  ],
  repassword: [
    { required: true, message: "请输入确认新密码", trigger: "blur" },
    { min: 5, message: "密码不能少于 5 个字符", trigger: "blur" },
    { validator: ruleValidatePasswordSame, trigger: "blur" },
  ]
})


const onSubmit = (formEl: FormInstance | undefined) => {
  if (!formEl) return

  formEl.validate(async (valid) => {
    if (!valid) return

    loading.value = true
    var res = await userApi.password(formModel)
    loading.value = false

    if (res.code != 0) return

    ElMessage.success('密码更新成功')

    formModel.password = ''
    formModel.repassword = ''
  })
}
</script>
