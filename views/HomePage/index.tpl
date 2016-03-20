<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
  <link rel="stylesheet" href="../../static/css/bootstrap.min.css">
  <link rel="stylesheet" href="../../static/css/jquery.mobile.min.css">
  <!-- <link rel="stylesheet" href="../../static/css/style.css"> -->
  <title>首页</title>
</head>
<body>

<div data-role="header" data-position="fixed" data-tap-toggle="false" style="padding: 0px;margin:0px ">
  <div class="row">
    <div class="col-xs-12">
      <div class="row">
          <div class="col-xs-1" style="padding: 5px 0px 0px 18px;">
            <a href="" class="btn btn-default" >
              <span class="glyphicon glyphicon-qrcode">
              </span>
            </a>
          </div>

           <div class="col-xs-8 " style="padding: 0px;margin:0px 0px 0px 25px;" > 
<!--              <input  type="search" name="search-mini" id="search-mini" value="" data-mini="true"40 placeholder="Search"></input> -->
    <form id="header-search" method="get" action="#" >
      <input  type="search" name="search-mini" id="search-mini" value="" data-mini="true"40 placeholder="Search">
      <button id="btnPost" class="hidden" type="submit" ><span class="glyphicon glyphicon-search"></span>
      </button>
    </form>
           </div>

          <div class="col-xs-1" style="padding: 5px 0px 0px 0px;"><a href="" class="btn btn-default" >
           <span class="glyphicon glyphicon-comment badge" ><sup><b style="font-size:15px;color: red">4</b></sup></span>
           </a>
          </div>
      </div>
    </div>
  </div>
</div>

<div class="carousel slide" id="Home_Carousel" data-ride="carousel">
  <ol class="carousel-indicators">
    <li data-target="#Home_Carousel" data-slide-to="0" class="active"></li>
    <li data-target="#Home_Carousel" data-slide-to="1" ></li>
    <li data-target="#Home_Carousel" data-slide-to="2" ></li>
  </ol>

  <div class="carousel-inner" role="listbox">
    <div class="item active">
      <img src="../../static/img/HomePage/home_carousel/carousel0.jpg" alt="展示一">
    </div>
    <div class="item ">
      <img src="../../static/img/HomePage/home_carousel/carousel1.jpg" alt="展示二">
    </div>
    <div class="item ">
      <img src="../../static/img/HomePage/home_carousel/carousel2.jpg" alt="展示三">
    </div>
  </div>
</div>


<div class="panel panel-default">
  <div class="panel-body">
    <b style="color: red; padding: 0px">最新活动：</b>
    <b class="well" style="padding: 0px">全新口味半价品尝，正在火热进行中......</b>
  </div>
</div>

<div class="row">
  <div class="col-xs-4 col-md-3" style="margin: 0px 0 30 0;padding: 0px">
    <a href="#" class="thumbnail">
      <img src="../../static/img/HomePage/home_imglabel/IMG_000.JPG" alt="..." >
    </a>
  </div>

  <div class="col-xs-4 col-md-3" style="margin: 0px;padding: 0px">
    <a href="#" class="thumbnail">
      <img src="../../static/img/HomePage/home_imglabel/IMG_001.JPG" alt="...">
    </a>
  </div>

  <div class="col-xs-4 col-md-3" style="margin: 0px;padding: 0px">
    <a href="#" class="thumbnail">
      <img src="../../static/img/HomePage/home_imglabel/IMG_002.JPG" alt="...">
    </a>
  </div>

  <div class="col-xs-4 col-md-3" style="margin: 0px;padding: 0px">
    <a href="#" class="thumbnail">
      <img src="../../static/img/HomePage/home_imglabel/IMG_003.JPG" alt="...">
    </a>
  </div>

  <div class="col-xs-4 col-md-3" style="margin: 0px;padding: 0px">
    <a href="#" class="thumbnail">
      <img src="../../static/img/HomePage/home_imglabel/IMG_004.JPG" alt="...">
    </a>
  </div>

    <div class="col-xs-4 col-md-3" style="margin: 0px;padding: 0px">
    <a href="#" class="thumbnail">
      <img src="../../static/img/HomePage/home_imglabel/IMG_005.JPG" alt="...">
    </a>
  </div>
</div>

<div data-role="footer" data-position="fixed" data-tap-toggle="false">
      <div data-role="navbar"data-iconpos="top" >
      <ul>
        <li><a href="#"  data-icon="home" class="ui-btn ui-btn-active"><h5>首页</h5></a></li>
        <li><a href="#"  data-icon="star" class="ui-btn"><h5>商店</h5></a></li>
        <li><a href="#"  data-icon="shop" class="ui-btn"><h5>购物车</h5></a></li>
        <li><a href="#"  data-icon="bullets" class="ui-btn"><h5>订单</h5></a></li>
        <li><a href="#"  data-icon="user" class="ui-btn"><h5>我的</h5></a></li>
      </ul>
    </div>  
</div>
 

    <script src="../../static/js/jquery.min.js"></script>
    <script src="../../static/js/jquery.mobile.min.js"></script>
    <script src="../../static/js/bootstrap.min.js"></script>
</body>
</html>