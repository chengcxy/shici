<template>
  <div class="header-box">
    <div class="middle-box">
      <div class="left alink">中国诗词</div>
      <div class="middle">
        <div class="nav">
          <div
            class="nav-item"
            v-for="(item, index) in $store.state.navMenuList"
            :key="index"
            :class="currentIndex == index ? 'active' : ''"
            @click="JumpToPage(item, index)"
          >
            {{ item.dynasty }}
          </div>
        </div>
      </div>
      <div class="right">
        <el-input placeholder="请输入内容" v-model="keyword">
          <i slot="prefix" class="el-input__icon el-icon-search"></i>
        </el-input>
      </div>
    </div>
  </div>
</template>

<script>
import { getNavList } from "../APIService/api/index";
export default {
  data() {
    return {
      keyword: "",
      currentIndex: "",
    };
  },
  mounted() {
    this.getNavMenu();
  },
  methods: {
    getNavMenu() {
      getNavList().then((res) => {
        var response = res.data;
        this.$store.commit("updateMenuList", response.datas);
      });
    },
    JumpToPage(value, index) {
      this.currentIndex = index;

      value["pageType"] = "dynasty"; // 作品
      this.$store.commit("changeCrumbsList", value);
      this.$router.push({
        path: `/dynasty/${value.dynasty_id}`,
      });
    },
  },
};
</script>

<style lang="less" scoped>
.header-box {
  padding: 40px 0;
  border-bottom: 1px solid #eee;
  .middle-box {
    display: flex;
    align-items: center;
  }
  .left {
    cursor: pointer;
    font-size: 30px;
    width: 15%;
  }
  .middle {
    flex: 1;
    .nav {
      display: flex;
      .nav-item {
        padding: 0 15px;
        font-size: 16px;
        text-align: center;
        cursor: pointer;
      }
    }
    .active {
      color: #0077dd;
    }
  }
  .right {
    width: 20%;
  }
}
</style>
