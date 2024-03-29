<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import router from '@/router';
import { ref, toRefs } from 'vue';
import type { VForm } from 'vuetify/components/VForm';

const form = ref<VForm | null>(null);

const authStore = useAuthStore();
const {
	profile: profile,
} = toRefs(authStore);

const valid = ref(null as boolean | null)

const companyRoles = [
	{ name: 'Company A', roles: 'CEO' },
	{ name: 'Company B', roles: 'CTO' },
	{ name: 'Company C', roles: 'CFO' },
	{ name: 'Company D', roles: 'COO' },
];

</script>

<template>
	<v-table height="300px" fixed-header>
		<thead>
			<tr>
				<th class="text-left">
					name
				</th>
				<th class="text-left">
					role
				</th>
			</tr>
		</thead>
		<tbody>
			<tr v-for="cr in companyRoles" :key="cr.name">
				<td>{{ cr.name }}</td>
				<td>{{ cr.roles }}</td>
			</tr>
		</tbody>
	</v-table>
</template>

<style scoped></style>
