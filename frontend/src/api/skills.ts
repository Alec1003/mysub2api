import { apiClient } from './client'
export interface Skill { id:number; name:string; description:string; effect_image:string; drive_url?:string; status:string; redeemed?:boolean; created_at:string; updated_at:string }
export interface Entitlement { total_recharge_points:number; earned_redemptions:number; used_redemptions:number; available_redemptions:number }
export async function list():Promise<{items:Skill[]; entitlement:Entitlement}>{const {data}=await apiClient.get('/skills');return data}
export async function redeem(id:number):Promise<Skill>{const {data}=await apiClient.post(`/skills/${id}/redeem`);return data}
export default {list,redeem}
