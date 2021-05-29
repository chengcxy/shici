<template>
  <div class="container">
    <!--面包靴导航-->
    <div class="crumbs-bar">
      <Crumbs />
    </div>

    <!--诗人列表-->
    <div class="person-box">
      <div class="title">诗人列表</div>

      <div class="author-item">
        <span class="author alink" style="margin-right:15px">李白</span>
        <span>
          丁仙（一作先）芝，曲阿人。登开元进士第，为馀杭尉。 诗十四首。</span
        >
      </div>
    </div>
  </div>
</template>

<script>
import Crumbs from "../components/Crumbs";
export default {
  components: {
    Crumbs,
  },
  middleware({ store, route, redirect, params, query, req, res }) {

    redirect('/dynasty/') // 默认跳转页面的路由
  },
  data() {
    return {};
  },
};
</script>

<style scoped lang="less">
.crumbs-bar{
  padding: 30px 0;
}

.title{
  font-size: 18px;
    line-height: 24px;

}
.author-item{
  padding: 15px  0;
  overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    color: #444444;
}
</style>



