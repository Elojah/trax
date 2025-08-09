<script setup lang="ts">
import { toRefs, ref, computed, onMounted } from "vue";
import { useAuthStore } from '@/stores/auth';
import { useErrorsStore } from '@/stores/errors';

import Button from 'primevue/button';
import FileUpload from 'primevue/fileupload';
import InputText from 'primevue/inputtext';
import { Form, type FormSubmitEvent } from "@primevue/forms";
import FormField from '@primevue/forms/formfield';
import Message from 'primevue/message';

import { zodResolver } from '@primevue/forms/resolvers/zod';
import { z } from 'zod';

import { logger } from "@/config";

const authStore = useAuthStore();
const errorsStore = useErrorsStore();

const {
	message,
	success,
	error,
} = toRefs(errorsStore);

// Form validation schema
const resolver = zodResolver(
	z.object({
		firstName: z.string().min(1, { message: 'First name is required.' }),
		lastName: z.string().min(1, { message: 'Last name is required.' }),
		avatarURL: z.string().optional(),
	})
);

// Password change validation schema
const passwordResolver = zodResolver(
	z.object({
		currentPassword: z.string().min(1, { message: 'Current password is required.' }),
		newPassword: z.string().min(8, { message: 'Password must be at least 8 characters long.' }),
		confirmPassword: z.string().min(1, { message: 'Password confirmation is required.' }),
	}).refine((data) => data.newPassword === data.confirmPassword, {
		message: "Passwords don't match",
		path: ["confirmPassword"],
	})
);

// Initialize form values with user data
const initialValues = computed(() => ({
	firstName: authStore.user?.firstName || '',
	lastName: authStore.user?.lastName || '',
	avatarURL: authStore.user?.avatarURL || '',
}));

const passwordInitialValues = {
	currentPassword: '',
	newPassword: '',
	confirmPassword: '',
};

// Loading state
const loading = ref(false);
const passwordLoading = ref(false);

// Upload avatar placeholder URL
const avatarPlaceholder = "https://fqjltiegiezfetthbags.supabase.co/storage/v1/object/public/block.images/blocks/formlayout/form-avatar.jpg";

// Update profile function
const updateProfile = async function (e: FormSubmitEvent) {
	if (!e.valid) {
		logger.error('Profile form is invalid', e);
		return;
	}

	loading.value = true;
	try {
		await authStore.updateUser(
			e.values.firstName,
			e.values.lastName,
			e.values.avatarURL
		);

		message.value = 'Profile updated successfully!';
		success.value = true;
	} catch (err: any) {
		logger.error('Failed to update profile', err);
		errorsStore.showGRPC(err);
	}
	loading.value = false;
};

// Change password function (placeholder - needs backend implementation)
const changePassword = async function (e: FormSubmitEvent) {
	if (!e.valid) {
		logger.error('Password form is invalid', e);
		return;
	}

	passwordLoading.value = true;
	try {
		// TODO: Implement password change endpoint
		// This would need a new endpoint in the auth service
		message.value = 'Password change functionality not yet implemented.';
		error.value = true;
	} catch (err: any) {
		logger.error('Failed to change password', err);
		errorsStore.showGRPC(err);
	}
	passwordLoading.value = false;
};

// Handle avatar upload (placeholder)
const handleAvatarUpload = (event: any) => {
	// TODO: Implement avatar upload functionality
	logger.info('Avatar upload triggered', event);
	message.value = 'Avatar upload functionality not yet implemented.';
	error.value = true;
};

