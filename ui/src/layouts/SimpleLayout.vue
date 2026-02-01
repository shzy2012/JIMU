<template>
  <q-layout view="hHh Lpr fFf">
    <q-header elevated>
      <q-toolbar style="height: 64px">
        <q-btn flat round dense icon="menu" @click="toggleLeftDrawer" />

        <q-toolbar-title class="row items-center no-wrap" @click="router.push('/')">
          <span class="text-weight-bold tracking-tight text-h6">
            TOLO
          </span>
        </q-toolbar-title>

        <q-space />

        <div class="row items-center q-gutter-x-sm">
          <q-separator vertical inset class="q-mx-md" />
          <HeaderUser />
        </div>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      side="left"
      show-if-above
      bordered
      class="bg-white menu-drawer"
      :width="260"
    >
      <q-scroll-area class="fit">
        <div class="q-pa-md">
          <div class="text-overline text-grey-6 q-px-sm q-mb-sm">MENU</div>

          <q-list class="q-gutter-y-xs">
            <template v-for="(item, index) in menuItems" :key="index">
              <q-expansion-item
                v-if="item.children && item.children.length > 0"
                :icon="item.icon"
                :label="item.title"
                :default-opened="hasActiveChild(item)"
                header-class="text-weight-medium text-grey-8 rounded-borders menu-item-header"
                expand-icon-class="text-grey-6"
                class="rounded-borders overflow-hidden"
              >
                <div class="q-pl-md q-pt-xs">
                    <EssentialLink
                      v-for="(child, childIndex) in item.children"
                      :key="childIndex"
                      v-bind="child"
                      class="rounded-borders menu-item-link"
                      active-class="bg-blue-1 text-primary text-weight-bold"
                      exact
                    />
                </div>
              </q-expansion-item>

              <EssentialLink
                v-else
                v-bind="item"
                class="rounded-borders menu-item-link"
                active-class="bg-blue-1 text-primary text-weight-bold"
                exact
              />
            </template>
          </q-list>
        </div>
      </q-scroll-area>
    </q-drawer>

    <q-drawer
      v-model="RightDrawerOpen"
      side="right"
      bordered
      width="400"
      overlay
      class="bg-white shadow-lg"
    >
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import EssentialLink from "components/EssentialLink.vue";
import { onMounted, ref } from "vue";
import { useRoute, useRouter } from "vue-router";
import { RightDrawerOpen } from "boot/config";
import HeaderUser from "components/HeaderUser.vue";
// icons 选择
// https://quasar.dev/vue-components/icon#introduction
// https://fonts.google.com/icons?selected=Material+Icons

const router = useRouter();
const route = useRoute();

const leftDrawerOpen = ref(false);

const menuItems = ref([]);

const toggleLeftDrawer = () => {
  leftDrawerOpen.value = !leftDrawerOpen.value;
};

// Check if any child link is currently active
const hasActiveChild = (item) => {
  if (!item.children) return false;
  return item.children.some((child) => child.link === route.path);
};

onMounted(() => {
  menuItems.value = [
    {
      title: "首页",
      icon: "home",
      link: "/",
    },
    {
      title: "示例",
      icon: "description",
      link: "/example",
    },
    {
      title: "出库",
      icon: "inventory_2",
      link: "/stockout",
    },
  ];
});
</script>

<style scoped>
.tracking-tight {
  letter-spacing: -0.025em;
}

/* Sidebar Menu Styling */
.menu-drawer :deep(.q-drawer__content) {
  background: #ffffff;
}

.menu-item-header {
  border-radius: 8px;
  transition: all 0.2s ease;
}
.menu-item-header:hover {
  background: #f8f9fa;
  color: #1976d2 !important;
}

.menu-item-link {
  transition: all 0.2s ease;
  margin-bottom: 2px;
  color: #5f6368;
}

.menu-item-link:hover {
  background: #f1f3f4;
  color: #1a1a1a;
}
</style>
