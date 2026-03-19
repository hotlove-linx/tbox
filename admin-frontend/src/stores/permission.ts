import { defineStore } from 'pinia'
import { ref } from 'vue'

export interface MenuItem {
  path: string
  title: string
  icon: string
  children?: MenuItem[]
  permission?: string
}

const allMenus: MenuItem[] = [
  {
    path: '/dashboard',
    title: '仪表盘',
    icon: 'Odometer',
    permission: 'dashboard:view'
  },
  {
    path: '/software',
    title: '软件管理',
    icon: 'Box',
    children: [
      { path: '/software/list', title: '软件列表', icon: 'List', permission: 'software:view' },
      { path: '/software/categories', title: '分类管理', icon: 'Folder', permission: 'category:view' },
      { path: '/software/tags', title: '标签管理', icon: 'PriceTag', permission: 'tag:view' }
    ]
  },
  {
    path: '/content',
    title: '内容运营',
    icon: 'Promotion',
    children: [
      { path: '/content/banners', title: 'Banner管理', icon: 'Picture', permission: 'banner:view' },
      { path: '/content/recommendations', title: '推荐管理', icon: 'Star', permission: 'recommend:view' },
      { path: '/content/announcements', title: '公告管理', icon: 'Bell', permission: 'announcement:view' }
    ]
  },
  {
    path: '/audit',
    title: '审核中心',
    icon: 'Checked',
    children: [
      { path: '/audit/software', title: '软件审核', icon: 'Document', permission: 'audit:software' },
      { path: '/audit/reviews', title: '评论审核', icon: 'ChatDotRound', permission: 'audit:review' },
      { path: '/audit/reports', title: '举报处理', icon: 'Warning', permission: 'audit:report' }
    ]
  },
  {
    path: '/user',
    title: '用户管理',
    icon: 'User',
    children: [
      { path: '/user/list', title: '用户列表', icon: 'UserFilled', permission: 'user:view' },
      { path: '/user/feedback', title: '用户反馈', icon: 'ChatLineSquare', permission: 'feedback:view' }
    ]
  },
  {
    path: '/statistics',
    title: '数据统计',
    icon: 'DataAnalysis',
    children: [
      { path: '/statistics/downloads', title: '下载统计', icon: 'Download', permission: 'stats:download' },
      { path: '/statistics/users', title: '用户统计', icon: 'TrendCharts', permission: 'stats:user' },
      { path: '/statistics/ranking', title: '软件排行', icon: 'Trophy', permission: 'stats:ranking' }
    ]
  },
  {
    path: '/system',
    title: '系统管理',
    icon: 'Setting',
    children: [
      { path: '/system/admins', title: '管理员管理', icon: 'Avatar', permission: 'admin:view' },
      { path: '/system/roles', title: '角色权限', icon: 'Lock', permission: 'role:view' },
      { path: '/system/logs', title: '操作日志', icon: 'Notebook', permission: 'log:view' },
      { path: '/system/config', title: '系统配置', icon: 'Tools', permission: 'config:view' }
    ]
  }
]

export const usePermissionStore = defineStore('permission', () => {
  const menus = ref<MenuItem[]>([])

  function generateMenus(permissions: string[]) {
    if (permissions.includes('*')) {
      menus.value = allMenus
      return
    }
    menus.value = allMenus
      .map((menu) => {
        if (menu.children) {
          const children = menu.children.filter(
            (child) => !child.permission || permissions.includes(child.permission)
          )
          if (children.length > 0) {
            return { ...menu, children }
          }
          return null
        }
        if (!menu.permission || permissions.includes(menu.permission)) {
          return menu
        }
        return null
      })
      .filter(Boolean) as MenuItem[]
  }

  return { menus, generateMenus }
})
