<template>
  <div
    class="fullscreen bg-blue text-white text-center q-pa-md flex flex-center"
  >
    <div>
      <div style="font-size: 30vh">404</div>

      <div class="text-h2" style="opacity: 0.4">Oops. Nothing here...</div>

      <q-btn
        class="q-mt-xl"
        color="white"
        text-color="blue"
        unelevated
        to="/"
        label="Go Home"
        no-caps
      />
    </div>
  </div>
</template>

<script setup>
import { onMounted } from "vue";
import { GetU } from "boot/config";

onMounted(() => {
  // 处理已登录用户: 无需登录,跳转到/
  const u = GetU();
  if (!u) {
    utils.goto("/login");
  }
});
</script>
