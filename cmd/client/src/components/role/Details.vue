<script setup lang="ts">
/*
Fetch all commands
Fetch all resources
-> Display table

Fetch all permissions of role
 */
</script>
<template v-slot:expanded-row="{ columns, item, props: itemProps }">
	<tr>
		<td :colspan="columns.length">
			<v-card outlined class="main-color-background">
				<v-table>
					<thead>
						<tr>
							<th>
							</th>
							<th class="text-left">
								Read
							</th>
							<th class="text-left">
								Create
							</th>
							<th class="text-left">
								Update
							</th>
							<th class="text-left">
								Delete
							</th>
						</tr>
					</thead>
					<tbody>
						<tr>
							<td>
								asset
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
						</tr>
						<tr>
							<td>
								entity
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
						</tr>
						<tr>
							<td>
								operation
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
						</tr>
						<tr>
							<td>
								role
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
						</tr>
						<tr>
							<td>
								user
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
							<td>
								<v-checkbox></v-checkbox>
							</td>
						</tr>
					</tbody>
				</v-table>
				<v-card-actions>
					<v-spacer></v-spacer>
					<v-btn color="primary" variant="text" @click="updatePermissions">
						Edit
					</v-btn>
					<v-btn color="error" variant="text" @click="delete_">
						Delete
					</v-btn>
				</v-card-actions>
			</v-card>
		</td>
	</tr>
</template>
