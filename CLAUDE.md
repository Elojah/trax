# CLAUDE.md - TRAX Project Guide

## Project Overview

TRAX is a **trading platform network** with a microservices architecture built in Go for the backend and Vue 3 + Primevue for the frontend.

**Note**: The `feat/primevue` branch is the first reference implementation using Primevue and PrimeBlocks, replacing Vuetify completely.

### Tech Stack

- **Backend**: Go, gRPC, PostgreSQL, Redis
- **Frontend**: Vue 3, TypeScript, Primevue, PrimeBlocks, Pinia, Vite
- **Protocol**: Protocol Buffers (protobuf) for API communication
- **Authentication**: JWT tokens, Google OAuth
- **Build**: Docker, Make

---

## Architecture Patterns

### Backend (Go)

- **Microservices**: admin, api, auth, web_client
- **Domain-Driven Design**: internal/user module with entities, roles, claims
- **Clean Architecture**: cmd → internal → pkg layers
- **Repository Pattern**: PostgreSQL implementations in `*/postgres/` directories

### Frontend (Vue 3)

- **Composition API**: All components use `<script setup lang="ts">`
- **State Management**: Pinia stores for each domain (auth, entity, user, role, errors)
- **Component Architecture**: View → Component → Store pattern

---

## Directory Structure

### Backend (`cmd/`)

```
cmd/
├── admin/     - Admin service (rotate keys, ping)
├── api/       - Main API service (CRUD operations)
├── auth/      - Authentication service (signin/signup)
└── web_client/ - Web client server
```

### Frontend (`cmd/client/`)

```
cmd/client/src/
├── components/  - Reusable Vue components
│   ├── entity/  - Entity-related components
│   ├── user/    - User management components
│   ├── role/    - Role management components
│   └── sign/    - Authentication components
├── stores/      - Pinia state stores
├── views/       - Page-level components
└── router/      - Vue Router configuration
```

### Shared Types (`internal/`)

```
internal/user/
├── entity.proto  - Entity definitions
├── user.proto    - User definitions
├── role.proto    - Role & permission definitions
├── dto/          - Data Transfer Objects
└── service/      - Business logic (renamed from agg/)
```

---

## Coding Standards

### Vue Components

```vue
<script setup lang="ts">
import { useStore } from '@/stores/store';
import { computed, ref, toRefs } from 'vue';

// Props with defaults
const props = withDefaults(defineProps<{
  showAction: boolean;
}>(), {
  showAction: false,
});

// Store usage
const store = useStore();
const { items, loading } = toRefs(store);

// Reactive refs
const search = ref('');
const dialog = ref(false);

// Computed properties
const filteredItems = computed(() =>
  items.value.filter(item => item.name.includes(search.value))
);
</script>

<template>
    <!-- Primevue components -->
</template>
```

### Pinia Stores

```typescript
export const useStore = defineStore('storeName', () => {
  // State
  const items = ref<Map<string, Type>>(new Map())
  const loading = ref(false)

  // Actions
  const create = async (data: CreateReq) => {
    try {
      const resp = await api.create(data, { meta: { token: token.value } })
      items.value.set(ulid(resp.id), resp)
    } catch (err) {
      logger.error(err)
      throw err
    }
  }

  return { items, loading, create }
})
```

### Error Handling

- **Frontend**: Use `errorsStore.showGRPC(e)` for consistent error display
- **gRPC Codes**: Map to user-friendly messages in `stores/errors.ts`
- **Success Messages**: Set `message.value` and `success.value = true`

---

## Data Flow Patterns

### CRUD Operations

1. **Component** triggers action (button click)
2. **Store** makes gRPC call to backend
3. **Backend** processes request, updates database
4. **Frontend** updates local state and shows success/error message
5. **Component** refreshes data list

### Authentication Flow

1. User signs in via `auth` service
2. JWT token stored in `authStore`
3. Token included in all API calls via `meta: { token }`
4. Automatic token refresh every 10 minutes
5. Route guards check authentication status

---

## Common Patterns

### Table Components

