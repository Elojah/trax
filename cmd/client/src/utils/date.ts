/**
 * Date formatting utilities for the TRAX application
 */

/**
 * Formats a bigint timestamp to a localized date string
 * @param timestamp - BigInt timestamp in seconds
 * @param options - Intl.DateTimeFormatOptions for customization
 * @returns Formatted date string
 */
export function formatDate(
  timestamp: bigint | undefined,
  options: Intl.DateTimeFormatOptions = {
    day: 'numeric',
    month: 'short',
    year: 'numeric'
  }
): string {
  if (!timestamp) {
    return 'Unknown';
  }

  return new Date(Number(timestamp) * 1000).toLocaleDateString('en-GB', options);
}

/**
 * Formats a timestamp to a full date and time string
 * @param timestamp - BigInt timestamp in seconds
 * @returns Formatted date and time string
 */
export function formatDateTime(timestamp: bigint | undefined): string {
  return formatDate(timestamp, {
    day: 'numeric',
    month: 'short',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  });
}

/**
 * Formats a timestamp to a relative time string (e.g., "2 hours ago")
 * @param timestamp - BigInt timestamp in seconds
 * @returns Relative time string
 */
export function formatRelativeTime(timestamp: bigint | undefined): string {
  if (!timestamp) {
    return 'Unknown';
  }

  const now = Date.now();
  const past = Number(timestamp) * 1000;
  const diffMs = now - past;

  const seconds = Math.floor(diffMs / 1000);
  const minutes = Math.floor(seconds / 60);
  const hours = Math.floor(minutes / 60);
  const days = Math.floor(hours / 24);

  if (days > 0) {
    return `${days} day${days > 1 ? 's' : ''} ago`;
  }
  if (hours > 0) {
    return `${hours} hour${hours > 1 ? 's' : ''} ago`;
  }
  if (minutes > 0) {
    return `${minutes} minute${minutes > 1 ? 's' : ''} ago`;
  }
  return 'Just now';
}
