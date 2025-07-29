/**
 * Composable for DataTable functionality
 * Provides reusable event handlers and state management for PrimeVue DataTable components
 */

import { ref, type Ref } from 'vue';
import type {
  DataTableFilterEvent,
  DataTablePageEvent,
  DataTableProps,
  DataTableSortEvent
} from 'primevue/datatable';

export interface UseDataTableOptions {
  /** Default number of rows per page */
  defaultRows?: number;
  /** Default sort field */
  defaultSortField?: string;
  /** Default sort order (-1 for desc, 1 for asc) */
  defaultSortOrder?: number;
}

export interface UseDataTableReturn {
  /** DataTable properties (for v-model binding) */
  properties: Ref<DataTableProps>;
  /** Loading state */
  loading: Ref<boolean>;
  /** Search query */
  search: Ref<string>;
  /** Event handlers */
  handlers: {
    onPage: (event: DataTablePageEvent) => void;
    onSort: (event: DataTableSortEvent) => void;
    onFilter: (event: DataTableFilterEvent) => void;
    onSearch: () => void;
  };
  /** Utility functions */
  utils: {
    /** Reset table to first page with current settings */
    resetToFirstPage: () => void;
    /** Get current page number (1-based) */
    getCurrentPage: () => number;
    /** Get sort configuration for API calls */
    getSortConfig: () => Array<{ key: string; order: 'asc' | 'desc' }>;
  };
}

/**
 * Composable for DataTable event handling and state management
 * @param listFunction - Function to call when data needs to be refreshed
 * @param options - Configuration options
 */
export function useDataTable(
  listFunction: (props: DataTableProps) => Promise<void> | void,
  options: UseDataTableOptions = {}
): UseDataTableReturn {
  const {
    defaultRows = 10,
    defaultSortField = 'created_at',
    defaultSortOrder = -1
  } = options;

  // Reactive state
  const properties = ref<DataTableProps>({
    first: 0,
    rows: defaultRows,
    sortField: defaultSortField,
    sortOrder: defaultSortOrder,
  });

  const loading = ref(false);
  const search = ref('');

  // Helper function to create DataTableProps from event
  const createProps = (event: DataTablePageEvent | DataTableSortEvent | DataTableFilterEvent): DataTableProps => ({
    first: event.first,
    rows: event.rows,
    sortField: event.sortField,
    sortOrder: event.sortOrder ?? 0,
    filters: event.filters,
  });

  // Event handlers
  const onPage = (event: DataTablePageEvent): void => {
    const props = createProps(event);
    listFunction(props);
  };

  const onSort = (event: DataTableSortEvent): void => {
    const props = createProps(event);
    listFunction(props);
  };

  const onFilter = (event: DataTableFilterEvent): void => {
    const props = createProps(event);
    listFunction(props);
  };

  const onSearch = (): void => {
    // Reset to first page when searching
    const props: DataTableProps = {
      first: 0,
      rows: properties.value.rows ?? defaultRows,
      sortField: properties.value.sortField,
      sortOrder: properties.value.sortOrder ?? defaultSortOrder,
      filters: properties.value.filters,
    };
    listFunction(props);
  };

  // Utility functions
  const resetToFirstPage = (): void => {
    const props: DataTableProps = {
      first: 0,
      rows: properties.value.rows ?? defaultRows,
      sortField: properties.value.sortField ?? defaultSortField,
      sortOrder: properties.value.sortOrder ?? defaultSortOrder,
      filters: properties.value.filters,
    };
    listFunction(props);
  };

  const getCurrentPage = (): number => {
    return Math.floor((properties.value.first ?? 0) / (properties.value.rows ?? defaultRows)) + 1;
  };

  const getSortConfig = (): Array<{ key: string; order: 'asc' | 'desc' }> => {
    if (!properties.value.sortField) {
      return [{ key: defaultSortField, order: 'desc' }];
    }

    return [{
      key: String(properties.value.sortField),
      order: properties.value.sortOrder === 1 ? 'asc' : 'desc'
    }];
  };

  return {
    properties,
    loading,
    search,
    handlers: {
      onPage,
      onSort,
      onFilter,
      onSearch
    },
    utils: {
      resetToFirstPage,
      getCurrentPage,
      getSortConfig
    }
  };
}
