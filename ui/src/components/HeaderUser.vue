<template>
  <div class="q-gutter-sm row items-center no-wrap">
    <q-btn dense flat no-wrap>
      <q-avatar rounded size="26px">
        <img src="~/assets/boy-avatar.png" />
      </q-avatar>
      <q-menu auto-close>
        <q-list dense>
          <q-item>
            <q-item-section>
              <div>
                当前用户 <strong>{{ name }}</strong>
              </div>
            </q-item-section>
          </q-item>
          <q-separator />
          <q-item clickable>
            <q-item-section>个人设置</q-item-section>
          </q-item>
          <q-item clickable @click="logout()">
            <q-item-section>退出</q-item-section>
          </q-item>
        </q-list>
      </q-menu>
    </q-btn>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { GetU, DelU } from "boot/config";
import { utils } from "src/utils";

const name = ref("");

const logout = function () {
  DelU();
  if (window.location.pathname.includes("open")) {
    utils.goto("/open");
  } else {
    utils.goto("/login");
  }
};

onMounted(() => {
  const u = GetU();
  if (u) {
    name.value = u.phone;
  }
});
</script>
