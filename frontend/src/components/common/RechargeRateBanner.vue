<template>
  <div
    role="status"
    class="border-b-4 border-slate-950 bg-amber-300 px-4 py-2.5 text-slate-950 dark:border-white dark:bg-amber-400"
  >
    <div class="mx-auto flex max-w-7xl items-center justify-center gap-2 text-center text-sm font-black leading-6 sm:text-base">
      <Icon name="creditCard" size="sm" class="shrink-0" aria-hidden="true" />
      <span>{{ t('home.brutal.rechargeBanner') }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import Icon from '@/components/icons/Icon.vue'

const { t } = useI18n()
</script>
