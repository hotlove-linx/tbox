import { createRouter, createWebHashHistory, type RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/auth/LoginRegister.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/',
    component: () => import('@/components/layout/AppLayout.vue'),
    redirect: '/home',
    children: [
      {
        path: 'home',
        name: 'Home',
        component: () => import('@/views/home/HomePage.vue'),
        meta: { title: '首页' }
      },
      {
        path: 'category',
        name: 'Category',
        component: () => import('@/views/category/CategoryBrowse.vue'),
        meta: { title: '分类浏览' }
      },
      {
        path: 'software/:id',
        name: 'SoftwareDetail',
        component: () => import('@/views/software/SoftwareDetail.vue'),
        meta: { title: '软件详情' }
      },
      {
        path: 'downloads',
        name: 'Downloads',
        component: () => import('@/views/download/DownloadManager.vue'),
        meta: { title: '下载管理' }
      },
      {
        path: 'install-records',
        name: 'InstallRecords',
        component: () => import('@/views/install/InstallRecord.vue'),
        meta: { title: '安装记录' }
      },
      {
        path: 'space',
        name: 'PersonalSpace',
        component: () => import('@/views/space/PersonalSpace.vue'),
        meta: { title: '个人空间' }
      },
      {
        path: 'profile',
        name: 'Profile',
        component: () => import('@/views/profile/UserProfile.vue'),
        meta: { title: '个人中心' }
      },
      {
        path: 'settings',
        name: 'Settings',
        component: () => import('@/views/settings/SettingsPage.vue'),
        meta: { title: '设置' }
      }
    ]
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

// Navigation guard
router.beforeEach((to, _from, next) => {
  const token = localStorage.getItem('tbox_token')
  if (to.meta.requiresAuth !== false && !token) {
    next('/login')
  } else {
    next()
  }
})

export default router