```vue
<!-- Standard data table pattern -->
<DataTable
  :value="views"
  :loading="loading"
  :totalRecords="Number(total)"
  :paginator="true"
  @page="list"
>
  <!-- Header definitions -->
  <!-- Column definitions -->
</DataTable>
```

### Dialog Components

```vue
<!-- Standard dialog pattern -->
<Dialog v-model:visible="dialogCreate" modal header="Create" :style="{ width: '50rem' }">
  <template #default>
    <form @submit.prevent="create">
      <!-- Form content -->
    </form>
  </template>
  <template #footer>
    <Button label="Cancel" icon="pi pi-times" text @click="dialogCreate = false" />
    <Button label="Create" icon="pi pi-check" @click="create" />
  </template>
</Dialog>
```

### Store Integration

```vue
<script setup lang="ts">
// Standard store pattern
const store = useStore();
const { items, total, selected } = toRefs(store);

const list = async (options = defaultOptions) => {
  loading.value = true;
  try {
    const newIDs = await store.list(createRequest(options));
    viewIDs.value = newIDs;
  } catch (e) {
    errorsStore.showGRPC(e);
  }
  loading.value = false;
};
</script>
```

---

## UI Guidelines

### Primevue Theme

- **Dark Theme**: Default theme is dark using PrimeOne design system
- **Primary Color**: Used for buttons, icons, highlights
- **Prime Icons**: Use `pi pi-*` icon names
- **Responsive**: Use Primevue Flex utilities and responsive grid system
- **PrimeBlocks**: Pre-built UI blocks for common layouts and components

### Color Classes

```scss
```

### Navigation

- **Sidebar**: Left navigation using Primevue Sidebar component
- **Views**: Full-screen components in `/views/`
- **Components**: Reusable pieces in `/components/` using Primevue components
- **Menu**: Primevue Menu and MenuBar components for navigation

---

## Development Workflow

### Make Commands

```bash
make proto        # Generate protobuf files
make admin        # Build admin service
make api          # Build API service
make auth         # Build auth service
make client       # Build frontend
make npm          # Install npm dependencies
```

### Frontend Development

```bash
cd cmd/client
npm run dev       # Start dev server
npm run build     # Build for production
npm run type-check # TypeScript checking
```

### Database Migrations

- Located in `db/postgres/`
- Timestamped SQL files for schema changes
- Applied via admin service

---

## Key Files Reference

### Configuration

- `cmd/client/vite.config.ts` - Frontend build config
- `cmd/client/package.json` - Dependencies
- `docker-compose.yml` - Local development setup

### Core Components

- `cmd/client/src/App.vue` - Root application
- `cmd/client/src/main.ts` - Application entry point
- `cmd/client/src/router/index.ts` - Route definitions

### Business Logic

- `cmd/client/src/stores/*.ts` - State management
- `internal/user/dto/*.proto` - API definitions
- `internal/user/service/` - Backend business logic

---

## Consistency Rules

### Naming Conventions

- **Files**: PascalCase for Vue components, camelCase for stores
- **Variables**: camelCase for JavaScript/TypeScript
- **Constants**: UPPER_SNAKE_CASE
- **Proto**: snake_case for fields, PascalCase for messages

### Import Organization

```typescript
// External libraries
import { computed, ref } from 'vue';
import { useRouter } from 'vue-router';

// Internal stores
import { useAuthStore } from '@/stores/auth';

// Components
import UserTable from '@/components/user/Table.vue';

// Types
import type { User } from '@internal/user/user';
```

### Component Structure

1. **Script setup** with TypeScript
2. **Template** with Primevue components
3. **Scoped styles** (minimal, prefer Primevue theme classes and PrimeBlocks)

---

## TODO Patterns

When implementing TODOs from README.md:

- Follow existing CRUD patterns in `Table.vue` components using Primevue DataTable
- Use permission-based `v-if` directives for action visibility
- Implement consistent error handling with `errorsStore`
- Maintain responsive design with Primevue Flex utilities and PrimeBlocks
- Follow protobuf-first API design
- Leverage PrimeBlocks for rapid UI development and consistent layouts
