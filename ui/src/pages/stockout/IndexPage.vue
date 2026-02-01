<template>
  <q-page class="bg-grey-2">
    <!-- View 1: Order List -->
    <div v-if="viewMode === 'list'" class="column full-height">
      <div class="bg-primary text-white q-pa-md shadow-2 z-top">
        <div class="text-h6">出库单列表</div>
        <div class="text-subtitle2 text-blue-2">Stockout Orders</div>
      </div>

      <div class="col scroll q-pa-sm">
        <q-list separator class="bg-white rounded-borders shadow-1">
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
              <div class="row items-center">
                <q-badge :color="order.status === '已完成' ? 'green' : 'orange'" rounded class="q-mr-sm">
                  {{ order.status || '待处理' }}
                </q-badge>
                <q-icon name="chevron_right" color="grey-5" />
              </div>
            </q-item-section>
          </q-item>
        </q-list>

        <div v-if="orderList.length === 0" class="text-center q-pa-lg text-grey">
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
        <div class="text-caption q-mb-sm text-grey-8 q-pl-xs">订单详情 ({{ currentOrderDetails.length }})</div>

        <div class="q-gutter-y-sm">
          <q-card v-for="(item, index) in currentOrderDetails" :key="index" flat bordered class="order-item-card">
            <q-card-section class="q-pb-xs">
              <div class="row no-wrap justify-between items-start">
                <div class="col">
                  <div class="text-body1 text-weight-bold">{{ item.product_name }}</div>
                  <div class="text-caption text-grey-7">{{ item.spec }}</div>
                </div>
                <div class="col-auto text-right">
                   <q-badge outline color="primary" :label="item.sku" />
                </div>
              </div>
            </q-card-section>

            <q-separator inset spaced />

            <q-card-section class="q-pt-xs">
              <div class="row q-col-gutter-sm">
                <div class="col-6">
                  <div class="text-caption text-grey">批号</div>
                  <div class="text-body2">{{ item.batch_no }}</div>
                </div>
                 <div class="col-6 text-right">
                  <div class="text-caption text-grey">配给 / 订单</div>
                  <div class="text-body1 text-weight-bold">
                    <span :class="item.allocated_qty < item.order_qty ? 'text-orange' : 'text-green'">
                      {{ item.allocated_qty }}
                    </span>
                    <span class="text-grey-5"> / {{ item.order_qty }}</span>
                  </div>
                </div>
              </div>
            </q-card-section>
          </q-card>
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
import { ref } from 'vue';
import { useQuasar } from 'quasar';

const $q = useQuasar();

// State
const viewMode = ref('list'); // 'list' | 'detail'
const currentOrder = ref({});
const currentOrderDetails = ref([]);

// Mock Data
const orderList = ref([
  { doc_no: 'CK-20231027-001', doc_date: '2023-10-27', status: '待处理' },
  { doc_no: 'CK-20231027-002', doc_date: '2023-10-27', status: '进行中' },
]);

// Details Mock Data Generator
const getMockDetails = (orderId) => {
  return [
    {
      sku: 'P001023',
      product_name: '高性能锂电池组',
      spec: '48V 20Ah / 黑色',
      batch_no: 'B20231001-A',
      order_qty: 10,
      allocated_qty: 10
    },
    {
      sku: 'A883211',
      product_name: '智能控制器V3',
      spec: '通用型',
      batch_no: 'B20230915-C',
      order_qty: 50,
      allocated_qty: 48 // Example of shortage
    },
    {
      sku: 'C772122',
      product_name: '连接线束',
      spec: '50cm / 防水',
      batch_no: 'B20231010-F',
      order_qty: 100,
      allocated_qty: 100
    },
  ];
};

// Actions
const openOrder = (order) => {
  currentOrder.value = order;
  // Simulate fetching details
  currentOrderDetails.value = getMockDetails(order.doc_no);
  viewMode.value = 'detail';
};

const completeOrder = () => {
  $q.dialog({
    title: '确认',
    message: '确定要完成该订单吗？完成后订单将从列表中移除。',
    cancel: true,
    persistent: true
  }).onOk(() => {
    // Determine the index of the current order
    const index = orderList.value.findIndex(o => o.doc_no === currentOrder.value.doc_no);
    if (index !== -1) {
      orderList.value.splice(index, 1);
      $q.notify({
        message: '订单已完成',
        color: 'positive',
        icon: 'check_circle',
        position: 'top'
      });
      viewMode.value = 'list';
    }
  });
};
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
