<template>
  <q-page class="bg-grey-2">
    <!-- View 1: Order List -->
    <div v-if="viewMode === 'list'" class="column full-height">
      <div class="bg-primary text-white q-pa-md shadow-2 z-top">
        <div class="text-h6">出库单列表</div>
        <div class="text-subtitle2 text-blue-2">Stockout Orders</div>
      </div>

      <div class="col scroll q-pa-sm">
        <q-inner-loading :showing="loading">
          <q-spinner color="primary" size="50px" />
        </q-inner-loading>

        <q-list separator class="bg-white rounded-borders shadow-1" v-if="!loading">
          <q-item
            v-for="order in orderList"
            :key="order.doc_no"
            clickable
            v-ripple
            @click="openOrder(order)"
            class="q-py-md"
          >
            <q-item-section avatar>
              <q-avatar color="blue-1" text-color="primary" icon="assignment_turned_in" />
            </q-item-section>

            <q-item-section>
              <q-item-label class="text-weight-bold text-body1">{{ order.doc_no }}</q-item-label>
              <q-item-label caption lines="1">
                <q-icon name="event" class="q-mr-xs" />{{ order.doc_date }}
              </q-item-label>
            </q-item-section>

            <q-item-section side>
              <q-icon name="chevron_right" color="grey-5" />
            </q-item-section>
          </q-item>
        </q-list>

        <div v-if="!loading && orderList.length === 0" class="text-center q-pa-lg text-grey">
          暂无数据
        </div>
      </div>
    </div>

    <!-- View 2: Order Detail -->
    <div v-else class="column full-height">
      <div class="bg-primary text-white q-pa-md shadow-2 z-top row items-center">
        <q-btn flat round icon="arrow_back" @click="viewMode = 'list'" class="q-mr-sm" dense />
        <div>
          <div class="text-subtitle1 text-weight-bold">{{ currentOrder.doc_no }}</div>
          <div class="text-caption text-blue-2">{{ currentOrder.doc_date }}</div>
        </div>
      </div>

      <div class="col scroll q-pa-sm">
        <q-inner-loading :showing="loading">
          <q-spinner color="primary" size="50px" />
        </q-inner-loading>

        <div v-if="!loading">
          <div class="text-caption q-mb-sm text-grey-8 q-pl-xs">
            订单详情 ({{ currentOrderDetails.length }})
          </div>

          <div class="q-gutter-y-sm">
            <q-card v-for="(item, index) in currentOrderDetails" :key="index" flat bordered class="order-item-card">
              <q-card-section class="q-pb-xs">
                <div class="row no-wrap justify-between items-start">
                  <div class="col">
                    <div class="text-body1 text-weight-bold">{{ item.product_name }}</div>
                    <div class="text-caption text-grey-7" v-if="item.spec && item.spec.trim()">
                      规格型号: {{ item.spec }}
                    </div>
                    <div class="text-caption text-grey-7" v-else>
                      物料编码: {{ item.material_no }}
                    </div>
                  </div>
                  <div class="col-auto text-right">
                    <q-badge outline color="primary" :label="`序号: ${item.seq}`" />
                  </div>
                </div>
              </q-card-section>

              <q-separator inset spaced />

              <q-card-section class="q-pt-xs">
                <div class="row q-col-gutter-sm">
                  <div class="col-6">
                    <div class="text-caption text-grey">批号</div>
                    <div class="text-body2">{{ item.batch_no || '无' }}</div>
                  </div>
                  <div class="col-6 text-right">
                    <div class="text-caption text-grey">单位</div>
                    <div class="text-body2">{{ item.unit || '无' }}</div>
                  </div>
                </div>
                <div class="row q-col-gutter-sm q-mt-sm">
                  <div class="col-6">
                    <div class="text-caption text-grey">数量</div>
                    <div class="text-body1 text-weight-bold text-primary">
                      {{ item.must_qty }} {{ item.unit }}
                    </div>
                  </div>
                  <div class="col-6 text-right">
                    <div class="text-caption text-grey">实发数量</div>
                    <div class="text-body1 text-weight-bold" :class="item.real_qty < item.must_qty ? 'text-orange' : 'text-green'">
                      {{ item.real_qty }} {{ item.unit }}
                    </div>
                  </div>
                </div>
              </q-card-section>
            </q-card>
          </div>

          <div v-if="currentOrderDetails.length === 0" class="text-center q-pa-lg text-grey">
            暂无明细数据
          </div>
        </div>
      </div>

      <div class="bg-white q-pa-md shadow-up-2">
         <q-btn
           color="primary"
           class="full-width"
           label="完成订单"
           size="lg"
           rounded
           unelevated
           @click="completeOrder"
         />
      </div>
    </div>
  </q-page>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useQuasar } from 'quasar';
