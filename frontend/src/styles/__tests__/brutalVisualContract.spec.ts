import { describe, expect, it } from 'vitest'
import { readFileSync } from 'node:fs'
import { resolve } from 'node:path'

const source = (path: string) => readFileSync(resolve(process.cwd(), path), 'utf8')

const authFiles = [
  'src/components/layout/AuthLayout.vue',
  'src/views/auth/LoginView.vue',
  'src/views/auth/RegisterView.vue',
]

const dashboardFiles = [
  'src/components/layout/AppLayout.vue',
  'src/components/layout/AppHeader.vue',
  'src/components/layout/AppSidebar.vue',
  'src/views/user/DashboardView.vue',
  'src/components/user/dashboard/UserDashboardStats.vue',
  'src/components/user/dashboard/UserDashboardCharts.vue',
  'src/components/user/dashboard/UserDashboardRecentUsage.vue',
  'src/components/user/dashboard/UserDashboardQuickActions.vue',
]

describe('MyWorkBuddy brutal visual contract', () => {
  it('marks every authentication surface as part of the brutal design system', () => {
    expect(source(authFiles[0])).toContain('data-ui="brutal-auth-shell"')
    expect(source(authFiles[1])).toContain('data-ui="brutal-login"')
    expect(source(authFiles[2])).toContain('data-ui="brutal-register"')
  })

  it('marks the dashboard shell and every dashboard module as brutal surfaces', () => {
    const requiredMarkers = [
      'data-ui="brutal-app-shell"',
      'data-ui="brutal-header"',
      'data-ui="brutal-sidebar"',
      'data-ui="brutal-dashboard"',
      'data-ui="brutal-dashboard-stats"',
      'data-ui="brutal-dashboard-charts"',
      'data-ui="brutal-dashboard-recent"',
      'data-ui="brutal-dashboard-actions"',
    ]

    dashboardFiles.forEach((file, index) => {
      expect(source(file), `${file} should expose its brutal UI marker`).toContain(requiredMarkers[index])
    })
  })

  it('defines hard borders, offset shadows, square controls and a blue accent centrally', () => {
    const styles = source('src/style.css')
    expect(styles).toContain('--brutal-ink:')
    expect(styles).toContain('--brutal-blue:')
    expect(styles).toContain('--brutal-shadow:')
    expect(styles).toContain('.brutal-surface')
    expect(styles).toContain('.brutal-control')
  })

  it('does not retain purple, violet, gradient or glow styling in redesigned surfaces', () => {
    const forbidden = /(?:purple|violet|gradient|shadow-glow)/i
    for (const file of [...authFiles, ...dashboardFiles]) {
      expect(source(file), `${file} still contains a forbidden legacy visual token`).not.toMatch(forbidden)
    }
  })
})
