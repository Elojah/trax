import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { Invitation, InvitationRole } from '@internal/user/invitation'
import { InvitationView, ListInvitationReq, ListInvitationResp } from '@internal/user/dto/invitation'
import { useAuthStore } from '@/stores/auth'
import { computed, ref } from 'vue'
import { parse, ulid } from '@/utils/ulid'
import type { U } from '@internal/user/user'

export const useInvitationStore = defineStore('invitation', () => {
  const invitations = ref<Map<string, InvitationView>>(new Map())
  const total = ref<bigint>(BigInt(0))

  const selected = ref<InvitationView[]>([])

  const api = new APIClient(
    new GrpcWebFetchTransport({
      baseUrl: config.api_url
    })
  )
  const authStore = useAuthStore()
  const token = computed(() => authStore.token)

  const create = async function (email: string, groupID: Uint8Array, roleIDs: Uint8Array[]) {
    try {
      const req = {
        email: email,
        groupID: groupID,
        roleIDs: roleIDs,
      }

      const resp: { response: U } = await api.createInvitation(req, { meta: { token: token.value } })

      return resp
    } catch (err: any) {
      logger.error(err)
      throw err
    }

  }

  // General list function for invitations
  const list = async function (req: ListInvitationReq): Promise<string[]> {
    try {
      const resp: { response: ListInvitationResp } = await api.listInvitation(req, { meta: { token: token.value } })

      resp.response.invitations?.forEach((invitationView: InvitationView) => {
        const id = ulid(invitationView.invitation?.iD)
        invitations.value?.set(id, invitationView)
      })

      if (req?.iDs.length === 0) {
        total.value = resp.response.total
      }

      return resp.response.invitations.map((invitationView: InvitationView) => ulid(invitationView.invitation?.iD))
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  // Mock function for deleting an invitation
  // This would need to be implemented in the backend
  const deleteInvitation = async function (invitationId: string) {
    try {
      // TODO: Implement backend endpoint for deleting invitations
      // For now, just remove from local state
      invitations.value?.delete(invitationId)
      logger.info(`Mock: Deleted invitation ${invitationId}`)
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  // Mock function for resending an invitation
  // This would need to be implemented in the backend
  const resendInvitation = async function (invitationId: string) {
    try {
      // TODO: Implement backend endpoint for resending invitations
      const invitation = invitations.value?.get(invitationId)
      if (invitation?.invitation) {
        invitation.invitation.updatedAt = BigInt(Date.now())
        invitations.value?.set(invitationId, invitation)
      }
      logger.info(`Mock: Resent invitation ${invitationId}`)
    } catch (err: any) {
      logger.error(err)
      throw err
    }
  }

  const populate = async function (ids: string[]) {
    const newIDs = ids.reduce((acc: string[], id: string) => {
      if (invitations.value?.has(id)) {
        return acc
      } else {
        return [...acc, id]
      }
    }, [])

    if (newIDs.length === 0) {
      return
    }

    // TODO: Implement actual populate logic when backend endpoint exists
    logger.info(`Mock: Would populate invitations ${newIDs.join(', ')}`)
  }

  return {
    invitations,
    total,
    selected,
    create,
    list,
    deleteInvitation,
    resendInvitation,
    populate
  }
})
