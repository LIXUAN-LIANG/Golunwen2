import { createRouter, createWebHistory } from 'vue-router'

const routes=[
    {
      path:"/",
      name:"login",
      component:() => import("../pages/login.vue"),
    },
    {
        path:"/index",
        name:"index",
        component:() => import("../pages/index.vue"),
    },
    {
        path:"/detail",
        name:"detail",
        component:() => import("../pages/detail.vue"),
    },
    {
        path:"/test",
        name:"test",
        component:() => import("../pages/test.vue"),
    }
]

export const router=createRouter({
history:createWebHistory(),
routes,
})
