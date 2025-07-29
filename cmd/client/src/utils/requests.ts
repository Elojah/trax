/**
 * Utility functions for building common API requests
 * Provides standardized request creation for list operations
 */

import type { DataTableProps } from 'primevue/datatable';

export interface PaginationConfig {
  start: bigint;
  end: bigint;
  sort: string;
  order: boolean; // true for asc, false for desc
}

/**
 * Creates pagination configuration from DataTable properties
 * @param props - DataTable properties
 * @param defaultSortField - Default field to sort by
 * @returns Pagination configuration
 */
export function createPagination(
  props: DataTableProps,
  defaultSortField: string = 'created_at'
): PaginationConfig {
  const page = Math.floor((props.first ?? 0) / (props.rows ?? 10)) + 1;
  const rows = props.rows ?? 10;

  // Handle sort configuration
  const sortField = props.sortField ? String(props.sortField) : defaultSortField;
  const sortOrder = props.sortOrder === 1 ? 'asc' : 'desc';

  return {
    start: BigInt(((page - 1) * rows) + 1),
    end: BigInt(page * rows),
    sort: sortField,
    order: sortOrder === 'asc'
  };
}

/**
 * Creates sort configuration array for API requests
 * @param props - DataTable properties
 * @param defaultSortField - Default field to sort by
 * @returns Array of sort configurations
 */
export function createSortConfig(
  props: DataTableProps,
  defaultSortField: string = 'created_at'
): Array<{ key: string; order: 'asc' | 'desc' }> {
  const sortField = props.sortField ? String(props.sortField) : defaultSortField;
  const sortOrder = props.sortOrder === 1 ? 'asc' : 'desc';

  return [{
    key: sortField,
    order: sortOrder
  }];
}
