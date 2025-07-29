/**
 * Composable for managing list data and pagination
 * Provides reusable functionality for DataTable-based list components
 */

import { ref, computed, type Ref, type ComputedRef } from 'vue';
import type { DataTableProps } from 'primevue/datatable';

export interface UseListManagerOptions<TRequest, TItem> {
  /** Function to create the list request from DataTable props and search */
  createRequest: (props: DataTableProps, search: string) => TRequest;
  /** Function to call the API list method */
  listMethod: (request: TRequest) => Promise<string[]>;
  /** Map containing all items keyed by ID */
  itemsMap: Ref<Map<string, TItem>>;
  /** Total count of items */
  total: Ref<bigint>;
  /** Default DataTable properties */
  defaultProps?: Partial<DataTableProps>;
}

export interface UseListManagerReturn<TItem> {
  /** Array of item IDs currently being displayed */
  viewIDs: Ref<string[]>;
  /** Computed array of items for the current view */
  views: ComputedRef<(TItem | undefined)[]>;
  /** Loading state */
  loading: Ref<boolean>;
  /** List function to fetch data */
  list: (props?: DataTableProps) => Promise<void>;
}

/**
 * Composable for managing list data with pagination
 * @param options - Configuration options
 */
export function useListManager<TRequest, TItem>(
  options: UseListManagerOptions<TRequest, TItem>
): UseListManagerReturn<TItem> {
  const {
    createRequest,
    listMethod,
    itemsMap,
    total,
    defaultProps = {
      first: 0,
      rows: 10,
      sortField: 'created_at',
      sortOrder: -1,
    }
  } = options;

  // State
  const viewIDs = ref<string[]>([]);
  const loading = ref(false);

  // Computed views based on current IDs
  const views = computed(() => {
    return viewIDs.value.map((id: string) => itemsMap.value?.get(id));
  });

  // List function
  const list = async (props: DataTableProps = defaultProps): Promise<void> => {
    loading.value = true;

    try {
      const request = createRequest(props, ''); // search handled by individual components
      const newIDs = await listMethod(request);
      viewIDs.value = newIDs;
    } catch (error) {
      // Error handling should be done by the calling component
      throw error;
    } finally {
      loading.value = false;
    }
  };

  return {
    viewIDs,
    views,
    loading,
    list
  };
}
