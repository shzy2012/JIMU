
<template>
  <q-layout>
    <q-page-container>
      <q-page
        class="bg-light-blue-2 window-height row justify-center items-center"
      >
        <q-card class="my-card">
          <q-card-section>
            <div class="text-h6 text-center">tolo</div>
          </q-card-section>
          <q-separator />
          <q-card-actions vertical>
            <q-card-section>
              <q-form class="q-gutter-md">
                <q-input
                  square
                  filled
                  v-model="phone"
                  type="text"
                  label="手机号"
                />
                <q-input
                  square
                  filled
                  v-model="password"
                  type="password"
                  label="密码"
                />
              </q-form>
            </q-card-section>
            <q-card-actions class="q-px-md">
              <q-btn
                unelevated
                color="primary"
                size="lg"
                class="full-width"
                label="Login"
                @click="onLogin()"
              />
            </q-card-actions>
          </q-card-actions>
        </q-card>
      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from "vue";
import { useQuasar } from "quasar";
import { UserService } from "boot/user";
import { sha256 } from "js-sha256";
import { GetU, SetU } from "boot/config";

const $q = useQuasar();
let timer;

onBeforeUnmount(() => {
  if (timer !== void 0) {
    clearTimeout(timer);
    $q.loading.hide();
  }
});

const phone = ref("");
const password = ref("");

const onLogin = async () => {
  $q.loading.show({
    message: "Login in progress. Hang on...",
  });

  const param = { phone: phone.value, passwd: sha256(password.value) };
  await UserService.login(param)
    .then((res) => {
      SetU(res.data);
      utils.goto("/");
    })
    .catch(({ response }) => {
      if (response) {
        $q.notify({
          message: response.data,
          icon: "report_problem",
        });
      }
    });

  $q.loading.hide();
};

onMounted(() => {
  // 处理已登录用户: 无需登录,跳转到/
  const u = GetU();
  if (u) {
    utils.goto("/");
  }
});
</script>

<style>
.q-card {
  width: 360px;
}
</style>
