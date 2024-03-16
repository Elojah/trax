<script setup lang="ts">
import Signup from '@/views/Signup.vue';
import Signin from '@/views/Signin.vue';
</script>

<template>
  <v-main>
    <v-container>
      <v-row>
        <v-col class="mx-auto px-12 py-12" cols="12" md="6">
          <v-sheet>
            <Signin />
          </v-sheet>
        </v-col>
        <v-divider vertical></v-divider>
        <v-col class="mx-auto px-12 py-12" cols="12" md="6">
          <v-sheet>
            <Signup />
          </v-sheet>
        </v-col>
      </v-row>
    </v-container>
  </v-main></template>
