<template>
    <div class="container">
        <el-form ref="formRef" :model="model" :rules="rules" status-icon label-width="80px">
            <el-form-item label="主机名" prop="name">
                <el-input v-model="model.name" placeholder="请输入主机名"></el-input>
            </el-form-item>

            <el-form-item label="ssh端口" prop="ssh_port">
                <el-input v-model.number="model.ssh_port" placeholder="请输入ssh端口号"></el-input>
            </el-form-item>
            <el-form-item label="ssh账号" prop="ssh_username">
                <el-input v-model="model.ssh_username" placeholder="请输入ssh账号"></el-input>
            </el-form-item>
            <el-form-item label="ssh密码" prop="ssh_password">
                <el-input v-model="model.ssh_password" placeholder="请输入ssh密码"></el-input>
            </el-form-item>

            <el-form-item label="备注" prop="remark">
                <el-input type="textarea" v-model="model.remark"></el-input>
            </el-form-item>

            <el-form-item>
                <el-button type="primary" @click="onSubmit(formRef)">立即提交</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>


<script lang="ts" setup>
import { reactive, ref } from 'vue'

import { FormInstance, FormRules } from 'element-plus'

import { IHost } from '@/intfs'


const props = defineProps<{ model: IHost }>()


const formRef = ref<FormInstance>()


const ruleAge = (rule: any, value: number, callback: any) => {
    if (!value) return callback()

    if (value < 1) {
        callback(new Error('ssh端口号不能小于 1'))
    } else if (value > 65535) {
        callback(new Error('ssh端口号不能大于 65535'))
    } else {
        callback()
    }
}

const rules = reactive<FormRules>({
    name: [
        { required: true, message: "请输入主机名", trigger: "blur" },
        { min: 2, message: "主机名不能少于 2 个字符", trigger: "blur" },
        { max: 30, message: "主机名太长", trigger: "blur" },
    ],

    ssh_port: [
        { type: 'number', message: 'ssh端口号必须为数字' },
        { validator: ruleAge, trigger: 'blur' },
    ],
    ssh_username: [
        { max: 30, message: "ssh账号太长", trigger: "blur" },
    ],
    ssh_password: [
        { max: 50, message: "ssh密码太长", trigger: "blur" },
    ],

    remark: [
        { required: true, message: "请输入备注信息", trigger: "blur" },
        { min: 2, message: "备注不能少于 2 个字符", trigger: "blur" },
    ],
})


const emit = defineEmits(['submit'])

const onSubmit = (formEl: FormInstance | undefined) => {
    if (!formEl) return

    formEl.validate(async (valid) => {
        if (!valid) return

        emit('submit', props.model)
    })
}
</script>
