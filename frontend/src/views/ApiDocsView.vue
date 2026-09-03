<template>
  <div class="min-h-screen bg-slate-50 text-slate-950 dark:bg-slate-950 dark:text-white">
    <header class="sticky top-0 z-40 border-b-4 border-slate-950 bg-white/95 backdrop-blur dark:border-white dark:bg-slate-950/95">
      <nav class="mx-auto flex max-w-7xl items-center justify-between gap-4 px-4 py-3 sm:px-8">
        <router-link to="/home" class="flex items-center gap-3" aria-label="Home">
          <span class="flex h-9 w-9 items-center justify-center overflow-hidden border-2 border-slate-950 bg-blue-600 text-lg font-black text-white dark:border-white">
            <img v-if="siteLogo" :src="siteLogo" :alt="siteName" class="h-full w-full object-contain" />
            <span v-else>{{ siteName.slice(0, 1).toUpperCase() }}</span>
          </span>
          <span class="font-black tracking-tight">{{ siteName.toUpperCase() }}<span class="text-blue-600 dark:text-blue-400">.</span></span>
        </router-link>
        <div class="flex items-center gap-2">
          <LocaleSwitcher />
          <button class="border-2 border-slate-950 p-2 text-slate-950 transition hover:bg-blue-100 dark:border-white dark:text-white dark:hover:bg-slate-800" :title="isDark ? 'Light theme' : 'Dark theme'" @click="toggleTheme">
            <Icon :name="isDark ? 'sun' : 'moon'" size="sm" />
          </button>
          <router-link :to="isAuthenticated ? '/dashboard' : '/login'" class="hidden border-2 border-slate-950 bg-blue-600 px-3 py-2 text-sm font-black text-white shadow-[3px_3px_0_0_#0f172a] sm:inline-flex dark:border-white dark:shadow-[3px_3px_0_0_#2563eb]">
            {{ isAuthenticated ? t('apiDocs.dashboard') : t('apiDocs.login') }}
          </router-link>
        </div>
      </nav>
    </header>

    <main class="mx-auto grid max-w-7xl gap-10 px-4 py-10 sm:px-8 lg:grid-cols-[190px_minmax(0,1fr)] lg:py-14">
      <aside class="hidden lg:block">
        <div class="sticky top-24 border-l-4 border-blue-600 pl-4">
          <p class="mb-4 text-xs font-black uppercase tracking-widest text-slate-500 dark:text-slate-400">{{ t('apiDocs.onThisPage') }}</p>
          <nav class="space-y-3 text-sm font-bold text-slate-600 dark:text-slate-300">
            <a href="#quick-start" class="block hover:text-blue-600 dark:hover:text-blue-400">{{ t('apiDocs.quickStart') }}</a>
            <a href="#workbuddy" class="block hover:text-blue-600 dark:hover:text-blue-400">{{ t('apiDocs.workbuddyTitle') }}</a>
            <a href="#examples" class="block hover:text-blue-600 dark:hover:text-blue-400">{{ t('apiDocs.examplesTitle') }}</a>
            <a href="#faq" class="block hover:text-blue-600 dark:hover:text-blue-400">{{ t('apiDocs.commonIssues') }}</a>
          </nav>
        </div>
      </aside>

      <div class="min-w-0">
        <section class="border-b-4 border-slate-950 pb-10 dark:border-white" aria-labelledby="docs-title">
          <p class="mb-4 inline-flex border-2 border-slate-950 bg-blue-200 px-3 py-1 text-xs font-black uppercase tracking-widest shadow-[3px_3px_0_0_#0f172a] dark:border-white dark:bg-blue-600 dark:text-white dark:shadow-[3px_3px_0_0_#2563eb]">{{ t('apiDocs.apiReference') }}</p>
          <h1 id="docs-title" class="max-w-3xl text-4xl font-black leading-tight sm:text-6xl">{{ t('apiDocs.introTitle') }}</h1>
          <p class="mt-5 max-w-2xl text-lg font-semibold leading-8 text-slate-600 dark:text-slate-300">{{ t('apiDocs.introDescription') }}</p>
          <div class="mt-8 flex flex-wrap gap-3">
            <a href="#quick-start" class="inline-flex items-center gap-2 border-2 border-slate-950 bg-blue-600 px-4 py-3 font-black text-white shadow-[5px_5px_0_0_#0f172a] transition hover:translate-x-1 hover:translate-y-1 hover:shadow-none dark:border-white dark:shadow-[5px_5px_0_0_#2563eb]">{{ t('apiDocs.getStarted') }} <Icon name="arrowRight" size="sm" /></a>
            <router-link to="/model-plaza" class="inline-flex items-center gap-2 border-2 border-slate-950 bg-white px-4 py-3 font-black dark:border-white dark:bg-slate-900">{{ t('apiDocs.viewModels') }}</router-link>
          </div>
        </section>

        <section id="quick-start" class="scroll-mt-24 pt-10">
          <div class="mb-6 flex items-end justify-between gap-4">
            <div><p class="text-sm font-black uppercase tracking-widest text-blue-600 dark:text-blue-400">01 / 03</p><h2 class="mt-2 text-3xl font-black">{{ t('apiDocs.quickStart') }}</h2></div>
            <router-link to="/keys" class="hidden border-b-2 border-blue-600 pb-1 text-sm font-black text-blue-600 sm:inline-flex dark:text-blue-400">{{ t('apiDocs.stepOneAction') }} <Icon name="arrowRight" size="sm" class="ml-1" /></router-link>
          </div>
          <div class="grid gap-4 md:grid-cols-3">
            <article class="border-4 border-slate-950 bg-blue-200 p-5 shadow-[5px_5px_0_0_#0f172a] dark:border-white dark:bg-blue-950 dark:shadow-[5px_5px_0_0_#2563eb]"><span class="text-xs font-black">{{ t('apiDocs.stepLabel', { number: 1 }) }}</span><h3 class="mt-8 text-xl font-black">{{ t('apiDocs.stepOneTitle') }}</h3><p class="mt-3 text-sm font-semibold leading-6">{{ t('apiDocs.stepOneDescription') }}</p></article>
            <article class="border-4 border-slate-950 bg-white p-5 shadow-[5px_5px_0_0_#2563eb] dark:border-white dark:bg-slate-900"><span class="text-xs font-black text-blue-600 dark:text-blue-400">{{ t('apiDocs.stepLabel', { number: 2 }) }}</span><h3 class="mt-8 text-xl font-black">{{ t('apiDocs.stepTwoTitle') }}</h3><p class="mt-3 text-sm font-semibold leading-6 text-slate-600 dark:text-slate-300">{{ t('apiDocs.stepTwoDescription') }}</p></article>
            <article class="border-4 border-slate-950 bg-slate-950 p-5 text-white shadow-[5px_5px_0_0_#2563eb] dark:border-white"><span class="text-xs font-black text-blue-400">{{ t('apiDocs.stepLabel', { number: 3 }) }}</span><h3 class="mt-8 text-xl font-black">{{ t('apiDocs.stepThreeTitle') }}</h3><p class="mt-3 text-sm font-semibold leading-6 text-slate-300">{{ t('apiDocs.stepThreeDescription') }}</p></article>
          </div>
          <div class="mt-6 grid gap-3 sm:grid-cols-3">
            <div v-for="item in connectionValues" :key="item.label" class="border-2 border-slate-950 bg-white p-3 dark:border-white dark:bg-slate-900"><div class="flex items-center justify-between gap-2"><span class="text-xs font-black uppercase tracking-wider text-slate-500">{{ item.label }}</span><button type="button" class="text-blue-600 hover:text-blue-800 dark:text-blue-400" :title="t('apiDocs.copy')" @click="copyValue(item.value)"><Icon name="copy" size="xs" /></button></div><code class="mt-2 block truncate text-sm font-bold">{{ item.value }}</code></div>
          </div>
        </section>

        <section id="workbuddy" class="scroll-mt-24 border-t-4 border-slate-950 pt-10 dark:border-white">
          <p class="text-sm font-black uppercase tracking-widest text-blue-600 dark:text-blue-400">02 / 03</p>
          <h2 class="mt-2 text-3xl font-black">{{ t('apiDocs.workbuddyTitle') }}</h2>
          <p class="mt-3 max-w-2xl font-semibold leading-7 text-slate-600 dark:text-slate-300">{{ t('apiDocs.workbuddyDescription') }}</p>
          <div class="mt-8 space-y-10">
            <article v-for="step in workbuddySteps" :key="step.image" class="grid gap-5 border-b-2 border-slate-200 pb-10 last:border-b-0 last:pb-0 dark:border-slate-700 lg:grid-cols-[minmax(0,0.9fr)_minmax(0,1.1fr)] lg:items-start">
              <div>
                <p class="text-xs font-black uppercase tracking-widest text-blue-600 dark:text-blue-400">{{ t('apiDocs.stepLabel', { number: step.number }) }}</p>
                <h3 class="mt-3 text-2xl font-black">{{ t(step.title) }}</h3>
                <p class="mt-3 text-sm font-semibold leading-6 text-slate-600 dark:text-slate-300">{{ t(step.description) }}</p>
              </div>
              <figure class="overflow-hidden border-4 border-slate-950 bg-white shadow-[6px_6px_0_0_#2563eb] dark:border-white dark:bg-slate-900">
                <img :src="step.image" :alt="t(step.title)" class="block h-auto w-full" loading="lazy" />
              </figure>
            </article>
          </div>
        </section>

        <section id="examples" class="scroll-mt-24 border-t-4 border-slate-950 pt-10 dark:border-white">
          <p class="text-sm font-black uppercase tracking-widest text-blue-600 dark:text-blue-400">03 / 03</p><h2 class="mt-2 text-3xl font-black">{{ t('apiDocs.examplesTitle') }}</h2><p class="mt-3 font-semibold text-slate-600 dark:text-slate-300">{{ t('apiDocs.examplesDescription') }}</p>
          <div class="mt-6 flex flex-wrap gap-2" role="tablist">
            <button v-for="tab in tabs" :key="tab.key" type="button" role="tab" :aria-selected="activeTab === tab.key" class="border-2 border-slate-950 px-3 py-2 text-sm font-black dark:border-white" :class="activeTab === tab.key ? 'bg-blue-600 text-white' : 'bg-white dark:bg-slate-900'" @click="activeTab = tab.key">{{ tab.label }}</button>
          </div>
          <div class="relative mt-3 border-4 border-slate-950 bg-slate-950 shadow-[6px_6px_0_0_#2563eb] dark:border-white">
            <button type="button" class="absolute right-3 top-3 inline-flex items-center gap-1 border-2 border-slate-600 px-2 py-1 text-xs font-black text-white hover:bg-slate-800" @click="copyValue(activeCode)"><Icon :name="copied ? 'check' : 'copy'" size="xs" />{{ copied ? t('apiDocs.copied') : t('apiDocs.copy') }}</button>
            <pre class="overflow-x-auto p-5 pt-14 text-sm leading-7 text-slate-200"><code>{{ activeCode }}</code></pre>
          </div>
        </section>

        <section id="faq" class="scroll-mt-24 border-t-4 border-slate-950 pt-10 dark:border-white"><h2 class="text-3xl font-black">{{ t('apiDocs.commonIssues') }}</h2><div class="mt-6 divide-y-2 divide-slate-200 border-y-2 border-slate-950 dark:divide-slate-700 dark:border-white"><details v-for="item in faqs" :key="item.title" class="group py-4"><summary class="flex cursor-pointer list-none items-center justify-between gap-4 font-black"><span>{{ item.title }}</span><Icon name="chevronDown" size="sm" class="transition group-open:rotate-180" /></summary><p class="max-w-3xl pt-3 text-sm font-semibold leading-6 text-slate-600 dark:text-slate-300">{{ item.answer }}</p></details></div><div class="mt-8 border-4 border-slate-950 bg-amber-200 p-5 dark:border-white dark:bg-amber-950"><div class="flex gap-3"><Icon name="shield" size="lg" /><div><h3 class="font-black">{{ t('apiDocs.securityTitle') }}</h3><p class="mt-1 text-sm font-semibold">{{ t('apiDocs.securityDescription') }}</p></div></div></div></section>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAuthStore, useAppStore } from '@/stores'
