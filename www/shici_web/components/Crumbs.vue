<template>
    <div class="crumbs">
        <el-breadcrumb separator="/">
          <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
          <el-breadcrumb-item :to="{path: $store.state.dynastyMenu.linkPath}">{{$store.state.dynastyMenu.title}}</el-breadcrumb-item>
          <el-breadcrumb-item
          v-if="$store.state.poeterMenu"
          :to="{path:$store.state.poeterMenu.linkPath}"
          >{{$store.state.poeterMenu.title}}</el-breadcrumb-item>

        <el-breadcrumb-item
        v-if="$store.state.poemMenu"
          :to="{path:$store.state.poemMenu.linkPath}"
          >{{$store.state.poemMenu.title}}</el-breadcrumb-item>

        </el-breadcrumb>


    </div>
</template>
