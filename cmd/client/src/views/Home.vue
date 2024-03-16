<script setup lang="ts">
import Signup from '@/components/Signup.vue';
import Signin from '@/components/Signin.vue';
</script>

<template>
  <v-main>
    <v-container fluid>
      <v-row>
        <v-col class="mx-auto px-12 py-12" cols="12" md="6">
          <v-sheet class="px-1" outlined color="primary" rounded>
            <Signin />
          </v-sheet>
        </v-col>
        <v-divider vertical :thickness="1"></v-divider>
        <v-col class="mx-auto px-12 py-12" cols="12" md="6">
          <v-sheet class="px-1" outlined color="primary" rounded>
            <Signup />
          </v-sheet>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<style scoped>
/* .bg {
  background: url('@/assets/img/home-background.jpg') no-repeat center fixed !important;
} */
</style>
