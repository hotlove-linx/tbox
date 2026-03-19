"use strict";
const electron = require("electron");
const path = require("path");
let mainWindow = null;
let tray = null;
const VITE_DEV_SERVER_URL = process.env.VITE_DEV_SERVER_URL;
function createWindow() {
  mainWindow = new electron.BrowserWindow({
    width: 1440,
    height: 900,
    minWidth: 1280,
    minHeight: 720,
    frame: false,
    webPreferences: {
      preload: path.join(__dirname, "preload.js"),
      contextIsolation: true,
      nodeIntegration: false
    },
    show: false
  });
  mainWindow.once("ready-to-show", () => {
    mainWindow == null ? void 0 : mainWindow.show();
  });
  if (VITE_DEV_SERVER_URL) {
    mainWindow.loadURL(VITE_DEV_SERVER_URL);
    mainWindow.webContents.openDevTools();
  } else {
    mainWindow.loadFile(path.join(__dirname, "../dist/index.html"));
  }
  mainWindow.on("close", (event) => {
    event.preventDefault();
    mainWindow == null ? void 0 : mainWindow.hide();
  });
  mainWindow.on("closed", () => {
    mainWindow = null;
  });
}
function createTray() {
  const icon = electron.nativeImage.createEmpty();
  tray = new electron.Tray(icon);
  const contextMenu = electron.Menu.buildFromTemplate([
    { label: "打开 TBox", click: () => mainWindow == null ? void 0 : mainWindow.show() },
    { type: "separator" },
    { label: "设置", click: () => {
      mainWindow == null ? void 0 : mainWindow.show();
      mainWindow == null ? void 0 : mainWindow.webContents.send("navigate", "/settings");
    } },
    { type: "separator" },
    { label: "退出", click: () => {
      mainWindow == null ? void 0 : mainWindow.destroy();
      electron.app.quit();
    } }
  ]);
  tray.setToolTip("TBox - 软件商城");
  tray.setContextMenu(contextMenu);
  tray.on("double-click", () => mainWindow == null ? void 0 : mainWindow.show());
}
electron.app.whenReady().then(() => {
  createWindow();
  createTray();
});
electron.app.on("window-all-closed", () => {
  if (process.platform !== "darwin") {
    electron.app.quit();
  }
});
electron.app.on("activate", () => {
  if (electron.BrowserWindow.getAllWindows().length === 0) {
    createWindow();
  }
});
electron.ipcMain.on("window-minimize", () => {
  mainWindow == null ? void 0 : mainWindow.minimize();
});
electron.ipcMain.on("window-maximize", () => {
  if (mainWindow == null ? void 0 : mainWindow.isMaximized()) {
    mainWindow.unmaximize();
  } else {
    mainWindow == null ? void 0 : mainWindow.maximize();
  }
});
electron.ipcMain.on("window-close", () => {
  mainWindow == null ? void 0 : mainWindow.hide();
});
electron.ipcMain.on("window-force-close", () => {
  mainWindow == null ? void 0 : mainWindow.destroy();
  electron.app.quit();
});
electron.ipcMain.handle("window-is-maximized", () => {
  return (mainWindow == null ? void 0 : mainWindow.isMaximized()) ?? false;
});
electron.ipcMain.handle("open-external", async (_event, url) => {
  await electron.shell.openExternal(url);
});
electron.ipcMain.handle("select-directory", async () => {
  const { dialog } = require("electron");
  const result = await dialog.showOpenDialog(mainWindow, {
    properties: ["openDirectory"]
  });
  return result.canceled ? null : result.filePaths[0];
});
