<script setup lang="ts">
import Signup from '@/views/Signup.vue';
import Signin from '@/views/Signin.vue';
</script>

<template>
  <v-layout class="rounded rounded-md">
    <v-main class="d-flex" justify="center" align="center" style="min-height: 300px;">
      <v-row>
        <v-col cols="4">
          <Signin></Signin>
        </v-col>
        <v-col cols="4">
          <Signup></Signup>
        </v-col>
      </v-row>
    </v-main>
  </v-layout>
</template>
