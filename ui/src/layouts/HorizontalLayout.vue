<template>
  <q-layout view="hHh lpR fFf" class="shadow-2 rounded-borders">
    <q-header elevated class="bg-primary text-white" height-hint="98">
      <!--
QToolbar: 如果你需要一个复杂且多功能的工具栏（如导航栏）
QBar: 如果你只需要一个简单的操作条（如确认/取消按钮）

  属性	        QToolbar	                  QBar
----------------------------------------------------
复杂性    适合复杂布局，支持多种组件	       更轻量，适合简单操作栏
常用场景  应用顶部工具栏（导航栏、标题栏等）   卡片、弹窗或底部的操作栏
搭配组件  通常与 QHeader 搭配使用	          通常与 QCard 或 QDialog 使用
灵活性    支持较复杂的内容，如标题、搜索框	   主要用于按钮和简单操作
----------------------------------------------------

https://quasar.dev/layout/header-and-footer#desktop-app-look-and-feel
      -->
      <q-toolbar style="height: 64px">
        <!-- <q-btn flat round dense icon="menu" class="q-mr-sm" /> -->
        <q-avatar>
          <img src="../assets/logo-mono-white.svg" />
        </q-avatar>

        <q-toolbar-title>tolo</q-toolbar-title>

        <HeaderUser2 />
        <q-btn dense flat round icon="menu" @click="toggleRightDrawer" />
      </q-toolbar>
      <!-- tabs demo -->
      <!-- <q-tabs align="left">
        <q-route-tab to="/page1" label="Page One" />
        <q-route-tab to="/page2" label="Page Two" />
        <q-route-tab to="/page3" label="Page Three" />
      </q-tabs> -->
    </q-header>

    <q-drawer v-model="rightDrawerOpen" side="right" bordered>
      <!-- drawer content -->
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>

    <!-- <q-footer elevated class="bg-grey-8 text-white">
      <q-toolbar>
        <q-toolbar-title>
          <q-avatar>
            <img src="https://cdn.quasar.dev/logo-v2/svg/logo-mono-white.svg" />
          </q-avatar>
          <div>Title</div>
        </q-toolbar-title>
      </q-toolbar>
    </q-footer> -->
  </q-layout>
</template>
// icons 选择
// https://quasar.dev/vue-components/icon#introduction
// https://fonts.google.com/icons?selected=Material+Icons


<script setup>
import { ref } from "vue";
import HeaderUser2 from "components/HeaderUser2.vue";
const rightDrawerOpen = ref(false);

const toggleRightDrawer = () => {
  rightDrawerOpen.value = !rightDrawerOpen.value;
};
</script>

