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
        <span>{{ item.poeter_desc == "NULL" ? "" : item.poeter_desc }}</span>
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
import { getDynastyPoeters, getNavList } from "@/APIService/api/index";

export default {
  components: {
    Crumbs,
  },
  data() {
    return {
      total:0,
      currentPage:1,
      pageSize:50,
      listQuery:{},
      dynasty_id: "",
      poeters: [],

    };
  },
  mounted() {
    if (this.$route.params.id) {
      this.dynasty_id = this.$route.params.id;
      this.ClickgetDynastyPoeters();
    } else {
      this.getNavMenu();
    }
  },
  methods: {
    getNavMenu() {
      getNavList().then((res) => {
        var response = res.data;
        this.$store.commit("updateMenuList", response.datas);
        this.dynasty_id = response.datas[0].dynasty_id;
        this.ClickgetDynastyPoeters();
      });
    },
    ClickgetDynastyPoeters() {
      let params = {
        dynasty_id: this.dynasty_id,
        currentPage:this.currentPage,
        pageSize:this.pageSize,
      };
      getDynastyPoeters(params).then((res) => {
        var response = res.data;
        this.poeters = response.datas;
        this.total = response.totalCount;
        this.currentPage = response.currentPage;
        this.pageSize = response.pageSize;
      });
    },
    ToPoeterDetail(item) {
      console.log(item);
      item["pageType"] = "poeter"; // 诗人
      this.$store.commit("changeCrumbsList", item);
      this.$router.push({
        path: `/poeter/${item.poeter_id}`,
      });
    },

    handleSizeChange(size) {
      this.pageSize = size;
    },
    handleCurrentChange(currentPage) {
      this.currentPage = currentPage;

      let params ={
          dynasty_id: this.dynasty_id ,
          currentPage: this.currentPage,
          pageSize: this.pageSize,
       };
      getDynastyPoeters(params).then((res) => {
        var response = res.data;
        this.poeters = response.datas;
        this.total = response.totalCount;
        this.currentPage = response.currentPage;
        this.pageSize = response.pageSize;
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
.author {
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
.pagebox {
  margin: 20px;
  text-align: right;
}

</style>
