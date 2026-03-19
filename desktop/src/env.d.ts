/// <reference types="vite/client" />

declare module '*.vue' {
  import type { DefineComponent } from 'vue'
  const component: DefineComponent<{}, {}, any>
  export default component
}

interface ElectronAPI {
  minimize: () => void
  maximize: () => void
  close: () => void
  forceClose: () => void
  isMaximized: () => Promise<boolean>
  openExternal: (url: string) => Promise<void>
  selectDirectory: () => Promise<string | null>
  onNavigate: (callback: (route: string) => void) => void
}

interface Window {
  electronAPI?: ElectronAPI
}
