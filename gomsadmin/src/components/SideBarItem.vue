<template>
  <template v-for="menu in menus">
    <el-sub-menu v-if="menu.children" :key="menu.name" :index="menu.name">
      <template #title>
        <span>{{ menu.desc }}</span>
      </template>
      <side-bar-item :menus="menu.children"></side-bar-item>
    </el-sub-menu>
    <el-menu-item v-else :key="menu.name + 'xx'" :index="menu.route" :route="menu.route">{{ menu.desc }}</el-menu-item>
  </template>
</template>


<script lang="ts" setup>
import { ElSubMenu } from 'element-plus';

import { IMenu } from '@/data/menu'


defineProps<{ menus: IMenu[] }>()
</script>
