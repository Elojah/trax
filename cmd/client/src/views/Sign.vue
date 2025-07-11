<script setup lang="ts">
import Signup from '@/components/sign/Signup.vue';
import Signin from '@/components/sign/Signin.vue';


</script>
<template>
    <div class="px-12 py-20 md:px-12 lg:px-20 bg-no-repeat bg-cover" style="
            background:
                linear-gradient(0deg, rgba(0, 0, 0, 0.3) 0%, rgba(0, 0, 0, 0.3) 100%),
                linear-gradient(0deg, var(--p-primary-500) 0%, var(--p-primary-500) 100%),
                url('https://fqjltiegiezfetthbags.supabase.co/storage/v1/object/public/block.images/blocks/signin/signin-feature.jpg') lightgray 50% / cover no-repeat;
            background-blend-mode: normal, overlay, normal;
        ">
        <div class="flex flex-wrap">
            <div
                class="md:flex hidden w-full md:w-6/12 p-16 bg-white/80 dark:bg-surface-950/70 backdrop-blur-sm rounded-l-2xl flex-col gap-10">
                <div class="flex items-center gap-4">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-14 w-14" width="33" height="32" viewBox="0 0 33 32"
                        fill="none">
                        <path fill-rule="evenodd" clip-rule="evenodd"
                            d="M7.09219 2.87829C5.94766 3.67858 4.9127 4.62478 4.01426 5.68992C7.6857 5.34906 12.3501 5.90564 17.7655 8.61335C23.5484 11.5047 28.205 11.6025 31.4458 10.9773C31.1517 10.087 30.7815 9.23135 30.343 8.41791C26.6332 8.80919 21.8772 8.29127 16.3345 5.51998C12.8148 3.76014 9.71221 3.03521 7.09219 2.87829ZM28.1759 5.33332C25.2462 2.06 20.9887 0 16.25 0C14.8584 0 13.5081 0.177686 12.2209 0.511584C13.9643 0.987269 15.8163 1.68319 17.7655 2.65781C21.8236 4.68682 25.3271 5.34013 28.1759 5.33332ZM32.1387 14.1025C28.2235 14.8756 22.817 14.7168 16.3345 11.4755C10.274 8.44527 5.45035 8.48343 2.19712 9.20639C2.0292 9.24367 1.86523 9.28287 1.70522 9.32367C1.2793 10.25 0.939308 11.2241 0.695362 12.2356C0.955909 12.166 1.22514 12.0998 1.50293 12.0381C5.44966 11.161 11.0261 11.1991 17.7655 14.5689C23.8261 17.5991 28.6497 17.561 31.9029 16.838C32.0144 16.8133 32.1242 16.7877 32.2322 16.7613C32.2441 16.509 32.25 16.2552 32.25 16C32.25 15.358 32.2122 14.7248 32.1387 14.1025ZM31.7098 20.1378C27.8326 20.8157 22.5836 20.5555 16.3345 17.431C10.274 14.4008 5.45035 14.439 2.19712 15.1619C1.475 15.3223 0.825392 15.5178 0.252344 15.7241C0.250782 15.8158 0.25 15.9078 0.25 16C0.25 24.8366 7.41344 32 16.25 32C23.6557 32 29.8862 26.9687 31.7098 20.1378Z"
                            class="fill-surface-900 dark:fill-surface-0" />
                    </svg>
                </div>
                <div class="flex flex-col gap-2">
                    <div class="text-surface-900 dark:text-surface-0 text-2xl font-semibold leading-tight">Welcome to
                        Tr@x</div>
                    <div class="text-surface-900/60 dark:text-surface-0/60 leading-normal">Tracks solar production in
                        real-time.</div>
                </div>
                <div class="flex flex-col gap-6">
                    <div class="flex gap-4">
                        <div
                            class="h-9 w-9 bg-surface-900 dark:bg-surface-50 rounded-lg flex items-center justify-center">
                            <i class="pi pi-inbox !text-base !leading-none text-white dark:text-surface-900" />
                        </div>
                        <div class="flex flex-col gap-2 flex-1">
                            <div class="text-surface-900 dark:text-surface-0 font-medium">Forecast</div>
                            <div class="text-surface-900/60 dark:text-surface-0/60 leading-normal">Anticipate solar
                                production and plan on year.
                            </div>
                        </div>
                    </div>
                    <div class="flex gap-4">
                        <div
                            class="h-9 w-9 bg-surface-900 dark:bg-surface-50 rounded-lg flex items-center justify-center">
                            <i class="pi pi-shield !text-base !leading-none text-white dark:text-surface-900" />
                        </div>
                        <div class="flex flex-col gap-2 flex-1">
                            <div class="text-surface-900 dark:text-surface-0 font-medium">Share information safely
                            </div>
                            <div class="text-surface-900/60 dark:text-surface-0/60 leading-normal">Share your insights
                                with who you decided. Only.</div>
                        </div>
                    </div>
                    <div class="flex gap-4">
                        <div
                            class="h-9 w-9 bg-surface-900 dark:bg-surface-50 rounded-lg flex items-center justify-center">
                            <i class="pi pi-globe !text-base !leading-none text-white dark:text-surface-900" />
                        </div>
                        <div class="flex flex-col gap-2 flex-1">
                            <div class="text-surface-900 dark:text-surface-0 font-medium">Data oriented</div>
                            <div class="text-surface-900/60 dark:text-surface-0/60 leading-normal">Organize your data
                                however you want.</div>
                        </div>
                    </div>
                </div>
            </div>
            <Signin />
        </div>
    </div>
</template>

<style scoped></style>
