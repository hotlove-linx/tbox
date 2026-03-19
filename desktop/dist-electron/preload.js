"use strict";
const electron = require("electron");
electron.contextBridge.exposeInMainWorld("electronAPI", {
  minimize: () => electron.ipcRenderer.send("window-minimize"),
  maximize: () => electron.ipcRenderer.send("window-maximize"),
  close: () => electron.ipcRenderer.send("window-close"),
  forceClose: () => electron.ipcRenderer.send("window-force-close"),
  isMaximized: () => electron.ipcRenderer.invoke("window-is-maximized"),
  openExternal: (url) => electron.ipcRenderer.invoke("open-external", url),
  selectDirectory: () => electron.ipcRenderer.invoke("select-directory"),
  onNavigate: (callback) => {
    electron.ipcRenderer.on("navigate", (_event, route) => callback(route));
  }
});
