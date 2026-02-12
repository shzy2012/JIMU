<template>
  <q-page class="q-pa-md bg-grey-1">
    <!-- 工具栏 -->
    <div class="row q-mb-md items-center">
      <div class="col">
        <div class="text-h6 text-weight-bold">标签打印</div>
        <div class="text-caption text-grey-7">TOLOBIO大标签 - 70mm × 30mm</div>
      </div>
      <div class="col-auto q-gutter-x-sm">
        <q-btn
          color="primary"
          label="打印"
          icon="print"
          @click="handlePrint"
        />
        <q-btn
          color="grey-7"
          label="预览"
          icon="preview"
          @click="showPreview = !showPreview"
        />
      </div>
    </div>

    <!-- 标签预览区域 -->
    <div class="row justify-center">
      <!-- 标签预览 -->
      <div class="col-12 col-md-8 col-lg-6">
        <q-card class="q-pa-md">
          <div class="text-subtitle1 q-mb-md">标签预览</div>

          <!-- 标签容器 - 70mm x 30mm 比例 -->
          <div class="label-container">
            <div
              id="label-preview"
              class="label-preview"
              :class="{ 'show-border': showPreview }"
            >
              <!-- 标签内容 -->
              <div class="label-content">
                <!-- 顶部：品牌/标题区域 -->
                <div class="label-header">
                  <div class="label-brand">TOLOBIO</div>
                  <div class="label-subtitle">大标签</div>
                </div>

                <!-- 中间：主要信息区域 -->
                <div class="label-body">
                  <div class="label-row">
                    <span class="label-field">产品：</span>
                    <span class="label-value">{{ labelData.productName }}</span>
                  </div>

                  <div class="label-row">
                    <span class="label-field">批次：</span>
                    <span class="label-value">{{ labelData.batchNo }}</span>
                  </div>

                  <div class="label-row">
                    <span class="label-field">规格：</span>
                    <span class="label-value">{{ labelData.specification }}</span>
                    <span class="label-unit">{{ labelData.unit }}</span>
                  </div>

                  <div class="label-row">
                    <span class="label-field">数量：</span>
                    <span class="label-value">{{ labelData.quantity }}</span>
                  </div>
                </div>

                <!-- 底部：日期信息 -->
                <div class="label-footer">
                  <div class="label-date">
                    生产：{{ formatDate(labelData.productionDate) }}
                  </div>
                  <div class="label-date">
                    有效期：{{ formatDate(labelData.expiryDate) }}
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- 尺寸说明 -->
          <div class="text-caption text-grey-6 q-mt-sm text-center">
            实际尺寸：70mm × 30mm
          </div>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script setup>
import { ref } from "vue";

const showPreview = ref(true);

const labelData = ref({
  productName: "TOLOBIO产品示例",
  batchNo: "BATCH20241218001",
  productionDate: "2024-12-18",
  expiryDate: "2025-12-18",
  quantity: "100",
  specification: "500ml",
  unit: "瓶",
});

const formatDate = (dateStr) => {
  if (!dateStr) return "";
  const date = new Date(dateStr);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const day = String(date.getDate()).padStart(2, "0");
  return `${year}-${month}-${day}`;
};

const handlePrint = () => {
  const printContent = document.getElementById("label-preview");
  const printWindow = window.open("", "_blank");

  printWindow.document.write(`
    <!DOCTYPE html>
    <html>
      <head>
        <title>标签打印</title>
        <style>
          @page {
            size: 70mm 30mm;
            margin: 0;
          }
          body {
            margin: 0;
            padding: 0;
            font-family: Arial, sans-serif;
          }
          ${document.getElementById("label-preview")?.innerHTML || ""}
        </style>
        <style>
          .label-preview {
            width: 70mm;
            height: 30mm;
            padding: 2mm;
            box-sizing: border-box;
            border: 1px solid #000;
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            font-size: 8pt;
          }
          .label-header {
            text-align: center;
            border-bottom: 1px solid #000;
            padding-bottom: 1mm;
            margin-bottom: 1mm;
          }
          .label-brand {
            font-weight: bold;
            font-size: 12pt;
          }
          .label-subtitle {
            font-size: 7pt;
            color: #666;
          }
          .label-body {
            flex: 1;
            display: flex;
            flex-direction: column;
            justify-content: space-around;
          }
          .label-row {
            display: flex;
            align-items: center;
            font-size: 7pt;
          }
          .label-field {
            font-weight: bold;
            margin-right: 2mm;
            min-width: 12mm;
          }
          .label-value {
            flex: 1;
          }
          .label-unit {
            margin-left: 1mm;
            color: #666;
          }
          .label-footer {
            display: flex;
            justify-content: space-between;
            font-size: 6pt;
            border-top: 1px solid #000;
            padding-top: 1mm;
            margin-top: 1mm;
          }
          .label-date {
            color: #666;
          }
        </style>
      </head>
      <body>
        ${printContent?.innerHTML || ""}
      </body>
    </html>
  `);

  printWindow.document.close();
  printWindow.focus();
  setTimeout(() => {
    printWindow.print();
    printWindow.close();
  }, 250);
};
</script>

<style scoped>
.label-container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
  background: #f5f5f5;
  border-radius: 4px;
  min-height: 300px;
}

.label-preview {
  width: 264.583px; /* 70mm at 96dpi */
  height: 113.386px; /* 30mm at 96dpi */
  padding: 7.559px; /* 2mm */
  box-sizing: border-box;
  background: white;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  font-size: 10px;
  transform-origin: top left;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.label-preview.show-border {
  border: 2px dashed #ccc;
}

.label-content {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.label-header {
  text-align: center;
  border-bottom: 1px solid #333;
  padding-bottom: 3px;
  margin-bottom: 3px;
}

.label-brand {
  font-weight: bold;
  font-size: 16px;
  letter-spacing: 1px;
  color: #000;
}

.label-subtitle {
  font-size: 9px;
  color: #666;
  margin-top: 2px;
}

.label-body {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  padding: 2px 0;
}

.label-row {
  display: flex;
  align-items: center;
  font-size: 9px;
  line-height: 1.4;
}

.label-field {
  font-weight: bold;
  margin-right: 4px;
  min-width: 32px;
  color: #333;
}

.label-value {
  flex: 1;
  color: #000;
}

.label-unit {
  margin-left: 2px;
  color: #666;
  font-size: 8px;
}

.label-footer {
  display: flex;
  justify-content: space-between;
  font-size: 8px;
  border-top: 1px solid #333;
  padding-top: 3px;
  margin-top: 3px;
}

.label-date {
  color: #666;
}

/* 打印样式 */
@media print {
  body {
    margin: 0;
    padding: 0;
  }

  .label-preview {
    width: 70mm;
    height: 30mm;
    padding: 2mm;
    border: none;
    box-shadow: none;
    page-break-inside: avoid;
  }

  .q-page {
    padding: 0;
  }
}
</style>