// Load user data on mount
onMounted(async () => {
	if (!authStore.user) {
		await authStore.refreshUser();
	}
});
</script>
<template>
	<div class="bg-surface-50 dark:bg-surface-950 px-8 md:px-20 py-12 md:py-20 lg:px-80">
		<div class="flex flex-col gap-2">
			<div class="text-surface-900 dark:text-surface-0 font-medium text-xl">Profile</div>
			<p class="m-0 p-0 text-surface-600 dark:text-surface-200 leading-normal">Modify your information.</p>
		</div>

		<!-- Display success/error messages -->
		<div v-if="success || error" class="mt-4">
			<Message v-if="success" severity="success" size="small">{{ message }}</Message>
			<Message v-if="error" severity="error" size="small">{{ message }}</Message>
		</div>

		<!-- Profile Information Form -->
		<div class="bg-surface-0 dark:bg-surface-900 p-7 shadow rounded-2xl mt-7">
			<Form v-slot="$form" :resolver :initialValues="initialValues" @submit="updateProfile"
				class="flex flex-col gap-8">
				<div class="flex flex-col gap-6">
					<h3 class="text-surface-900 dark:text-surface-0 font-medium text-lg m-0">Personal Information</h3>

					<div class="grid grid-cols-12 gap-7">
						<FormField v-slot="$field" name="firstName"
							class="col-span-12 md:col-span-6 flex flex-col gap-2">
							<label for="firstName" class="font-medium text-surface-900 dark:text-surface-0">First
								Name</label>
							<InputText id="firstName" name="firstName" type="text"
								class="w-full rounded-md shadow-sm" />
							<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
								$field.error?.message }}</Message>
						</FormField>

						<FormField v-slot="$field" name="lastName"
							class="col-span-12 md:col-span-6 flex flex-col gap-2">
							<label for="lastName" class="font-medium text-surface-900 dark:text-surface-0">Last
								Name</label>
							<InputText id="lastName" name="lastName" type="text" class="w-full rounded-md shadow-sm" />
							<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
								$field.error?.message }}</Message>
						</FormField>
					</div>

					<div class="grid grid-cols-12 gap-6">
						<div class="col-span-12 md:col-span-6 flex flex-col gap-2">
							<label for="email" class="font-medium text-surface-500 dark:text-surface-400">Email</label>
							<InputText id="email" type="text"
								class="w-full bg-surface-100 dark:bg-surface-800 text-surface-500 dark:text-surface-400"
								:value="authStore.user?.email || ''" readonly disabled />
							<small class="text-surface-500 dark:text-surface-400">Email cannot be changed</small>
						</div>

						<div class="col-span-12 md:col-span-6 flex flex-col gap-2">
							<label for="avatar" class="font-medium text-surface-900 dark:text-surface-0">Avatar</label>
							<div class="flex items-center gap-6">
								<img :src="authStore.user?.avatarURL || avatarPlaceholder"
									class="w-14 h-14 rounded-lg object-cover border border-surface-200 dark:border-surface-700"
									:alt="`${authStore.user?.firstName || 'User'} avatar`" />
								<div class="flex flex-col gap-2">
									<FileUpload mode="basic" name="avatar" accept="image/*" :custom-upload="true" auto
										class="p-button-outlined p-button-secondary" :max-file-size="1000000"
										choose-label="Upload" choose-icon="pi pi-upload"
										@uploader="handleAvatarUpload" />
									<FormField v-slot="$field" name="avatarURL" class="hidden">
										<InputText id="avatarURL" name="avatarURL" type="hidden" />
									</FormField>
								</div>
							</div>
							<small class="text-surface-500 dark:text-surface-400">Maximum file size: 1MB</small>
						</div>
					</div>
				</div>

				<div class="flex justify-start">
					<Button label="Save Changes" type="submit" :loading="loading" icon="pi pi-check" />
				</div>
			</Form>
		</div>

		<!-- Change Password Form -->
		<div class="bg-surface-0 dark:bg-surface-900 p-7 shadow rounded-2xl mt-7">
			<Form v-slot="$passwordForm" :resolver="passwordResolver" :initialValues="passwordInitialValues"
				@submit="changePassword" class="flex flex-col gap-8">
				<div class="flex flex-col gap-6">
					<h3 class="text-surface-900 dark:text-surface-0 font-medium text-lg m-0">Change Password</h3>

					<FormField v-slot="$field" name="currentPassword" class="flex flex-col gap-2">
						<label for="currentPassword" class="font-medium text-surface-900 dark:text-surface-0">Current
							Password</label>
						<InputText id="currentPassword" name="currentPassword" type="password"
							class="w-full rounded-md shadow-sm" />
						<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
							$field.error?.message }}</Message>
					</FormField>

					<div class="grid grid-cols-12 gap-7">
						<FormField v-slot="$field" name="newPassword"
							class="col-span-12 md:col-span-6 flex flex-col gap-2">
							<label for="newPassword" class="font-medium text-surface-900 dark:text-surface-0">New
								Password</label>
							<InputText id="newPassword" name="newPassword" type="password"
								class="w-full rounded-md shadow-sm" />
							<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
								$field.error?.message }}</Message>
						</FormField>

						<FormField v-slot="$field" name="confirmPassword"
							class="col-span-12 md:col-span-6 flex flex-col gap-2">
							<label for="confirmPassword"
								class="font-medium text-surface-900 dark:text-surface-0">Confirm New Password</label>
							<InputText id="confirmPassword" name="confirmPassword" type="password"
								class="w-full rounded-md shadow-sm" />
							<Message v-if="$field?.invalid" severity="error" size="small" variant="simple">{{
								$field.error?.message }}</Message>
						</FormField>
					</div>
				</div>

				<div class="flex justify-start">
					<Button label="Change Password" type="submit" :loading="passwordLoading" icon="pi pi-key"
						severity="secondary" />
				</div>
			</Form>
		</div>
	</div>
</template>
