import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { usePermissionStore } from '@/stores/permission'

const routes: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/login/LoginPage.vue'),
    meta: { title: '登录', requiresAuth: false }
  },
  {
    path: '/',
    component: () => import('@/components/layout/AdminLayout.vue'),
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/views/dashboard/DashboardPage.vue'),
        meta: { title: '仪表盘', icon: 'Odometer' }
      },
      {
        path: 'software/list',
        name: 'SoftwareList',
        component: () => import('@/views/software/SoftwareList.vue'),
        meta: { title: '软件列表', parent: '软件管理' }
      },
      {
        path: 'software/detail/:id',
        name: 'SoftwareDetail',
        component: () => import('@/views/software/SoftwareDetail.vue'),
        meta: { title: '软件详情', parent: '软件管理', hidden: true }
      },
      {
        path: 'software/categories',
        name: 'CategoryManage',
        component: () => import('@/views/software/CategoryManage.vue'),
        meta: { title: '分类管理', parent: '软件管理' }
      },
      {
        path: 'software/tags',
        name: 'TagManage',
        component: () => import('@/views/software/TagManage.vue'),
        meta: { title: '标签管理', parent: '软件管理' }
      },
      {
        path: 'content/banners',
        name: 'BannerManage',
        component: () => import('@/views/content/BannerManage.vue'),
        meta: { title: 'Banner管理', parent: '内容运营' }
      },
      {
        path: 'content/recommendations',
        name: 'RecommendManage',
        component: () => import('@/views/content/RecommendManage.vue'),
        meta: { title: '推荐管理', parent: '内容运营' }
      },
      {
        path: 'content/announcements',
        name: 'AnnouncementManage',
        component: () => import('@/views/content/AnnouncementManage.vue'),
        meta: { title: '公告管理', parent: '内容运营' }
      },
      {
        path: 'audit/software',
        name: 'SoftwareAudit',
        component: () => import('@/views/audit/SoftwareAudit.vue'),
        meta: { title: '软件审核', parent: '审核中心' }
      },
      {
        path: 'audit/reviews',
        name: 'ReviewAudit',
        component: () => import('@/views/audit/ReviewAudit.vue'),
        meta: { title: '评论审核', parent: '审核中心' }
      },
      {
        path: 'audit/reports',
        name: 'ReportHandle',
        component: () => import('@/views/audit/ReportHandle.vue'),
        meta: { title: '举报处理', parent: '审核中心' }
      },
      {
        path: 'user/list',
        name: 'UserList',
        component: () => import('@/views/user/UserList.vue'),
        meta: { title: '用户列表', parent: '用户管理' }
      },
      {
        path: 'user/detail/:id',
        name: 'UserDetail',
        component: () => import('@/views/user/UserDetail.vue'),
        meta: { title: '用户详情', parent: '用户管理', hidden: true }
      },
      {
        path: 'user/feedback',
        name: 'FeedbackManage',
        component: () => import('@/views/user/FeedbackManage.vue'),
        meta: { title: '用户反馈', parent: '用户管理' }
      },
      {
        path: 'statistics/downloads',
        name: 'DownloadStats',
        component: () => import('@/views/statistics/DownloadStats.vue'),
        meta: { title: '下载统计', parent: '数据统计' }
      },
      {
        path: 'statistics/users',
        name: 'UserStats',
        component: () => import('@/views/statistics/UserStats.vue'),
        meta: { title: '用户统计', parent: '数据统计' }
      },
      {
        path: 'statistics/ranking',
        name: 'SoftwareRanking',
        component: () => import('@/views/statistics/SoftwareRanking.vue'),
        meta: { title: '软件排行', parent: '数据统计' }
      },
      {
        path: 'system/admins',
        name: 'AdminManage',
        component: () => import('@/views/system/AdminManage.vue'),
        meta: { title: '管理员管理', parent: '系统管理' }
      },
      {
        path: 'system/roles',
        name: 'RolePermission',
        component: () => import('@/views/system/RolePermission.vue'),
        meta: { title: '角色权限', parent: '系统管理' }
      },
      {
        path: 'system/logs',
        name: 'OperationLog',
        component: () => import('@/views/system/OperationLog.vue'),
        meta: { title: '操作日志', parent: '系统管理' }
      },
      {
        path: 'system/config',
        name: 'SystemConfig',
        component: () => import('@/views/system/SystemConfig.vue'),
        meta: { title: '系统配置', parent: '系统管理' }
      }
    ]
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/dashboard'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach(async (to, _from, next) => {
  document.title = `${to.meta.title || 'TBox'} - TBox Admin`

  if (to.meta.requiresAuth === false) {
    next()
    return
  }

  const token = localStorage.getItem('admin_token')
  if (!token) {
    next('/login')
    return
  }

  const authStore = useAuthStore()
  if (!authStore.adminInfo) {
    await authStore.fetchAdminInfo()
    const permissionStore = usePermissionStore()
    permissionStore.generateMenus(authStore.permissions)
  }

  next()
})

export default router
