<template>
    <div class="footer-box">
        <div class="middle-box">
                <div class="left">
                    中国诗词©2021
                </div>
                <div class="right">
                    中国诗词收录了上至先秦，下至当代共计六万余首诗词，完全免费并开放所有诗词内容和程序源代码，您可自由复制、修改、传播诗词内容和程序源码。如果您发现收录的诗词存在错误，请直接提交您的修改请求。
                </div>
        </div>
    </div>
</template>

<script>
export default {
    data () {
        return {

        }
    }
}
</script>

<style lang="less" scoped>
.footer-box{


    border-top: 1px solid #eee;
    padding: 40px 0;
    position: fixed;
    left: 0;
    right: 0;
    bottom: 0;

    background-color: #fff;
    .middle-box{
        display: flex;
        align-items: center;
        .left{
            width: 25%;
        }

    }
}
</style>