import { useTheme } from '@/composables/useTheme'
import { sanitizeUrl } from '@/utils/url'
import LocaleSwitcher from '@/components/common/LocaleSwitcher.vue'
import Icon from '@/components/icons/Icon.vue'

const { t, locale } = useI18n()
const authStore = useAuthStore()
const appStore = useAppStore()
const { isDark, toggleTheme } = useTheme()
const isAuthenticated = computed(() => authStore.isAuthenticated)
const siteName = computed(() => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'Sub2API')
const siteLogo = computed(() => sanitizeUrl(appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || '', { allowRelative: true, allowDataUrl: true }))
const baseUrl = 'https://myworkbuddy.vip/v1'
const connectionValues = computed(() => [
  { label: t('apiDocs.baseUrl'), value: baseUrl },
  { label: t('apiDocs.apiKey'), value: 'YOUR_API_KEY' },
  { label: t('apiDocs.model'), value: 'YOUR_MODEL' },
])
const workbuddySteps = [
  { number: 1, title: 'apiDocs.workbuddyStepOneTitle', description: 'apiDocs.workbuddyStepOneDescription', image: '/docs/第一步-点击自定义模型.png' },
  { number: 2, title: 'apiDocs.workbuddyStepTwoTitle', description: 'apiDocs.workbuddyStepTwoDescription', image: '/docs/第二步-点击自定义提供商.png' },
  { number: 3, title: 'apiDocs.workbuddyStepThreeTitle', description: 'apiDocs.workbuddyStepThreeDescription', image: '/docs/第三步-配置url-apikey-model.png' },
]
const activeTab = ref<'curl' | 'javascript' | 'python'>('curl')
const copied = ref(false)
const tabs = computed(() => [
  { key: 'curl' as const, label: t('apiDocs.curl') },
  { key: 'javascript' as const, label: t('apiDocs.javascript') },
  { key: 'python' as const, label: t('apiDocs.python') },
])
const activeCode = computed(() => ({
  curl: `curl ${baseUrl}/chat/completions \\\n+  -H "Authorization: Bearer YOUR_API_KEY" \\\n+  -H "Content-Type: application/json" \\\n+  -d '{"model":"YOUR_MODEL","messages":[{"role":"user","content":"Hello"}]}'`,
  javascript: `const client = new OpenAI({\n  apiKey: process.env.SUB2API_API_KEY,\n  baseURL: '${baseUrl}'\n})\n\nconst response = await client.chat.completions.create({\n  model: 'YOUR_MODEL',\n  messages: [{ role: 'user', content: 'Hello' }]\n})\nconsole.log(response.choices[0].message.content)`,
  python: `from openai import OpenAI\n\nclient = OpenAI(\n    api_key='YOUR_API_KEY',\n    base_url='${baseUrl}'\n)\n\nresponse = client.chat.completions.create(\n    model='YOUR_MODEL',\n    messages=[{'role': 'user', 'content': 'Hello'}]\n)\nprint(response.choices[0].message.content)`,
}[activeTab.value]))
const faqs = computed(() => [
  { title: t('apiDocs.faqKeyTitle'), answer: t('apiDocs.faqKeyAnswer') },
  { title: t('apiDocs.faq401Title'), answer: t('apiDocs.faq401Answer') },
  { title: t('apiDocs.faq404Title'), answer: t('apiDocs.faq404Answer') },
  { title: t('apiDocs.faqUrlTitle'), answer: t('apiDocs.faqUrlAnswer') },
])

async function copyValue(value: string) {
  try { await navigator.clipboard.writeText(value) } catch { /* clipboard may be unavailable on HTTP */ }
  copied.value = true
  window.setTimeout(() => { copied.value = false }, 1600)
}

onMounted(() => {
  authStore.checkAuth()
  if (!appStore.publicSettingsLoaded) appStore.fetchPublicSettings()
  document.documentElement.setAttribute('lang', locale.value)
})
</script>
