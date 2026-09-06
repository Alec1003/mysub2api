<template>
  <AppLayout>
    <div class="space-y-5">
      <div class="flex flex-wrap items-center justify-between gap-4">
        <div>
          <h1 class="text-2xl font-semibold text-gray-900 dark:text-white">{{ t('skills.adminTitle') }}</h1>
          <p class="mt-1 text-sm text-gray-500 dark:text-dark-400">{{ t('skills.adminHint') }}</p>
        </div>
        <button class="btn btn-primary shrink-0" @click="open()">{{ t('common.create') }}</button>
      </div>

      <div class="overflow-x-auto rounded-lg border border-gray-200 dark:border-dark-700">
        <table class="w-full min-w-[760px] table-fixed divide-y divide-gray-200 dark:divide-dark-700">
          <colgroup>
            <col class="w-[34%]" />
            <col class="w-[24%]" />
            <col class="w-[14%]" />
            <col class="w-[28%]" />
          </colgroup>
          <thead class="bg-gray-50 dark:bg-dark-800/50">
            <tr>
              <th class="px-4 py-3 text-left text-sm font-semibold text-gray-700 dark:text-dark-200">{{ t('skills.name') }}</th>
              <th class="px-4 py-3 text-left text-sm font-semibold text-gray-700 dark:text-dark-200">{{ t('skills.effects') }}</th>
              <th class="px-4 py-3 text-left text-sm font-semibold text-gray-700 dark:text-dark-200">{{ t('common.status') }}</th>
              <th class="px-4 py-3 text-left text-sm font-semibold text-gray-700 dark:text-dark-200">{{ t('common.actions') }}</th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-100 dark:divide-dark-800">
            <tr v-for="skill in items" :key="skill.id" class="align-middle transition-colors hover:bg-gray-50 dark:hover:bg-dark-800/30">
              <td class="px-4 py-4">
                <div class="min-w-0">
                  <p class="truncate font-medium text-gray-900 dark:text-white">{{ skill.name }}</p>
                  <p class="mt-1 line-clamp-2 text-sm leading-5 text-gray-500 dark:text-dark-400">{{ skill.description }}</p>
                </div>
              </td>
              <td class="px-4 py-4">
                <img
                  v-if="skill.effect_image"
                  :src="skill.effect_image"
                  :alt="skill.name"
                  class="aspect-video h-20 w-32 rounded-md border border-gray-200 object-cover dark:border-dark-600"
                />
                <span v-else class="text-sm text-gray-400 dark:text-dark-500">-</span>
              </td>
              <td class="px-4 py-4">
                <span class="badge" :class="skill.status === 'active' ? 'badge-success' : 'badge-gray'">
                  {{ skill.status === 'active' ? t('skills.statusActive') : t('skills.statusOffline') }}
                </span>
              </td>
              <td class="px-4 py-4">
                <div class="flex flex-wrap items-center gap-2">
                  <button class="btn btn-secondary" @click="open(skill)">{{ t('common.edit') }}</button>
                  <button class="btn btn-danger" @click="remove(skill)">{{ t('common.delete') }}</button>
                </div>
              </td>
            </tr>
            <tr v-if="items.length === 0">
              <td colspan="4" class="px-4 py-10 text-center text-sm text-gray-500 dark:text-dark-400">{{ t('common.noData') }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="form" class="fixed inset-0 z-50 flex items-center justify-center bg-black/40 p-4">
        <form class="w-full max-w-lg space-y-4 bg-white p-6 dark:bg-dark-800" @submit.prevent="save">
          <h2 class="text-lg font-semibold">{{ form.id ? t('common.edit') : t('common.create') }} Skill</h2>
          <input v-model="form.name" class="input" required :placeholder="t('skills.name')" />
          <input v-model="form.description" class="input" required :placeholder="t('skills.description')" />
          <div>
            <label class="input-label">{{ t('skills.effects') }}</label>
            <input type="file" accept="image/*" class="input mt-1" :required="!form.id" @change="handleImageUpload" />
            <img v-if="form.effect_image" :src="form.effect_image" :alt="t('skills.effects')" class="mt-3 aspect-video w-full object-cover" />
          </div>
          <input v-model="form.drive_url" class="input" required placeholder="网盘链接" />
          <select v-model="form.status" class="input">
            <option value="active">{{ t('skills.statusActive') }}</option>
            <option value="offline">{{ t('skills.statusOffline') }}</option>
          </select>
          <div class="flex justify-end gap-2">
            <button type="button" class="btn btn-secondary" @click="form = null">{{ t('common.cancel') }}</button>
            <button class="btn btn-primary" :disabled="saving">{{ saving ? t('common.saving') : t('common.save') }}</button>
          </div>
        </form>
      </div>
    </div>
  </AppLayout>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import AppLayout from '@/components/layout/AppLayout.vue'
import api from '@/api/admin/skills'
import type { Skill } from '@/api/skills'
import { useAppStore } from '@/stores/app'
import { extractApiErrorMessage } from '@/utils/apiError'

const { t } = useI18n()
const app = useAppStore()
const items = ref<Skill[]>([])
const form = ref<Partial<Skill> | null>(null)
const saving = ref(false)
async function load() { try { items.value = await api.list() } catch (error) { app.showError(extractApiErrorMessage(error, t('common.error'))) } }
function open(skill?: Skill) { form.value = skill ? { ...skill } : { name: '', description: '', effect_image: '', drive_url: '', status: 'active' } }
function handleImageUpload(event: Event) { const file = (event.target as HTMLInputElement).files?.[0]; if (!file || !file.type.startsWith('image/')) return; if (file.size > 6 * 1024 * 1024) { app.showError(t('skills.imageTooLarge')); return }; const reader = new FileReader(); reader.onload = () => { if (form.value && typeof reader.result === 'string') form.value.effect_image = reader.result }; reader.readAsDataURL(file) }
async function save() { if (!form.value) return; saving.value = true; try { const payload = { name: form.value.name || '', description: form.value.description || '', effect_image: form.value.effect_image || '', drive_url: form.value.drive_url || '', status: form.value.status || 'active' }; if (form.value.id) await api.update(form.value.id, payload); else await api.create(payload); form.value = null; await load() } catch (error) { app.showError(extractApiErrorMessage(error, t('common.error'))) } finally { saving.value = false } }
async function remove(skill: Skill) { if (!confirm(t('common.confirm'))) return; try { await api.remove(skill.id); await load() } catch (error) { app.showError(extractApiErrorMessage(error, t('common.error'))) } }
onMounted(load)
</script>
