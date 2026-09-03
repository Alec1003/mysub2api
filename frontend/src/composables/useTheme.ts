import { ref } from 'vue'

export type Theme = 'light' | 'dark'

const THEME_KEY = 'theme'
const isDark = ref(false)
let initialized = false
let mediaQuery: MediaQueryList | null = null

function applyTheme(theme: Theme) {
  isDark.value = theme === 'dark'
  document.documentElement.classList.toggle('dark', isDark.value)
}

function initTheme() {
  if (typeof window === 'undefined' || typeof document === 'undefined') return
  const saved = localStorage.getItem(THEME_KEY)
  const systemDark = window.matchMedia('(prefers-color-scheme: dark)').matches
  applyTheme(saved === 'dark' || (!saved && systemDark) ? 'dark' : 'light')
  if (initialized) return
  initialized = true
  mediaQuery = window.matchMedia('(prefers-color-scheme: dark)')
  mediaQuery.addEventListener?.('change', (event) => {
    if (!localStorage.getItem(THEME_KEY)) applyTheme(event.matches ? 'dark' : 'light')
  })
  window.addEventListener('storage', (event) => {
    if (event.key === THEME_KEY) applyTheme(event.newValue === 'dark' ? 'dark' : 'light')
  })
}

function setTheme(theme: Theme) {
  initTheme()
  applyTheme(theme)
  localStorage.setItem(THEME_KEY, theme)
}

function toggleTheme() {
  setTheme(isDark.value ? 'light' : 'dark')
}

export function useTheme() {
  initTheme()
  return { isDark, setTheme, toggleTheme }
}

export { initTheme }
