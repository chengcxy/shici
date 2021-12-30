<template>
  <div class="container">
    <!--面包靴导航-->
    <div class="crumbs-bar">
      <Crumbs />
    </div>

    <!--作品详情页面  根据作品id获取作品详情-->
    <div class="person-box">
      <div class="author-item">
        <h1>{{poemItem.poem_name}}</h1>

        <h3 style="margin:15px 0">{{ poemItem.dynasty}} · {{poemItem.poeter_name}}</h3>

        <div >
          <p class="contents" v-for="item,index in contents" :key="index" style="margin:10px 0;color:#444"> {{ item }} </p>
          </div>
      </div>

    </div>
  </div>
</template>

<script>
import Crumbs from "@/components/Crumbs";
import { getPoemDetail } from "@/APIService/api/index";

export default {
  components: {
    Crumbs,
  },
  data() {
    return {
      poem_id: "",
      poemItem:{},
      content:"",
      contents:[]

    };
  },
  mounted() {
    this.poem_id = this.$route.params.id;
    this.ClickgetPoemDetail();
  },
  methods: {
    ClickgetPoemDetail() {
       let params = { poem_id: this.poem_id };
       getPoemDetail(params).then((res) => {
        this.poemItem = res.data;
        this.content = this.poemItem.contents
        this.contents = this.content.split("\n")
      });
    },
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
.author-item {
  padding: 15px 0;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  color: #444444;
}
.contents {
  
}
</style>
