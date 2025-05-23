<script setup>
import '../css/test.css'

</script>

<template>
  <header>
    <a href="#" class="logo">秋秋商城</a>
    <nav class="navbar">
      <a href="#">主页</a>
      <a href="#">平台简介</a>
      <a href="#">秒杀商品</a>
      <a href="#">用户评价</a>
      <a href="#">关于我们</a>
    </nav>

    <div class="icons">
      <a href="#" class="fas fa-heart"></a>
      <a href="#" class="fas fa-shopping-cart"></a>
      <a href="#">
        <img src="../images/Sy/qiuqiu.jpg"  class="qiu" alt="picture">
      </a>
    </div>
  </header>

  <!-- header section ends -->

  <!-- 商品详细页面-->
  <div class="container">
 <div class="box-container">
      <div class="left-box" >
      <img src="../images/Sy/test.png"  alt="picture">
      </div>
      <div class="right-box">
        <div class="radio-inputs">
          <label class="radio">
            <input type="radio" name="radio" checked="">
            <span class="name">首页/商品/手机</span>
          </label>
        </div>
      <h1>Ipad 2024款</h1>
      <div class="price-box">
        <div class="price-front">现价:998.00</div>
        <div class="price-behind">
          <div>原价:</div><div class="behind-line-price">￥999.99</div>
        </div>

      </div>
        <div class="detail">
          iPad Pro 2021，是苹果公司推出的一款平板电脑，北京时间2021年4月21日凌晨，iPad Pro在2021苹果春季新品发布会上正式发布。iPad Pro 2021分为11英寸和12.9英寸两种版本，颜色有深空灰色和银色两种颜色可选，搭载最新的M1芯片，后置一颗1200万像素广角摄像头和一颗1000万像素超广角摄像头，前置采用一颗1200万像素深感摄像头，支持人物居中和面容ID。
        </div>
        <div><button id="bottone5">立即购买</button></div>
      </div>

    </div>
    <div class="downBox">
      <div>规格选择:</div>
      <!-- 选择开始-->
      <div class="wrapper">
        <div class="card">
          <input class="input" type="radio" name="card" value="basic">
          <span class="check"></span>
          <label class="label">
            <div class="title">基础款</div>
            <div class="price">
              <span class="span">￥</span>
              998
            </div>
          </label>
        </div>
        <div class="card">
          <input class="input" type="radio" name="card" value="standart">
          <span class="check"></span>
          <label class="label">
            <div class="title">高配款</div>
            <div class="price">
              <span class="span">￥</span>
              1299
            </div>
          </label>
        </div>
        <div class="card">
          <input class="input" type="radio" name="card" value="premium">
          <span class="check"></span>
          <label class="label">
            <div class="title">豪华版</div>
            <div class="price">
              <span class="span">￥</span>
              999999
            </div>
          </label>
        </div>
      </div>
      <!--选择结束-->
    </div>

  </div>
  <!--商品详细页面末尾-->




  <!-- 末尾 starts  -->
  <section class="indexend">
    <div class="relative isolate flex items-center gap-x-6 overflow-hidden bg-gray-50 px-6 py-2.5 sm:px-3.5 sm:before:flex-1">
      <div class="absolute top-1/2 left-[max(-7rem,calc(50%-52rem))] -z-10 -translate-y-1/2 transform-gpu blur-2xl" aria-hidden="true">
        <div class="aspect-577/310 w-144.25 bg-linear-to-r from-[#ff80b5] to-[#9089fc] opacity-30" style="clip-path: polygon(74.8% 41.9%, 97.2% 73.2%, 100% 34.9%, 92.5% 0.4%, 87.5% 0%, 75% 28.6%, 58.5% 54.6%, 50.1% 56.8%, 46.9% 44%, 48.3% 17.4%, 24.7% 53.9%, 0% 27.9%, 11.9% 74.2%, 24.9% 54.1%, 68.6% 100%, 74.8% 41.9%)"></div>
      </div>
      <div class="absolute top-1/2 left-[max(45rem,calc(50%+8rem))] -z-10 -translate-y-1/2 transform-gpu blur-2xl" aria-hidden="true">
        <div class="aspect-577/310 w-144.25 bg-linear-to-r from-[#ff80b5] to-[#9089fc] opacity-30" style="clip-path: polygon(74.8% 41.9%, 97.2% 73.2%, 100% 34.9%, 92.5% 0.4%, 87.5% 0%, 75% 28.6%, 58.5% 54.6%, 50.1% 56.8%, 46.9% 44%, 48.3% 17.4%, 24.7% 53.9%, 0% 27.9%, 11.9% 74.2%, 24.9% 54.1%, 68.6% 100%, 74.8% 41.9%)"></div>
      </div>
      <div class="flex flex-wrap items-center gap-x-4 gap-y-2">
        <p class="text-sm/6 text-gray-900">
          <strong class="font-semibold">秋秋商场 2025</strong><svg viewBox="0 0 2 2" class="mx-2 inline size-0.5 fill-current" aria-hidden="true"><circle cx="1" cy="1" r="1" /></svg>在5-7月加入我们 会有季节优惠卡！
        </p>
        <a href="#" class="flex-none rounded-full bg-gray-900 px-3.5 py-1 text-sm font-semibold text-white shadow-xs hover:bg-gray-700 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gray-900">Register now <span aria-hidden="true">&rarr;</span></a>
      </div>
      <div class="flex flex-1 justify-end">
        <button type="button" class="-m-3 p-3 focus-visible:-outline-offset-4">
          <span class="sr-only">Dismiss</span>
          <svg class="size-5 text-gray-900" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true" data-slot="icon">
            <path d="M6.28 5.22a.75.75 0 0 0-1.06 1.06L8.94 10l-3.72 3.72a.75.75 0 1 0 1.06 1.06L10 11.06l3.72 3.72a.75.75 0 1 0 1.06-1.06L11.06 10l3.72-3.72a.75.75 0 0 0-1.06-1.06L10 8.94 6.28 5.22Z" />
          </svg>
        </button>
      </div>
    </div>
  </section>

  <!-- 末尾 ends -->
</template>

<style scoped>
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css');
</style>