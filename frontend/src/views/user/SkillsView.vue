<template>
  <AppLayout>
    <div class="mx-auto max-w-6xl space-y-6">
      <div class="flex flex-wrap items-end justify-between gap-4">
        <div><h1 class="text-2xl font-semibold text-gray-900 dark:text-white">{{ t('skills.title') }}</h1><p class="mt-1 text-sm text-gray-500">{{ t('skills.rule') }}</p></div>
        <div class="flex gap-5 text-sm"><span>{{ t('skills.total') }} <b>{{ Math.floor(ent.total_recharge_points).toLocaleString() }}</b></span><span>{{ t('skills.available') }} <b class="text-primary-600">{{ ent.available_redemptions }}</b></span></div>
      </div>
      <div v-if="loading" class="py-16 text-center text-gray-500">{{ t('common.loading') }}</div>
      <div v-else class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
        <div v-for="skill in items" :key="skill.id" class="border border-gray-200 bg-white p-5 shadow-sm dark:border-dark-700 dark:bg-dark-800">
          <div class="flex items-start justify-between gap-3"><h2 class="font-semibold text-gray-900 dark:text-white">{{ skill.name }}</h2><span :class="skill.redeemed ? 'badge-success' : 'badge-gray'" class="badge">{{ skill.redeemed ? t('skills.redeemed') : t('skills.notRedeemed') }}</span></div>
          <p class="mt-3 text-sm text-gray-600 dark:text-gray-300">{{ skill.description }}</p>
          <button type="button" class="group mt-4 block w-full overflow-hidden text-left" @click="openPreview(skill.effect_image)"><img :src="skill.effect_image" :alt="skill.name" class="aspect-video w-full object-cover transition duration-200 group-hover:scale-[1.02]" /><span class="mt-2 block text-xs text-gray-500">{{ t('skills.effects') }} · {{ t('skills.clickToPreview') }}</span></button>
          <div class="mt-5"><a v-if="skill.redeemed && skill.drive_url" :href="skill.drive_url" target="_blank" rel="noopener" class="btn btn-secondary w-full">{{ t('skills.openMaterial') }}</a><button v-else class="btn btn-primary w-full" :disabled="ent.available_redemptions < 1 || redeeming === skill.id" @click="handleRedeem(skill)">{{ redeeming === skill.id ? t('common.processing') : t('skills.redeem') }}</button></div>
        </div>
      </div>
    </div>
    <div v-if="previewImage" class="fixed inset-0 z-50 flex items-center justify-center bg-black/80 p-4" @click.self="closePreview"><button type="button" class="absolute right-5 top-5 text-3xl leading-none text-white" :aria-label="t('common.close')" @click="closePreview">&times;</button><img :src="previewImage" class="max-h-[90vh] max-w-[94vw] object-contain" :alt="t('skills.effects')" /></div>
  </AppLayout>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import AppLayout from '@/components/layout/AppLayout.vue'
import api, { type Entitlement, type Skill } from '@/api/skills'
import { useAppStore } from '@/stores/app'
import { extractApiErrorMessage } from '@/utils/apiError'

const { t } = useI18n()
const app = useAppStore()
const items = ref<Skill[]>([])
const ent = ref<Entitlement>({ total_recharge_points: 0, earned_redemptions: 0, used_redemptions: 0, available_redemptions: 0 })
const loading = ref(false)
const redeeming = ref<number | null>(null)
const previewImage = ref('')
async function load() { loading.value = true; try { const result = await api.list(); items.value = result.items; ent.value = result.entitlement } catch (error) { app.showError(extractApiErrorMessage(error, t('common.error'))) } finally { loading.value = false } }
async function handleRedeem(skill: Skill) { redeeming.value = skill.id; try { await api.redeem(skill.id); app.showSuccess(t('skills.redeemSuccess')); await load() } catch (error) { app.showError(extractApiErrorMessage(error, t('common.error'))) } finally { redeeming.value = null } }
function openPreview(image: string) { previewImage.value = image }
function closePreview() { previewImage.value = '' }
onMounted(load)
</script>
