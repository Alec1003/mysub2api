# Brutal 前端重构审查

审查日期：2026-09-02

## 结论

当前不是全站 Brutal 重构完成，而是“认证页 + 首页 + 用户 Dashboard 外壳”基本完成。管理员 Dashboard、共享布局内部控件，以及用户 Dashboard 子组件内部仍保留旧设计系统，属于混合状态。

## 已验证

- 登录页：运行时可见粗黑边框、偏移硬阴影、蓝色主按钮和网格背景。
- 注册页：运行时与登录页视觉统一；当前业务状态显示注册功能暂时关闭。
- 首页：运行时可见蓝色 Hero、粗边框、偏移阴影，以及模型、价格和 FAQ 区块。
- Docker：`sub2api`、PostgreSQL、Redis 均在运行，`/health` 返回 HTTP 200。

## 未完成或混合部分

- `frontend/src/views/admin/DashboardView.vue` 未纳入 Brutal 视觉迁移，仍大量使用 `rounded-lg`、`bg-purple-*`、`text-purple-*`、灰色卡片和渐进式 hover 样式。
- `frontend/src/components/user/dashboard/` 的四个子组件虽然添加了 Brutal marker，但内部仍有 `rounded-xl`、`rounded-full`、`bg-gray-*`、`bg-primary-*` 等旧 utility。
- `frontend/src/components/layout/AppHeader.vue` 只完成外层覆盖，内部仍有圆角按钮、圆形徽章、灰色下拉菜单等旧样式。
- `frontend/src/components/layout/AppSidebar.vue` 需要继续检查并统一内部导航项、状态和图标容器。
- `frontend/src/style.css` 仍保留旧全局 `.btn`、`.input` 和渐变定义；当前主要依赖 `.brutal-auth` / `.brutal-shell` 的后置覆盖，因此新旧设计系统会同时生效。

## 测试限制

`frontend/src/styles/__tests__/brutalVisualContract.spec.ts` 目前只检查 marker、变量和少量源码关键词，没有覆盖管理员 Dashboard，也没有验证运行时计算样式或截图。因此测试通过不能证明全站视觉已经统一。

本次测试尝试被 pnpm 的 ignored build scripts（`esbuild` / `vue-demi`）阻断，未得到 Vitest 最终结果。

## 建议顺序

1. 先迁移管理员 Dashboard，补齐对应 Brutal marker 和运行时截图测试。
2. 再统一 Header、Sidebar 和四个用户 Dashboard 子组件内部的圆角、颜色和阴影。
3. 最后收敛 `style.css` 的全局按钮、输入框和渐变规则，避免依赖页面级覆盖。
