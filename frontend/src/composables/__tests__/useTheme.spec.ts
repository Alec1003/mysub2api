import { beforeEach, describe, expect, it, vi } from 'vitest'

describe('useTheme', () => {
  beforeEach(() => {
    vi.resetModules()
    localStorage.clear()
    document.documentElement.classList.remove('dark')
    vi.stubGlobal('matchMedia', vi.fn(() => ({
      matches: false,
      addEventListener: vi.fn(),
    })))
  })

  it('keeps every consumer synchronized with the document theme', async () => {
    const { useTheme } = await import('../useTheme')
    const first = useTheme()
    const second = useTheme()

    first.toggleTheme()

    expect(first.isDark.value).toBe(true)
    expect(second.isDark.value).toBe(true)
    expect(document.documentElement.classList.contains('dark')).toBe(true)
    expect(localStorage.getItem('theme')).toBe('dark')
  })

  it('uses the system preference only when no saved theme exists', async () => {
    vi.stubGlobal('matchMedia', vi.fn(() => ({
      matches: true,
      addEventListener: vi.fn(),
    })))
    const { useTheme } = await import('../useTheme')

    expect(useTheme().isDark.value).toBe(true)
  })
})
