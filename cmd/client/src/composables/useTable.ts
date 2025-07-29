/**
 * Composable combining DataTable and ListManager functionality
 * Provides a complete solution for table-based list components
 */

import { type Ref } from 'vue';
import type { DataTableProps } from 'primevue/datatable';
import { useDataTable, type UseDataTableOptions } from './useDataTable';
import { useListManager, type UseListManagerOptions } from './useListManager';

export interface UseTableOptions<TRequest, TItem> extends UseDataTableOptions {
  /** Function to create the list request from DataTable props and search */
  createRequest: (props: DataTableProps, search: string) => TRequest;
  /** Function to call the API list method */
  listMethod: (request: TRequest) => Promise<string[]>;
  /** Map containing all items keyed by ID */
  itemsMap: Ref<Map<string, TItem>>;
  /** Total count of items */
  total: Ref<bigint>;
}

/**
 * Combined composable for table functionality
 * @param options - Configuration options
 */
export function useTable<TRequest, TItem>(options: UseTableOptions<TRequest, TItem>) {
  const {
    createRequest,
    listMethod,
    itemsMap,
    total,
    defaultRows = 10,
    defaultSortField = 'created_at',
    defaultSortOrder = -1
  } = options;

  // Create list manager
  const listManager = useListManager({
    createRequest: (props: DataTableProps) => createRequest(props, search.value),
    listMethod,
    itemsMap,
    total,
    defaultProps: {
      first: 0,
      rows: defaultRows,
      sortField: defaultSortField,
      sortOrder: defaultSortOrder,
    }
  });

  // Create DataTable handler
  const dataTable = useDataTable(listManager.list, {
    defaultRows,
    defaultSortField,
    defaultSortOrder
  });

  // Override the search functionality to include search in request creation
  const { search } = dataTable;

  return {
    // List manager exports
    viewIDs: listManager.viewIDs,
    views: listManager.views,

    // DataTable exports
    properties: dataTable.properties,
    loading: listManager.loading, // Use loading from list manager
    search,
    handlers: dataTable.handlers,
    utils: dataTable.utils,

    // Combined functionality
    list: listManager.list
  };
}
