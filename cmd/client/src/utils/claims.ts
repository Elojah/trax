import { ulid } from './ulid';
import { Command, Resource } from '@internal/user/role';
import type { ClaimAuth } from '@internal/user/claims';

/**
 * Check if a user has a specific claim (permission) for a group, resource, and command
 * @param claims - The user's claims object
 * @param groupId - The group ID (as Uint8Array)
 * @param resource - The resource type
 * @param command - The command type
 * @returns boolean - Whether the claim exists or not
 */
export function hasClaim(
  claims: ClaimAuth | null | undefined,
  groupId: Uint8Array,
  resource: Resource,
  command: Command
): boolean {
  if (!claims || !groupId) {
    return false;
  }

  return claims.groups?.[ulid(groupId)]?.resources?.[Resource[resource]]?.commands?.[Command[command]] !== undefined;
}

// Example usage:
// import { hasClaim } from '@/utils/claims';
// import { Resource, Command } from '@internal/user/role';
//
// const canReadRoles = hasClaim(claims.value, group.value?.group?.iD, Resource.R_role, Command.C_read);
// if (canReadRoles) {
//   // User has permission to read roles in this group
// }
