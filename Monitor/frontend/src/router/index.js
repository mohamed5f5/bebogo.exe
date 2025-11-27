import { createRouter, createWebHistory } from 'vue-router'

import Dashboard from '../pages/Dashboard.vue'
import Devices from '../pages/Devices.vue'
import Scan from '../pages/Scan.vue'
import Diagnostics from '../pages/Diagnostics.vue'
import MapView from '../pages/MapView.vue'
import HistoryView from '../pages/HistoryView.vue'
import SettingsView from '../pages/SettingsView.vue'
import AboutView from '../pages/AboutView.vue'

const routes = [
  { path: '/', redirect: '/dashboard' },
  { path: '/dashboard', component: Dashboard },
  { path: '/devices', component: Devices },
  { path: '/scan', component: Scan },
  { path: '/diagnostics', component: Diagnostics },
  { path: '/map', component: MapView },
  { path: '/history', component: HistoryView },
  { path: '/settings', component: SettingsView },
  { path: '/about', component: AboutView }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router