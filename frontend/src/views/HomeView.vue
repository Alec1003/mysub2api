<template>
  <div v-if="hasHomeContent" class="min-h-screen">
    <iframe v-if="isHomeContentUrl" :src="homeContent.trim()" class="h-screen w-full border-0" allowfullscreen />
    <div v-else v-html="homeContent"></div>
  </div>

  <div v-else-if="compactHomeEnabled" data-testid="compact-home" class="flex min-h-screen flex-col bg-slate-50 text-slate-950 dark:bg-slate-950 dark:text-white">
    <header class="border-b-4 border-slate-950 px-4 py-4 dark:border-white"><nav class="mx-auto flex max-w-5xl items-center justify-between gap-4"><div class="font-black">{{ siteName }}</div><div class="flex items-center gap-2"><LocaleSwitcher /><router-link v-if="showModelPlazaEntry" to="/model-plaza" class="border-2 border-slate-950 px-3 py-2 text-sm font-bold">{{ t('home.brutal.navModelPlaza') }}</router-link><router-link :to="isAuthenticated ? dashboardPath : '/login'" class="border-2 border-slate-950 bg-blue-600 px-3 py-2 text-sm font-bold text-white">{{ isAuthenticated ? t('home.brutal.goToDashboard') : t('home.login') }}</router-link></div></nav></header>
    <main class="flex flex-1 items-center justify-center px-5 py-16"><div class="text-center"><h1 class="text-4xl font-black">{{ siteName }}</h1><p class="mt-4 text-slate-600 dark:text-slate-300">{{ appStore.cachedPublicSettings?.site_subtitle || 'AI API Gateway Platform' }}</p><router-link :to="isAuthenticated ? dashboardPath : '/login'" class="mt-8 inline-flex border-2 border-slate-950 bg-blue-600 px-5 py-3 font-bold text-white">{{ isAuthenticated ? t('home.goToDashboard') : t('home.login') }}</router-link></div></main>
  </div>

  <div v-else class="min-h-screen bg-slate-50 text-slate-950 dark:bg-slate-950 dark:text-white">
    <header class="sticky top-0 z-30 border-b-4 border-slate-950 bg-blue-600 dark:border-white">
      <nav class="mx-auto flex max-w-7xl items-center justify-between gap-4 px-4 py-4 sm:px-8">
        <router-link to="/home" class="flex items-center gap-3" aria-label="MyWorkBuddy home">
          <span class="flex h-10 w-10 items-center justify-center overflow-hidden border-2 border-slate-950 bg-white text-xl font-black text-blue-600 shadow-[4px_4px_0_0_#0f172a]"><img v-if="siteLogo" :src="siteLogo" :alt="siteName" class="h-full w-full object-contain" /><span v-else>M</span></span>
          <span class="text-lg font-black tracking-tight text-white">{{ siteName.toUpperCase() }}<span class="text-blue-200">.</span></span>
        </router-link>
        <div class="hidden items-center gap-7 text-sm font-bold text-white md:flex">
          <a href="#features" class="hover:text-blue-100">{{ t('home.brutal.navFeatures') }}</a>
          <a href="#models" class="hover:text-blue-100">{{ t('home.brutal.navModels') }}</a>
          <a href="#pricing" class="hover:text-blue-100">{{ t('home.brutal.navPricing') }}</a>
          <a href="#faq" class="hover:text-blue-100">{{ t('home.brutal.navFaq') }}</a>
          <router-link v-if="showModelPlazaEntry" to="/model-plaza" class="hover:text-blue-100">{{ t('home.brutal.navModelPlaza') }}</router-link>
        </div>
        <div class="flex items-center gap-2">
          <LocaleSwitcher />
          <button class="border-2 border-slate-950 bg-white p-2 text-slate-950 shadow-[3px_3px_0_0_#0f172a] hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none" :title="isDark ? t('home.switchToLight') : t('home.switchToDark')" @click="toggleTheme">
            <Icon v-if="isDark" name="sun" size="sm" /><Icon v-else name="moon" size="sm" />
          </button>
          <router-link :to="isAuthenticated ? dashboardPath : '/login'" class="hidden border-2 border-slate-950 bg-white px-4 py-2 text-sm font-black text-slate-950 shadow-[4px_4px_0_0_#0f172a] transition hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none sm:inline-flex">
            {{ isAuthenticated ? t('home.brutal.goToDashboard') : t('home.login') }}
          </router-link>
        </div>
      </nav>
    </header>

    <main>
      <div class="terminal-container hidden" aria-hidden="true"></div>
      <section class="border-b-4 border-slate-950 bg-blue-600 px-4 py-16 sm:px-8 sm:py-24 dark:border-white">
        <div class="mx-auto grid max-w-7xl items-center gap-12 lg:grid-cols-2">
          <div>
            <div class="mb-7 inline-block border-2 border-slate-950 bg-white px-3 py-1 text-xs font-black tracking-widest text-slate-950 shadow-[4px_4px_0_0_#0f172a]">{{ t('home.brutal.heroTag') }}</div>
            <h1 class="max-w-3xl text-5xl font-black leading-[0.95] tracking-tight text-white sm:text-7xl">{{ t('home.brutal.heroLine1') }}<br />{{ t('home.brutal.heroLine2') }}<br /><span class="text-blue-200">{{ t('home.brutal.heroLine3') }}</span></h1>
            <p class="mt-7 max-w-xl text-lg font-bold leading-8 text-blue-50 sm:text-xl">{{ t('home.brutal.heroDescription') }}</p>
            <div class="mt-9 flex flex-wrap gap-4">
              <router-link :to="isAuthenticated ? dashboardPath : '/register'" class="inline-flex items-center gap-2 border-2 border-slate-950 bg-white px-5 py-3 font-black text-slate-950 shadow-[6px_6px_0_0_#0f172a] transition hover:translate-x-[3px] hover:translate-y-[3px] hover:shadow-none">{{ isAuthenticated ? t('home.brutal.goToDashboard') : t('home.brutal.freeRegister') }} <Icon name="arrowRight" size="sm" /></router-link>
              <a href="#pricing" class="inline-flex border-2 border-slate-950 bg-blue-800 px-5 py-3 font-black text-white shadow-[6px_6px_0_0_#0f172a] transition hover:translate-x-[3px] hover:translate-y-[3px] hover:shadow-none">{{ t('home.brutal.viewPricing') }}</a>
            </div>
            <div class="mt-8 flex flex-wrap gap-x-6 gap-y-2 text-sm font-bold text-blue-100"><span>✓ {{ t('home.brutal.trustNoCard') }}</span><span>✓ {{ t('home.brutal.trustSdk') }}</span><span>✓ {{ t('home.brutal.trustUsage') }}</span></div>
          </div>
          <div class="relative mx-auto w-full max-w-xl lg:ml-auto">
            <div class="rotate-2 border-4 border-slate-950 bg-white p-3 shadow-[12px_12px_0_0_#0f172a]">
              <div class="border-4 border-slate-950 bg-slate-950 p-5 font-mono text-sm text-white sm:p-7"><div class="mb-8 flex gap-2"><i class="h-3 w-3 rounded-full bg-red-400"></i><i class="h-3 w-3 rounded-full bg-amber-400"></i><i class="h-3 w-3 rounded-full bg-emerald-400"></i></div><p><span class="text-blue-400">$ curl</span> -X POST /v1/chat/completions</p><p class="mt-4 text-slate-400">model: <span class="text-white">claude-sonnet</span></p><p class="text-slate-400">stream: <span class="text-white">true</span></p><p class="mt-5 border-l-4 border-blue-500 pl-3 text-slate-300">One API. Many possibilities.</p><p class="mt-6"><span class="bg-emerald-500 px-2 py-1 font-bold text-slate-950">200 OK</span> <span class="ml-2 text-slate-500">1.24s</span></p></div>
            </div>
            <div class="absolute -bottom-7 -left-3 -rotate-6 border-4 border-slate-950 bg-blue-200 px-4 py-3 text-sm font-black text-slate-950 shadow-[6px_6px_0_0_#0f172a] sm:-left-8">{{ t('home.brutal.builtForBuilders') }}</div>
          </div>
        </div>
      </section>

      <section class="border-b-4 border-slate-950 bg-slate-950 px-4 py-9 text-white dark:border-white sm:px-8"><div class="mx-auto grid max-w-7xl grid-cols-2 gap-7 text-center sm:grid-cols-4"><div><strong class="block text-3xl font-black text-blue-400 sm:text-4xl">1</strong><span class="text-xs font-bold uppercase tracking-wider text-slate-400">{{ t('home.brutal.statUnified') }}</span></div><div><strong class="block text-3xl font-black text-blue-400 sm:text-4xl">24/7</strong><span class="text-xs font-bold uppercase tracking-wider text-slate-400">{{ t('home.brutal.statStable') }}</span></div><div><strong class="block text-3xl font-black text-blue-400 sm:text-4xl">0</strong><span class="text-xs font-bold uppercase tracking-wider text-slate-400">{{ t('home.brutal.statNoFees') }}</span></div><div><strong class="block text-3xl font-black text-blue-400 sm:text-4xl">∞</strong><span class="text-xs font-bold uppercase tracking-wider text-slate-400">{{ t('home.brutal.statPossibility') }}</span></div></div></section>

      <section id="features" class="mx-auto max-w-7xl px-4 py-20 sm:px-8 sm:py-24"><div class="mb-12 flex flex-col justify-between gap-5 sm:flex-row sm:items-end"><div><span class="inline-block border-2 border-slate-950 bg-blue-200 px-3 py-1 text-xs font-black text-slate-950 shadow-[3px_3px_0_0_#0f172a]">{{ t('home.brutal.featuresEyebrow') }}</span><h2 class="mt-6 text-4xl font-black tracking-tight sm:text-5xl">{{ t('home.brutal.featuresTitleLine1') }}<br />{{ t('home.brutal.featuresTitleLine2') }}</h2></div><p class="max-w-md font-semibold leading-7 text-slate-600 dark:text-slate-300">{{ t('home.brutal.featuresDescription') }}</p></div><div class="grid gap-7 md:grid-cols-3"><article class="border-4 border-slate-950 bg-blue-200 p-7 shadow-[8px_8px_0_0_#0f172a] dark:border-white dark:shadow-[8px_8px_0_0_#2563eb]"><Icon name="server" size="lg" /><h3 class="mt-8 text-2xl font-black">{{ t('home.brutal.featureUnifiedTitle') }}</h3><p class="mt-3 font-semibold leading-7">{{ t('home.brutal.featureUnifiedDesc') }}</p></article><article class="border-4 border-slate-950 bg-white p-7 shadow-[8px_8px_0_0_#2563eb] dark:border-white dark:bg-slate-900"><Icon name="shield" size="lg" class="text-blue-600" /><h3 class="mt-8 text-2xl font-black">{{ t('home.brutal.featureRoutingTitle') }}</h3><p class="mt-3 font-semibold leading-7 text-slate-600 dark:text-slate-300">{{ t('home.brutal.featureRoutingDesc') }}</p></article><article class="border-4 border-slate-950 bg-blue-600 p-7 text-white shadow-[8px_8px_0_0_#0f172a] dark:border-white"><Icon name="chart" size="lg" /><h3 class="mt-8 text-2xl font-black">{{ t('home.brutal.featureUsageTitle') }}</h3><p class="mt-3 font-semibold leading-7 text-blue-50">{{ t('home.brutal.featureUsageDesc') }}</p></article></div></section>

      <section id="models" class="border-y-4 border-slate-950 bg-white px-4 py-20 sm:px-8 sm:py-24 dark:border-white dark:bg-slate-900"><div class="mx-auto max-w-7xl"><span class="inline-block border-2 border-slate-950 bg-blue-200 px-3 py-1 text-xs font-black text-slate-950 shadow-[3px_3px_0_0_#0f172a]">{{ t('home.brutal.modelsEyebrow') }}</span><h2 class="mt-6 text-4xl font-black tracking-tight sm:text-5xl">{{ t('home.brutal.modelsTitleLine1') }}<br />{{ t('home.brutal.modelsTitleLine2') }}</h2><div class="mt-12 grid gap-5 sm:grid-cols-2 lg:grid-cols-4"><div v-for="model in models" :key="model.name" class="flex items-center gap-4 border-4 border-slate-950 bg-slate-50 p-5 shadow-[5px_5px_0_0_#2563eb] dark:border-white dark:bg-slate-950"><span class="flex h-11 w-11 items-center justify-center bg-blue-600 text-lg font-black text-white">{{ model.mark }}</span><div><p class="font-black">{{ model.name }}</p><p class="text-xs font-bold text-slate-500 dark:text-slate-400">{{ model.detail }}</p></div></div></div></div></section>

      <section id="pricing" class="mx-auto max-w-7xl px-4 py-20 sm:px-8 sm:py-24"><div class="grid gap-10 lg:grid-cols-[0.8fr_1.2fr] lg:items-start"><div><span class="inline-block border-2 border-slate-950 bg-blue-200 px-3 py-1 text-xs font-black text-slate-950 shadow-[3px_3px_0_0_#0f172a]">{{ t('home.brutal.pricingEyebrow') }}</span><h2 class="mt-6 text-4xl font-black tracking-tight sm:text-5xl">{{ t('home.brutal.pricingTitleLine1') }}<br />{{ t('home.brutal.pricingTitleLine2') }}</h2><p class="mt-5 font-semibold leading-7 text-slate-600 dark:text-slate-300">{{ t('home.brutal.pricingDescription') }}</p></div><div class="border-4 border-slate-950 bg-blue-600 p-7 text-white shadow-[9px_9px_0_0_#0f172a] dark:border-white"><div class="flex items-center justify-between border-b-2 border-blue-300 pb-5"><h3 class="text-2xl font-black">{{ t('home.brutal.paygTitle') }}</h3><span class="border-2 border-slate-950 bg-white px-2 py-1 text-xs font-black text-slate-950">FAIR</span></div><div class="mt-7 grid gap-5 sm:grid-cols-2"><div class="border-2 border-blue-300 p-4"><p class="text-sm font-bold text-blue-100">{{ t('home.brutal.inputOutput') }}</p><p class="mt-2 text-xl font-black">{{ t('home.brutal.realtimePricing') }}</p></div><div class="border-2 border-blue-300 p-4"><p class="text-sm font-bold text-blue-100">{{ t('home.brutal.accountBalance') }}</p><p class="mt-2 text-xl font-black">{{ t('home.brutal.alwaysCheck') }}</p></div></div><router-link :to="isAuthenticated ? dashboardPath : '/register'" class="mt-7 inline-flex border-2 border-slate-950 bg-white px-5 py-3 font-black text-slate-950 shadow-[5px_5px_0_0_#0f172a] hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none">{{ t('home.brutal.startUsing') }} →</router-link></div></div></section>

      <section id="faq" class="border-t-4 border-slate-950 bg-blue-200 px-4 py-20 sm:px-8 sm:py-24 dark:border-white"><div class="mx-auto max-w-4xl"><span class="inline-block border-2 border-slate-950 bg-white px-3 py-1 text-xs font-black text-slate-950 shadow-[3px_3px_0_0_#0f172a]">{{ t('home.brutal.navFaq') }}</span><h2 class="mt-6 text-4xl font-black tracking-tight sm:text-5xl">{{ t('home.brutal.faqTitle') }}</h2><div class="mt-10 space-y-4"><div v-for="(item, index) in faqs" :key="item.question" class="border-4 border-slate-950 bg-white shadow-[5px_5px_0_0_#0f172a]"><button class="flex w-full items-center justify-between gap-4 p-5 text-left font-black" :aria-expanded="openFaq === index" @click="openFaq = openFaq === index ? -1 : index"><span>{{ item.question }}</span><Icon :name="openFaq === index ? 'chevronUp' : 'chevronDown'" size="sm" /></button><p v-if="openFaq === index" class="border-t-4 border-slate-950 p-5 font-semibold leading-7 text-slate-600">{{ item.answer }}</p></div></div></div></section>
    </main>

      <footer class="border-t-4 border-slate-950 bg-slate-950 px-4 py-10 text-white dark:border-white sm:px-8"><div class="mx-auto flex max-w-7xl flex-col gap-5 sm:flex-row sm:items-center sm:justify-between"><div><p class="text-xl font-black">{{ siteName.toUpperCase() }}<span class="text-blue-400">.</span></p><p class="mt-2 text-sm font-semibold text-slate-400">{{ t('home.brutal.footerTagline') }} · © {{ currentYear }}</p></div><div class="flex gap-5 text-sm font-bold text-slate-300"><router-link to="/login" class="hover:text-blue-400">{{ t('home.login') }}</router-link></div></div></footer>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAuthStore, useAppStore } from '@/stores'
