<template>
  <div class="container">
    <el-table :data="datas" border stripe style="width: 100%">
    <el-table-column prop="id" label="ID" fixed> </el-table-column>

      <el-table-column prop="name" label="主机名" fixed> </el-table-column>
      <el-table-column prop="created_at" label="添加时间" :formatter="formatDate"></el-table-column>
      <el-table-column prop="remark" label="备注"> </el-table-column>

      <el-table-column fixed="right" align="center" label="操作" width="180">
        <template #header>
          <el-button @click="onClickAdd" type="primary" :icon="IconPlus" size="small">新增</el-button>
        </template>
        <template #default="scope">
          <el-button @click="onClickEdit(scope.row)" type="primary" :icon="IconEdit" circle size="small"></el-button>
          <el-button @click="onClickDelete(scope.row)" type="danger" :icon="IconDelete" circle size="small"></el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :page-sizes="[5, 10, 20]"
      :page-size="pageSize" :current-page="pageCurrent" :total="count" layout="total, sizes, prev, pager, next, jumper">
    </el-pagination>

    <el-dialog :title="model != null ? '主机更新' : '主机添加'" v-model="editShow" center width="30%"
      :close-on-click-modal="false">
      <host-edit :model="model" @submit="editSubmit" />
    </el-dialog>
  </div>
</template>


<script lang="ts" setup>
import { ref } from 'vue'
import moment from 'moment'

import { Action, ElMessage, ElMessageBox } from 'element-plus'
import { Plus as IconPlus, Edit as IconEdit, Delete as IconDelete } from '@element-plus/icons-vue'

import HostEdit from './components/HostEdit.vue'

import { IHost } from '@/intfs'
import * as hostApi from '@/api/host'


const pageCurrent = ref<number>(1)
const pageSize = ref<number>(10)

const count = ref<number>(0)
const datas = ref<IHost[]>([])

const editShow = ref<boolean>(false)
const model = ref<IHost>(<IHost>{})



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
  model.value = <IHost>{}
}

const onClickEdit = (row: IHost) => {
  const data = getDataById(row.id)
  if (data == undefined) return

  editShow.value = true

  model.value = <IHost>data
}

const editSubmit = async (model: IHost) => {
  const model2 = { ...model }

  if (model2.id == undefined) {
    const res = await hostApi.add(model2)
    if (res.code != 0) return

    editShow.value = false
    loadDatas()

    ElMessage.success('主机添加成功')
  } else {
    const res = await hostApi.edit(model2.id, model2)
    if (res.code != 0) return

    editShow.value = false
    loadDatas()

    ElMessage.success('主机更新成功')
  }
}

const onClickDelete = (row: IHost) => {
  ElMessageBox.alert(row.name + " 主机将永久删除，是否继续？", '提示', {
    type: 'warning',
    confirmButtonText: '确定',
    showCancelButton: true,
    cancelButtonText: '取消',
    callback: async (action: Action) => {
      if (action != 'confirm') return

      const res = await hostApi.del(row.id)
      if (res.code != 0) return

      loadDatas()

      ElMessage.success('主机删除成功')
    },
  })
}


const loadDatas = async () => {
  const res = await hostApi.getAll(pageCurrent.value, pageSize.value)
  if (res.code != 0) return

  count.value = res.data.count
  datas.value = <IHost[]>res.data.datas
}

const getDataById = (id: number): IHost | undefined => {
  for (let i = 0; i < datas.value.length; i++) {
    if (datas.value[i].id == id) {
      return datas.value[i];
    }
  }
}

const formatDate = (row: any, column: any) => {
  return moment(row[column.property] * 1000).format("YYYY-MM-DD HH:mm:ss")
}


loadDatas()

</script>
