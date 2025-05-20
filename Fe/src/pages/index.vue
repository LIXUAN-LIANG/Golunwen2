<script setup>
import '../css/style2.css'
</script>

<template>
  <header>
    <input type="checkbox" name="" id="toggler">
    <label for="toggler" class="fas fa-bars"></label>

    <a href="#" class="logo">秋秋商城</a>

    <nav class="navbar">
      <a href="#home">主页</a>
      <a href="#about">平台简介</a>
      <a href="#products">秒杀商品</a>
      <a href="#review">用户评价</a>
      <a href="#contact">关于我们</a>
    </nav>

    <div class="icons">
      <a href="#" class="fas fa-heart"></a>
      <a href="#" class="fas fa-shopping-cart"></a>
      <a href="#">
        <img src="../images/Sy/qiuqiu.jpg"  class="qiu">
      </a>
    </div>
  </header>

  <!-- header section ends -->

  <!-- home section starts  -->

  <section class="home" id="home">

    <div class="content">
      <h3>秋秋抢购商城</h3>
      <span> 琳琅满目的商品</span>
      <p>秒杀狂欢季，抢到就是赚到！全网爆款直降50%，大牌1折起，限时开抢！库存告急，手慢无！每日10点、20点准时开闸，超值好物1元秒杀、限量神券免费领！拼手速更要拼网速，点击即享闪电配送，错过再等一整年！你想要的爆品、低价、惊喜，这里统统有——立即加入购物车，做今夜抢购赢家！</p>
      <a href="#" class="btn">下单直购</a>
    </div>

  </section>

  <!-- home section ends -->

  <!-- about section starts  -->

  <section class="about" id="about">

    <h1 class="heading"> <span> 平台 </span> 选择 </h1>

    <div class="row">

      <div class="video-container">
        <video src="../images/Sy/xlarge_2x.mp4" loop autoplay muted></video>
        <h3>best shopping platform</h3>
      </div>

      <div class="content">
        <h3>为何秋秋商场？</h3>
        <p>极速购，畅快省！XX商城0.1秒闪电加载，下单流畅无卡顿；爆款直降3折起，每日限量秒杀+百元券免费领！智能推荐精准匹配，24小时极速售后，省钱省心更省时！新人注册立享1元特权，满额再赠豪礼——点击即达，好货不等人，抢到就是赚翻天！</p>
        <p>最好的售后  最好的用户体验 只等你来！</p>
        <a href="#" class="btn">了解更多</a>
      </div>

    </div>

  </section>

  <!-- about section ends -->

  <!-- icons section starts  -->

  <section class="icons-container">

    <div class="icons">
      <img src="../images/Sy/icon-1.png" alt="">
      <div class="info">
        <h3>免邮费</h3>
        <span>无论何时何地</span>
      </div>
    </div>

    <div class="icons">
      <img src="../images/Sy/icon-2.png" alt="">
      <div class="info">
        <h3>七天包退换</h3>
        <span>退货退款有保障</span>
      </div>
    </div>

    <div class="icons">
      <img src="../images/Sy/icon-3.png" alt="">
      <div class="info">
        <h3>促销 & 赠品</h3>
        <span>超多选择</span>
      </div>
    </div>

    <div class="icons">
      <img src="../images/Sy/icon-4.png" alt="">
      <div class="info">
        <h3>安全支付</h3>
        <span>安全的支付接口</span>
      </div>
    </div>

  </section>

  <!-- icons section ends -->

  <!-- prodcuts section starts  -->

  <section class="products" id="products">

    <h1 class="heading"> 秒杀 <span>商品</span> </h1>

    <div class="box-container">

      <div class="box">
        <span class="discount">-10%</span>
        <div class="image">
          <img src="../images/Sy/img-1.jpg" alt="">
          <div class="icons">
            <a href="#" class="fas fa-heart"></a>
            <a href="#" class="cart-btn">add to cart</a>
            <a href="#" class="fas fa-share"></a>
          </div>
        </div>
        <div class="content">
          <h3>flower pot</h3>
          <div class="price"> $12.99 <span>$15.99</span> </div>
        </div>
      </div>

      <div class="box">
        <span class="discount">-15%</span>
        <div class="image">
          <img src="../images/Sy/img-2.jpg" alt="">
          <div class="icons">
            <a href="#" class="fas fa-heart"></a>
            <a href="#" class="cart-btn">add to cart</a>
            <a href="#" class="fas fa-share"></a>
          </div>
        </div>
        <div class="content">
          <h3>flower pot</h3>
          <div class="price"> $12.99 <span>$15.99</span> </div>
        </div>
      </div>

      <div class="box">
        <span class="discount">-5%</span>
        <div class="image">
          <img src="../images/Sy/img-3.jpg" alt="">
          <div class="icons">
            <a href="#" class="fas fa-heart"></a>
            <a href="#" class="cart-btn">add to cart</a>
            <a href="#" class="fas fa-share"></a>
          </div>
        </div>
        <div class="content">
          <h3>flower pot</h3>
          <div class="price"> $12.99 <span>$15.99</span> </div>
        </div>
      </div>

      <div class="box">
        <span class="discount">-20%</span>
        <div class="image">
          <img src="../images/Sy/img-4.jpg" alt="">
          <div class="icons">
            <a href="#" class="fas fa-heart"></a>
            <a href="#" class="cart-btn">add to cart</a>
            <a href="#" class="fas fa-share"></a>
          </div>
        </div>
        <div class="content">
          <h3>flower pot</h3>
          <div class="price"> $12.99 <span>$15.99</span> </div>
        </div>
      </div>

      <div class="box">
        <span class="discount">-17%</span>
        <div class="image">
          <img src="../images/Sy/img-5.jpg" alt="">
          <div class="icons">
            <a href="#" class="fas fa-heart"></a>
            <a href="#" class="cart-btn">add to cart</a>
            <a href="#" class="fas fa-share"></a>
          </div>
        </div>
        <div class="content">
          <h3>flower pot</h3>
          <div class="price"> $12.99 <span>$15.99</span> </div>
        </div>
      </div>

      <div class="box">
        <span class="discount">-3%</span>
        <div class="image">
          <img src="../images/Sy/img-6.jpg" alt="">
          <div class="icons">
            <a href="#" class="fas fa-heart"></a>
            <a href="#" class="cart-btn">add to cart</a>
            <a href="#" class="fas fa-share"></a>
          </div>
        </div>
        <div class="content">
          <h3>flower pot</h3>
          <div class="price"> $12.99 <span>$15.99</span> </div>
        </div>
      </div>

      <div class="box">
        <span class="discount">-18%</span>
        <div class="image">
          <img src="../images/Sy/img-7.jpg" alt="">
          <div class="icons">
            <a href="#" class="fas fa-heart"></a>
            <a href="#" class="cart-btn">add to cart</a>
            <a href="#" class="fas fa-share"></a>
          </div>
        </div>
        <div class="content">
          <h3>flower pot</h3>
          <div class="price"> $12.99 <span>$15.99</span> </div>
        </div>
      </div>

      <div class="box">
        <span class="discount">-10%</span>
        <div class="image">
          <img src="../images/Sy/img-8.jpg" alt="">
          <div class="icons">
            <a href="#" class="fas fa-heart"></a>
            <a href="#" class="cart-btn">add to cart</a>
            <a href="#" class="fas fa-share"></a>
          </div>
        </div>
        <div class="content">
          <h3>flower pot</h3>
          <div class="price"> $12.99 <span>$15.99</span> </div>
        </div>
      </div>

      <div class="box">
        <span class="discount">-5%</span>
        <div class="image">
          <img src="../images/Sy/img-9.jpg" alt="">
          <div class="icons">
            <a href="#" class="fas fa-heart"></a>
            <a href="#" class="cart-btn">add to cart</a>
            <a href="#" class="fas fa-share"></a>
          </div>
        </div>
        <div class="content">
          <h3>flower pot</h3>
          <div class="price"> $12.99 <span>$15.99</span> </div>
        </div>
      </div>

    </div>

  </section>

  <!-- prodcuts section ends -->

  <!-- review section starts  -->

  <section class="review" id="review">

    <h1 class="heading"> 顾客 <span>评论</span> </h1>

    <div class="box-container">

      <div class="box">
        <div class="stars">
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
        </div>
        <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti asperiores laboriosam praesentium enim maiores? Ad repellat voluptates alias facere repudiandae dolor accusamus enim ut odit, aliquam nesciunt eaque nulla dignissimos.</p>
        <div class="user">
          <img src="../images/Sy/pic-1.png" alt="">
          <div class="user-info">
            <h3>john deo</h3>
            <span>happy customer</span>
          </div>
        </div>
        <span class="fas fa-quote-right"></span>
      </div>

      <div class="box">
        <div class="stars">
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
        </div>
        <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti asperiores laboriosam praesentium enim maiores? Ad repellat voluptates alias facere repudiandae dolor accusamus enim ut odit, aliquam nesciunt eaque nulla dignissimos.</p>
        <div class="user">
          <img src="../images/Sy/pic-2.png" alt="">
          <div class="user-info">
            <h3>john deo</h3>
            <span>happy customer</span>
          </div>
        </div>
        <span class="fas fa-quote-right"></span>
      </div>

      <div class="box">
        <div class="stars">
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
          <i class="fas fa-star"></i>
        </div>
        <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Corrupti asperiores laboriosam praesentium enim maiores? Ad repellat voluptates alias facere repudiandae dolor accusamus enim ut odit, aliquam nesciunt eaque nulla dignissimos.</p>
        <div class="user">
          <img src="../images/Sy/pic-3.png" alt="">
          <div class="user-info">
            <h3>john deo</h3>
            <span>happy customer</span>
          </div>
        </div>
        <span class="fas fa-quote-right"></span>
      </div>

    </div>

  </section>

  <!-- review section ends -->

  <!-- contact section starts  -->

  <section class="contact" id="contact">

    <h1 class="heading"> <span> contact </span> us </h1>

    <div class="row">

      <form action="">
        <input type="text" placeholder="name" class="box">
        <input type="email" placeholder="email" class="box">
        <input type="number" placeholder="number" class="box">
        <textarea name="" class="box" placeholder="message" id="" cols="30" rows="10"></textarea>
        <input type="submit" value="send message" class="btn">
      </form>

      <div class="image">
        <img src="../images/Sy/contact-img.svg" alt="">
      </div>

    </div>

  </section>

  <!-- contact section ends -->

  <!-- footer section starts  -->

  <section class="footer">

    <div class="box-container">

      <div class="box">
        <h3>quick links</h3>
        <a href="#">home</a>
        <a href="#">about</a>
        <a href="#">products</a>
        <a href="#">review</a>
        <a href="#">contact</a>
      </div>

      <div class="box">
        <h3>extra links</h3>
        <a href="#">my account</a>
        <a href="#">my order</a>
        <a href="#">my favorite</a>
      </div>

      <div class="box">
        <h3>locations</h3>
        <a href="#">india</a>
        <a href="#">USA</a>
        <a href="#">japan</a>
        <a href="#">france</a>
      </div>

      <div class="box">
        <h3>contact info</h3>
        <a href="#">+123-456-7890</a>
        <a href="#">example@gmail.com</a>
        <a href="#">mumbai, india - 400104</a>
        <img src="../images/Sy/payment.png" alt="">
      </div>

    </div>

    <div class="credit"> created by <span> mr. web designer </span> | all rights reserved </div>

  </section>
</template>

<style scoped>
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css');
</style>