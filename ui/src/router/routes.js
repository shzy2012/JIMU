const routes = [
  {
    path: "/",
    component: () => import("layouts/SimpleLayout.vue"),
    children: [{ path: "", component: () => import("pages/IndexPage.vue") }],
  },
  {
    path: "/example",
    component: () => import("layouts/SimpleLayout.vue"),
    children: [
      { path: "", component: () => import("pages/example/IndexPage.vue") },
    ],
  },
  {
    path: "/stockout",
    component: () => import("layouts/EmptyLayout.vue"),
    children: [
      { path: "", component: () => import("pages/stockout/IndexPage.vue") },
    ],
  },
  {
    path: "/login",
    component: () => import("layouts/EmptyLayout.vue"),
    children: [
      { path: "", component: () => import("src/pages/LoginPage.vue") },
    ],
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: "/:catchAll(.*)*",
    component: () => import("src/pages/ErrorNotFound.vue"),
  },
];

export default routes;
