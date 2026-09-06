import { apiClient } from '../client'
import type { Skill } from '../skills'
export async function list():Promise<Skill[]>{const {data}=await apiClient.get('/admin/skills');return data}
export async function create(input:Omit<Skill,'id'|'created_at'|'updated_at'|'redeemed'>):Promise<Skill>{const {data}=await apiClient.post('/admin/skills',input);return data}
export async function update(id:number,input:Omit<Skill,'id'|'created_at'|'updated_at'|'redeemed'>):Promise<Skill>{const {data}=await apiClient.put(`/admin/skills/${id}`,input);return data}
export async function remove(id:number){return apiClient.delete(`/admin/skills/${id}`)}
export default {list,create,update,remove}
