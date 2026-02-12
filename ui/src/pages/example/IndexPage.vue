<template>
  <div class="q-pa-lg">
    <div class="row">
      <div class="col col-md-6 col-sm-6 col-xs-12">
        <q-input label="股票代码/名称" v-model="filter.ts_code"> </q-input>
      </div>
      <div class="col col-md-6 col-sm-6 col-xs-12">
        <q-btn
          color="primary"
          label="搜索"
          class="q-mr-xs"
          @click="onRefresh()"
        />
        <q-btn
          color="deep-orange"
          label="清空"
          class="q-mr-xs"
          @click="onClear()"
        />
      </div>
    </div>
  </div>

  <div class="q-pa-md">
    <q-table
      title="股票列表示例"
      :rows="rows"
      :columns="columns"
      row-key="id"
      :loading="loading"
      @request="onRequest"
      v-model:pagination="pagination"
      :filter="filter"
      binary-state-sort
    >
      <template v-slot:top-right>
        <q-btn
          label="刷新"
          color="primary"
          class="q-mr-xs"
          icon="refresh"
          @click="onRefresh()"
        />
      </template>
    </q-table>
  </div>
</template>

<script setup>
import { useQuasar } from "quasar";
import { ref, onMounted } from "vue";

const $q = useQuasar();
const filter = ref({
  ts_code: "",
});
const loading = ref(false);
const pagination = ref({
  sortBy: "desc",
  descending: false,
  page: 0,
  rowsPerPage: 30,
  rowsNumber: 0,
});

const rows = ref([]);
const columns = [
  { name: "ts_code", label: "股票代码", field: "ts_code", align: "left" },
  { name: "name", label: "名称", field: "name", align: "left" },
  { name: "area", label: "地域", field: "area", align: "left" },
  { name: "industry", label: "所属行业", field: "industry", align: "left" },
  { name: "trade_date", label: "交易日期", field: "trade_date", align: "left" },
  { name: "close", label: "收盘价", field: "close", align: "left" },
  { name: "pct_chg", label: "涨跌幅", field: "pct_chg", align: "left" },
  { name: "vol", label: "成交量", field: "vol", align: "left" },
];

const onRefresh = () => {
  onRequest({
    pagination: pagination.value,
  });
};

const onRequest = async (props) => {
  loading.value = true;
  try {
    // 示例页面：暂无数据源
    rows.value = [];
    pagination.value.rowsNumber = 0;
    const { page, rowsPerPage, sortBy, descending } = props.pagination;
    pagination.value.page = page;
    pagination.value.rowsPerPage = rowsPerPage;
    pagination.value.sortBy = sortBy;
    pagination.value.descending = descending;
  } catch (error) {
    $q.notify({
      message: error?.response?.data || "加载失败",
      icon: "report_problem",
    });
  } finally {
    loading.value = false;
  }
};

const onClear = () => {
  filter.value.ts_code = "";
};

onMounted(() => {
  onRequest({
    pagination: pagination.value,
  });
});
</script>
