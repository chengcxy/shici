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

      <!--分页-->
      <div class="pagebox">
        <el-pagination
          background
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="currentPage"
          :page-sizes="[20,50,100]"
          :page-size="pageSize"
          layout="total, sizes, prev, pager, next"
          :total="total"
        ></el-pagination>
      </div>



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
      total:0,
      currentPage: 1,
      pageSize:50,
    };
  },
  mounted() {
    this.poeter_id = this.$route.params.id;
    this.ClickgetPoeterPoems();
  },
  methods: {
    ClickgetPoeterPoems() {
       let params = {
          poeter_id: this.poeter_id,
         currentPage: this.currentPage,
          pageSize: this.pageSize
      };
       getPoeterPoems(params).then((res) => {
        var response = res.data;
        this.poems = response.datas;
        this.total = response.totalCount;
        this.currentPage = response.currentPage;
        this.pageSize = response.pageSize;

      });
    },


    handleSizeChange(size) {
      this.pageSize = size;
    },
    handleCurrentChange(currentPage) {
      this.currentPage = currentPage;

      let params ={
          poeter_id: this.poeter_id ,
          currentPage: this.currentPage,
          pageSize: this.pageSize,
       };
      getPoeterPoems(params).then((res) => {
        var response = res.data;
        this.poeters = response.datas;
        this.total = response.totalCount;
        this.currentPage = response.currentPage;
        this.pageSize = response.pageSize;
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
