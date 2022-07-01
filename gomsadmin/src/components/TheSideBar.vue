<template>
  <el-menu background-color="#393D49" text-color="#fff" active-text-color="#ffd04b" router :default-active="$route.path"
    @select="onSelect">
    <side-bar-item :menus="menus"></side-bar-item>
  </el-menu>
</template>


<script lang="ts" setup>
import SideBarItem from "@/components/SideBarItem.vue";
import { IMenu } from "@/data/menu";


defineProps<{ menus: IMenu[] }>()

const emit = defineEmits(['select'])


const onSelect = (key: string, keyPath: string[]) => {
  emit('select', key, keyPath)
}
</script>
