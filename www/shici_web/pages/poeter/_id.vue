<template>
  <div class="container">
    <!--面包靴导航-->
    <div class="crumbs-bar">
      <Crumbs />
    </div>
  <h3>诗人作品集</h3>
    <!--根据诗人获取作品列表-->
    <div class="person-box">
      <div class="author-item" v-for="(item, index) in poems" :key="index">
        <span @click="ToPoemDetail(item)" class="author alink" >{{item.poem_name}}</span>
         <span>{{ item.contents }}</span>

      </div>

      <!-- <template>
      <div class="author-item">
        <span class="author alink" style="margin-right:15px">李白</span>
        <span>
          丁仙（一作先）芝，曲阿人。登开元进士第，为馀杭尉。 诗十四首。</span
        >
      </div>
      </template> -->
    </div>
  </div>
</template>

<script>
import Crumbs from "@/components/Crumbs";
import { getPoeterPoems } from "@/APIService/api/index";

export default {
  components: {
    Crumbs,
  },
  data() {
    return {
      poeter_id: "",
      poems: [],
    };
  },
  mounted() {
    this.poeter_id = this.$route.params.id;
    this.ClickgetPoeterPoems();
  },
  methods: {
    ClickgetPoeterPoems() {
       let params = { poeter_id: this.poeter_id };
       getPoeterPoems(params).then((res) => {
        var response = res.data;
        this.poems = response.datas;

      });
    },
    ToPoemDetail(item){
            console.log(item)
            item['pageType'] = 'poem' // 作品
               this.$store.commit('changeCrumbsList',item)
            this.$router.push({
                path:`/poem/${item.poem_id}`
            })
        }


  },
};
</script>

<style scoped lang="less">
.crumbs-bar {
  padding: 30px 0;
}

.title {
  font-size: 18px;
  line-height: 24px;
}
.author{
  display: inline-block;
  width: 130px;
  cursor: pointer;
}
.author-item {
  padding: 15px 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: #444444;
}
</style>
