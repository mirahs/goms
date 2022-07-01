<template>
  <div class="container">
    <el-table :data="datas" border stripe style="width: 100%">
      <el-table-column prop="account" label="账号" fixed> </el-table-column>
      <el-table-column prop="type" label="类型" :formatter="formatType"> </el-table-column>
      <el-table-column prop="created_at" label="创建时间" :formatter="formatDate">
      </el-table-column>
      <el-table-column prop="login_times" label="登录次数"> </el-table-column>
      <el-table-column prop="login_at" label="最后登陆时间" :formatter="formatDate">
      </el-table-column>
      <el-table-column prop="login_ip" label="最后登陆IP"> </el-table-column>

      <el-table-column prop="locked" label="锁定">
        <template #default="scope">
          <el-switch v-model="scope.row.locked" active-color="#13ce66" @change="onChangeLocked(scope.row.id)" />
        </template>
      </el-table-column>

      <el-table-column prop="remark" label="备注"> </el-table-column>

      <el-table-column fixed="right" align="center" label="操作" width="180">
        <template #header>
          <el-button @click="onClickAdd" type="primary" :icon="IconPlus" size="small">新增</el-button>
        </template>
        <template #default="scope">
          <el-button @click="onClickReset(scope.row)" type="warning" size="small">重置密码</el-button>
          <el-button @click="onClickEdit(scope.row)" type="primary" :icon="IconEdit" circle size="small"></el-button>
          <el-button @click="onClickDelete(scope.row)" type="danger" :icon="IconDelete" circle size="small"></el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :page-sizes="[5, 10, 20]"
      :page-size="pageSize" :current-page="pageCurrent" :total="count" layout="total, sizes, prev, pager, next, jumper">
    </el-pagination>

    <el-dialog :title="model != null ? '管理员更新' : '管理员添加'" v-model="editShow" center width="30%"
      :close-on-click-modal="false">
      <adm-user-edit :model="model" @submit="editSubmit" />
    </el-dialog>
  </div>
</template>


<script lang="ts" setup>
import { ref } from 'vue'
import moment from 'moment'
import useClipboard from 'vue-clipboard3'

import { Action, ElMessage, ElMessageBox } from 'element-plus'
import { Plus as IconPlus, Edit as IconEdit, Delete as IconDelete } from '@element-plus/icons-vue'

import AdmUserEdit from './components/AdmUserEdit.vue'

import useAuthStore from '@/sotre/auth'
import { IUser } from '@/intfs'
import Constant from '@/constant'
import * as userApi from '@/api/user'


const authStore = useAuthStore()

const pageCurrent = ref<number>(1)
const pageSize = ref<number>(10)

const count = ref<number>(0)
const datas = ref<IUser[]>([])

const editShow = ref<boolean>(false)
const model = ref<IUser>(<IUser>{})



const handleSizeChange = (val: number) => {
  pageSize.value = val
  loadDatas()
}

const handleCurrentChange = (val: number) => {
  pageCurrent.value = val
  loadDatas()
}


const onClickAdd = () => {
  editShow.value = true
  model.value = <IUser>{}
}

const onClickEdit = (row: IUser) => {
  const data = getDataById(row.id)
  if (data == undefined) return

  editShow.value = true

  data.type2 = data.type + ''

  model.value = <IUser>data
}

const editSubmit = async (model: IUser) => {
  const model2 = { ...model }
  model2.type = parseInt(model2.type2)

  if (model2.id == undefined) {
    const res = await userApi.add(model2)
    if (res.code != 0) return

    editShow.value = false
    loadDatas()

    const password = res.data
    ElMessageBox.alert("密码为 " + password, "管理员添加成功", {
      type: 'success',
      confirmButtonText: '确定',
      callback: async (action: Action) => {
        const { toClipboard } = useClipboard()
        try {
          await toClipboard(password)
          ElMessage.success('密码已复制到剪贴板')
        } catch (e) {
          ElMessage.error('密码复制失败')
        }
      },
    })
  } else {
    const res = await userApi.edit(model2.id, model2)
    if (res.code != 0) return

    editShow.value = false
    loadDatas()

    ElMessage.success('管理员更新成功')
  }
}

const onClickReset = (row: IUser) => {
  if (row.id == authStore.userInfo.id) {
    ElMessage.error('不能操作自己')
    return
  }

  ElMessageBox.alert(row.account + " 用户密码将重置，是否继续？", '提示', {
    type: 'warning',
    confirmButtonText: '确定',
    showCancelButton: true,
    cancelButtonText: '取消',
    callback: async (action: Action) => {
      if (action != 'confirm') return

      const res = await userApi.reset(row.id)
      if (res.code != 0) return

      const password = res.data

      ElMessageBox.alert("密码为 " + password, "重置用户密码成功", {
        type: 'success',
        confirmButtonText: '确定',
        callback: async (action: Action) => {
          const { toClipboard } = useClipboard()
          try {
            await toClipboard(password)
            ElMessage.success('密码已复制到剪贴板')
          } catch (e) {
            ElMessage.error('密码复制失败')
          }
        },
      })
    },
  })
}

const onClickDelete = (row: IUser) => {
  if (row.id == authStore.userInfo.id) {
    ElMessage.error('不能操作自己')
    return
  }

  ElMessageBox.alert(row.account + " 用户将永久删除，是否继续？", '提示', {
    type: 'warning',
    confirmButtonText: '确定',
    showCancelButton: true,
    cancelButtonText: '取消',
    callback: async (action: Action) => {
      if (action != 'confirm') return

      const res = await userApi.del(row.id)
      if (res.code != 0) return

      loadDatas()

      ElMessage.success('用户删除成功')
    },
  })
}

const onChangeLocked = async (id: number) => {
  // 回调时数据已经改变了，这里把改变过后的值发送给后端
  const data = getDataById(id)
  if (data == undefined) return

  const res = await userApi.lock(id)
  if (res.code != 0) {
    data.locked = !data.locked
  }
}


const loadDatas = async () => {
  const res = await userApi.getAll(pageCurrent.value, pageSize.value)
  if (res.code != 0) return

  count.value = res.data.count
  datas.value = <IUser[]>res.data.datas
}

const getDataById = (id: number): IUser | undefined => {
  for (let i = 0; i < datas.value.length; i++) {
    if (datas.value[i].id == id) {
      return datas.value[i];
    }
  }
}

const formatType = (row: any, column: any) => {
  return Constant.adminUserTypeDesc[row[column.property]]
}

const formatDate = (row: any, column: any) => {
  return moment(row[column.property] * 1000).format("YYYY-MM-DD HH:mm:ss")
}


loadDatas()

</script>
