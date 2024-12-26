<!DOCTYPE html>
<html lang="en">
<head>
 <title>Beego Web App</title>
 <meta http-equiv="Content-Type" content="text/html; charset=utf-8">	
 <script src="https://code.jquery.com/jquery-3.4.1.min.js" integrity="sha256-CSXorXvZcTkaix6Yvo6HppcZGetbYMGWSFlBw8HfCJo=" crossorigin="anonymous"></script>
 <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css">
 <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap-theme.min.css">
 <link rel="stylesheet" href="/static/css/bootstrap.min.css">
 <link rel="stylesheet" href="/static/css/bootstrap-theme.min.css">
 <link href="/static/css/starter-template.css" rel="stylesheet">
</head>  	
<body>
  <div class="navbar navbar-inverse navbar-fixed-top" role="navigation">
    <div class="container">
      <div class="navbar-header">
        <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
          <span class="sr-only">Toggle navigation</span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </button>
        <a class="navbar-brand" href="#">Beego Web App</a>
      </div>
      <div class="collapse navbar-collapse">
        <ul class="nav navbar-nav">
          <li><a href="/">Home</a></li>
          <li class="dropdown">
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">Manage Products <b class="caret"></b></a>
            <ul class="dropdown-menu">
              <li><a href="/manage/add">Add Product</a></li>
              <li><a href="/manage/view">View All</a></li>
            </ul>
          </li>
        </ul>
      </div>s
    </div>
  </div>