import { SaloutstockService } from 'boot/saloutstock';

const $q = useQuasar();

// State
const viewMode = ref('list'); // 'list' | 'detail'
const currentOrder = ref({});
const currentOrderDetails = ref([]);
const orderList = ref([]);
const loading = ref(false);

// 加载订单列表
const loadOrderList = async () => {
  loading.value = true;
  try {
    const response = await SaloutstockService.list({ page: 1, size: 100 });
    // 后端返回分页格式：{ total: number, data: [] }
    const pageData = response.data || {};
    const data = pageData.data || [];
    // 转换数据格式：将后端返回的数据转换为前端需要的格式
    orderList.value = data.map(item => ({
      doc_no: item.BillNo || item.bill_no,
      doc_date: item.Date || item.date,
    }));
  } catch (error) {
    console.error('加载订单列表失败:', error);
    $q.notify({
      message: '加载订单列表失败: ' + (error.response?.data?.error || error.message),
      color: 'negative',
      icon: 'error',
      position: 'top'
    });
  } finally {
    loading.value = false;
  }
};

// 打开订单详情
const openOrder = async (order) => {
  currentOrder.value = order;
  loading.value = true;
  try {
    const response = await SaloutstockService.get(order.doc_no);
    const detail = response.data;

    // 使用新的DTO格式
    if (detail && detail.items) {
      currentOrderDetails.value = detail.items.map((item) => ({
        seq: item.seq,
        material_no: item.material_no,
        product_name: item.material_name,
        spec: item.specification,
        must_qty: item.must_qty,
        real_qty: item.real_qty,
        unit: item.unit,
        batch_no: item.lot_no,
      }));

      // 保存备注信息
      if (detail.note) {
        currentOrder.value.note = detail.note;
      }
    } else {
      currentOrderDetails.value = [];
    }

    viewMode.value = 'detail';
  } catch (error) {
    console.error('加载订单详情失败:', error);
    $q.notify({
      message: '加载订单详情失败: ' + (error.response?.data?.error || error.message),
      color: 'negative',
      icon: 'error',
      position: 'top'
    });
  } finally {
    loading.value = false;
  }
};

// 完成订单
const completeOrder = () => {
  $q.dialog({
    title: '确认',
    message: '确定要完成该订单吗？',
    cancel: true,
    persistent: true
  }).onOk(async () => {
    loading.value = true;
    try {
      await SaloutstockService.completeOrder({
        bill_no: currentOrder.value.doc_no,
      });

      // 从列表中移除
      const index = orderList.value.findIndex(o => o.doc_no === currentOrder.value.doc_no);
      if (index !== -1) {
        orderList.value.splice(index, 1);
      }

      $q.notify({
        message: '订单已完成',
        color: 'positive',
        icon: 'check_circle',
        position: 'top'
      });
      viewMode.value = 'list';
    } catch (error) {
      console.error('完成订单失败:', error);
      $q.notify({
        message: '完成订单失败: ' + (error.response?.data?.error || error.message),
        color: 'negative',
        icon: 'error',
        position: 'top'
      });
    } finally {
      loading.value = false;
    }
  });
};

// 初始化加载
onMounted(() => {
  loadOrderList();
});
</script>

<style scoped>
.order-item-card {
  border-radius: 12px;
  border: 1px solid #f0f0f0;
}
/* Ensure the page takes full height for the flex column layout to work */
.q-page {
  height: calc(100vh - 50px); /* Adjust based on actual header height if needed, or use 100vh */
  overflow: hidden;
}
</style>
