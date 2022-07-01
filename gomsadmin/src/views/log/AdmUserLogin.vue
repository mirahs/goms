<template>
  <div class="container">
    <el-form :inline="true" :model="search">
      <el-form-item label="账号">
        <el-input v-model="search.account" placeholder="管理员账号" style="width:100px"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSearch">查询</el-button>
      </el-form-item>
    </el-form>

    <el-table :data="datas" border stripe style="width: 100%">
      <el-table-column prop="account" label="账号" fixed> </el-table-column>
      <el-table-column prop="type" label="登录类型" fixed :formatter="formatLoginType"> </el-table-column>
      <el-table-column prop="time" label="时间" fixed :formatter="formatDate"></el-table-column>

      <el-table-column prop="status" label="状态" :formatter="formatStatus"> </el-table-column>
      <el-table-column prop="ip" label="IP地址"></el-table-column>
      <el-table-column prop="address" label="地址"></el-table-column>

      <el-table-column fixed="right" align="center" label="操作" width="60">
        <template #default="scope">
          <el-button @click="onClickDelete(scope.row)" type="danger" :icon="IconDelete" circle size="small"></el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :page-sizes="[5, 10, 20]"
      :page-size="pageSize" :current-page="pageCurrent" :total="count" layout="total, sizes, prev, pager, next, jumper">
    </el-pagination>
  </div>
</template>


<script lang="ts" setup>
import { ref } from 'vue'
import moment from 'moment'

import { ElMessage } from 'element-plus'
import { Delete as IconDelete } from '@element-plus/icons-vue'

import useAuthStore from '@/sotre/auth'
import { IUserLoginLog, IUserSearch } from '@/intfs'
import * as logApi from '@/api/log'


const search = ref<IUserSearch>(<IUserSearch>({}))

const pageCurrent = ref<number>(1)
const pageSize = ref<number>(10)

const count = ref<number>(0)
const datas = ref<IUserLoginLog[]>([])


const onSearch = async () => {
  await loadDatas()
}


const handleSizeChange = (val: number) => {
  pageSize.value = val
  loadDatas()
}

const handleCurrentChange = (val: number) => {
  pageCurrent.value = val
  loadDatas()
}


const onClickDelete = async (row: IUserLoginLog) => {
  const res = await logApi.delLoginAdmUser(row.id)
  if (res.code != 0) return

  loadDatas()

  ElMessage.success('日志删除成功')
}


const loadDatas = async () => {
  const res = await logApi.getLoginAdmUser(pageCurrent.value, pageSize.value, search.value)
  if (res.code != 0) return

  count.value = res.data.count
  datas.value = <IUserLoginLog[]>res.data.datas
}

const formatLoginType = (row: any, column: any) => {
  return row[column.property] == 1 ? "账号登录" : "token登录";
}

const formatStatus = (row: any, column: any) => {
  return row[column.property] == 0 ? "失败" : "成功";
}

const formatDate = (row: any, column: any) => {
  return moment(row[column.property] * 1000).format("YYYY-MM-DD HH:mm:ss")
}


loadDatas()

</script>
