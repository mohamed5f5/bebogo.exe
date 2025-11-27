<template>
  <div class="layout">
    <aside class="sidebar">
      <div class="logo">
        <div class="logo-mark">JMS</div>
        <div class="logo-text">
          <div class="title">Smart Monitor</div>
          <div class="subtitle">Network Intelligence</div>
        </div>
      </div>
      <nav class="nav">
        <RouterLink to="/dashboard">Dashboard</RouterLink>
        <RouterLink to="/devices">Devices</RouterLink>
        <RouterLink to="/scan">Scan</RouterLink>
        <RouterLink to="/diagnostics">Diagnostics</RouterLink>
        <RouterLink to="/map">Map</RouterLink>
        <RouterLink to="/history">History</RouterLink>
        <RouterLink to="/settings">Settings</RouterLink>
        <RouterLink to="/about">About</RouterLink>
      </nav>
    </aside>
    <main class="content">
      <header class="topbar">
        <div class="topbar-title">JMS Smart Network Monitor</div>
      </header>
      <section class="page">
        <slot />
      </section>
    </main>
  </div>
</template>

<script setup>
import { RouterLink } from 'vue-router'
</script>

<style scoped>
.layout {
  display: flex;
  min-height: 100vh;
}
.sidebar {
  width: 240px;
  background: #020617;
  padding: 1rem;
  border-right: 1px solid #1f2937;
  display: flex;
  flex-direction: column;
}
.logo {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 2rem;
}
.logo-mark {
  width: 36px;
  height: 36px;
  border-radius: 0.75rem;
  border: 1px solid #22c55e;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 0.9rem;
}
.logo-text .title {
  font-weight: 600;
  font-size: 0.9rem;
}
.logo-text .subtitle {
  font-size: 0.75rem;
  color: #9ca3af;
}
.nav {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}
.nav a {
  text-decoration: none;
  color: #9ca3af;
  font-size: 0.9rem;
  padding: 0.35rem 0.5rem;
  border-radius: 0.5rem;
}
.nav a.router-link-active {
  background: #0f172a;
  color: #e5e7eb;
}
.content {
  flex: 1;
  display: flex;
  flex-direction: column;
}
.topbar {
  height: 52px;
  border-bottom: 1px solid #1f2937;
  display: flex;
  align-items: center;
  padding: 0 1.5rem;
}
.topbar-title {
  font-size: 0.95rem;
  font-weight: 500;
}
.page {
  padding: 1.5rem;
}
</style>