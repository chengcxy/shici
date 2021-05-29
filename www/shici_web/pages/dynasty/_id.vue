<template>
  <div class="container">
    <!--面包靴导航-->
    <div class="crumbs-bar">
      <Crumbs />
    </div>

    <!--诗人列表页 根据朝代获取诗人列表-->
    <div class="person-box">
      <div class="author-item" v-for="(item, index) in poeters" :key="index">
        <span @click="ToPoeterDetail(item)" class="author alink">{{
          item.poeter_name
        }}</span>
        <span>{{ item.poeter_desc=='NULL'?"":item.poeter_desc }}</span>
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
import { getDynastyPoeters , getNavList } from "@/APIService/api/index";

export default {
  components: {
    Crumbs,
  },
  data() {
    return {
      dynasty_id: "",
      poeters: [],
    };
  },
  mounted() {
   if(this.$route.params.id){
      this.dynasty_id = this.$route.params.id;
      this.ClickgetDynastyPoeters();
   }else{
     this.getNavMenu()
   }


  },
  methods: {
        getNavMenu() {
            getNavList().then(res=>{
              var response = res.data
              this.$store.commit('updateMenuList',response.datas)

              this.dynasty_id = response.datas[0].dynasty_id
              this.ClickgetDynastyPoeters();
            })
        },
    ClickgetDynastyPoeters() {

       let params = { dynasty_id: this.dynasty_id };
       getDynastyPoeters(params).then((res) => {
        var response = res.data;
        this.poeters = response.datas;

      });
    },
    ToPoeterDetail(item){
            console.log(item)
            item['pageType'] = 'poeter' // 诗人
             this.$store.commit('changeCrumbsList',item)
            this.$router.push({
                path:`/poeter/${item.poeter_id}`
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
  width: 100px;
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