import LocaleSwitcher from '@/components/common/LocaleSwitcher.vue'
import Icon from '@/components/icons/Icon.vue'
import { sanitizeUrl } from '@/utils/url'
import { FeatureFlags, isFeatureFlagEnabled } from '@/utils/featureFlags'
import { useTheme } from '@/composables/useTheme'

const { t } = useI18n()
const authStore = useAuthStore()
const appStore = useAppStore()
const siteName = computed(() => appStore.cachedPublicSettings?.site_name || appStore.siteName || 'MyWorkBuddy')
const siteLogo = computed(() => sanitizeUrl(appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || '', { allowRelative: true, allowDataUrl: true }))
const homeContent = computed(() => appStore.cachedPublicSettings?.home_content || '')
const hasHomeContent = computed(() => homeContent.value.trim().length > 0)
const compactHomeEnabled = computed(() => appStore.cachedPublicSettings?.compact_home_enabled === true)
const isHomeContentUrl = computed(() => /^https?:\/\//.test(homeContent.value.trim()))
const { isDark, toggleTheme } = useTheme()
const isAuthenticated = computed(() => authStore.isAuthenticated)
const dashboardPath = computed(() => authStore.isAdmin ? '/admin/dashboard' : '/dashboard')
const modelPlazaEnabled = computed(() => isFeatureFlagEnabled(FeatureFlags.modelPlaza))
const modelPlazaRequiresAuth = computed(() => appStore.cachedPublicSettings?.model_plaza_require_auth === true)
const showModelPlazaEntry = computed(() => modelPlazaEnabled.value && (isAuthenticated.value || !modelPlazaRequiresAuth.value))
const currentYear = computed(() => new Date().getFullYear())
const openFaq = ref(-1)
const models = computed(() => [{ mark: 'C', name: 'Claude', detail: 'Anthropic' }, { mark: 'G', name: 'GPT', detail: 'OpenAI' }, { mark: 'G', name: 'Gemini', detail: 'Google' }, { mark: '+', name: t('home.brutal.moreModels'), detail: t('home.brutal.ongoing') }])
const faqs = computed(() => [{ question: t('home.brutal.faqStart'), answer: t('home.brutal.faqStartAnswer') }, { question: t('home.brutal.faqFormats'), answer: t('home.brutal.faqFormatsAnswer') }, { question: t('home.brutal.faqUsage'), answer: t('home.brutal.faqUsageAnswer') }])
onMounted(() => { authStore.checkAuth(); if (!appStore.publicSettingsLoaded) appStore.fetchPublicSettings() })
</script>
