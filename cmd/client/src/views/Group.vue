<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { FilterMatchMode, FilterOperator } from '@primevue/core/api';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import MultiSelect from 'primevue/multiselect';
import Select from 'primevue/select';
import Tag from 'primevue/tag';

const customers = ref();
const selectedCustomer = ref();
const filters = ref(
	{
		global: { value: null, matchMode: FilterMatchMode.CONTAINS },
		name: { operator: FilterOperator.AND, constraints: [{ value: null, matchMode: FilterMatchMode.STARTS_WITH }] },
		'country.name': { operator: FilterOperator.AND, constraints: [{ value: null, matchMode: FilterMatchMode.STARTS_WITH }] },
		representative: { value: null, matchMode: FilterMatchMode.IN },
		status: { operator: FilterOperator.OR, constraints: [{ value: null, matchMode: FilterMatchMode.EQUALS }] }
	}
);
const representatives = ref([
	{ name: 'Amy Elsner', image: 'amyelsner.png' },
	{ name: 'Anna Fali', image: 'annafali.png' },
	{ name: 'Asiya Javayant', image: 'asiyajavayant.png' },
	{ name: 'Bernardo Dominic', image: 'bernardodominic.png' },
	{ name: 'Elwin Sharvill', image: 'elwinsharvill.png' },
	{ name: 'Ioni Bowcher', image: 'ionibowcher.png' },
	{ name: 'Ivan Magalhaes', image: 'ivanmagalhaes.png' },
	{ name: 'Onyama Limba', image: 'onyamalimba.png' },
	{ name: 'Stephen Shaw', image: 'stephenshaw.png' },
	{ name: 'XuXue Feng', image: 'xuxuefeng.png' }
]);
const statuses = ref(['unqualified', 'qualified', 'new', 'negotiation', 'renewal', 'proposal']);

const search = ref('');

customers.value = [
	{
		id: '1000',
		name: 'Bamboo Watch',
		country: { name: 'USA', code: 'US' },
		status: 'unqualified',
		representative: representatives.value[0]
	},
	{
		id: '1001',
		name: 'Black Watch',
		country: { name: 'USA', code: 'US' },
		status: 'qualified',
		representative: representatives.value[1]
	},
	{
		id: '1002',
		name: 'Blue Band',
		country: { name: 'France', code: 'FR' },
		status: 'new',
		representative: representatives.value[2]
	},
	{
		id: '1003',
		name: 'Brown Purse',
		country: { name: 'UK', code: 'GB' },
		status: 'negotiation',
		representative: representatives.value[3]
	},
];

onMounted(() => {
	// CustomerService.getCustomersSmall().then((data) => (customers.value = data));
});

const getSeverity = (status: string) => {
	switch (status) {
		case 'unqualified':
			return 'danger';

		case 'qualified':
			return 'success';

		case 'new':
			return 'info';

		case 'negotiation':
			return 'warn';

		case 'renewal':
			return null;
	}
};

// [Info, Roles, User]
</script>

<template>
	<div class="min-h-screen flex flex-col relative flex-auto">
		<div
			class="flex justify-between items-center px-8 py-4 bg-surface-0 dark:bg-surface-950 relative lg:static border-b border-surface gap-8">
			<div class="flex flex-1 gap-4">
				<a v-styleclass="{
					selector: '#app-sidebar-9',
					enterFromClass: 'hidden',
					enterActiveClass: 'animate-fadeinleft',
					leaveToClass: 'hidden',
					leaveActiveClass: 'animate-fadeoutleft',
					hideOnOutsideClick: true
				}" class="cursor-pointer flex items-center justify-center lg:hidden text-surface-700 dark:text-surface-100">
					<i class="pi pi-bars !text-xl" />
				</a>
				<IconField icon-position="left">
					<InputIcon class="pi pi-search text-surface-700 dark:text-surface-100" />
					<InputText type="text" class="border-0 w-full sm:w-80" placeholder="Search" />
				</IconField>
			</div>
		</div>
		<div class="p-8 flex flex-col flex-auto">
			<DataTable v-model:filters="filters" v-model:selection="selectedCustomer" :value="customers"
				stateStorage="session" stateKey="dt-state-demo-session" paginator :rows="5" filterDisplay="menu"
				selectionMode="single" dataKey="id"
				:globalFilterFields="['name', 'country.name', 'representative.name', 'status']"
				tableStyle="min-width: 50rem">
				<template #header>
					<IconField>
						<InputIcon>
							<i class="pi pi-search" />
						</InputIcon>
						<InputText v-model="filters['global'].value" placeholder="Search user" />
					</IconField>
				</template>
				<Column field="name" header="Name" sortable style="width: 25%">
					<template #filter="{ filterModel }">
						<InputText v-model="filterModel.value" type="text" placeholder="Search by name" />
					</template>
				</Column>
				<Column header="Country" sortable sortField="country.name" filterField="country.name"
					filterMatchMode="contains" style="width: 25%">
					<template #body="{ data }">
						<div class="flex items-center gap-2">
							<img alt="flag" src="https://primefaces.org/cdn/primevue/images/flag/flag_placeholder.png"
								:class="`flag flag-${data.country.code}`" style="width: 24px" />
							<span>{{ data.country.name }}</span>
						</div>
					</template>
					<template #filter="{ filterModel }">
						<InputText v-model="filterModel.value" type="text" placeholder="Search by country" />
					</template>
				</Column>
				<Column header="Representative" sortable sortField="representative.name" filterField="representative"
					:showFilterMatchModes="false" :filterMenuStyle="{ width: '14rem' }" style="width: 25%">
					<template #body="{ data }">
						<div class="flex items-center gap-2">
							<img :alt="data.representative.name"
								:src="`https://primefaces.org/cdn/primevue/images/avatar/${data.representative.image}`"
								style="width: 32px" />
							<span>{{ data.representative.name }}</span>
						</div>
					</template>
					<template #filter="{ filterModel }">
						<MultiSelect v-model="filterModel.value" :options="representatives" optionLabel="name"
							placeholder="Any">
							<template #option="slotProps">
								<div class="flex items-center gap-2">
									<img :alt="slotProps.option.name"
										:src="`https://primefaces.org/cdn/primevue/images/avatar/${slotProps.option.image}`"
										style="width: 32px" />
									<span>{{ slotProps.option.name }}</span>
								</div>
							</template>
						</MultiSelect>
					</template>
				</Column>
				<Column field="status" header="Status" sortable filterMatchMode="equals" style="width: 25%">
					<template #body="{ data }">
						<Tag :value="data.status" :severity="getSeverity(data.status)" />
					</template>
					<template #filter="{ filterModel }">
						<Select v-model="filterModel.value" :options="statuses" placeholder="Select One" showClear>
							<template #option="slotProps">
								<Tag :value="slotProps.option" :severity="getSeverity(slotProps.option)" />
							</template>
						</Select>
					</template>
				</Column>
				<template #empty> No customers found. </template>
			</DataTable>
		</div>
	</div>
</template>
<style scoped></style>
