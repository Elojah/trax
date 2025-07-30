import { defineStore } from 'pinia'
import { config, logger } from '@/config'
import { APIClient } from '@api/api.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'
import { Invitation, InvitationRole } from '@internal/user/invitation'
import { useAuthStore } from '@/stores/auth'
import { computed, ref } from 'vue'
import { parse, ulid } from '@/utils/ulid'

export interface InvitationView {
  invitation?: Invitation
  roles?: string[] // Role names
  roleCount?: number
}

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

  // Mock function for listing invitations by group ID
  // This would need to be implemented in the backend
  const listByGroup = async function (groupId: string): Promise<string[]> {
    try {
      // TODO: Implement backend endpoint for listing invitations by group
      // For now, return mock data
      const mockInvitations: InvitationView[] = [
        {
          invitation: {
            iD: new Uint8Array([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]),
            email: 'john.doe@example.com',
            createdAt: BigInt(Date.now() - 86400000), // 1 day ago
            updatedAt: BigInt(Date.now() - 86400000)
          },
          roles: ['Admin', 'Editor'],
          roleCount: 2
        },
        {
          invitation: {
            iD: new Uint8Array([2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]),
            email: 'jane.smith@example.com',
            createdAt: BigInt(Date.now() - 172800000), // 2 days ago
            updatedAt: BigInt(Date.now() - 172800000)
          },
          roles: ['Viewer'],
          roleCount: 1
        }
      ]

      mockInvitations.forEach((invitation: InvitationView) => {
        const id = ulid(invitation.invitation?.iD)
        invitations.value?.set(id, invitation)
      })

      total.value = BigInt(mockInvitations.length)

      return mockInvitations.map((invitation: InvitationView) => ulid(invitation.invitation?.iD))
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
    listByGroup,
    deleteInvitation,
    resendInvitation,
    populate
  }
})
