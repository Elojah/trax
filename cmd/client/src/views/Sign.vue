<script setup lang="ts">
import Signup from '@/components/sign/Signup.vue';
import Signin from '@/components/sign/Signin.vue';
import Menu from '@/components/Menu.vue';
</script>

<template>
	  <div class="common-layout">
    <el-container>
      <el-header>Header</el-header>
      <el-container>
		  <el-aside width="200px">
			<Menu />
		</el-aside>
        <el-main>Main</el-main>
      </el-container>
    </el-container>
  </div>
</template>

<style scoped>
</style>
